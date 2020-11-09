package node

import (
	"encoding/json"

	"github.com/HhhuYu/schedule-framework/pkg/utils"
)

const (
	nodeAPI = "api/v1/nodes"
)

func parseNodeInfo(nodeInfoSource string) (Node, error) {
	var nodeInfo = Node{}

	json.Unmarshal([]byte(nodeInfoSource), &nodeInfo)

	return nodeInfo, nil
}

// GetNodeInfo get metics info
func GetNodeInfo(nodeName string) (Node, error) {
	nodeInfoSource, err := utils.GetInfo(nodeAPI + `/` + nodeName)
	if err != nil {
		return Node{}, err
	}
	nodeInfo, err := parseNodeInfo(nodeInfoSource)
	if err != nil {
		return Node{}, err
	}

	return nodeInfo, nil
}

// GetMetricsInfo get metics info
