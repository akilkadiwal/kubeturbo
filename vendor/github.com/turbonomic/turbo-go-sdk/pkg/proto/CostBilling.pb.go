// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: CostBilling.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CloudBillingData_CloudBillingBucket_Granularity int32

const (
	CloudBillingData_CloudBillingBucket_HOURLY  CloudBillingData_CloudBillingBucket_Granularity = 0
	CloudBillingData_CloudBillingBucket_DAILY   CloudBillingData_CloudBillingBucket_Granularity = 1
	CloudBillingData_CloudBillingBucket_MONTHLY CloudBillingData_CloudBillingBucket_Granularity = 2
)

// Enum value maps for CloudBillingData_CloudBillingBucket_Granularity.
var (
	CloudBillingData_CloudBillingBucket_Granularity_name = map[int32]string{
		0: "HOURLY",
		1: "DAILY",
		2: "MONTHLY",
	}
	CloudBillingData_CloudBillingBucket_Granularity_value = map[string]int32{
		"HOURLY":  0,
		"DAILY":   1,
		"MONTHLY": 2,
	}
)

func (x CloudBillingData_CloudBillingBucket_Granularity) Enum() *CloudBillingData_CloudBillingBucket_Granularity {
	p := new(CloudBillingData_CloudBillingBucket_Granularity)
	*p = x
	return p
}

func (x CloudBillingData_CloudBillingBucket_Granularity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CloudBillingData_CloudBillingBucket_Granularity) Descriptor() protoreflect.EnumDescriptor {
	return file_CostBilling_proto_enumTypes[0].Descriptor()
}

func (CloudBillingData_CloudBillingBucket_Granularity) Type() protoreflect.EnumType {
	return &file_CostBilling_proto_enumTypes[0]
}

func (x CloudBillingData_CloudBillingBucket_Granularity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CloudBillingData_CloudBillingBucket_Granularity) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CloudBillingData_CloudBillingBucket_Granularity(num)
	return nil
}

// Deprecated: Use CloudBillingData_CloudBillingBucket_Granularity.Descriptor instead.
func (CloudBillingData_CloudBillingBucket_Granularity) EnumDescriptor() ([]byte, []int) {
	return file_CostBilling_proto_rawDescGZIP(), []int{0, 1, 0}
}

type CloudBillingDataPoint_CostCategory int32

const (
	CloudBillingDataPoint_COMPUTE           CloudBillingDataPoint_CostCategory = 1
	CloudBillingDataPoint_LICENSE           CloudBillingDataPoint_CostCategory = 2
	CloudBillingDataPoint_STORAGE           CloudBillingDataPoint_CostCategory = 3
	CloudBillingDataPoint_COMMITMENT_CHARGE CloudBillingDataPoint_CostCategory = 4
	// Represents the cost of a commitment based on proportional usage by an entity/scope.
	// This is useful in creating a charge-back model. For GCP, COMMITMENT_USAGE will be calculated
	// in the bill for us. It will represent the used portion of a CUD cost. The COMMITMENT_CHARGE
	// cost will be the unused CUD cost charged to the purchasing account. For AWS and Azure,
	// the COMMITMENT_CHARGE cost will always be the full cost of the RI/SP in the purchasing account,
	// while COMMITMENT_USAGE will be an accounting cost (not actually billed), based on where it is
	// used.
	CloudBillingDataPoint_COMMITMENT_USAGE       CloudBillingDataPoint_CostCategory = 5
	CloudBillingDataPoint_COMPUTE_LICENSE_BUNDLE CloudBillingDataPoint_CostCategory = 6
	CloudBillingDataPoint_UNCATEGORIZED          CloudBillingDataPoint_CostCategory = 2047
)

// Enum value maps for CloudBillingDataPoint_CostCategory.
var (
	CloudBillingDataPoint_CostCategory_name = map[int32]string{
		1:    "COMPUTE",
		2:    "LICENSE",
		3:    "STORAGE",
		4:    "COMMITMENT_CHARGE",
		5:    "COMMITMENT_USAGE",
		6:    "COMPUTE_LICENSE_BUNDLE",
		2047: "UNCATEGORIZED",
	}
	CloudBillingDataPoint_CostCategory_value = map[string]int32{
		"COMPUTE":                1,
		"LICENSE":                2,
		"STORAGE":                3,
		"COMMITMENT_CHARGE":      4,
		"COMMITMENT_USAGE":       5,
		"COMPUTE_LICENSE_BUNDLE": 6,
		"UNCATEGORIZED":          2047,
	}
)

