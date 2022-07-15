# LoyaltyTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**ProgramID** | **int32** | The ID of the loyalty program that owns this entity. | 
**Name** | **string** | The name of the tier | 
**MinPoints** | **float32** | The minimum amount of points required to be eligible for the tier. | 

## Methods

### NewLoyaltyTier

`func NewLoyaltyTier(id int32, created time.Time, programID int32, name string, minPoints float32, ) *LoyaltyTier`

NewLoyaltyTier instantiates a new LoyaltyTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyTierWithDefaults

`func NewLoyaltyTierWithDefaults() *LoyaltyTier`

NewLoyaltyTierWithDefaults instantiates a new LoyaltyTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoyaltyTier) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoyaltyTier) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoyaltyTier) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *LoyaltyTier) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LoyaltyTier) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LoyaltyTier) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetProgramID

`func (o *LoyaltyTier) GetProgramID() int32`

GetProgramID returns the ProgramID field if non-nil, zero value otherwise.

### GetProgramIDOk

`func (o *LoyaltyTier) GetProgramIDOk() (*int32, bool)`

GetProgramIDOk returns a tuple with the ProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramID

`func (o *LoyaltyTier) SetProgramID(v int32)`

SetProgramID sets ProgramID field to given value.


### GetName

`func (o *LoyaltyTier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoyaltyTier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoyaltyTier) SetName(v string)`

SetName sets Name field to given value.


### GetMinPoints

`func (o *LoyaltyTier) GetMinPoints() float32`

GetMinPoints returns the MinPoints field if non-nil, zero value otherwise.

### GetMinPointsOk

`func (o *LoyaltyTier) GetMinPointsOk() (*float32, bool)`

GetMinPointsOk returns a tuple with the MinPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPoints

`func (o *LoyaltyTier) SetMinPoints(v float32)`

SetMinPoints sets MinPoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


