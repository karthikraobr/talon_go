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

// ErrorSource The source of the current error, exactly one of `pointer`, `parameter` or `line` will be defined. 
type ErrorSource struct {
	// Pointer to the path in the payload that caused this error.
	Pointer *string `json:"pointer,omitempty"`
	// Query parameter that caused this error.
	Parameter *string `json:"parameter,omitempty"`
	// Line number in uploaded multipart file that caused this error. 'N/A' if unknown.
	Line *string `json:"line,omitempty"`
	// Pointer to the resource that caused this error.
	Resource *string `json:"resource,omitempty"`
}

// NewErrorSource instantiates a new ErrorSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorSource() *ErrorSource {
	this := ErrorSource{}
	return &this
}

// NewErrorSourceWithDefaults instantiates a new ErrorSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorSourceWithDefaults() *ErrorSource {
	this := ErrorSource{}
	return &this
}

// GetPointer returns the Pointer field value if set, zero value otherwise.
func (o *ErrorSource) GetPointer() string {
	if o == nil || o.Pointer == nil {
		var ret string
		return ret
	}
	return *o.Pointer
}

// GetPointerOk returns a tuple with the Pointer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorSource) GetPointerOk() (*string, bool) {
	if o == nil || o.Pointer == nil {
		return nil, false
	}
	return o.Pointer, true
}

// HasPointer returns a boolean if a field has been set.
func (o *ErrorSource) HasPointer() bool {
	if o != nil && o.Pointer != nil {
		return true
	}

	return false
}

// SetPointer gets a reference to the given string and assigns it to the Pointer field.
func (o *ErrorSource) SetPointer(v string) {
	o.Pointer = &v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *ErrorSource) GetParameter() string {
	if o == nil || o.Parameter == nil {
		var ret string
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorSource) GetParameterOk() (*string, bool) {
	if o == nil || o.Parameter == nil {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *ErrorSource) HasParameter() bool {
	if o != nil && o.Parameter != nil {
		return true
	}

	return false
}

// SetParameter gets a reference to the given string and assigns it to the Parameter field.
func (o *ErrorSource) SetParameter(v string) {
	o.Parameter = &v
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *ErrorSource) GetLine() string {
	if o == nil || o.Line == nil {
		var ret string
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorSource) GetLineOk() (*string, bool) {
	if o == nil || o.Line == nil {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *ErrorSource) HasLine() bool {
	if o != nil && o.Line != nil {
		return true
	}

	return false
}

// SetLine gets a reference to the given string and assigns it to the Line field.
func (o *ErrorSource) SetLine(v string) {
	o.Line = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ErrorSource) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorSource) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ErrorSource) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *ErrorSource) SetResource(v string) {
	o.Resource = &v
}

func (o ErrorSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pointer != nil {
		toSerialize["pointer"] = o.Pointer
	}
	if o.Parameter != nil {
		toSerialize["parameter"] = o.Parameter
	}
	if o.Line != nil {
		toSerialize["line"] = o.Line
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullableErrorSource struct {
	value *ErrorSource
	isSet bool
}

func (v NullableErrorSource) Get() *ErrorSource {
	return v.value
}

func (v *NullableErrorSource) Set(val *ErrorSource) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorSource) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorSource(val *ErrorSource) *NullableErrorSource {
	return &NullableErrorSource{value: val, isSet: true}
}

func (v NullableErrorSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


