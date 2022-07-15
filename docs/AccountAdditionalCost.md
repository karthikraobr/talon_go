# AccountAdditionalCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**AccountId** | **int32** | The ID of the account that owns this entity. | 
**Name** | **string** | The additional cost name that will be used in API requests and Talang. E.g. if &#x60;name &#x3D;&#x3D; \&quot;shipping\&quot;&#x60; then you would set the shipping additional cost by including an &#x60;additionalCosts.shipping&#x60; property in your request payload. | 
**Title** | **string** | The human-readable name for the additional cost that will be shown in the Campaign Manager. Like &#x60;name&#x60;, the combination of entity and title must also be unique. | 
**Description** | **string** | A description of this additional cost. | 
**SubscribedApplicationsIds** | Pointer to **[]int32** | A list of the IDs of the applications that are subscribed to this additional cost. | [optional] 
**Type** | Pointer to **string** | The type of additional cost. The following options can be chosen: - &#x60;session&#x60;: Additional cost will be added per session, - &#x60;item&#x60;: Additional cost will be added per item, - &#x60;both&#x60;: Additional cost will be added per item and session.  | [optional] [default to "session"]

## Methods

### NewAccountAdditionalCost

`func NewAccountAdditionalCost(id int32, created time.Time, accountId int32, name string, title string, description string, ) *AccountAdditionalCost`

NewAccountAdditionalCost instantiates a new AccountAdditionalCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAdditionalCostWithDefaults

`func NewAccountAdditionalCostWithDefaults() *AccountAdditionalCost`

NewAccountAdditionalCostWithDefaults instantiates a new AccountAdditionalCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountAdditionalCost) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountAdditionalCost) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountAdditionalCost) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *AccountAdditionalCost) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AccountAdditionalCost) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AccountAdditionalCost) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetAccountId

`func (o *AccountAdditionalCost) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountAdditionalCost) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountAdditionalCost) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *AccountAdditionalCost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountAdditionalCost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountAdditionalCost) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *AccountAdditionalCost) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AccountAdditionalCost) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AccountAdditionalCost) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *AccountAdditionalCost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountAdditionalCost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountAdditionalCost) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSubscribedApplicationsIds

`func (o *AccountAdditionalCost) GetSubscribedApplicationsIds() []int32`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *AccountAdditionalCost) GetSubscribedApplicationsIdsOk() (*[]int32, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplicationsIds

`func (o *AccountAdditionalCost) SetSubscribedApplicationsIds(v []int32)`

SetSubscribedApplicationsIds sets SubscribedApplicationsIds field to given value.

### HasSubscribedApplicationsIds

`func (o *AccountAdditionalCost) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.

### GetType

`func (o *AccountAdditionalCost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountAdditionalCost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountAdditionalCost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccountAdditionalCost) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


