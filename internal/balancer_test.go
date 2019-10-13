package internal

import (
	"context"
	"testing"

	"github.com/golang/protobuf/proto"

	options "binchencoder.com/ease-gateway/gateway/options"
	"binchencoder.com/letsgo/hashring"
)

const DefaultHashKey = "8daad76a-dbb6-4f95-855d-7cfceb89afa1"

type msgA struct {
	StringValue string `protobuf:"bytes,1,opt,name=string_value" json:"string_value,omitempty"`
}

func (ma *msgA) Reset()         { *ma = msgA{} }
func (ma *msgA) String() string { return proto.CompactTextString(ma) }
func (*msgA) ProtoMessage()     {}

type msgB struct {
	Nested *msgA `protobuf:"bytes,1,opt,name=nested" json:"nested,omitempty"`
}

func (mb *msgB) Reset()         { *mb = msgB{} }
func (mb *msgB) String() string { return proto.CompactTextString(mb) }
func (*msgB) ProtoMessage()     {}

func (mb *msgB) GetNested() *msgA {
	if mb != nil {
		return mb.Nested
	}
	return nil
}

type msgC struct {
	Nested *msgB `protobuf:"bytes,1,opt,name=nested" json:"nested,omitempty"`
}

func (mc *msgC) Reset()         { *mc = msgC{} }
func (mc *msgC) String() string { return proto.CompactTextString(mc) }
func (*msgC) ProtoMessage()     {}

func (mc *msgC) GetNested() *msgB {
	if mc != nil {
		return mc.Nested
	}
	return nil
}

func TestGetProtoFiledValue(t *testing.T) {
	a := msgA{
		StringValue: "foo",
	}
	v := getProtoFiledValue(&a, "string_value")
	if v.String() != "foo" {
		t.Errorf("Expect string %s but got %s", "foo", v.String())
	}

	b := msgB{
		Nested: &a,
	}
	v = getProtoFiledValue(&b, "nested.string_value")
	if v.String() != "foo" {
		t.Errorf("Expect string %s but got %s", "foo", v.String())
	}

	c := msgC{
		Nested: &b,
	}
	v = getProtoFiledValue(&c, "nested.nested.string_value")
	if v.String() != "foo" {
		t.Errorf("Expect string %s but got %s", "foo", v.String())
	}
}

func TestPreLoadBalance(t *testing.T) {
	// Generate UUID.
	req := msgA{
		StringValue: DefaultHashKey,
	}
	ctx := PreLoadBalance(context.Background(), options.LoadBalancer_CONSISTENT.String(), hashKeyUUID, &req)
	key := hashring.GetHashKeyOrEmpty(ctx)
	if len(key) != 36 {
		t.Errorf("Expect getting hash key with length 36s but got %s", key)
	}

	// Proto field.
	req = msgA{
		StringValue: DefaultHashKey,
	}
	ctx = PreLoadBalance(context.Background(), options.LoadBalancer_CONSISTENT.String(), "string_value", &req)
	key = hashring.GetHashKeyOrEmpty(ctx)
	if key != DefaultHashKey {
		t.Errorf("Expect getting hash key %s but got %s", DefaultHashKey, key)
	}
}
