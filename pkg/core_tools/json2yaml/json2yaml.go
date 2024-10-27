package json2yaml

import (
	"bytes"
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func ConvertJson2Yaml(jsonStr string) (string, error) {
	var jsonData map[string]interface{}
	var b bytes.Buffer
	encoder := yaml.NewEncoder(&b)
	encoder.SetIndent(2)

	if err := json.Unmarshal([]byte(jsonStr), &jsonData); err != nil {
		return "", err
	}

	err := encoder.Encode(jsonData)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
