# CustomEffect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**AccountId** | **int32** | The ID of the account that owns this entity. | 
**Modified** | **time.Time** | The exact moment this entity was last modified. | 
**ApplicationIds** | **[]int32** | The IDs of the applications that are related to this entity. | 
**Name** | **string** | The name of this effect. | 
**Title** | **string** | The title of this effect. | 
**Payload** | **string** | The JSON payload of this effect. | 
**Description** | Pointer to **string** | The description of this effect. | [optional] 
**Enabled** | **bool** | Determines if this effect is active. | 
**Params** | Pointer to [**[]TemplateArgDef**](TemplateArgDef.md) | Array of template argument definitions. | [optional] 
**ModifiedBy** | Pointer to **int32** | ID of the user who last updated this effect if available. | [optional] 
**CreatedBy** | **int32** | ID of the user who created this effect. | 

## Methods

### NewCustomEffect

`func NewCustomEffect(id int32, created time.Time, accountId int32, modified time.Time, applicationIds []int32, name string, title string, payload string, enabled bool, createdBy int32, ) *CustomEffect`

NewCustomEffect instantiates a new CustomEffect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEffectWithDefaults

`func NewCustomEffectWithDefaults() *CustomEffect`

NewCustomEffectWithDefaults instantiates a new CustomEffect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomEffect) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomEffect) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomEffect) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *CustomEffect) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomEffect) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomEffect) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetAccountId

`func (o *CustomEffect) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CustomEffect) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CustomEffect) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetModified

`func (o *CustomEffect) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CustomEffect) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *CustomEffect) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetApplicationIds

`func (o *CustomEffect) GetApplicationIds() []int32`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *CustomEffect) GetApplicationIdsOk() (*[]int32, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *CustomEffect) SetApplicationIds(v []int32)`

SetApplicationIds sets ApplicationIds field to given value.


### GetName

`func (o *CustomEffect) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEffect) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomEffect) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *CustomEffect) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomEffect) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomEffect) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPayload

`func (o *CustomEffect) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CustomEffect) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CustomEffect) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetDescription

`func (o *CustomEffect) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomEffect) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomEffect) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomEffect) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CustomEffect) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustomEffect) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustomEffect) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetParams

`func (o *CustomEffect) GetParams() []TemplateArgDef`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CustomEffect) GetParamsOk() (*[]TemplateArgDef, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CustomEffect) SetParams(v []TemplateArgDef)`

SetParams sets Params field to given value.

### HasParams

`func (o *CustomEffect) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetModifiedBy

`func (o *CustomEffect) GetModifiedBy() int32`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *CustomEffect) GetModifiedByOk() (*int32, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *CustomEffect) SetModifiedBy(v int32)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *CustomEffect) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CustomEffect) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CustomEffect) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CustomEffect) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


