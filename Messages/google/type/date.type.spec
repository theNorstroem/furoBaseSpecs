name: date
type: Date
description: 'Date, https://github.com/googleapis/googleapis/blob/master/google/date.proto '
__proto:
    package: google.type
    targetfile: date.proto
    imports: []
    options:
        cc_enable_arenas : true
        go_package : "google.golang.org/genproto/googleapis/type/date;date"
        java_multiple_files : true
        java_outer_classname : "DateProto"
        java_package : "com.google.type"
        objc_class_prefix : "GTP"
fields:
    display_name:
        type: string
        description: Localized String representation of date
        __proto:
            number: 4
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: Datum
            options: null
            readonly: true
            repeated: false
            typespecific: null
        constraints: {}
    year:
        type: int32
        description: Year of date. Must be from 1 to 9999, or 0 if specifying a date without a year.
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: ""
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    month:
        type: int32
        description: Month of year. Must be from 1 to 12, or 0 if specifying a year without a month and day.
        __proto:
            number: 2
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: ""
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    day:
        type: int32
        description: Day of month. Must be from 1 to 31 and valid for the year and month, or 0. if specifying a year by itself or a year and month where the day is not significant.
        __proto:
            number: 3
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: ""
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
