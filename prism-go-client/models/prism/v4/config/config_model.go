/*
 * Generated file models/prism/v4/config/config_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix Prism APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Configure Tasks and Monitoring
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import5 "github.com/orichter/package-publishing-example-go/prism-go-client/v4/models/clustermgmt/v4/config"
	import2 "github.com/orichter/package-publishing-example-go/prism-go-client/v4/models/common/v1/config"
	import1 "github.com/orichter/package-publishing-example-go/prism-go-client/v4/models/common/v1/response"
	import4 "github.com/orichter/package-publishing-example-go/prism-go-client/v4/models/prism/v4/error"
	import3 "github.com/orichter/package-publishing-example-go/prism-go-client/v4/models/vmm/v4/ahv/config"
	"time"
)

type AssociationDetail struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier for the given category that is used across all v4 apis/entities/resources where categories are referenced.<br>
	The field is in UUID format.<br>
	A type 4 UUID is generated during category creation.
	*/
	CategoryId *string `json:"categoryId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	ResourceGroup *ResourceGroup `json:"resourceGroup,omitempty"`
	/*
	  The UUID of the entity or policy associated with the particular category.
	*/
	ResourceId *string `json:"resourceId,omitempty"`

	ResourceType *ResourceType `json:"resourceType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewAssociationDetail() *AssociationDetail {
	p := new(AssociationDetail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.AssociationDetail"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
This attribute contains a list of entities which have been assigned the given category.<br>
These entities are grouped by entity types (like VM or HOST) or policy types (like PROTECTION_POLICY or
NGT_POLICY).<br>
Each associated object contains the total entities belonging to the given entity type, category extId, and
references (for example, for VM it will be VM UUID).
*/
type AssociationDetailOld struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier for the given category.
	*/
	CategoryId *string `json:"categoryId,omitempty"`
	/*
	  Denotes the type of a category.<br>
	There are three types of categories: SYSTEM, INTERNAL, and USER.<br>
	This field is immutable.
	*/
	Count *int `json:"count,omitempty"`

	ResourceGroup *ResourceGroup `json:"resourceGroup,omitempty"`

	ResourceReferences []Reference `json:"resourceReferences,omitempty"`

	ResourceType *ResourceType `json:"resourceType,omitempty"`
}

func NewAssociationDetailOld() *AssociationDetailOld {
	p := new(AssociationDetailOld)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.AssociationDetailOld"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type AssociationDetailOldProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier for the given category.
	*/
	CategoryId *string `json:"categoryId,omitempty"`
	/*
	  Denotes the type of a category.<br>
	There are three types of categories: SYSTEM, INTERNAL, and USER.<br>
	This field is immutable.
	*/
	Count *int `json:"count,omitempty"`

	ResourceGroup *ResourceGroup `json:"resourceGroup,omitempty"`

	ResourceReferences []Reference `json:"resourceReferences,omitempty"`

	ResourceType *ResourceType `json:"resourceType,omitempty"`
}

