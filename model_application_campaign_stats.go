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

// ApplicationCampaignStats Provides statistics regarding an application's campaigns.
type ApplicationCampaignStats struct {
	// Number of draft campaigns.
	Draft int32 `json:"draft"`
	// Number of disabled campaigns.
	Disabled int32 `json:"disabled"`
	// Number of scheduled campaigns.
	Scheduled int32 `json:"scheduled"`
	// Number of running campaigns.
	Running int32 `json:"running"`
	// Number of expired campaigns.
	Expired int32 `json:"expired"`
	// Number of archived campaigns.
	Archived int32 `json:"archived"`
}

// NewApplicationCampaignStats instantiates a new ApplicationCampaignStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCampaignStats(draft int32, disabled int32, scheduled int32, running int32, expired int32, archived int32) *ApplicationCampaignStats {
	this := ApplicationCampaignStats{}
	this.Draft = draft
	this.Disabled = disabled
	this.Scheduled = scheduled
	this.Running = running
	this.Expired = expired
	this.Archived = archived
	return &this
}

// NewApplicationCampaignStatsWithDefaults instantiates a new ApplicationCampaignStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCampaignStatsWithDefaults() *ApplicationCampaignStats {
	this := ApplicationCampaignStats{}
	return &this
}

// GetDraft returns the Draft field value
func (o *ApplicationCampaignStats) GetDraft() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Draft
}

// GetDraftOk returns a tuple with the Draft field value
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignStats) GetDraftOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Draft, true
}

// SetDraft sets field value
func (o *ApplicationCampaignStats) SetDraft(v int32) {
	o.Draft = v
}

// GetDisabled returns the Disabled field value
func (o *ApplicationCampaignStats) GetDisabled() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignStats) GetDisabledOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *ApplicationCampaignStats) SetDisabled(v int32) {
	o.Disabled = v
}

// GetScheduled returns the Scheduled field value
func (o *ApplicationCampaignStats) GetScheduled() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Scheduled
}

// GetScheduledOk returns a tuple with the Scheduled field value
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignStats) GetScheduledOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scheduled, true
}

// SetScheduled sets field value
func (o *ApplicationCampaignStats) SetScheduled(v int32) {
	o.Scheduled = v
}

// GetRunning returns the Running field value
func (o *ApplicationCampaignStats) GetRunning() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Running
}

// GetRunningOk returns a tuple with the Running field value
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignStats) GetRunningOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Running, true
}

// SetRunning sets field value
func (o *ApplicationCampaignStats) SetRunning(v int32) {
	o.Running = v
}

// GetExpired returns the Expired field value
func (o *ApplicationCampaignStats) GetExpired() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignStats) GetExpiredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expired, true
}

// SetExpired sets field value
func (o *ApplicationCampaignStats) SetExpired(v int32) {
	o.Expired = v
}

// GetArchived returns the Archived field value
func (o *ApplicationCampaignStats) GetArchived() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignStats) GetArchivedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *ApplicationCampaignStats) SetArchived(v int32) {
	o.Archived = v
}

func (o ApplicationCampaignStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["draft"] = o.Draft
	}
	if true {
		toSerialize["disabled"] = o.Disabled
	}
	if true {
		toSerialize["scheduled"] = o.Scheduled
	}
	if true {
		toSerialize["running"] = o.Running
	}
	if true {
		toSerialize["expired"] = o.Expired
	}
	if true {
		toSerialize["archived"] = o.Archived
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationCampaignStats struct {
	value *ApplicationCampaignStats
	isSet bool
}

func (v NullableApplicationCampaignStats) Get() *ApplicationCampaignStats {
	return v.value
}

func (v *NullableApplicationCampaignStats) Set(val *ApplicationCampaignStats) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCampaignStats) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCampaignStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCampaignStats(val *ApplicationCampaignStats) *NullableApplicationCampaignStats {
	return &NullableApplicationCampaignStats{value: val, isSet: true}
}

func (v NullableApplicationCampaignStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCampaignStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


