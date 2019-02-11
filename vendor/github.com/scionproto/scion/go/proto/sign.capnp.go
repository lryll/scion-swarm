// Code generated by capnpc-go. DO NOT EDIT.

package proto

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type SignedBlob struct{ capnp.Struct }

// SignedBlob_TypeID is the unique identifier for the type SignedBlob.
const SignedBlob_TypeID = 0x9f32478537fae352

func NewSignedBlob(s *capnp.Segment) (SignedBlob, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SignedBlob{st}, err
}

func NewRootSignedBlob(s *capnp.Segment) (SignedBlob, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SignedBlob{st}, err
}

func ReadRootSignedBlob(msg *capnp.Message) (SignedBlob, error) {
	root, err := msg.RootPtr()
	return SignedBlob{root.Struct()}, err
}

func (s SignedBlob) String() string {
	str, _ := text.Marshal(0x9f32478537fae352, s.Struct)
	return str
}

func (s SignedBlob) Blob() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s SignedBlob) HasBlob() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SignedBlob) SetBlob(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s SignedBlob) Sign() (Sign, error) {
	p, err := s.Struct.Ptr(1)
	return Sign{Struct: p.Struct()}, err
}

func (s SignedBlob) HasSign() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SignedBlob) SetSign(v Sign) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewSign sets the sign field to a newly
// allocated Sign struct, preferring placement in s's segment.
func (s SignedBlob) NewSign() (Sign, error) {
	ss, err := NewSign(s.Struct.Segment())
	if err != nil {
		return Sign{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// SignedBlob_List is a list of SignedBlob.
type SignedBlob_List struct{ capnp.List }

// NewSignedBlob creates a new list of SignedBlob.
func NewSignedBlob_List(s *capnp.Segment, sz int32) (SignedBlob_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return SignedBlob_List{l}, err
}

func (s SignedBlob_List) At(i int) SignedBlob { return SignedBlob{s.List.Struct(i)} }

func (s SignedBlob_List) Set(i int, v SignedBlob) error { return s.List.SetStruct(i, v.Struct) }

func (s SignedBlob_List) String() string {
	str, _ := text.MarshalList(0x9f32478537fae352, s.List)
	return str
}

// SignedBlob_Promise is a wrapper for a SignedBlob promised by a client call.
type SignedBlob_Promise struct{ *capnp.Pipeline }

func (p SignedBlob_Promise) Struct() (SignedBlob, error) {
	s, err := p.Pipeline.Struct()
	return SignedBlob{s}, err
}

func (p SignedBlob_Promise) Sign() Sign_Promise {
	return Sign_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type Sign struct{ capnp.Struct }

// Sign_TypeID is the unique identifier for the type Sign.
const Sign_TypeID = 0x844d9464f44e810a

func NewSign(s *capnp.Segment) (Sign, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Sign{st}, err
}

func NewRootSign(s *capnp.Segment) (Sign, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Sign{st}, err
}

func ReadRootSign(msg *capnp.Message) (Sign, error) {
	root, err := msg.RootPtr()
	return Sign{root.Struct()}, err
}

func (s Sign) String() string {
	str, _ := text.Marshal(0x844d9464f44e810a, s.Struct)
	return str
}

func (s Sign) Type() SignType {
	return SignType(s.Struct.Uint16(0))
}

func (s Sign) SetType(v SignType) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s Sign) Src() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Sign) HasSrc() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Sign) SetSrc(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s Sign) Signature() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s Sign) HasSignature() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Sign) SetSignature(v []byte) error {
	return s.Struct.SetData(1, v)
}

func (s Sign) Timestamp() uint32 {
	return s.Struct.Uint32(4)
}

func (s Sign) SetTimestamp(v uint32) {
	s.Struct.SetUint32(4, v)
}

// Sign_List is a list of Sign.
type Sign_List struct{ capnp.List }

// NewSign creates a new list of Sign.
func NewSign_List(s *capnp.Segment, sz int32) (Sign_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return Sign_List{l}, err
}

func (s Sign_List) At(i int) Sign { return Sign{s.List.Struct(i)} }

func (s Sign_List) Set(i int, v Sign) error { return s.List.SetStruct(i, v.Struct) }

func (s Sign_List) String() string {
	str, _ := text.MarshalList(0x844d9464f44e810a, s.List)
	return str
}

