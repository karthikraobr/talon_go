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

// LoyaltyLedgerEntry A single row of the ledger, describing one addition or deduction.
type LoyaltyLedgerEntry struct {
	Created time.Time `json:"created"`
	ProgramID int32 `json:"programID"`
	CustomerProfileID *string `json:"customerProfileID,omitempty"`
	CardID *int32 `json:"cardID,omitempty"`
	CustomerSessionID *string `json:"customerSessionID,omitempty"`
	EventID *int32 `json:"eventID,omitempty"`
	// The type of the ledger transaction. Possible values are addition, subtraction, expire or expiring (for expiring points ledgers) 
	Type string `json:"type"`
	Amount float32 `json:"amount"`
	StartDate *time.Time `json:"startDate,omitempty"`
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// A name referencing the condition or effect that added this entry, or the specific name provided in an API call.
	Name string `json:"name"`
	// This specifies if we are adding loyalty points to the main ledger or a subledger.
	SubLedgerID string `json:"subLedgerID"`
	// This is the ID of the user who created this entry, if the addition or subtraction was done manually.
	UserID *int32 `json:"userID,omitempty"`
}

// NewLoyaltyLedgerEntry instantiates a new LoyaltyLedgerEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoyaltyLedgerEntry(created time.Time, programID int32, type_ string, amount float32, name string, subLedgerID string) *LoyaltyLedgerEntry {
	this := LoyaltyLedgerEntry{}
	this.Created = created
	this.ProgramID = programID
	this.Type = type_
	this.Amount = amount
	this.Name = name
	this.SubLedgerID = subLedgerID
	return &this
}

// NewLoyaltyLedgerEntryWithDefaults instantiates a new LoyaltyLedgerEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoyaltyLedgerEntryWithDefaults() *LoyaltyLedgerEntry {
	this := LoyaltyLedgerEntry{}
	return &this
}

// GetCreated returns the Created field value
func (o *LoyaltyLedgerEntry) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerEntry) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *LoyaltyLedgerEntry) SetCreated(v time.Time) {
	o.Created = v
}

// GetProgramID returns the ProgramID field value
func (o *LoyaltyLedgerEntry) GetProgramID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProgramID
}

// GetProgramIDOk returns a tuple with the ProgramID field value
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerEntry) GetProgramIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProgramID, true
}

// SetProgramID sets field value
func (o *LoyaltyLedgerEntry) SetProgramID(v int32) {
	o.ProgramID = v
}

// GetCustomerProfileID returns the CustomerProfileID field value if set, zero value otherwise.
func (o *LoyaltyLedgerEntry) GetCustomerProfileID() string {
	if o == nil || o.CustomerProfileID == nil {
		var ret string
		return ret
	}
	return *o.CustomerProfileID
}

// GetCustomerProfileIDOk returns a tuple with the CustomerProfileID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerEntry) GetCustomerProfileIDOk() (*string, bool) {
	if o == nil || o.CustomerProfileID == nil {
		return nil, false
	}
	return o.CustomerProfileID, true
}

// HasCustomerProfileID returns a boolean if a field has been set.
func (o *LoyaltyLedgerEntry) HasCustomerProfileID() bool {
	if o != nil && o.CustomerProfileID != nil {
		return true
	}

	return false
}

// SetCustomerProfileID gets a reference to the given string and assigns it to the CustomerProfileID field.
func (o *LoyaltyLedgerEntry) SetCustomerProfileID(v string) {
	o.CustomerProfileID = &v
}

// GetCardID returns the CardID field value if set, zero value otherwise.
func (o *LoyaltyLedgerEntry) GetCardID() int32 {
	if o == nil || o.CardID == nil {
		var ret int32
		return ret
	}
	return *o.CardID
}

// GetCardIDOk returns a tuple with the CardID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerEntry) GetCardIDOk() (*int32, bool) {
	if o == nil || o.CardID == nil {
		return nil, false
	}
	return o.CardID, true
}

// HasCardID returns a boolean if a field has been set.
func (o *LoyaltyLedgerEntry) HasCardID() bool {
	if o != nil && o.CardID != nil {
		return true
	}

	return false
}

// SetCardID gets a reference to the given int32 and assigns it to the CardID field.
func (o *LoyaltyLedgerEntry) SetCardID(v int32) {
	o.CardID = &v
}

// GetCustomerSessionID returns the CustomerSessionID field value if set, zero value otherwise.
func (o *LoyaltyLedgerEntry) GetCustomerSessionID() string {
	if o == nil || o.CustomerSessionID == nil {
		var ret string
		return ret
	}
	return *o.CustomerSessionID
}

// GetCustomerSessionIDOk returns a tuple with the CustomerSessionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerEntry) GetCustomerSessionIDOk() (*string, bool) {
	if o == nil || o.CustomerSessionID == nil {
		return nil, false
	}
	return o.CustomerSessionID, true
}

