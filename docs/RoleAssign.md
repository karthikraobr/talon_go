# RoleAssign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | **[]int32** | An array of userIDs. | 
**Roles** | **[]int32** | An array of roleIDs. | 

## Methods

### NewRoleAssign

`func NewRoleAssign(users []int32, roles []int32, ) *RoleAssign`

NewRoleAssign instantiates a new RoleAssign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleAssignWithDefaults

`func NewRoleAssignWithDefaults() *RoleAssign`

NewRoleAssignWithDefaults instantiates a new RoleAssign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *RoleAssign) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *RoleAssign) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *RoleAssign) SetUsers(v []int32)`

SetUsers sets Users field to given value.


### GetRoles

`func (o *RoleAssign) GetRoles() []int32`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *RoleAssign) GetRolesOk() (*[]int32, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *RoleAssign) SetRoles(v []int32)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


