name: metafield
type: MetaField
description: fields of meta info
__proto:
    package: furo
    targetfile: meta.proto
    imports: []
    options: {}
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
