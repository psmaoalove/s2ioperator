// +build !ignore_autogenerated

/*
Copyright 2019 The Kubesphere Authors.

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfig) DeepCopyInto(out *AuthConfig) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfig.
func (in *AuthConfig) DeepCopy() *AuthConfig {
	if in == nil {
		return nil
	}
	out := new(AuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CGroupLimits) DeepCopyInto(out *CGroupLimits) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CGroupLimits.
func (in *CGroupLimits) DeepCopy() *CGroupLimits {
	if in == nil {
		return nil
	}
	out := new(CGroupLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerConfig) DeepCopyInto(out *ContainerConfig) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerConfig.
func (in *ContainerConfig) DeepCopy() *ContainerConfig {
	if in == nil {
		return nil
	}
	out := new(ContainerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerInfo) DeepCopyInto(out *ContainerInfo) {
	*out = *in
	if in.RuntimeArtifacts != nil {
		in, out := &in.RuntimeArtifacts, &out.RuntimeArtifacts
		*out = make([]VolumeSpec, len(*in))
		copy(*out, *in)
	}
	if in.BuildVolumes != nil {
		in, out := &in.BuildVolumes, &out.BuildVolumes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerInfo.
func (in *ContainerInfo) DeepCopy() *ContainerInfo {
	if in == nil {
		return nil
	}
	out := new(ContainerInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerConfig) DeepCopyInto(out *DockerConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerConfig.
func (in *DockerConfig) DeepCopy() *DockerConfig {
	if in == nil {
		return nil
	}
	out := new(DockerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerConfigEntry) DeepCopyInto(out *DockerConfigEntry) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerConfigEntry.
func (in *DockerConfigEntry) DeepCopy() *DockerConfigEntry {
	if in == nil {
		return nil
	}
	out := new(DockerConfigEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerConfigJson) DeepCopyInto(out *DockerConfigJson) {
	*out = *in
	if in.Auths != nil {
		in, out := &in.Auths, &out.Auths
		*out = make(DockerConfigMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerConfigJson.
func (in *DockerConfigJson) DeepCopy() *DockerConfigJson {
	if in == nil {
		return nil
	}
	out := new(DockerConfigJson)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in DockerConfigMap) DeepCopyInto(out *DockerConfigMap) {
	{
		in := &in
		*out = make(DockerConfigMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerConfigMap.
func (in DockerConfigMap) DeepCopy() DockerConfigMap {
	if in == nil {
		return nil
	}
	out := new(DockerConfigMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpec) DeepCopyInto(out *EnvironmentSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpec.
func (in *EnvironmentSpec) DeepCopy() *EnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parameter) DeepCopyInto(out *Parameter) {
	*out = *in
	if in.OptValues != nil {
		in, out := &in.OptValues, &out.OptValues
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parameter.
func (in *Parameter) DeepCopy() *Parameter {
	if in == nil {
		return nil
	}
	out := new(Parameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyConfig) DeepCopyInto(out *ProxyConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyConfig.
func (in *ProxyConfig) DeepCopy() *ProxyConfig {
	if in == nil {
		return nil
	}
	out := new(ProxyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iAutoScale) DeepCopyInto(out *S2iAutoScale) {
	*out = *in
	if in.InitReplicas != nil {
		in, out := &in.InitReplicas, &out.InitReplicas
		*out = new(int32)
		**out = **in
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iAutoScale.
func (in *S2iAutoScale) DeepCopy() *S2iAutoScale {
	if in == nil {
		return nil
	}
	out := new(S2iAutoScale)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iBuildResult) DeepCopyInto(out *S2iBuildResult) {
	*out = *in
	if in.ImageRepoTags != nil {
		in, out := &in.ImageRepoTags, &out.ImageRepoTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iBuildResult.
func (in *S2iBuildResult) DeepCopy() *S2iBuildResult {
	if in == nil {
		return nil
	}
	out := new(S2iBuildResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iBuildSource) DeepCopyInto(out *S2iBuildSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iBuildSource.
func (in *S2iBuildSource) DeepCopy() *S2iBuildSource {
	if in == nil {
		return nil
	}
	out := new(S2iBuildSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iBuilder) DeepCopyInto(out *S2iBuilder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iBuilder.
func (in *S2iBuilder) DeepCopy() *S2iBuilder {
	if in == nil {
		return nil
	}
	out := new(S2iBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S2iBuilder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iBuilderList) DeepCopyInto(out *S2iBuilderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S2iBuilder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iBuilderList.
func (in *S2iBuilderList) DeepCopy() *S2iBuilderList {
	if in == nil {
		return nil
	}
	out := new(S2iBuilderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S2iBuilderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iBuilderSpec) DeepCopyInto(out *S2iBuilderSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(S2iConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.FromTemplate != nil {
		in, out := &in.FromTemplate, &out.FromTemplate
		*out = new(UserDefineTemplate)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iBuilderSpec.
func (in *S2iBuilderSpec) DeepCopy() *S2iBuilderSpec {
	if in == nil {
		return nil
	}
	out := new(S2iBuilderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iBuilderStatus) DeepCopyInto(out *S2iBuilderStatus) {
	*out = *in
	if in.LastRunName != nil {
		in, out := &in.LastRunName, &out.LastRunName
		*out = new(string)
		**out = **in
	}
	if in.LastRunStartTime != nil {
		in, out := &in.LastRunStartTime, &out.LastRunStartTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iBuilderStatus.
func (in *S2iBuilderStatus) DeepCopy() *S2iBuilderStatus {
	if in == nil {
		return nil
	}
	out := new(S2iBuilderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iBuilderTemplate) DeepCopyInto(out *S2iBuilderTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iBuilderTemplate.
func (in *S2iBuilderTemplate) DeepCopy() *S2iBuilderTemplate {
	if in == nil {
		return nil
	}
	out := new(S2iBuilderTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S2iBuilderTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iBuilderTemplateList) DeepCopyInto(out *S2iBuilderTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S2iBuilderTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iBuilderTemplateList.
func (in *S2iBuilderTemplateList) DeepCopy() *S2iBuilderTemplateList {
	if in == nil {
		return nil
	}
	out := new(S2iBuilderTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S2iBuilderTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iBuilderTemplateSpec) DeepCopyInto(out *S2iBuilderTemplateSpec) {
	*out = *in
	if in.ContainerInfo != nil {
		in, out := &in.ContainerInfo, &out.ContainerInfo
		*out = make([]ContainerInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]Parameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iBuilderTemplateSpec.
func (in *S2iBuilderTemplateSpec) DeepCopy() *S2iBuilderTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(S2iBuilderTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iBuilderTemplateStatus) DeepCopyInto(out *S2iBuilderTemplateStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iBuilderTemplateStatus.
func (in *S2iBuilderTemplateStatus) DeepCopy() *S2iBuilderTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(S2iBuilderTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iConfig) DeepCopyInto(out *S2iConfig) {
	*out = *in
	if in.RuntimeAuthentication != nil {
		in, out := &in.RuntimeAuthentication, &out.RuntimeAuthentication
		*out = new(AuthConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.RuntimeArtifacts != nil {
		in, out := &in.RuntimeArtifacts, &out.RuntimeArtifacts
		*out = make([]VolumeSpec, len(*in))
		copy(*out, *in)
	}
	if in.DockerConfig != nil {
		in, out := &in.DockerConfig, &out.DockerConfig
		*out = new(DockerConfig)
		**out = **in
	}
	if in.PullAuthentication != nil {
		in, out := &in.PullAuthentication, &out.PullAuthentication
		*out = new(AuthConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.PushAuthentication != nil {
		in, out := &in.PushAuthentication, &out.PushAuthentication
		*out = new(AuthConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.IncrementalAuthentication != nil {
		in, out := &in.IncrementalAuthentication, &out.IncrementalAuthentication
		*out = new(AuthConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = make([]EnvironmentSpec, len(*in))
		copy(*out, *in)
	}
	if in.Injections != nil {
		in, out := &in.Injections, &out.Injections
		*out = make([]VolumeSpec, len(*in))
		copy(*out, *in)
	}
	if in.CGroupLimits != nil {
		in, out := &in.CGroupLimits, &out.CGroupLimits
		*out = new(CGroupLimits)
		**out = **in
	}
	if in.DropCapabilities != nil {
		in, out := &in.DropCapabilities, &out.DropCapabilities
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ScriptDownloadProxyConfig != nil {
		in, out := &in.ScriptDownloadProxyConfig, &out.ScriptDownloadProxyConfig
		*out = new(ProxyConfig)
		**out = **in
	}
	if in.BuildVolumes != nil {
		in, out := &in.BuildVolumes, &out.BuildVolumes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityOpt != nil {
		in, out := &in.SecurityOpt, &out.SecurityOpt
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AddHost != nil {
		in, out := &in.AddHost, &out.AddHost
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.GitSecretRef != nil {
		in, out := &in.GitSecretRef, &out.GitSecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.NodeAffinityValues != nil {
		in, out := &in.NodeAffinityValues, &out.NodeAffinityValues
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iConfig.
func (in *S2iConfig) DeepCopy() *S2iConfig {
	if in == nil {
		return nil
	}
	out := new(S2iConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iRun) DeepCopyInto(out *S2iRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iRun.
func (in *S2iRun) DeepCopy() *S2iRun {
	if in == nil {
		return nil
	}
	out := new(S2iRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S2iRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iRunList) DeepCopyInto(out *S2iRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S2iRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iRunList.
func (in *S2iRunList) DeepCopy() *S2iRunList {
	if in == nil {
		return nil
	}
	out := new(S2iRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S2iRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iRunSpec) DeepCopyInto(out *S2iRunSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iRunSpec.
func (in *S2iRunSpec) DeepCopy() *S2iRunSpec {
	if in == nil {
		return nil
	}
	out := new(S2iRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iRunStatus) DeepCopyInto(out *S2iRunStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.S2iBuildResult != nil {
		in, out := &in.S2iBuildResult, &out.S2iBuildResult
		*out = new(S2iBuildResult)
		(*in).DeepCopyInto(*out)
	}
	if in.S2iBuildSource != nil {
		in, out := &in.S2iBuildSource, &out.S2iBuildSource
		*out = new(S2iBuildSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iRunStatus.
func (in *S2iRunStatus) DeepCopy() *S2iRunStatus {
	if in == nil {
		return nil
	}
	out := new(S2iRunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserDefineTemplate) DeepCopyInto(out *UserDefineTemplate) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]Parameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserDefineTemplate.
func (in *UserDefineTemplate) DeepCopy() *UserDefineTemplate {
	if in == nil {
		return nil
	}
	out := new(UserDefineTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSpec) DeepCopyInto(out *VolumeSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSpec.
func (in *VolumeSpec) DeepCopy() *VolumeSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeSpec)
	in.DeepCopyInto(out)
	return out
}
