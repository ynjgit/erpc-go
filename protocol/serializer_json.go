package protocol

import "encoding/json"

type jsonSerializer struct{}

func (j *jsonSerializer) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (j *jsonSerializer) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
