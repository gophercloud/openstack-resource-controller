//go:build !ignore_autogenerated

/*
Copyright 2023.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackCloud) DeepCopyInto(out *OpenStackCloud) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackCloud.
func (in *OpenStackCloud) DeepCopy() *OpenStackCloud {
	if in == nil {
		return nil
	}
	out := new(OpenStackCloud)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackCloud) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackCloudCredentials) DeepCopyInto(out *OpenStackCloudCredentials) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackCloudCredentials.
func (in *OpenStackCloudCredentials) DeepCopy() *OpenStackCloudCredentials {
	if in == nil {
		return nil
	}
	out := new(OpenStackCloudCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackCloudCredentialsSecretRef) DeepCopyInto(out *OpenStackCloudCredentialsSecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackCloudCredentialsSecretRef.
func (in *OpenStackCloudCredentialsSecretRef) DeepCopy() *OpenStackCloudCredentialsSecretRef {
	if in == nil {
		return nil
	}
	out := new(OpenStackCloudCredentialsSecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackCloudList) DeepCopyInto(out *OpenStackCloudList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackCloud, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackCloudList.
func (in *OpenStackCloudList) DeepCopy() *OpenStackCloudList {
	if in == nil {
		return nil
	}
	out := new(OpenStackCloudList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackCloudList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackCloudSpec) DeepCopyInto(out *OpenStackCloudSpec) {
	*out = *in
	out.Credentials = in.Credentials
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackCloudSpec.
func (in *OpenStackCloudSpec) DeepCopy() *OpenStackCloudSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackCloudSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackCloudStatus) DeepCopyInto(out *OpenStackCloudStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackCloudStatus.
func (in *OpenStackCloudStatus) DeepCopy() *OpenStackCloudStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackCloudStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackFlavor) DeepCopyInto(out *OpenStackFlavor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackFlavor.
func (in *OpenStackFlavor) DeepCopy() *OpenStackFlavor {
	if in == nil {
		return nil
	}
	out := new(OpenStackFlavor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackFlavor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackFlavorList) DeepCopyInto(out *OpenStackFlavorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackFlavor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackFlavorList.
func (in *OpenStackFlavorList) DeepCopy() *OpenStackFlavorList {
	if in == nil {
		return nil
	}
	out := new(OpenStackFlavorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackFlavorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackFlavorSpec) DeepCopyInto(out *OpenStackFlavorSpec) {
	*out = *in
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		*out = new(int)
		**out = **in
	}
	if in.Swap != nil {
		in, out := &in.Swap, &out.Swap
		*out = new(int)
		**out = **in
	}
	if in.IsPublic != nil {
		in, out := &in.IsPublic, &out.IsPublic
		*out = new(bool)
		**out = **in
	}
	if in.Ephemeral != nil {
		in, out := &in.Ephemeral, &out.Ephemeral
		*out = new(int)
		**out = **in
	}
	if in.Unmanaged != nil {
		in, out := &in.Unmanaged, &out.Unmanaged
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackFlavorSpec.
func (in *OpenStackFlavorSpec) DeepCopy() *OpenStackFlavorSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackFlavorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackFlavorStatus) DeepCopyInto(out *OpenStackFlavorStatus) {
	*out = *in
	if in.ExtraSpecs != nil {
		in, out := &in.ExtraSpecs, &out.ExtraSpecs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackFlavorStatus.
func (in *OpenStackFlavorStatus) DeepCopy() *OpenStackFlavorStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackFlavorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackFloatingIP) DeepCopyInto(out *OpenStackFloatingIP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackFloatingIP.
func (in *OpenStackFloatingIP) DeepCopy() *OpenStackFloatingIP {
	if in == nil {
		return nil
	}
	out := new(OpenStackFloatingIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackFloatingIP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackFloatingIPList) DeepCopyInto(out *OpenStackFloatingIPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackFloatingIP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackFloatingIPList.
func (in *OpenStackFloatingIPList) DeepCopy() *OpenStackFloatingIPList {
	if in == nil {
		return nil
	}
	out := new(OpenStackFloatingIPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackFloatingIPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackFloatingIPSpec) DeepCopyInto(out *OpenStackFloatingIPSpec) {
	*out = *in
	if in.Unmanaged != nil {
		in, out := &in.Unmanaged, &out.Unmanaged
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackFloatingIPSpec.
func (in *OpenStackFloatingIPSpec) DeepCopy() *OpenStackFloatingIPSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackFloatingIPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackFloatingIPStatus) DeepCopyInto(out *OpenStackFloatingIPStatus) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackFloatingIPStatus.
func (in *OpenStackFloatingIPStatus) DeepCopy() *OpenStackFloatingIPStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackFloatingIPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackImage) DeepCopyInto(out *OpenStackImage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackImage.
func (in *OpenStackImage) DeepCopy() *OpenStackImage {
	if in == nil {
		return nil
	}
	out := new(OpenStackImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackImage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackImageList) DeepCopyInto(out *OpenStackImageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackImageList.
func (in *OpenStackImageList) DeepCopy() *OpenStackImageList {
	if in == nil {
		return nil
	}
	out := new(OpenStackImageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackImageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackImageSpec) DeepCopyInto(out *OpenStackImageSpec) {
	*out = *in
	if in.Protected != nil {
		in, out := &in.Protected, &out.Protected
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(string)
		**out = **in
	}
	if in.Unmanaged != nil {
		in, out := &in.Unmanaged, &out.Unmanaged
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackImageSpec.
func (in *OpenStackImageSpec) DeepCopy() *OpenStackImageSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackImageStatus) DeepCopyInto(out *OpenStackImageStatus) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ImportMethods != nil {
		in, out := &in.ImportMethods, &out.ImportMethods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StoreIDs != nil {
		in, out := &in.StoreIDs, &out.StoreIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackImageStatus.
func (in *OpenStackImageStatus) DeepCopy() *OpenStackImageStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackImageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackNetwork) DeepCopyInto(out *OpenStackNetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackNetwork.
func (in *OpenStackNetwork) DeepCopy() *OpenStackNetwork {
	if in == nil {
		return nil
	}
	out := new(OpenStackNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackNetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackNetworkList) DeepCopyInto(out *OpenStackNetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackNetworkList.
func (in *OpenStackNetworkList) DeepCopy() *OpenStackNetworkList {
	if in == nil {
		return nil
	}
	out := new(OpenStackNetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackNetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackNetworkSpec) DeepCopyInto(out *OpenStackNetworkSpec) {
	*out = *in
	if in.Unmanaged != nil {
		in, out := &in.Unmanaged, &out.Unmanaged
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackNetworkSpec.
func (in *OpenStackNetworkSpec) DeepCopy() *OpenStackNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackNetworkStatus) DeepCopyInto(out *OpenStackNetworkStatus) {
	*out = *in
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AvailabilityZoneHints != nil {
		in, out := &in.AvailabilityZoneHints, &out.AvailabilityZoneHints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackNetworkStatus.
func (in *OpenStackNetworkStatus) DeepCopy() *OpenStackNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackNetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackSecurityGroup) DeepCopyInto(out *OpenStackSecurityGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackSecurityGroup.
func (in *OpenStackSecurityGroup) DeepCopy() *OpenStackSecurityGroup {
	if in == nil {
		return nil
	}
	out := new(OpenStackSecurityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackSecurityGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackSecurityGroupList) DeepCopyInto(out *OpenStackSecurityGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackSecurityGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackSecurityGroupList.
func (in *OpenStackSecurityGroupList) DeepCopy() *OpenStackSecurityGroupList {
	if in == nil {
		return nil
	}
	out := new(OpenStackSecurityGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackSecurityGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackSecurityGroupRule) DeepCopyInto(out *OpenStackSecurityGroupRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackSecurityGroupRule.
func (in *OpenStackSecurityGroupRule) DeepCopy() *OpenStackSecurityGroupRule {
	if in == nil {
		return nil
	}
	out := new(OpenStackSecurityGroupRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackSecurityGroupRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackSecurityGroupRuleList) DeepCopyInto(out *OpenStackSecurityGroupRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackSecurityGroupRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackSecurityGroupRuleList.
func (in *OpenStackSecurityGroupRuleList) DeepCopy() *OpenStackSecurityGroupRuleList {
	if in == nil {
		return nil
	}
	out := new(OpenStackSecurityGroupRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackSecurityGroupRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackSecurityGroupRuleSpec) DeepCopyInto(out *OpenStackSecurityGroupRuleSpec) {
	*out = *in
	if in.Unmanaged != nil {
		in, out := &in.Unmanaged, &out.Unmanaged
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackSecurityGroupRuleSpec.
func (in *OpenStackSecurityGroupRuleSpec) DeepCopy() *OpenStackSecurityGroupRuleSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackSecurityGroupRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackSecurityGroupRuleStatus) DeepCopyInto(out *OpenStackSecurityGroupRuleStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackSecurityGroupRuleStatus.
func (in *OpenStackSecurityGroupRuleStatus) DeepCopy() *OpenStackSecurityGroupRuleStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackSecurityGroupRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackSecurityGroupSpec) DeepCopyInto(out *OpenStackSecurityGroupSpec) {
	*out = *in
	if in.Unmanaged != nil {
		in, out := &in.Unmanaged, &out.Unmanaged
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackSecurityGroupSpec.
func (in *OpenStackSecurityGroupSpec) DeepCopy() *OpenStackSecurityGroupSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackSecurityGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackSecurityGroupStatus) DeepCopyInto(out *OpenStackSecurityGroupStatus) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackSecurityGroupStatus.
func (in *OpenStackSecurityGroupStatus) DeepCopy() *OpenStackSecurityGroupStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackSecurityGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackServer) DeepCopyInto(out *OpenStackServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackServer.
func (in *OpenStackServer) DeepCopy() *OpenStackServer {
	if in == nil {
		return nil
	}
	out := new(OpenStackServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackServerList) DeepCopyInto(out *OpenStackServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackServerList.
func (in *OpenStackServerList) DeepCopy() *OpenStackServerList {
	if in == nil {
		return nil
	}
	out := new(OpenStackServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackServerSpec) DeepCopyInto(out *OpenStackServerSpec) {
	*out = *in
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Unmanaged != nil {
		in, out := &in.Unmanaged, &out.Unmanaged
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackServerSpec.
func (in *OpenStackServerSpec) DeepCopy() *OpenStackServerSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackServerStatus) DeepCopyInto(out *OpenStackServerStatus) {
	*out = *in
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AttachedVolumeIDs != nil {
		in, out := &in.AttachedVolumeIDs, &out.AttachedVolumeIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServerGroupIDs != nil {
		in, out := &in.ServerGroupIDs, &out.ServerGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackServerStatus.
func (in *OpenStackServerStatus) DeepCopy() *OpenStackServerStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackSubnet) DeepCopyInto(out *OpenStackSubnet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackSubnet.
func (in *OpenStackSubnet) DeepCopy() *OpenStackSubnet {
	if in == nil {
		return nil
	}
	out := new(OpenStackSubnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackSubnet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackSubnetAllocationPool) DeepCopyInto(out *OpenStackSubnetAllocationPool) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackSubnetAllocationPool.
func (in *OpenStackSubnetAllocationPool) DeepCopy() *OpenStackSubnetAllocationPool {
	if in == nil {
		return nil
	}
	out := new(OpenStackSubnetAllocationPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackSubnetHostRoute) DeepCopyInto(out *OpenStackSubnetHostRoute) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackSubnetHostRoute.
func (in *OpenStackSubnetHostRoute) DeepCopy() *OpenStackSubnetHostRoute {
	if in == nil {
		return nil
	}
	out := new(OpenStackSubnetHostRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackSubnetList) DeepCopyInto(out *OpenStackSubnetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackSubnet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackSubnetList.
func (in *OpenStackSubnetList) DeepCopy() *OpenStackSubnetList {
	if in == nil {
		return nil
	}
	out := new(OpenStackSubnetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackSubnetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackSubnetSpec) DeepCopyInto(out *OpenStackSubnetSpec) {
	*out = *in
	if in.AllocationPools != nil {
		in, out := &in.AllocationPools, &out.AllocationPools
		*out = make([]OpenStackSubnetAllocationPool, len(*in))
		copy(*out, *in)
	}
	if in.GatewayIP != nil {
		in, out := &in.GatewayIP, &out.GatewayIP
		*out = new(string)
		**out = **in
	}
	if in.EnableDHCP != nil {
		in, out := &in.EnableDHCP, &out.EnableDHCP
		*out = new(bool)
		**out = **in
	}
	if in.DNSNameservers != nil {
		in, out := &in.DNSNameservers, &out.DNSNameservers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServiceTypes != nil {
		in, out := &in.ServiceTypes, &out.ServiceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HostRoutes != nil {
		in, out := &in.HostRoutes, &out.HostRoutes
		*out = make([]OpenStackSubnetHostRoute, len(*in))
		copy(*out, *in)
	}
	if in.Unmanaged != nil {
		in, out := &in.Unmanaged, &out.Unmanaged
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackSubnetSpec.
func (in *OpenStackSubnetSpec) DeepCopy() *OpenStackSubnetSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackSubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackSubnetStatus) DeepCopyInto(out *OpenStackSubnetStatus) {
	*out = *in
	if in.DNSNameservers != nil {
		in, out := &in.DNSNameservers, &out.DNSNameservers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServiceTypes != nil {
		in, out := &in.ServiceTypes, &out.ServiceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllocationPools != nil {
		in, out := &in.AllocationPools, &out.AllocationPools
		*out = make([]OpenStackSubnetAllocationPool, len(*in))
		copy(*out, *in)
	}
	if in.HostRoutes != nil {
		in, out := &in.HostRoutes, &out.HostRoutes
		*out = make([]OpenStackSubnetHostRoute, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackSubnetStatus.
func (in *OpenStackSubnetStatus) DeepCopy() *OpenStackSubnetStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackSubnetStatus)
	in.DeepCopyInto(out)
	return out
}
