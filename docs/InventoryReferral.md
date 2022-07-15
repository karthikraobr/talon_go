# InventoryReferral

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**StartDate** | Pointer to **time.Time** | Timestamp at which point the referral code becomes valid. | [optional] 
**ExpiryDate** | Pointer to **time.Time** | Expiry date of the referral code. Referral never expires if this is omitted, zero, or negative. | [optional] 
**UsageLimit** | **int32** | The number of times a referral code can be used. &#x60;0&#x60; means no limit but any campaign usage limits will still apply.  | 
**CampaignId** | **int32** | ID of the campaign from which the referral received the referral code. | 
**AdvocateProfileIntegrationId** | **string** | The Integration ID of the Advocate&#39;s Profile. | 
**FriendProfileIntegrationId** | Pointer to **string** | An optional Integration ID of the Friend&#39;s Profile. | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | Arbitrary properties associated with this item. | [optional] 
**ImportId** | Pointer to **int32** | The ID of the Import which created this referral. | [optional] 
**Code** | **string** | The referral code. | 
**UsageCounter** | **int32** | The number of times this referral code has been successfully used. | 
**BatchId** | Pointer to **string** | The ID of the batch the referrals belong to. | [optional] 
**ReferredCustomers** | **[]string** | An array of referred customers. | 

## Methods

### NewInventoryReferral

`func NewInventoryReferral(id int32, created time.Time, usageLimit int32, campaignId int32, advocateProfileIntegrationId string, code string, usageCounter int32, referredCustomers []string, ) *InventoryReferral`

NewInventoryReferral instantiates a new InventoryReferral object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryReferralWithDefaults

`func NewInventoryReferralWithDefaults() *InventoryReferral`

NewInventoryReferralWithDefaults instantiates a new InventoryReferral object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InventoryReferral) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryReferral) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryReferral) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *InventoryReferral) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InventoryReferral) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InventoryReferral) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetStartDate

`func (o *InventoryReferral) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InventoryReferral) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InventoryReferral) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InventoryReferral) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *InventoryReferral) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *InventoryReferral) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *InventoryReferral) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *InventoryReferral) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetUsageLimit

`func (o *InventoryReferral) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *InventoryReferral) GetUsageLimitOk() (*int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *InventoryReferral) SetUsageLimit(v int32)`

SetUsageLimit sets UsageLimit field to given value.


### GetCampaignId

`func (o *InventoryReferral) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *InventoryReferral) GetCampaignIdOk() (*int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *InventoryReferral) SetCampaignId(v int32)`

SetCampaignId sets CampaignId field to given value.


### GetAdvocateProfileIntegrationId

`func (o *InventoryReferral) GetAdvocateProfileIntegrationId() string`

GetAdvocateProfileIntegrationId returns the AdvocateProfileIntegrationId field if non-nil, zero value otherwise.

### GetAdvocateProfileIntegrationIdOk

`func (o *InventoryReferral) GetAdvocateProfileIntegrationIdOk() (*string, bool)`

GetAdvocateProfileIntegrationIdOk returns a tuple with the AdvocateProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvocateProfileIntegrationId

`func (o *InventoryReferral) SetAdvocateProfileIntegrationId(v string)`

SetAdvocateProfileIntegrationId sets AdvocateProfileIntegrationId field to given value.


### GetFriendProfileIntegrationId

`func (o *InventoryReferral) GetFriendProfileIntegrationId() string`

GetFriendProfileIntegrationId returns the FriendProfileIntegrationId field if non-nil, zero value otherwise.

### GetFriendProfileIntegrationIdOk

`func (o *InventoryReferral) GetFriendProfileIntegrationIdOk() (*string, bool)`

GetFriendProfileIntegrationIdOk returns a tuple with the FriendProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendProfileIntegrationId

`func (o *InventoryReferral) SetFriendProfileIntegrationId(v string)`

SetFriendProfileIntegrationId sets FriendProfileIntegrationId field to given value.

### HasFriendProfileIntegrationId

`func (o *InventoryReferral) HasFriendProfileIntegrationId() bool`

HasFriendProfileIntegrationId returns a boolean if a field has been set.

### GetAttributes

`func (o *InventoryReferral) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InventoryReferral) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InventoryReferral) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *InventoryReferral) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetImportId

`func (o *InventoryReferral) GetImportId() int32`

GetImportId returns the ImportId field if non-nil, zero value otherwise.

### GetImportIdOk

`func (o *InventoryReferral) GetImportIdOk() (*int32, bool)`

GetImportIdOk returns a tuple with the ImportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportId

`func (o *InventoryReferral) SetImportId(v int32)`

SetImportId sets ImportId field to given value.

### HasImportId

`func (o *InventoryReferral) HasImportId() bool`

HasImportId returns a boolean if a field has been set.

### GetCode

`func (o *InventoryReferral) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InventoryReferral) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InventoryReferral) SetCode(v string)`

SetCode sets Code field to given value.


### GetUsageCounter

`func (o *InventoryReferral) GetUsageCounter() int32`

GetUsageCounter returns the UsageCounter field if non-nil, zero value otherwise.

### GetUsageCounterOk

`func (o *InventoryReferral) GetUsageCounterOk() (*int32, bool)`

GetUsageCounterOk returns a tuple with the UsageCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCounter

`func (o *InventoryReferral) SetUsageCounter(v int32)`

SetUsageCounter sets UsageCounter field to given value.


### GetBatchId

`func (o *InventoryReferral) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *InventoryReferral) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *InventoryReferral) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *InventoryReferral) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetReferredCustomers

`func (o *InventoryReferral) GetReferredCustomers() []string`

GetReferredCustomers returns the ReferredCustomers field if non-nil, zero value otherwise.

### GetReferredCustomersOk

`func (o *InventoryReferral) GetReferredCustomersOk() (*[]string, bool)`

GetReferredCustomersOk returns a tuple with the ReferredCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredCustomers

`func (o *InventoryReferral) SetReferredCustomers(v []string)`

SetReferredCustomers sets ReferredCustomers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