// Sign_Promise is a wrapper for a Sign promised by a client call.
type Sign_Promise struct{ *capnp.Pipeline }

func (p Sign_Promise) Struct() (Sign, error) {
	s, err := p.Pipeline.Struct()
	return Sign{s}, err
}

type SignType uint16

// SignType_TypeID is the unique identifier for the type SignType.
const SignType_TypeID = 0xf6b5bc42e3072fc9

// Values of SignType.
const (
	SignType_none    SignType = 0
	SignType_ed25519 SignType = 1
)

// String returns the enum's constant name.
func (c SignType) String() string {
	switch c {
	case SignType_none:
		return "none"
	case SignType_ed25519:
		return "ed25519"

	default:
		return ""
	}
}

// SignTypeFromString returns the enum value with a name,
// or the zero value if there's no such value.
func SignTypeFromString(c string) SignType {
	switch c {
	case "none":
		return SignType_none
	case "ed25519":
		return SignType_ed25519

	default:
		return 0
	}
}

type SignType_List struct{ capnp.List }

func NewSignType_List(s *capnp.Segment, sz int32) (SignType_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return SignType_List{l.List}, err
}

func (l SignType_List) At(i int) SignType {
	ul := capnp.UInt16List{List: l.List}
	return SignType(ul.At(i))
}

func (l SignType_List) Set(i int, v SignType) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

const schema_99440334ec0946a0 = "x\xdal\x91A\xab\x12Q\x1c\xc5\xcf\xf9_m4\x91" +
	"\x99\xcb\xb8,\xa2hQ\x82\x95\xa2H\xadB\xa2VE" +
	"\xd7\xda\x04\x05\x8dz\xd1\x01\x9d\x99t\"\x84 \x8a\xfa" +
	"\x06mZ\x15}\x84\xa0]\xbbV}\x0e\x17mZF" +
	"P\x10\x133=\xf5\xc1{\xbb\xc3\xe1\xc7=\xbf\xcb\xdf" +
	"\xfbz]\xda\xe5X\x00s\xaa|\";\xf9\xf2\xce\xcf" +
	"\xc9\xdb\xdb\xafajd\xf6\xe1f\xf5GW\xddx\x87" +
	"\xb28\x80\x7f\x9a\xaf\xfc\xb3\xfc\x9f\xbe\x83\xd9p\xf3\xa7" +
	"\xff\xe6V\xe7=t\xed\x08\xfb\x9b\x9f|\x16\xe9/\x9f" +
	"\x81\xd9\xb7\xcb\xcef\xf0\xe5\xf3/\xe8\x9a\xecY\xd0\x7f" +
	" \x1f\xfd\xa0\x00\x1fI\x1f\xadl\x15N\xa3K\xe3 " +
	"a\x94\\\xbb\x17N#\xdc%\x8d\xa7J@\x89\x80\x0e" +
	"\x9a\x80y\xa8hfB\xb2\xc1\xbc\xb3\xe7\x00\xf3X\xd1" +
	"\xcc\x85Z\xd8\xa0\x00:\x1c\x02f\xa6hR\xa1Vl" +
	"P\x01\xfaI^&\x8a\xe6\xb9\xd0M\xd7\x89\xa5\xbb7" +
	"\x03\xe9\x82\xcej9f\x1d\xc2:X\xc8\x04\xe9\xd3%" +
	"hw]\x1a.\xec*\x0d\x16`\xc2\x0a\x84\x95\x03\xee" +
	"\xb0\xb4=3\x19\xcc\xe3Q\xae^\xd9\xa9_\xcc\xd5\xcf" +
	"+\x9a+B\xbduo\xe5\xe5\x05E\xd3\x15\xba\xa3y" +
	"<\xda\xee\xb8\xf9\x9b\xf4\xf6\x17\x01\xe9\x1d3u\xdfY" +
	"'\xb6\x18*\xbe\xad\x9b9\xa8\xab\x03\xc0\x8d\xe2\xc8\xbe" +
	"\xb0\x93N\xaf\xd7\xbe\xfa/\x00\x00\xff\xff\x8eKl\x07"

func init() {
	schemas.Register(schema_99440334ec0946a0,
		0x844d9464f44e810a,
		0x9f32478537fae352,
		0xf6b5bc42e3072fc9)
}
