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

// ShowNotificationEffectProps The properties specific to the \"showNotification\" effect. This gets triggered whenever a validated rule contained a \"show notification\" effect.
type ShowNotificationEffectProps struct {
	// The type of notification that should be shown (e.g. error/warning/info).
	NotificationType string `json:"notificationType"`
	// Title of the notification.
	Title string `json:"title"`
	// Body of the notification.
	Body string `json:"body"`
}

// NewShowNotificationEffectProps instantiates a new ShowNotificationEffectProps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShowNotificationEffectProps(notificationType string, title string, body string) *ShowNotificationEffectProps {
	this := ShowNotificationEffectProps{}
	this.NotificationType = notificationType
	this.Title = title
	this.Body = body
	return &this
}

// NewShowNotificationEffectPropsWithDefaults instantiates a new ShowNotificationEffectProps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShowNotificationEffectPropsWithDefaults() *ShowNotificationEffectProps {
	this := ShowNotificationEffectProps{}
	return &this
}

// GetNotificationType returns the NotificationType field value
func (o *ShowNotificationEffectProps) GetNotificationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value
// and a boolean to check if the value has been set.
func (o *ShowNotificationEffectProps) GetNotificationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationType, true
}

// SetNotificationType sets field value
func (o *ShowNotificationEffectProps) SetNotificationType(v string) {
	o.NotificationType = v
}

// GetTitle returns the Title field value
func (o *ShowNotificationEffectProps) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ShowNotificationEffectProps) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ShowNotificationEffectProps) SetTitle(v string) {
	o.Title = v
}

// GetBody returns the Body field value
func (o *ShowNotificationEffectProps) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *ShowNotificationEffectProps) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *ShowNotificationEffectProps) SetBody(v string) {
	o.Body = v
}

func (o ShowNotificationEffectProps) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["notificationType"] = o.NotificationType
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableShowNotificationEffectProps struct {
	value *ShowNotificationEffectProps
	isSet bool
}

func (v NullableShowNotificationEffectProps) Get() *ShowNotificationEffectProps {
	return v.value
}

func (v *NullableShowNotificationEffectProps) Set(val *ShowNotificationEffectProps) {
	v.value = val
	v.isSet = true
}

func (v NullableShowNotificationEffectProps) IsSet() bool {
	return v.isSet
}

func (v *NullableShowNotificationEffectProps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShowNotificationEffectProps(val *ShowNotificationEffectProps) *NullableShowNotificationEffectProps {
	return &NullableShowNotificationEffectProps{value: val, isSet: true}
}

func (v NullableShowNotificationEffectProps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShowNotificationEffectProps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

