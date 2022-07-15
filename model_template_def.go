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

// TemplateDef 
type TemplateDef struct {
	// Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`
	// Campaigner-friendly name for the template that will be shown in the rule editor.
	Title string `json:"title"`
	// A short description of the template that will be shown in the rule editor.
	Description string `json:"description"`
	// Extended help text for the template.
	Help string `json:"help"`
	// Used for grouping templates in the rule editor sidebar.
	Category string `json:"category"`
	// A Talang expression that contains variable bindings referring to args.
	Expr []map[string]interface{} `json:"expr"`
	// An array of argument definitions.
	Args []TemplateArgDef `json:"args"`
	// A flag to control exposure in Rule Builder.
	Expose *bool `json:"expose,omitempty"`
	// The template name used in Talang.
	Name string `json:"name"`
}

// NewTemplateDef instantiates a new TemplateDef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateDef(id int32, created time.Time, applicationId int32, title string, description string, help string, category string, expr []map[string]interface{}, args []TemplateArgDef, name string) *TemplateDef {
	this := TemplateDef{}
	this.Id = id
	this.Created = created
	this.ApplicationId = applicationId
	this.Title = title
	this.Description = description
	this.Help = help
	this.Category = category
	this.Expr = expr
	this.Args = args
	var expose bool = false
	this.Expose = &expose
	this.Name = name
	return &this
}

// NewTemplateDefWithDefaults instantiates a new TemplateDef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateDefWithDefaults() *TemplateDef {
	this := TemplateDef{}
	var expose bool = false
	this.Expose = &expose
	return &this
}

// GetId returns the Id field value
func (o *TemplateDef) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TemplateDef) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TemplateDef) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *TemplateDef) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *TemplateDef) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *TemplateDef) SetCreated(v time.Time) {
	o.Created = v
}

// GetApplicationId returns the ApplicationId field value
func (o *TemplateDef) GetApplicationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *TemplateDef) GetApplicationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *TemplateDef) SetApplicationId(v int32) {
	o.ApplicationId = v
}

// GetTitle returns the Title field value
func (o *TemplateDef) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *TemplateDef) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *TemplateDef) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *TemplateDef) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TemplateDef) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TemplateDef) SetDescription(v string) {
	o.Description = v
}

// GetHelp returns the Help field value
func (o *TemplateDef) GetHelp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Help
}

// GetHelpOk returns a tuple with the Help field value
// and a boolean to check if the value has been set.
func (o *TemplateDef) GetHelpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Help, true
}

// SetHelp sets field value
func (o *TemplateDef) SetHelp(v string) {
	o.Help = v
}

// GetCategory returns the Category field value
func (o *TemplateDef) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *TemplateDef) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *TemplateDef) SetCategory(v string) {
	o.Category = v
}

// GetExpr returns the Expr field value
func (o *TemplateDef) GetExpr() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Expr
}

// GetExprOk returns a tuple with the Expr field value
// and a boolean to check if the value has been set.
func (o *TemplateDef) GetExprOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expr, true
}

// SetExpr sets field value
func (o *TemplateDef) SetExpr(v []map[string]interface{}) {
	o.Expr = v
}

// GetArgs returns the Args field value
func (o *TemplateDef) GetArgs() []TemplateArgDef {
	if o == nil {
		var ret []TemplateArgDef
		return ret
	}

	return o.Args
}

// GetArgsOk returns a tuple with the Args field value
// and a boolean to check if the value has been set.
func (o *TemplateDef) GetArgsOk() ([]TemplateArgDef, bool) {
	if o == nil {
		return nil, false
	}
	return o.Args, true
}

// SetArgs sets field value
func (o *TemplateDef) SetArgs(v []TemplateArgDef) {
	o.Args = v
}

// GetExpose returns the Expose field value if set, zero value otherwise.
func (o *TemplateDef) GetExpose() bool {
	if o == nil || o.Expose == nil {
		var ret bool
		return ret
	}
	return *o.Expose
}

// GetExposeOk returns a tuple with the Expose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDef) GetExposeOk() (*bool, bool) {
	if o == nil || o.Expose == nil {
		return nil, false
	}
	return o.Expose, true
}

// HasExpose returns a boolean if a field has been set.
func (o *TemplateDef) HasExpose() bool {
	if o != nil && o.Expose != nil {
		return true
	}

	return false
}

// SetExpose gets a reference to the given bool and assigns it to the Expose field.
func (o *TemplateDef) SetExpose(v bool) {
	o.Expose = &v
}

// GetName returns the Name field value
func (o *TemplateDef) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateDef) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateDef) SetName(v string) {
	o.Name = v
}

func (o TemplateDef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["help"] = o.Help
	}
	if true {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["expr"] = o.Expr
	}
	if true {
		toSerialize["args"] = o.Args
	}
	if o.Expose != nil {
		toSerialize["expose"] = o.Expose
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateDef struct {
	value *TemplateDef
	isSet bool
}

func (v NullableTemplateDef) Get() *TemplateDef {
	return v.value
}

func (v *NullableTemplateDef) Set(val *TemplateDef) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateDef) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateDef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateDef(val *TemplateDef) *NullableTemplateDef {
	return &NullableTemplateDef{value: val, isSet: true}
}

func (v NullableTemplateDef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateDef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


