/*
Integration API reference docs

Use the Integration API to push data to and retrieve data from Talon.One in real time. For more background information about this API, see [Integration API overview](/docs/dev/integration-api/overview)  For example, use this API to share shopping cart information as a session with Talon.One and evaluate promotion rules. You can also create custom events to track specific actions that do not fit into the session data model.  Ensure you [authenticate](#section/Authentication) to make requests to the API.  <div class=\"redoc-section\">   <p class=\"title\">Are you looking for a different API?</p>    If you need the API to:    - Interact with the Campaign Manager for backoffice operations, see [the Management API reference docs](https://docs.talon.one/management-api).   - Integrate with Talon.One from a CEP or CDP platform, see [the Third-party API reference docs](https://docs.talon.one/third-party-api).  </div>  # Authentication  <SecurityDefinitions /> 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package talon

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)


// CustomerProfilesApiService CustomerProfilesApi service
type CustomerProfilesApiService service

type ApiDeleteCustomerDataRequest struct {
	ctx context.Context
	ApiService *CustomerProfilesApiService
	integrationId string
}

func (r ApiDeleteCustomerDataRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCustomerDataExecute(r)
}

/*
DeleteCustomerData Delete customer's personal data

Delete all attributes on the customer profile and on entities that reference this customer profile.

**Important:** To preserve performance, we recommend avoiding deleting customer data during peak-traffic hours.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param integrationId The integration ID of the customer profile. You can get the `integrationId` of a profile using: - A customer session integration Id with the [Update customer session endpoint](https://docs.talon.one/integration-api/#operation/updateCustomerSessionV2). - The Management API with the [List application's customers endpoint](https://docs.talon.one/management-api/#operation/getApplicationCustomers). 
 @return ApiDeleteCustomerDataRequest
*/
func (a *CustomerProfilesApiService) DeleteCustomerData(ctx context.Context, integrationId string) ApiDeleteCustomerDataRequest {
	return ApiDeleteCustomerDataRequest{
		ApiService: a,
		ctx: ctx,
		integrationId: integrationId,
	}
}

// Execute executes the request
func (a *CustomerProfilesApiService) DeleteCustomerDataExecute(r ApiDeleteCustomerDataRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerProfilesApiService.DeleteCustomerData")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/customer_data/{integrationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"integrationId"+"}", url.PathEscape(parameterToString(r.integrationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key_v1"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponseWithStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponseWithStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetCustomerInventoryRequest struct {
	ctx context.Context
	ApiService *CustomerProfilesApiService
	integrationId string
	profile *bool
	referrals *bool
	coupons *bool
	loyalty *bool
	giveaways *bool
	loyaltyProjectionEndDate *time.Time
}

// Set to &#x60;true&#x60; to include customer profile information in the response.
func (r ApiGetCustomerInventoryRequest) Profile(profile bool) ApiGetCustomerInventoryRequest {
	r.profile = &profile
	return r
}

// Set to &#x60;true&#x60; to include referral information in the response.
func (r ApiGetCustomerInventoryRequest) Referrals(referrals bool) ApiGetCustomerInventoryRequest {
	r.referrals = &referrals
	return r
}

// Set to &#x60;true&#x60; to include coupon information in the response.
func (r ApiGetCustomerInventoryRequest) Coupons(coupons bool) ApiGetCustomerInventoryRequest {
	r.coupons = &coupons
	return r
}

// Set to &#x60;true&#x60; to include loyalty information in the response.
func (r ApiGetCustomerInventoryRequest) Loyalty(loyalty bool) ApiGetCustomerInventoryRequest {
	r.loyalty = &loyalty
	return r
}

// Set to &#x60;true&#x60; to include giveaways information in the response.
func (r ApiGetCustomerInventoryRequest) Giveaways(giveaways bool) ApiGetCustomerInventoryRequest {
	r.giveaways = &giveaways
	return r
}

