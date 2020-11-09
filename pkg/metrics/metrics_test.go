package metrics

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseMetricsInfo(t *testing.T) {
	data, err := ioutil.ReadFile("metrics.json")
	if err != nil {
		return
	}
	metricsInfo, err := parseMetricsInfo(string(data))

	fmt.Printf("metrics info: %v", metricsInfo)

}

func TestGetValue(t *testing.T) {
	data, err := ioutil.ReadFile("metrics.json")
	if err != nil {
		return
	}
	metricsInfo, err := parseMetricsInfo(string(data))

	cpu, mem, _ := metricsInfo.Usage.GetValue()
	fmt.Printf("metrics info: %v", metricsInfo)

	fmt.Println("cpu value:", cpu)
	fmt.Println("mem value:", mem)

}
