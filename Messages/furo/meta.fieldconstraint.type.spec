name: fieldconstraint
type: FieldConstraint
description: a single fieldconstraint
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
    is:
        type: string
        description: the constraint value as string, even it is a number
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: the constraint value as string, even it is a number
            label: is
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    message:
        type: string
        description: The message to display on constraint violation
        __proto:
            number: 2
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: message
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
