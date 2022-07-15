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

// WebhookLogEntry Log of webhook api calls.
type WebhookLogEntry struct {
	// UUID reference of the webhook request.
	Id string `json:"id"`
	// UUID reference of the integration request linked to this webhook request.
	IntegrationRequestUuid string `json:"integrationRequestUuid"`
	// ID of the webhook that triggered the request.
	WebhookId int32 `json:"webhookId"`
	// ID of the application that triggered the webhook.
	ApplicationId *int32 `json:"applicationId,omitempty"`
	// Target url of request
	Url string `json:"url"`
	// Request message
	Request string `json:"request"`
	// Response message
	Response *string `json:"response,omitempty"`
	// HTTP status code of response.
	Status *int32 `json:"status,omitempty"`
	// Timestamp of request
	RequestTime time.Time `json:"requestTime"`
	// Timestamp of response
	ResponseTime *time.Time `json:"responseTime,omitempty"`
}

// NewWebhookLogEntry instantiates a new WebhookLogEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookLogEntry(id string, integrationRequestUuid string, webhookId int32, url string, request string, requestTime time.Time) *WebhookLogEntry {
	this := WebhookLogEntry{}
	this.Id = id
	this.IntegrationRequestUuid = integrationRequestUuid
	this.WebhookId = webhookId
	this.Url = url
	this.Request = request
	this.RequestTime = requestTime
	return &this
}

// NewWebhookLogEntryWithDefaults instantiates a new WebhookLogEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookLogEntryWithDefaults() *WebhookLogEntry {
	this := WebhookLogEntry{}
	return &this
}

// GetId returns the Id field value
func (o *WebhookLogEntry) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhookLogEntry) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhookLogEntry) SetId(v string) {
	o.Id = v
}

// GetIntegrationRequestUuid returns the IntegrationRequestUuid field value
func (o *WebhookLogEntry) GetIntegrationRequestUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationRequestUuid
}

// GetIntegrationRequestUuidOk returns a tuple with the IntegrationRequestUuid field value
// and a boolean to check if the value has been set.
func (o *WebhookLogEntry) GetIntegrationRequestUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationRequestUuid, true
}

// SetIntegrationRequestUuid sets field value
func (o *WebhookLogEntry) SetIntegrationRequestUuid(v string) {
	o.IntegrationRequestUuid = v
}

// GetWebhookId returns the WebhookId field value
func (o *WebhookLogEntry) GetWebhookId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value
// and a boolean to check if the value has been set.
func (o *WebhookLogEntry) GetWebhookIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebhookId, true
}

// SetWebhookId sets field value
func (o *WebhookLogEntry) SetWebhookId(v int32) {
	o.WebhookId = v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *WebhookLogEntry) GetApplicationId() int32 {
	if o == nil || o.ApplicationId == nil {
		var ret int32
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookLogEntry) GetApplicationIdOk() (*int32, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *WebhookLogEntry) HasApplicationId() bool {
	if o != nil && o.ApplicationId != nil {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.
func (o *WebhookLogEntry) SetApplicationId(v int32) {
	o.ApplicationId = &v
}

// GetUrl returns the Url field value
func (o *WebhookLogEntry) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WebhookLogEntry) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WebhookLogEntry) SetUrl(v string) {
	o.Url = v
}

// GetRequest returns the Request field value
func (o *WebhookLogEntry) GetRequest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *WebhookLogEntry) GetRequestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Request, true
}

// SetRequest sets field value
func (o *WebhookLogEntry) SetRequest(v string) {
	o.Request = v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *WebhookLogEntry) GetResponse() string {
	if o == nil || o.Response == nil {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookLogEntry) GetResponseOk() (*string, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *WebhookLogEntry) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *WebhookLogEntry) SetResponse(v string) {
	o.Response = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WebhookLogEntry) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookLogEntry) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WebhookLogEntry) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *WebhookLogEntry) SetStatus(v int32) {
	o.Status = &v
}

// GetRequestTime returns the RequestTime field value
func (o *WebhookLogEntry) GetRequestTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.RequestTime
}

// GetRequestTimeOk returns a tuple with the RequestTime field value
// and a boolean to check if the value has been set.
func (o *WebhookLogEntry) GetRequestTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestTime, true
}

// SetRequestTime sets field value
func (o *WebhookLogEntry) SetRequestTime(v time.Time) {
	o.RequestTime = v
}

// GetResponseTime returns the ResponseTime field value if set, zero value otherwise.
func (o *WebhookLogEntry) GetResponseTime() time.Time {
	if o == nil || o.ResponseTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ResponseTime
}

// GetResponseTimeOk returns a tuple with the ResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookLogEntry) GetResponseTimeOk() (*time.Time, bool) {
	if o == nil || o.ResponseTime == nil {
		return nil, false
	}
	return o.ResponseTime, true
}

// HasResponseTime returns a boolean if a field has been set.
func (o *WebhookLogEntry) HasResponseTime() bool {
	if o != nil && o.ResponseTime != nil {
		return true
	}

	return false
}

// SetResponseTime gets a reference to the given time.Time and assigns it to the ResponseTime field.
func (o *WebhookLogEntry) SetResponseTime(v time.Time) {
	o.ResponseTime = &v
}

func (o WebhookLogEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["integrationRequestUuid"] = o.IntegrationRequestUuid
	}
	if true {
		toSerialize["webhookId"] = o.WebhookId
	}
	if o.ApplicationId != nil {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["request"] = o.Request
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["requestTime"] = o.RequestTime
	}
	if o.ResponseTime != nil {
		toSerialize["responseTime"] = o.ResponseTime
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookLogEntry struct {
	value *WebhookLogEntry
	isSet bool
}

func (v NullableWebhookLogEntry) Get() *WebhookLogEntry {
	return v.value
}

func (v *NullableWebhookLogEntry) Set(val *WebhookLogEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookLogEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookLogEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookLogEntry(val *WebhookLogEntry) *NullableWebhookLogEntry {
	return &NullableWebhookLogEntry{value: val, isSet: true}
}

func (v NullableWebhookLogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookLogEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

