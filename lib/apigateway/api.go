package apigateway

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/jsonapi"
)

type Response struct {
	Data       []byte
	StatusCode int
}

func (resp *Response) ParseJsonApi(model interface{}) error {
	return jsonapi.UnmarshalPayload(bytes.NewReader(resp.Data), model)
}

func (resp *Response) ParseJsonObject() map[string]interface{} {
	var jsonObj interface{}
	json.Unmarshal(resp.Data, &jsonObj)
	return jsonObj.(map[string]interface{})
}

type PreparedRequest struct {
	body      string
	headers   map[string]string
	targetURL string
}

func (p *PreparedRequest) sendRequest(method string) (*Response, error) {
	req, _ := http.NewRequest(method, p.targetURL, strings.NewReader(p.body))
	for k, v := range p.headers {
		req.Header.Set(k, v)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	return &Response{data, resp.StatusCode}, err
}

func (p *PreparedRequest) Post(url string) (*Response, error) {
	p.targetURL = url
	return p.sendRequest(http.MethodPost)
}

func (p *PreparedRequest) Get(url string) (*Response, error) {
	p.targetURL = url
	return p.sendRequest(http.MethodGet)
}

func Prepare(body string, headers map[string]string) *PreparedRequest {
	return &PreparedRequest{body, headers, ""}
}
