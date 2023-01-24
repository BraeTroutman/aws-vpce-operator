//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDns) DeepCopyInto(out *CustomDns) {
	*out = *in
	out.Route53PrivateHostedZone = in.Route53PrivateHostedZone
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDns.
func (in *CustomDns) DeepCopy() *CustomDns {
	if in == nil {
		return nil
	}
	out := new(CustomDns)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalNameService) DeepCopyInto(out *ExternalNameService) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalNameService.
func (in *ExternalNameService) DeepCopy() *ExternalNameService {
	if in == nil {
		return nil
	}
	out := new(ExternalNameService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53HostedZoneRecord) DeepCopyInto(out *Route53HostedZoneRecord) {
	*out = *in
	out.ExternalNameService = in.ExternalNameService
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53HostedZoneRecord.
func (in *Route53HostedZoneRecord) DeepCopy() *Route53HostedZoneRecord {
	if in == nil {
		return nil
	}
	out := new(Route53HostedZoneRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53PrivateHostedZone) DeepCopyInto(out *Route53PrivateHostedZone) {
	*out = *in
	out.Record = in.Record
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53PrivateHostedZone.
func (in *Route53PrivateHostedZone) DeepCopy() *Route53PrivateHostedZone {
	if in == nil {
		return nil
	}
	out := new(Route53PrivateHostedZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroup) DeepCopyInto(out *SecurityGroup) {
	*out = *in
	if in.IngressRules != nil {
		in, out := &in.IngressRules, &out.IngressRules
		*out = make([]SecurityGroupRule, len(*in))
		copy(*out, *in)
	}
	if in.EgressRules != nil {
		in, out := &in.EgressRules, &out.EgressRules
		*out = make([]SecurityGroupRule, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroup.
func (in *SecurityGroup) DeepCopy() *SecurityGroup {
	if in == nil {
		return nil
	}
	out := new(SecurityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupRule) DeepCopyInto(out *SecurityGroupRule) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupRule.
func (in *SecurityGroupRule) DeepCopy() *SecurityGroupRule {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vpc) DeepCopyInto(out *Vpc) {
	*out = *in
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vpc.
func (in *Vpc) DeepCopy() *Vpc {
	if in == nil {
		return nil
	}
	out := new(Vpc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpoint) DeepCopyInto(out *VpcEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpoint.
func (in *VpcEndpoint) DeepCopy() *VpcEndpoint {
	if in == nil {
		return nil
	}
	out := new(VpcEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointList) DeepCopyInto(out *VpcEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpcEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointList.
func (in *VpcEndpointList) DeepCopy() *VpcEndpointList {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointSpec) DeepCopyInto(out *VpcEndpointSpec) {
	*out = *in
	in.SecurityGroup.DeepCopyInto(&out.SecurityGroup)
	in.Vpc.DeepCopyInto(&out.Vpc)
	out.CustomDns = in.CustomDns
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointSpec.
func (in *VpcEndpointSpec) DeepCopy() *VpcEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointStatus) DeepCopyInto(out *VpcEndpointStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointStatus.
func (in *VpcEndpointStatus) DeepCopy() *VpcEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointStatus)
	in.DeepCopyInto(out)
	return out
}