// Set an end date to query the projected loyalty balances. You can project results up to 31 days from today.
func (r ApiGetCustomerInventoryRequest) LoyaltyProjectionEndDate(loyaltyProjectionEndDate time.Time) ApiGetCustomerInventoryRequest {
	r.loyaltyProjectionEndDate = &loyaltyProjectionEndDate
	return r
}

func (r ApiGetCustomerInventoryRequest) Execute() (*CustomerInventory, *http.Response, error) {
	return r.ApiService.GetCustomerInventoryExecute(r)
}

/*
GetCustomerInventory List customer data

Return the customer inventory regarding entities referencing this customer profile's `integrationId`.

Typical entities returned are: customer profile information, referral codes, loyalty points and reserved coupons.
Reserved coupons also include redeemed coupons.

You can also use this endpoint to get the projected loyalty balances in order to notify
your customers about points that are about to expire, or to remind them how many points they have.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param integrationId The integration ID of the customer profile. You can get the `integrationId` of a profile using: - A customer session integration Id with the [Update customer session endpoint](https://docs.talon.one/integration-api/#operation/updateCustomerSessionV2). - The Management API with the [List application's customers endpoint](https://docs.talon.one/management-api/#operation/getApplicationCustomers). 
 @return ApiGetCustomerInventoryRequest
*/
func (a *CustomerProfilesApiService) GetCustomerInventory(ctx context.Context, integrationId string) ApiGetCustomerInventoryRequest {
	return ApiGetCustomerInventoryRequest{
		ApiService: a,
		ctx: ctx,
		integrationId: integrationId,
	}
}

