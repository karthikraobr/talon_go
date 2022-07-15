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

// TemplateArgDef 
type TemplateArgDef struct {
	// The type of value this argument expects.
	Type string `json:"type"`
	// A campaigner-friendly description of the argument, this will also be shown in the rule editor.
	Description string `json:"description"`
	// A campaigner friendly name for the argument, this will be shown in the rule editor.
	Title string `json:"title"`
	// Arbitrary metadata that may be used to render an input for this argument.
	Ui map[string]interface{} `json:"ui"`
}

// NewTemplateArgDef instantiates a new TemplateArgDef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateArgDef(type_ string, description string, title string, ui map[string]interface{}) *TemplateArgDef {
	this := TemplateArgDef{}
	this.Type = type_
	this.Description = description
	this.Title = title
	this.Ui = ui
	return &this
}

// NewTemplateArgDefWithDefaults instantiates a new TemplateArgDef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateArgDefWithDefaults() *TemplateArgDef {
	this := TemplateArgDef{}
	return &this
}

// GetType returns the Type field value
func (o *TemplateArgDef) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TemplateArgDef) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TemplateArgDef) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value
func (o *TemplateArgDef) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TemplateArgDef) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TemplateArgDef) SetDescription(v string) {
	o.Description = v
}

// GetTitle returns the Title field value
func (o *TemplateArgDef) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *TemplateArgDef) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *TemplateArgDef) SetTitle(v string) {
	o.Title = v
}

// GetUi returns the Ui field value
func (o *TemplateArgDef) GetUi() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Ui
}

// GetUiOk returns a tuple with the Ui field value
// and a boolean to check if the value has been set.
func (o *TemplateArgDef) GetUiOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ui, true
}

// SetUi sets field value
func (o *TemplateArgDef) SetUi(v map[string]interface{}) {
	o.Ui = v
}

func (o TemplateArgDef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["ui"] = o.Ui
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateArgDef struct {
	value *TemplateArgDef
	isSet bool
}

func (v NullableTemplateArgDef) Get() *TemplateArgDef {
	return v.value
}

func (v *NullableTemplateArgDef) Set(val *TemplateArgDef) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateArgDef) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateArgDef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateArgDef(val *TemplateArgDef) *NullableTemplateArgDef {
	return &NullableTemplateArgDef{value: val, isSet: true}
}

func (v NullableTemplateArgDef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateArgDef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


