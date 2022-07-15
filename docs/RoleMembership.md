# RoleMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleID** | **int32** | ID of role. | 
**UserID** | **int32** | ID of User. | 

## Methods

### NewRoleMembership

`func NewRoleMembership(roleID int32, userID int32, ) *RoleMembership`

NewRoleMembership instantiates a new RoleMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleMembershipWithDefaults

`func NewRoleMembershipWithDefaults() *RoleMembership`

NewRoleMembershipWithDefaults instantiates a new RoleMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleID

`func (o *RoleMembership) GetRoleID() int32`

GetRoleID returns the RoleID field if non-nil, zero value otherwise.

### GetRoleIDOk

`func (o *RoleMembership) GetRoleIDOk() (*int32, bool)`

GetRoleIDOk returns a tuple with the RoleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleID

`func (o *RoleMembership) SetRoleID(v int32)`

SetRoleID sets RoleID field to given value.


### GetUserID

`func (o *RoleMembership) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *RoleMembership) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *RoleMembership) SetUserID(v int32)`

SetUserID sets UserID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


