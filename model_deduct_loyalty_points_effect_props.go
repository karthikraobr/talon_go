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

// DeductLoyaltyPointsEffectProps The properties specific to the \"deductLoyaltyPoints\" effect. This gets triggered whenever a validated rule contained a condition to only trigger when the given number of loyalty points could be deduced. These points are automatically stored and managed inside Talon.One.
type DeductLoyaltyPointsEffectProps struct {
	// The title of the rule that contained triggered this points deduction.
	RuleTitle string `json:"ruleTitle"`
	// The ID of the loyalty program where these points were added.
	ProgramId int32 `json:"programId"`
	// The ID of the subledger within the loyalty program where these points were added.
	SubLedgerId string `json:"subLedgerId"`
	// The amount of points that were deducted.
	Value float32 `json:"value"`
	// The identifier of this deduction in the loyalty ledger.
	TransactionUUID string `json:"transactionUUID"`
	// The name property gets one of the following two values. It can be the loyalty program name or it can represent a reason for the respective deduction of loyalty points. The latter is an optional value defined in a deduction rule. 
	Name string `json:"name"`
	// The card on which these points were added.
	CardIdentifier *string `json:"cardIdentifier,omitempty"`
}

// NewDeductLoyaltyPointsEffectProps instantiates a new DeductLoyaltyPointsEffectProps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeductLoyaltyPointsEffectProps(ruleTitle string, programId int32, subLedgerId string, value float32, transactionUUID string, name string) *DeductLoyaltyPointsEffectProps {
	this := DeductLoyaltyPointsEffectProps{}
	this.RuleTitle = ruleTitle
	this.ProgramId = programId
	this.SubLedgerId = subLedgerId
	this.Value = value
	this.TransactionUUID = transactionUUID
	this.Name = name
	return &this
}

// NewDeductLoyaltyPointsEffectPropsWithDefaults instantiates a new DeductLoyaltyPointsEffectProps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeductLoyaltyPointsEffectPropsWithDefaults() *DeductLoyaltyPointsEffectProps {
	this := DeductLoyaltyPointsEffectProps{}
	return &this
}

// GetRuleTitle returns the RuleTitle field value
func (o *DeductLoyaltyPointsEffectProps) GetRuleTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleTitle
}

// GetRuleTitleOk returns a tuple with the RuleTitle field value
// and a boolean to check if the value has been set.
func (o *DeductLoyaltyPointsEffectProps) GetRuleTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleTitle, true
}

// SetRuleTitle sets field value
func (o *DeductLoyaltyPointsEffectProps) SetRuleTitle(v string) {
	o.RuleTitle = v
}

// GetProgramId returns the ProgramId field value
func (o *DeductLoyaltyPointsEffectProps) GetProgramId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProgramId
}

// GetProgramIdOk returns a tuple with the ProgramId field value
// and a boolean to check if the value has been set.
func (o *DeductLoyaltyPointsEffectProps) GetProgramIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProgramId, true
}

// SetProgramId sets field value
func (o *DeductLoyaltyPointsEffectProps) SetProgramId(v int32) {
	o.ProgramId = v
}

// GetSubLedgerId returns the SubLedgerId field value
func (o *DeductLoyaltyPointsEffectProps) GetSubLedgerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubLedgerId
}

// GetSubLedgerIdOk returns a tuple with the SubLedgerId field value
// and a boolean to check if the value has been set.
func (o *DeductLoyaltyPointsEffectProps) GetSubLedgerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubLedgerId, true
}

// SetSubLedgerId sets field value
func (o *DeductLoyaltyPointsEffectProps) SetSubLedgerId(v string) {
	o.SubLedgerId = v
}

// GetValue returns the Value field value
func (o *DeductLoyaltyPointsEffectProps) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *DeductLoyaltyPointsEffectProps) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *DeductLoyaltyPointsEffectProps) SetValue(v float32) {
	o.Value = v
}

// GetTransactionUUID returns the TransactionUUID field value
func (o *DeductLoyaltyPointsEffectProps) GetTransactionUUID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionUUID
}

// GetTransactionUUIDOk returns a tuple with the TransactionUUID field value
// and a boolean to check if the value has been set.
func (o *DeductLoyaltyPointsEffectProps) GetTransactionUUIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionUUID, true
}

// SetTransactionUUID sets field value
func (o *DeductLoyaltyPointsEffectProps) SetTransactionUUID(v string) {
	o.TransactionUUID = v
}

// GetName returns the Name field value
func (o *DeductLoyaltyPointsEffectProps) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeductLoyaltyPointsEffectProps) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeductLoyaltyPointsEffectProps) SetName(v string) {
	o.Name = v
}

// GetCardIdentifier returns the CardIdentifier field value if set, zero value otherwise.
func (o *DeductLoyaltyPointsEffectProps) GetCardIdentifier() string {
	if o == nil || o.CardIdentifier == nil {
		var ret string
		return ret
	}
	return *o.CardIdentifier
}

// GetCardIdentifierOk returns a tuple with the CardIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeductLoyaltyPointsEffectProps) GetCardIdentifierOk() (*string, bool) {
	if o == nil || o.CardIdentifier == nil {
		return nil, false
	}
	return o.CardIdentifier, true
}

// HasCardIdentifier returns a boolean if a field has been set.
func (o *DeductLoyaltyPointsEffectProps) HasCardIdentifier() bool {
	if o != nil && o.CardIdentifier != nil {
		return true
	}

	return false
}

// SetCardIdentifier gets a reference to the given string and assigns it to the CardIdentifier field.
func (o *DeductLoyaltyPointsEffectProps) SetCardIdentifier(v string) {
	o.CardIdentifier = &v
}

func (o DeductLoyaltyPointsEffectProps) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ruleTitle"] = o.RuleTitle
	}
	if true {
		toSerialize["programId"] = o.ProgramId
	}
	if true {
		toSerialize["subLedgerId"] = o.SubLedgerId
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["transactionUUID"] = o.TransactionUUID
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.CardIdentifier != nil {
		toSerialize["cardIdentifier"] = o.CardIdentifier
	}
	return json.Marshal(toSerialize)
}

type NullableDeductLoyaltyPointsEffectProps struct {
	value *DeductLoyaltyPointsEffectProps
	isSet bool
}

func (v NullableDeductLoyaltyPointsEffectProps) Get() *DeductLoyaltyPointsEffectProps {
	return v.value
}

func (v *NullableDeductLoyaltyPointsEffectProps) Set(val *DeductLoyaltyPointsEffectProps) {
	v.value = val
	v.isSet = true
}

func (v NullableDeductLoyaltyPointsEffectProps) IsSet() bool {
	return v.isSet
}

func (v *NullableDeductLoyaltyPointsEffectProps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeductLoyaltyPointsEffectProps(val *DeductLoyaltyPointsEffectProps) *NullableDeductLoyaltyPointsEffectProps {
	return &NullableDeductLoyaltyPointsEffectProps{value: val, isSet: true}
}

func (v NullableDeductLoyaltyPointsEffectProps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeductLoyaltyPointsEffectProps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

