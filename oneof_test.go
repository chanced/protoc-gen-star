package pgs

import (
	"testing"

	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/stretchr/testify/assert"
)

func TestOneof_Name(t *testing.T) {
	t.Parallel()

	o := &oneof{desc: &descriptor.OneofDescriptorProto{Name: proto.String("foo")}}
	assert.Equal(t, "foo", o.Name().String())
}

func TestOneOf_FullyQualifiedName(t *testing.T) {
	t.Parallel()

	o := &oneof{fqn: "one_of"}
	assert.Equal(t, o.fqn, o.FullyQualifiedName())
}

func TestOneof_Syntax(t *testing.T) {
	t.Parallel()

	m := dummyMsg()
	o := &oneof{}
	m.AddOneOf(o)

	assert.Equal(t, m.Syntax(), o.Syntax())
}

func TestOneof_Package(t *testing.T) {
	t.Parallel()

	m := dummyMsg()
	o := &oneof{}
	m.AddOneOf(o)

	assert.NotNil(t, o.Package())
	assert.Equal(t, m.Package(), o.Package())
}

func TestOneof_File(t *testing.T) {
	t.Parallel()

	m := dummyMsg()
	o := &oneof{}
	m.AddOneOf(o)

	assert.NotNil(t, o.File())
	assert.Equal(t, m.File(), o.File())
}

func TestOneof_BuildTarget(t *testing.T) {
	t.Parallel()

	m := dummyMsg()
	o := &oneof{}
	m.AddOneOf(o)

	assert.False(t, o.BuildTarget())
	m.SetParent(&file{buildTarget: true})
	assert.True(t, o.BuildTarget())
}

func TestOneof_Descriptor(t *testing.T) {
	t.Parallel()

	o := &oneof{desc: &descriptor.OneofDescriptorProto{}}

	assert.Equal(t, o.desc, o.Descriptor())
}

func TestOneof_Message(t *testing.T) {
	t.Parallel()

	m := dummyMsg()
	o := &oneof{}
	m.AddOneOf(o)

	assert.Equal(t, m, o.Message())
}

func TestOneof_Imports(t *testing.T) {
	t.Parallel()

	o := &oneof{}
	assert.Empty(t, o.Imports())

	o.AddField(&mockField{i: []File{&file{}, &file{}}, Field: &field{}})
	assert.Len(t, o.Imports(), 1)

	f := &file{desc: &descriptor.FileDescriptorProto{
		Name: proto.String("foobar"),
	}}
	o.AddField(&mockField{i: []File{f}, Field: &field{}})
	assert.Len(t, o.Imports(), 2)
}

func TestOneof_Extension(t *testing.T) {
	// cannot be parallel

	o := &oneof{desc: &descriptor.OneofDescriptorProto{}}
	assert.NotPanics(t, func() { o.Extension(nil, nil) })
}

func TestOneof_Fields(t *testing.T) {
	t.Parallel()

	o := &oneof{}
	assert.Empty(t, o.Fields())

	o.AddField(&field{})
	assert.Len(t, o.Fields(), 1)
}

func TestOneof_Accept(t *testing.T) {
	t.Parallel()

	o := &oneof{}
	assert.NoError(t, o.Accept(nil))

	v := &mockVisitor{err: errors.New("")}
	assert.Error(t, o.Accept(v))
	assert.Equal(t, 1, v.oneof)
}

func TestOneof_ChildAtPath(t *testing.T) {
	t.Parallel()

	o := &oneof{}
	assert.Equal(t, o, o.ChildAtPath(nil))
	assert.Nil(t, o.ChildAtPath([]int32{1}))
}

type mockOneOf struct {
	OneOf
	i   []File
	m   Message
	err error
}

func (o *mockOneOf) Imports() []File { return o.i }

func (o *mockOneOf) setMessage(m Message) { o.m = m }

func (o *mockOneOf) Accept(v Visitor) error {
	_, err := v.VisitOneOf(o)
	if o.err != nil {
		return o.err
	}
	return err
}

func dummyOneof() *oneof {
	m := dummyMsg()
	o := &oneof{desc: &descriptor.OneofDescriptorProto{Name: proto.String("oneof")}}
	m.AddOneOf(o)
	return o
}
