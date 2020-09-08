name: meta
type: Meta
description: meta info
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
    fields:
        type: map<string, furo.MetaField>
        description: fields of meta info
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta: null
        constraints: {}
