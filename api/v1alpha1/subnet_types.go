/*
Copyright 2024 The ORC Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

// TODO validations:
//
// * IP addresses in CIDR, AllocationPools, Gateway, DNSNameserver(?), and
//   HostRoutes match the version in IPVersion (Spec and SubnetFilter)
// * IPv6 may only be set if IPVersion is 6 (Spec and SubnetFilter)
// * AllocationPools must be in CIDR

type SubnetRefs struct {
	// networkRef is a reference to the ORC Network which this subnet is associated with.
	// +required
	NetworkRef KubernetesNameRef `json:"networkRef"`
}

// SubnetFilter specifies a filter to select a subnet. At least one parameter must be specified.
// +kubebuilder:validation:MinProperties:=1
type SubnetFilter struct {
	Name *OpenStackName `json:"name,omitempty"`

	// description of the existing resource
	// +optional
	Description *NeutronDescription `json:"description,omitempty"`

	IPVersion *IPVersion   `json:"ipVersion,omitempty"`
	GatewayIP *IPvAny      `json:"gatewayIP,omitempty"`
	CIDR      *CIDR        `json:"cidr,omitempty"`
	IPv6      *IPv6Options `json:"ipv6,omitempty"`

	FilterByNeutronTags `json:",inline"`
}

type SubnetResourceSpec struct {
	// name is a human-readable name of the subnet. If not set, the object's name will be used.
	// +optional
	Name *OpenStackName `json:"name,omitempty"`

	// description is a human-readable description for the resource.
	// +optional
	Description *NeutronDescription `json:"description,omitempty"`

	// tags is a list of tags which will be applied to the subnet.
	// +kubebuilder:validation:MaxItems:=32
	// +listType=set
	Tags []NeutronTag `json:"tags,omitempty"`

	// ipVersion is the IP version for the subnet.
	// +required
	IPVersion IPVersion `json:"ipVersion"`

	// cidr is the address CIDR of the subnet. It must match the IP version specified in IPVersion.
	// +required
	CIDR CIDR `json:"cidr"`

	// allocationPools are IP Address pools that will be available for DHCP. IP
	// addresses must be in CIDR.
	// +kubebuilder:validation:MaxItems:=32
	// +listType=atomic
	AllocationPools []AllocationPool `json:"allocationPools,omitempty"`

	// gateway specifies the default gateway of the subnet. If not specified,
	// neutron will add one automatically. To disable this behaviour, specify a
	// gateway with a type of None.
	// +optional
	Gateway *SubnetGateway `json:"gateway,omitempty"`

	// enableDHCP will either enable to disable the DHCP service.
	// +optional
	EnableDHCP *bool `json:"enableDHCP,omitempty"`

	// dnsNameservers are the nameservers to be set via DHCP.
	// +kubebuilder:validation:MaxItems:=16
	// +listType=set
	DNSNameservers []IPvAny `json:"dnsNameservers,omitempty"`

	// dnsPublishFixedIP will either enable or disable the publication of fixed IPs to the DNS
	// +optional
	DNSPublishFixedIP *bool `json:"dnsPublishFixedIP,omitempty"`

	// hostRoutes are any static host routes to be set via DHCP.
	// +kubebuilder:validation:MaxItems:=256
	// +listType=atomic
	HostRoutes []HostRoute `json:"hostRoutes,omitempty"`

	// ipv6 contains IPv6-specific options. It may only be set if IPVersion is 6.
	IPv6 *IPv6Options `json:"ipv6,omitempty"`

	// routerRef specifies a router to attach the subnet to
	// +optional
	RouterRef *KubernetesNameRef `json:"routerRef,omitempty"`

	// TODO: Support service types
	// TODO: Support subnet pools
}

type AllocationPoolStatus struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type HostRouteStatus struct {
	Destination string `json:"destination"`
	NextHop     string `json:"nextHop"`
}

type SubnetResourceStatus struct {
	// name is the human-readable name of the subnet. Might not be unique.
	Name string `json:"name"`

	// description is a human-readable description for the resource.
	// +optional
	Description string `json:"description,omitempty"`

	// ipVersion specifies IP version, either `4' or `6'.
	IPVersion int `json:"ipVersion"`

	// cidr representing IP range for this subnet, based on IP version.
	CIDR string `json:"cidr"`

	// gatewayIP is the default gateway used by devices in this subnet, if any.
	// +optional
	GatewayIP string `json:"gatewayIP,omitempty"`

	// dnsNameservers is a list of name servers used by hosts in this subnet.
	// +listType=atomic
	DNSNameservers []string `json:"dnsNameservers,omitempty"`

	// dnsPublishFixedIP specifies whether the fixed IP addresses are published to the DNS.
	DNSPublishFixedIP bool `json:"dnsPublishFixedIP,omitempty"`

	// allocationPools is a list of sub-ranges within CIDR available for dynamic
	// allocation to ports.
	// +listType=atomic
	AllocationPools []AllocationPoolStatus `json:"allocationPools,omitempty"`

	// hostRoutes is a list of routes that should be used by devices with IPs
	// from this subnet (not including local subnet route).
	// +listType=atomic
	HostRoutes []HostRouteStatus `json:"hostRoutes,omitempty"`

	// Specifies whether DHCP is enabled for this subnet or not.
	EnableDHCP bool `json:"enableDHCP"`

	// projectID is the project owner of the subnet.
	ProjectID string `json:"projectID,omitempty"`

	// The IPv6 address modes specifies mechanisms for assigning IPv6 IP addresses.
	// +optional
	IPv6AddressMode string `json:"ipv6AddressMode,omitempty"`

	// The IPv6 router advertisement specifies whether the networking service
	// should transmit ICMPv6 packets.
	// +optional
	IPv6RAMode string `json:"ipv6RAMode,omitempty"`

	// subnetPoolID is the id of the subnet pool associated with the subnet.
	// +optional
	SubnetPoolID string `json:"subnetPoolID,omitempty"`

	// tags optionally set via extensions/attributestags
	// +listType=atomic
	Tags []string `json:"tags,omitempty"`

	NeutronStatusMetadata `json:",inline"`
}

// +kubebuilder:validation:Enum:=slaac;dhcpv6-stateful;dhcpv6-stateless
type IPv6AddressMode string

const (
	IPv6AddressModeSLAAC           = "slaac"
	IPv6AddressModeDHCPv6Stateful  = "dhcpv6-stateful"
	IPv6AddressModeDHCPv6Stateless = "dhcpv6-stateless"
)

// +kubebuilder:validation:Enum:=slaac;dhcpv6-stateful;dhcpv6-stateless
type IPv6RAMode string

const (
	IPv6RAModeSLAAC           = "slaac"
	IPv6RAModeDHCPv6Stateful  = "dhcpv6-stateful"
	IPv6RAModeDHCPv6Stateless = "dhcpv6-stateless"
)

// +kubebuilder:validation:MinProperties:=1
type IPv6Options struct {
	// addressMode specifies mechanisms for assigning IPv6 IP addresses.
	AddressMode *IPv6AddressMode `json:"addressMode,omitempty"`

	// raMode specifies the IPv6 router advertisement mode. It specifies whether
	// the networking service should transmit ICMPv6 packets.
	RAMode *IPv6RAMode `json:"raMode,omitempty"`
}

type SubnetGatewayType string

const (
	SubnetGatewayTypeNone      = "None"
	SubnetGatewayTypeAutomatic = "Automatic"
	SubnetGatewayTypeIP        = "IP"
)

type SubnetGateway struct {
	// type specifies how the default gateway will be created. `Automatic`
	// specifies that neutron will automatically add a default gateway. This is
	// also the default if no Gateway is specified. `None` specifies that the
	// subnet will not have a default gateway. `IP` specifies that the subnet
	// will use a specific address as the default gateway, which must be
	// specified in `IP`.
	// +kubebuilder:validation:Enum:=None;Automatic;IP
	// +required
	Type SubnetGatewayType `json:"type"`

	// ip is the IP address of the default gateway, which must be specified if
	// Type is `IP`. It must be a valid IP address, either IPv4 or IPv6,
	// matching the IPVersion in SubnetResourceSpec.
	// +optional
	IP *IPvAny `json:"ip,omitempty"`
}

type AllocationPool struct {
	// +required
	Start IPvAny `json:"start"`

	// +required
	End IPvAny `json:"end"`
}

type HostRoute struct {
	Destination CIDR   `json:"destination"`
	NextHop     IPvAny `json:"nextHop"`
}
