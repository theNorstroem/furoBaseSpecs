name: reference
type: Reference
description: reference
__proto:
    package: furo
    targetfile: reference.proto
    imports:
        - furo/link.proto
    options: {}
fields:
    display_name:
        type: string
        description: String representation of the reference
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: ""
            options: null
            readonly: true
            repeated: false
            typespecific: null
        constraints: {}
    id:
        type: string
        description: Id of the reference
        __proto:
            number: 2
            oneof: ""
        __ui: null
        meta: null
        constraints: {}
    link:
        type: furo.Link
        description: Hateoas link
        __proto:
            number: 3
            oneof: ""
        __ui: null
        meta: null
        constraints: {}
