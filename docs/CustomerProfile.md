# CustomerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. The exact moment this entity was created. | 
**IntegrationId** | **string** | The integration ID set by your integration layer. | 
**Attributes** | **map[string]interface{}** | Arbitrary properties associated with this item. | 
**AccountId** | **int32** | The ID of the Talon.One account that owns this profile. | 
**ClosedSessions** | **int32** | The total amount of closed sessions by a customer. A closed session is a successful purchase. | 
**TotalSales** | **float32** | Sum of all purchases made by this customer. | 
**LoyaltyMemberships** | Pointer to [**[]LoyaltyMembership**](LoyaltyMembership.md) | **DEPRECATED** A list of loyalty programs joined by the customer.  | [optional] 
**AudienceMemberships** | Pointer to [**[]AudienceMembership**](AudienceMembership.md) | A list of audiences the customer belongs to. | [optional] 
**LastActivity** | **time.Time** | Timestamp of the most recent event received from this customer. This field is updated on calls that trigger the rule-engine and that are not [dry requests](https://docs.talon.one/docs/dev/integration-api/dry-requests/#overlay).  For example, [reserving a coupon](https://docs.talon.one/integration-api/#operation/createCouponReservation) for a customer doesn&#39;t impact this field.  | 

## Methods

### NewCustomerProfile

`func NewCustomerProfile(id int32, created time.Time, integrationId string, attributes map[string]interface{}, accountId int32, closedSessions int32, totalSales float32, lastActivity time.Time, ) *CustomerProfile`

NewCustomerProfile instantiates a new CustomerProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerProfileWithDefaults

`func NewCustomerProfileWithDefaults() *CustomerProfile`

NewCustomerProfileWithDefaults instantiates a new CustomerProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerProfile) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *CustomerProfile) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomerProfile) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomerProfile) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetIntegrationId

`func (o *CustomerProfile) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *CustomerProfile) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *CustomerProfile) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetAttributes

`func (o *CustomerProfile) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerProfile) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerProfile) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetAccountId

`func (o *CustomerProfile) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CustomerProfile) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CustomerProfile) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetClosedSessions

`func (o *CustomerProfile) GetClosedSessions() int32`

GetClosedSessions returns the ClosedSessions field if non-nil, zero value otherwise.

### GetClosedSessionsOk

`func (o *CustomerProfile) GetClosedSessionsOk() (*int32, bool)`

GetClosedSessionsOk returns a tuple with the ClosedSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedSessions

`func (o *CustomerProfile) SetClosedSessions(v int32)`

SetClosedSessions sets ClosedSessions field to given value.


### GetTotalSales

`func (o *CustomerProfile) GetTotalSales() float32`

GetTotalSales returns the TotalSales field if non-nil, zero value otherwise.

### GetTotalSalesOk

`func (o *CustomerProfile) GetTotalSalesOk() (*float32, bool)`

GetTotalSalesOk returns a tuple with the TotalSales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSales

`func (o *CustomerProfile) SetTotalSales(v float32)`

SetTotalSales sets TotalSales field to given value.


### GetLoyaltyMemberships

`func (o *CustomerProfile) GetLoyaltyMemberships() []LoyaltyMembership`

GetLoyaltyMemberships returns the LoyaltyMemberships field if non-nil, zero value otherwise.

### GetLoyaltyMembershipsOk

`func (o *CustomerProfile) GetLoyaltyMembershipsOk() (*[]LoyaltyMembership, bool)`

GetLoyaltyMembershipsOk returns a tuple with the LoyaltyMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyMemberships

`func (o *CustomerProfile) SetLoyaltyMemberships(v []LoyaltyMembership)`

SetLoyaltyMemberships sets LoyaltyMemberships field to given value.

### HasLoyaltyMemberships

`func (o *CustomerProfile) HasLoyaltyMemberships() bool`

HasLoyaltyMemberships returns a boolean if a field has been set.

### GetAudienceMemberships

`func (o *CustomerProfile) GetAudienceMemberships() []AudienceMembership`

GetAudienceMemberships returns the AudienceMemberships field if non-nil, zero value otherwise.

### GetAudienceMembershipsOk

`func (o *CustomerProfile) GetAudienceMembershipsOk() (*[]AudienceMembership, bool)`

GetAudienceMembershipsOk returns a tuple with the AudienceMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceMemberships

`func (o *CustomerProfile) SetAudienceMemberships(v []AudienceMembership)`

SetAudienceMemberships sets AudienceMemberships field to given value.

### HasAudienceMemberships

`func (o *CustomerProfile) HasAudienceMemberships() bool`

HasAudienceMemberships returns a boolean if a field has been set.

### GetLastActivity

`func (o *CustomerProfile) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *CustomerProfile) GetLastActivityOk() (*time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *CustomerProfile) SetLastActivity(v time.Time)`

SetLastActivity sets LastActivity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