func (x CloudBillingDataPoint_CostCategory) Enum() *CloudBillingDataPoint_CostCategory {
	p := new(CloudBillingDataPoint_CostCategory)
	*p = x
	return p
}

func (x CloudBillingDataPoint_CostCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CloudBillingDataPoint_CostCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_CostBilling_proto_enumTypes[1].Descriptor()
}

func (CloudBillingDataPoint_CostCategory) Type() protoreflect.EnumType {
	return &file_CostBilling_proto_enumTypes[1]
}

func (x CloudBillingDataPoint_CostCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CloudBillingDataPoint_CostCategory) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CloudBillingDataPoint_CostCategory(num)
	return nil
}

// Deprecated: Use CloudBillingDataPoint_CostCategory.Descriptor instead.
func (CloudBillingDataPoint_CostCategory) EnumDescriptor() ([]byte, []int) {
	return file_CostBilling_proto_rawDescGZIP(), []int{2, 0}
}

type CloudBillingData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// For AWS, this should be the master account ID
	// For Azure, this should be the EA number
	// For GCP, this should be the billing account ID.
	// Note: this may possible remain a string in the XL platform, as it will
	// only be used as an identifier (not a reference to an entity or group)
	BillingIdentifier *string                                `protobuf:"bytes,1,opt,name=billing_identifier,json=billingIdentifier" json:"billing_identifier,omitempty"`
	CloudCostBuckets  []*CloudBillingData_CloudBillingBucket `protobuf:"bytes,2,rep,name=cloud_cost_buckets,json=cloudCostBuckets" json:"cloud_cost_buckets,omitempty"`
	CostTagGroupMap   map[int64]*CostTagGroup                `protobuf:"bytes,3,rep,name=cost_tag_group_map,json=costTagGroupMap" json:"cost_tag_group_map,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (x *CloudBillingData) Reset() {
	*x = CloudBillingData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CostBilling_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudBillingData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudBillingData) ProtoMessage() {}

func (x *CloudBillingData) ProtoReflect() protoreflect.Message {
	mi := &file_CostBilling_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudBillingData.ProtoReflect.Descriptor instead.
func (*CloudBillingData) Descriptor() ([]byte, []int) {
	return file_CostBilling_proto_rawDescGZIP(), []int{0}
}

func (x *CloudBillingData) GetBillingIdentifier() string {
	if x != nil && x.BillingIdentifier != nil {
		return *x.BillingIdentifier
	}
	return ""
}

func (x *CloudBillingData) GetCloudCostBuckets() []*CloudBillingData_CloudBillingBucket {
	if x != nil {
		return x.CloudCostBuckets
	}
	return nil
}

func (x *CloudBillingData) GetCostTagGroupMap() map[int64]*CostTagGroup {
	if x != nil {
		return x.CostTagGroupMap
	}
	return nil
}

// Provides a single reference per data point to the set of tags
type CostTagGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags map[string]string `protobuf:"bytes,1,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (x *CostTagGroup) Reset() {
	*x = CostTagGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CostBilling_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostTagGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostTagGroup) ProtoMessage() {}

