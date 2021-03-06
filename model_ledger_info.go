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

// LedgerInfo 
type LedgerInfo struct {
	// Sum of currently active points.
	CurrentBalance float32 `json:"currentBalance"`
	// Sum of pending points.
	PendingBalance float32 `json:"pendingBalance"`
	// Sum of expired points.
	ExpiredBalance float32 `json:"expiredBalance"`
	// Sum of spent points.
	SpentBalance float32 `json:"spentBalance"`
	// Sum of currently active points, including points added and deducted in open sessions.
	TentativeCurrentBalance float32 `json:"tentativeCurrentBalance"`
	CurrentTier *Tier `json:"currentTier,omitempty"`
	// Points required to move up a tier.
	PointsToNextTier *float32 `json:"pointsToNextTier,omitempty"`
	Projection *LoyaltyProjection `json:"projection,omitempty"`
}

// NewLedgerInfo instantiates a new LedgerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLedgerInfo(currentBalance float32, pendingBalance float32, expiredBalance float32, spentBalance float32, tentativeCurrentBalance float32) *LedgerInfo {
	this := LedgerInfo{}
	this.CurrentBalance = currentBalance
	this.PendingBalance = pendingBalance
	this.ExpiredBalance = expiredBalance
	this.SpentBalance = spentBalance
	this.TentativeCurrentBalance = tentativeCurrentBalance
	return &this
}

// NewLedgerInfoWithDefaults instantiates a new LedgerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLedgerInfoWithDefaults() *LedgerInfo {
	this := LedgerInfo{}
	return &this
}

// GetCurrentBalance returns the CurrentBalance field value
func (o *LedgerInfo) GetCurrentBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CurrentBalance
}

// GetCurrentBalanceOk returns a tuple with the CurrentBalance field value
// and a boolean to check if the value has been set.
func (o *LedgerInfo) GetCurrentBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentBalance, true
}

// SetCurrentBalance sets field value
func (o *LedgerInfo) SetCurrentBalance(v float32) {
	o.CurrentBalance = v
}

// GetPendingBalance returns the PendingBalance field value
func (o *LedgerInfo) GetPendingBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PendingBalance
}

// GetPendingBalanceOk returns a tuple with the PendingBalance field value
// and a boolean to check if the value has been set.
func (o *LedgerInfo) GetPendingBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingBalance, true
}

// SetPendingBalance sets field value
func (o *LedgerInfo) SetPendingBalance(v float32) {
	o.PendingBalance = v
}

// GetExpiredBalance returns the ExpiredBalance field value
func (o *LedgerInfo) GetExpiredBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ExpiredBalance
}

// GetExpiredBalanceOk returns a tuple with the ExpiredBalance field value
// and a boolean to check if the value has been set.
func (o *LedgerInfo) GetExpiredBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiredBalance, true
}

// SetExpiredBalance sets field value
func (o *LedgerInfo) SetExpiredBalance(v float32) {
	o.ExpiredBalance = v
}

// GetSpentBalance returns the SpentBalance field value
func (o *LedgerInfo) GetSpentBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SpentBalance
}

// GetSpentBalanceOk returns a tuple with the SpentBalance field value
// and a boolean to check if the value has been set.
func (o *LedgerInfo) GetSpentBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpentBalance, true
}

// SetSpentBalance sets field value
func (o *LedgerInfo) SetSpentBalance(v float32) {
	o.SpentBalance = v
}

// GetTentativeCurrentBalance returns the TentativeCurrentBalance field value
func (o *LedgerInfo) GetTentativeCurrentBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TentativeCurrentBalance
}

// GetTentativeCurrentBalanceOk returns a tuple with the TentativeCurrentBalance field value
// and a boolean to check if the value has been set.
func (o *LedgerInfo) GetTentativeCurrentBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TentativeCurrentBalance, true
}

// SetTentativeCurrentBalance sets field value
func (o *LedgerInfo) SetTentativeCurrentBalance(v float32) {
	o.TentativeCurrentBalance = v
}

