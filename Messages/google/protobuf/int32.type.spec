name: int32value
type: Int32Value
description: Wrapper message for `int32`.  https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/wrappers.proto
__proto:
    package: google.protobuf
    targetfile: wrappers.proto
    imports: []
    options: {}
fields:
    value:
        type: int32
        description: The JSON representation for `Int32Value` is JSON number
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta: null
        constraints:
            max:
                is: "2147483647"
                message: out of range
            min:
                is: −2147483648
                message: out of range
