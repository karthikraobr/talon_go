# Audience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int32** | The ID of the account that owns this entity. | 
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**Name** | **string** | The human-friendly display name for this audience. | 
**Integration** | Pointer to **string** | The Talon.One-supported [3rd-party platform](https://docs.talon.one/docs/dev/technology-partners/overview) that this audience was created in.  For example, &#x60;mParticle&#x60;, &#x60;Segment&#x60;, &#x60;Selligent&#x60;, &#x60;Braze&#x60;, or &#x60;Iterable&#x60;.  **Note:** If you do not integrate with any of these platforms, do not use this property.  | [optional] 
**IntegrationId** | Pointer to **string** | The ID of this audience in the third-party integration.  **Note:** To create an audience that doesn&#39;t come from a 3rd party platform, do not use this property.  | [optional] 
**Sandbox** | Pointer to **bool** | Indicates if this is a live or sandbox Application. | [optional] 
**CreatedIn3rdParty** | Pointer to **bool** | Determines if this audience is a 3rd party audience or not. | [optional] 

## Methods

### NewAudience

`func NewAudience(accountId int32, id int32, created time.Time, name string, ) *Audience`

NewAudience instantiates a new Audience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceWithDefaults

`func NewAudienceWithDefaults() *Audience`

NewAudienceWithDefaults instantiates a new Audience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Audience) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Audience) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Audience) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetId

`func (o *Audience) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Audience) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Audience) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *Audience) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Audience) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Audience) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetName

`func (o *Audience) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Audience) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Audience) SetName(v string)`

SetName sets Name field to given value.


### GetIntegration

`func (o *Audience) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *Audience) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *Audience) SetIntegration(v string)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *Audience) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetIntegrationId

`func (o *Audience) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *Audience) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *Audience) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *Audience) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetSandbox

`func (o *Audience) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *Audience) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *Audience) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *Audience) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### GetCreatedIn3rdParty

`func (o *Audience) GetCreatedIn3rdParty() bool`

GetCreatedIn3rdParty returns the CreatedIn3rdParty field if non-nil, zero value otherwise.

### GetCreatedIn3rdPartyOk

`func (o *Audience) GetCreatedIn3rdPartyOk() (*bool, bool)`

GetCreatedIn3rdPartyOk returns a tuple with the CreatedIn3rdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedIn3rdParty

`func (o *Audience) SetCreatedIn3rdParty(v bool)`

SetCreatedIn3rdParty sets CreatedIn3rdParty field to given value.

### HasCreatedIn3rdParty

`func (o *Audience) HasCreatedIn3rdParty() bool`

HasCreatedIn3rdParty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


