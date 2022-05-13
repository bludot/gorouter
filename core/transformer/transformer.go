package transformer

import (
	"encoding/json"
)

// transforms data to types (json, xml, yaml)
type TransformerService struct {
}

var transformer Transformer

func (t *TransformerService) ToJson(data interface{}) (string, error) {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(jsonByte), nil
}

func GetTransformer() Transformer {
	if transformer == nil {
		transformer = &TransformerService{}
	}

	return transformer
}
