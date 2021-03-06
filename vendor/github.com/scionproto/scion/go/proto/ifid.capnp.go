// Code generated by capnpc-go. DO NOT EDIT.

package proto

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type IFID struct{ capnp.Struct }

// IFID_TypeID is the unique identifier for the type IFID.
const IFID_TypeID = 0x9d95fb13f80529b9

func NewIFID(s *capnp.Segment) (IFID, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return IFID{st}, err
}

func NewRootIFID(s *capnp.Segment) (IFID, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return IFID{st}, err
}

func ReadRootIFID(msg *capnp.Message) (IFID, error) {
	root, err := msg.RootPtr()
	return IFID{root.Struct()}, err
}

func (s IFID) String() string {
	str, _ := text.Marshal(0x9d95fb13f80529b9, s.Struct)
	return str
}

func (s IFID) OrigIF() uint64 {
	return s.Struct.Uint64(0)
}

func (s IFID) SetOrigIF(v uint64) {
	s.Struct.SetUint64(0, v)
}

// IFID_List is a list of IFID.
type IFID_List struct{ capnp.List }

// NewIFID creates a new list of IFID.
func NewIFID_List(s *capnp.Segment, sz int32) (IFID_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return IFID_List{l}, err
}

func (s IFID_List) At(i int) IFID { return IFID{s.List.Struct(i)} }

func (s IFID_List) Set(i int, v IFID) error { return s.List.SetStruct(i, v.Struct) }

func (s IFID_List) String() string {
	str, _ := text.MarshalList(0x9d95fb13f80529b9, s.List)
	return str
}

// IFID_Promise is a wrapper for a IFID promised by a client call.
type IFID_Promise struct{ *capnp.Pipeline }

func (p IFID_Promise) Struct() (IFID, error) {
	s, err := p.Pipeline.Struct()
	return IFID{s}, err
}

const schema_9cb1ca08a160c787 = "x\xda\x12\xf0s`2d\xdd\xcf\xc8\xc0\x10(\xc2\xca" +
	"\xf6\x7f\xa7&\xeb\x0f\xe1\xdfS\xe72\x04r32\xfe" +
	"o?\x9e\xb0\x90\xe3\xd4\xc69\x0c,\xec\x0c\x0c\x82G" +
	"\x9b\x04O\x82i{\x06\xdd\xff\x99i\x99)z\xc9\x89" +
	"\x05\x8cy\x05V\x9en\x9e.\x0c\x01\x8c\x8c\x81,\xcc" +
	",\x0c\x0c,\x8c\x0c\x0c\x82\xbcV\x0c\x0c\x81\x1c\xcc\x8c" +
	"\x81\"L\x8c\xf6\xf9E\x99\xe9\x9en\x8c\x9c\x0cL\x8c" +
	"\x9c\x0c\x8c\x80\x00\x00\x00\xff\xff+O\x1b\x0e"

func init() {
	schemas.Register(schema_9cb1ca08a160c787,
		0x9d95fb13f80529b9)
}
