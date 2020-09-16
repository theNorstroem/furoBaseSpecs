name: any
type: Any
description: |-
    Any contains an arbitrary serialized protocol buffer message along with a
    // URL that describes the type of the serialized message. client uses type `ArrayBuffer` for the value field .  https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/any.proto
__proto:
    package: google.protobuf
    targetfile: any.proto
    imports: []
    options:
        csharp_namespace: Google.Protobuf.WellKnownTypes
        go_package: google.golang.org/protobuf/types/known/anypb
        java_multiple_files: "true"
        java_outer_classname: AnyProto
        java_package: com.google.protobuf
        objc_class_prefix: GPB
fields:
    type_url:
        type: string
        description: ""
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta: null
        constraints: {}
    value:
        type: bytes
        description: ""
        __proto:
            number: 2
            oneof: ""
        __ui: null
        meta: null
        constraints: {}
