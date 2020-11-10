package priority

import (
	"github.com/HhhuYu/schedule-framework/pkg/metrics"
	"github.com/HhhuYu/schedule-framework/pkg/node"
	"k8s.io/klog"
)

// CalculateFraction calcilate the fraction
func CalculateFraction(nodeInfo *node.Node, metricsInfo *metrics.Metrics) (int64, error) {
	cpuM, memM, err := metricsInfo.Usage.GetValue()
	if err != nil {
		return 0, err
	}

	cpuN, memN, err := nodeInfo.Status.Allocatable.GetValue()
	if err != nil {
		return 0, err
	}

	klog.V(3).Infof("Node: %s", nodeInfo.Metadata.Name)
	klog.V(3).Infof("Metrics:  cpu: %d memory: %d\n", cpuM, memM)
	klog.V(3).Infof("Node:  cpu: %d memory: %d\n", cpuN, memN)
	klog.V(3).Infof("Fraction:  cpu: %f memory: %f", float64(cpuM)/float64(cpuN)*100, float64(memM)*1.0/float64(memN)*100)
	score := (1-(float64(cpuM)/float64(cpuN)))*6 + (1-(float64(memM)/float64(memN)))*4
	return int64(score), nil
}
