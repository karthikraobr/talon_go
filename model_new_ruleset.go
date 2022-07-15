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

// NewRuleset struct for NewRuleset
type NewRuleset struct {
	// Set of rules to apply.
	Rules []Rule `json:"rules"`
	// An array that provides objects with variable names (name) and talang expressions to whose result they are bound (expression) during rule evaluation. The order of the evaluation is decided by the position in the array.
	Bindings []Binding `json:"bindings"`
	// The version of the rulebuilder used to create this ruleset.
	RbVersion *string `json:"rbVersion,omitempty"`
	// Indicates whether this created ruleset should be activated for the campaign that owns it.
	Activate *bool `json:"activate,omitempty"`
}

// NewNewRuleset instantiates a new NewRuleset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewRuleset(rules []Rule, bindings []Binding) *NewRuleset {
	this := NewRuleset{}
	this.Rules = rules
	this.Bindings = bindings
	return &this
}

// NewNewRulesetWithDefaults instantiates a new NewRuleset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewRulesetWithDefaults() *NewRuleset {
	this := NewRuleset{}
	return &this
}

// GetRules returns the Rules field value
func (o *NewRuleset) GetRules() []Rule {
	if o == nil {
		var ret []Rule
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *NewRuleset) GetRulesOk() ([]Rule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *NewRuleset) SetRules(v []Rule) {
	o.Rules = v
}

// GetBindings returns the Bindings field value
func (o *NewRuleset) GetBindings() []Binding {
	if o == nil {
		var ret []Binding
		return ret
	}

	return o.Bindings
}

// GetBindingsOk returns a tuple with the Bindings field value
// and a boolean to check if the value has been set.
func (o *NewRuleset) GetBindingsOk() ([]Binding, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bindings, true
}

// SetBindings sets field value
func (o *NewRuleset) SetBindings(v []Binding) {
	o.Bindings = v
}

// GetRbVersion returns the RbVersion field value if set, zero value otherwise.
func (o *NewRuleset) GetRbVersion() string {
	if o == nil || o.RbVersion == nil {
		var ret string
		return ret
	}
	return *o.RbVersion
}

// GetRbVersionOk returns a tuple with the RbVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewRuleset) GetRbVersionOk() (*string, bool) {
	if o == nil || o.RbVersion == nil {
		return nil, false
	}
	return o.RbVersion, true
}

// HasRbVersion returns a boolean if a field has been set.
func (o *NewRuleset) HasRbVersion() bool {
	if o != nil && o.RbVersion != nil {
		return true
	}

	return false
}

// SetRbVersion gets a reference to the given string and assigns it to the RbVersion field.
func (o *NewRuleset) SetRbVersion(v string) {
	o.RbVersion = &v
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *NewRuleset) GetActivate() bool {
	if o == nil || o.Activate == nil {
		var ret bool
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewRuleset) GetActivateOk() (*bool, bool) {
	if o == nil || o.Activate == nil {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *NewRuleset) HasActivate() bool {
	if o != nil && o.Activate != nil {
		return true
	}

	return false
}

// SetActivate gets a reference to the given bool and assigns it to the Activate field.
func (o *NewRuleset) SetActivate(v bool) {
	o.Activate = &v
}

func (o NewRuleset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	if true {
		toSerialize["bindings"] = o.Bindings
	}
	if o.RbVersion != nil {
		toSerialize["rbVersion"] = o.RbVersion
	}
	if o.Activate != nil {
		toSerialize["activate"] = o.Activate
	}
	return json.Marshal(toSerialize)
}

type NullableNewRuleset struct {
	value *NewRuleset
	isSet bool
}

func (v NullableNewRuleset) Get() *NewRuleset {
	return v.value
}

func (v *NullableNewRuleset) Set(val *NewRuleset) {
	v.value = val
	v.isSet = true
}

func (v NullableNewRuleset) IsSet() bool {
	return v.isSet
}

func (v *NullableNewRuleset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewRuleset(val *NewRuleset) *NullableNewRuleset {
	return &NullableNewRuleset{value: val, isSet: true}
}

func (v NullableNewRuleset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewRuleset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

