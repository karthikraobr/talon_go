# Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**ApplicationId** | **int32** | The ID of the application that owns this entity. | 
**Slots** | [**[]SlotDef**](SlotDef.md) | The slots defined for this application. | 
**Functions** | [**[]FunctionDef**](FunctionDef.md) | The functions defined for this application. | 
**Templates** | [**[]TemplateDef**](TemplateDef.md) | The templates defined for this application. | 
**Variables** | **string** | A stringified version of the environment&#39;s Talang variables scope. | 
**GiveawaysPools** | Pointer to [**[]GiveawaysPool**](GiveawaysPool.md) | The giveaways pools that the application is subscribed to. | [optional] 
**LoyaltyPrograms** | Pointer to [**[]LoyaltyProgram**](LoyaltyProgram.md) | The loyalty programs that the application is subscribed to. | [optional] 
**Attributes** | Pointer to [**[]Attribute**](Attribute.md) | The attributes that the application is subscribed to. | [optional] 
**AdditionalCosts** | Pointer to [**[]AccountAdditionalCost**](AccountAdditionalCost.md) | The additional costs that the application is subscribed to. | [optional] 
**Audiences** | Pointer to [**[]Audience**](Audience.md) | The audiences contained in the account which the application belongs to. | [optional] 
**Collections** | Pointer to [**[]Collection**](Collection.md) | The account-level collections that the application is subscribed to. | [optional] 

## Methods

### NewEnvironment

`func NewEnvironment(id int32, created time.Time, applicationId int32, slots []SlotDef, functions []FunctionDef, templates []TemplateDef, variables string, ) *Environment`

NewEnvironment instantiates a new Environment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentWithDefaults

`func NewEnvironmentWithDefaults() *Environment`

NewEnvironmentWithDefaults instantiates a new Environment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Environment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Environment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Environment) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *Environment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Environment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Environment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetApplicationId

`func (o *Environment) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Environment) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *Environment) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.


### GetSlots

`func (o *Environment) GetSlots() []SlotDef`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *Environment) GetSlotsOk() (*[]SlotDef, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *Environment) SetSlots(v []SlotDef)`

SetSlots sets Slots field to given value.


### GetFunctions

`func (o *Environment) GetFunctions() []FunctionDef`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *Environment) GetFunctionsOk() (*[]FunctionDef, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *Environment) SetFunctions(v []FunctionDef)`

SetFunctions sets Functions field to given value.


### GetTemplates

`func (o *Environment) GetTemplates() []TemplateDef`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *Environment) GetTemplatesOk() (*[]TemplateDef, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *Environment) SetTemplates(v []TemplateDef)`

SetTemplates sets Templates field to given value.


### GetVariables

`func (o *Environment) GetVariables() string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Environment) GetVariablesOk() (*string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Environment) SetVariables(v string)`

SetVariables sets Variables field to given value.


### GetGiveawaysPools

`func (o *Environment) GetGiveawaysPools() []GiveawaysPool`

GetGiveawaysPools returns the GiveawaysPools field if non-nil, zero value otherwise.

### GetGiveawaysPoolsOk

`func (o *Environment) GetGiveawaysPoolsOk() (*[]GiveawaysPool, bool)`

GetGiveawaysPoolsOk returns a tuple with the GiveawaysPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveawaysPools

`func (o *Environment) SetGiveawaysPools(v []GiveawaysPool)`

SetGiveawaysPools sets GiveawaysPools field to given value.

### HasGiveawaysPools

`func (o *Environment) HasGiveawaysPools() bool`

HasGiveawaysPools returns a boolean if a field has been set.

### GetLoyaltyPrograms

`func (o *Environment) GetLoyaltyPrograms() []LoyaltyProgram`

GetLoyaltyPrograms returns the LoyaltyPrograms field if non-nil, zero value otherwise.

### GetLoyaltyProgramsOk

`func (o *Environment) GetLoyaltyProgramsOk() (*[]LoyaltyProgram, bool)`

GetLoyaltyProgramsOk returns a tuple with the LoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyPrograms

`func (o *Environment) SetLoyaltyPrograms(v []LoyaltyProgram)`

SetLoyaltyPrograms sets LoyaltyPrograms field to given value.

### HasLoyaltyPrograms

`func (o *Environment) HasLoyaltyPrograms() bool`

HasLoyaltyPrograms returns a boolean if a field has been set.

### GetAttributes

`func (o *Environment) GetAttributes() []Attribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Environment) GetAttributesOk() (*[]Attribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Environment) SetAttributes(v []Attribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Environment) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetAdditionalCosts

`func (o *Environment) GetAdditionalCosts() []AccountAdditionalCost`

GetAdditionalCosts returns the AdditionalCosts field if non-nil, zero value otherwise.

### GetAdditionalCostsOk

`func (o *Environment) GetAdditionalCostsOk() (*[]AccountAdditionalCost, bool)`

GetAdditionalCostsOk returns a tuple with the AdditionalCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCosts

`func (o *Environment) SetAdditionalCosts(v []AccountAdditionalCost)`

SetAdditionalCosts sets AdditionalCosts field to given value.

### HasAdditionalCosts

`func (o *Environment) HasAdditionalCosts() bool`

HasAdditionalCosts returns a boolean if a field has been set.

### GetAudiences

`func (o *Environment) GetAudiences() []Audience`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *Environment) GetAudiencesOk() (*[]Audience, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *Environment) SetAudiences(v []Audience)`

SetAudiences sets Audiences field to given value.

### HasAudiences

`func (o *Environment) HasAudiences() bool`

HasAudiences returns a boolean if a field has been set.

### GetCollections

`func (o *Environment) GetCollections() []Collection`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *Environment) GetCollectionsOk() (*[]Collection, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *Environment) SetCollections(v []Collection)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *Environment) HasCollections() bool`

HasCollections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


