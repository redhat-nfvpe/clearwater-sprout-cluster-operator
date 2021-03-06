// +build !ignore_autogenerated

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SproutCluster) DeepCopyInto(out *SproutCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SproutCluster.
func (in *SproutCluster) DeepCopy() *SproutCluster {
	if in == nil {
		return nil
	}
	out := new(SproutCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SproutCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SproutClusterList) DeepCopyInto(out *SproutClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SproutCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SproutClusterList.
func (in *SproutClusterList) DeepCopy() *SproutClusterList {
	if in == nil {
		return nil
	}
	out := new(SproutClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SproutClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SproutClusterSpec) DeepCopyInto(out *SproutClusterSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SproutClusterSpec.
func (in *SproutClusterSpec) DeepCopy() *SproutClusterSpec {
	if in == nil {
		return nil
	}
	out := new(SproutClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SproutClusterStatus) DeepCopyInto(out *SproutClusterStatus) {
	*out = *in
	if in.ShardNodes != nil {
		in, out := &in.ShardNodes, &out.ShardNodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BonoNodes != nil {
		in, out := &in.BonoNodes, &out.BonoNodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SproutClusterStatus.
func (in *SproutClusterStatus) DeepCopy() *SproutClusterStatus {
	if in == nil {
		return nil
	}
	out := new(SproutClusterStatus)
	in.DeepCopyInto(out)
	return out
}
