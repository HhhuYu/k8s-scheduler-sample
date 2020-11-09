package node

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestNodeParse(t *testing.T) {

}

func TestGetValue(t *testing.T) {
	data, err := ioutil.ReadFile("node.json")
	if err != nil {
		return
	}
	nodeInfo, err := parseNodeInfo(string(data))

	cpu, mem, _ := nodeInfo.Status.Allocatable.GetValue()
	fmt.Printf("metrics info: %v", nodeInfo)

	fmt.Println("cpu value:", cpu)
	fmt.Println("mem value:", mem)
}
