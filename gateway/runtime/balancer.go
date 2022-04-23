package runtime

import (
	"context"

	"github.com/pborman/uuid"
	"google.golang.org/protobuf/proto"

	options "github.com/binchencoder/janus-gateway/httpoptions"
	"github.com/binchencoder/letsgo/grpc"
	"github.com/binchencoder/letsgo/hashring"
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
			// hashKey := fmt.Sprintf("%v", getProtoFiledValue(req, hashHeyType))
			hashKey := ""
			ctx = hashring.WithHashKey(ctx, hashKey)
		}
	}

	return ctx
}

// func getProtoFiledValue(msg proto.Message, fieldPathStr string) reflect.Value {
// 	fieldPath := strings.Split(fieldPathStr, ".")
// 	v := reflect.ValueOf(msg).Elem()
// 	for _, fieldName := range fieldPath {
// 		f, _, err := fieldByProtoName(v, fieldName)
// 		if err != nil {
// 			grpclog.Printf("field not found in %T: %s, %v", msg, strings.Join(fieldPath, "."), err)
// 			return reflect.Value{}
// 		}
// 		if !f.IsValid() {
// 			grpclog.Printf("field not found in %T: %s", msg, strings.Join(fieldPath, "."))
// 			return reflect.Value{}
// 		}

// 		switch f.Kind() {
// 		case reflect.Bool, reflect.Float32, reflect.Float64, reflect.Int32, reflect.Int64, reflect.String, reflect.Uint32, reflect.Uint64:
// 			v = f
// 		case reflect.Ptr:
// 			if f.IsNil() {
// 				grpclog.Printf("field is nil in %T: %s", msg, strings.Join(fieldPath, "."))
// 				return reflect.Value{}
// 			}
// 			v = f.Elem()
// 			continue
// 		case reflect.Struct:
// 			v = f
// 			continue
// 		default:
// 			grpclog.Printf("unexpected type %s in %T", f.Type(), msg)
// 			return reflect.Value{}
// 		}
// 	}
// 	return v
// }
