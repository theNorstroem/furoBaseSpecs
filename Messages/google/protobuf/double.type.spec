name: doublevalue
type: DoubleValue
description: Wrapper message for `double`.  https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/wrappers.proto
__proto:
    package: google.protobuf
    targetfile: wrappers.proto
    imports: []
    options:
        cc_enable_arenas: "true"
        csharp_namespace: Google.Protobuf.WellKnownTypes
        go_package: google.golang.org/protobuf/types/known/wrapperspb
        java_multiple_files: "true"
        java_outer_classname: WrappersProto
        java_package: com.google.protobuf
        objc_class_prefix: GPB
fields:
    value:
        type: double
        description: The JSON representation for `DoubleValue` is JSON number
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta: null
        constraints: {}
