package runtime

import (
	"context"
	"testing"

	pb "github.com/binchencoder/janus-gateway/gateway/runtime/internal/examplepb"
	"google.golang.org/protobuf/proto"

	options "github.com/binchencoder/janus-gateway/httpoptions"
	"github.com/binchencoder/letsgo/hashring"
)

const DefaultHashKey = "8daad76a-dbb6-4f95-855d-7cfceb89afa1"

type msgA struct {
	pb proto.Message
}

// func TestGetProtoFiledValue(t *testing.T) {
// 	a := msgA{
// 		StringValue: "foo",
// 	}
// 	v := getProtoFiledValue(&a, "string_value")
// 	if v.String() != "foo" {
// 		t.Errorf("Expect string %s but got %s", "foo", v.String())
// 	}

// 	b := msgB{
// 		Nested: &a,
// 	}
// 	v = getProtoFiledValue(&b, "nested.string_value")
// 	if v.String() != "foo" {
// 		t.Errorf("Expect string %s but got %s", "foo", v.String())
// 	}

// 	c := msgC{
// 		Nested: &b,
// 	}
// 	v = getProtoFiledValue(&c, "nested.nested.string_value")
// 	if v.String() != "foo" {
// 		t.Errorf("Expect string %s but got %s", "foo", v.String())
// 	}
// }

func TestPreLoadBalance(t *testing.T) {
	// Generate UUID.
	req := msgA{
		pb: &pb.Proto3Message{
			StringValue: DefaultHashKey,
		},
	}

	ctx := PreLoadBalance(context.Background(), options.LoadBalancer_CONSISTENT.String(), hashKeyUUID, req.pb)
	key := hashring.GetHashKeyOrEmpty(ctx)
	if len(key) != 36 {
		t.Errorf("Expect getting hash key with length 36s but got %s", key)
	}

	// Proto field.
	req = msgA{
		pb: &pb.Proto3Message{
			StringValue: DefaultHashKey,
		},
	}
	ctx = PreLoadBalance(context.Background(), options.LoadBalancer_CONSISTENT.String(), "string_value", req.pb)
	key = hashring.GetHashKeyOrEmpty(ctx)
	if key != DefaultHashKey {
		t.Errorf("Expect getting hash key %s but got %s", DefaultHashKey, key)
	}
}
