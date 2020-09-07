name: integer_property
type: IntegerProperty
description: Integer type with embedded meta
__proto:
    package: furo
    targetfile: property.proto
    imports: []
    options: {}
fields:
    data:
        type: int32
        description: Integer data part
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta: null
        constraints:
            step:
                is: "1"
                message: ""
