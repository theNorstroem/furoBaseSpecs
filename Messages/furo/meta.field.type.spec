name: metafield
type: MetaField
description: fields of meta info
__proto:
    package: furo
    targetfile: meta.proto
    imports: []
    options:
        cc_enable_arenas: "true"
        csharp_namespace: Furo.Meta
        go_package: github.com/theNorstroem/FuroBaseSpecs/dist/pb/furo/meta;metapb
        java_multiple_files: "true"
        java_outer_classname: MetaProto
        java_package: pro.furo.meta
        objc_class_prefix: FPB
fields:
    meta:
        type: furo.FieldMeta
        description: meta informatioxn of a field
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta: null
        constraints: {}
    constraints:
        type: map<string,furo.FieldConstraint>
        description: constraints for a field
        __proto:
            number: 2
            oneof: ""
        __ui: null
        meta: null
        constraints: {}
