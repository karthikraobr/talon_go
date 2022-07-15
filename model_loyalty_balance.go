/*
Integration API reference docs

Use the Integration API to push data to and retrieve data from Talon.One in real time. For more background information about this API, see [Integration API overview](/docs/dev/integration-api/overview)  For example, use this API to share shopping cart information as a session with Talon.One and evaluate promotion rules. You can also create custom events to track specific actions that do not fit into the session data model.  Ensure you [authenticate](#section/Authentication) to make requests to the API.  <div class=\"redoc-section\">   <p class=\"title\">Are you looking for a different API?</p>    If you need the API to:    - Interact with the Campaign Manager for backoffice operations, see [the Management API reference docs](https://docs.talon.one/management-api).   - Integrate with Talon.One from a CEP or CDP platform, see [the Third-party API reference docs](https://docs.talon.one/third-party-api).  </div>  # Authentication  <SecurityDefinitions /> 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package talon

import (
	"encoding/json"
)

// LoyaltyBalance Point balance of a ledger in the Loyalty Program.
type LoyaltyBalance struct {
	// Total amount of points awarded to this customer and available to spend.
	ActivePoints *float32 `json:"activePoints,omitempty"`
	// Total amount of points awarded to this customer but not available until their start date.
	PendingPoints *float32 `json:"pendingPoints,omitempty"`
	// Total amount of points already spent by this customer.
	SpentPoints *float32 `json:"spentPoints,omitempty"`
	// Total amount of points awarded but never redeemed. They cannot be used anymore.
	ExpiredPoints *float32 `json:"expiredPoints,omitempty"`
}

// NewLoyaltyBalance instantiates a new LoyaltyBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoyaltyBalance() *LoyaltyBalance {
	this := LoyaltyBalance{}
	return &this
}

// NewLoyaltyBalanceWithDefaults instantiates a new LoyaltyBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoyaltyBalanceWithDefaults() *LoyaltyBalance {
	this := LoyaltyBalance{}
	return &this
}

// GetActivePoints returns the ActivePoints field value if set, zero value otherwise.
func (o *LoyaltyBalance) GetActivePoints() float32 {
	if o == nil || o.ActivePoints == nil {
		var ret float32
		return ret
	}
	return *o.ActivePoints
}

// GetActivePointsOk returns a tuple with the ActivePoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyBalance) GetActivePointsOk() (*float32, bool) {
	if o == nil || o.ActivePoints == nil {
		return nil, false
	}
	return o.ActivePoints, true
}

// HasActivePoints returns a boolean if a field has been set.
func (o *LoyaltyBalance) HasActivePoints() bool {
	if o != nil && o.ActivePoints != nil {
		return true
	}

	return false
}

// SetActivePoints gets a reference to the given float32 and assigns it to the ActivePoints field.
func (o *LoyaltyBalance) SetActivePoints(v float32) {
	o.ActivePoints = &v
}

// GetPendingPoints returns the PendingPoints field value if set, zero value otherwise.
func (o *LoyaltyBalance) GetPendingPoints() float32 {
	if o == nil || o.PendingPoints == nil {
		var ret float32
		return ret
	}
	return *o.PendingPoints
}

// GetPendingPointsOk returns a tuple with the PendingPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyBalance) GetPendingPointsOk() (*float32, bool) {
	if o == nil || o.PendingPoints == nil {
		return nil, false
	}
	return o.PendingPoints, true
}

// HasPendingPoints returns a boolean if a field has been set.
func (o *LoyaltyBalance) HasPendingPoints() bool {
	if o != nil && o.PendingPoints != nil {
		return true
	}

	return false
}

// SetPendingPoints gets a reference to the given float32 and assigns it to the PendingPoints field.
func (o *LoyaltyBalance) SetPendingPoints(v float32) {
	o.PendingPoints = &v
}

// GetSpentPoints returns the SpentPoints field value if set, zero value otherwise.
func (o *LoyaltyBalance) GetSpentPoints() float32 {
	if o == nil || o.SpentPoints == nil {
		var ret float32
		return ret
	}
	return *o.SpentPoints
}

// GetSpentPointsOk returns a tuple with the SpentPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyBalance) GetSpentPointsOk() (*float32, bool) {
	if o == nil || o.SpentPoints == nil {
		return nil, false
	}
	return o.SpentPoints, true
}

// HasSpentPoints returns a boolean if a field has been set.
func (o *LoyaltyBalance) HasSpentPoints() bool {
	if o != nil && o.SpentPoints != nil {
		return true
	}

	return false
}

// SetSpentPoints gets a reference to the given float32 and assigns it to the SpentPoints field.
func (o *LoyaltyBalance) SetSpentPoints(v float32) {
	o.SpentPoints = &v
}

// GetExpiredPoints returns the ExpiredPoints field value if set, zero value otherwise.
func (o *LoyaltyBalance) GetExpiredPoints() float32 {
	if o == nil || o.ExpiredPoints == nil {
		var ret float32
		return ret
	}
	return *o.ExpiredPoints
}

// GetExpiredPointsOk returns a tuple with the ExpiredPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyBalance) GetExpiredPointsOk() (*float32, bool) {
	if o == nil || o.ExpiredPoints == nil {
		return nil, false
	}
	return o.ExpiredPoints, true
}

// HasExpiredPoints returns a boolean if a field has been set.
func (o *LoyaltyBalance) HasExpiredPoints() bool {
	if o != nil && o.ExpiredPoints != nil {
		return true
	}

	return false
}

// SetExpiredPoints gets a reference to the given float32 and assigns it to the ExpiredPoints field.
func (o *LoyaltyBalance) SetExpiredPoints(v float32) {
	o.ExpiredPoints = &v
}

func (o LoyaltyBalance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActivePoints != nil {
		toSerialize["activePoints"] = o.ActivePoints
	}
	if o.PendingPoints != nil {
		toSerialize["pendingPoints"] = o.PendingPoints
	}
	if o.SpentPoints != nil {
		toSerialize["spentPoints"] = o.SpentPoints
	}
	if o.ExpiredPoints != nil {
		toSerialize["expiredPoints"] = o.ExpiredPoints
	}
	return json.Marshal(toSerialize)
}

type NullableLoyaltyBalance struct {
	value *LoyaltyBalance
	isSet bool
}

func (v NullableLoyaltyBalance) Get() *LoyaltyBalance {
	return v.value
}

func (v *NullableLoyaltyBalance) Set(val *LoyaltyBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableLoyaltyBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableLoyaltyBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoyaltyBalance(val *LoyaltyBalance) *NullableLoyaltyBalance {
	return &NullableLoyaltyBalance{value: val, isSet: true}
}

func (v NullableLoyaltyBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoyaltyBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

