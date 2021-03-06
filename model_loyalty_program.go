/*
Integration API reference docs

Use the Integration API to push data to and retrieve data from Talon.One in real time. For more background information about this API, see [Integration API overview](/docs/dev/integration-api/overview)  For example, use this API to share shopping cart information as a session with Talon.One and evaluate promotion rules. You can also create custom events to track specific actions that do not fit into the session data model.  Ensure you [authenticate](#section/Authentication) to make requests to the API.  <div class=\"redoc-section\">   <p class=\"title\">Are you looking for a different API?</p>    If you need the API to:    - Interact with the Campaign Manager for backoffice operations, see [the Management API reference docs](https://docs.talon.one/management-api).   - Integrate with Talon.One from a CEP or CDP platform, see [the Third-party API reference docs](https://docs.talon.one/third-party-api).  </div>  # Authentication  <SecurityDefinitions /> 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package talon

import (
	"encoding/json"
	"time"
)

// LoyaltyProgram 
type LoyaltyProgram struct {
	// The ID of loyalty program. Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The display title for the Loyalty Program.
	Title string `json:"title"`
	// Description of our Loyalty Program.
	Description string `json:"description"`
	// A list containing the IDs of all applications that are subscribed to this Loyalty Program.
	SubscribedApplications []int32 `json:"subscribedApplications"`
	// Indicates the default duration after which new loyalty points should expire. The format is a number, followed by one letter indicating the unit; like '1h' or '40m'.
	DefaultValidity string `json:"defaultValidity"`
	// Indicates the default duration for the pending time, after which points will be valid. The format is a number followed by a duration unit, like '1h' or '40m'.
	DefaultPending string `json:"defaultPending"`
	// Indicates if this program supports subledgers inside the program.
	AllowSubledger bool `json:"allowSubledger"`
	// The max amount of user profiles with whom a card can be shared. This can be set to 0 for no limit. This property is only used when `cardBased` is `true`. 
	UsersPerCardLimit *int32 `json:"usersPerCardLimit,omitempty"`
	// The ID of the Talon.One account that owns this program.
	AccountID int32 `json:"accountID"`
	// The internal name for the Loyalty Program. This is an immutable value.
	Name string `json:"name"`
	// The tiers in this loyalty program.
	Tiers []LoyaltyTier `json:"tiers,omitempty"`
	// A string containing an IANA timezone descriptor.
	Timezone string `json:"timezone"`
	// Defines the type of loyalty program: - `true`: the program is a card-based. - `false`: the program is profile-based. 
	CardBased bool `json:"cardBased"`
}

// NewLoyaltyProgram instantiates a new LoyaltyProgram object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoyaltyProgram(id int32, created time.Time, title string, description string, subscribedApplications []int32, defaultValidity string, defaultPending string, allowSubledger bool, accountID int32, name string, timezone string, cardBased bool) *LoyaltyProgram {
	this := LoyaltyProgram{}
	this.Id = id
	this.Created = created
	this.Title = title
	this.Description = description
	this.SubscribedApplications = subscribedApplications
	this.DefaultValidity = defaultValidity
	this.DefaultPending = defaultPending
	this.AllowSubledger = allowSubledger
	this.AccountID = accountID
	this.Name = name
	this.Timezone = timezone
	this.CardBased = cardBased
	return &this
}

// NewLoyaltyProgramWithDefaults instantiates a new LoyaltyProgram object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoyaltyProgramWithDefaults() *LoyaltyProgram {
	this := LoyaltyProgram{}
	var cardBased bool = false
	this.CardBased = cardBased
	return &this
}

// GetId returns the Id field value
func (o *LoyaltyProgram) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LoyaltyProgram) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LoyaltyProgram) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *LoyaltyProgram) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *LoyaltyProgram) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *LoyaltyProgram) SetCreated(v time.Time) {
	o.Created = v
}

// GetTitle returns the Title field value
func (o *LoyaltyProgram) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *LoyaltyProgram) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *LoyaltyProgram) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *LoyaltyProgram) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *LoyaltyProgram) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *LoyaltyProgram) SetDescription(v string) {
	o.Description = v
}

// GetSubscribedApplications returns the SubscribedApplications field value
func (o *LoyaltyProgram) GetSubscribedApplications() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.SubscribedApplications
}

// GetSubscribedApplicationsOk returns a tuple with the SubscribedApplications field value
// and a boolean to check if the value has been set.
func (o *LoyaltyProgram) GetSubscribedApplicationsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscribedApplications, true
}

// SetSubscribedApplications sets field value
func (o *LoyaltyProgram) SetSubscribedApplications(v []int32) {
	o.SubscribedApplications = v
}

// GetDefaultValidity returns the DefaultValidity field value
func (o *LoyaltyProgram) GetDefaultValidity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultValidity
}

// GetDefaultValidityOk returns a tuple with the DefaultValidity field value
// and a boolean to check if the value has been set.
func (o *LoyaltyProgram) GetDefaultValidityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultValidity, true
}

// SetDefaultValidity sets field value
func (o *LoyaltyProgram) SetDefaultValidity(v string) {
	o.DefaultValidity = v
}

// GetDefaultPending returns the DefaultPending field value
func (o *LoyaltyProgram) GetDefaultPending() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultPending
}

// GetDefaultPendingOk returns a tuple with the DefaultPending field value
// and a boolean to check if the value has been set.
func (o *LoyaltyProgram) GetDefaultPendingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultPending, true
}

// SetDefaultPending sets field value
func (o *LoyaltyProgram) SetDefaultPending(v string) {
	o.DefaultPending = v
}

// GetAllowSubledger returns the AllowSubledger field value
func (o *LoyaltyProgram) GetAllowSubledger() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowSubledger
}

// GetAllowSubledgerOk returns a tuple with the AllowSubledger field value
// and a boolean to check if the value has been set.
func (o *LoyaltyProgram) GetAllowSubledgerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowSubledger, true
}

// SetAllowSubledger sets field value
func (o *LoyaltyProgram) SetAllowSubledger(v bool) {
	o.AllowSubledger = v
}

// GetUsersPerCardLimit returns the UsersPerCardLimit field value if set, zero value otherwise.
func (o *LoyaltyProgram) GetUsersPerCardLimit() int32 {
	if o == nil || o.UsersPerCardLimit == nil {
		var ret int32
		return ret
	}
	return *o.UsersPerCardLimit
}

// GetUsersPerCardLimitOk returns a tuple with the UsersPerCardLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyProgram) GetUsersPerCardLimitOk() (*int32, bool) {
	if o == nil || o.UsersPerCardLimit == nil {
		return nil, false
	}
	return o.UsersPerCardLimit, true
}

// HasUsersPerCardLimit returns a boolean if a field has been set.
func (o *LoyaltyProgram) HasUsersPerCardLimit() bool {
	if o != nil && o.UsersPerCardLimit != nil {
		return true
	}

	return false
}

// SetUsersPerCardLimit gets a reference to the given int32 and assigns it to the UsersPerCardLimit field.
func (o *LoyaltyProgram) SetUsersPerCardLimit(v int32) {
	o.UsersPerCardLimit = &v
}

// GetAccountID returns the AccountID field value
func (o *LoyaltyProgram) GetAccountID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountID
}

// GetAccountIDOk returns a tuple with the AccountID field value
// and a boolean to check if the value has been set.
func (o *LoyaltyProgram) GetAccountIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountID, true
}

// SetAccountID sets field value
func (o *LoyaltyProgram) SetAccountID(v int32) {
	o.AccountID = v
}

// GetName returns the Name field value
func (o *LoyaltyProgram) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LoyaltyProgram) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LoyaltyProgram) SetName(v string) {
	o.Name = v
}

// GetTiers returns the Tiers field value if set, zero value otherwise.
func (o *LoyaltyProgram) GetTiers() []LoyaltyTier {
	if o == nil || o.Tiers == nil {
		var ret []LoyaltyTier
		return ret
	}
	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyProgram) GetTiersOk() ([]LoyaltyTier, bool) {
	if o == nil || o.Tiers == nil {
		return nil, false
	}
	return o.Tiers, true
}

// HasTiers returns a boolean if a field has been set.
func (o *LoyaltyProgram) HasTiers() bool {
	if o != nil && o.Tiers != nil {
		return true
	}

	return false
}

// SetTiers gets a reference to the given []LoyaltyTier and assigns it to the Tiers field.
func (o *LoyaltyProgram) SetTiers(v []LoyaltyTier) {
	o.Tiers = v
}

// GetTimezone returns the Timezone field value
func (o *LoyaltyProgram) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *LoyaltyProgram) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *LoyaltyProgram) SetTimezone(v string) {
	o.Timezone = v
}

// GetCardBased returns the CardBased field value
func (o *LoyaltyProgram) GetCardBased() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CardBased
}

// GetCardBasedOk returns a tuple with the CardBased field value
// and a boolean to check if the value has been set.
func (o *LoyaltyProgram) GetCardBasedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardBased, true
}

// SetCardBased sets field value
func (o *LoyaltyProgram) SetCardBased(v bool) {
	o.CardBased = v
}

func (o LoyaltyProgram) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["subscribedApplications"] = o.SubscribedApplications
	}
	if true {
		toSerialize["defaultValidity"] = o.DefaultValidity
	}
	if true {
		toSerialize["defaultPending"] = o.DefaultPending
	}
	if true {
		toSerialize["allowSubledger"] = o.AllowSubledger
	}
	if o.UsersPerCardLimit != nil {
		toSerialize["usersPerCardLimit"] = o.UsersPerCardLimit
	}
	if true {
		toSerialize["accountID"] = o.AccountID
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Tiers != nil {
		toSerialize["tiers"] = o.Tiers
	}
	if true {
		toSerialize["timezone"] = o.Timezone
	}
	if true {
		toSerialize["cardBased"] = o.CardBased
	}
	return json.Marshal(toSerialize)
}

type NullableLoyaltyProgram struct {
	value *LoyaltyProgram
	isSet bool
}

func (v NullableLoyaltyProgram) Get() *LoyaltyProgram {
	return v.value
}

func (v *NullableLoyaltyProgram) Set(val *LoyaltyProgram) {
	v.value = val
	v.isSet = true
}

func (v NullableLoyaltyProgram) IsSet() bool {
	return v.isSet
}

func (v *NullableLoyaltyProgram) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoyaltyProgram(val *LoyaltyProgram) *NullableLoyaltyProgram {
	return &NullableLoyaltyProgram{value: val, isSet: true}
}

func (v NullableLoyaltyProgram) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoyaltyProgram) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


