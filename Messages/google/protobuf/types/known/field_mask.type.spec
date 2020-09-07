name: field_mask
type: FieldMask
description: A field mask in update operations specifies which fields of the targeted resource are going to be updated. https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
__proto:
    package: google.protobuf.types.known
    targetfile: field_mask.proto
    imports: []
    options:
        csharp_namespace : "Google.Protobuf.WellKnownTypes"
        java_package : "com.google.protobuf"
        java_outer_classname : "FieldMaskProto"
        java_multiple_files : true
        objc_class_prefix : "GPB"
        go_package : "google.golang.org/protobuf/types/known/fieldmaskpb"
        cc_enable_arenas : true
fields:
    paths:
        type: string
        description: The implementation of any API method which has a FieldMask type field in the request should verify the included field paths, and return an `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
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
            repeated: true
            typespecific: null
        constraints: {}