// HasCustomerSessionID returns a boolean if a field has been set.
func (o *LoyaltyLedgerEntry) HasCustomerSessionID() bool {
	if o != nil && o.CustomerSessionID != nil {
		return true
	}

	return false
}

// SetCustomerSessionID gets a reference to the given string and assigns it to the CustomerSessionID field.
func (o *LoyaltyLedgerEntry) SetCustomerSessionID(v string) {
	o.CustomerSessionID = &v
}

// GetEventID returns the EventID field value if set, zero value otherwise.
func (o *LoyaltyLedgerEntry) GetEventID() int32 {
	if o == nil || o.EventID == nil {
		var ret int32
		return ret
	}
	return *o.EventID
}

// GetEventIDOk returns a tuple with the EventID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerEntry) GetEventIDOk() (*int32, bool) {
	if o == nil || o.EventID == nil {
		return nil, false
	}
	return o.EventID, true
}

// HasEventID returns a boolean if a field has been set.
func (o *LoyaltyLedgerEntry) HasEventID() bool {
	if o != nil && o.EventID != nil {
		return true
	}

	return false
}

// SetEventID gets a reference to the given int32 and assigns it to the EventID field.
func (o *LoyaltyLedgerEntry) SetEventID(v int32) {
	o.EventID = &v
}

// GetType returns the Type field value
func (o *LoyaltyLedgerEntry) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerEntry) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LoyaltyLedgerEntry) SetType(v string) {
	o.Type = v
}

// GetAmount returns the Amount field value
func (o *LoyaltyLedgerEntry) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerEntry) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *LoyaltyLedgerEntry) SetAmount(v float32) {
	o.Amount = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *LoyaltyLedgerEntry) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerEntry) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *LoyaltyLedgerEntry) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *LoyaltyLedgerEntry) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *LoyaltyLedgerEntry) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerEntry) GetExpiryDateOk() (*time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *LoyaltyLedgerEntry) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *LoyaltyLedgerEntry) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetName returns the Name field value
func (o *LoyaltyLedgerEntry) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerEntry) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LoyaltyLedgerEntry) SetName(v string) {
	o.Name = v
}

// GetSubLedgerID returns the SubLedgerID field value
func (o *LoyaltyLedgerEntry) GetSubLedgerID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubLedgerID
}

// GetSubLedgerIDOk returns a tuple with the SubLedgerID field value
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerEntry) GetSubLedgerIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubLedgerID, true
}

// SetSubLedgerID sets field value
func (o *LoyaltyLedgerEntry) SetSubLedgerID(v string) {
	o.SubLedgerID = v
}

// GetUserID returns the UserID field value if set, zero value otherwise.
func (o *LoyaltyLedgerEntry) GetUserID() int32 {
	if o == nil || o.UserID == nil {
		var ret int32
		return ret
	}
	return *o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerEntry) GetUserIDOk() (*int32, bool) {
	if o == nil || o.UserID == nil {
		return nil, false
	}
	return o.UserID, true
}

// HasUserID returns a boolean if a field has been set.
func (o *LoyaltyLedgerEntry) HasUserID() bool {
	if o != nil && o.UserID != nil {
		return true
	}

	return false
}

// SetUserID gets a reference to the given int32 and assigns it to the UserID field.
func (o *LoyaltyLedgerEntry) SetUserID(v int32) {
	o.UserID = &v
}

func (o LoyaltyLedgerEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["programID"] = o.ProgramID
	}
	if o.CustomerProfileID != nil {
		toSerialize["customerProfileID"] = o.CustomerProfileID
	}
	if o.CardID != nil {
		toSerialize["cardID"] = o.CardID
	}
	if o.CustomerSessionID != nil {
		toSerialize["customerSessionID"] = o.CustomerSessionID
	}
	if o.EventID != nil {
		toSerialize["eventID"] = o.EventID
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.ExpiryDate != nil {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["subLedgerID"] = o.SubLedgerID
	}
	if o.UserID != nil {
		toSerialize["userID"] = o.UserID
	}
	return json.Marshal(toSerialize)
}

type NullableLoyaltyLedgerEntry struct {
	value *LoyaltyLedgerEntry
	isSet bool
}

func (v NullableLoyaltyLedgerEntry) Get() *LoyaltyLedgerEntry {
	return v.value
}

func (v *NullableLoyaltyLedgerEntry) Set(val *LoyaltyLedgerEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableLoyaltyLedgerEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableLoyaltyLedgerEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoyaltyLedgerEntry(val *LoyaltyLedgerEntry) *NullableLoyaltyLedgerEntry {
	return &NullableLoyaltyLedgerEntry{value: val, isSet: true}
}

func (v NullableLoyaltyLedgerEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoyaltyLedgerEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


