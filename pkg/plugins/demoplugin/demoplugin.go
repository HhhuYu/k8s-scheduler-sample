package demoplugin

import (
	"sync"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

const (
	// Name plugins name
	Name = "demoplugin"
)

// Args config args
type Args struct {
	KubeConfig string `json:"kubeconfig,omitempty"`
	Master     string `json:"master,omitempty"`
}

// DemoPlugin object
type DemoPlugin struct {
	args   *Args
	handle framework.FrameworkHandle

	numRuns int
	dpState map[int]string
	mu      sync.RWMutex
}

var (
	_ framework.FilterPlugin  = &DemoPlugin{}
	_ framework.PreBindPlugin = &DemoPlugin{}
	_ framework.ReservePlugin = &DemoPlugin{}
)

// Name plugins name
func (dp *DemoPlugin) Name() string {
	return Name
}

// Reserve Reserve extenders
func (dp *DemoPlugin) Reserve(pc *framework.PluginContext, pod *v1.Pod, nodeName string) *framework.Status {
	if pod == nil {
		return framework.NewStatus(framework.Error, "pod cannot be nil")
	}

	if pod.Name == "pod-test" {
		pc.Lock()
		pc.Write(framework.ContextKey(pod.Name), "never bind")
		pc.Unlock()

	}

	return nil
}

// Filter plugins interface
func (dp *DemoPlugin) Filter(pc *framework.PluginContext, pod *v1.Pod, nodeName string) *framework.Status {
	klog.V(3).Infof("filter pod: %v", pod.Name)
	if pod.Name == "pod-test" {
		return framework.NewStatus(framework.Error, "the pod-test cannot pass")
	}
	return framework.NewStatus(framework.Success, "")
}

// PreBind plugins interface
func (dp *DemoPlugin) PreBind(pc *framework.PluginContext, pod *v1.Pod, nodeName string) *framework.Status {

	dp.mu.Lock()
	defer dp.mu.Unlock()

	if pod == nil {
		return framework.NewStatus(framework.Error, "pod must not be nil")
	}
	return framework.NewStatus(framework.Success, "")
}

// New plugins constructor
func New(plArgs *runtime.Unknown, handle framework.FrameworkHandle) (framework.Plugin, error) {
	args := &Args{}
	if err := framework.DecodeInto(plArgs, args); err != nil {
		return nil, err
	}
	klog.V(3).Infof("####-> args: <-#### %+v", args)
	return &DemoPlugin{
		args:   args,
		handle: handle,
	}, nil
}
