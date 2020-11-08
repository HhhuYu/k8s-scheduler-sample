package metrics

import "github.com/HhhuYu/schedule-framework/pkg/utils"

// MetricsInfoMap node metrics info map
type MetricsInfoMap map[string]Resource

// NodeInfoMap node allocatable info map
type NodeInfoMap map[string]Resource

// NodeScoreMap Node score map
type NodeScoreMap map[string]float64

// Resource node resource include cpu & memory
type Resource struct {
	cpu    uint64
	memory uint64
}

const (
	metricsAPI = "apis/metrics.k8s.io/v1beta1/nodes"
	nodeAPI    = "api/v1/nodes"
)

func parseMetricsInfo(metricsInfoSource string) (MetricsInfoMap, error) {

	return MetricsInfoMap{}, nil
}

// GetMetricsInfo get metics info
func GetMetricsInfo() (string, error) {
	metricsInfoSource, err := utils.GetInfo(metricsAPI)
	if err != nil {
		return "", err
	}

	return metricsInfoSource, nil
}

// GetNodeInfo get metics info
func GetNodeInfo() (string, error) {
	nodeInfoSource, err := utils.GetInfo(nodeAPI)
	if err != nil {
		return "", err
	}

	return nodeInfoSource, nil
}