func (x *CostTagGroup) ProtoReflect() protoreflect.Message {
	mi := &file_CostBilling_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostTagGroup.ProtoReflect.Descriptor instead.
func (*CostTagGroup) Descriptor() ([]byte, []int) {
	return file_CostBilling_proto_rawDescGZIP(), []int{1}
}

func (x *CostTagGroup) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type CloudBillingDataPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudServiceId *string `protobuf:"bytes,1,opt,name=cloud_service_id,json=cloudServiceId" json:"cloud_service_id,omitempty"`
	AccountId      *string `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	RegionId       *string `protobuf:"bytes,3,opt,name=region_id,json=regionId" json:"region_id,omitempty"`
	// For VM/DB/Volume ON_DEMAND, this will be associated cloud tier
	// For VM COMMITMENT_DISCOUNT, this will be the commitment ID
	// For entity ID = commitment ID and COMMITMENT_CHARGE, the cloud tier?
	ProviderId   *string                             `protobuf:"bytes,4,opt,name=provider_id,json=providerId" json:"provider_id,omitempty"`
	EntityType   *int32                              `protobuf:"varint,5,opt,name=entity_type,json=entityType" json:"entity_type,omitempty"`
	EntityId     *string                             `protobuf:"bytes,6,opt,name=entity_id,json=entityId" json:"entity_id,omitempty"`
	PriceModel   *PriceModel                         `protobuf:"varint,7,opt,name=price_model,json=priceModel,enum=common_dto.PriceModel" json:"price_model,omitempty"`
	CostCategory *CloudBillingDataPoint_CostCategory `protobuf:"varint,8,opt,name=cost_category,json=costCategory,enum=common_dto.CloudBillingDataPoint_CostCategory" json:"cost_category,omitempty"`
	// The purchase commodity from the provider ID.
	// For GCP VM's, this may be NUM_VCORE or VMEM
	// For an AWS Volume, this may be STORAGE_AMOUNT, STORAGE_ACCESS (IOPS), or IO_THROUGHPUT
	PurchasedCommodity *CommodityDTO_CommodityType `protobuf:"varint,9,opt,name=purchased_commodity,json=purchasedCommodity,enum=common_dto.CommodityDTO_CommodityType" json:"purchased_commodity,omitempty"`
	UsageAmount        *float64                    `protobuf:"fixed64,10,opt,name=usage_amount,json=usageAmount" json:"usage_amount,omitempty"`
	Cost               *CurrencyAmount             `protobuf:"bytes,11,opt,name=cost" json:"cost,omitempty"`
	CostTagGroupId     *int64                      `protobuf:"varint,12,opt,name=cost_tag_group_id,json=costTagGroupId" json:"cost_tag_group_id,omitempty"`
	// In Azure, this would be the meter ID
	// In GCP, this would be the SKU ID
	// In AWS, this would be product/sku (also present in offer sheet)
	// SKU will provide clarity on the rates we discover (debugging or answering customer questions). Does not necessarily need to be persisted
	// in the cost DB.
	SkuId        *string `protobuf:"bytes,13,opt,name=sku_id,json=skuId" json:"sku_id,omitempty"`
	ProviderType *int32  `protobuf:"varint,14,opt,name=provider_type,json=providerType" json:"provider_type,omitempty"`
}

func (x *CloudBillingDataPoint) Reset() {
	*x = CloudBillingDataPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CostBilling_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudBillingDataPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudBillingDataPoint) ProtoMessage() {}

func (x *CloudBillingDataPoint) ProtoReflect() protoreflect.Message {
	mi := &file_CostBilling_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudBillingDataPoint.ProtoReflect.Descriptor instead.
func (*CloudBillingDataPoint) Descriptor() ([]byte, []int) {
	return file_CostBilling_proto_rawDescGZIP(), []int{2}
}

func (x *CloudBillingDataPoint) GetCloudServiceId() string {
	if x != nil && x.CloudServiceId != nil {
		return *x.CloudServiceId
	}
	return ""
}

func (x *CloudBillingDataPoint) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *CloudBillingDataPoint) GetRegionId() string {
	if x != nil && x.RegionId != nil {
		return *x.RegionId
	}
	return ""
}

func (x *CloudBillingDataPoint) GetProviderId() string {
	if x != nil && x.ProviderId != nil {
		return *x.ProviderId
	}
	return ""
}

func (x *CloudBillingDataPoint) GetEntityType() int32 {
	if x != nil && x.EntityType != nil {
		return *x.EntityType
	}
	return 0
}

func (x *CloudBillingDataPoint) GetEntityId() string {
	if x != nil && x.EntityId != nil {
		return *x.EntityId
	}
	return ""
}

func (x *CloudBillingDataPoint) GetPriceModel() PriceModel {
	if x != nil && x.PriceModel != nil {
		return *x.PriceModel
	}
	return PriceModel_ON_DEMAND
}

func (x *CloudBillingDataPoint) GetCostCategory() CloudBillingDataPoint_CostCategory {
	if x != nil && x.CostCategory != nil {
		return *x.CostCategory
	}
	return CloudBillingDataPoint_COMPUTE
}

func (x *CloudBillingDataPoint) GetPurchasedCommodity() CommodityDTO_CommodityType {
	if x != nil && x.PurchasedCommodity != nil {
		return *x.PurchasedCommodity
	}
	return CommodityDTO_CLUSTER
}

func (x *CloudBillingDataPoint) GetUsageAmount() float64 {
	if x != nil && x.UsageAmount != nil {
		return *x.UsageAmount
	}
	return 0
}

func (x *CloudBillingDataPoint) GetCost() *CurrencyAmount {
	if x != nil {
		return x.Cost
	}
	return nil
}

func (x *CloudBillingDataPoint) GetCostTagGroupId() int64 {
	if x != nil && x.CostTagGroupId != nil {
		return *x.CostTagGroupId
	}
	return 0
}

func (x *CloudBillingDataPoint) GetSkuId() string {
	if x != nil && x.SkuId != nil {
		return *x.SkuId
	}
	return ""
}

func (x *CloudBillingDataPoint) GetProviderType() int32 {
	if x != nil && x.ProviderType != nil {
		return *x.ProviderType
	}
	return 0
}

// This structure mirrors CloudCommitmentData
type CloudBillingData_CloudBillingBucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimestampUtcMillis *int64                                           `protobuf:"varint,1,opt,name=timestamp_utc_millis,json=timestampUtcMillis" json:"timestamp_utc_millis,omitempty"`
	Samples            []*CloudBillingDataPoint                         `protobuf:"bytes,2,rep,name=samples" json:"samples,omitempty"`
	Granularity        *CloudBillingData_CloudBillingBucket_Granularity `protobuf:"varint,3,opt,name=granularity,enum=common_dto.CloudBillingData_CloudBillingBucket_Granularity,def=0" json:"granularity,omitempty"`
}

// Default values for CloudBillingData_CloudBillingBucket fields.
const (
	Default_CloudBillingData_CloudBillingBucket_Granularity = CloudBillingData_CloudBillingBucket_HOURLY
)

func (x *CloudBillingData_CloudBillingBucket) Reset() {
	*x = CloudBillingData_CloudBillingBucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CostBilling_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudBillingData_CloudBillingBucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudBillingData_CloudBillingBucket) ProtoMessage() {}

func (x *CloudBillingData_CloudBillingBucket) ProtoReflect() protoreflect.Message {
	mi := &file_CostBilling_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudBillingData_CloudBillingBucket.ProtoReflect.Descriptor instead.
func (*CloudBillingData_CloudBillingBucket) Descriptor() ([]byte, []int) {
	return file_CostBilling_proto_rawDescGZIP(), []int{0, 1}
}

func (x *CloudBillingData_CloudBillingBucket) GetTimestampUtcMillis() int64 {
	if x != nil && x.TimestampUtcMillis != nil {
		return *x.TimestampUtcMillis
	}
	return 0
}

func (x *CloudBillingData_CloudBillingBucket) GetSamples() []*CloudBillingDataPoint {
	if x != nil {
		return x.Samples
	}
	return nil
}

func (x *CloudBillingData_CloudBillingBucket) GetGranularity() CloudBillingData_CloudBillingBucket_Granularity {
	if x != nil && x.Granularity != nil {
		return *x.Granularity
	}
	return Default_CloudBillingData_CloudBillingBucket_Granularity
}

var File_CostBilling_proto protoreflect.FileDescriptor

var file_CostBilling_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x6f, 0x73, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x1a,
	0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x44, 0x54, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x43, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xfe, 0x04, 0x0a, 0x10, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x12, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x5d, 0x0a, 0x12, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f,
	0x63, 0x6f, 0x73, 0x74, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x10, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x5e, 0x0a, 0x12, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x61,
	0x67, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x43, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x63, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4d, 0x61, 0x70, 0x1a, 0x5c, 0x0a, 0x14, 0x43, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x73, 0x74,
	0x54, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x9d, 0x02, 0x0a, 0x12, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x75, 0x74, 0x63, 0x5f, 0x6d, 0x69, 0x6c, 0x6c,
	0x69, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x55, 0x74, 0x63, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x12, 0x3b, 0x0a, 0x07,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x65, 0x0a, 0x0b, 0x67, 0x72, 0x61,
	0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x47, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x3a, 0x06, 0x48, 0x4f, 0x55,
	0x52, 0x4c, 0x59, 0x52, 0x0b, 0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79,
	0x22, 0x31, 0x0a, 0x0b, 0x47, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x0a, 0x0a, 0x06, 0x48, 0x4f, 0x55, 0x52, 0x4c, 0x59, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44,
	0x41, 0x49, 0x4c, 0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x4c,
	0x59, 0x10, 0x02, 0x22, 0x7f, 0x0a, 0x0c, 0x43, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x54, 0x61, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x54,
	0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x92, 0x06, 0x0a, 0x15, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x28,
	0x0a, 0x10, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x52, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x53, 0x0a, 0x0d,
	0x63, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x0c, 0x63, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x57, 0x0a, 0x13, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x69, 0x74, 0x79, 0x44, 0x54, 0x4f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69,
	0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x12, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0b, 0x75, 0x73, 0x61, 0x67, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a,
	0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x29, 0x0a,
	0x11, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x6f, 0x73, 0x74, 0x54, 0x61,
	0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f,
	0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4d, 0x50, 0x55, 0x54, 0x45,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x10, 0x02, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11,
	0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x47,
	0x45, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x10, 0x05, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x4d,
	0x50, 0x55, 0x54, 0x45, 0x5f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x42, 0x55, 0x4e,
	0x44, 0x4c, 0x45, 0x10, 0x06, 0x12, 0x12, 0x0a, 0x0d, 0x55, 0x4e, 0x43, 0x41, 0x54, 0x45, 0x47,
	0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0xff, 0x0f, 0x42, 0x4f, 0x0a, 0x1f, 0x63, 0x6f, 0x6d,
	0x2e, 0x76, 0x6d, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5a, 0x2c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x6e, 0x6f,
	0x6d, 0x69, 0x63, 0x2f, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_CostBilling_proto_rawDescOnce sync.Once
	file_CostBilling_proto_rawDescData = file_CostBilling_proto_rawDesc
)

func file_CostBilling_proto_rawDescGZIP() []byte {
	file_CostBilling_proto_rawDescOnce.Do(func() {
		file_CostBilling_proto_rawDescData = protoimpl.X.CompressGZIP(file_CostBilling_proto_rawDescData)
	})
	return file_CostBilling_proto_rawDescData
}

var file_CostBilling_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_CostBilling_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_CostBilling_proto_goTypes = []interface{}{
	(CloudBillingData_CloudBillingBucket_Granularity)(0), // 0: common_dto.CloudBillingData.CloudBillingBucket.Granularity
	(CloudBillingDataPoint_CostCategory)(0),              // 1: common_dto.CloudBillingDataPoint.CostCategory
	(*CloudBillingData)(nil),                             // 2: common_dto.CloudBillingData
	(*CostTagGroup)(nil),                                 // 3: common_dto.CostTagGroup
	(*CloudBillingDataPoint)(nil),                        // 4: common_dto.CloudBillingDataPoint
	nil,                                                  // 5: common_dto.CloudBillingData.CostTagGroupMapEntry
	(*CloudBillingData_CloudBillingBucket)(nil),          // 6: common_dto.CloudBillingData.CloudBillingBucket
	nil,                             // 7: common_dto.CostTagGroup.TagsEntry
	(PriceModel)(0),                 // 8: common_dto.PriceModel
	(CommodityDTO_CommodityType)(0), // 9: common_dto.CommodityDTO.CommodityType
	(*CurrencyAmount)(nil),          // 10: common_dto.CurrencyAmount
}
var file_CostBilling_proto_depIdxs = []int32{
	6,  // 0: common_dto.CloudBillingData.cloud_cost_buckets:type_name -> common_dto.CloudBillingData.CloudBillingBucket
	5,  // 1: common_dto.CloudBillingData.cost_tag_group_map:type_name -> common_dto.CloudBillingData.CostTagGroupMapEntry
	7,  // 2: common_dto.CostTagGroup.tags:type_name -> common_dto.CostTagGroup.TagsEntry
	8,  // 3: common_dto.CloudBillingDataPoint.price_model:type_name -> common_dto.PriceModel
	1,  // 4: common_dto.CloudBillingDataPoint.cost_category:type_name -> common_dto.CloudBillingDataPoint.CostCategory
	9,  // 5: common_dto.CloudBillingDataPoint.purchased_commodity:type_name -> common_dto.CommodityDTO.CommodityType
	10, // 6: common_dto.CloudBillingDataPoint.cost:type_name -> common_dto.CurrencyAmount
	3,  // 7: common_dto.CloudBillingData.CostTagGroupMapEntry.value:type_name -> common_dto.CostTagGroup
	4,  // 8: common_dto.CloudBillingData.CloudBillingBucket.samples:type_name -> common_dto.CloudBillingDataPoint
	0,  // 9: common_dto.CloudBillingData.CloudBillingBucket.granularity:type_name -> common_dto.CloudBillingData.CloudBillingBucket.Granularity
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_CostBilling_proto_init() }
func file_CostBilling_proto_init() {
	if File_CostBilling_proto != nil {
		return
	}
	file_CommonDTO_proto_init()
	file_CommonCost_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CostBilling_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudBillingData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_CostBilling_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CostTagGroup); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_CostBilling_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudBillingDataPoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_CostBilling_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudBillingData_CloudBillingBucket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_CostBilling_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CostBilling_proto_goTypes,
		DependencyIndexes: file_CostBilling_proto_depIdxs,
		EnumInfos:         file_CostBilling_proto_enumTypes,
		MessageInfos:      file_CostBilling_proto_msgTypes,
	}.Build()
	File_CostBilling_proto = out.File
	file_CostBilling_proto_rawDesc = nil
	file_CostBilling_proto_goTypes = nil
	file_CostBilling_proto_depIdxs = nil
}
