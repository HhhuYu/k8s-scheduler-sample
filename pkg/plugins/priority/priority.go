package priority

import (
	"github.com/HhhuYu/schedule-framework/pkg/metrics"
	"github.com/HhhuYu/schedule-framework/pkg/node"
)

// CalculateFraction calcilate the fraction
func CalculateFraction(nodeInfo *node.Node, metricsInfo *metrics.Metrics) (int64, error) {
	cpuM, meroM, err := metricsInfo.Usage.GetValue()
	if err != nil {
		return 0, err
	}

	cpuN, meroN, err := nodeInfo.Status.Allocatable.GetValue()
	if err != nil {
		return 0, err
	}

	score := (1-cpuM*1.0/cpuN)*60 + (1-meroM*1.0/meroN)*40
	return score, nil
}
