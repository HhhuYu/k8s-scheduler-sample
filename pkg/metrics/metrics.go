package metrics

import (
	"encoding/json"

	"github.com/HhhuYu/schedule-framework/pkg/utils"
)

const (
	metricsAPI = "apis/metrics.k8s.io/v1beta1/nodes"
)

func parseMetricsInfo(metricsInfoSource string) (Metrics, error) {
	var metricsInfo = Metrics{}

	json.Unmarshal([]byte(metricsInfoSource), &metricsInfo)

	return metricsInfo, nil
}

// GetMetricsInfo get metics info
func GetMetricsInfo(nodeName string) (Metrics, error) {
	metricsInfoSource, err := utils.GetInfo(metricsAPI + `/` + nodeName)
	if err != nil {
		return Metrics{}, err
	}

	metricsInfo, err := parseMetricsInfo(metricsInfoSource)
	if err != nil {
		return Metrics{}, err
	}

	return metricsInfo, nil
}
