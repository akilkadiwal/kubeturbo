package probe

import (
	"fmt"
	"strconv"

	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/labels"

	"github.com/turbonomic/turbo-go-sdk/pkg/builder"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"

	"github.com/golang/glog"
)

// Pods Getter is such func that gets all the pods match the provided namespace, labels and fiels.
type ServiceGetter func(namespace string, selector labels.Selector) ([]*api.Service, error)

type EndpointGetter func(namespace string, selector labels.Selector) ([]*api.Endpoints, error)

type ServiceProbe struct {
	serviceGetter  ServiceGetter
	endpointGetter EndpointGetter
}

func NewServiceProbe(serviceGetter ServiceGetter, epGetter EndpointGetter) *ServiceProbe {
	return &ServiceProbe{
		serviceGetter:  serviceGetter,
		endpointGetter: epGetter,
	}
}

type VMTServiceGetter struct {
	kubeClient *client.Client
}

func NewVMTServiceGetter(kubeClient *client.Client) *VMTServiceGetter {
	return &VMTServiceGetter{
		kubeClient: kubeClient,
	}
}

// Get service match specified namesapce and label.
func (this *VMTServiceGetter) GetService(namespace string, selector labels.Selector) ([]*api.Service, error) {
	listOption := &api.ListOptions{
		LabelSelector: selector,
	}
	serviceList, err := this.kubeClient.Services(namespace).List(*listOption)
	if err != nil {
		return nil, fmt.Errorf("Error listing services: %s", err)
	}

	var serviceItems []*api.Service
	for _, service := range serviceList.Items {
		s := service
		serviceItems = append(serviceItems, &s)
	}

	glog.V(2).Infof("Discovering Services, now the cluster has " + strconv.Itoa(len(serviceItems)) + " services")

	return serviceItems, nil
}

type VMTEndpointGetter struct {
	kubeClient *client.Client
}

func NewVMTEndpointGetter(kubeClient *client.Client) *VMTEndpointGetter {
	return &VMTEndpointGetter{
		kubeClient: kubeClient,
	}
}

// Get endpoints match specified namesapce and label.
func (this *VMTEndpointGetter) GetEndpoints(namespace string, selector labels.Selector) ([]*api.Endpoints, error) {
	listOption := &api.ListOptions{
		LabelSelector: selector,
	}
	epList, err := this.kubeClient.Endpoints(namespace).List(*listOption)
	if err != nil {
		return nil, fmt.Errorf("Error listing endpoints: %s", err)
	}

	var epItems []*api.Endpoints
	for _, endpoint := range epList.Items {
		ep := endpoint
		epItems = append(epItems, &ep)
	}

	return epItems, nil
}

func (this *ServiceProbe) GetService(namespace string, selector labels.Selector) ([]*api.Service, error) {
	if this.serviceGetter == nil {
		return nil, fmt.Errorf("Service getter is not set")
	}
	return this.serviceGetter(namespace, selector)
}

func (this *ServiceProbe) GetEndpoints(namespace string, selector labels.Selector) ([]*api.Endpoints, error) {
	if this.endpointGetter == nil {
		return nil, fmt.Errorf("Endpoint getter is not set")
	}
	return this.endpointGetter(namespace, selector)
}

// Parse Services inside Kubernetes and build entityDTO as VApp.
func (svcProbe *ServiceProbe) ParseService(serviceList []*api.Service, endpointList []*api.Endpoints) (result []*proto.EntityDTO, err error) {
	// first make a endpoint map, key is endpoints cluster ID; value is endpoint object
	endpointMap := make(map[string]*api.Endpoints)
	for _, endpoint := range endpointList {
		nameWithNamespace := endpoint.Namespace + "/" + endpoint.Name
		endpointMap[nameWithNamespace] = endpoint
	}

	for _, service := range serviceList {
		podClusterIDs := svcProbe.findPodEndpoints(service, endpointMap)
		if len(podClusterIDs) < 1 {
			continue
		}
		glog.V(4).Infof("service %s has the following pod as endpoints %v",
			service.Namespace+"/"+service.Name, podClusterIDs)

		// create the commodities bought map.
		commoditiesBoughtMap, err := svcProbe.getCommoditiesBought(podClusterIDs)
		if err != nil {
			return nil, err
		}

		// build the entityDTO.
		entityDto, err := svcProbe.buildServiceEntityDTO(service, commoditiesBoughtMap)
		if err != nil {
			return nil, err
		}

		result = append(result, entityDto)
	}

	return
}

