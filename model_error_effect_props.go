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

// ErrorEffectProps Whenever an error occurred during evaluation, we return an error effect. This should never happen for rules created in the rule builder.
type ErrorEffectProps struct {
	// The error message.
	Message string `json:"message"`
}

// NewErrorEffectProps instantiates a new ErrorEffectProps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorEffectProps(message string) *ErrorEffectProps {
	this := ErrorEffectProps{}
	this.Message = message
	return &this
}

// NewErrorEffectPropsWithDefaults instantiates a new ErrorEffectProps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorEffectPropsWithDefaults() *ErrorEffectProps {
	this := ErrorEffectProps{}
	return &this
}

// GetMessage returns the Message field value
func (o *ErrorEffectProps) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorEffectProps) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorEffectProps) SetMessage(v string) {
	o.Message = v
}

func (o ErrorEffectProps) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableErrorEffectProps struct {
	value *ErrorEffectProps
	isSet bool
}

func (v NullableErrorEffectProps) Get() *ErrorEffectProps {
	return v.value
}

func (v *NullableErrorEffectProps) Set(val *ErrorEffectProps) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorEffectProps) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorEffectProps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorEffectProps(val *ErrorEffectProps) *NullableErrorEffectProps {
	return &NullableErrorEffectProps{value: val, isSet: true}
}

func (v NullableErrorEffectProps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorEffectProps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