func NewAssociationDetailOldProjection() *AssociationDetailOldProjection {
	p := new(AssociationDetailOldProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.AssociationDetailOldProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type AssociationDetailProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier for the given category that is used across all v4 apis/entities/resources where categories are referenced.<br>
	The field is in UUID format.<br>
	A type 4 UUID is generated during category creation.
	*/
	CategoryId *string `json:"categoryId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	ResourceGroup *ResourceGroup `json:"resourceGroup,omitempty"`
	/*
	  The UUID of the entity or policy associated with the particular category.
	*/
	ResourceId *string `json:"resourceId,omitempty"`

	ResourceType *ResourceType `json:"resourceType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewAssociationDetailProjection() *AssociationDetailProjection {
	p := new(AssociationDetailProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.AssociationDetailProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
This attribute contains a list of entities and policies that have been assigned the given category.<br>
These entities are grouped by entity types (like VM or HOST) or policy types (like PROTECTION_POLICY or NGT_POLICY).<br>
Each associated object contains the total entities belonging to the given entity type, count, category extId, and
references (for example, for VM it will be VM UUID).
*/
type AssociationSummary struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier for the given category that is used across all v4 apis/entities/resources where categories are referenced.<br>
	The field is in UUID format.<br>
	A type 4 UUID is generated during category creation.
	*/
	CategoryId *string `json:"categoryId,omitempty"`
	/*
	  Count of associations of a particular type of entity or policy
	*/
	Count *int `json:"count,omitempty"`

	ResourceGroup *ResourceGroup `json:"resourceGroup,omitempty"`

	ResourceType *ResourceType `json:"resourceType,omitempty"`
}

func NewAssociationSummary() *AssociationSummary {
	p := new(AssociationSummary)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.AssociationSummary"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type AssociationSummaryProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier for the given category that is used across all v4 apis/entities/resources where categories are referenced.<br>
	The field is in UUID format.<br>
	A type 4 UUID is generated during category creation.
	*/
	CategoryId *string `json:"categoryId,omitempty"`
	/*
	  Count of associations of a particular type of entity or policy
	*/
	Count *int `json:"count,omitempty"`

	ResourceGroup *ResourceGroup `json:"resourceGroup,omitempty"`

	ResourceType *ResourceType `json:"resourceType,omitempty"`
}

func NewAssociationSummaryProjection() *AssociationSummaryProjection {
	p := new(AssociationSummaryProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.AssociationSummaryProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
This model would abstract away the common attributes as part of internal and external networks.
*/
type BaseNetwork struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DefaultGateway *import2.IPAddressOrFQDN `json:"defaultGateway"`
	/*
	  Range of IPs used for Prism Central network setup.
	*/
	IpRanges []import2.IpRange `json:"ipRanges"`

	SubnetMask *import2.IPAddressOrFQDN `json:"subnetMask"`
}

func (p *BaseNetwork) MarshalJSON() ([]byte, error) {
	type BaseNetworkProxy BaseNetwork
	return json.Marshal(struct {
		*BaseNetworkProxy
		DefaultGateway *import2.IPAddressOrFQDN `json:"defaultGateway,omitempty"`
		IpRanges       []import2.IpRange        `json:"ipRanges,omitempty"`
		SubnetMask     *import2.IPAddressOrFQDN `json:"subnetMask,omitempty"`
	}{
		BaseNetworkProxy: (*BaseNetworkProxy)(p),
		DefaultGateway:   p.DefaultGateway,
		IpRanges:         p.IpRanges,
		SubnetMask:       p.SubnetMask,
	})
}

func NewBaseNetwork() *BaseNetwork {
	p := new(BaseNetwork)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.BaseNetwork"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Bootstrap configuration details for the domain manager (Prism Central).
*/
type BootstrapConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of cloud-init commands required to bootstrap the domain manager (Prism Central) cluster on a startup.
	*/
	CloudInitConfig []import3.CloudInit `json:"cloudInitConfig,omitempty"`

	EnvironmentInfo *EnvironmentInfo `json:"environmentInfo,omitempty"`
}

func NewBootstrapConfig() *BootstrapConfig {
	p := new(BootstrapConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.BootstrapConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0/config/tasks/{taskExtId}/$actions/cancel Post operation
*/
type CancelTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCancelTaskApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCancelTaskApiResponse() *CancelTaskApiResponse {
	p := new(CancelTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.CancelTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CancelTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CancelTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCancelTaskApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

type Category struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  This field gives basic information about resources that are associated with the category.<br>
	The results present under this field summarize the counts of various kinds of resources associated with the category.<br>
	For more detailed information about the UUIDs of the resources, please look into the field `detailedAssociations`.<br>
	This field will be ignored, if given in the payload of `updateCategoryById` or `createCategory` APIs.<br>
	This field will not be present by default in `listCategories` API, unless the parameter $expand=associations is present in the URL.
	*/
	Associations []AssociationSummary `json:"associations,omitempty"`
	/*
	  A string consisting of the description of the category as defined by the user.<br>
	Description can be optionally provided in the payload of `createCategory` and `updateCategoryById` APIs.<br>
	Description field can be updated through `updateCategoryById` API.<br>
	The server does not validate this value nor does it enforce the uniqueness or any other constraints.<br>
	It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
	this field.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  This field gives detailed information about the resources which are associated with the category.<br>
	The results present under this field contain the UUIDs of the entities and policies of various kinds associated with the category.<br>
	This field will be ignored, if given in the payload of `updateCategoryById` or `createCategory` APIs.<br>
	This field will not be present by default in `listCategories` or `getCategoryById` APIs, unless the parameter $expand=detailedAssociations is present in the URL.
	*/
	DetailedAssociations []AssociationDetail `json:"detailedAssociations,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The key of a category when it is represented in `key:value` format.

	Constraints applicable when field is given in the payload during create and update:
	* A string of maxlength of 64
	* Character at the start cannot be `$`
	* Character `/` is not allowed anywhere

	It is a mandatory field in the payload of `createCategory` and `updateCategoryById` APIs.<br>
	This field can't be updated through `updateCategoryById` API.
	*/
	Key *string `json:"key"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  This field contains the UUID of a user who owns the category.<br>
	This field will be ignored if given in the payload of `createCategory` API. Hence, when a category is created, the logged-in user automatically becomes the owner of the category.<br>
	This field can be updated through `updateCategoryById` API, in which case, should be provided, UUID of a valid user is present in the system.<br>
	Validity of the user UUID can be checked by invoking the API: authn/users/{extId} in the 'Identity and Access Management' or 'IAM' namespace.<br>
	It is used for enabling RBAC access to self-owned categories.
	*/
	OwnerUuid *string `json:"ownerUuid,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *CategoryType `json:"type,omitempty"`
	/*
	  The value of a category when it is represented in `key:value` format.

	Constraints applicable when the field is given in the payload during create and update:
	* A string of max length 64
	* Character at the start cannot be `$`
	* Character `/` is not allowed anywhere

	It is a mandatory input field in the payload of `createCategory` and `updateCategoryById` APIs.<br>
	This field can be updated through `updateCategoryById` API.<br>
	Updating the value will not change the extId of the category.
	*/
	Value *string `json:"value"`
}

func (p *Category) MarshalJSON() ([]byte, error) {
	type CategoryProxy Category
	return json.Marshal(struct {
		*CategoryProxy
		Key   *string `json:"key,omitempty"`
		Value *string `json:"value,omitempty"`
	}{
		CategoryProxy: (*CategoryProxy)(p),
		Key:           p.Key,
		Value:         p.Value,
	})
}

func NewCategory() *Category {
	p := new(Category)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.Category"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
This attribute contains a list of entities that have been assigned the given category.<br>
These entities are grouped by entity types (like VM or HOST).<br>
Each associated object contains the total entities belonging to the given entity type, and category extId.
*/
type CategoryAssociationSummaryOld struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier for the given category.
	*/
	CategoryId *string `json:"categoryId,omitempty"`
	/*
	  Denotes the type of a category.<br>
	There are three types of categories: SYSTEM, INTERNAL, and USER.<br>
	This field is immutable.
	*/
	Count *int `json:"count,omitempty"`

	ResourceGroup *ResourceGroup `json:"resourceGroup,omitempty"`

	ResourceType *ResourceType `json:"resourceType,omitempty"`
}

func NewCategoryAssociationSummaryOld() *CategoryAssociationSummaryOld {
	p := new(CategoryAssociationSummaryOld)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.CategoryAssociationSummaryOld"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type CategoryAssociationSummaryOldProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier for the given category.
	*/
	CategoryId *string `json:"categoryId,omitempty"`
	/*
	  Denotes the type of a category.<br>
	There are three types of categories: SYSTEM, INTERNAL, and USER.<br>
	This field is immutable.
	*/
	Count *int `json:"count,omitempty"`

	ResourceGroup *ResourceGroup `json:"resourceGroup,omitempty"`

	ResourceType *ResourceType `json:"resourceType,omitempty"`
}

func NewCategoryAssociationSummaryOldProjection() *CategoryAssociationSummaryOldProjection {
	p := new(CategoryAssociationSummaryOldProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.CategoryAssociationSummaryOldProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Denotes the type of a category.<br>
There are three types of categories: SYSTEM, INTERNAL, and USER.<br>
This field is immutable.
*/
type CategoryOld struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Associations []CategoryAssociationSummaryOld `json:"associations,omitempty"`
	/*
	  This attribute contains the list of all the categories for which this category is the parent.<br>
	The parentExtId attributes of each child category is set as the extId of this category.<br>
	Note that this list only contains the Summary view of each child category.
	*/
	ChildCategories []CategorySummaryOld `json:"childCategories,omitempty"`
	/*
	  A string consisting of the description of the category as defined by the user.

	The server does not validate this value nor does it enforce the uniqueness or any other constraints.<br>
	It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
	this field.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The fully qualified name of this category. It is unique for each category.<br>
	It is a read-only field.
	The service constructs it from the name-parentExtId combination.
	An example of a fqName would be `Location/Bangalore`, where `Location` is the
	parent category's name and `Bangalore` is the category name.<br>
	This field is immutable.<br>
	*/
	FqName *string `json:"fqName,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Opaque metadata which can be associated to a category.<br>
	It is a list of key-value pairs.<br>
	For example, for a category 'California/SanJose' we can associate a geographical coordinate based metadata
	like: {'latitude': '37.3382° N', 'longitude': '121.8863° W'}.

	The server does not validate this value nor does it enforce the uniqueness or any other constraints.
	It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
	this field.
	*/
	Metadata []import2.KVPair `json:"metadata,omitempty"`
	/*
	  The short name of this category. It may not be unique for each category.<br>
	It is a mandatory field that must be specified inside the post/put request
	body.<br>
	This field is immutable.
	*/
	Name *string `json:"name"`
	/*
	  It is a read-only field inserted into category entity at the time of category creation, and which contains the UUID of
	the user who created this category. It is used for enabling authorization of a particular kind where the user has no
	access to view/create/update/delete any categories other than the category created by oneself.
	*/
	OwnerUuid *string `json:"ownerUuid,omitempty"`
	/*
	  The parent category of this category (may be null if this category is not part of a hierarchy).<br>
	Each category can have at most one parent.<br>
	A parent cannot be deleted until all the children categories are deleted first.<br>
	Must be specified inside the post/put request body for child categories (if not specified, the service assumes
	the category to be a parent category).<br>
	This field is immutable.
	*/
	ParentExtId *string `json:"parentExtId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *CategoryType `json:"type,omitempty"`
	/*
	  The user-specified name is a string that the user can specify; with syntax and semantics controlled by the user.

	The server does not validate this value nor does it enforce the uniqueness or any other constraints.<br>
	It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
	this field. Unlike the name of the categories, which are immutable, the user name can be changed by the user to meet their needs.
	*/
	UserSpecifiedName *string `json:"userSpecifiedName,omitempty"`
}

func (p *CategoryOld) MarshalJSON() ([]byte, error) {
	type CategoryOldProxy CategoryOld
	return json.Marshal(struct {
		*CategoryOldProxy
		Name *string `json:"name,omitempty"`
	}{
		CategoryOldProxy: (*CategoryOldProxy)(p),
		Name:             p.Name,
	})
}

func NewCategoryOld() *CategoryOld {
	p := new(CategoryOld)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.CategoryOld"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type CategoryOldProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Associations []CategoryAssociationSummaryOld `json:"associations,omitempty"`

	CategoryAssociationSummaryOldProjection []CategoryAssociationSummaryOldProjection `json:"categoryAssociationSummaryOldProjection,omitempty"`

	CategorySummaryOldProjection []CategorySummaryOldProjection `json:"categorySummaryOldProjection,omitempty"`
	/*
	  This attribute contains the list of all the categories for which this category is the parent.<br>
	The parentExtId attributes of each child category is set as the extId of this category.<br>
	Note that this list only contains the Summary view of each child category.
	*/
	ChildCategories []CategorySummaryOld `json:"childCategories,omitempty"`
	/*
	  A string consisting of the description of the category as defined by the user.

	The server does not validate this value nor does it enforce the uniqueness or any other constraints.<br>
	It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
	this field.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The fully qualified name of this category. It is unique for each category.<br>
	It is a read-only field.
	The service constructs it from the name-parentExtId combination.
	An example of a fqName would be `Location/Bangalore`, where `Location` is the
	parent category's name and `Bangalore` is the category name.<br>
	This field is immutable.<br>
	*/
	FqName *string `json:"fqName,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Opaque metadata which can be associated to a category.<br>
	It is a list of key-value pairs.<br>
	For example, for a category 'California/SanJose' we can associate a geographical coordinate based metadata
	like: {'latitude': '37.3382° N', 'longitude': '121.8863° W'}.

	The server does not validate this value nor does it enforce the uniqueness or any other constraints.
	It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
	this field.
	*/
	Metadata []import2.KVPair `json:"metadata,omitempty"`
	/*
	  The short name of this category. It may not be unique for each category.<br>
	It is a mandatory field that must be specified inside the post/put request
	body.<br>
	This field is immutable.
	*/
	Name *string `json:"name"`
	/*
	  It is a read-only field inserted into category entity at the time of category creation, and which contains the UUID of
	the user who created this category. It is used for enabling authorization of a particular kind where the user has no
	access to view/create/update/delete any categories other than the category created by oneself.
	*/
	OwnerUuid *string `json:"ownerUuid,omitempty"`
	/*
	  The parent category of this category (may be null if this category is not part of a hierarchy).<br>
	Each category can have at most one parent.<br>
	A parent cannot be deleted until all the children categories are deleted first.<br>
	Must be specified inside the post/put request body for child categories (if not specified, the service assumes
	the category to be a parent category).<br>
	This field is immutable.
	*/
	ParentExtId *string `json:"parentExtId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *CategoryType `json:"type,omitempty"`
	/*
	  The user-specified name is a string that the user can specify; with syntax and semantics controlled by the user.

	The server does not validate this value nor does it enforce the uniqueness or any other constraints.<br>
	It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
	this field. Unlike the name of the categories, which are immutable, the user name can be changed by the user to meet their needs.
	*/
	UserSpecifiedName *string `json:"userSpecifiedName,omitempty"`
}

func NewCategoryOldProjection() *CategoryOldProjection {
	p := new(CategoryOldProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.CategoryOldProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type CategoryProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AssociationDetailProjection []AssociationDetailProjection `json:"associationDetailProjection,omitempty"`

	AssociationSummaryProjection []AssociationSummaryProjection `json:"associationSummaryProjection,omitempty"`
	/*
	  This field gives basic information about resources that are associated with the category.<br>
	The results present under this field summarize the counts of various kinds of resources associated with the category.<br>
	For more detailed information about the UUIDs of the resources, please look into the field `detailedAssociations`.<br>
	This field will be ignored, if given in the payload of `updateCategoryById` or `createCategory` APIs.<br>
	This field will not be present by default in `listCategories` API, unless the parameter $expand=associations is present in the URL.
	*/
	Associations []AssociationSummary `json:"associations,omitempty"`
	/*
	  A string consisting of the description of the category as defined by the user.<br>
	Description can be optionally provided in the payload of `createCategory` and `updateCategoryById` APIs.<br>
	Description field can be updated through `updateCategoryById` API.<br>
	The server does not validate this value nor does it enforce the uniqueness or any other constraints.<br>
	It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
	this field.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  This field gives detailed information about the resources which are associated with the category.<br>
	The results present under this field contain the UUIDs of the entities and policies of various kinds associated with the category.<br>
	This field will be ignored, if given in the payload of `updateCategoryById` or `createCategory` APIs.<br>
	This field will not be present by default in `listCategories` or `getCategoryById` APIs, unless the parameter $expand=detailedAssociations is present in the URL.
	*/
	DetailedAssociations []AssociationDetail `json:"detailedAssociations,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The key of a category when it is represented in `key:value` format.

	Constraints applicable when field is given in the payload during create and update:
	* A string of maxlength of 64
	* Character at the start cannot be `$`
	* Character `/` is not allowed anywhere

	It is a mandatory field in the payload of `createCategory` and `updateCategoryById` APIs.<br>
	This field can't be updated through `updateCategoryById` API.
	*/
	Key *string `json:"key"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  This field contains the UUID of a user who owns the category.<br>
	This field will be ignored if given in the payload of `createCategory` API. Hence, when a category is created, the logged-in user automatically becomes the owner of the category.<br>
	This field can be updated through `updateCategoryById` API, in which case, should be provided, UUID of a valid user is present in the system.<br>
	Validity of the user UUID can be checked by invoking the API: authn/users/{extId} in the 'Identity and Access Management' or 'IAM' namespace.<br>
	It is used for enabling RBAC access to self-owned categories.
	*/
	OwnerUuid *string `json:"ownerUuid,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *CategoryType `json:"type,omitempty"`
	/*
	  The value of a category when it is represented in `key:value` format.

	Constraints applicable when the field is given in the payload during create and update:
	* A string of max length 64
	* Character at the start cannot be `$`
	* Character `/` is not allowed anywhere

	It is a mandatory input field in the payload of `createCategory` and `updateCategoryById` APIs.<br>
	This field can be updated through `updateCategoryById` API.<br>
	Updating the value will not change the extId of the category.
	*/
	Value *string `json:"value"`
}

func (p *CategoryProjection) MarshalJSON() ([]byte, error) {
	type CategoryProjectionProxy CategoryProjection
	return json.Marshal(struct {
		*CategoryProjectionProxy
		Key   *string `json:"key,omitempty"`
		Value *string `json:"value,omitempty"`
	}{
		CategoryProjectionProxy: (*CategoryProjectionProxy)(p),
		Key:                     p.Key,
		Value:                   p.Value,
	})
}

func NewCategoryProjection() *CategoryProjection {
	p := new(CategoryProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.CategoryProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type CategorySummaryOld struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Associations []CategoryAssociationSummaryOld `json:"associations,omitempty"`
	/*
	  This attribute contains the list of all the categories for which this category is the parent.<br>
	The parentExtId attributes of each child category is set as the extId of this category.<br>
	Note that this list only contains the Summary view of each child category.
	*/
	ChildCategories []CategorySummaryOld `json:"childCategories,omitempty"`
	/*
	  A string consisting of the description of the category as defined by the user.

	The server does not validate this value nor does it enforce the uniqueness or any other constraints.<br>
	It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
	this field.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The fully qualified name of this category. It is unique for each category.<br>
	It is a read-only field.
	The service constructs it from the name-parentExtId combination.
	An example of a fqName would be `Location/Bangalore`, where `Location` is the
	parent category's name and `Bangalore` is the category name.<br>
	This field is immutable.<br>
	*/
	FqName *string `json:"fqName,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  The short name of this category. It may not be unique for each category.<br>
	It is a mandatory field that must be specified inside the post/put request
	body.<br>
	This field is immutable.
	*/
	Name *string `json:"name"`
	/*
	  It is a read-only field inserted into category entity at the time of category creation, and which contains the UUID of
	the user who created this category. It is used for enabling authorization of a particular kind where the user has no
	access to view/create/update/delete any categories other than the category created by oneself.
	*/
	OwnerUuid *string `json:"ownerUuid,omitempty"`
	/*
	  The parent category of this category (may be null if this category is not part of a hierarchy).<br>
	Each category can have at most one parent.<br>
	A parent cannot be deleted until all the children categories are deleted first.<br>
	Must be specified inside the post/put request body for child categories (if not specified, the service assumes
	the category to be a parent category).<br>
	This field is immutable.
	*/
	ParentExtId *string `json:"parentExtId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *CategoryType `json:"type,omitempty"`
	/*
	  The user-specified name is a string that the user can specify; with syntax and semantics controlled by the user.

	The server does not validate this value nor does it enforce the uniqueness or any other constraints.<br>
	It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
	this field. Unlike the name of the categories, which are immutable, the user name can be changed by the user to meet their needs.
	*/
	UserSpecifiedName *string `json:"userSpecifiedName,omitempty"`
}

func (p *CategorySummaryOld) MarshalJSON() ([]byte, error) {
	type CategorySummaryOldProxy CategorySummaryOld
	return json.Marshal(struct {
		*CategorySummaryOldProxy
		Name *string `json:"name,omitempty"`
	}{
		CategorySummaryOldProxy: (*CategorySummaryOldProxy)(p),
		Name:                    p.Name,
	})
}

func NewCategorySummaryOld() *CategorySummaryOld {
	p := new(CategorySummaryOld)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.CategorySummaryOld"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type CategorySummaryOldProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Associations []CategoryAssociationSummaryOld `json:"associations,omitempty"`

	CategoryAssociationSummaryOldProjection []CategoryAssociationSummaryOldProjection `json:"categoryAssociationSummaryOldProjection,omitempty"`

	CategorySummaryOldProjection []CategorySummaryOldProjection `json:"categorySummaryOldProjection,omitempty"`
	/*
	  This attribute contains the list of all the categories for which this category is the parent.<br>
	The parentExtId attributes of each child category is set as the extId of this category.<br>
	Note that this list only contains the Summary view of each child category.
	*/
	ChildCategories []CategorySummaryOld `json:"childCategories,omitempty"`
	/*
	  A string consisting of the description of the category as defined by the user.

	The server does not validate this value nor does it enforce the uniqueness or any other constraints.<br>
	It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
	this field.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The fully qualified name of this category. It is unique for each category.<br>
	It is a read-only field.
	The service constructs it from the name-parentExtId combination.
	An example of a fqName would be `Location/Bangalore`, where `Location` is the
	parent category's name and `Bangalore` is the category name.<br>
	This field is immutable.<br>
	*/
	FqName *string `json:"fqName,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  The short name of this category. It may not be unique for each category.<br>
	It is a mandatory field that must be specified inside the post/put request
	body.<br>
	This field is immutable.
	*/
	Name *string `json:"name"`
	/*
	  It is a read-only field inserted into category entity at the time of category creation, and which contains the UUID of
	the user who created this category. It is used for enabling authorization of a particular kind where the user has no
	access to view/create/update/delete any categories other than the category created by oneself.
	*/
	OwnerUuid *string `json:"ownerUuid,omitempty"`
	/*
	  The parent category of this category (may be null if this category is not part of a hierarchy).<br>
	Each category can have at most one parent.<br>
	A parent cannot be deleted until all the children categories are deleted first.<br>
	Must be specified inside the post/put request body for child categories (if not specified, the service assumes
	the category to be a parent category).<br>
	This field is immutable.
	*/
	ParentExtId *string `json:"parentExtId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *CategoryType `json:"type,omitempty"`
	/*
	  The user-specified name is a string that the user can specify; with syntax and semantics controlled by the user.

	The server does not validate this value nor does it enforce the uniqueness or any other constraints.<br>
	It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
	this field. Unlike the name of the categories, which are immutable, the user name can be changed by the user to meet their needs.
	*/
	UserSpecifiedName *string `json:"userSpecifiedName,omitempty"`
}

func (p *CategorySummaryOldProjection) MarshalJSON() ([]byte, error) {
	type CategorySummaryOldProjectionProxy CategorySummaryOldProjection
	return json.Marshal(struct {
		*CategorySummaryOldProjectionProxy
		Name *string `json:"name,omitempty"`
	}{
		CategorySummaryOldProjectionProxy: (*CategorySummaryOldProjectionProxy)(p),
		Name:                              p.Name,
	})
}

func NewCategorySummaryOldProjection() *CategorySummaryOldProjection {
	p := new(CategorySummaryOldProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.CategorySummaryOldProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Denotes the type of a category.<br>
There are three types of categories: SYSTEM, INTERNAL, and USER.<br>
This field is immutable.
*/
type CategoryType int

const (
	CATEGORYTYPE_UNKNOWN  CategoryType = 0
	CATEGORYTYPE_REDACTED CategoryType = 1
	CATEGORYTYPE_USER     CategoryType = 2
	CATEGORYTYPE_SYSTEM   CategoryType = 3
	CATEGORYTYPE_INTERNAL CategoryType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *CategoryType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"USER",
		"SYSTEM",
		"INTERNAL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e CategoryType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"USER",
		"SYSTEM",
		"INTERNAL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *CategoryType) index(name string) CategoryType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"USER",
		"SYSTEM",
		"INTERNAL",
	}
	for idx := range names {
		if names[idx] == name {
			return CategoryType(idx)
		}
	}
	return CATEGORYTYPE_UNKNOWN
}

func (e *CategoryType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CategoryType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CategoryType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CategoryType) Ref() *CategoryType {
	return &e
}

/*
REST response for all response codes in API path /prism/v4.0/config/categories Post operation
*/
type CreateCategoryApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateCategoryApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateCategoryApiResponse() *CreateCategoryApiResponse {
	p := new(CreateCategoryApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.CreateCategoryApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateCategoryApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateCategoryApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateCategoryApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /prism/v4.0/config/domain-managers Post operation
*/
type CreateDomainManagerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateDomainManagerApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateDomainManagerApiResponse() *CreateDomainManagerApiResponse {
	p := new(CreateDomainManagerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.CreateDomainManagerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateDomainManagerApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateDomainManagerApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateDomainManagerApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /prism/v4.0/config/categories/{extId} Delete operation
*/
type DeleteCategoryApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteCategoryApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteCategoryApiResponse() *DeleteCategoryApiResponse {
	p := new(DeleteCategoryApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.DeleteCategoryApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteCategoryApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteCategoryApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteCategoryApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
Domain manager (Prism Central) details.
*/
type DomainManager struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Config *DomainManagerClusterConfig `json:"config"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The external identifier of the cluster hosting the domain manager (Prism Central) instance.
	*/
	HostingClusterExtId *string `json:"hostingClusterExtId,omitempty"`
	/*
	  Boolean value indicating if the domain manager (Prism Central) is registered with the hosting cluster, that is, Prism Element.
	*/
	IsRegisteredWithHostingCluster *bool `json:"isRegisteredWithHostingCluster,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	Network *DomainManagerNetwork `json:"network"`
	/*
	  Domain manager (Prism Central) nodes external identifier.
	*/
	NodeExtIds []string `json:"nodeExtIds,omitempty"`
	/*
	  This configuration enables Prism Central to be deployed in scale-out mode.
	*/
	ShouldEnableHighAvailability *bool `json:"shouldEnableHighAvailability,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *DomainManager) MarshalJSON() ([]byte, error) {
	type DomainManagerProxy DomainManager
	return json.Marshal(struct {
		*DomainManagerProxy
		Config  *DomainManagerClusterConfig `json:"config,omitempty"`
		Network *DomainManagerNetwork       `json:"network,omitempty"`
	}{
		DomainManagerProxy: (*DomainManagerProxy)(p),
		Config:             p.Config,
		Network:            p.Network,
	})
}

func NewDomainManager() *DomainManager {
	p := new(DomainManager)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.DomainManager"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldEnableHighAvailability = new(bool)
	*p.ShouldEnableHighAvailability = false

	return p
}

/*
Domain manager (Prism Central) cluster configuration details.
*/
type DomainManagerClusterConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BootstrapConfig *BootstrapConfig `json:"bootstrapConfig,omitempty"`

	BuildInfo *import5.BuildInfo `json:"buildInfo,omitempty"`
	/*
	  The credentials consist of a username and password for a particular user like admin. Users can pass the credentials of admin users currently which will be configured in the create domain manager operation.
	*/
	Credentials []import2.BasicAuth `json:"credentials,omitempty"`
	/*
	  Name of the domain manager (Prism Central).
	*/
	Name *string `json:"name,omitempty"`

	ResourceConfig *DomainManagerResourceConfig `json:"resourceConfig,omitempty"`
	/*
	  A boolean value indicating whether to enable lockdown mode for a cluster.
	*/
	ShouldEnableLockdownMode *bool `json:"shouldEnableLockdownMode,omitempty"`

	Size *Size `json:"size,omitempty"`
}

func NewDomainManagerClusterConfig() *DomainManagerClusterConfig {
	p := new(DomainManagerClusterConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.DomainManagerClusterConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Domain manager (Prism Central) network configuration details.
*/
type DomainManagerNetwork struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ExternalAddress *import2.IPAddress `json:"externalAddress,omitempty"`
	/*
	  This configuration is used to manage Prism Central.
	*/
	ExternalNetworks []ExternalNetwork `json:"externalNetworks,omitempty"`
	/*
	  Cluster fully qualified domain name. This is part of payload for cluster update operation only.
	*/
	Fqdn *string `json:"fqdn,omitempty"`
	/*
	  This configuration is used to internally manage Prism Central network.
	*/
	InternalNetworks []BaseNetwork `json:"internalNetworks,omitempty"`
	/*
	  List of name servers on a cluster. This is part of payload for both cluster create & update operations. For create operation, only ipv4 address / fqdn values are supported currently.
	*/
	NameServers []import2.IPAddressOrFQDN `json:"nameServers,omitempty"`
	/*
	  List of NTP servers on a cluster. This is part of payload for both cluster create & update operations. For create operation, only ipv4 address / fqdn values are supported currently.
	*/
	NtpServers []import2.IPAddressOrFQDN `json:"ntpServers,omitempty"`
}

func NewDomainManagerNetwork() *DomainManagerNetwork {
	p := new(DomainManagerNetwork)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.DomainManagerNetwork"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
This configuration is used to provide the resource-related details like container external identifiers, number of VCPUs, memory size, data disk size of the domain manager (Prism Central). In the case of a multi-node setup, the sum of resources like number of VCPUs, memory size and data disk size are provided.
*/
type DomainManagerResourceConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the container that will be used to create the domain manager (Prism Central) cluster.
	*/
	ContainerExtIds []string `json:"containerExtIds,omitempty"`
	/*
	  This property is used for readOnly purposes to display Prism Central data disk size allocation at a cluster level.
	*/
	DataDiskSizeBytes *int64 `json:"dataDiskSizeBytes,omitempty"`
	/*
	  This property is used for readOnly purposes to display Prism Central RAM allocation at the cluster level.
	*/
	MemorySizeBytes *int64 `json:"memorySizeBytes,omitempty"`
	/*
	  This property is used for readOnly purposes to display Prism Central number of VCPUs allocation.
	*/
	NumVcpus *int `json:"numVcpus,omitempty"`
}

func NewDomainManagerResourceConfig() *DomainManagerResourceConfig {
	p := new(DomainManagerResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.DomainManagerResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Details of the entity.
*/
type EntityReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of the entity.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Name of the entity.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Entity type identified as 'namespace:module[:submodule]:entityType'. For example - vmm:ahv:vm, where vmm is the namepsace, ahv is the module and vm is the entitytype.
	*/
	Rel *string `json:"rel,omitempty"`
}

func NewEntityReference() *EntityReference {
	p := new(EntityReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.EntityReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
An object denoting the environment information of the PC. It contains the following fields:<br>
- type: Enums denoting the environment type of the PC.<br>
- providerType: Enums denoting the provider of the cloud PC.<br>
- instanceObj: Enums denoting the instance type of the cloud PC.<br>
*/
type EnvironmentInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ProviderType *ProviderType `json:"providerType,omitempty"`

	ProvisioningType *ProvisioningType `json:"provisioningType,omitempty"`

	Type *EnvironmentType `json:"type,omitempty"`
}

func NewEnvironmentInfo() *EnvironmentInfo {
	p := new(EnvironmentInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.EnvironmentInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Enums denoting the environment type of the PC, that is, on-prem PC or cloud PC.<br>
Following are the supported entity types:
- ONPREM
- NTNX_CLOUD
*/
type EnvironmentType int

const (
	ENVIRONMENTTYPE_UNKNOWN    EnvironmentType = 0
	ENVIRONMENTTYPE_REDACTED   EnvironmentType = 1
	ENVIRONMENTTYPE_ONPREM     EnvironmentType = 2
	ENVIRONMENTTYPE_NTNX_CLOUD EnvironmentType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EnvironmentType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ONPREM",
		"NTNX_CLOUD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e EnvironmentType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ONPREM",
		"NTNX_CLOUD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *EnvironmentType) index(name string) EnvironmentType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ONPREM",
		"NTNX_CLOUD",
	}
	for idx := range names {
		if names[idx] == name {
			return EnvironmentType(idx)
		}
	}
	return ENVIRONMENTTYPE_UNKNOWN
}

func (e *EnvironmentType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EnvironmentType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EnvironmentType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EnvironmentType) Ref() *EnvironmentType {
	return &e
}

/*
This configuration is used to manage Prism Central.
*/
type ExternalNetwork struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DefaultGateway *import2.IPAddressOrFQDN `json:"defaultGateway"`
	/*
	  Range of IPs used for Prism Central network setup.
	*/
	IpRanges []import2.IpRange `json:"ipRanges"`
	/*
	  The network external identifier to which Domain Manager (Prism Central) is to be deployed or is already configured.
	*/
	NetworkExtId *string `json:"networkExtId"`

	SubnetMask *import2.IPAddressOrFQDN `json:"subnetMask"`
}

func (p *ExternalNetwork) MarshalJSON() ([]byte, error) {
	type ExternalNetworkProxy ExternalNetwork
	return json.Marshal(struct {
		*ExternalNetworkProxy
		DefaultGateway *import2.IPAddressOrFQDN `json:"defaultGateway,omitempty"`
		IpRanges       []import2.IpRange        `json:"ipRanges,omitempty"`
		NetworkExtId   *string                  `json:"networkExtId,omitempty"`
		SubnetMask     *import2.IPAddressOrFQDN `json:"subnetMask,omitempty"`
	}{
		ExternalNetworkProxy: (*ExternalNetworkProxy)(p),
		DefaultGateway:       p.DefaultGateway,
		IpRanges:             p.IpRanges,
		NetworkExtId:         p.NetworkExtId,
		SubnetMask:           p.SubnetMask,
	})
}

func NewExternalNetwork() *ExternalNetwork {
	p := new(ExternalNetwork)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.ExternalNetwork"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0/config/categories/{extId} Get operation
*/
type GetCategoryApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetCategoryApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetCategoryApiResponse() *GetCategoryApiResponse {
	p := new(GetCategoryApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.GetCategoryApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetCategoryApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetCategoryApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetCategoryApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /prism/v4.0/config/domain-managers/{extId} Get operation
*/
type GetDomainManagerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetDomainManagerApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetDomainManagerApiResponse() *GetDomainManagerApiResponse {
	p := new(GetDomainManagerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.GetDomainManagerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetDomainManagerApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetDomainManagerApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetDomainManagerApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /prism/v4.0/config/tasks/{extId} Get operation
*/
type GetTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetTaskApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetTaskApiResponse() *GetTaskApiResponse {
	p := new(GetTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.GetTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetTaskApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /prism/v4.0/config/categories Get operation
*/
type ListCategoriesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListCategoriesApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListCategoriesApiResponse() *ListCategoriesApiResponse {
	p := new(ListCategoriesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.ListCategoriesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListCategoriesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListCategoriesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListCategoriesApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /prism/v4.0/config/domain-managers Get operation
*/
type ListDomainManagerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListDomainManagerApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListDomainManagerApiResponse() *ListDomainManagerApiResponse {
	p := new(ListDomainManagerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.ListDomainManagerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListDomainManagerApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListDomainManagerApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListDomainManagerApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /prism/v4.0/config/tasks Get operation
*/
type ListTasksApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListTasksApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListTasksApiResponse() *ListTasksApiResponse {
	p := new(ListTasksApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.ListTasksApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListTasksApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListTasksApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListTasksApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
Reference to the owner of the task.
*/
type OwnerReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of the task owner.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Username of the task owner.
	*/
	Name *string `json:"name,omitempty"`
}

func NewOwnerReference() *OwnerReference {
	p := new(OwnerReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.OwnerReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Enums denoting the provider of the cloud, in case of environment type a cloud PC.<br>
The service currently supports the following providers:
- NTNX
- AZURE
- AWS
- GCP
- VSPHERE
*/
type ProviderType int

const (
	PROVIDERTYPE_UNKNOWN  ProviderType = 0
	PROVIDERTYPE_REDACTED ProviderType = 1
	PROVIDERTYPE_NTNX     ProviderType = 2
	PROVIDERTYPE_AZURE    ProviderType = 3
	PROVIDERTYPE_AWS      ProviderType = 4
	PROVIDERTYPE_GCP      ProviderType = 5
	PROVIDERTYPE_VSPHERE  ProviderType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ProviderType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTNX",
		"AZURE",
		"AWS",
		"GCP",
		"VSPHERE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ProviderType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTNX",
		"AZURE",
		"AWS",
		"GCP",
		"VSPHERE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ProviderType) index(name string) ProviderType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTNX",
		"AZURE",
		"AWS",
		"GCP",
		"VSPHERE",
	}
	for idx := range names {
		if names[idx] == name {
			return ProviderType(idx)
		}
	}
	return PROVIDERTYPE_UNKNOWN
}

func (e *ProviderType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ProviderType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ProviderType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ProviderType) Ref() *ProviderType {
	return &e
}

/*
Enums denoting the instance type of the cloud PC. It indicates whether the PC is created on bare-metal
or on a cloud-provisioned VM. Hence, it supports two possible values:
- NTNX
- NATIVE
*/
type ProvisioningType int

const (
	PROVISIONINGTYPE_UNKNOWN  ProvisioningType = 0
	PROVISIONINGTYPE_REDACTED ProvisioningType = 1
	PROVISIONINGTYPE_NTNX     ProvisioningType = 2
	PROVISIONINGTYPE_NATIVE   ProvisioningType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ProvisioningType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTNX",
		"NATIVE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ProvisioningType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTNX",
		"NATIVE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ProvisioningType) index(name string) ProvisioningType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTNX",
		"NATIVE",
	}
	for idx := range names {
		if names[idx] == name {
			return ProvisioningType(idx)
		}
	}
	return PROVISIONINGTYPE_UNKNOWN
}

func (e *ProvisioningType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ProvisioningType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ProvisioningType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ProvisioningType) Ref() *ProvisioningType {
	return &e
}

/*
Contains references for entities given in EntityAssociation.
This contains the entity ID and a list of links to fetch the associated entities.
*/
type Reference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the resource which uniquely identifies it.
	*/
	ResourceId *string `json:"resourceId,omitempty"`
}

func NewReference() *Reference {
	p := new(Reference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.Reference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
An enum denoting the resource group.<br>
Resources can be organised into either an entity or a policy. Hence, it supports two possible values:
  * ENTITY
  * POLICY
*/
type ResourceGroup int

const (
	RESOURCEGROUP_UNKNOWN  ResourceGroup = 0
	RESOURCEGROUP_REDACTED ResourceGroup = 1
	RESOURCEGROUP_ENTITY   ResourceGroup = 2
	RESOURCEGROUP_POLICY   ResourceGroup = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ResourceGroup) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENTITY",
		"POLICY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ResourceGroup) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENTITY",
		"POLICY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ResourceGroup) index(name string) ResourceGroup {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENTITY",
		"POLICY",
	}
	for idx := range names {
		if names[idx] == name {
			return ResourceGroup(idx)
		}
	}
	return RESOURCEGROUP_UNKNOWN
}

func (e *ResourceGroup) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ResourceGroup:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ResourceGroup) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ResourceGroup) Ref() *ResourceGroup {
	return &e
}

/*
An enum denoting the associated resource types.<br>
Resource types are further grouped into two types: entity or policy.
*/
type ResourceType int

const (
	RESOURCETYPE_UNKNOWN                    ResourceType = 0
	RESOURCETYPE_REDACTED                   ResourceType = 1
	RESOURCETYPE_VM                         ResourceType = 2
	RESOURCETYPE_MH_VM                      ResourceType = 3
	RESOURCETYPE_IMAGE                      ResourceType = 4
	RESOURCETYPE_SUBNET                     ResourceType = 5
	RESOURCETYPE_CLUSTER                    ResourceType = 6
	RESOURCETYPE_HOST                       ResourceType = 7
	RESOURCETYPE_REPORT                     ResourceType = 8
	RESOURCETYPE_MARKETPLACE_ITEM           ResourceType = 9
	RESOURCETYPE_BLUEPRINT                  ResourceType = 10
	RESOURCETYPE_APP                        ResourceType = 11
	RESOURCETYPE_VOLUMEGROUP                ResourceType = 12
	RESOURCETYPE_IMAGE_PLACEMENT_POLICY     ResourceType = 13
	RESOURCETYPE_NETWORK_SECURITY_POLICY    ResourceType = 14
	RESOURCETYPE_NETWORK_SECURITY_RULE      ResourceType = 15
	RESOURCETYPE_VM_HOST_AFFINITY_POLICY    ResourceType = 16
	RESOURCETYPE_VM_VM_ANTI_AFFINITY_POLICY ResourceType = 17
	RESOURCETYPE_QOS_POLICY                 ResourceType = 18
	RESOURCETYPE_NGT_POLICY                 ResourceType = 19
	RESOURCETYPE_PROTECTION_RULE            ResourceType = 20
	RESOURCETYPE_ACCESS_CONTROL_POLICY      ResourceType = 21
	RESOURCETYPE_STORAGE_POLICY             ResourceType = 22
	RESOURCETYPE_IMAGE_RATE_LIMIT           ResourceType = 23
	RESOURCETYPE_RECOVERY_PLAN              ResourceType = 24
	RESOURCETYPE_BUNDLE                     ResourceType = 25
	RESOURCETYPE_POLICY_SCHEMA              ResourceType = 26
	RESOURCETYPE_HOST_NIC                   ResourceType = 27
	RESOURCETYPE_ACTION_RULE                ResourceType = 28
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ResourceType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"MH_VM",
		"IMAGE",
		"SUBNET",
		"CLUSTER",
		"HOST",
		"REPORT",
		"MARKETPLACE_ITEM",
		"BLUEPRINT",
		"APP",
		"VOLUMEGROUP",
		"IMAGE_PLACEMENT_POLICY",
		"NETWORK_SECURITY_POLICY",
		"NETWORK_SECURITY_RULE",
		"VM_HOST_AFFINITY_POLICY",
		"VM_VM_ANTI_AFFINITY_POLICY",
		"QOS_POLICY",
		"NGT_POLICY",
		"PROTECTION_RULE",
		"ACCESS_CONTROL_POLICY",
		"STORAGE_POLICY",
		"IMAGE_RATE_LIMIT",
		"RECOVERY_PLAN",
		"BUNDLE",
		"POLICY_SCHEMA",
		"HOST_NIC",
		"ACTION_RULE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ResourceType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"MH_VM",
		"IMAGE",
		"SUBNET",
		"CLUSTER",
		"HOST",
		"REPORT",
		"MARKETPLACE_ITEM",
		"BLUEPRINT",
		"APP",
		"VOLUMEGROUP",
		"IMAGE_PLACEMENT_POLICY",
		"NETWORK_SECURITY_POLICY",
		"NETWORK_SECURITY_RULE",
		"VM_HOST_AFFINITY_POLICY",
		"VM_VM_ANTI_AFFINITY_POLICY",
		"QOS_POLICY",
		"NGT_POLICY",
		"PROTECTION_RULE",
		"ACCESS_CONTROL_POLICY",
		"STORAGE_POLICY",
		"IMAGE_RATE_LIMIT",
		"RECOVERY_PLAN",
		"BUNDLE",
		"POLICY_SCHEMA",
		"HOST_NIC",
		"ACTION_RULE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ResourceType) index(name string) ResourceType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"MH_VM",
		"IMAGE",
		"SUBNET",
		"CLUSTER",
		"HOST",
		"REPORT",
		"MARKETPLACE_ITEM",
		"BLUEPRINT",
		"APP",
		"VOLUMEGROUP",
		"IMAGE_PLACEMENT_POLICY",
		"NETWORK_SECURITY_POLICY",
		"NETWORK_SECURITY_RULE",
		"VM_HOST_AFFINITY_POLICY",
		"VM_VM_ANTI_AFFINITY_POLICY",
		"QOS_POLICY",
		"NGT_POLICY",
		"PROTECTION_RULE",
		"ACCESS_CONTROL_POLICY",
		"STORAGE_POLICY",
		"IMAGE_RATE_LIMIT",
		"RECOVERY_PLAN",
		"BUNDLE",
		"POLICY_SCHEMA",
		"HOST_NIC",
		"ACTION_RULE",
	}
	for idx := range names {
		if names[idx] == name {
			return ResourceType(idx)
		}
	}
	return RESOURCETYPE_UNKNOWN
}

func (e *ResourceType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ResourceType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ResourceType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ResourceType) Ref() *ResourceType {
	return &e
}

/*
Domain manager (Prism Central) size is an enumeration of starter, small, large, or extra large starter values.
*/
type Size int

const (
	SIZE_UNKNOWN    Size = 0
	SIZE_REDACTED   Size = 1
	SIZE_STARTER    Size = 2
	SIZE_SMALL      Size = 3
	SIZE_LARGE      Size = 4
	SIZE_EXTRALARGE Size = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Size) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STARTER",
		"SMALL",
		"LARGE",
		"EXTRALARGE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Size) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STARTER",
		"SMALL",
		"LARGE",
		"EXTRALARGE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Size) index(name string) Size {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STARTER",
		"SMALL",
		"LARGE",
		"EXTRALARGE",
	}
	for idx := range names {
		if names[idx] == name {
			return Size(idx)
		}
	}
	return SIZE_UNKNOWN
}

func (e *Size) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Size:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Size) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Size) Ref() *Size {
	return &e
}

/*
The task object tracking an asynchronous operation.
*/
type Task struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of globally unique identifiers for clusters associated with the task or any of its subtasks.
	*/
	ClusterExtIds []string `json:"clusterExtIds,omitempty"`
	/*
	  UTC date and time in RFC-3339 format when the task was completed.
	*/
	CompletedTime *time.Time `json:"completedTime,omitempty"`
	/*
	  Additional details on the task to aid the user with further actions post completion of the task.
	*/
	CompletionDetails []import2.KVPair `json:"completionDetails,omitempty"`
	/*
	  UTC date and time in RFC-3339 format when the task was created.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Reference to entities associated with the task.
	*/
	EntitiesAffected []EntityReference `json:"entitiesAffected,omitempty"`
	/*
	  Error details explaining a task failure. These would be populated only in the case of task failures.
	*/
	ErrorMessages []import4.AppMessage `json:"errorMessages,omitempty"`
	/*
	  A globally unique identifier of a task.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Signifies if the task is a background task or not.
	*/
	IsBackgroundTask *bool `json:"isBackgroundTask,omitempty"`
	/*
	  Signifies if the task can be cancelled.
	*/
	IsCancelable *bool `json:"isCancelable,omitempty"`
	/*
	  UTC date and time in RFC-3339 format when the task was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  Provides an error message in the absence of a well-defined error message for the tasks created through legacy APIs.
	*/
	LegacyErrorMessage *string `json:"legacyErrorMessage,omitempty"`
	/*
	  Number of entities associated with the task.
	*/
	NumberOfEntitiesAffected *int `json:"numberOfEntitiesAffected,omitempty"`
	/*
	  Number of tasks spawned as children of the current task.
	*/
	NumberOfSubtasks *int `json:"numberOfSubtasks,omitempty"`
	/*
	  The operation name being tracked by the task.
	*/
	Operation *string `json:"operation,omitempty"`
	/*
	  Description of the operation being tracked by the task.
	*/
	OperationDescription *string `json:"operationDescription,omitempty"`

	OwnedBy *OwnerReference `json:"ownedBy,omitempty"`

	ParentTask *TaskReferenceInternal `json:"parentTask,omitempty"`
	/*
	  Task progress expressed as a percentage.
	*/
	ProgressPercentage *int `json:"progressPercentage,omitempty"`

	RootTask *TaskReferenceInternal `json:"rootTask,omitempty"`
	/*
	  UTC date and time in RFC-3339 format when the task was started.
	*/
	StartedTime *time.Time `json:"startedTime,omitempty"`

	Status *TaskStatus `json:"status,omitempty"`
	/*
	  List of steps completed as part of the task.
	*/
	SubSteps []TaskStep `json:"subSteps,omitempty"`
	/*
	  Reference to tasks spawned as children of the current task. The task get response would contain a limited number of subtask references. To get the entire list of subtasks for a task, use the parent task filter in the task list API.
	*/
	SubTasks []TaskReferenceInternal `json:"subTasks,omitempty"`
	/*
	  Warning messages to alert the user of issues which did not directly cause task failure. These can be populated for any task.
	*/
	Warnings []import4.AppMessage `json:"warnings,omitempty"`
}

func NewTask() *Task {
	p := new(Task)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.Task"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A reference to a task tracking an asynchronous operation. The status of the task can be queried by making a GET request to the task URI provided in the metadata section of the API response.
*/
type TaskReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of a task.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func NewTaskReference() *TaskReference {
	p := new(TaskReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.TaskReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to the parent task associated with the current task.
*/
type TaskReferenceInternal struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of the task.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The URL at which the entity described by the link can be accessed.
	*/
	Href *string `json:"href,omitempty"`
	/*
	  A name that identifies the relationship of the link to the object that is returned by the URL.  The unique value of "self" identifies the URL for the object.
	*/
	Rel *string `json:"rel,omitempty"`
}

func NewTaskReferenceInternal() *TaskReferenceInternal {
	p := new(TaskReferenceInternal)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.TaskReferenceInternal"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Status of the task.
*/
type TaskStatus int

const (
	TASKSTATUS_UNKNOWN   TaskStatus = 0
	TASKSTATUS_REDACTED  TaskStatus = 1
	TASKSTATUS_QUEUED    TaskStatus = 2
	TASKSTATUS_RUNNING   TaskStatus = 3
	TASKSTATUS_CANCELING TaskStatus = 4
	TASKSTATUS_SUCCEEDED TaskStatus = 5
	TASKSTATUS_FAILED    TaskStatus = 6
	TASKSTATUS_CANCELED  TaskStatus = 7
	TASKSTATUS_SUSPENDED TaskStatus = 8
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *TaskStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"QUEUED",
		"RUNNING",
		"CANCELING",
		"SUCCEEDED",
		"FAILED",
		"CANCELED",
		"SUSPENDED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e TaskStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"QUEUED",
		"RUNNING",
		"CANCELING",
		"SUCCEEDED",
		"FAILED",
		"CANCELED",
		"SUSPENDED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *TaskStatus) index(name string) TaskStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"QUEUED",
		"RUNNING",
		"CANCELING",
		"SUCCEEDED",
		"FAILED",
		"CANCELED",
		"SUSPENDED",
	}
	for idx := range names {
		if names[idx] == name {
			return TaskStatus(idx)
		}
	}
	return TASKSTATUS_UNKNOWN
}

func (e *TaskStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for TaskStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *TaskStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e TaskStatus) Ref() *TaskStatus {
	return &e
}

/*
A single step in the task.
*/
type TaskStep struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Message describing the completed steps for the task.
	*/
	Name *string `json:"name,omitempty"`
}

func NewTaskStep() *TaskStep {
	p := new(TaskStep)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.TaskStep"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0/config/categories/{extId} Put operation
*/
type UpdateCategoryApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateCategoryApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateCategoryApiResponse() *UpdateCategoryApiResponse {
	p := new(UpdateCategoryApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.UpdateCategoryApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateCategoryApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateCategoryApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateCategoryApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

type OneOfCancelTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 *import4.AppMessage    `json:"-"`
}

func NewOneOfCancelTaskApiResponseData() *OneOfCancelTaskApiResponseData {
	p := new(OneOfCancelTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCancelTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCancelTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import4.AppMessage:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import4.AppMessage)
		}
		*p.oneOfType2001 = v.(import4.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCancelTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfCancelTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new(import4.AppMessage)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.error.AppMessage" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import4.AppMessage)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCancelTaskApiResponseData"))
}

func (p *OneOfCancelTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfCancelTaskApiResponseData")
}

type OneOfGetTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 *Task                  `json:"-"`
}

func NewOneOfGetTaskApiResponseData() *OneOfGetTaskApiResponseData {
	p := new(OneOfGetTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Task:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Task)
		}
		*p.oneOfType2001 = v.(Task)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new(Task)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.Task" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Task)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetTaskApiResponseData"))
}

func (p *OneOfGetTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetTaskApiResponseData")
}

type OneOfUpdateCategoryApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []import4.AppMessage   `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfUpdateCategoryApiResponseData() *OneOfUpdateCategoryApiResponseData {
	p := new(OneOfUpdateCategoryApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateCategoryApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateCategoryApiResponseData is nil"))
	}
	switch v.(type) {
	case []import4.AppMessage:
		p.oneOfType0 = v.([]import4.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<prism.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.error.AppMessage>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateCategoryApiResponseData) GetValue() interface{} {
	if "List<prism.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateCategoryApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]import4.AppMessage)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "prism.v4.error.AppMessage" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.error.AppMessage>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateCategoryApiResponseData"))
}

func (p *OneOfUpdateCategoryApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<prism.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateCategoryApiResponseData")
}

type OneOfDeleteCategoryApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfDeleteCategoryApiResponseData() *OneOfDeleteCategoryApiResponseData {
	p := new(OneOfDeleteCategoryApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteCategoryApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteCategoryApiResponseData is nil"))
	}
	if nil == v {
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(interface{})
		}
		*p.oneOfType1 = nil
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "EMPTY"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "EMPTY"
		return nil
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDeleteCategoryApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteCategoryApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(interface{})
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if nil == *vOneOfType1 {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(interface{})
			}
			*p.oneOfType1 = nil
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "EMPTY"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "EMPTY"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteCategoryApiResponseData"))
}

func (p *OneOfDeleteCategoryApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteCategoryApiResponseData")
}

type OneOfGetDomainManagerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 *DomainManager         `json:"-"`
}

func NewOneOfGetDomainManagerApiResponseData() *OneOfGetDomainManagerApiResponseData {
	p := new(OneOfGetDomainManagerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetDomainManagerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetDomainManagerApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case DomainManager:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(DomainManager)
		}
		*p.oneOfType2001 = v.(DomainManager)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetDomainManagerApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetDomainManagerApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new(DomainManager)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.DomainManager" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(DomainManager)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetDomainManagerApiResponseData"))
}

func (p *OneOfGetDomainManagerApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetDomainManagerApiResponseData")
}

type OneOfGetCategoryApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Category              `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetCategoryApiResponseData() *OneOfGetCategoryApiResponseData {
	p := new(OneOfGetCategoryApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetCategoryApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetCategoryApiResponseData is nil"))
	}
	switch v.(type) {
	case Category:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Category)
		}
		*p.oneOfType0 = v.(Category)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetCategoryApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetCategoryApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Category)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.Category" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Category)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetCategoryApiResponseData"))
}

func (p *OneOfGetCategoryApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetCategoryApiResponseData")
}

type OneOfListTasksApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 []Task                 `json:"-"`
}

func NewOneOfListTasksApiResponseData() *OneOfListTasksApiResponseData {
	p := new(OneOfListTasksApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListTasksApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListTasksApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []Task:
		p.oneOfType2001 = v.([]Task)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<prism.v4.config.Task>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.config.Task>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListTasksApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<prism.v4.config.Task>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListTasksApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new([]Task)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "prism.v4.config.Task" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.config.Task>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.config.Task>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListTasksApiResponseData"))
}

func (p *OneOfListTasksApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<prism.v4.config.Task>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListTasksApiResponseData")
}

type OneOfCreateDomainManagerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *TaskReference         `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCreateDomainManagerApiResponseData() *OneOfCreateDomainManagerApiResponseData {
	p := new(OneOfCreateDomainManagerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateDomainManagerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateDomainManagerApiResponseData is nil"))
	}
	switch v.(type) {
	case TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(TaskReference)
		}
		*p.oneOfType2001 = v.(TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCreateDomainManagerApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateDomainManagerApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(TaskReference)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateDomainManagerApiResponseData"))
}

func (p *OneOfCreateDomainManagerApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateDomainManagerApiResponseData")
}

type OneOfCreateCategoryApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Category              `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCreateCategoryApiResponseData() *OneOfCreateCategoryApiResponseData {
	p := new(OneOfCreateCategoryApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateCategoryApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateCategoryApiResponseData is nil"))
	}
	switch v.(type) {
	case Category:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Category)
		}
		*p.oneOfType0 = v.(Category)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCreateCategoryApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateCategoryApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Category)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.Category" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Category)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateCategoryApiResponseData"))
}

func (p *OneOfCreateCategoryApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateCategoryApiResponseData")
}

type OneOfListCategoriesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType401  []CategoryProjection   `json:"-"`
	oneOfType0    []Category             `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListCategoriesApiResponseData() *OneOfListCategoriesApiResponseData {
	p := new(OneOfListCategoriesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListCategoriesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListCategoriesApiResponseData is nil"))
	}
	switch v.(type) {
	case []CategoryProjection:
		p.oneOfType401 = v.([]CategoryProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<prism.v4.config.CategoryProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.config.CategoryProjection>"
	case []Category:
		p.oneOfType0 = v.([]Category)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<prism.v4.config.Category>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.config.Category>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListCategoriesApiResponseData) GetValue() interface{} {
	if "List<prism.v4.config.CategoryProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if "List<prism.v4.config.Category>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListCategoriesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType401 := new([]CategoryProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "prism.v4.config.CategoryProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.config.CategoryProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.config.CategoryProjection>"
			return nil
		}
	}
	vOneOfType0 := new([]Category)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "prism.v4.config.Category" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.config.Category>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.config.Category>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListCategoriesApiResponseData"))
}

func (p *OneOfListCategoriesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<prism.v4.config.CategoryProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if "List<prism.v4.config.Category>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListCategoriesApiResponseData")
}

type OneOfListDomainManagerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 []DomainManager        `json:"-"`
}

func NewOneOfListDomainManagerApiResponseData() *OneOfListDomainManagerApiResponseData {
	p := new(OneOfListDomainManagerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListDomainManagerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListDomainManagerApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []DomainManager:
		p.oneOfType2001 = v.([]DomainManager)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<prism.v4.config.DomainManager>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.config.DomainManager>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListDomainManagerApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<prism.v4.config.DomainManager>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListDomainManagerApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new([]DomainManager)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "prism.v4.config.DomainManager" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.config.DomainManager>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.config.DomainManager>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListDomainManagerApiResponseData"))
}

func (p *OneOfListDomainManagerApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<prism.v4.config.DomainManager>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListDomainManagerApiResponseData")
}

type FileDetail struct {
	Path        *string `json:"-"`
	ObjectType_ *string `json:"-"`
}

func NewFileDetail() *FileDetail {
	p := new(FileDetail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "FileDetail"

	return p
}
