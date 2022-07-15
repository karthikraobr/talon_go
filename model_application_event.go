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

// ApplicationEvent 
type ApplicationEvent struct {
	// Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`
	// The globally unique Talon.One ID of the customer that created this entity.
	ProfileId *int32 `json:"profileId,omitempty"`
	// The globally unique Talon.One ID of the session that contains this event.
	SessionId *int32 `json:"sessionId,omitempty"`
	// A string representing the event. Must not be a reserved event name.
	Type string `json:"type"`
	// Additional JSON serialized data associated with the event.
	Attributes map[string]interface{} `json:"attributes"`
	// An array containing the effects that were applied as a result of this event.
	Effects []map[string]interface{} `json:"effects"`
	// An array containing the rule failure reasons which happened during this event.
	RuleFailureReasons []RuleFailureReason `json:"ruleFailureReasons,omitempty"`
}

// NewApplicationEvent instantiates a new ApplicationEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationEvent(id int32, created time.Time, applicationId int32, type_ string, attributes map[string]interface{}, effects []map[string]interface{}) *ApplicationEvent {
	this := ApplicationEvent{}
	this.Id = id
	this.Created = created
	this.ApplicationId = applicationId
	this.Type = type_
	this.Attributes = attributes
	this.Effects = effects
	return &this
}

// NewApplicationEventWithDefaults instantiates a new ApplicationEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationEventWithDefaults() *ApplicationEvent {
	this := ApplicationEvent{}
	return &this
}

// GetId returns the Id field value
func (o *ApplicationEvent) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationEvent) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationEvent) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *ApplicationEvent) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *ApplicationEvent) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *ApplicationEvent) SetCreated(v time.Time) {
	o.Created = v
}

// GetApplicationId returns the ApplicationId field value
func (o *ApplicationEvent) GetApplicationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *ApplicationEvent) GetApplicationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *ApplicationEvent) SetApplicationId(v int32) {
	o.ApplicationId = v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *ApplicationEvent) GetProfileId() int32 {
	if o == nil || o.ProfileId == nil {
		var ret int32
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEvent) GetProfileIdOk() (*int32, bool) {
	if o == nil || o.ProfileId == nil {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *ApplicationEvent) HasProfileId() bool {
	if o != nil && o.ProfileId != nil {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given int32 and assigns it to the ProfileId field.
func (o *ApplicationEvent) SetProfileId(v int32) {
	o.ProfileId = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *ApplicationEvent) GetSessionId() int32 {
	if o == nil || o.SessionId == nil {
		var ret int32
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEvent) GetSessionIdOk() (*int32, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *ApplicationEvent) HasSessionId() bool {
	if o != nil && o.SessionId != nil {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given int32 and assigns it to the SessionId field.
func (o *ApplicationEvent) SetSessionId(v int32) {
	o.SessionId = &v
}

// GetType returns the Type field value
func (o *ApplicationEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationEvent) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ApplicationEvent) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ApplicationEvent) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *ApplicationEvent) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetEffects returns the Effects field value
func (o *ApplicationEvent) GetEffects() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Effects
}

// GetEffectsOk returns a tuple with the Effects field value
// and a boolean to check if the value has been set.
func (o *ApplicationEvent) GetEffectsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Effects, true
}

// SetEffects sets field value
func (o *ApplicationEvent) SetEffects(v []map[string]interface{}) {
	o.Effects = v
}

// GetRuleFailureReasons returns the RuleFailureReasons field value if set, zero value otherwise.
func (o *ApplicationEvent) GetRuleFailureReasons() []RuleFailureReason {
	if o == nil || o.RuleFailureReasons == nil {
		var ret []RuleFailureReason
		return ret
	}
	return o.RuleFailureReasons
}

// GetRuleFailureReasonsOk returns a tuple with the RuleFailureReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEvent) GetRuleFailureReasonsOk() ([]RuleFailureReason, bool) {
	if o == nil || o.RuleFailureReasons == nil {
		return nil, false
	}
	return o.RuleFailureReasons, true
}

// HasRuleFailureReasons returns a boolean if a field has been set.
func (o *ApplicationEvent) HasRuleFailureReasons() bool {
	if o != nil && o.RuleFailureReasons != nil {
		return true
	}

	return false
}

// SetRuleFailureReasons gets a reference to the given []RuleFailureReason and assigns it to the RuleFailureReasons field.
func (o *ApplicationEvent) SetRuleFailureReasons(v []RuleFailureReason) {
	o.RuleFailureReasons = v
}

func (o ApplicationEvent) MarshalJSON() ([]byte, error) {
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
	if o.ProfileId != nil {
		toSerialize["profileId"] = o.ProfileId
	}
	if o.SessionId != nil {
		toSerialize["sessionId"] = o.SessionId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["effects"] = o.Effects
	}
	if o.RuleFailureReasons != nil {
		toSerialize["ruleFailureReasons"] = o.RuleFailureReasons
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationEvent struct {
	value *ApplicationEvent
	isSet bool
}

func (v NullableApplicationEvent) Get() *ApplicationEvent {
	return v.value
}

func (v *NullableApplicationEvent) Set(val *ApplicationEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationEvent(val *ApplicationEvent) *NullableApplicationEvent {
	return &NullableApplicationEvent{value: val, isSet: true}
}

func (v NullableApplicationEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

