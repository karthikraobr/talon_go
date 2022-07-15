# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**ApplicationId** | **int32** | The ID of the application that owns this entity. | 
**ProfileId** | Pointer to **string** | ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known &#x60;profileId&#x60;, we recommend you use a guest &#x60;profileId&#x60;.  | [optional] 
**Type** | **string** | A string representing the event. Must not be a reserved event name. | 
**Attributes** | **map[string]interface{}** | Arbitrary additional JSON data associated with the event. | 
**SessionId** | Pointer to **string** | The ID of the session that this event occurred in. | [optional] 
**Effects** | **[]map[string]interface{}** | An array of effects generated by the rules of the enabled campaigns of the Application.  You decide how to apply them in your system. See the list of [API effects](/docs/dev/integration-api/api-effects).  | 
**LedgerEntries** | [**[]LedgerEntry**](LedgerEntry.md) | Ledger entries for the event. | 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewEvent

`func NewEvent(id int32, created time.Time, applicationId int32, type_ string, attributes map[string]interface{}, effects []map[string]interface{}, ledgerEntries []LedgerEntry, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Event) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Event) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Event) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *Event) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Event) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Event) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetApplicationId

`func (o *Event) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Event) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *Event) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.


### GetProfileId

`func (o *Event) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *Event) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *Event) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *Event) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetType

`func (o *Event) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Event) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Event) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *Event) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Event) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Event) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetSessionId

`func (o *Event) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *Event) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *Event) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *Event) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetEffects

`func (o *Event) GetEffects() []map[string]interface{}`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *Event) GetEffectsOk() (*[]map[string]interface{}, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *Event) SetEffects(v []map[string]interface{})`

SetEffects sets Effects field to given value.


### GetLedgerEntries

`func (o *Event) GetLedgerEntries() []LedgerEntry`

GetLedgerEntries returns the LedgerEntries field if non-nil, zero value otherwise.

### GetLedgerEntriesOk

`func (o *Event) GetLedgerEntriesOk() (*[]LedgerEntry, bool)`

GetLedgerEntriesOk returns a tuple with the LedgerEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerEntries

`func (o *Event) SetLedgerEntries(v []LedgerEntry)`

SetLedgerEntries sets LedgerEntries field to given value.


### GetMeta

`func (o *Event) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Event) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Event) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Event) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