// GetCurrentTier returns the CurrentTier field value if set, zero value otherwise.
func (o *LedgerInfo) GetCurrentTier() Tier {
	if o == nil || o.CurrentTier == nil {
		var ret Tier
		return ret
	}
	return *o.CurrentTier
}

// GetCurrentTierOk returns a tuple with the CurrentTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerInfo) GetCurrentTierOk() (*Tier, bool) {
	if o == nil || o.CurrentTier == nil {
		return nil, false
	}
	return o.CurrentTier, true
}

// HasCurrentTier returns a boolean if a field has been set.
func (o *LedgerInfo) HasCurrentTier() bool {
	if o != nil && o.CurrentTier != nil {
		return true
	}

	return false
}

// SetCurrentTier gets a reference to the given Tier and assigns it to the CurrentTier field.
func (o *LedgerInfo) SetCurrentTier(v Tier) {
	o.CurrentTier = &v
}

// GetPointsToNextTier returns the PointsToNextTier field value if set, zero value otherwise.
func (o *LedgerInfo) GetPointsToNextTier() float32 {
	if o == nil || o.PointsToNextTier == nil {
		var ret float32
		return ret
	}
	return *o.PointsToNextTier
}

// GetPointsToNextTierOk returns a tuple with the PointsToNextTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerInfo) GetPointsToNextTierOk() (*float32, bool) {
	if o == nil || o.PointsToNextTier == nil {
		return nil, false
	}
	return o.PointsToNextTier, true
}

// HasPointsToNextTier returns a boolean if a field has been set.
func (o *LedgerInfo) HasPointsToNextTier() bool {
	if o != nil && o.PointsToNextTier != nil {
		return true
	}

	return false
}

// SetPointsToNextTier gets a reference to the given float32 and assigns it to the PointsToNextTier field.
func (o *LedgerInfo) SetPointsToNextTier(v float32) {
	o.PointsToNextTier = &v
}

// GetProjection returns the Projection field value if set, zero value otherwise.
func (o *LedgerInfo) GetProjection() LoyaltyProjection {
	if o == nil || o.Projection == nil {
		var ret LoyaltyProjection
		return ret
	}
	return *o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerInfo) GetProjectionOk() (*LoyaltyProjection, bool) {
	if o == nil || o.Projection == nil {
		return nil, false
	}
	return o.Projection, true
}

// HasProjection returns a boolean if a field has been set.
func (o *LedgerInfo) HasProjection() bool {
	if o != nil && o.Projection != nil {
		return true
	}

	return false
}

// SetProjection gets a reference to the given LoyaltyProjection and assigns it to the Projection field.
func (o *LedgerInfo) SetProjection(v LoyaltyProjection) {
	o.Projection = &v
}

func (o LedgerInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["currentBalance"] = o.CurrentBalance
	}
	if true {
		toSerialize["pendingBalance"] = o.PendingBalance
	}
	if true {
		toSerialize["expiredBalance"] = o.ExpiredBalance
	}
	if true {
		toSerialize["spentBalance"] = o.SpentBalance
	}
	if true {
		toSerialize["tentativeCurrentBalance"] = o.TentativeCurrentBalance
	}
	if o.CurrentTier != nil {
		toSerialize["currentTier"] = o.CurrentTier
	}
	if o.PointsToNextTier != nil {
		toSerialize["pointsToNextTier"] = o.PointsToNextTier
	}
	if o.Projection != nil {
		toSerialize["projection"] = o.Projection
	}
	return json.Marshal(toSerialize)
}

type NullableLedgerInfo struct {
	value *LedgerInfo
	isSet bool
}

func (v NullableLedgerInfo) Get() *LedgerInfo {
	return v.value
}

func (v *NullableLedgerInfo) Set(val *LedgerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLedgerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLedgerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLedgerInfo(val *LedgerInfo) *NullableLedgerInfo {
	return &NullableLedgerInfo{value: val, isSet: true}
}

func (v NullableLedgerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLedgerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


