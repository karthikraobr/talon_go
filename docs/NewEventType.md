# NewEventType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The human-friendly display name for this event type. Use a short, past-tense, description of the event. | 
**Name** | **string** | The machine-friendly canonical name for this event type. This will be used in URLs, and cannot be changed after an event type has been created. | 
**Description** | Pointer to **string** | An explanation of when the event type is triggered. Write this with a campaign manager in mind. For example:  &gt; The \&quot;Payment Accepted\&quot; event is triggered after successful processing of a payment by our payment gateway.  | [optional] 

## Methods

### NewNewEventType

`func NewNewEventType(title string, name string, ) *NewEventType`

NewNewEventType instantiates a new NewEventType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewEventTypeWithDefaults

`func NewNewEventTypeWithDefaults() *NewEventType`

NewNewEventTypeWithDefaults instantiates a new NewEventType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *NewEventType) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewEventType) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewEventType) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetName

`func (o *NewEventType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewEventType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewEventType) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NewEventType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewEventType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewEventType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewEventType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


