package utils

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
	"time"
)

const (
	token  = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	cacert = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
)

func createURL(sourceURL string) string {
	host := os.Getenv("KUBERNETES_SERVICE_HOST")
	port := os.Getenv("KUBERRETES_SERVICE_PORT")
	hostURL := fmt.Sprintf("https://%s:%s/", host, port)
	url := hostURL + sourceURL
	return url
}

func loadCA(caFile string) (*x509.CertPool, error) {
	pool := x509.NewCertPool()

	if ca, e := ioutil.ReadFile(caFile); e != nil {
		log.Fatal("ReadFile: ", e)
		return nil, e
	} else {
		pool.AppendCertsFromPEM(ca)
	}
	return pool, nil
}

func get(url string) (string, error) {
	cert, e := loadCA(cacert)
	if e != nil {
		return "", e
	}
	client := &http.Client{
		// 超时时间：5秒
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: cert},
		}}
	req, err := http.NewRequest("GET", url, nil)
	token, err := ioutil.ReadFile(token)
	if err != nil {
		return "", e
	}

	req.Header.Add("Authorization", `Bearer `+string(token))

	resp, err := client.Do(req)
	if err != nil {
		// panic(err)
		return "", err
	}
	defer resp.Body.Close()
	var buffer [512 * 1024]byte
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

	return result.String(), nil
}

// GetInfo export info get
func GetInfo(api string) (string, error) {
	url := createURL(api)
	data, err := get(url)
	if err != nil {
		return "", err
	}
	return data, nil
}
