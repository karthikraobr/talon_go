# ApplicationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**ApplicationId** | **int32** | The ID of the application that owns this entity. | 
**ProfileId** | Pointer to **int32** | The globally unique Talon.One ID of the customer that created this entity. | [optional] 
**SessionId** | Pointer to **int32** | The globally unique Talon.One ID of the session that contains this event. | [optional] 
**Type** | **string** | A string representing the event. Must not be a reserved event name. | 
**Attributes** | **map[string]interface{}** | Additional JSON serialized data associated with the event. | 
**Effects** | **[]map[string]interface{}** | An array containing the effects that were applied as a result of this event. | 
**RuleFailureReasons** | Pointer to [**[]RuleFailureReason**](RuleFailureReason.md) | An array containing the rule failure reasons which happened during this event. | [optional] 

## Methods

### NewApplicationEvent

`func NewApplicationEvent(id int32, created time.Time, applicationId int32, type_ string, attributes map[string]interface{}, effects []map[string]interface{}, ) *ApplicationEvent`

NewApplicationEvent instantiates a new ApplicationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationEventWithDefaults

`func NewApplicationEventWithDefaults() *ApplicationEvent`

NewApplicationEventWithDefaults instantiates a new ApplicationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationEvent) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *ApplicationEvent) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApplicationEvent) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApplicationEvent) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetApplicationId

`func (o *ApplicationEvent) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationEvent) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationEvent) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.


### GetProfileId

`func (o *ApplicationEvent) GetProfileId() int32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *ApplicationEvent) GetProfileIdOk() (*int32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *ApplicationEvent) SetProfileId(v int32)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *ApplicationEvent) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetSessionId

`func (o *ApplicationEvent) GetSessionId() int32`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ApplicationEvent) GetSessionIdOk() (*int32, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ApplicationEvent) SetSessionId(v int32)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ApplicationEvent) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetType

`func (o *ApplicationEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationEvent) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ApplicationEvent) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationEvent) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationEvent) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetEffects

`func (o *ApplicationEvent) GetEffects() []map[string]interface{}`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *ApplicationEvent) GetEffectsOk() (*[]map[string]interface{}, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *ApplicationEvent) SetEffects(v []map[string]interface{})`

SetEffects sets Effects field to given value.


### GetRuleFailureReasons

`func (o *ApplicationEvent) GetRuleFailureReasons() []RuleFailureReason`

GetRuleFailureReasons returns the RuleFailureReasons field if non-nil, zero value otherwise.

### GetRuleFailureReasonsOk

`func (o *ApplicationEvent) GetRuleFailureReasonsOk() (*[]RuleFailureReason, bool)`

GetRuleFailureReasonsOk returns a tuple with the RuleFailureReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleFailureReasons

`func (o *ApplicationEvent) SetRuleFailureReasons(v []RuleFailureReason)`

SetRuleFailureReasons sets RuleFailureReasons field to given value.

### HasRuleFailureReasons

`func (o *ApplicationEvent) HasRuleFailureReasons() bool`

HasRuleFailureReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


