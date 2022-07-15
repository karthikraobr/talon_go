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

// UserFeedNotifications Notifications to notify CAMA user about.
type UserFeedNotifications struct {
	// Timestamp of the last request for this list.
	LastUpdate time.Time `json:"lastUpdate"`
	// List of all notifications to notify the user about.
	Notifications []FeedNotification `json:"notifications"`
}

// NewUserFeedNotifications instantiates a new UserFeedNotifications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFeedNotifications(lastUpdate time.Time, notifications []FeedNotification) *UserFeedNotifications {
	this := UserFeedNotifications{}
	this.LastUpdate = lastUpdate
	this.Notifications = notifications
	return &this
}

// NewUserFeedNotificationsWithDefaults instantiates a new UserFeedNotifications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFeedNotificationsWithDefaults() *UserFeedNotifications {
	this := UserFeedNotifications{}
	return &this
}

// GetLastUpdate returns the LastUpdate field value
func (o *UserFeedNotifications) GetLastUpdate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value
// and a boolean to check if the value has been set.
func (o *UserFeedNotifications) GetLastUpdateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdate, true
}

// SetLastUpdate sets field value
func (o *UserFeedNotifications) SetLastUpdate(v time.Time) {
	o.LastUpdate = v
}

// GetNotifications returns the Notifications field value
func (o *UserFeedNotifications) GetNotifications() []FeedNotification {
	if o == nil {
		var ret []FeedNotification
		return ret
	}

	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value
// and a boolean to check if the value has been set.
func (o *UserFeedNotifications) GetNotificationsOk() ([]FeedNotification, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notifications, true
}

// SetNotifications sets field value
func (o *UserFeedNotifications) SetNotifications(v []FeedNotification) {
	o.Notifications = v
}

func (o UserFeedNotifications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["lastUpdate"] = o.LastUpdate
	}
	if true {
		toSerialize["notifications"] = o.Notifications
	}
	return json.Marshal(toSerialize)
}

type NullableUserFeedNotifications struct {
	value *UserFeedNotifications
	isSet bool
}

func (v NullableUserFeedNotifications) Get() *UserFeedNotifications {
	return v.value
}

func (v *NullableUserFeedNotifications) Set(val *UserFeedNotifications) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFeedNotifications) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFeedNotifications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFeedNotifications(val *UserFeedNotifications) *NullableUserFeedNotifications {
	return &NullableUserFeedNotifications{value: val, isSet: true}
}

func (v NullableUserFeedNotifications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFeedNotifications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


