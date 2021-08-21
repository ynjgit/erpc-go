package protocol

// Serializer the serializer interface
type Serializer interface {
	Unmarshal(data []byte, v interface{}) error
	Marshal(v interface{}) ([]byte, error)
}

var defaultSerializer = new(jsonSerializer)

// Unmarshal serialize unmarshal
func Unmarshal(data []byte, v interface{}) error {
	return defaultSerializer.Unmarshal(data, v)
}

// Marshal serialize marshal
func Marshal(v interface{}) ([]byte, error) {
	return defaultSerializer.Marshal(v)
}
