package proto

// Marshaler is implemented by messages that can marshal themselves.
// This interface is used by the following functions: Size, Marshal,
// Buffer.Marshal, and Buffer.EncodeMessage.
//
// Deprecated: Do not implement.
type Marshaler interface {
	// Marshal formats the encoded bytes of the message.
	// It should be deterministic and emit valid protobuf wire data.
	// The caller takes ownership of the returned buffer.
	Marshal() ([]byte, error)
}

// Unmarshaler is implemented by messages that can unmarshal themselves.
// This interface is used by the following functions: Unmarshal, UnmarshalMerge,
// Buffer.Unmarshal, Buffer.DecodeMessage, and Buffer.DecodeGroup.
//
// Deprecated: Do not implement.
type Unmarshaler interface {
	// Unmarshal parses the encoded bytes of the protobuf wire input.
	// The provided buffer is only valid for during method call.
	// It should not reset the receiver message.
	Unmarshal([]byte) error
}
