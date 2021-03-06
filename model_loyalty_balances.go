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

// LoyaltyBalances List of loyalty balances for a ledger and its subledgers.
type LoyaltyBalances struct {
	Balance *LoyaltyBalance `json:"balance,omitempty"`
	// Map of the loyalty balances of the subledgers of a ledger.
	SubledgerBalances *map[string]LoyaltyBalance `json:"subledgerBalances,omitempty"`
}

// NewLoyaltyBalances instantiates a new LoyaltyBalances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoyaltyBalances() *LoyaltyBalances {
	this := LoyaltyBalances{}
	return &this
}

// NewLoyaltyBalancesWithDefaults instantiates a new LoyaltyBalances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoyaltyBalancesWithDefaults() *LoyaltyBalances {
	this := LoyaltyBalances{}
	return &this
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *LoyaltyBalances) GetBalance() LoyaltyBalance {
	if o == nil || o.Balance == nil {
		var ret LoyaltyBalance
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyBalances) GetBalanceOk() (*LoyaltyBalance, bool) {
	if o == nil || o.Balance == nil {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *LoyaltyBalances) HasBalance() bool {
	if o != nil && o.Balance != nil {
		return true
	}

	return false
}

// SetBalance gets a reference to the given LoyaltyBalance and assigns it to the Balance field.
func (o *LoyaltyBalances) SetBalance(v LoyaltyBalance) {
	o.Balance = &v
}

// GetSubledgerBalances returns the SubledgerBalances field value if set, zero value otherwise.
func (o *LoyaltyBalances) GetSubledgerBalances() map[string]LoyaltyBalance {
	if o == nil || o.SubledgerBalances == nil {
		var ret map[string]LoyaltyBalance
		return ret
	}
	return *o.SubledgerBalances
}

// GetSubledgerBalancesOk returns a tuple with the SubledgerBalances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyBalances) GetSubledgerBalancesOk() (*map[string]LoyaltyBalance, bool) {
	if o == nil || o.SubledgerBalances == nil {
		return nil, false
	}
	return o.SubledgerBalances, true
}

// HasSubledgerBalances returns a boolean if a field has been set.
func (o *LoyaltyBalances) HasSubledgerBalances() bool {
	if o != nil && o.SubledgerBalances != nil {
		return true
	}

	return false
}

// SetSubledgerBalances gets a reference to the given map[string]LoyaltyBalance and assigns it to the SubledgerBalances field.
func (o *LoyaltyBalances) SetSubledgerBalances(v map[string]LoyaltyBalance) {
	o.SubledgerBalances = &v
}

func (o LoyaltyBalances) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Balance != nil {
		toSerialize["balance"] = o.Balance
	}
	if o.SubledgerBalances != nil {
		toSerialize["subledgerBalances"] = o.SubledgerBalances
	}
	return json.Marshal(toSerialize)
}

type NullableLoyaltyBalances struct {
	value *LoyaltyBalances
	isSet bool
}

func (v NullableLoyaltyBalances) Get() *LoyaltyBalances {
	return v.value
}

func (v *NullableLoyaltyBalances) Set(val *LoyaltyBalances) {
	v.value = val
	v.isSet = true
}

func (v NullableLoyaltyBalances) IsSet() bool {
	return v.isSet
}

func (v *NullableLoyaltyBalances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoyaltyBalances(val *LoyaltyBalances) *NullableLoyaltyBalances {
	return &NullableLoyaltyBalances{value: val, isSet: true}
}

func (v NullableLoyaltyBalances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoyaltyBalances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


