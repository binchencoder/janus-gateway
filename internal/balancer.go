package internal

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/pborman/uuid"
	"google.golang.org/grpc/grpclog"

	options "binchencoder.com/ease-gateway/gateway/options"
	"binchencoder.com/letsgo/grpc"
	"binchencoder.com/letsgo/hashring"
)

const (
	hashKeyUUID    = "@uuid"
	hashKeySession = "@session"
)

// PreLoadBalance processes context to affect the load balancer.
func PreLoadBalance(ctx context.Context, balancer, hashHeyType string, req proto.Message) context.Context {
	if balancer == "" || balancer == options.LoadBalancer_ROUND_ROBIN.String() {
		return ctx
	}

	if balancer == options.LoadBalancer_CONSISTENT.String() {
		if hashHeyType == hashKeyUUID {
			hashKey := uuid.New()
			ctx = hashring.WithHashKey(ctx, hashKey)
			// Also put to gRPC metadata.
			ctx = grpc.ToMetadataOutgoing(ctx, "")
		} else if hashHeyType == hashKeySession {
			// TODO(chenbin): implement it.
			return ctx
		} else {
			// Hash key is a proto field.
			hashKey := fmt.Sprintf("%v", getProtoFiledValue(req, hashHeyType))
			ctx = hashring.WithHashKey(ctx, hashKey)
		}
	}

	return ctx
}

func getProtoFiledValue(msg proto.Message, fieldPathStr string) reflect.Value {
	fieldPath := strings.Split(fieldPathStr, ".")
	v := reflect.ValueOf(msg).Elem()
	for _, fieldName := range fieldPath {
		f, _, err := fieldByProtoName(v, fieldName)
		if err != nil {
			grpclog.Printf("field not found in %T: %s, %v", msg, strings.Join(fieldPath, "."), err)
			return reflect.Value{}
		}
		if !f.IsValid() {
			grpclog.Printf("field not found in %T: %s", msg, strings.Join(fieldPath, "."))
			return reflect.Value{}
		}

		switch f.Kind() {
		case reflect.Bool, reflect.Float32, reflect.Float64, reflect.Int32, reflect.Int64, reflect.String, reflect.Uint32, reflect.Uint64:
			v = f
		case reflect.Ptr:
			if f.IsNil() {
				grpclog.Printf("field is nil in %T: %s", msg, strings.Join(fieldPath, "."))
				return reflect.Value{}
			}
			v = f.Elem()
			continue
		case reflect.Struct:
			v = f
			continue
		default:
			grpclog.Printf("unexpected type %s in %T", f.Type(), msg)
			return reflect.Value{}
		}
	}
	return v
}

// fieldByProtoName looks up a field whose corresponding protobuf field name is "name".
// "m" must be a struct value. It returns zero reflect.Value if no such field found.
func fieldByProtoName(m reflect.Value, name string) (reflect.Value, *proto.Properties, error) {
	props := proto.GetProperties(m.Type())

	// look up field name in oneof map
	if op, ok := props.OneofTypes[name]; ok {
		v := reflect.New(op.Type.Elem())
		field := m.Field(op.Field)
		if !field.IsNil() {
			return reflect.Value{}, nil, fmt.Errorf("field already set for %s oneof", props.Prop[op.Field].OrigName)
		}
		field.Set(v)
		return v.Elem().Field(0), op.Prop, nil
	}

	for _, p := range props.Prop {
		if p.OrigName == name {
			return m.FieldByName(p.Name), p, nil
		}
		if p.JSONName == name {
			return m.FieldByName(p.Name), p, nil
		}
	}
	return reflect.Value{}, nil, nil
}
