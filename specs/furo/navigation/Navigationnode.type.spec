name: Navigationnode
type: Navigationnode
description: This is an example signature of a navigation node, which can be used for the components of @furo/navigation
__proto:
    package: furo.navigation
    targetfile: navigation.proto
    imports:
        - furo/furo.proto
        - google/protobuf/any.proto
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
            label: furo.navigation.Navigationnode.id.label
            options:
                flags: []
                list: []
            readonly: true
            repeated: false
            typespecific: null
        constraints: {}
    display_name:
        type: string
        description: String representation of the node
        __proto:
            number: 2
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: furo.navigation.Navigationnode.display_name.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    secondary_text:
        type: string
        description: Secondary text of the node [optional]
        __proto:
            number: 3
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: furo.navigation.Navigationnode.secondary_text.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    icon:
        type: string
        description: icon of a node. When used in furo-tree it will be displayed as leading icon [optional]
        __proto:
            number: 5
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: furo.navigation.Navigationnode.icon.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    panel:
        type: string
        description: Which panel (i.e. view, edit, display) opens the node type (which is defined in property link). The value of this field must correspond to your registry.
        __proto:
            number: 6
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: furo.navigation.Navigationnode.panel.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    key_words:
        type: string
        description: key words of the node
        __proto:
            number: 7
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: furo.navigation.Navigationnode.key_words.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    has_error:
        type: bool
        description: if node has error
        __proto:
            number: 8
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: furo.navigation.Navigationnode.has_error.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    open:
        type: bool
        description: node is open or not
        __proto:
            number: 9
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: furo.navigation.Navigationnode.open.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    link:
        type: furo.Link
        description: Deeplink information of this node
        __proto:
            number: 10
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: furo.navigation.Navigationnode.link.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    is_group_label:
        type: bool
        description: This node is a group label
        __proto:
            number: 11
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: furo.navigation.Navigationnode.is_group_label.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    children:
        type: furo.navigation.Navigationnode
        description: Children of this node
        __proto:
            number: 12
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: furo.navigation.Navigationnode.children.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
    flags:
        type: string
        description: Attribute flags e.g. important, negative, positive. Can be used for custom annotations for styling, logic,...
        __proto:
            number: 13
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: furo.navigation.Navigationnode.flags.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
    payload:
        type: google.protobuf.Any
        description: Optional payload
        __proto:
            number: 14
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: furo.navigation.Navigationnode.payload.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
