name: empty
type: Empty
description: Furo annotated type wrapper message for `empty`. Empty has no values and only contains the labels and attributes
__proto:
    package: furo.fat
    targetfile: fat.proto
    imports: []
    options: {}
fields:
    labels:
        type: string
        description: Labels / flags for the value, something like unspecified, empty, confidential, absent,... Can be used for AI, UI, Business Logic,...
        __proto:
            number: 2
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: ""
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
    attributes:
        type: map<string,string>
        description: 'Attributes for a value, something like confidential-msg: you are not allowed to see this value '
        __proto:
            number: 3
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: ""
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
