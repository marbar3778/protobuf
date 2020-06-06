package proto

// Unmarshal parses the protocol buffer representation in buf and places the
// decoded result in pb.  If the struct underlying pb does not match
// the data in buf, the results can be unpredictable.
//
// Unmarshal resets pb before starting to unmarshal, so any
// existing data in pb is always removed. Use UnmarshalMerge
// to preserve and append to existing data.
func Unmarshal(buf []byte, pb Message) error {
	pb.Reset()
	if u, ok := pb.(newUnmarshaler); ok {
		return u.XXX_Unmarshal(buf)
	}
	if u, ok := pb.(Unmarshaler); ok {
		return u.Unmarshal(buf)
	}
	return NewBuffer(buf).Unmarshal(pb)
}

// UnmarshalMerge parses the protocol buffer representation in buf and
// writes the decoded result to pb.  If the struct underlying pb does not match
// the data in buf, the results can be unpredictable.
//
// UnmarshalMerge merges into existing data in pb.
// Most code should use Unmarshal instead.
func UnmarshalMerge(buf []byte, pb Message) error {
	if u, ok := pb.(newUnmarshaler); ok {
		return u.XXX_Unmarshal(buf)
	}
	if u, ok := pb.(Unmarshaler); ok {
		// NOTE: The history of proto have unfortunately been inconsistent
		// whether Unmarshaler should or should not implicitly clear itself.
		// Some implementations do, most do not.
		// Thus, calling this here may or may not do what people want.
		//
		// See https://github.com/golang/protobuf/issues/424
		return u.Unmarshal(buf)
	}
	return NewBuffer(buf).Unmarshal(pb)
}
