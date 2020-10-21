package nodelabel

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

const (
	// Name plugin name
	Name = "nodelabel"
	// Error plugin error
	Error = "node didn't have the requested Labels"
)

// Args plugins args
type Args struct {
	KubeConfig       string   `json:"kubeconfig,omitempty"`
	Master           string   `json:"master,omitempty"`
	LabelsPreference []string `json:"labelspreference,omitempty"`
	LabelsAvoid      []string `json:"labelsavoid,omitempty"`
}

// NodeLabel plugin struct
type NodeLabel struct {
	handle framework.FrameworkHandle
	args   *Args
}

var _ framework.FilterPlugin = &NodeLabel{}
var _ framework.ScorePlugin = &NodeLabel{}

// Name plugins name
func (nl *NodeLabel) Name() string {
	return Name
}

// Filter filter extenter
func (nl *NodeLabel) Filter(pc *framework.PluginContext, pod *v1.Pod, nodeName string) *framework.Status {
	node := nl.handle.NodeInfoSnapshot().NodeInfoMap[nodeName]
	if node == nil {
		return framework.NewStatus(framework.Error, "node not found")
	}
	nodeLabels := labels.Set(node.Node().Labels)
	check := func(labels []string, prefer bool) bool {
		for _, label := range labels {
			exists := nodeLabels.Has(label)
			if (exists && !prefer) || (!exists && prefer) {
				return false
			}
		}
		return true
	}
	if check(nl.args.LabelsPreference, true) && check(nl.args.LabelsAvoid, false) {
		return nil
	}

	return framework.NewStatus(framework.UnschedulableAndUnresolvable, Error)
}

// Score plugin score
func (nl *NodeLabel) Score(pc *framework.PluginContext, p *v1.Pod, nodeName string) (int, *framework.Status) {
	nodeInfo := nl.handle.NodeInfoSnapshot().NodeInfoMap[nodeName]
	if nodeInfo.Node() == nil {
		return 0, framework.NewStatus(framework.Error, fmt.Sprintf("getting node %q , node is nil: %v", nodeName, nodeInfo.Node() == nil))
	}

	node := nodeInfo.Node()
	score := int(0)
	for _, label := range nl.args.LabelsPreference {
		if labels.Set(node.Labels).Has(label) {
			score += framework.MaxNodeScore
		}
	}

	score /= int(len(nl.args.LabelsPreference))
	klog.V(3).Infof("pod %+v, node %+v, score: %+v", p.Name, nodeName, score)
	return score, nil
}

// New plugins constructor
func New(plArgs *runtime.Unknown, handle framework.FrameworkHandle) (framework.Plugin, error) {
	args := &Args{}
	if err := framework.DecodeInto(plArgs, args); err != nil {
		return nil, err
	}
	klog.V(1).Infof("node label initing......")
	klog.V(3).Infof("####-> args: <-#### %+v", args)
	return &NodeLabel{
		args:   args,
		handle: handle,
	}, nil
}
