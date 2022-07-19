name: MessageContainerGroup
type: MessageContainerGroup
description: MessageContainerGroup Helps you grouping the errors and warnings for the UI.
lifecycle: null
__proto:
    package: furo
    targetfile: furo.proto
    imports:
        - google/protobuf/any.proto
    options:
        go_package: github.com/theNorstroem/FuroBaseSpecs/dist/pb/furo;furopb
        java_multiple_files: "true"
        java_outer_classname: FuroProto
        java_package: furo
fields:
    details:
        type: google.protobuf.Any
        description: A list of messages that carry the message details.
        __proto:
            number: 1
        __ui: null
        meta:
            default: ""
            placeholder: furo.messagecontainergroup.details.placeholder
            hint: ""
            label: furo.messagecontainergroup.details.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
    title:
        type: string
        description: Title of the group, should be a localized string.
        __proto:
            number: 2
        __ui: null
        meta:
            default: ""
            placeholder: furo.messagecontainergroup.title.placeholder
            hint: ""
            label: furo.messagecontainergroup.title.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
