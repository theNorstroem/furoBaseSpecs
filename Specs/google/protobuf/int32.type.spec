name: int32value
type: Int32Value
description: Wrapper message for `int32`.  https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/wrappers.proto
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
        type: int32
        description: The JSON representation for `Int32Value` is JSON number
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta: null
        constraints:
            max:
                is: "2147483647"
                message: out of range
            min:
                is: −2147483648
                message: out of range
