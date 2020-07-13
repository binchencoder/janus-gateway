# Common Validation Annotations by Spring

## @NotNull

The annotated element must not be null. Accepts any type.

## @Min

The annotated element must be a number whose value must be higher or equal
to the specified minimum.

Supported types are: BigDecimal, BigInteger, byte, short, int, long, and their
respective wrappers. Note that double and float are not supported due to
rounding errors (some providers might provide some approximative support).

null elements are considered valid.

## @Max

The annotated element must be a number whose value must be lower or equal
to the specified maximum.

Supported types are: BigDecimal, BigInteger, byte, short, int, long, and their
respective wrappers. Note that double and float are not supported due to
rounding errors (some providers might provide some approximative support).

null elements are considered valid.

## @Size(min = 0, max = 100)

The annotated element size must be between the specified boundaries (included).

Supported types are:

- String (string length is evaludated)
- Collection (collection size is evaluated)
- Map (map size is evaluated)
- Array (array length is evaluated)

null elements are considered valid.

## @Pattern(regex = "XXXX")

The annotated String must match the following regular expression.
The regular expression follows the Java regular expression conventions
see Pattern. Accepts String.

null elements are considered valid.

## Reference:

[Package javax.validation.constraints ](http://docs.oracle.com/javaee/6/api/javax/validation/constraints/package-summary.html)
