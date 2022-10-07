name: FeatureToggle
type: FeatureToggle
description: FeatureToggle is the state for a feature toggle identified by the key field
lifecycle: null
__proto:
    package: furo
    targetfile: furo.proto
    imports: []
    options:
        go_package: github.com/theNorstroem/FuroBaseSpecs/dist/pb/furo;furopb
        java_multiple_files: "true"
        java_outer_classname: FuroProto
        java_package: pro.furo.furo
fields:
    key:
        type: string
        description: Id of the furo.FeatureToggle.
        __proto:
            number: 1
        __ui: null
        meta:
            default: ""
            placeholder: furo.featuretoggle.key.placeholder
            hint: ""
            label: furo.featuretoggle.key.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    state:
        type: bool
        description: True for allowed, false for denied.
        __proto:
            number: 2
        __ui: null
        meta:
            default: ""
            placeholder: furo.featuretoggle.state.placeholder
            hint: ""
            label: furo.featuretoggle.state.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
