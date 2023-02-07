/*
 * Generated file models/vmm/v4/ahv/config/config_model.go.
 *
 * Product version: 4.0.2-alpha-1
 *
 * Part of the Nutanix Vmm Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
 *
 */

/*
  Configure Acropolis hypervisor based Virtual Machines, Disks, Nics etc
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/orichter/package-publishing-example-go/vmm-go-client/v4/models/common/v1/config"
	import4 "github.com/orichter/package-publishing-example-go/vmm-go-client/v4/models/common/v1/response"
	import5 "github.com/orichter/package-publishing-example-go/vmm-go-client/v4/models/dataprotection/v4/config"
	import2 "github.com/orichter/package-publishing-example-go/vmm-go-client/v4/models/prism/v4/config"
	import3 "github.com/orichter/package-publishing-example-go/vmm-go-client/v4/models/vmm/v4/error"
	"time"
)

/**
ACPI shutdown mode options.
*/
type ACPIPowerOptions struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
}

func NewACPIPowerOptions() *ACPIPowerOptions {
	p := new(ACPIPowerOptions)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ACPIPowerOptions"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.ACPIPowerOptions"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
The disk migration plan defines how disks are migrated.
*/
type ADSFDiskMigrationPlan struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  The disks included in the migration plan.
	*/
	Disks []VmDiskReference `json:"disks,omitempty"`

	StorageContainer *VmDiskContainerReference `json:"storageContainer,omitempty"`
}

func NewADSFDiskMigrationPlan() *ADSFDiskMigrationPlan {
	p := new(ADSFDiskMigrationPlan)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ADSFDiskMigrationPlan"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.ADSFDiskMigrationPlan"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Storage configuration for VM.
*/
type ADSFVmStorageConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Indicates whether the virtual disk is pinned to the hot tier or not.
	*/
	FlashModeEnabled *bool `json:"flashModeEnabled,omitempty"`

	QosConfig *QosConfig `json:"qosConfig,omitempty"`
}

