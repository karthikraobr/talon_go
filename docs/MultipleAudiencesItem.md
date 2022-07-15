# MultipleAudiencesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**Name** | **string** | The human-friendly display name for this audience. | 
**IntegrationId** | **string** | The ID of this audience in the third-party integration. | 
**Status** | **string** | Indicates whether the audience is new, updated or unmodified by the request.  | 

## Methods

### NewMultipleAudiencesItem

`func NewMultipleAudiencesItem(id int32, created time.Time, name string, integrationId string, status string, ) *MultipleAudiencesItem`

NewMultipleAudiencesItem instantiates a new MultipleAudiencesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleAudiencesItemWithDefaults

`func NewMultipleAudiencesItemWithDefaults() *MultipleAudiencesItem`

NewMultipleAudiencesItemWithDefaults instantiates a new MultipleAudiencesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MultipleAudiencesItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MultipleAudiencesItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MultipleAudiencesItem) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *MultipleAudiencesItem) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MultipleAudiencesItem) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MultipleAudiencesItem) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetName

`func (o *MultipleAudiencesItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MultipleAudiencesItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MultipleAudiencesItem) SetName(v string)`

SetName sets Name field to given value.


### GetIntegrationId

`func (o *MultipleAudiencesItem) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *MultipleAudiencesItem) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *MultipleAudiencesItem) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetStatus

`func (o *MultipleAudiencesItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MultipleAudiencesItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MultipleAudiencesItem) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