// For every service, find the pods serve this service.
func (this *ServiceProbe) findPodEndpoints(service *api.Service, endpointMap map[string]*api.Endpoints) []string {
	serviceNameWithNamespace := service.Namespace + "/" + service.Name
	serviceEndpoint := endpointMap[serviceNameWithNamespace]
	if serviceEndpoint == nil {
		return nil
	}
	subsets := serviceEndpoint.Subsets
	var podClusterIDList []string
	for _, endpointSubset := range subsets {
		addresses := endpointSubset.Addresses
		for _, address := range addresses {
			target := address.TargetRef
			if target == nil {
				continue
			}
			podName := target.Name
			podNamespace := target.Namespace
			podClusterID := GetPodClusterID(podNamespace, podName)
			// get the pod name and the service name
			podClusterIDList = append(podClusterIDList, podClusterID)
		}
	}
	return podClusterIDList
}

func (this *ServiceProbe) buildServiceEntityDTO(service *api.Service, commoditiesBoughtMap map[*builder.ProviderDTO][]*proto.CommodityDTO) (*proto.EntityDTO, error) {
	serviceEntityType := proto.EntityDTO_VIRTUAL_APPLICATION
	id := string(service.UID)
	displayName := "vApp-" + service.Namespace + "/" + service.Name + "-" + ClusterID
	entityDTOBuilder := builder.NewEntityDTOBuilder(serviceEntityType, id)

	entityDTOBuilder = entityDTOBuilder.DisplayName(displayName)
	for provider, commodities := range commoditiesBoughtMap {
		entityDTOBuilder.Provider(provider)
		entityDTOBuilder.BuysCommodities(commodities)
	}
	entityDto, err := entityDTOBuilder.Create()
	if err != nil {
		return nil, fmt.Errorf("Failed to build EntityDTO for service %s: %s", id, err)
	}

	glog.V(4).Infof("created a service entityDTO %v", entityDto)
	return entityDto, nil
}

func (this *ServiceProbe) getCommoditiesBought(podIDList []string) (
	map[*builder.ProviderDTO][]*proto.CommodityDTO, error) {
	commoditiesBoughtMap := make(map[*builder.ProviderDTO][]*proto.CommodityDTO)

	for _, podID := range podIDList {
		serviceResourceStat := getServiceResourceStat(podTransactionCountMap, podID)
		turboPodUUID, exist := turboPodUUIDMap[podID]
		if !exist {
			return nil, fmt.Errorf("Cannot build commodityBought based on give pod identifier: %s. "+
				"Failed to find Turbo UUID.", podID)
		}
		// Here it is consisted with the ID when we build the application entityDTO in ApplicationProbe/
		appID := appPrefix + turboPodUUID
		appType := podAppTypeMap[podID]
		// We might want to check here if the appID exist.
		appProvider := builder.CreateProvider(proto.EntityDTO_APPLICATION, appID)
		var commoditiesBoughtFromApp []*proto.CommodityDTO
		transactionCommBought, err := builder.NewCommodityDTOBuilder(proto.CommodityDTO_TRANSACTION).
			Key(appType + "-" + ClusterID).
			Used(serviceResourceStat.transactionBought).
			Create()
		if err != nil {
			return nil, err
		}
		commoditiesBoughtFromApp = append(commoditiesBoughtFromApp, transactionCommBought)

		commoditiesBoughtMap[appProvider] = commoditiesBoughtFromApp

	}
	return commoditiesBoughtMap, nil
}

func getServiceResourceStat(transactionCountMap map[string]float64, podID string) *ServiceResourceStat {
	transactionBought := float64(0)

	count, ok := transactionCountMap[podID]
	if ok {
		transactionBought = count
		glog.V(4).Infof("Transaction bought from pod %s is %d", podID, count)
	} else {
		glog.V(4).Infof("No transaction value for applications on pod %s", podID)
	}

	return &ServiceResourceStat{
		transactionBought: transactionBought,
	}
}
