# Ease-gateway Validation Rule

## Rules Defination

Validation Rule定义可以在这里找到：

> ease-gateway/httpoptions/annotations.proto

```protobuf
// The opertaion type.
enum OperatorType {
	OPERATOR_TYPE_UNKNOWN = 0;

	GT      = 1; // Greater than
	LT      = 2; // Less than
	EQ      = 3; // Equals
	MATCH   = 4; // String pattern match.
	NON_NIL = 5; // Not nil
	LEN_GT  = 6; // String length great than
	LEN_LT  = 7; // String length less than
	LEN_EQ  = 8; // String length equals
}

// The supported function type list
enum FunctionType{
	FUNCTION_TYPE_UNKNOWN = 0;

	TRIM = 1; // String trim.
}

// ValueType is the type of the field.
enum ValueType {
	VALUE_TYPE_UNKNOWN = 0;

	NUMBER = 1; // Represent all number type like int,real
	STRING = 2; // String
	OBJ    = 3;
}

// ValidationRule defines the rule to validate the input value.
message ValidationRule {
	OperatorType operator = 1;
	ValueType    type     = 2;
	string       value    = 3;
	FunctionType function = 4;
}
```

> NOTE
>
> Validation Rule是作为Message's field的Extension Attribute。

每一条Rule包括：

* 类型
* 操作符
* 期待值
* 函数（可选）

## Samples

```protobuf
// The request message to be greeted.
message Payment {
    message SubPayment {
        PaymentType type = 1;
        // 100 > paied_amount > 10
        int64 paied_amount = 2 [
           (ease.api.rules) = {
                rules: {
                    type:NUMBER,
                    operator: GT,
                    value:"10",
                },
                rules: {
                    type:NUMBER,
                    operator: LT,
                    value:"100",
                 },
             }
        ];
    };

    PaymentType type = 1;
    // 长度=10
    string message_value_len_eq = 3 [
        (ease.api.rules) = {
               rules: {
               type:STRING,
                operator: LEN_EQ,
                value:"10",
            },
        }
    ];

    // 长度Trim之后小于21
    string message_value_len_gt = 4 [
        (ease.api.rules) = {
            rules: {
                type:STRING,
                operator: LEN_LT,
                value:"21",
                function: TRIM,
            },
        }
    ];
}
```

更多例子可以以下目录找到：

> ease-gateway/proto/examples/...

## Implemention Details

Validation Rule的定义会被编译到相呼应的xxx.gw.go文件中。每个类型会有一个Validation方法，如果这个类型包括任何Rule. 如果一个Field引用另外一个文件的类型（并且有Rule），那么生成的文件会引用相应的go package.

这是一个生成代码的例子：

```go

func Validate__sharedproto_Payment(v *Payment) error {
	if v == nil {
		return nil
	}
	// Validation for each Fields

	// Validation Field: Type

	// Validation Field: Amount

	if v.Amount <= 10 {
		return janus_demo_proto_sharedproto_shared_error
	}

	if v.Amount >= 100 {
		return janus_demo_proto_sharedproto_shared_error
	}
... ...
	return nil
}

```

## Add Validation Rule

加一个新的操作符或者函数是比较简单的。可以查阅有关的CR.代码在：

### Defination

> ease-gateway/httpoptions/annotations.proto

### Implementation Validation Rule

> ease-gateway/gateway/protoc-gen-grpc-gateway/internal/gengateway/template.go

所有Validation Rule的实现都在template.go 文件中,  搜索`validatorTemplate`， 参照现有实现Coding