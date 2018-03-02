package apigateway

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/jsonapi"
)

type Response *http.Response

func (resp *Response) ParseJsonApi(model interface{}) error {
	return jsonapi.UnmarshalPayload(resp.Body, model)
}

func (resp *Response) ParseJsonObject() map[string]interface{} {
	jsonObj := make(map[string]interface{})
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, jsonObj)
	return jsonObj
}

type PreparedRequest struct {
	body      string
	headers   map[string]string
	targetURL string
}

func (p *PreparedRequest) sendRequest(method string) (Response, error) {
	req, err := http.NewRequest(method, p.targetURL, strings.NewReader(p.body))
	if err != nil {
		return nil, err
	}
	for k, v := range p.headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	return Response(resp), err
}

func (p *PreparedRequest) Post(url string) (Response, error) {
	return p.sendRequest(http.MethodPost)
}

func (p *PreparedRequest) Get(url string) (Response, error) {
	return p.sendRequest(http.MethodGet)
}

func Prepare(body string, headers map[string]string) *PreparedRequest {
	return &PreparedRequest{body, headers}
}
