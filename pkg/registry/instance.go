package registry

import "encoding/json"

type EndpointInfo struct {
	Ip       string                 `json:"ip,omitempty"`
	Port     string                 `json:"port,omitempty"`
	MetaData map[string]interface{} `json:"meta_data,omitempty"`
}

func (e *EndpointInfo) Marshal() (string, error) {
	data, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func UnMarshal(data []byte) (*EndpointInfo, error) {
	ep := &EndpointInfo{}
	if err := json.Unmarshal(data, ep); err != nil {
		return nil, err
	}
	return ep, nil
}
