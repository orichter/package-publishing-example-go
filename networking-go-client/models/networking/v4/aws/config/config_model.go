/*
 * Generated file models/networking/v4/aws/config/config_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix Networking APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module networking.v4.aws.config of Nutanix Networking APIs
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/orichter/package-publishing-example-go/networking-go-client/v4/models/common/v1/response"
	import2 "github.com/orichter/package-publishing-example-go/networking-go-client/v4/models/networking/v4/error"
)

/*
NC2A Subnet in the given VPC.
*/
type AwsSubnet struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Annotation string for cloud resources.
	*/
	Annotation *string `json:"annotation,omitempty"`
	/*
	  Cloud subnet mask.
	*/
	Cidr *string `json:"cidr"`

	CloudType *CloudType `json:"cloudType"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Cloud subnet gateway IP.
	*/
	GatewayIp *string `json:"gatewayIp"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Cloud subnet Id.
	*/
	SubnetId *string `json:"subnetId"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  AWS VPC ID where cluster is deployed.
	*/
	VpcId *string `json:"vpcId"`
}

func (p *AwsSubnet) MarshalJSON() ([]byte, error) {
	type AwsSubnetProxy AwsSubnet
	return json.Marshal(struct {
		*AwsSubnetProxy
		Cidr      *string    `json:"cidr,omitempty"`
		CloudType *CloudType `json:"cloudType,omitempty"`
		GatewayIp *string    `json:"gatewayIp,omitempty"`
		SubnetId  *string    `json:"subnetId,omitempty"`
		VpcId     *string    `json:"vpcId,omitempty"`
	}{
		AwsSubnetProxy: (*AwsSubnetProxy)(p),
		Cidr:           p.Cidr,
		CloudType:      p.CloudType,
		GatewayIp:      p.GatewayIp,
		SubnetId:       p.SubnetId,
		VpcId:          p.VpcId,
	})
}

func NewAwsSubnet() *AwsSubnet {
	p := new(AwsSubnet)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.aws.config.AwsSubnet"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
NC2 AWS VPC object.
*/
type AwsVpc struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Annotation string for cloud resources.
	*/
	Annotation *string `json:"annotation,omitempty"`
	/*
	  List of subnet cidrs associated with the AWS VPC.
	*/
	Cidrs []string `json:"cidrs"`

	CloudType *CloudType `json:"cloudType"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  AWS VPC ID where cluster is deployed.
	*/
	VpcId *string `json:"vpcId"`
}

func (p *AwsVpc) MarshalJSON() ([]byte, error) {
	type AwsVpcProxy AwsVpc
	return json.Marshal(struct {
		*AwsVpcProxy
		Cidrs     []string   `json:"cidrs,omitempty"`
		CloudType *CloudType `json:"cloudType,omitempty"`
		VpcId     *string    `json:"vpcId,omitempty"`
	}{
		AwsVpcProxy: (*AwsVpcProxy)(p),
		Cidrs:       p.Cidrs,
		CloudType:   p.CloudType,
		VpcId:       p.VpcId,
	})
}

func NewAwsVpc() *AwsVpc {
	p := new(AwsVpc)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.aws.config.AwsVpc"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Cloud capabilities schema.
*/
type Capabilities struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	Subnet *SubnetCapabilities `json:"subnet"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Capabilities) MarshalJSON() ([]byte, error) {
	type CapabilitiesProxy Capabilities
	return json.Marshal(struct {
		*CapabilitiesProxy
		Subnet *SubnetCapabilities `json:"subnet,omitempty"`
	}{
		CapabilitiesProxy: (*CapabilitiesProxy)(p),
		Subnet:            p.Subnet,
	})
}

func NewCapabilities() *Capabilities {
	p := new(Capabilities)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.aws.config.Capabilities"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Enum defining different cloud platforms.
*/
type CloudType int

const (
	CLOUDTYPE_UNKNOWN  CloudType = 0
	CLOUDTYPE_REDACTED CloudType = 1
	CLOUDTYPE_AWS      CloudType = 2
	CLOUDTYPE_AZURE    CloudType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *CloudType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AWS",
		"AZURE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e CloudType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AWS",
		"AZURE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *CloudType) index(name string) CloudType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AWS",
		"AZURE",
	}
	for idx := range names {
		if names[idx] == name {
			return CloudType(idx)
		}
	}
	return CLOUDTYPE_UNKNOWN
}

func (e *CloudType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CloudType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CloudType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CloudType) Ref() *CloudType {
	return &e
}

/*
REST response for all response codes in API path /networking/v4.0/aws/config/subnets Get operation
*/
type ListAwsSubnetsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListAwsSubnetsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListAwsSubnetsApiResponse() *ListAwsSubnetsApiResponse {
	p := new(ListAwsSubnetsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.aws.config.ListAwsSubnetsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListAwsSubnetsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListAwsSubnetsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListAwsSubnetsApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/aws/config/vpcs Get operation
*/
type ListAwsVpcsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListAwsVpcsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListAwsVpcsApiResponse() *ListAwsVpcsApiResponse {
	p := new(ListAwsVpcsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.aws.config.ListAwsVpcsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListAwsVpcsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListAwsVpcsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListAwsVpcsApiResponseData()
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
AwsSubnet capabilities schema.
*/
type SubnetCapabilities struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  This value defines the create, read, update and delete capabilities for the resource. It will be a 4-bit integer with the most significant bit defining delete capability, followed by update and create capability and finally least significant defining the read capability.
	*/
	Crud *int `json:"crud,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewSubnetCapabilities() *SubnetCapabilities {
	p := new(SubnetCapabilities)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.aws.config.SubnetCapabilities"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfListAwsVpcsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []AwsVpc               `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfListAwsVpcsApiResponseData() *OneOfListAwsVpcsApiResponseData {
	p := new(OneOfListAwsVpcsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListAwsVpcsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListAwsVpcsApiResponseData is nil"))
	}
	switch v.(type) {
	case []AwsVpc:
		p.oneOfType0 = v.([]AwsVpc)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.aws.config.AwsVpc>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.aws.config.AwsVpc>"
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfListAwsVpcsApiResponseData) GetValue() interface{} {
	if "List<networking.v4.aws.config.AwsVpc>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListAwsVpcsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]AwsVpc)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.aws.config.AwsVpc" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.aws.config.AwsVpc>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.aws.config.AwsVpc>"
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListAwsVpcsApiResponseData"))
}

func (p *OneOfListAwsVpcsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<networking.v4.aws.config.AwsVpc>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListAwsVpcsApiResponseData")
}

type OneOfListAwsSubnetsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    []AwsSubnet            `json:"-"`
}

func NewOneOfListAwsSubnetsApiResponseData() *OneOfListAwsSubnetsApiResponseData {
	p := new(OneOfListAwsSubnetsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListAwsSubnetsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListAwsSubnetsApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []AwsSubnet:
		p.oneOfType0 = v.([]AwsSubnet)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.aws.config.AwsSubnet>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.aws.config.AwsSubnet>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListAwsSubnetsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.aws.config.AwsSubnet>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListAwsSubnetsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType0 := new([]AwsSubnet)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.aws.config.AwsSubnet" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.aws.config.AwsSubnet>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.aws.config.AwsSubnet>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListAwsSubnetsApiResponseData"))
}

func (p *OneOfListAwsSubnetsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.aws.config.AwsSubnet>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListAwsSubnetsApiResponseData")
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
