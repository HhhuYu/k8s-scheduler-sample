package scoreplugin

import (
	"github.com/HhhuYu/schedule-framework/pkg/metrics"
	"github.com/HhhuYu/schedule-framework/pkg/node"
	"github.com/HhhuYu/schedule-framework/pkg/plugins/priority"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

const (
	// Name score plugin name
	Name    = "moscoreplugin"
	version = "0.0.1"
)

type Args struct {
	KubeConfig string `json:"kubeconfig,omitempty"`
	Master     string `json:"master,omitempty"`
}

// ScorePlugin the demo plugin of score
type ScorePlugin struct {
	args   *Args
	handle framework.FrameworkHandle
}

var (
	_ framework.ScorePlugin = &ScorePlugin{}
)

// Name return name of plugin
func (sp *ScorePlugin) Name() string {
	return Name
}

// Score plugins interface
func (sp *ScorePlugin) Score(pc *framework.PluginContext, p *v1.Pod, nodeName string) (int, *framework.Status) {
	metricsInfo, err := metrics.GetMetricsInfo(nodeName)
	if err != nil {
		return 0, framework.NewStatus(framework.Error, "get metrics info error")
	}
	nodeInfo, err := node.GetNodeInfo(nodeName)
	if err != nil {
		return 0, framework.NewStatus(framework.Error, "get node info error")
	}
	score, err := priority.CalculateFraction(&nodeInfo, &metricsInfo)
	if err != nil {
		return 0, framework.NewStatus(framework.Error, err.Error())
	}
	klog.V(3).Infof("%s: score:%v", nodeName, score)
	return int(score), framework.NewStatus(framework.Success, "")
}

// New plugins constructor
func New(plArgs *runtime.Unknown, handle framework.FrameworkHandle) (framework.Plugin, error) {
	args := &Args{}
	if err := framework.DecodeInto(plArgs, args); err != nil {
		return nil, err
	}
	klog.V(3).Infof("Args: %+v", args)

	klog.V(3).Infof("Scoring plugin version: %s", version)
	return &ScorePlugin{
		args:   args,
		handle: handle,
	}, nil
}
