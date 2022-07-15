# NewInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the user being invited. | [optional] 
**Email** | **string** |  | 
**Acl** | **string** | The &#x60;Access Control List&#x60; json defining the role of the user. This represents the actual access control on the user level. Use one of the following: - normal user: &#x60;{\&quot;Role\&quot;: 0}&#x60; - admin: &#x60;{\&quot;Role\&quot;: 127}&#x60;  | 
**Roles** | Pointer to **[]int32** | An array of roleIDs to assign the new user to. | [optional] 

## Methods

### NewNewInvitation

`func NewNewInvitation(email string, acl string, ) *NewInvitation`

NewNewInvitation instantiates a new NewInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewInvitationWithDefaults

`func NewNewInvitationWithDefaults() *NewInvitation`

NewNewInvitationWithDefaults instantiates a new NewInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewInvitation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewInvitation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewInvitation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NewInvitation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *NewInvitation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NewInvitation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NewInvitation) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAcl

`func (o *NewInvitation) GetAcl() string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *NewInvitation) GetAclOk() (*string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *NewInvitation) SetAcl(v string)`

SetAcl sets Acl field to given value.


### GetRoles

`func (o *NewInvitation) GetRoles() []int32`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *NewInvitation) GetRolesOk() (*[]int32, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *NewInvitation) SetRoles(v []int32)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *NewInvitation) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


