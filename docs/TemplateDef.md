# TemplateDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**ApplicationId** | **int32** | The ID of the application that owns this entity. | 
**Title** | **string** | Campaigner-friendly name for the template that will be shown in the rule editor. | 
**Description** | **string** | A short description of the template that will be shown in the rule editor. | 
**Help** | **string** | Extended help text for the template. | 
**Category** | **string** | Used for grouping templates in the rule editor sidebar. | 
**Expr** | **[]map[string]interface{}** | A Talang expression that contains variable bindings referring to args. | 
**Args** | [**[]TemplateArgDef**](TemplateArgDef.md) | An array of argument definitions. | 
**Expose** | Pointer to **bool** | A flag to control exposure in Rule Builder. | [optional] [default to false]
**Name** | **string** | The template name used in Talang. | 

## Methods

### NewTemplateDef

`func NewTemplateDef(id int32, created time.Time, applicationId int32, title string, description string, help string, category string, expr []map[string]interface{}, args []TemplateArgDef, name string, ) *TemplateDef`

NewTemplateDef instantiates a new TemplateDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateDefWithDefaults

`func NewTemplateDefWithDefaults() *TemplateDef`

NewTemplateDefWithDefaults instantiates a new TemplateDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateDef) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateDef) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateDef) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *TemplateDef) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TemplateDef) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TemplateDef) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetApplicationId

`func (o *TemplateDef) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *TemplateDef) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *TemplateDef) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.


### GetTitle

`func (o *TemplateDef) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TemplateDef) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TemplateDef) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *TemplateDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateDef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateDef) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHelp

`func (o *TemplateDef) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *TemplateDef) GetHelpOk() (*string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *TemplateDef) SetHelp(v string)`

SetHelp sets Help field to given value.


### GetCategory

`func (o *TemplateDef) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TemplateDef) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TemplateDef) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetExpr

`func (o *TemplateDef) GetExpr() []map[string]interface{}`

GetExpr returns the Expr field if non-nil, zero value otherwise.

### GetExprOk

`func (o *TemplateDef) GetExprOk() (*[]map[string]interface{}, bool)`

GetExprOk returns a tuple with the Expr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpr

`func (o *TemplateDef) SetExpr(v []map[string]interface{})`

SetExpr sets Expr field to given value.


### GetArgs

`func (o *TemplateDef) GetArgs() []TemplateArgDef`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *TemplateDef) GetArgsOk() (*[]TemplateArgDef, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *TemplateDef) SetArgs(v []TemplateArgDef)`

SetArgs sets Args field to given value.


### GetExpose

`func (o *TemplateDef) GetExpose() bool`

GetExpose returns the Expose field if non-nil, zero value otherwise.

### GetExposeOk

`func (o *TemplateDef) GetExposeOk() (*bool, bool)`

GetExposeOk returns a tuple with the Expose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpose

`func (o *TemplateDef) SetExpose(v bool)`

SetExpose sets Expose field to given value.

### HasExpose

`func (o *TemplateDef) HasExpose() bool`

HasExpose returns a boolean if a field has been set.

### GetName

`func (o *TemplateDef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateDef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateDef) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


