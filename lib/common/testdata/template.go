package testdata

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/google/jsonapi"
)

type TestData struct {
	Payload []byte
}

func (d *TestData) AssignValues(model interface{}) (string, error) {
	t := template.Must(template.New("payload").Parse(string(d.Payload)))
	var payload bytes.Buffer
	if err := t.Execute(&payload, model); err != nil {
		return "", err
	}
	return payload.String(), nil
}

func New(filePath string) *TestData {
	payload, _ := ioutil.ReadFile(filePath)
	return &TestData{payload}
}

func ParseJsonApi(filePath string, model interface{}) error {
	if file, err := os.Open(filePath); err == nil {
		return jsonapi.UnmarshalPayload(file, model)
	} else {
		return err
	}
}
