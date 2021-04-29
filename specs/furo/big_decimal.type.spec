name: big_decimal
type: BigDecimal
description: |
    A BigDecimal is defined by two values: an arbitrary precision integer and a 32-bit integer scale.
    The value of the BigDecimal is defined to be unscaledValue*10^{-scale}.'

    This java-oriented BigDecimal consists of an arbitrary precision integer unscaled value and a 32-bit integer scale.
    If zero or positive, the scale is the number of digits to the right of the decimal point.
    If negative, the unscaled value of the number is multiplied by ten to the power of the negation of the scale.
    The value of the number represented by the BigDecimal is therefore (unscaledValue × 10-scale).
    See also https://docs.oracle.com/en/java/javase/13/docs/api/java.base/java/math/BigDecimal.html.
__proto:
    package: furo
    targetfile: furo.proto
    imports: []
    options:
        cc_enable_arenas: "true"
        csharp_namespace: Furo
        go_package: github.com/theNorstroem/FuroBaseSpecs/dist/pb/furo;furopb
        java_multiple_files: "true"
        java_outer_classname: FuroProto
        java_package: pro.furo
        objc_class_prefix: FPB
fields:
    unscaled_value:
        type: sint64
        description: |
            This java-oriented BigDecimal consists of an arbitrary precision integer unscaled value and a 32-bit integer scale.
            If zero or positive, the scale is the number of digits to the right of the decimal point.
            If negative, the unscaled value of the number is multiplied by ten to the power of the negation of the scale.
            The value of the number represented by the BigDecimal is therefore (unscaledValue × 10-scale).
            See also https://docs.oracle.com/en/java/javase/13/docs/api/java.base/java/math/BigDecimal.html.
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: ""
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: { }
    scale:
        type: int32
        description: |
            Number of digits to the right of the decimal point.

            If zero or positive, the scale is the number of digits to the right of the decimal point.
            If negative, the unscaled value of the number is multiplied by ten to the power of the
            negation of the scale.
        __proto:
            number: 2
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: ""
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