// Execute executes the request
//  @return CustomerInventory
func (a *CustomerProfilesApiService) GetCustomerInventoryExecute(r ApiGetCustomerInventoryRequest) (*CustomerInventory, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomerInventory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerProfilesApiService.GetCustomerInventory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/customer_profiles/{integrationId}/inventory"
	localVarPath = strings.Replace(localVarPath, "{"+"integrationId"+"}", url.PathEscape(parameterToString(r.integrationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.profile != nil {
		localVarQueryParams.Add("profile", parameterToString(*r.profile, ""))
	}
	if r.referrals != nil {
		localVarQueryParams.Add("referrals", parameterToString(*r.referrals, ""))
	}
	if r.coupons != nil {
		localVarQueryParams.Add("coupons", parameterToString(*r.coupons, ""))
	}
	if r.loyalty != nil {
		localVarQueryParams.Add("loyalty", parameterToString(*r.loyalty, ""))
	}
	if r.giveaways != nil {
		localVarQueryParams.Add("giveaways", parameterToString(*r.giveaways, ""))
	}
	if r.loyaltyProjectionEndDate != nil {
		localVarQueryParams.Add("loyaltyProjectionEndDate", parameterToString(*r.loyaltyProjectionEndDate, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key_v1"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponseWithStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponseWithStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateCustomerProfileV2Request struct {
	ctx context.Context
	ApiService *CustomerProfilesApiService
	integrationId string
	body *CustomerProfileIntegrationRequestV2
	runRuleEngine *bool
	dry *bool
}

func (r ApiUpdateCustomerProfileV2Request) Body(body CustomerProfileIntegrationRequestV2) ApiUpdateCustomerProfileV2Request {
	r.body = &body
	return r
}

// Indicates whether to run the Rule Engine.  If &#x60;true&#x60;, the response includes: - The effects generated by the triggered campaigns are returned in the &#x60;effects&#x60; property. - The created coupons and referral objects.  If &#x60;false&#x60;: - The rules are not executed and the &#x60;effects&#x60; property is always empty. - The response time improves. - You cannot use &#x60;responseContent&#x60; in the body. 
func (r ApiUpdateCustomerProfileV2Request) RunRuleEngine(runRuleEngine bool) ApiUpdateCustomerProfileV2Request {
	r.runRuleEngine = &runRuleEngine
	return r
}

// Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;.  This property only works when &#x60;runRuleEngine&#x3D;true&#x60;. 
func (r ApiUpdateCustomerProfileV2Request) Dry(dry bool) ApiUpdateCustomerProfileV2Request {
	r.dry = &dry
	return r
}

func (r ApiUpdateCustomerProfileV2Request) Execute() (*IntegrationStateV2, *http.Response, error) {
	return r.ApiService.UpdateCustomerProfileV2Execute(r)
}

/*
UpdateCustomerProfileV2 Update customer profile

Update or create a [Customer Profile](/docs/dev/concepts/entities#customer-profile). This endpoint triggers the Rule Builder.

You can use this endpoint to:
- Set attributes on the given customer profile. Ensure you create the attributes in the Campaign Manager, first.
- Modify the audience the customer profile is a member of.

<div class="redoc-section">
  <p class="title">Performance tips</p>

  Updating a customer profile returns a response with the requested integration state.

  You can use the `responseContent` property to save yourself extra API calls. For example, you can get
  the customer profile details directly without extra requests.
</div>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param integrationId The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier. 
 @return ApiUpdateCustomerProfileV2Request
*/
func (a *CustomerProfilesApiService) UpdateCustomerProfileV2(ctx context.Context, integrationId string) ApiUpdateCustomerProfileV2Request {
	return ApiUpdateCustomerProfileV2Request{
		ApiService: a,
		ctx: ctx,
		integrationId: integrationId,
	}
}

// Execute executes the request
//  @return IntegrationStateV2
func (a *CustomerProfilesApiService) UpdateCustomerProfileV2Execute(r ApiUpdateCustomerProfileV2Request) (*IntegrationStateV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IntegrationStateV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerProfilesApiService.UpdateCustomerProfileV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/customer_profiles/{integrationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"integrationId"+"}", url.PathEscape(parameterToString(r.integrationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.runRuleEngine != nil {
		localVarQueryParams.Add("runRuleEngine", parameterToString(*r.runRuleEngine, ""))
	}
	if r.dry != nil {
		localVarQueryParams.Add("dry", parameterToString(*r.dry, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key_v1"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponseWithStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponseWithStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateCustomerProfilesV2Request struct {
	ctx context.Context
	ApiService *CustomerProfilesApiService
	body *MultipleCustomerProfileIntegrationRequest
	silent *string
}

func (r ApiUpdateCustomerProfilesV2Request) Body(body MultipleCustomerProfileIntegrationRequest) ApiUpdateCustomerProfilesV2Request {
	r.body = &body
	return r
}

// Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the perfomance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains essential data such as the updated customer profiles and session-related information. 
func (r ApiUpdateCustomerProfilesV2Request) Silent(silent string) ApiUpdateCustomerProfilesV2Request {
	r.silent = &silent
	return r
}

func (r ApiUpdateCustomerProfilesV2Request) Execute() (*MultipleCustomerProfileIntegrationResponseV2, *http.Response, error) {
	return r.ApiService.UpdateCustomerProfilesV2Execute(r)
}

/*
UpdateCustomerProfilesV2 Update multiple customer profiles

Update (or create) up to 1000 [customer profiles](/docs/dev/concepts/entities#customer-profile) in 1 request.

The `integrationId` must be any identifier that remains stable for
the customer. Do not use an ID that the customer can update
themselves. For example, you can use a database ID.

A customer profile [can be linked to one or more sessions](/integration-api/#tag/Customer-sessions).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateCustomerProfilesV2Request
*/
func (a *CustomerProfilesApiService) UpdateCustomerProfilesV2(ctx context.Context) ApiUpdateCustomerProfilesV2Request {
	return ApiUpdateCustomerProfilesV2Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MultipleCustomerProfileIntegrationResponseV2
func (a *CustomerProfilesApiService) UpdateCustomerProfilesV2Execute(r ApiUpdateCustomerProfilesV2Request) (*MultipleCustomerProfileIntegrationResponseV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MultipleCustomerProfileIntegrationResponseV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerProfilesApiService.UpdateCustomerProfilesV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/customer_profiles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.silent != nil {
		localVarQueryParams.Add("silent", parameterToString(*r.silent, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key_v1"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponseWithStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponseWithStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}