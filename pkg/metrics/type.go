package metrics

import (
	"regexp"
	"strconv"
)

type Metrics struct {
	Kind       string   `json:"kind"`
	APIVersion string   `json:"apiVersion"`
	Metadata   Metadata `json:"metadata"`
	Timestamp  string   `json:"timestamp"`
	Window     string   `json:"window"`
	Usage      Usage    `json:"usage"`
}

type Metadata struct {
	Name              string `json:"name"`
	SelfLink          string `json:"selfLink"`
	CreationTimestamp string `json:"creationTimestamp"`
}

type Usage struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}

// GetValue get cpu: ms, memory: kb, err
func (c Usage) GetValue() (int64, int64, error) {
	re := regexp.MustCompile(`([0-9]+)([A-Za-z]+)`)
	var str = re.FindStringSubmatch(c.CPU)
	cpuValue, err := strconv.ParseInt(str[1], 10, 64)
	if err != nil {
		panic(err)
	}
	cpuUnit := str[2]
	if cpuUnit == `n` {
		cpuValue /= 1000000
	}

	str = re.FindStringSubmatch(c.Memory)
	memValue, err := strconv.ParseInt(str[1], 10, 64)
	if err != nil {
		panic(err)
	}

	return cpuValue, memValue, nil
}
