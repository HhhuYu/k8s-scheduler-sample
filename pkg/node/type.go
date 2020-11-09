package node

import (
	"regexp"
	"strconv"
)

type Node struct {
	Kind       string   `json:"kind"`
	APIVersion string   `json:"apiVersion"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
	Status     Status   `json:"status"`
}

type Metadata struct {
	Name              string            `json:"name"`
	SelfLink          string            `json:"selfLink"`
	Uid               string            `json:"uid"`
	ResourceVersion   string            `json:"resourceVersion"`
	CreationTimestamp string            `json:"creationTimestamp"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
}

type Spec struct {
	PodCIDR string `json:"podCIDR"`
}

type Status struct {
	Capacity        Allocatable     `json:"capacity"`
	Allocatable     Allocatable     `json:"allocatable"`
	Conditions      []Condition     `json:"conditions"`
	Addresses       []Address       `json:"addresses"`
	DaemonEndpoints DaemonEndpoints `json:"daemonEndpoints"`
	NodeInfo        NodeInfo        `json:"nodeInfo"`
	Images          []Image         `json:"images"`
}

type Address struct {
	Type    string `json:"type"`
	Address string `json:"address"`
}

type Allocatable struct {
	CPU              string `json:"cpu"`
	EphemeralStorage string `json:"ephemeral-storage"`
	Hugepages1Gi     string `json:"hugepages-1Gi"`
	Hugepages2Mi     string `json:"hugepages-2Mi"`
	Memory           string `json:"memory"`
	Pods             string `json:"pods"`
}

type Condition struct {
	Type               string `json:"type"`
	Status             string `json:"status"`
	LastHeartbeatTime  string `json:"lastHeartbeatTime"`
	LastTransitionTime string `json:"lastTransitionTime"`
	Reason             string `json:"reason"`
	Message            string `json:"message"`
}

type DaemonEndpoints struct {
	KubeletEndpoint KubeletEndpoint `json:"kubeletEndpoint"`
}

type KubeletEndpoint struct {
	Port int64 `json:"Port"`
}

type Image struct {
	Names     []string `json:"names"`
	SizeBytes int64    `json:"sizeBytes"`
}

type NodeInfo struct {
	MachineID               string `json:"machineID"`
	SystemUUID              string `json:"systemUUID"`
	BootID                  string `json:"bootID"`
	KernelVersion           string `json:"kernelVersion"`
	OSImage                 string `json:"osImage"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
	KubeletVersion          string `json:"kubeletVersion"`
	KubeProxyVersion        string `json:"kubeProxyVersion"`
	OperatingSystem         string `json:"operatingSystem"`
	Architecture            string `json:"architecture"`
}

// GetValue get cpu: ms, memory: kb, err,
func (a Allocatable) GetValue() (int64, int64, error) {
	re := regexp.MustCompile(`([0-9]+)([A-Za-z]+)`)
	var str = re.FindStringSubmatch(a.CPU)
	cpuValue, err := strconv.ParseInt(a.CPU, 10, 64)

	if err != nil {
		panic(err)
	}
	cpuValue *= 1000

	str = re.FindStringSubmatch(a.Memory)
	memValue, err := strconv.ParseInt(str[1], 10, 64)
	if err != nil {
		panic(err)
	}

	return cpuValue, memValue, nil
}
