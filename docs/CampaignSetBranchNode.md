# CampaignSetBranchNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Indicates the node type. | 
**Name** | **string** | Name of the set | 
**Operator** | **string** | How does the set operates on its elements. | 
**Elements** | [**[]CampaignSetNode**](CampaignSetNode.md) | Child elements of this set. | 

## Methods

### NewCampaignSetBranchNode

`func NewCampaignSetBranchNode(type_ string, name string, operator string, elements []CampaignSetNode, ) *CampaignSetBranchNode`

NewCampaignSetBranchNode instantiates a new CampaignSetBranchNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignSetBranchNodeWithDefaults

`func NewCampaignSetBranchNodeWithDefaults() *CampaignSetBranchNode`

NewCampaignSetBranchNodeWithDefaults instantiates a new CampaignSetBranchNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CampaignSetBranchNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignSetBranchNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignSetBranchNode) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CampaignSetBranchNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignSetBranchNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignSetBranchNode) SetName(v string)`

SetName sets Name field to given value.


### GetOperator

`func (o *CampaignSetBranchNode) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CampaignSetBranchNode) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CampaignSetBranchNode) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetElements

`func (o *CampaignSetBranchNode) GetElements() []CampaignSetNode`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *CampaignSetBranchNode) GetElementsOk() (*[]CampaignSetNode, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *CampaignSetBranchNode) SetElements(v []CampaignSetNode)`

SetElements sets Elements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