func NewADSFVmStorageConfig() *ADSFVmStorageConfig {
	p := new(ADSFVmStorageConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ADSFVmStorageConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.ADSFVmStorageConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Reference to an ADSF Volume Group.
*/
type ADSFVolumeGroupReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Globally unique identifier of an ADSF volume group. It should be of type UUID.
	*/
	VolumeGroupExtId *string `json:"volumeGroupExtId,omitempty"`
}

func NewADSFVolumeGroupReference() *ADSFVolumeGroupReference {
	p := new(ADSFVolumeGroupReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ADSFVolumeGroupReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.ADSFVolumeGroupReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
IP address assignment to a VM NIC.
*/
type AssignIp struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	IpAddress *import1.IPv4Address `json:"ipAddress,omitempty"`
}

func NewAssignIp() *AssignIp {
	p := new(AssignIp)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.AssignIp"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.AssignIp"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/nics/{extId}/$actions/assign-ip Post operation
*/
type AssignIpResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAssignIpResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAssignIpResponse() *AssignIpResponse {
	p := new(AssignIpResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.AssignIpResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.AssignIpResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AssignIpResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AssignIpResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAssignIpResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/categories/$actions/attach Post operation
*/
type AttachVmCategoryResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAttachVmCategoryResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAttachVmCategoryResponse() *AttachVmCategoryResponse {
	p := new(AttachVmCategoryResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.AttachVmCategoryResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.AttachVmCategoryResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AttachVmCategoryResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AttachVmCategoryResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAttachVmCategoryResponseData()
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

/**
Reference to an availability zone.
*/
type AvailabilityZoneReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Globally unique identifier of an availability zone. It should be of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func NewAvailabilityZoneReference() *AvailabilityZoneReference {
	p := new(AvailabilityZoneReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.AvailabilityZoneReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.AvailabilityZoneReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
The disk that will be used to boot the VM.
*/
type BootDeviceDisk struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DiskAddress *DiskAddress `json:"diskAddress,omitempty"`
}

func NewBootDeviceDisk() *BootDeviceDisk {
	p := new(BootDeviceDisk)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.BootDeviceDisk"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.BootDeviceDisk"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
The NIC that will be used to boot the VM.
*/
type BootDeviceNic struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	MacAddress *string `json:"macAddress,omitempty"`
}

func NewBootDeviceNic() *BootDeviceNic {
	p := new(BootDeviceNic)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.BootDeviceNic"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.BootDeviceNic"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
The type of boot device. Valid values are CDROM, NIC, and DISK.
*/
type BootDeviceType int

const (
	BOOTDEVICETYPE_UNKNOWN  BootDeviceType = 0
	BOOTDEVICETYPE_REDACTED BootDeviceType = 1
	BOOTDEVICETYPE_CDROM    BootDeviceType = 2
	BOOTDEVICETYPE_DISK     BootDeviceType = 3
	BOOTDEVICETYPE_NETWORK  BootDeviceType = 4
)

// returns the name of the enum given an ordinal number
func (e *BootDeviceType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CDROM",
		"DISK",
		"NETWORK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *BootDeviceType) index(name string) BootDeviceType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CDROM",
		"DISK",
		"NETWORK",
	}
	for idx := range names {
		if names[idx] == name {
			return BootDeviceType(idx)
		}
	}
	return BOOTDEVICETYPE_UNKNOWN
}

func (e *BootDeviceType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for BootDeviceType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *BootDeviceType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e BootDeviceType) Ref() *BootDeviceType {
	return &e
}

/**
Reference to a category.
*/
type CategoryReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Globally unique identifier of a VM category. It should be of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func NewCategoryReference() *CategoryReference {
	p := new(CategoryReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CategoryReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.CategoryReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Virtual Machine CD-ROM.
*/
type Cdrom struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BackingInfo *VmDisk `json:"backingInfo,omitempty"`

	DiskAddress *CdromAddress `json:"diskAddress,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewCdrom() *Cdrom {
	p := new(Cdrom)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Cdrom"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.Cdrom"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Virtual Machine disk (VM disk).
*/
type CdromAddress struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BusType *CdromBusType `json:"busType,omitempty"`
	/**
	  Device index on the bus. This field is ignored unless the bus details are specified.
	*/
	Index *int `json:"index,omitempty"`
}

func NewCdromAddress() *CdromAddress {
	p := new(CdromAddress)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CdromAddress"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.CdromAddress"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Bus type for the device. The acceptable values are: IDE, SATA.
*/
type CdromBusType int

const (
	CDROMBUSTYPE_UNKNOWN  CdromBusType = 0
	CDROMBUSTYPE_REDACTED CdromBusType = 1
	CDROMBUSTYPE_IDE      CdromBusType = 2
	CDROMBUSTYPE_SATA     CdromBusType = 3
)

// returns the name of the enum given an ordinal number
func (e *CdromBusType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IDE",
		"SATA",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *CdromBusType) index(name string) CdromBusType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IDE",
		"SATA",
	}
	for idx := range names {
		if names[idx] == name {
			return CdromBusType(idx)
		}
	}
	return CDROMBUSTYPE_UNKNOWN
}

func (e *CdromBusType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CdromBusType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CdromBusType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CdromBusType) Ref() *CdromBusType {
	return &e
}

/**
Virtual Machine CD-ROM.
*/
type CdromInsert struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BackingInfo *VmDisk `json:"backingInfo,omitempty"`
}

func NewCdromInsert() *CdromInsert {
	p := new(CdromInsert)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CdromInsert"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.CdromInsert"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
VM clone override specification.
*/
type CloneOverride struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	BootConfigItemDiscriminator_ *string `json:"$bootConfigItemDiscriminator,omitempty"`
	/**
	  Indicates the order of device types in which the VM should try to boot from. If the boot device order is not provided the system will decide an appropriate boot device order.
	*/
	BootConfig *OneOfCloneOverrideBootConfig `json:"bootConfig,omitempty"`

	GuestCustomization *GuestCustomization `json:"guestCustomization,omitempty"`
	/**
	  Memory size in bytes.
	*/
	MemorySizeBytes *int64 `json:"memorySizeBytes,omitempty"`
	/**
	  VM name.
	*/
	Name *string `json:"name,omitempty"`
	/**
	  NICs attached to the VM.
	*/
	Nics []Nic `json:"nics,omitempty"`
	/**
	  Number of cores per socket.
	*/
	NumCoresPerSocket *int `json:"numCoresPerSocket,omitempty"`
	/**
	  Number of vCPU sockets.
	*/
	NumSockets *int `json:"numSockets,omitempty"`
	/**
	  Number of cores per socket.
	*/
	NumThreadsPerCore *int `json:"numThreadsPerCore,omitempty"`
}

func NewCloneOverride() *CloneOverride {
	p := new(CloneOverride)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CloneOverride"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.CloneOverride"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CloneOverride) GetBootConfig() interface{} {
	if nil == p.BootConfig {
		return nil
	}
	return p.BootConfig.GetValue()
}

func (p *CloneOverride) SetBootConfig(v interface{}) error {
	if nil == p.BootConfig {
		p.BootConfig = NewOneOfCloneOverrideBootConfig()
	}
	e := p.BootConfig.SetValue(v)
	if nil == e {
		if nil == p.BootConfigItemDiscriminator_ {
			p.BootConfigItemDiscriminator_ = new(string)
		}
		*p.BootConfigItemDiscriminator_ = *p.BootConfig.Discriminator
	}
	return e
}

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/clone Post operation
*/
type CloneVmResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCloneVmResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCloneVmResponse() *CloneVmResponse {
	p := new(CloneVmResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CloneVmResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.CloneVmResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CloneVmResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CloneVmResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCloneVmResponseData()
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

/**
If this field is set, the guest will be customized using cloud-init. Either user_data or custom_key_values should be provided. If custom_key_ves are provided then the user data will be generated using these key-value pairs.
*/
type CloudInit struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	CloudInitScriptItemDiscriminator_ *string `json:"$cloudInitScriptItemDiscriminator,omitempty"`
	/**
	  The script to use for cloud-init.
	*/
	CloudInitScript *OneOfCloudInitCloudInitScript `json:"cloudInitScript,omitempty"`

	DatasourceType *CloudInitDataSourceType `json:"datasourceType,omitempty"`
	/**
	  The contents of the meta_data configuration for cloud-init. This can be formatted as YAML or JSON. The value must be base64 encoded.
	*/
	Metadata *string `json:"metadata,omitempty"`
}

func NewCloudInit() *CloudInit {
	p := new(CloudInit)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CloudInit"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.CloudInit"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CloudInit) GetCloudInitScript() interface{} {
	if nil == p.CloudInitScript {
		return nil
	}
	return p.CloudInitScript.GetValue()
}

func (p *CloudInit) SetCloudInitScript(v interface{}) error {
	if nil == p.CloudInitScript {
		p.CloudInitScript = NewOneOfCloudInitCloudInitScript()
	}
	e := p.CloudInitScript.SetValue(v)
	if nil == e {
		if nil == p.CloudInitScriptItemDiscriminator_ {
			p.CloudInitScriptItemDiscriminator_ = new(string)
		}
		*p.CloudInitScriptItemDiscriminator_ = *p.CloudInitScript.Discriminator
	}
	return e
}

/**
Type of datasource. Default: CONFIG_DRIVE_V2
*/
type CloudInitDataSourceType int

const (
	CLOUDINITDATASOURCETYPE_UNKNOWN         CloudInitDataSourceType = 0
	CLOUDINITDATASOURCETYPE_REDACTED        CloudInitDataSourceType = 1
	CLOUDINITDATASOURCETYPE_CONFIG_DRIVE_V2 CloudInitDataSourceType = 2
)

// returns the name of the enum given an ordinal number
func (e *CloudInitDataSourceType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONFIG_DRIVE_V2",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *CloudInitDataSourceType) index(name string) CloudInitDataSourceType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONFIG_DRIVE_V2",
	}
	for idx := range names {
		if names[idx] == name {
			return CloudInitDataSourceType(idx)
		}
	}
	return CLOUDINITDATASOURCETYPE_UNKNOWN
}

func (e *CloudInitDataSourceType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CloudInitDataSourceType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CloudInitDataSourceType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CloudInitDataSourceType) Ref() *CloudInitDataSourceType {
	return &e
}

/**
Reference to a cluster.
*/
type ClusterReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Globally unique identifier of a cluster. It should be of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func NewClusterReference() *ClusterReference {
	p := new(ClusterReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ClusterReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.ClusterReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
The list of additional CPU features to be enabled. HardwareVirtualization: Indicates whether hardware assisted virtualization should be enabled for the Guest OS or not. Once enabled, the Guest OS can deploy a nested hypervisor.
*/
type CpuFeature int

const (
	CPUFEATURE_UNKNOWN                 CpuFeature = 0
	CPUFEATURE_REDACTED                CpuFeature = 1
	CPUFEATURE_HARDWARE_VIRTUALIZATION CpuFeature = 2
)

// returns the name of the enum given an ordinal number
func (e *CpuFeature) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HARDWARE_VIRTUALIZATION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *CpuFeature) index(name string) CpuFeature {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HARDWARE_VIRTUALIZATION",
	}
	for idx := range names {
		if names[idx] == name {
			return CpuFeature(idx)
		}
	}
	return CPUFEATURE_UNKNOWN
}

func (e *CpuFeature) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CpuFeature:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CpuFeature) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CpuFeature) Ref() *CpuFeature {
	return &e
}

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/cdroms Post operation
*/
type CreateCdromResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateCdromResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateCdromResponse() *CreateCdromResponse {
	p := new(CreateCdromResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CreateCdromResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.CreateCdromResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateCdromResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateCdromResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateCdromResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/disks Post operation
*/
type CreateDiskResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateDiskResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateDiskResponse() *CreateDiskResponse {
	p := new(CreateDiskResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CreateDiskResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.CreateDiskResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateDiskResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateDiskResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateDiskResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/gpus Post operation
*/
type CreateGpuResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateGpuResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateGpuResponse() *CreateGpuResponse {
	p := new(CreateGpuResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CreateGpuResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.CreateGpuResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateGpuResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateGpuResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateGpuResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/nics Post operation
*/
type CreateNicResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateNicResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateNicResponse() *CreateNicResponse {
	p := new(CreateNicResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CreateNicResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.CreateNicResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateNicResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateNicResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateNicResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/serial-ports Post operation
*/
type CreateSerialPortResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateSerialPortResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateSerialPortResponse() *CreateSerialPortResponse {
	p := new(CreateSerialPortResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CreateSerialPortResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.CreateSerialPortResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateSerialPortResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateSerialPortResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateSerialPortResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms Post operation
*/
type CreateVmResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateVmResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateVmResponse() *CreateVmResponse {
	p := new(CreateVmResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CreateVmResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.CreateVmResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateVmResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateVmResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateVmResponseData()
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

/**
Sign in credentials for the server.
*/
type Credential struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Password *string `json:"password,omitempty"`

	Username *string `json:"username,omitempty"`
}

func NewCredential() *Credential {
	p := new(Credential)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Credential"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.Credential"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
A collection of key/value pairs.
*/
type CustomKeyValues struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  The list of the individual KeyValuePair elements.
	*/
	KeyValuePairs []import1.KVPair `json:"keyValuePairs,omitempty"`
}

func NewCustomKeyValues() *CustomKeyValues {
	p := new(CustomKeyValues)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CustomKeyValues"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.CustomKeyValues"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/customize-guest Post operation
*/
type CustomizeGuestVmResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCustomizeGuestVmResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCustomizeGuestVmResponse() *CustomizeGuestVmResponse {
	p := new(CustomizeGuestVmResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CustomizeGuestVmResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.CustomizeGuestVmResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CustomizeGuestVmResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CustomizeGuestVmResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCustomizeGuestVmResponseData()
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

/**
A reference to a disk or image that contains the contents of a disk.
*/
type DataSource struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	ReferenceItemDiscriminator_ *string `json:"$referenceItemDiscriminator,omitempty"`

	Reference *OneOfDataSourceReference `json:"reference,omitempty"`
}

func NewDataSource() *DataSource {
	p := new(DataSource)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DataSource"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.DataSource"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DataSource) GetReference() interface{} {
	if nil == p.Reference {
		return nil
	}
	return p.Reference.GetValue()
}

func (p *DataSource) SetReference(v interface{}) error {
	if nil == p.Reference {
		p.Reference = NewOneOfDataSourceReference()
	}
	e := p.Reference.SetValue(v)
	if nil == e {
		if nil == p.ReferenceItemDiscriminator_ {
			p.ReferenceItemDiscriminator_ = new(string)
		}
		*p.ReferenceItemDiscriminator_ = *p.Reference.Discriminator
	}
	return e
}

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/cdroms/{extId} Delete operation
*/
type DeleteCdromResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteCdromResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteCdromResponse() *DeleteCdromResponse {
	p := new(DeleteCdromResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DeleteCdromResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.DeleteCdromResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteCdromResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteCdromResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteCdromResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/disks/{extId} Delete operation
*/
type DeleteDiskResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteDiskResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteDiskResponse() *DeleteDiskResponse {
	p := new(DeleteDiskResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DeleteDiskResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.DeleteDiskResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteDiskResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteDiskResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteDiskResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/gpus/{extId} Delete operation
*/
type DeleteGpuResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteGpuResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteGpuResponse() *DeleteGpuResponse {
	p := new(DeleteGpuResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DeleteGpuResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.DeleteGpuResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteGpuResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteGpuResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteGpuResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/nics/{extId} Delete operation
*/
type DeleteNicResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteNicResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteNicResponse() *DeleteNicResponse {
	p := new(DeleteNicResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DeleteNicResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.DeleteNicResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteNicResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteNicResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteNicResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/serial-ports/{extId} Delete operation
*/
type DeleteSerialPortResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteSerialPortResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteSerialPortResponse() *DeleteSerialPortResponse {
	p := new(DeleteSerialPortResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DeleteSerialPortResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.DeleteSerialPortResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteSerialPortResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteSerialPortResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteSerialPortResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId} Delete operation
*/
type DeleteVmResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteVmResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteVmResponse() *DeleteVmResponse {
	p := new(DeleteVmResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DeleteVmResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.DeleteVmResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteVmResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteVmResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteVmResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/categories/$actions/detach Post operation
*/
type DetachVmCategoryResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDetachVmCategoryResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDetachVmCategoryResponse() *DetachVmCategoryResponse {
	p := new(DetachVmCategoryResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DetachVmCategoryResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.DetachVmCategoryResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DetachVmCategoryResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DetachVmCategoryResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDetachVmCategoryResponseData()
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

/**
Virtual Machine disk (VM disk).
*/
type Disk struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	BackingInfoItemDiscriminator_ *string `json:"$backingInfoItemDiscriminator,omitempty"`
	/**
	  Supporting storage to create virtual disk on.
	*/
	BackingInfo *OneOfDiskBackingInfo `json:"backingInfo,omitempty"`

	DiskAddress *DiskAddress `json:"diskAddress,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/**
	  Indicates whether the SCSI passthrough (true) is used or not (false).
	*/
	ScsiPassthroughEnabled *bool `json:"scsiPassthroughEnabled,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewDisk() *Disk {
	p := new(Disk)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Disk"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.Disk"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ScsiPassthroughEnabled = new(bool)
	*p.ScsiPassthroughEnabled = true

	return p
}

/**
Disk address.
*/
type DiskAddress struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BusType *DiskBusType `json:"busType,omitempty"`
	/**
	  Device index on the bus. This field is ignored unless the bus details are specified.
	*/
	Index *int `json:"index,omitempty"`
}

func NewDiskAddress() *DiskAddress {
	p := new(DiskAddress)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DiskAddress"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.DiskAddress"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Bus type for the device. The acceptable values are: SCSI, IDE, PCI, SATA, SPAPR (only PPC).
*/
type DiskBusType int

const (
	DISKBUSTYPE_UNKNOWN  DiskBusType = 0
	DISKBUSTYPE_REDACTED DiskBusType = 1
	DISKBUSTYPE_SCSI     DiskBusType = 2
	DISKBUSTYPE_IDE      DiskBusType = 3
	DISKBUSTYPE_PCI      DiskBusType = 4
	DISKBUSTYPE_SATA     DiskBusType = 5
	DISKBUSTYPE_SPAPR    DiskBusType = 6
)

// returns the name of the enum given an ordinal number
func (e *DiskBusType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SCSI",
		"IDE",
		"PCI",
		"SATA",
		"SPAPR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *DiskBusType) index(name string) DiskBusType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SCSI",
		"IDE",
		"PCI",
		"SATA",
		"SPAPR",
	}
	for idx := range names {
		if names[idx] == name {
			return DiskBusType(idx)
		}
	}
	return DISKBUSTYPE_UNKNOWN
}

func (e *DiskBusType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for DiskBusType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *DiskBusType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e DiskBusType) Ref() *DiskBusType {
	return &e
}

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/cdroms/{extId}/$actions/eject Post operation
*/
type EjectCdromResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfEjectCdromResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewEjectCdromResponse() *EjectCdromResponse {
	p := new(EjectCdromResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.EjectCdromResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.EjectCdromResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *EjectCdromResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *EjectCdromResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfEjectCdromResponseData()
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

/**
Defines a NIC emulated by the hypervisor
*/
type EmulatedNic struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Indicates whether the NIC is connected or not. Default is True.
	*/
	IsConnected *bool `json:"isConnected,omitempty"`
	/**
	  MAC address of the emulated NIC.
	*/
	MacAddress *string `json:"macAddress,omitempty"`

	Model *EmulatedNicModel `json:"model,omitempty"`
	/**
	  The number of Tx/Rx queue pairs for this NIC.
	*/
	NumQueues *int `json:"numQueues,omitempty"`
}

func NewEmulatedNic() *EmulatedNic {
	p := new(EmulatedNic)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.EmulatedNic"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.EmulatedNic"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsConnected = new(bool)
	*p.IsConnected = true
	p.NumQueues = new(int)
	*p.NumQueues = 1

	return p
}

/**
Options for the NIC emulation.
*/
type EmulatedNicModel int

const (
	EMULATEDNICMODEL_UNKNOWN  EmulatedNicModel = 0
	EMULATEDNICMODEL_REDACTED EmulatedNicModel = 1
	EMULATEDNICMODEL_VIRTIO   EmulatedNicModel = 2
	EMULATEDNICMODEL_E1000    EmulatedNicModel = 3
)

// returns the name of the enum given an ordinal number
func (e *EmulatedNicModel) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VIRTIO",
		"E1000",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *EmulatedNicModel) index(name string) EmulatedNicModel {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VIRTIO",
		"E1000",
	}
	for idx := range names {
		if names[idx] == name {
			return EmulatedNicModel(idx)
		}
	}
	return EMULATEDNICMODEL_UNKNOWN
}

func (e *EmulatedNicModel) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EmulatedNicModel:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EmulatedNicModel) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EmulatedNicModel) Ref() *EmulatedNicModel {
	return &e
}

/**
Policy enforcement mode informs us about what the policy engine is currently doing to enforce the policy on the entity. Monitoring indicates that the policy engine is simply monitoring the entity state. Enforcing means that the policy engine is currently trying to enforce the policy on the entity. Enforcement failed indicates that the policy engine encountered a non-transient error and requires user intervention to fix the problem, the error message gives the reason for the error in this case.
*/
type EnforcementMode int

const (
	ENFORCEMENTMODE_UNKNOWN            EnforcementMode = 0
	ENFORCEMENTMODE_REDACTED           EnforcementMode = 1
	ENFORCEMENTMODE_MONITORING         EnforcementMode = 2
	ENFORCEMENTMODE_ENFORCING          EnforcementMode = 3
	ENFORCEMENTMODE_ENFORCEMENT_FAILED EnforcementMode = 4
)

// returns the name of the enum given an ordinal number
func (e *EnforcementMode) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MONITORING",
		"ENFORCING",
		"ENFORCEMENT_FAILED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *EnforcementMode) index(name string) EnforcementMode {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MONITORING",
		"ENFORCING",
		"ENFORCEMENT_FAILED",
	}
	for idx := range names {
		if names[idx] == name {
			return EnforcementMode(idx)
		}
	}
	return ENFORCEMENTMODE_UNKNOWN
}

func (e *EnforcementMode) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EnforcementMode:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EnforcementMode) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EnforcementMode) Ref() *EnforcementMode {
	return &e
}

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/cdroms/{extId} Get operation
*/
type GetCdromResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetCdromResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetCdromResponse() *GetCdromResponse {
	p := new(GetCdromResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GetCdromResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.GetCdromResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetCdromResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetCdromResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetCdromResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/disks/{extId} Get operation
*/
type GetDiskResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetDiskResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetDiskResponse() *GetDiskResponse {
	p := new(GetDiskResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GetDiskResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.GetDiskResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetDiskResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetDiskResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetDiskResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/gpus/{extId} Get operation
*/
type GetGpuResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetGpuResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetGpuResponse() *GetGpuResponse {
	p := new(GetGpuResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GetGpuResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.GetGpuResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetGpuResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetGpuResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetGpuResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/guest-tools-config Get operation
*/
type GetGuestToolsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetGuestToolsResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetGuestToolsResponse() *GetGuestToolsResponse {
	p := new(GetGuestToolsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GetGuestToolsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.GetGuestToolsResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetGuestToolsResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetGuestToolsResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetGuestToolsResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/nics/{extId} Get operation
*/
type GetNicResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetNicResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetNicResponse() *GetNicResponse {
	p := new(GetNicResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GetNicResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.GetNicResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetNicResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetNicResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetNicResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/serial-ports/{extId} Get operation
*/
type GetSerialPortResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetSerialPortResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetSerialPortResponse() *GetSerialPortResponse {
	p := new(GetSerialPortResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GetSerialPortResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.GetSerialPortResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetSerialPortResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetSerialPortResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetSerialPortResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/categories Get operation
*/
type GetVmCategoriesResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVmCategoriesResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetVmCategoriesResponse() *GetVmCategoriesResponse {
	p := new(GetVmCategoriesResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GetVmCategoriesResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.GetVmCategoriesResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVmCategoriesResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVmCategoriesResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVmCategoriesResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/ownership-info Get operation
*/
type GetVmOwnershipInfoResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVmOwnershipInfoResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetVmOwnershipInfoResponse() *GetVmOwnershipInfoResponse {
	p := new(GetVmOwnershipInfoResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GetVmOwnershipInfoResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.GetVmOwnershipInfoResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVmOwnershipInfoResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVmOwnershipInfoResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVmOwnershipInfoResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId} Get operation
*/
type GetVmResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVmResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetVmResponse() *GetVmResponse {
	p := new(GetVmResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GetVmResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.GetVmResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVmResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVmResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVmResponseData()
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

/**
Graphics resource information for the Virtual Machine.
*/
type Gpu struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  The device Id of the GPU.
	*/
	DeviceId *int `json:"deviceId,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  Fraction of the physical GPU assigned.
	*/
	Fraction *int `json:"fraction,omitempty"`
	/**
	  GPU frame buffer size in bytes.
	*/
	FrameBufferSizeBytes *int64 `json:"frameBufferSizeBytes,omitempty"`
	/**
	  Last determined guest driver version.
	*/
	GuestDriverVersion *string `json:"guestDriverVersion,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`

	Mode *GpuMode `json:"mode,omitempty"`
	/**
	  Name of the GPU resource.
	*/
	Name *string `json:"name,omitempty"`
	/**
	  Number of supported virtual display heads.
	*/
	NumVirtualDisplayHeads *int `json:"numVirtualDisplayHeads,omitempty"`

	PciAddress *SBDF `json:"pciAddress,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Vendor *GpuVendor `json:"vendor,omitempty"`
}

func NewGpu() *Gpu {
	p := new(Gpu)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Gpu"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.Gpu"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
The mode of this GPU.
*/
type GpuMode int

const (
	GPUMODE_UNKNOWN              GpuMode = 0
	GPUMODE_REDACTED             GpuMode = 1
	GPUMODE_PASSTHROUGH_GRAPHICS GpuMode = 2
	GPUMODE_PASSTHROUGH_COMPUTE  GpuMode = 3
	GPUMODE_VIRTUAL              GpuMode = 4
)

// returns the name of the enum given an ordinal number
func (e *GpuMode) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PASSTHROUGH_GRAPHICS",
		"PASSTHROUGH_COMPUTE",
		"VIRTUAL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *GpuMode) index(name string) GpuMode {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PASSTHROUGH_GRAPHICS",
		"PASSTHROUGH_COMPUTE",
		"VIRTUAL",
	}
	for idx := range names {
		if names[idx] == name {
			return GpuMode(idx)
		}
	}
	return GPUMODE_UNKNOWN
}

func (e *GpuMode) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for GpuMode:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *GpuMode) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e GpuMode) Ref() *GpuMode {
	return &e
}

/**
The vendor of the GPU.
*/
type GpuVendor int

const (
	GPUVENDOR_UNKNOWN  GpuVendor = 0
	GPUVENDOR_REDACTED GpuVendor = 1
	GPUVENDOR_NVIDIA   GpuVendor = 2
	GPUVENDOR_INTEL    GpuVendor = 3
	GPUVENDOR_AMD      GpuVendor = 4
)

// returns the name of the enum given an ordinal number
func (e *GpuVendor) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NVIDIA",
		"INTEL",
		"AMD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *GpuVendor) index(name string) GpuVendor {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NVIDIA",
		"INTEL",
		"AMD",
	}
	for idx := range names {
		if names[idx] == name {
			return GpuVendor(idx)
		}
	}
	return GPUVENDOR_UNKNOWN
}

func (e *GpuVendor) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for GpuVendor:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *GpuVendor) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e GpuVendor) Ref() *GpuVendor {
	return &e
}

/**
Customize a guest VM.
*/
type GuestCustomization struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	ConfigItemDiscriminator_ *string `json:"$configItemDiscriminator,omitempty"`
	/**
	  The Nutanix Guest Tools customization settings.
	*/
	Config *OneOfGuestCustomizationConfig `json:"config,omitempty"`
}

func NewGuestCustomization() *GuestCustomization {
	p := new(GuestCustomization)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GuestCustomization"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.GuestCustomization"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GuestCustomization) GetConfig() interface{} {
	if nil == p.Config {
		return nil
	}
	return p.Config.GetValue()
}

func (p *GuestCustomization) SetConfig(v interface{}) error {
	if nil == p.Config {
		p.Config = NewOneOfGuestCustomizationConfig()
	}
	e := p.Config.SetValue(v)
	if nil == e {
		if nil == p.ConfigItemDiscriminator_ {
			p.ConfigItemDiscriminator_ = new(string)
		}
		*p.ConfigItemDiscriminator_ = *p.Config.Discriminator
	}
	return e
}

/**
Nutanix Guest Tools shutdown mode options.
*/
type GuestPowerOptions struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	GuestPowerStateTransitionConfig *GuestPowerStateTransitionConfig `json:"guestPowerStateTransitionConfig,omitempty"`
}

func NewGuestPowerOptions() *GuestPowerOptions {
	p := new(GuestPowerOptions)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GuestPowerOptions"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.GuestPowerOptions"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Additional configuration for Nutanix Gust Tools power state transition.
*/
type GuestPowerStateTransitionConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Indicates whether to run the set script before the VM shutdowns/restarts.
	*/
	EnableScriptExec *bool `json:"enableScriptExec,omitempty"`
	/**
	  Indicates whether to abort VM shutdown/restart if the script fails.
	*/
	ShouldFailOnScriptFailure *bool `json:"shouldFailOnScriptFailure,omitempty"`
}

func NewGuestPowerStateTransitionConfig() *GuestPowerStateTransitionConfig {
	p := new(GuestPowerStateTransitionConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GuestPowerStateTransitionConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.GuestPowerStateTransitionConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
The details about Nutanix Guest Tools for a VM.
*/
type GuestTools struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Version of Nutanix Guest Tools available on the cluster.
	*/
	AvailableVersion *string `json:"availableVersion,omitempty"`
	/**
	  The list of the application names that are enabled on the guest VM.
	*/
	Capabilities []NgtCapability `json:"capabilities,omitempty"`
	/**
	  Version of the operating system on the VM.
	*/
	GuestOsVersion *string `json:"guestOsVersion,omitempty"`
	/**
	  Indicates whether Nutanix Guest Tools is enabled or not.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`
	/**
	  Indicates whether Nutanix Guest Tools is installed on the VM or not.
	*/
	IsInstalled *bool `json:"isInstalled,omitempty"`
	/**
	  Indicates whether Nutanix Guest Tools ISO is inserted or not.
	*/
	IsIsoInserted *bool `json:"isIsoInserted,omitempty"`
	/**
	  Indicates whether the communication from VM to CVM is active or not.
	*/
	IsReachable *bool `json:"isReachable,omitempty"`
	/**
	  Indicates whether the VM mobility drivers are installed on the VM or not.
	*/
	IsVmMobilityDriversInstalled *bool `json:"isVmMobilityDriversInstalled,omitempty"`
	/**
	  Indicates whether the VM is configured to take VSS snapshots through NGT or not.
	*/
	IsVssSnapshotCapable *bool `json:"isVssSnapshotCapable,omitempty"`
	/**
	  Version of Nutanix Guest Tools installed on the VM.
	*/
	Version *string `json:"version,omitempty"`
}

func NewGuestTools() *GuestTools {
	p := new(GuestTools)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GuestTools"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.GuestTools"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Argument for inserting a Nutanix Guest Tools ISO into an available slot.
*/
type GuestToolsInsertConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  The list of the application names that are enabled on the guest VM.
	*/
	Capabilities []NgtCapability `json:"capabilities,omitempty"`
}

func NewGuestToolsInsertConfig() *GuestToolsInsertConfig {
	p := new(GuestToolsInsertConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GuestToolsInsertConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.GuestToolsInsertConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Arguments for installing Nutanix Guest Tools.
*/
type GuestToolsInstallConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  The list of the application names that are enabled on the guest VM.
	*/
	Capabilities []NgtCapability `json:"capabilities,omitempty"`

	Credential *Credential `json:"credential,omitempty"`

	RebootPreference *RebootPreference `json:"rebootPreference,omitempty"`
}

func NewGuestToolsInstallConfig() *GuestToolsInstallConfig {
	p := new(GuestToolsInstallConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GuestToolsInstallConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.GuestToolsInstallConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Arguments for upgrading Nutanix Guest Tools.
*/
type GuestToolsUpgradeConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	RebootPreference *RebootPreference `json:"rebootPreference,omitempty"`
}

func NewGuestToolsUpgradeConfig() *GuestToolsUpgradeConfig {
	p := new(GuestToolsUpgradeConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GuestToolsUpgradeConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.GuestToolsUpgradeConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Reference to the host, the VM is running on.
*/
type HostReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Globally unique identifier of a host. It should be of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func NewHostReference() *HostReference {
	p := new(HostReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.HostReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.HostReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Clone a disk from an image.
*/
type ImageReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Globally unique identifier of an image. It should be of type UUID.
	*/
	ImageExtId *string `json:"imageExtId,omitempty"`
}

func NewImageReference() *ImageReference {
	p := new(ImageReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ImageReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.ImageReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/cdroms/{extId}/$actions/insert Post operation
*/
type InsertCdromResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfInsertCdromResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewInsertCdromResponse() *InsertCdromResponse {
	p := new(InsertCdromResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.InsertCdromResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.InsertCdromResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *InsertCdromResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *InsertCdromResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfInsertCdromResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/guest-tools/$actions/insert-iso Post operation
*/
type InsertVmGuestToolsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfInsertVmGuestToolsResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewInsertVmGuestToolsResponse() *InsertVmGuestToolsResponse {
	p := new(InsertVmGuestToolsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.InsertVmGuestToolsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.InsertVmGuestToolsResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *InsertVmGuestToolsResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *InsertVmGuestToolsResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfInsertVmGuestToolsResponseData()
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

/**
Indicates whether the guest will be freshly installed using this unattend configuration, or this unattend configuration will be applied to a pre-prepared image. Default is 'PREPARED'.
*/
type InstallType int

const (
	INSTALLTYPE_UNKNOWN  InstallType = 0
	INSTALLTYPE_REDACTED InstallType = 1
	INSTALLTYPE_FRESH    InstallType = 2
	INSTALLTYPE_PREPARED InstallType = 3
)

// returns the name of the enum given an ordinal number
func (e *InstallType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FRESH",
		"PREPARED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *InstallType) index(name string) InstallType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FRESH",
		"PREPARED",
	}
	for idx := range names {
		if names[idx] == name {
			return InstallType(idx)
		}
	}
	return INSTALLTYPE_UNKNOWN
}

func (e *InstallType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for InstallType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *InstallType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e InstallType) Ref() *InstallType {
	return &e
}

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/guest-tools/$actions/install Post operation
*/
type InstallVmGuestToolsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfInstallVmGuestToolsResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewInstallVmGuestToolsResponse() *InstallVmGuestToolsResponse {
	p := new(InstallVmGuestToolsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.InstallVmGuestToolsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.InstallVmGuestToolsResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *InstallVmGuestToolsResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *InstallVmGuestToolsResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfInstallVmGuestToolsResponseData()
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

/**
The IP address configurations.
*/
type Ipv4Config struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	IpAddress *import1.IPv4Address `json:"ipAddress,omitempty"`
	/**
	  Secondary IP addresses for the NIC.
	*/
	SecondaryIpAddressList []import1.IPv4Address `json:"secondaryIpAddressList,omitempty"`
	/**
	  If set to true (default value), an IP address must be assigned to the VM NIC - either the one explicitly specified by the user or allocated automatically by the IPAM service by not specifying the IP address. If false, then no IP assignment is required for this VM NIC.
	*/
	ShouldAssignIp *bool `json:"shouldAssignIp,omitempty"`
}

func NewIpv4Config() *Ipv4Config {
	p := new(Ipv4Config)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Ipv4Config"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.Ipv4Config"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
The runtime IP address information of the NIC.
*/
type Ipv4Info struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  The IP addresses as found on the guest VM for the NIC.
	*/
	LearnedIpAddresses []import1.IPv4Address `json:"learnedIpAddresses,omitempty"`
}

func NewIpv4Info() *Ipv4Info {
	p := new(Ipv4Info)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Ipv4Info"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.Ipv4Info"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Legacy boot mode and its associated configuration.
*/
type LegacyBoot struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	BootDeviceItemDiscriminator_ *string `json:"$bootDeviceItemDiscriminator,omitempty"`

	BootDevice *OneOfLegacyBootBootDevice `json:"bootDevice,omitempty"`
	/**
	  Indicates the order of device types in which the VM should try to boot from. If the boot device order is not provided the system will decide an appropriate boot device order.
	*/
	BootOrder []BootDeviceType `json:"bootOrder,omitempty"`
}

func NewLegacyBoot() *LegacyBoot {
	p := new(LegacyBoot)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.LegacyBoot"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.LegacyBoot"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *LegacyBoot) GetBootDevice() interface{} {
	if nil == p.BootDevice {
		return nil
	}
	return p.BootDevice.GetValue()
}

func (p *LegacyBoot) SetBootDevice(v interface{}) error {
	if nil == p.BootDevice {
		p.BootDevice = NewOneOfLegacyBootBootDevice()
	}
	e := p.BootDevice.SetValue(v)
	if nil == e {
		if nil == p.BootDeviceItemDiscriminator_ {
			p.BootDeviceItemDiscriminator_ = new(string)
		}
		*p.BootDeviceItemDiscriminator_ = *p.BootDevice.Discriminator
	}
	return e
}

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/cdroms Get operation
*/
type ListCdromsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListCdromsResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListCdromsResponse() *ListCdromsResponse {
	p := new(ListCdromsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ListCdromsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.ListCdromsResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListCdromsResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListCdromsResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListCdromsResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/disks Get operation
*/
type ListDisksResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListDisksResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListDisksResponse() *ListDisksResponse {
	p := new(ListDisksResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ListDisksResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.ListDisksResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListDisksResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListDisksResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListDisksResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/gpus Get operation
*/
type ListGpusResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListGpusResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListGpusResponse() *ListGpusResponse {
	p := new(ListGpusResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ListGpusResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.ListGpusResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListGpusResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListGpusResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListGpusResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/nics Get operation
*/
type ListNicsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListNicsResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListNicsResponse() *ListNicsResponse {
	p := new(ListNicsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ListNicsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.ListNicsResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListNicsResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListNicsResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListNicsResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/serial-ports Get operation
*/
type ListSerialPortsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListSerialPortsResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListSerialPortsResponse() *ListSerialPortsResponse {
	p := new(ListSerialPortsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ListSerialPortsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.ListSerialPortsResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListSerialPortsResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListSerialPortsResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListSerialPortsResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms Get operation
*/
type ListVmsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListVmsResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListVmsResponse() *ListVmsResponse {
	p := new(ListVmsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ListVmsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.ListVmsResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListVmsResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListVmsResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListVmsResponseData()
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

/**
Machine type for the VM. Machine type Q35 is required for secure boot and does not support IDE disks.
*/
type MachineType int

const (
	MACHINETYPE_UNKNOWN  MachineType = 0
	MACHINETYPE_REDACTED MachineType = 1
	MACHINETYPE_PC       MachineType = 2
	MACHINETYPE_PSERIES  MachineType = 3
	MACHINETYPE_Q35      MachineType = 4
)

// returns the name of the enum given an ordinal number
func (e *MachineType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"PSERIES",
		"Q35",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *MachineType) index(name string) MachineType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"PSERIES",
		"Q35",
	}
	for idx := range names {
		if names[idx] == name {
			return MachineType(idx)
		}
	}
	return MACHINETYPE_UNKNOWN
}

func (e *MachineType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MachineType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MachineType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MachineType) Ref() *MachineType {
	return &e
}

/**
Configurations for migrating a VM NIC.
*/
type MigrateNicConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	IpAddress *import1.IPv4Address `json:"ipAddress,omitempty"`

	MigrateType *MigrateNicType `json:"migrateType"`

	Subnet *SubnetReference `json:"subnet"`
}

func (p *MigrateNicConfig) MarshalJSON() ([]byte, error) {
	type MigrateNicConfigProxy MigrateNicConfig
	return json.Marshal(struct {
		*MigrateNicConfigProxy
		MigrateType *MigrateNicType  `json:"migrateType,omitempty"`
		Subnet      *SubnetReference `json:"subnet,omitempty"`
	}{
		MigrateNicConfigProxy: (*MigrateNicConfigProxy)(p),
		MigrateType:           p.MigrateType,
		Subnet:                p.Subnet,
	})
}

func NewMigrateNicConfig() *MigrateNicConfig {
	p := new(MigrateNicConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.MigrateNicConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.MigrateNicConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/nics/{extId}/$actions/migrate Post operation
*/
type MigrateNicResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfMigrateNicResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewMigrateNicResponse() *MigrateNicResponse {
	p := new(MigrateNicResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.MigrateNicResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.MigrateNicResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *MigrateNicResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *MigrateNicResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfMigrateNicResponseData()
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

/**
The type of IP address management for NIC migration.
*/
type MigrateNicType int

const (
	MIGRATENICTYPE_UNKNOWN    MigrateNicType = 0
	MIGRATENICTYPE_REDACTED   MigrateNicType = 1
	MIGRATENICTYPE_ASSIGN_IP  MigrateNicType = 2
	MIGRATENICTYPE_RELEASE_IP MigrateNicType = 3
)

// returns the name of the enum given an ordinal number
func (e *MigrateNicType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ASSIGN_IP",
		"RELEASE_IP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *MigrateNicType) index(name string) MigrateNicType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ASSIGN_IP",
		"RELEASE_IP",
	}
	for idx := range names {
		if names[idx] == name {
			return MigrateNicType(idx)
		}
	}
	return MIGRATENICTYPE_UNKNOWN
}

func (e *MigrateNicType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MigrateNicType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MigrateNicType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MigrateNicType) Ref() *MigrateNicType {
	return &e
}

/**
The network function chain associates with the NIC. Only valid if nic_type is NORMAL_NIC.
*/
type NetworkFunctionChainReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Globally unique identifier of a network function chain. It should be of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func NewNetworkFunctionChainReference() *NetworkFunctionChainReference {
	p := new(NetworkFunctionChainReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.NetworkFunctionChainReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.NetworkFunctionChainReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
The type of this Network function NIC. Defaults to INGRESS.
*/
type NetworkFunctionNicType int

const (
	NETWORKFUNCTIONNICTYPE_UNKNOWN  NetworkFunctionNicType = 0
	NETWORKFUNCTIONNICTYPE_REDACTED NetworkFunctionNicType = 1
	NETWORKFUNCTIONNICTYPE_INGRESS  NetworkFunctionNicType = 2
	NETWORKFUNCTIONNICTYPE_EGRESS   NetworkFunctionNicType = 3
	NETWORKFUNCTIONNICTYPE_TAP      NetworkFunctionNicType = 4
)

// returns the name of the enum given an ordinal number
func (e *NetworkFunctionNicType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INGRESS",
		"EGRESS",
		"TAP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *NetworkFunctionNicType) index(name string) NetworkFunctionNicType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INGRESS",
		"EGRESS",
		"TAP",
	}
	for idx := range names {
		if names[idx] == name {
			return NetworkFunctionNicType(idx)
		}
	}
	return NETWORKFUNCTIONNICTYPE_UNKNOWN
}

func (e *NetworkFunctionNicType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NetworkFunctionNicType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NetworkFunctionNicType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NetworkFunctionNicType) Ref() *NetworkFunctionNicType {
	return &e
}

/**
The capabilities of the Nutanix Guest Tools in the VM.
*/
type NgtCapability int

const (
	NGTCAPABILITY_UNKNOWN              NgtCapability = 0
	NGTCAPABILITY_REDACTED             NgtCapability = 1
	NGTCAPABILITY_SELF_SERVICE_RESTORE NgtCapability = 2
	NGTCAPABILITY_VSS_SNAPSHOT         NgtCapability = 3
)

// returns the name of the enum given an ordinal number
func (e *NgtCapability) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SELF_SERVICE_RESTORE",
		"VSS_SNAPSHOT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *NgtCapability) index(name string) NgtCapability {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SELF_SERVICE_RESTORE",
		"VSS_SNAPSHOT",
	}
	for idx := range names {
		if names[idx] == name {
			return NgtCapability(idx)
		}
	}
	return NGTCAPABILITY_UNKNOWN
}

func (e *NgtCapability) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NgtCapability:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NgtCapability) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NgtCapability) Ref() *NgtCapability {
	return &e
}

/**
Virtual Machine NIC.
*/
type Nic struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BackingInfo *EmulatedNic `json:"backingInfo,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`

	NetworkInfo *NicNetworkInfo `json:"networkInfo,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewNic() *Nic {
	p := new(Nic)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Nic"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.Nic"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Network information for a NIC.
*/
type NicNetworkInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Indicates whether an unknown unicast traffic is forwarded to this NIC or not. This is applicable only for the NICs on the overlay subnets.
	*/
	AllowUnknownMacs *bool `json:"allowUnknownMacs,omitempty"`

	Ipv4Config *Ipv4Config `json:"ipv4Config,omitempty"`

	Ipv4Info *Ipv4Info `json:"ipv4Info,omitempty"`

	NetworkFunctionChain *NetworkFunctionChainReference `json:"networkFunctionChain,omitempty"`

	NetworkFunctionNicType *NetworkFunctionNicType `json:"networkFunctionNicType,omitempty"`

	NicType *NicType `json:"nicType,omitempty"`

	Subnet *SubnetReference `json:"subnet,omitempty"`
	/**
	  List of networks to trunk if VLAN mode is marked as TRUNKED. If empty and VLAN mode is set to TRUNKED, all the VLANs are trunked.
	*/
	TrunkedVlans []int `json:"trunkedVlans,omitempty"`

	VlanMode *VlanMode `json:"vlanMode,omitempty"`
}

func NewNicNetworkInfo() *NicNetworkInfo {
	p := new(NicNetworkInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.NicNetworkInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.NicNetworkInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
NIC type. Defaults to NORMAL_NIC.
*/
type NicType int

const (
	NICTYPE_UNKNOWN              NicType = 0
	NICTYPE_REDACTED             NicType = 1
	NICTYPE_NORMAL_NIC           NicType = 2
	NICTYPE_DIRECT_NIC           NicType = 3
	NICTYPE_NETWORK_FUNCTION_NIC NicType = 4
	NICTYPE_SPAN_DESTINATION_NIC NicType = 5
)

// returns the name of the enum given an ordinal number
func (e *NicType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NORMAL_NIC",
		"DIRECT_NIC",
		"NETWORK_FUNCTION_NIC",
		"SPAN_DESTINATION_NIC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *NicType) index(name string) NicType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NORMAL_NIC",
		"DIRECT_NIC",
		"NETWORK_FUNCTION_NIC",
		"SPAN_DESTINATION_NIC",
	}
	for idx := range names {
		if names[idx] == name {
			return NicType(idx)
		}
	}
	return NICTYPE_UNKNOWN
}

func (e *NicType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NicType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NicType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NicType) Ref() *NicType {
	return &e
}

/**
Configuration for NVRAM to be presented to the VM.
*/
type NvramDevice struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BackingStorageInfo *VmDisk `json:"backingStorageInfo,omitempty"`
}

func NewNvramDevice() *NvramDevice {
	p := new(NvramDevice)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.NvramDevice"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.NvramDevice"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Owner reference.
*/
type OwnerReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Globally unique identifier of a VM owner. It should be of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func NewOwnerReference() *OwnerReference {
	p := new(OwnerReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.OwnerReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.OwnerReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Ownership information for the VM.
*/
type OwnershipInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Owner *OwnerReference `json:"owner,omitempty"`
}

func NewOwnershipInfo() *OwnershipInfo {
	p := new(OwnershipInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.OwnershipInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.OwnershipInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Compliance state enum.
*/
type PolicyComplianceState int

const (
	POLICYCOMPLIANCESTATE_UNKNOWN       PolicyComplianceState = 0
	POLICYCOMPLIANCESTATE_REDACTED      PolicyComplianceState = 1
	POLICYCOMPLIANCESTATE_COMPLIANT     PolicyComplianceState = 2
	POLICYCOMPLIANCESTATE_NON_COMPLIANT PolicyComplianceState = 3
	POLICYCOMPLIANCESTATE_UNKNOWN_STATE PolicyComplianceState = 4
)

// returns the name of the enum given an ordinal number
func (e *PolicyComplianceState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"COMPLIANT",
		"NON_COMPLIANT",
		"UNKNOWN_STATE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *PolicyComplianceState) index(name string) PolicyComplianceState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"COMPLIANT",
		"NON_COMPLIANT",
		"UNKNOWN_STATE",
	}
	for idx := range names {
		if names[idx] == name {
			return PolicyComplianceState(idx)
		}
	}
	return POLICYCOMPLIANCESTATE_UNKNOWN
}

func (e *PolicyComplianceState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PolicyComplianceState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PolicyComplianceState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PolicyComplianceState) Ref() *PolicyComplianceState {
	return &e
}

/**
Reference to the policy object in use.
*/
type PolicyReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Globally unique identifier of an instance. It should be of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func NewPolicyReference() *PolicyReference {
	p := new(PolicyReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.PolicyReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.PolicyReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/power-cycle Post operation
*/
type PowerCycleVmResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfPowerCycleVmResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewPowerCycleVmResponse() *PowerCycleVmResponse {
	p := new(PowerCycleVmResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.PowerCycleVmResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.PowerCycleVmResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *PowerCycleVmResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *PowerCycleVmResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfPowerCycleVmResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/power-off Post operation
*/
type PowerOffVmResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfPowerOffVmResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewPowerOffVmResponse() *PowerOffVmResponse {
	p := new(PowerOffVmResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.PowerOffVmResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.PowerOffVmResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *PowerOffVmResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *PowerOffVmResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfPowerOffVmResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/power-on Post operation
*/
type PowerOnVmResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfPowerOnVmResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewPowerOnVmResponse() *PowerOnVmResponse {
	p := new(PowerOnVmResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.PowerOnVmResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.PowerOnVmResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *PowerOnVmResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *PowerOnVmResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfPowerOnVmResponseData()
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

/**
The current power state of the VM.
*/
type PowerState int

const (
	POWERSTATE_UNKNOWN      PowerState = 0
	POWERSTATE_REDACTED     PowerState = 1
	POWERSTATE_TRUE         PowerState = 2
	POWERSTATE_FALSE        PowerState = 3
	POWERSTATE_PAUSED       PowerState = 4
	POWERSTATE_UNDETERMINED PowerState = 5
)

// returns the name of the enum given an ordinal number
func (e *PowerState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"true",
		"false",
		"PAUSED",
		"UNDETERMINED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *PowerState) index(name string) PowerState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"true",
		"false",
		"PAUSED",
		"UNDETERMINED",
	}
	for idx := range names {
		if names[idx] == name {
			return PowerState(idx)
		}
	}
	return POWERSTATE_UNKNOWN
}

func (e *PowerState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PowerState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PowerState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PowerState) Ref() *PowerState {
	return &e
}

/**
Information about the state of the policy.
*/
type ProtectionInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ReplicationStatus *ProtectionInfoReplicationStatus `json:"replicationStatus,omitempty"`
}

func NewProtectionInfo() *ProtectionInfo {
	p := new(ProtectionInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ProtectionInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.ProtectionInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Replication status of the entity.
*/
type ProtectionInfoReplicationStatus int

const (
	PROTECTIONINFOREPLICATIONSTATUS_UNKNOWN     ProtectionInfoReplicationStatus = 0
	PROTECTIONINFOREPLICATIONSTATUS_REDACTED    ProtectionInfoReplicationStatus = 1
	PROTECTIONINFOREPLICATIONSTATUS_SYNCED      ProtectionInfoReplicationStatus = 2
	PROTECTIONINFOREPLICATIONSTATUS_SYNCING     ProtectionInfoReplicationStatus = 3
	PROTECTIONINFOREPLICATIONSTATUS_OUT_OF_SYNC ProtectionInfoReplicationStatus = 4
)

// returns the name of the enum given an ordinal number
func (e *ProtectionInfoReplicationStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SYNCED",
		"SYNCING",
		"OUT_OF_SYNC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *ProtectionInfoReplicationStatus) index(name string) ProtectionInfoReplicationStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SYNCED",
		"SYNCING",
		"OUT_OF_SYNC",
	}
	for idx := range names {
		if names[idx] == name {
			return ProtectionInfoReplicationStatus(idx)
		}
	}
	return PROTECTIONINFOREPLICATIONSTATUS_UNKNOWN
}

func (e *ProtectionInfoReplicationStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ProtectionInfoReplicationStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ProtectionInfoReplicationStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ProtectionInfoReplicationStatus) Ref() *ProtectionInfoReplicationStatus {
	return &e
}

/**
Status of protection policy applied to this VM.
*/
type ProtectionPolicyState struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ComplianceStatus *PolicyComplianceState `json:"complianceStatus,omitempty"`

	EnforcementMode *EnforcementMode `json:"enforcementMode,omitempty"`

	Policy *PolicyReference `json:"policy,omitempty"`

	PolicyInfo *ProtectionInfo `json:"policyInfo,omitempty"`
}

func NewProtectionPolicyState() *ProtectionPolicyState {
	p := new(ProtectionPolicyState)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ProtectionPolicyState"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.ProtectionPolicyState"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
The type of protection applied on a VM. PD_PROTECTED indicates a VM is protected using the Prism Element. RULE_PROTECTED indicates a VM protection using the Prism Central.
*/
type ProtectionType int

const (
	PROTECTIONTYPE_UNKNOWN        ProtectionType = 0
	PROTECTIONTYPE_REDACTED       ProtectionType = 1
	PROTECTIONTYPE_UNPROTECTED    ProtectionType = 2
	PROTECTIONTYPE_PD_PROTECTED   ProtectionType = 3
	PROTECTIONTYPE_RULE_PROTECTED ProtectionType = 4
)

// returns the name of the enum given an ordinal number
func (e *ProtectionType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UNPROTECTED",
		"PD_PROTECTED",
		"RULE_PROTECTED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *ProtectionType) index(name string) ProtectionType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UNPROTECTED",
		"PD_PROTECTED",
		"RULE_PROTECTED",
	}
	for idx := range names {
		if names[idx] == name {
			return ProtectionType(idx)
		}
	}
	return PROTECTIONTYPE_UNKNOWN
}

func (e *ProtectionType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ProtectionType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ProtectionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ProtectionType) Ref() *ProtectionType {
	return &e
}

/**
QoS parameters to be enforced.
*/
type QosConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Throttled IOPS for the governed entities. The block size for the I/O is 32 kB.
	*/
	ThrottledIops *int `json:"throttledIops,omitempty"`
}

func NewQosConfig() *QosConfig {
	p := new(QosConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.QosConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.QosConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
The restart schedule after installing Nutanix Guest Tools.
*/
type RebootPreference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Schedule *RebootPreferenceSchedule `json:"schedule,omitempty"`

	ScheduleType *ScheduleType `json:"scheduleType,omitempty"`
}

func NewRebootPreference() *RebootPreference {
	p := new(RebootPreference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.RebootPreference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.RebootPreference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Restart schedule.
*/
type RebootPreferenceSchedule struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  The start time for a scheduled restart.
	*/
	StartTime *time.Time `json:"startTime,omitempty"`
}

func NewRebootPreferenceSchedule() *RebootPreferenceSchedule {
	p := new(RebootPreferenceSchedule)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.RebootPreferenceSchedule"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.RebootPreferenceSchedule"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/guest-reboot Post operation
*/
type RebootVmResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRebootVmResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRebootVmResponse() *RebootVmResponse {
	p := new(RebootVmResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.RebootVmResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.RebootVmResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RebootVmResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RebootVmResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRebootVmResponseData()
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

/**
Status of a recovery plan.
*/
type RecoveryPlanState struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ComplianceStatus *PolicyComplianceState `json:"complianceStatus,omitempty"`

	EnforcementMode *EnforcementMode `json:"enforcementMode,omitempty"`

	Policy *PolicyReference `json:"policy,omitempty"`

	PolicyInfo *ProtectionInfo `json:"policyInfo,omitempty"`
}

func NewRecoveryPlanState() *RecoveryPlanState {
	p := new(RecoveryPlanState)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.RecoveryPlanState"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.RecoveryPlanState"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/nics/{extId}/$actions/release-ip Post operation
*/
type ReleaseIpResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfReleaseIpResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewReleaseIpResponse() *ReleaseIpResponse {
	p := new(ReleaseIpResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ReleaseIpResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.ReleaseIpResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ReleaseIpResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ReleaseIpResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfReleaseIpResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/reset Post operation
*/
type ResetVmResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfResetVmResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewResetVmResponse() *ResetVmResponse {
	p := new(ResetVmResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ResetVmResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.ResetVmResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ResetVmResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ResetVmResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfResetVmResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/revert Post operation
*/
type RevertVmResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRevertVmResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRevertVmResponse() *RevertVmResponse {
	p := new(RevertVmResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.RevertVmResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.RevertVmResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RevertVmResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RevertVmResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRevertVmResponseData()
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

/**
The (S)egment:(B)us:(D)evice.(F)unction hardware address. See https://wiki.xen.org/wiki/Bus:Device.Function_(BDF)_Notation for more details.
*/
type SBDF struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Bus *int `json:"bus,omitempty"`

	Device *int `json:"device,omitempty"`

	Func *int `json:"func,omitempty"`

	Segment *int `json:"segment,omitempty"`
}

func NewSBDF() *SBDF {
	p := new(SBDF)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.SBDF"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.SBDF"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Schedule type. One of SKIP, IMMEDIATE, or LATER.
*/
type ScheduleType int

const (
	SCHEDULETYPE_UNKNOWN   ScheduleType = 0
	SCHEDULETYPE_REDACTED  ScheduleType = 1
	SCHEDULETYPE_SKIP      ScheduleType = 2
	SCHEDULETYPE_IMMEDIATE ScheduleType = 3
	SCHEDULETYPE_LATER     ScheduleType = 4
)

// returns the name of the enum given an ordinal number
func (e *ScheduleType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SKIP",
		"IMMEDIATE",
		"LATER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *ScheduleType) index(name string) ScheduleType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SKIP",
		"IMMEDIATE",
		"LATER",
	}
	for idx := range names {
		if names[idx] == name {
			return ScheduleType(idx)
		}
	}
	return SCHEDULETYPE_UNKNOWN
}

func (e *ScheduleType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ScheduleType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ScheduleType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ScheduleType) Ref() *ScheduleType {
	return &e
}

/**
Indicates the configuration of serial ports of the VM.
*/
type SerialPort struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  Index of the serial port.
	*/
	Index *int `json:"index,omitempty"`
	/**
	  Indicates whether the serial port is connected or not.
	*/
	IsConnected *bool `json:"isConnected,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewSerialPort() *SerialPort {
	p := new(SerialPort)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.SerialPort"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.SerialPort"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/guest-shutdown Post operation
*/
type ShutdownVmResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfShutdownVmResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewShutdownVmResponse() *ShutdownVmResponse {
	p := new(ShutdownVmResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ShutdownVmResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.ShutdownVmResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ShutdownVmResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ShutdownVmResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfShutdownVmResponseData()
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

/**
Network identifier for this adapter. Only valid if nic_type is NORMAL_NIC or DIRECT_NIC.
*/
type SubnetReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Globally unique identifier of a subnet. It should be of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func NewSubnetReference() *SubnetReference {
	p := new(SubnetReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.SubnetReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.SubnetReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
If this field is set, the guest will be customized using Sysprep. Either unattend_xml or custom_key_values should be provided. If custom_key_values are provided then the unattended answer file will be generated using these key-value pairs.
*/
type Sysprep struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	InstallType *InstallType `json:"installType,omitempty"`
	/**

	 */
	SysprepScriptItemDiscriminator_ *string `json:"$sysprepScriptItemDiscriminator,omitempty"`

	SysprepScript *OneOfSysprepSysprepScript `json:"sysprepScript,omitempty"`
}

func NewSysprep() *Sysprep {
	p := new(Sysprep)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Sysprep"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.Sysprep"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Sysprep) GetSysprepScript() interface{} {
	if nil == p.SysprepScript {
		return nil
	}
	return p.SysprepScript.GetValue()
}

func (p *Sysprep) SetSysprepScript(v interface{}) error {
	if nil == p.SysprepScript {
		p.SysprepScript = NewOneOfSysprepSysprepScript()
	}
	e := p.SysprepScript.SetValue(v)
	if nil == e {
		if nil == p.SysprepScriptItemDiscriminator_ {
			p.SysprepScriptItemDiscriminator_ = new(string)
		}
		*p.SysprepScriptItemDiscriminator_ = *p.SysprepScript.Discriminator
	}
	return e
}

/**
UEFI boot mode and its associated configuration.
*/
type UefiBoot struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	NvramDevice *NvramDevice `json:"nvramDevice,omitempty"`
	/**
	  Indicate whether to enable secure boot or not.
	*/
	SecureBootEnabled *bool `json:"secureBootEnabled,omitempty"`
}

func NewUefiBoot() *UefiBoot {
	p := new(UefiBoot)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.UefiBoot"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.UefiBoot"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
This field contains a Sysprep unattend xml definition, as a string. The value must be base64 encoded.
*/
type Unattendxml struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Value *string `json:"value,omitempty"`
}

func NewUnattendxml() *Unattendxml {
	p := new(Unattendxml)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Unattendxml"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.Unattendxml"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/guest-tools/$actions/uninstall Post operation
*/
type UninstallVmGuestToolsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUninstallVmGuestToolsResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUninstallVmGuestToolsResponse() *UninstallVmGuestToolsResponse {
	p := new(UninstallVmGuestToolsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.UninstallVmGuestToolsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.UninstallVmGuestToolsResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UninstallVmGuestToolsResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UninstallVmGuestToolsResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUninstallVmGuestToolsResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/cdroms/{extId} Put operation
*/
type UpdateCdromResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateCdromResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateCdromResponse() *UpdateCdromResponse {
	p := new(UpdateCdromResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.UpdateCdromResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.UpdateCdromResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateCdromResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateCdromResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateCdromResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/disks/{extId} Put operation
*/
type UpdateDiskResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateDiskResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateDiskResponse() *UpdateDiskResponse {
	p := new(UpdateDiskResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.UpdateDiskResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.UpdateDiskResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateDiskResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateDiskResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateDiskResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/guest-tools-config Put operation
*/
type UpdateGuestToolsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateGuestToolsResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateGuestToolsResponse() *UpdateGuestToolsResponse {
	p := new(UpdateGuestToolsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.UpdateGuestToolsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.UpdateGuestToolsResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateGuestToolsResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateGuestToolsResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateGuestToolsResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/nics/{extId} Put operation
*/
type UpdateNicResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateNicResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateNicResponse() *UpdateNicResponse {
	p := new(UpdateNicResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.UpdateNicResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.UpdateNicResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateNicResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateNicResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateNicResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{vmExtId}/serial-ports/{extId} Put operation
*/
type UpdateSerialPortResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateSerialPortResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateSerialPortResponse() *UpdateSerialPortResponse {
	p := new(UpdateSerialPortResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.UpdateSerialPortResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.UpdateSerialPortResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateSerialPortResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateSerialPortResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateSerialPortResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/categories Put operation
*/
type UpdateVmCategoriesResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateVmCategoriesResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateVmCategoriesResponse() *UpdateVmCategoriesResponse {
	p := new(UpdateVmCategoriesResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.UpdateVmCategoriesResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.UpdateVmCategoriesResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateVmCategoriesResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateVmCategoriesResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateVmCategoriesResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/ownership-info Put operation
*/
type UpdateVmOwnershipInfoResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateVmOwnershipInfoResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateVmOwnershipInfoResponse() *UpdateVmOwnershipInfoResponse {
	p := new(UpdateVmOwnershipInfoResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.UpdateVmOwnershipInfoResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.UpdateVmOwnershipInfoResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateVmOwnershipInfoResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateVmOwnershipInfoResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateVmOwnershipInfoResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId} Put operation
*/
type UpdateVmResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateVmResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateVmResponse() *UpdateVmResponse {
	p := new(UpdateVmResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.UpdateVmResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.UpdateVmResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateVmResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateVmResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateVmResponseData()
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

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/guest-tools/$actions/upgrade Post operation
*/
type UpgradeVmGuestToolsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpgradeVmGuestToolsResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpgradeVmGuestToolsResponse() *UpgradeVmGuestToolsResponse {
	p := new(UpgradeVmGuestToolsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.UpgradeVmGuestToolsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.UpgradeVmGuestToolsResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpgradeVmGuestToolsResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpgradeVmGuestToolsResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpgradeVmGuestToolsResponseData()
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

/**
The contents of the user_data configuration for cloud-init. This can be formatted as YAML, JSON, or could be a shell script. The value must be base64 encoded.
*/
type Userdata struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  The value for the cloud-init user_data.
	*/
	Value *string `json:"value,omitempty"`
}

func NewUserdata() *Userdata {
	p := new(Userdata)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Userdata"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.Userdata"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
By default, all the virtual NICs are created in ACCESS mode, which permits only one VLAN per virtual network. TRUNKED mode allows multiple VLANs on a single VM NIC for network-aware user VMs.
*/
type VlanMode int

const (
	VLANMODE_UNKNOWN  VlanMode = 0
	VLANMODE_REDACTED VlanMode = 1
	VLANMODE_ACCESS   VlanMode = 2
	VLANMODE_TRUNK    VlanMode = 3
)

// returns the name of the enum given an ordinal number
func (e *VlanMode) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACCESS",
		"TRUNK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *VlanMode) index(name string) VlanMode {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACCESS",
		"TRUNK",
	}
	for idx := range names {
		if names[idx] == name {
			return VlanMode(idx)
		}
	}
	return VLANMODE_UNKNOWN
}

func (e *VlanMode) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for VlanMode:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *VlanMode) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e VlanMode) Ref() *VlanMode {
	return &e
}

/**
VM configuration.
*/
type Vm struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AvailabilityZone *AvailabilityZoneReference `json:"availabilityZone,omitempty"`
	/**

	 */
	BootConfigItemDiscriminator_ *string `json:"$bootConfigItemDiscriminator,omitempty"`
	/**
	  Indicates the order of device types in which the VM should try to boot from. If the boot device order is not provided the system will decide an appropriate boot device order.
	*/
	BootConfig *OneOfVmBootConfig `json:"bootConfig,omitempty"`
	/**
	  Indicates whether to remove AHV branding from VM firmware tables or not.
	*/
	BrandingEnabled *bool `json:"brandingEnabled,omitempty"`
	/**
	  Categories for the VM.
	*/
	Categories []CategoryReference `json:"categories,omitempty"`
	/**
	  CD-ROMs attached to the VM.
	*/
	Cdroms []Cdrom `json:"cdroms,omitempty"`

	Cluster *ClusterReference `json:"cluster,omitempty"`
	/**
	  Indicates whether to passthrough the host CPU features to the guest or not. Enabling this will make VM incapable of live migration.
	*/
	CpuPassthroughEnabled *bool `json:"cpuPassthroughEnabled,omitempty"`
	/**
	  VM creation time.
	*/
	CreateTime *time.Time `json:"createTime,omitempty"`
	/**
	  VM description.
	*/
	Description *string `json:"description,omitempty"`
	/**
	  Disks attached to the VM.
	*/
	Disks []Disk `json:"disks,omitempty"`
	/**
	  The list of additional CPU features to be enabled. HardwareVirtualization: Indicates whether hardware assisted virtualization should be enabled for the Guest OS or not. Once enabled, the Guest OS can deploy a nested hypervisor.
	*/
	EnabledCpuFeatures []CpuFeature `json:"enabledCpuFeatures,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  Indicates whether the vGPU console is enabled or not.
	*/
	GpuConsoleEnabled *bool `json:"gpuConsoleEnabled,omitempty"`
	/**
	  GPUs attached to the VM.
	*/
	Gpus []Gpu `json:"gpus,omitempty"`

	GuestCustomization *GuestCustomization `json:"guestCustomization,omitempty"`

	GuestTools *GuestTools `json:"guestTools,omitempty"`
	/**
	  VM hardware clock timezone in IANA TZDB format (America/Los_Angeles).
	*/
	HardwareClockTimezone *string `json:"hardwareClockTimezone,omitempty"`

	Host *HostReference `json:"host,omitempty"`
	/**
	  Indicates whether the VM is an agent VM or not. When their host enters maintenance mode, once the normal VMs are evacuated, the agent VMs are powered off. When the host is restored, agent VMs are powered on before the normal VMs are restored. In other words, agent VMs cannot be HA-protected or live migrated.
	*/
	IsAgentVm *bool `json:"isAgentVm,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/**
	  Indicates whether the VM is capable of live migrating to another host or not, based on its configuration.
	*/
	LiveMigrateCapable *bool `json:"liveMigrateCapable,omitempty"`

	MachineType *MachineType `json:"machineType,omitempty"`
	/**
	  Indicates whether the memory overcommit feature should be enabled for the VM or not. If enabled, parts of the VM memory may reside outside of the hypervisor physical memory. Once enabled, it should be expected that the VM may suffer performance degradation.
	*/
	MemoryOvercommitEnabled *bool `json:"memoryOvercommitEnabled,omitempty"`
	/**
	  Memory size in bytes.
	*/
	MemorySizeBytes *int64 `json:"memorySizeBytes,omitempty"`
	/**
	  VM name.
	*/
	Name *string `json:"name,omitempty"`
	/**
	  NICs attached to the VM.
	*/
	Nics []Nic `json:"nics,omitempty"`
	/**
	  Number of cores per socket.
	*/
	NumCoresPerSocket *int `json:"numCoresPerSocket,omitempty"`
	/**
	  Number of vCPU sockets.
	*/
	NumSockets *int `json:"numSockets,omitempty"`
	/**
	  Number of threads per core.
	*/
	NumThreadsPerCore *int `json:"numThreadsPerCore,omitempty"`
	/**
	  Number of vNUMA nodes. 0 means vNUMA is disabled.
	*/
	NumVnumaNodes *int `json:"numVnumaNodes,omitempty"`

	OwnershipInfo *OwnershipInfo `json:"ownershipInfo,omitempty"`

	PowerState *PowerState `json:"powerState,omitempty"`

	ProtectionPolicyState *ProtectionPolicyState `json:"protectionPolicyState,omitempty"`

	ProtectionType *ProtectionType `json:"protectionType,omitempty"`
	/**
	  Status of the recovery plans associated with the VM.
	*/
	RecoveryPlanStates []RecoveryPlanState `json:"recoveryPlanStates,omitempty"`
	/**
	  Serial ports configured on the VM.
	*/
	SerialPorts []SerialPort `json:"serialPorts,omitempty"`

	Source *VmSourceReference `json:"source,omitempty"`

	StorageConfig *ADSFVmStorageConfig `json:"storageConfig,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/**
	  VM last updated time.
	*/
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	/**
	  Indicates whether the vCPUs should be hard pinned to specific pCPUs or not.
	*/
	VcpuHardPinningEnabled *bool `json:"vcpuHardPinningEnabled,omitempty"`
	/**
	  Indicates whether the VGA console should be disabled or not.
	*/
	VgaConsoleEnabled *bool `json:"vgaConsoleEnabled,omitempty"`

	VtpmConfig *VtpmConfig `json:"vtpmConfig,omitempty"`
}

func NewVm() *Vm {
	p := new(Vm)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Vm"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.Vm"}
	p.UnknownFields_ = map[string]interface{}{}

	p.BrandingEnabled = new(bool)
	*p.BrandingEnabled = true
	p.HardwareClockTimezone = new(string)
	*p.HardwareClockTimezone = "UTC"
	p.NumCoresPerSocket = new(int)
	*p.NumCoresPerSocket = 1
	p.NumThreadsPerCore = new(int)
	*p.NumThreadsPerCore = 1
	p.VgaConsoleEnabled = new(bool)
	*p.VgaConsoleEnabled = true

	return p
}

/**
Cross cluster migration overrides the VM configuration.
*/
type VmCrossClusterMigrateOverrides struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  The NIC configuration to apply on the destination cluster when migrating a VM from one cluster to another.
	*/
	OverrideNicList []Nic `json:"overrideNicList,omitempty"`
}

func NewVmCrossClusterMigrateOverrides() *VmCrossClusterMigrateOverrides {
	p := new(VmCrossClusterMigrateOverrides)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmCrossClusterMigrateOverrides"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.VmCrossClusterMigrateOverrides"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Input on how to migrate a VM across cluster
*/
type VmCrossClusterMigrateParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Indicates if the migration is performed with a running VM.
	*/
	IsLiveMigration *bool `json:"isLiveMigration,omitempty"`

	Overrides *VmCrossClusterMigrateOverrides `json:"overrides,omitempty"`

	TargetAvailabilityZone *AvailabilityZoneReference `json:"targetAvailabilityZone,omitempty"`

	TargetCluster *ClusterReference `json:"targetCluster,omitempty"`
}

func NewVmCrossClusterMigrateParams() *VmCrossClusterMigrateParams {
	p := new(VmCrossClusterMigrateParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmCrossClusterMigrateParams"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.VmCrossClusterMigrateParams"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/migrate Post operation
*/
type VmCrossClusterMigrateWithStorageResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfVmCrossClusterMigrateWithStorageResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewVmCrossClusterMigrateWithStorageResponse() *VmCrossClusterMigrateWithStorageResponse {
	p := new(VmCrossClusterMigrateWithStorageResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmCrossClusterMigrateWithStorageResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.VmCrossClusterMigrateWithStorageResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *VmCrossClusterMigrateWithStorageResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *VmCrossClusterMigrateWithStorageResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfVmCrossClusterMigrateWithStorageResponseData()
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

/**
Storage provided by Nutanix ADSF
*/
type VmDisk struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DataSource *DataSource `json:"dataSource,omitempty"`
	/**
	  Globally unique identifier of a VM disk. It should be of type UUID.
	*/
	DiskExtId *string `json:"diskExtId,omitempty"`
	/**
	  Size of the disk in Bytes
	*/
	DiskSizeBytes *int64 `json:"diskSizeBytes,omitempty"`

	StorageConfig *VmDiskStorageConfig `json:"storageConfig,omitempty"`

	StorageContainer *VmDiskContainerReference `json:"storageContainer,omitempty"`
}

func NewVmDisk() *VmDisk {
	p := new(VmDisk)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmDisk"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.VmDisk"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
This reference is for disk level storage container preference. This preference specifies the storage container to which this disk belongs.
*/
type VmDiskContainerReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Globally unique identifier of a VM disk container. It should be of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func NewVmDiskContainerReference() *VmDiskContainerReference {
	p := new(VmDiskContainerReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmDiskContainerReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.VmDiskContainerReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Reference to an existing VM disk.
*/
type VmDiskReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DiskAddress *DiskAddress `json:"diskAddress,omitempty"`
	/**
	  Globally unique identifier of a VM disk. It should be of type UUID.
	*/
	DiskExtId *string `json:"diskExtId,omitempty"`

	VmReference *VmReference `json:"vmReference,omitempty"`
}

func NewVmDiskReference() *VmDiskReference {
	p := new(VmDiskReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmDiskReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.VmDiskReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Storage configuration for VM disks.
*/
type VmDiskStorageConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Indicates whether the virtual disk is pinned to the hot tier or not.
	*/
	FlashModeEnabled *bool `json:"flashModeEnabled,omitempty"`
}

func NewVmDiskStorageConfig() *VmDiskStorageConfig {
	p := new(VmDiskStorageConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmDiskStorageConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.VmDiskStorageConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
A model that represents VM Recovery Point properties
*/
type VmRecoveryPoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  The external identifier that can be used to retrieve the consistency group using its URL.
	*/
	ConsistencyGroupExtId *string `json:"consistencyGroupExtId,omitempty"`
	/**
	  Name of the Consistency group.
	*/
	ConsistencyGroupName *string `json:"consistencyGroupName,omitempty"`
	/**
	  The UTC date and time in ISO-8601 format when the Recovery point is created.
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`
	/**
	  The UTC date and time in ISO-8601 format when the current Recovery point expires and will be garbage collected.
	*/
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/**
	  Location agnostic identifier of the Recovery point.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/**
	  The name of the Recovery point.
	*/
	Name *string `json:"name,omitempty"`

	RecoveryPointType *import5.RecoveryPointType `json:"recoveryPointType,omitempty"`

	Status *import5.RecoveryPointStatus `json:"status,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/**
	  List of additional metadata provided by the client at the time of Recovery point creation.
	*/
	VendorSpecificProperties []import5.VendorSpecificProperty `json:"vendorSpecificProperties,omitempty"`

	Vm *Vm `json:"vm,omitempty"`
}

func NewVmRecoveryPoint() *VmRecoveryPoint {
	p := new(VmRecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmRecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.VmRecoveryPoint"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
This is a reference to a VM.
*/
type VmReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Globally unique identifier of a VM. It should be of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func NewVmReference() *VmReference {
	p := new(VmReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.VmReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Specify the VM Recovery Point Id to which the VM would be reverted.
*/
type VmRevert struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  The external identifier of the VM Recovery Point.
	*/
	VmRecoveryPointExtId *string `json:"vmRecoveryPointExtId"`
}

func (p *VmRevert) MarshalJSON() ([]byte, error) {
	type VmRevertProxy VmRevert
	return json.Marshal(struct {
		*VmRevertProxy
		VmRecoveryPointExtId *string `json:"vmRecoveryPointExtId,omitempty"`
	}{
		VmRevertProxy:        (*VmRevertProxy)(p),
		VmRecoveryPointExtId: p.VmRecoveryPointExtId,
	})
}

func NewVmRevert() *VmRevert {
	p := new(VmRevert)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmRevert"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.VmRevert"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Reference to an entity that the VM should be cloned or created from.
*/
type VmSourceReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EntityType *VmSourceReferenceEntityType `json:"entityType,omitempty"`
	/**
	  Globally unique identifier of a VM. It should be of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func NewVmSourceReference() *VmSourceReference {
	p := new(VmSourceReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmSourceReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.VmSourceReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type VmSourceReferenceEntityType int

const (
	VMSOURCEREFERENCEENTITYTYPE_UNKNOWN           VmSourceReferenceEntityType = 0
	VMSOURCEREFERENCEENTITYTYPE_REDACTED          VmSourceReferenceEntityType = 1
	VMSOURCEREFERENCEENTITYTYPE_VM                VmSourceReferenceEntityType = 2
	VMSOURCEREFERENCEENTITYTYPE_VM_RECOVERY_POINT VmSourceReferenceEntityType = 3
)

// returns the name of the enum given an ordinal number
func (e *VmSourceReferenceEntityType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"VM_RECOVERY_POINT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *VmSourceReferenceEntityType) index(name string) VmSourceReferenceEntityType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"VM_RECOVERY_POINT",
	}
	for idx := range names {
		if names[idx] == name {
			return VmSourceReferenceEntityType(idx)
		}
	}
	return VMSOURCEREFERENCEENTITYTYPE_UNKNOWN
}

func (e *VmSourceReferenceEntityType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for VmSourceReferenceEntityType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *VmSourceReferenceEntityType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e VmSourceReferenceEntityType) Ref() *VmSourceReferenceEntityType {
	return &e
}

/**
Indicates how the vTPM for the VM should be configured.
*/
type VtpmConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Virtual trusted platform module version.
	*/
	Version *string `json:"version,omitempty"`
	/**
	  Indicates whether the virtual trusted platform module is enabled for the Guest OS or not.
	*/
	VtpmEnabled *bool `json:"vtpmEnabled,omitempty"`
}

func NewVtpmConfig() *VtpmConfig {
	p := new(VtpmConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VtpmConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.ahv.config.VtpmConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfListDisksResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 []Disk                 `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfListDisksResponseData() *OneOfListDisksResponseData {
	p := new(OneOfListDisksResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListDisksResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListDisksResponseData is nil"))
	}
	switch v.(type) {
	case []Disk:
		p.oneOfType2001 = v.([]Disk)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.config.Disk>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.config.Disk>"
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfListDisksResponseData) GetValue() interface{} {
	if "List<vmm.v4.ahv.config.Disk>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListDisksResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new([]Disk)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.config.Disk" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.config.Disk>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.config.Disk>"
			return nil

		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListDisksResponseData"))
}

func (p *OneOfListDisksResponseData) MarshalJSON() ([]byte, error) {
	if "List<vmm.v4.ahv.config.Disk>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListDisksResponseData")
}

type OneOfCreateDiskResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfCreateDiskResponseData() *OneOfCreateDiskResponseData {
	p := new(OneOfCreateDiskResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateDiskResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateDiskResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfCreateDiskResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateDiskResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateDiskResponseData"))
}

func (p *OneOfCreateDiskResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateDiskResponseData")
}

type OneOfLegacyBootBootDevice struct {
	Discriminator *string         `json:"-"`
	ObjectType_   *string         `json:"-"`
	oneOfType2002 *BootDeviceNic  `json:"-"`
	oneOfType2001 *BootDeviceDisk `json:"-"`
}

func NewOneOfLegacyBootBootDevice() *OneOfLegacyBootBootDevice {
	p := new(OneOfLegacyBootBootDevice)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfLegacyBootBootDevice) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfLegacyBootBootDevice is nil"))
	}
	switch v.(type) {
	case BootDeviceNic:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(BootDeviceNic)
		}
		*p.oneOfType2002 = v.(BootDeviceNic)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	case BootDeviceDisk:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(BootDeviceDisk)
		}
		*p.oneOfType2001 = v.(BootDeviceDisk)
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

func (p *OneOfLegacyBootBootDevice) GetValue() interface{} {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfLegacyBootBootDevice) UnmarshalJSON(b []byte) error {
	vOneOfType2002 := new(BootDeviceNic)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.config.BootDeviceNic" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(BootDeviceNic)
			}
			*p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2002.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new(BootDeviceDisk)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.BootDeviceDisk" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(BootDeviceDisk)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfLegacyBootBootDevice"))
}

func (p *OneOfLegacyBootBootDevice) MarshalJSON() ([]byte, error) {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfLegacyBootBootDevice")
}

type OneOfListGpusResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 []Gpu                  `json:"-"`
}

func NewOneOfListGpusResponseData() *OneOfListGpusResponseData {
	p := new(OneOfListGpusResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListGpusResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListGpusResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []Gpu:
		p.oneOfType2001 = v.([]Gpu)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.config.Gpu>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.config.Gpu>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListGpusResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<vmm.v4.ahv.config.Gpu>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListGpusResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType2001 := new([]Gpu)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.config.Gpu" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.config.Gpu>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.config.Gpu>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListGpusResponseData"))
}

func (p *OneOfListGpusResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<vmm.v4.ahv.config.Gpu>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListGpusResponseData")
}

type OneOfDetachVmCategoryResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfDetachVmCategoryResponseData() *OneOfDetachVmCategoryResponseData {
	p := new(OneOfDetachVmCategoryResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDetachVmCategoryResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDetachVmCategoryResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfDetachVmCategoryResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDetachVmCategoryResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDetachVmCategoryResponseData"))
}

func (p *OneOfDetachVmCategoryResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDetachVmCategoryResponseData")
}

type OneOfPowerCycleVmResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfPowerCycleVmResponseData() *OneOfPowerCycleVmResponseData {
	p := new(OneOfPowerCycleVmResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfPowerCycleVmResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfPowerCycleVmResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfPowerCycleVmResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfPowerCycleVmResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPowerCycleVmResponseData"))
}

func (p *OneOfPowerCycleVmResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfPowerCycleVmResponseData")
}

type OneOfListSerialPortsResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 []SerialPort           `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfListSerialPortsResponseData() *OneOfListSerialPortsResponseData {
	p := new(OneOfListSerialPortsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListSerialPortsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListSerialPortsResponseData is nil"))
	}
	switch v.(type) {
	case []SerialPort:
		p.oneOfType2001 = v.([]SerialPort)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.config.SerialPort>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.config.SerialPort>"
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfListSerialPortsResponseData) GetValue() interface{} {
	if "List<vmm.v4.ahv.config.SerialPort>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListSerialPortsResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new([]SerialPort)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.config.SerialPort" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.config.SerialPort>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.config.SerialPort>"
			return nil

		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListSerialPortsResponseData"))
}

func (p *OneOfListSerialPortsResponseData) MarshalJSON() ([]byte, error) {
	if "List<vmm.v4.ahv.config.SerialPort>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListSerialPortsResponseData")
}

type OneOfShutdownVmResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfShutdownVmResponseData() *OneOfShutdownVmResponseData {
	p := new(OneOfShutdownVmResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfShutdownVmResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfShutdownVmResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfShutdownVmResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfShutdownVmResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfShutdownVmResponseData"))
}

func (p *OneOfShutdownVmResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfShutdownVmResponseData")
}

type OneOfGetGpuResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *Gpu                   `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfGetGpuResponseData() *OneOfGetGpuResponseData {
	p := new(OneOfGetGpuResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetGpuResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetGpuResponseData is nil"))
	}
	switch v.(type) {
	case Gpu:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Gpu)
		}
		*p.oneOfType2001 = v.(Gpu)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfGetGpuResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetGpuResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(Gpu)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.Gpu" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Gpu)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetGpuResponseData"))
}

func (p *OneOfGetGpuResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetGpuResponseData")
}

type OneOfPowerOnVmResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfPowerOnVmResponseData() *OneOfPowerOnVmResponseData {
	p := new(OneOfPowerOnVmResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfPowerOnVmResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfPowerOnVmResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfPowerOnVmResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfPowerOnVmResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPowerOnVmResponseData"))
}

func (p *OneOfPowerOnVmResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfPowerOnVmResponseData")
}

type OneOfUpgradeVmGuestToolsResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfUpgradeVmGuestToolsResponseData() *OneOfUpgradeVmGuestToolsResponseData {
	p := new(OneOfUpgradeVmGuestToolsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpgradeVmGuestToolsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpgradeVmGuestToolsResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfUpgradeVmGuestToolsResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpgradeVmGuestToolsResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpgradeVmGuestToolsResponseData"))
}

func (p *OneOfUpgradeVmGuestToolsResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpgradeVmGuestToolsResponseData")
}

type OneOfReleaseIpResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfReleaseIpResponseData() *OneOfReleaseIpResponseData {
	p := new(OneOfReleaseIpResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfReleaseIpResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfReleaseIpResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfReleaseIpResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfReleaseIpResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfReleaseIpResponseData"))
}

func (p *OneOfReleaseIpResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfReleaseIpResponseData")
}

type OneOfGuestCustomizationConfig struct {
	Discriminator *string    `json:"-"`
	ObjectType_   *string    `json:"-"`
	oneOfType2001 *Sysprep   `json:"-"`
	oneOfType2002 *CloudInit `json:"-"`
}

func NewOneOfGuestCustomizationConfig() *OneOfGuestCustomizationConfig {
	p := new(OneOfGuestCustomizationConfig)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGuestCustomizationConfig) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGuestCustomizationConfig is nil"))
	}
	switch v.(type) {
	case Sysprep:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Sysprep)
		}
		*p.oneOfType2001 = v.(Sysprep)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case CloudInit:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(CloudInit)
		}
		*p.oneOfType2002 = v.(CloudInit)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGuestCustomizationConfig) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	return nil
}

func (p *OneOfGuestCustomizationConfig) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(Sysprep)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.Sysprep" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Sysprep)
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
	vOneOfType2002 := new(CloudInit)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.config.CloudInit" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(CloudInit)
			}
			*p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2002.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGuestCustomizationConfig"))
}

func (p *OneOfGuestCustomizationConfig) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	return nil, errors.New("No value to marshal for OneOfGuestCustomizationConfig")
}

type OneOfDeleteCdromResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfDeleteCdromResponseData() *OneOfDeleteCdromResponseData {
	p := new(OneOfDeleteCdromResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteCdromResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteCdromResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfDeleteCdromResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteCdromResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteCdromResponseData"))
}

func (p *OneOfDeleteCdromResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteCdromResponseData")
}

type OneOfInsertCdromResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfInsertCdromResponseData() *OneOfInsertCdromResponseData {
	p := new(OneOfInsertCdromResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfInsertCdromResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfInsertCdromResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfInsertCdromResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfInsertCdromResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfInsertCdromResponseData"))
}

func (p *OneOfInsertCdromResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfInsertCdromResponseData")
}

type OneOfRevertVmResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfRevertVmResponseData() *OneOfRevertVmResponseData {
	p := new(OneOfRevertVmResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRevertVmResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRevertVmResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfRevertVmResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRevertVmResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRevertVmResponseData"))
}

func (p *OneOfRevertVmResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRevertVmResponseData")
}

type OneOfListVmsResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 []Vm                   `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfListVmsResponseData() *OneOfListVmsResponseData {
	p := new(OneOfListVmsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListVmsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListVmsResponseData is nil"))
	}
	switch v.(type) {
	case []Vm:
		p.oneOfType2001 = v.([]Vm)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.config.Vm>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.config.Vm>"
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfListVmsResponseData) GetValue() interface{} {
	if "List<vmm.v4.ahv.config.Vm>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListVmsResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new([]Vm)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.config.Vm" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.config.Vm>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.config.Vm>"
			return nil

		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListVmsResponseData"))
}

func (p *OneOfListVmsResponseData) MarshalJSON() ([]byte, error) {
	if "List<vmm.v4.ahv.config.Vm>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListVmsResponseData")
}

type OneOfUpdateVmOwnershipInfoResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfUpdateVmOwnershipInfoResponseData() *OneOfUpdateVmOwnershipInfoResponseData {
	p := new(OneOfUpdateVmOwnershipInfoResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateVmOwnershipInfoResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateVmOwnershipInfoResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfUpdateVmOwnershipInfoResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateVmOwnershipInfoResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateVmOwnershipInfoResponseData"))
}

func (p *OneOfUpdateVmOwnershipInfoResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateVmOwnershipInfoResponseData")
}

type OneOfDeleteNicResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfDeleteNicResponseData() *OneOfDeleteNicResponseData {
	p := new(OneOfDeleteNicResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteNicResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteNicResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfDeleteNicResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteNicResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteNicResponseData"))
}

func (p *OneOfDeleteNicResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteNicResponseData")
}

type OneOfDeleteGpuResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfDeleteGpuResponseData() *OneOfDeleteGpuResponseData {
	p := new(OneOfDeleteGpuResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteGpuResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteGpuResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfDeleteGpuResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteGpuResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteGpuResponseData"))
}

func (p *OneOfDeleteGpuResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteGpuResponseData")
}

type OneOfGetVmCategoriesResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 []CategoryReference    `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfGetVmCategoriesResponseData() *OneOfGetVmCategoriesResponseData {
	p := new(OneOfGetVmCategoriesResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVmCategoriesResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVmCategoriesResponseData is nil"))
	}
	switch v.(type) {
	case []CategoryReference:
		p.oneOfType2001 = v.([]CategoryReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.config.CategoryReference>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.config.CategoryReference>"
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfGetVmCategoriesResponseData) GetValue() interface{} {
	if "List<vmm.v4.ahv.config.CategoryReference>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetVmCategoriesResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new([]CategoryReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.config.CategoryReference" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.config.CategoryReference>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.config.CategoryReference>"
			return nil

		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVmCategoriesResponseData"))
}

func (p *OneOfGetVmCategoriesResponseData) MarshalJSON() ([]byte, error) {
	if "List<vmm.v4.ahv.config.CategoryReference>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetVmCategoriesResponseData")
}

type OneOfCreateGpuResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfCreateGpuResponseData() *OneOfCreateGpuResponseData {
	p := new(OneOfCreateGpuResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateGpuResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateGpuResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfCreateGpuResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateGpuResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateGpuResponseData"))
}

func (p *OneOfCreateGpuResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateGpuResponseData")
}

type OneOfCloneVmResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfCloneVmResponseData() *OneOfCloneVmResponseData {
	p := new(OneOfCloneVmResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCloneVmResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCloneVmResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfCloneVmResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCloneVmResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCloneVmResponseData"))
}

func (p *OneOfCloneVmResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCloneVmResponseData")
}

type OneOfUpdateCdromResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfUpdateCdromResponseData() *OneOfUpdateCdromResponseData {
	p := new(OneOfUpdateCdromResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateCdromResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateCdromResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfUpdateCdromResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateCdromResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateCdromResponseData"))
}

func (p *OneOfUpdateCdromResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateCdromResponseData")
}

type OneOfUpdateGuestToolsResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfUpdateGuestToolsResponseData() *OneOfUpdateGuestToolsResponseData {
	p := new(OneOfUpdateGuestToolsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateGuestToolsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateGuestToolsResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfUpdateGuestToolsResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateGuestToolsResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateGuestToolsResponseData"))
}

func (p *OneOfUpdateGuestToolsResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateGuestToolsResponseData")
}

type OneOfVmCrossClusterMigrateWithStorageResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfVmCrossClusterMigrateWithStorageResponseData() *OneOfVmCrossClusterMigrateWithStorageResponseData {
	p := new(OneOfVmCrossClusterMigrateWithStorageResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVmCrossClusterMigrateWithStorageResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVmCrossClusterMigrateWithStorageResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfVmCrossClusterMigrateWithStorageResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfVmCrossClusterMigrateWithStorageResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVmCrossClusterMigrateWithStorageResponseData"))
}

func (p *OneOfVmCrossClusterMigrateWithStorageResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfVmCrossClusterMigrateWithStorageResponseData")
}

type OneOfGetCdromResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *Cdrom                 `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfGetCdromResponseData() *OneOfGetCdromResponseData {
	p := new(OneOfGetCdromResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetCdromResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetCdromResponseData is nil"))
	}
	switch v.(type) {
	case Cdrom:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Cdrom)
		}
		*p.oneOfType2001 = v.(Cdrom)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfGetCdromResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetCdromResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(Cdrom)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.Cdrom" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Cdrom)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetCdromResponseData"))
}

func (p *OneOfGetCdromResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetCdromResponseData")
}

type OneOfEjectCdromResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfEjectCdromResponseData() *OneOfEjectCdromResponseData {
	p := new(OneOfEjectCdromResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfEjectCdromResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfEjectCdromResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfEjectCdromResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfEjectCdromResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfEjectCdromResponseData"))
}

func (p *OneOfEjectCdromResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfEjectCdromResponseData")
}

type OneOfAssignIpResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfAssignIpResponseData() *OneOfAssignIpResponseData {
	p := new(OneOfAssignIpResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAssignIpResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAssignIpResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfAssignIpResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAssignIpResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAssignIpResponseData"))
}

func (p *OneOfAssignIpResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAssignIpResponseData")
}

type OneOfGetVmResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *Vm                    `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfGetVmResponseData() *OneOfGetVmResponseData {
	p := new(OneOfGetVmResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVmResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVmResponseData is nil"))
	}
	switch v.(type) {
	case Vm:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Vm)
		}
		*p.oneOfType2001 = v.(Vm)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfGetVmResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetVmResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(Vm)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.Vm" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Vm)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVmResponseData"))
}

func (p *OneOfGetVmResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetVmResponseData")
}

type OneOfUpdateVmResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfUpdateVmResponseData() *OneOfUpdateVmResponseData {
	p := new(OneOfUpdateVmResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateVmResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateVmResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfUpdateVmResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateVmResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateVmResponseData"))
}

func (p *OneOfUpdateVmResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateVmResponseData")
}

type OneOfInsertVmGuestToolsResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfInsertVmGuestToolsResponseData() *OneOfInsertVmGuestToolsResponseData {
	p := new(OneOfInsertVmGuestToolsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfInsertVmGuestToolsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfInsertVmGuestToolsResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfInsertVmGuestToolsResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfInsertVmGuestToolsResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfInsertVmGuestToolsResponseData"))
}

func (p *OneOfInsertVmGuestToolsResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfInsertVmGuestToolsResponseData")
}

type OneOfVmBootConfig struct {
	Discriminator *string     `json:"-"`
	ObjectType_   *string     `json:"-"`
	oneOfType2001 *LegacyBoot `json:"-"`
	oneOfType2002 *UefiBoot   `json:"-"`
}

func NewOneOfVmBootConfig() *OneOfVmBootConfig {
	p := new(OneOfVmBootConfig)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVmBootConfig) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVmBootConfig is nil"))
	}
	switch v.(type) {
	case LegacyBoot:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(LegacyBoot)
		}
		*p.oneOfType2001 = v.(LegacyBoot)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case UefiBoot:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(UefiBoot)
		}
		*p.oneOfType2002 = v.(UefiBoot)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfVmBootConfig) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	return nil
}

func (p *OneOfVmBootConfig) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(LegacyBoot)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.LegacyBoot" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(LegacyBoot)
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
	vOneOfType2002 := new(UefiBoot)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.config.UefiBoot" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(UefiBoot)
			}
			*p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2002.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVmBootConfig"))
}

func (p *OneOfVmBootConfig) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	return nil, errors.New("No value to marshal for OneOfVmBootConfig")
}

type OneOfRebootVmResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfRebootVmResponseData() *OneOfRebootVmResponseData {
	p := new(OneOfRebootVmResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRebootVmResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRebootVmResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfRebootVmResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRebootVmResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRebootVmResponseData"))
}

func (p *OneOfRebootVmResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRebootVmResponseData")
}

type OneOfAttachVmCategoryResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfAttachVmCategoryResponseData() *OneOfAttachVmCategoryResponseData {
	p := new(OneOfAttachVmCategoryResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAttachVmCategoryResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAttachVmCategoryResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfAttachVmCategoryResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAttachVmCategoryResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAttachVmCategoryResponseData"))
}

func (p *OneOfAttachVmCategoryResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAttachVmCategoryResponseData")
}

type OneOfUpdateDiskResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfUpdateDiskResponseData() *OneOfUpdateDiskResponseData {
	p := new(OneOfUpdateDiskResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateDiskResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateDiskResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfUpdateDiskResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateDiskResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateDiskResponseData"))
}

func (p *OneOfUpdateDiskResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateDiskResponseData")
}

type OneOfUninstallVmGuestToolsResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfUninstallVmGuestToolsResponseData() *OneOfUninstallVmGuestToolsResponseData {
	p := new(OneOfUninstallVmGuestToolsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUninstallVmGuestToolsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUninstallVmGuestToolsResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfUninstallVmGuestToolsResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUninstallVmGuestToolsResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUninstallVmGuestToolsResponseData"))
}

func (p *OneOfUninstallVmGuestToolsResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUninstallVmGuestToolsResponseData")
}

type OneOfDeleteSerialPortResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfDeleteSerialPortResponseData() *OneOfDeleteSerialPortResponseData {
	p := new(OneOfDeleteSerialPortResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteSerialPortResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteSerialPortResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfDeleteSerialPortResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteSerialPortResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteSerialPortResponseData"))
}

func (p *OneOfDeleteSerialPortResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteSerialPortResponseData")
}

type OneOfListCdromsResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 []Cdrom                `json:"-"`
}

func NewOneOfListCdromsResponseData() *OneOfListCdromsResponseData {
	p := new(OneOfListCdromsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListCdromsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListCdromsResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []Cdrom:
		p.oneOfType2001 = v.([]Cdrom)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.config.Cdrom>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.config.Cdrom>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListCdromsResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<vmm.v4.ahv.config.Cdrom>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListCdromsResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType2001 := new([]Cdrom)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.config.Cdrom" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.config.Cdrom>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.config.Cdrom>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListCdromsResponseData"))
}

func (p *OneOfListCdromsResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<vmm.v4.ahv.config.Cdrom>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListCdromsResponseData")
}

type OneOfMigrateNicResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfMigrateNicResponseData() *OneOfMigrateNicResponseData {
	p := new(OneOfMigrateNicResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfMigrateNicResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfMigrateNicResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfMigrateNicResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfMigrateNicResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfMigrateNicResponseData"))
}

func (p *OneOfMigrateNicResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfMigrateNicResponseData")
}

type OneOfPowerOffVmResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfPowerOffVmResponseData() *OneOfPowerOffVmResponseData {
	p := new(OneOfPowerOffVmResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfPowerOffVmResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfPowerOffVmResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfPowerOffVmResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfPowerOffVmResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPowerOffVmResponseData"))
}

func (p *OneOfPowerOffVmResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfPowerOffVmResponseData")
}

type OneOfCustomizeGuestVmResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfCustomizeGuestVmResponseData() *OneOfCustomizeGuestVmResponseData {
	p := new(OneOfCustomizeGuestVmResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCustomizeGuestVmResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCustomizeGuestVmResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfCustomizeGuestVmResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCustomizeGuestVmResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCustomizeGuestVmResponseData"))
}

func (p *OneOfCustomizeGuestVmResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCustomizeGuestVmResponseData")
}

type OneOfUpdateSerialPortResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfUpdateSerialPortResponseData() *OneOfUpdateSerialPortResponseData {
	p := new(OneOfUpdateSerialPortResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateSerialPortResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateSerialPortResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfUpdateSerialPortResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateSerialPortResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateSerialPortResponseData"))
}

func (p *OneOfUpdateSerialPortResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateSerialPortResponseData")
}

type OneOfGetNicResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *Nic                   `json:"-"`
}

func NewOneOfGetNicResponseData() *OneOfGetNicResponseData {
	p := new(OneOfGetNicResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetNicResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetNicResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Nic:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Nic)
		}
		*p.oneOfType2001 = v.(Nic)
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

func (p *OneOfGetNicResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetNicResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType2001 := new(Nic)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.Nic" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Nic)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetNicResponseData"))
}

func (p *OneOfGetNicResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetNicResponseData")
}

type OneOfUpdateVmCategoriesResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfUpdateVmCategoriesResponseData() *OneOfUpdateVmCategoriesResponseData {
	p := new(OneOfUpdateVmCategoriesResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateVmCategoriesResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateVmCategoriesResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfUpdateVmCategoriesResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateVmCategoriesResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateVmCategoriesResponseData"))
}

func (p *OneOfUpdateVmCategoriesResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateVmCategoriesResponseData")
}

type OneOfGetVmOwnershipInfoResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *OwnershipInfo         `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfGetVmOwnershipInfoResponseData() *OneOfGetVmOwnershipInfoResponseData {
	p := new(OneOfGetVmOwnershipInfoResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVmOwnershipInfoResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVmOwnershipInfoResponseData is nil"))
	}
	switch v.(type) {
	case OwnershipInfo:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(OwnershipInfo)
		}
		*p.oneOfType2001 = v.(OwnershipInfo)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfGetVmOwnershipInfoResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetVmOwnershipInfoResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(OwnershipInfo)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.OwnershipInfo" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(OwnershipInfo)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVmOwnershipInfoResponseData"))
}

func (p *OneOfGetVmOwnershipInfoResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetVmOwnershipInfoResponseData")
}

type OneOfGetDiskResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *Disk                  `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfGetDiskResponseData() *OneOfGetDiskResponseData {
	p := new(OneOfGetDiskResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetDiskResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetDiskResponseData is nil"))
	}
	switch v.(type) {
	case Disk:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Disk)
		}
		*p.oneOfType2001 = v.(Disk)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfGetDiskResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetDiskResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(Disk)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.Disk" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Disk)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetDiskResponseData"))
}

func (p *OneOfGetDiskResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetDiskResponseData")
}

type OneOfCreateVmResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfCreateVmResponseData() *OneOfCreateVmResponseData {
	p := new(OneOfCreateVmResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateVmResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateVmResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfCreateVmResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateVmResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateVmResponseData"))
}

func (p *OneOfCreateVmResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateVmResponseData")
}

type OneOfCreateSerialPortResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfCreateSerialPortResponseData() *OneOfCreateSerialPortResponseData {
	p := new(OneOfCreateSerialPortResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateSerialPortResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateSerialPortResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfCreateSerialPortResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateSerialPortResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateSerialPortResponseData"))
}

func (p *OneOfCreateSerialPortResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateSerialPortResponseData")
}

type OneOfInstallVmGuestToolsResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfInstallVmGuestToolsResponseData() *OneOfInstallVmGuestToolsResponseData {
	p := new(OneOfInstallVmGuestToolsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfInstallVmGuestToolsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfInstallVmGuestToolsResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfInstallVmGuestToolsResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfInstallVmGuestToolsResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfInstallVmGuestToolsResponseData"))
}

func (p *OneOfInstallVmGuestToolsResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfInstallVmGuestToolsResponseData")
}

type OneOfGetGuestToolsResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *GuestTools            `json:"-"`
}

func NewOneOfGetGuestToolsResponseData() *OneOfGetGuestToolsResponseData {
	p := new(OneOfGetGuestToolsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetGuestToolsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetGuestToolsResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case GuestTools:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(GuestTools)
		}
		*p.oneOfType2001 = v.(GuestTools)
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

func (p *OneOfGetGuestToolsResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetGuestToolsResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType2001 := new(GuestTools)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.GuestTools" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(GuestTools)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetGuestToolsResponseData"))
}

func (p *OneOfGetGuestToolsResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetGuestToolsResponseData")
}

type OneOfDiskBackingInfo struct {
	Discriminator *string                   `json:"-"`
	ObjectType_   *string                   `json:"-"`
	oneOfType2001 *VmDisk                   `json:"-"`
	oneOfType2002 *ADSFVolumeGroupReference `json:"-"`
}

func NewOneOfDiskBackingInfo() *OneOfDiskBackingInfo {
	p := new(OneOfDiskBackingInfo)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDiskBackingInfo) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDiskBackingInfo is nil"))
	}
	switch v.(type) {
	case VmDisk:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(VmDisk)
		}
		*p.oneOfType2001 = v.(VmDisk)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case ADSFVolumeGroupReference:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(ADSFVolumeGroupReference)
		}
		*p.oneOfType2002 = v.(ADSFVolumeGroupReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDiskBackingInfo) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	return nil
}

func (p *OneOfDiskBackingInfo) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(VmDisk)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.VmDisk" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(VmDisk)
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
	vOneOfType2002 := new(ADSFVolumeGroupReference)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.config.ADSFVolumeGroupReference" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(ADSFVolumeGroupReference)
			}
			*p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2002.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDiskBackingInfo"))
}

func (p *OneOfDiskBackingInfo) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	return nil, errors.New("No value to marshal for OneOfDiskBackingInfo")
}

type OneOfSysprepSysprepScript struct {
	Discriminator *string          `json:"-"`
	ObjectType_   *string          `json:"-"`
	oneOfType2002 *CustomKeyValues `json:"-"`
	oneOfType2001 *Unattendxml     `json:"-"`
}

func NewOneOfSysprepSysprepScript() *OneOfSysprepSysprepScript {
	p := new(OneOfSysprepSysprepScript)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfSysprepSysprepScript) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfSysprepSysprepScript is nil"))
	}
	switch v.(type) {
	case CustomKeyValues:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(CustomKeyValues)
		}
		*p.oneOfType2002 = v.(CustomKeyValues)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	case Unattendxml:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Unattendxml)
		}
		*p.oneOfType2001 = v.(Unattendxml)
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

func (p *OneOfSysprepSysprepScript) GetValue() interface{} {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfSysprepSysprepScript) UnmarshalJSON(b []byte) error {
	vOneOfType2002 := new(CustomKeyValues)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.config.CustomKeyValues" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(CustomKeyValues)
			}
			*p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2002.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new(Unattendxml)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.Unattendxml" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Unattendxml)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSysprepSysprepScript"))
}

func (p *OneOfSysprepSysprepScript) MarshalJSON() ([]byte, error) {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfSysprepSysprepScript")
}

type OneOfCreateNicResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfCreateNicResponseData() *OneOfCreateNicResponseData {
	p := new(OneOfCreateNicResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateNicResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateNicResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfCreateNicResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateNicResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateNicResponseData"))
}

func (p *OneOfCreateNicResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateNicResponseData")
}

type OneOfDeleteVmResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfDeleteVmResponseData() *OneOfDeleteVmResponseData {
	p := new(OneOfDeleteVmResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteVmResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteVmResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfDeleteVmResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteVmResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteVmResponseData"))
}

func (p *OneOfDeleteVmResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteVmResponseData")
}

type OneOfGetSerialPortResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *SerialPort            `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfGetSerialPortResponseData() *OneOfGetSerialPortResponseData {
	p := new(OneOfGetSerialPortResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetSerialPortResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetSerialPortResponseData is nil"))
	}
	switch v.(type) {
	case SerialPort:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(SerialPort)
		}
		*p.oneOfType2001 = v.(SerialPort)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfGetSerialPortResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetSerialPortResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(SerialPort)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.SerialPort" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(SerialPort)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetSerialPortResponseData"))
}

func (p *OneOfGetSerialPortResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetSerialPortResponseData")
}

type OneOfCloudInitCloudInitScript struct {
	Discriminator *string          `json:"-"`
	ObjectType_   *string          `json:"-"`
	oneOfType2002 *CustomKeyValues `json:"-"`
	oneOfType2001 *Userdata        `json:"-"`
}

func NewOneOfCloudInitCloudInitScript() *OneOfCloudInitCloudInitScript {
	p := new(OneOfCloudInitCloudInitScript)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCloudInitCloudInitScript) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCloudInitCloudInitScript is nil"))
	}
	switch v.(type) {
	case CustomKeyValues:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(CustomKeyValues)
		}
		*p.oneOfType2002 = v.(CustomKeyValues)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	case Userdata:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Userdata)
		}
		*p.oneOfType2001 = v.(Userdata)
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

func (p *OneOfCloudInitCloudInitScript) GetValue() interface{} {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfCloudInitCloudInitScript) UnmarshalJSON(b []byte) error {
	vOneOfType2002 := new(CustomKeyValues)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.config.CustomKeyValues" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(CustomKeyValues)
			}
			*p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2002.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new(Userdata)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.Userdata" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Userdata)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCloudInitCloudInitScript"))
}

func (p *OneOfCloudInitCloudInitScript) MarshalJSON() ([]byte, error) {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfCloudInitCloudInitScript")
}

type OneOfUpdateNicResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfUpdateNicResponseData() *OneOfUpdateNicResponseData {
	p := new(OneOfUpdateNicResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateNicResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateNicResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfUpdateNicResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateNicResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateNicResponseData"))
}

func (p *OneOfUpdateNicResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateNicResponseData")
}

type OneOfDeleteDiskResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfDeleteDiskResponseData() *OneOfDeleteDiskResponseData {
	p := new(OneOfDeleteDiskResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteDiskResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteDiskResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfDeleteDiskResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteDiskResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteDiskResponseData"))
}

func (p *OneOfDeleteDiskResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteDiskResponseData")
}

type OneOfCreateCdromResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfCreateCdromResponseData() *OneOfCreateCdromResponseData {
	p := new(OneOfCreateCdromResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateCdromResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateCdromResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfCreateCdromResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateCdromResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateCdromResponseData"))
}

func (p *OneOfCreateCdromResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateCdromResponseData")
}

type OneOfCloneOverrideBootConfig struct {
	Discriminator *string     `json:"-"`
	ObjectType_   *string     `json:"-"`
	oneOfType2001 *LegacyBoot `json:"-"`
	oneOfType2002 *UefiBoot   `json:"-"`
}

func NewOneOfCloneOverrideBootConfig() *OneOfCloneOverrideBootConfig {
	p := new(OneOfCloneOverrideBootConfig)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCloneOverrideBootConfig) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCloneOverrideBootConfig is nil"))
	}
	switch v.(type) {
	case LegacyBoot:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(LegacyBoot)
		}
		*p.oneOfType2001 = v.(LegacyBoot)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case UefiBoot:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(UefiBoot)
		}
		*p.oneOfType2002 = v.(UefiBoot)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCloneOverrideBootConfig) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	return nil
}

func (p *OneOfCloneOverrideBootConfig) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(LegacyBoot)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.LegacyBoot" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(LegacyBoot)
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
	vOneOfType2002 := new(UefiBoot)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.config.UefiBoot" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(UefiBoot)
			}
			*p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2002.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCloneOverrideBootConfig"))
}

func (p *OneOfCloneOverrideBootConfig) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	return nil, errors.New("No value to marshal for OneOfCloneOverrideBootConfig")
}

type OneOfListNicsResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 []Nic                  `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfListNicsResponseData() *OneOfListNicsResponseData {
	p := new(OneOfListNicsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListNicsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListNicsResponseData is nil"))
	}
	switch v.(type) {
	case []Nic:
		p.oneOfType2001 = v.([]Nic)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.config.Nic>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.config.Nic>"
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfListNicsResponseData) GetValue() interface{} {
	if "List<vmm.v4.ahv.config.Nic>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListNicsResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new([]Nic)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.config.Nic" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.config.Nic>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.config.Nic>"
			return nil

		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListNicsResponseData"))
}

func (p *OneOfListNicsResponseData) MarshalJSON() ([]byte, error) {
	if "List<vmm.v4.ahv.config.Nic>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListNicsResponseData")
}

type OneOfResetVmResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfResetVmResponseData() *OneOfResetVmResponseData {
	p := new(OneOfResetVmResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfResetVmResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfResetVmResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfResetVmResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfResetVmResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfResetVmResponseData"))
}

func (p *OneOfResetVmResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfResetVmResponseData")
}

type OneOfDataSourceReference struct {
	Discriminator *string          `json:"-"`
	ObjectType_   *string          `json:"-"`
	oneOfType2002 *VmDiskReference `json:"-"`
	oneOfType2001 *ImageReference  `json:"-"`
}

func NewOneOfDataSourceReference() *OneOfDataSourceReference {
	p := new(OneOfDataSourceReference)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDataSourceReference) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDataSourceReference is nil"))
	}
	switch v.(type) {
	case VmDiskReference:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(VmDiskReference)
		}
		*p.oneOfType2002 = v.(VmDiskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	case ImageReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(ImageReference)
		}
		*p.oneOfType2001 = v.(ImageReference)
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

func (p *OneOfDataSourceReference) GetValue() interface{} {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfDataSourceReference) UnmarshalJSON(b []byte) error {
	vOneOfType2002 := new(VmDiskReference)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.config.VmDiskReference" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(VmDiskReference)
			}
			*p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2002.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new(ImageReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.ImageReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(ImageReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDataSourceReference"))
}

func (p *OneOfDataSourceReference) MarshalJSON() ([]byte, error) {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfDataSourceReference")
}
