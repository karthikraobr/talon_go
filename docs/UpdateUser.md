# UpdateUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The user name. | [optional] 
**Policy** | Pointer to **string** | ACL Json. | [optional] 
**State** | Pointer to **string** | New state (\&quot;deactivated\&quot; or \&quot;active\&quot;) for the user. Only usable by admins for the user. | [optional] 
**Roles** | Pointer to **[]int32** | List of roles to assign to the user. | [optional] 
**ApplicationNotificationSubscriptions** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateUser

`func NewUpdateUser() *UpdateUser`

NewUpdateUser instantiates a new UpdateUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserWithDefaults

`func NewUpdateUserWithDefaults() *UpdateUser`

NewUpdateUserWithDefaults instantiates a new UpdateUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicy

`func (o *UpdateUser) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UpdateUser) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *UpdateUser) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *UpdateUser) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetState

`func (o *UpdateUser) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateUser) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateUser) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpdateUser) HasState() bool`

HasState returns a boolean if a field has been set.

### GetRoles

`func (o *UpdateUser) GetRoles() []int32`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateUser) GetRolesOk() (*[]int32, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateUser) SetRoles(v []int32)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpdateUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetApplicationNotificationSubscriptions

`func (o *UpdateUser) GetApplicationNotificationSubscriptions() map[string]interface{}`

GetApplicationNotificationSubscriptions returns the ApplicationNotificationSubscriptions field if non-nil, zero value otherwise.

### GetApplicationNotificationSubscriptionsOk

`func (o *UpdateUser) GetApplicationNotificationSubscriptionsOk() (*map[string]interface{}, bool)`

GetApplicationNotificationSubscriptionsOk returns a tuple with the ApplicationNotificationSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNotificationSubscriptions

`func (o *UpdateUser) SetApplicationNotificationSubscriptions(v map[string]interface{})`

SetApplicationNotificationSubscriptions sets ApplicationNotificationSubscriptions field to given value.

### HasApplicationNotificationSubscriptions

`func (o *UpdateUser) HasApplicationNotificationSubscriptions() bool`

HasApplicationNotificationSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


