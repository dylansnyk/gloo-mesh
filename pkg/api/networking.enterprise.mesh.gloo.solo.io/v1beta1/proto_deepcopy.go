// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package v1beta1

import (
	proto "github.com/golang/protobuf/proto"
	"github.com/solo-io/protoc-gen-ext/pkg/clone"
)

// DeepCopyInto for the WasmDeployment.Spec
func (in *WasmDeploymentSpec) DeepCopyInto(out *WasmDeploymentSpec) {
	var p *WasmDeploymentSpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*WasmDeploymentSpec)
	} else {
		p = proto.Clone(in).(*WasmDeploymentSpec)
	}
	*out = *p
}

// DeepCopyInto for the WasmDeployment.Status
func (in *WasmDeploymentStatus) DeepCopyInto(out *WasmDeploymentStatus) {
	var p *WasmDeploymentStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*WasmDeploymentStatus)
	} else {
		p = proto.Clone(in).(*WasmDeploymentStatus)
	}
	*out = *p
}

// DeepCopyInto for the RateLimitClientConfig.Spec
func (in *RateLimitClientConfigSpec) DeepCopyInto(out *RateLimitClientConfigSpec) {
	var p *RateLimitClientConfigSpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*RateLimitClientConfigSpec)
	} else {
		p = proto.Clone(in).(*RateLimitClientConfigSpec)
	}
	*out = *p
}

// DeepCopyInto for the RateLimitClientConfig.Status
func (in *RateLimitClientConfigStatus) DeepCopyInto(out *RateLimitClientConfigStatus) {
	var p *RateLimitClientConfigStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*RateLimitClientConfigStatus)
	} else {
		p = proto.Clone(in).(*RateLimitClientConfigStatus)
	}
	*out = *p
}

// DeepCopyInto for the RateLimitServerConfig.Spec
func (in *RateLimitServerConfigSpec) DeepCopyInto(out *RateLimitServerConfigSpec) {
	var p *RateLimitServerConfigSpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*RateLimitServerConfigSpec)
	} else {
		p = proto.Clone(in).(*RateLimitServerConfigSpec)
	}
	*out = *p
}

// DeepCopyInto for the RateLimitServerConfig.Status
func (in *RateLimitServerConfigStatus) DeepCopyInto(out *RateLimitServerConfigStatus) {
	var p *RateLimitServerConfigStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*RateLimitServerConfigStatus)
	} else {
		p = proto.Clone(in).(*RateLimitServerConfigStatus)
	}
	*out = *p
}

// DeepCopyInto for the VirtualDestination.Spec
func (in *VirtualDestinationSpec) DeepCopyInto(out *VirtualDestinationSpec) {
	var p *VirtualDestinationSpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*VirtualDestinationSpec)
	} else {
		p = proto.Clone(in).(*VirtualDestinationSpec)
	}
	*out = *p
}

// DeepCopyInto for the VirtualDestination.Status
func (in *VirtualDestinationStatus) DeepCopyInto(out *VirtualDestinationStatus) {
	var p *VirtualDestinationStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*VirtualDestinationStatus)
	} else {
		p = proto.Clone(in).(*VirtualDestinationStatus)
	}
	*out = *p
}

// DeepCopyInto for the VirtualGateway.Spec
func (in *VirtualGatewaySpec) DeepCopyInto(out *VirtualGatewaySpec) {
	var p *VirtualGatewaySpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*VirtualGatewaySpec)
	} else {
		p = proto.Clone(in).(*VirtualGatewaySpec)
	}
	*out = *p
}

// DeepCopyInto for the VirtualGateway.Status
func (in *VirtualGatewayStatus) DeepCopyInto(out *VirtualGatewayStatus) {
	var p *VirtualGatewayStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*VirtualGatewayStatus)
	} else {
		p = proto.Clone(in).(*VirtualGatewayStatus)
	}
	*out = *p
}

// DeepCopyInto for the VirtualHost.Spec
func (in *VirtualHostSpec) DeepCopyInto(out *VirtualHostSpec) {
	var p *VirtualHostSpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*VirtualHostSpec)
	} else {
		p = proto.Clone(in).(*VirtualHostSpec)
	}
	*out = *p
}

// DeepCopyInto for the VirtualHost.Status
func (in *VirtualHostStatus) DeepCopyInto(out *VirtualHostStatus) {
	var p *VirtualHostStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*VirtualHostStatus)
	} else {
		p = proto.Clone(in).(*VirtualHostStatus)
	}
	*out = *p
}

// DeepCopyInto for the RouteTable.Spec
func (in *RouteTableSpec) DeepCopyInto(out *RouteTableSpec) {
	var p *RouteTableSpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*RouteTableSpec)
	} else {
		p = proto.Clone(in).(*RouteTableSpec)
	}
	*out = *p
}

// DeepCopyInto for the RouteTable.Status
func (in *RouteTableStatus) DeepCopyInto(out *RouteTableStatus) {
	var p *RouteTableStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*RouteTableStatus)
	} else {
		p = proto.Clone(in).(*RouteTableStatus)
	}
	*out = *p
}

// DeepCopyInto for the ServiceDependency.Spec
func (in *ServiceDependencySpec) DeepCopyInto(out *ServiceDependencySpec) {
	var p *ServiceDependencySpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*ServiceDependencySpec)
	} else {
		p = proto.Clone(in).(*ServiceDependencySpec)
	}
	*out = *p
}

// DeepCopyInto for the ServiceDependency.Status
func (in *ServiceDependencyStatus) DeepCopyInto(out *ServiceDependencyStatus) {
	var p *ServiceDependencyStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*ServiceDependencyStatus)
	} else {
		p = proto.Clone(in).(*ServiceDependencyStatus)
	}
	*out = *p
}

// DeepCopyInto for the CertificateVerification.Spec
func (in *CertificateVerificationSpec) DeepCopyInto(out *CertificateVerificationSpec) {
	var p *CertificateVerificationSpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*CertificateVerificationSpec)
	} else {
		p = proto.Clone(in).(*CertificateVerificationSpec)
	}
	*out = *p
}

// DeepCopyInto for the CertificateVerification.Status
func (in *CertificateVerificationStatus) DeepCopyInto(out *CertificateVerificationStatus) {
	var p *CertificateVerificationStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*CertificateVerificationStatus)
	} else {
		p = proto.Clone(in).(*CertificateVerificationStatus)
	}
	*out = *p
}
