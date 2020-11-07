name: Tree
type: Tree
description: A menu tree
__proto:
    package: furo.navigation
    targetfile: navigation.proto
    imports: []
    options:
        go_package: github.com/veith/furoc-gen-u33e/dist/pb/furo/navigation;navigationpb
        java_multiple_files: "true"
        java_outer_classname: NavigationProto
        java_package: com.furo.basefuro.navigation
fields:
    id:
        type: string
        description: Id of the node, this field value must be unique
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: furo.navigation.Tree.id.label
            options:
                flags: []
                list: []
            readonly: true
            repeated: false
            typespecific: null
        constraints: {}
    display_name:
        type: string
        description: String representation of the tree
        __proto:
            number: 2
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: furo.navigation.Tree.display_name.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    labels:
        type: string
        description: Labels / flags for the value, something like unspecified, empty, confidential, absent,... Can be used for AI, UI, Business Logic,...
        __proto:
            number: 3
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: furo.navigation.Tree.labels.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
    attributes:
        type: map<string,string>
        description: 'Attributes for a value, something like confidential-msg: you are not allowed to see this value'
        __proto:
            number: 4
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: furo.navigation.Tree.attributes.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    root:
        type: furo.navigation.Navigationnode
        description: The root of the tree
        __proto:
            number: 5
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: furo.navigation.Tree.root.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
