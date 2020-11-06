package demoplugin

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

const (
	// Name plugins name
	Name = "demoplugin"
	// metricsServerURL = "https://kubernetes/apis/metrics.k8s.io/v1beta1/nodes"

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

func getURL() string {
	host := os.Getenv("KUBERNETES_SERVICE_HOST")
	port := os.Getenv("KUBERRETES_SERVICE_PORT")
	url := fmt.Sprintf("https://%s:%s/apis/metrics.k8s.io/v1beta1/nodes", host, port)
	// url := "https://" + HOST + ":" + PORT + "/apis/metrics.k8s.io/v1beta1/nodes"
	return url
}

func loadCA(caFile string) *x509.CertPool {
	pool := x509.NewCertPool()

	if ca, e := ioutil.ReadFile(caFile); e != nil {
		log.Fatal("ReadFile: ", e)
	} else {
		pool.AppendCertsFromPEM(ca)
	}
	return pool
}

// Get Http get
func Get(url string) string {

	client := &http.Client{
		// 超时时间：5秒
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: loadCA("/var/run/secrets/kubernetes.io/serviceaccount/ca.crt")},
		}}
	req, err := http.NewRequest("GET", url, nil)
	token, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	req.Header.Add("Authorization", `Bearer `+string(token))

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
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

	klog.V(3).Infof("Write pod ", pod.Name, " in ContextKey")

	return nil
}

// Filter plugins interface
func (dp *DemoPlugin) Filter(pc *framework.PluginContext, pod *v1.Pod, nodeName string) *framework.Status {
	klog.V(3).Infof("filter pod: %v", pod.Name)

	metricsInfo := Get(getURL())
	klog.V(3).Infof("metrics-server-info: %v", metricsInfo)

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
