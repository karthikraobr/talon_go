# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**Modified** | **time.Time** | The exact moment this entity was last modified. | 
**Email** | **string** | The email address associated with your account. | 
**AccountId** | **int32** | The ID of the account that owns this entity. | 
**InviteToken** | **string** | Invite token, empty if the user as already accepted their invite. | 
**State** | **string** | Current user state. | 
**Name** | **string** | Full name | 
**Policy** | **map[string]interface{}** | User ACL Policy | 
**LatestFeedTimestamp** | Pointer to **time.Time** | Latest timestamp the user has been notified for feed. | [optional] 
**Roles** | Pointer to **[]int32** | Contains a list of all roles the user is a member of. | [optional] 
**ApplicationNotificationSubscriptions** | Pointer to **map[string]interface{}** |  | [optional] 
**AuthMethod** | Pointer to **string** | The Authentication method for this user. | [optional] 

## Methods

### NewUser

`func NewUser(id int32, created time.Time, modified time.Time, email string, accountId int32, inviteToken string, state string, name string, policy map[string]interface{}, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *User) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *User) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *User) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *User) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *User) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *User) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAccountId

`func (o *User) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *User) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *User) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetInviteToken

`func (o *User) GetInviteToken() string`

GetInviteToken returns the InviteToken field if non-nil, zero value otherwise.

### GetInviteTokenOk

`func (o *User) GetInviteTokenOk() (*string, bool)`

GetInviteTokenOk returns a tuple with the InviteToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteToken

`func (o *User) SetInviteToken(v string)`

SetInviteToken sets InviteToken field to given value.


### GetState

`func (o *User) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *User) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *User) SetState(v string)`

SetState sets State field to given value.


### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.


### GetPolicy

`func (o *User) GetPolicy() map[string]interface{}`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *User) GetPolicyOk() (*map[string]interface{}, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *User) SetPolicy(v map[string]interface{})`

SetPolicy sets Policy field to given value.


### GetLatestFeedTimestamp

`func (o *User) GetLatestFeedTimestamp() time.Time`

GetLatestFeedTimestamp returns the LatestFeedTimestamp field if non-nil, zero value otherwise.

### GetLatestFeedTimestampOk

`func (o *User) GetLatestFeedTimestampOk() (*time.Time, bool)`

GetLatestFeedTimestampOk returns a tuple with the LatestFeedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestFeedTimestamp

`func (o *User) SetLatestFeedTimestamp(v time.Time)`

SetLatestFeedTimestamp sets LatestFeedTimestamp field to given value.

### HasLatestFeedTimestamp

`func (o *User) HasLatestFeedTimestamp() bool`

HasLatestFeedTimestamp returns a boolean if a field has been set.

### GetRoles

`func (o *User) GetRoles() []int32`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *User) GetRolesOk() (*[]int32, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *User) SetRoles(v []int32)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *User) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetApplicationNotificationSubscriptions

`func (o *User) GetApplicationNotificationSubscriptions() map[string]interface{}`

GetApplicationNotificationSubscriptions returns the ApplicationNotificationSubscriptions field if non-nil, zero value otherwise.

### GetApplicationNotificationSubscriptionsOk

`func (o *User) GetApplicationNotificationSubscriptionsOk() (*map[string]interface{}, bool)`

GetApplicationNotificationSubscriptionsOk returns a tuple with the ApplicationNotificationSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNotificationSubscriptions

`func (o *User) SetApplicationNotificationSubscriptions(v map[string]interface{})`

SetApplicationNotificationSubscriptions sets ApplicationNotificationSubscriptions field to given value.

### HasApplicationNotificationSubscriptions

`func (o *User) HasApplicationNotificationSubscriptions() bool`

HasApplicationNotificationSubscriptions returns a boolean if a field has been set.

### GetAuthMethod

`func (o *User) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *User) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *User) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *User) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


