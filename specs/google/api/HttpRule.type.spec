name: HttpRule
type: HttpRule
description: |-
    # gRPC Transcoding

     gRPC Transcoding is a feature for mapping between a gRPC method and one or
     more HTTP REST endpoints. It allows developers to build a single API service
     that supports both gRPC APIs and REST APIs. Many systems, including [Google
     APIs](https://github.com/googleapis/googleapis),
     [Cloud Endpoints](https://cloud.google.com/endpoints), [gRPC
     Gateway](https://github.com/grpc-ecosystem/grpc-gateway),
     and [Envoy](https://github.com/envoyproxy/envoy) proxy support this feature
     and use it for large scale production services.

     `HttpRule` defines the schema of the gRPC/REST mapping. The mapping specifies
     how different portions of the gRPC request message are mapped to the URL
     path, URL query parameters, and HTTP request body. It also controls how the
     gRPC response message is mapped to the HTTP response body. `HttpRule` is
     typically specified as an `google.api.http` annotation on the gRPC method.

     Each mapping specifies a URL path template and an HTTP method. The path
     template may refer to one or more fields in the gRPC request message, as long
     as each field is a non-repeated field with a primitive (non-message) type.
     The path template controls how fields of the request message are mapped to
     the URL path.

     Example:

         service Messaging {
           rpc GetMessage(GetMessageRequest) returns (Message) {
             option (google.api.http) = {
                 get: "/v1/{name=messages/*}"
             };
           }
         }
         message GetMessageRequest {
           string name = 1; // Mapped to URL path.
         }
         message Message {
           string text = 1; // The resource content.
         }

     This enables an HTTP REST to gRPC mapping as below:

     HTTP | gRPC
     -----|-----
     `GET /v1/messages/123456`  | `GetMessage(name: "messages/123456")`

     Any fields in the request message which are not bound by the path template
     automatically become HTTP query parameters if there is no HTTP request body.
     For example:

         service Messaging {
           rpc GetMessage(GetMessageRequest) returns (Message) {
             option (google.api.http) = {
                 get:"/v1/messages/{message_id}"
             };
           }
         }
         message GetMessageRequest {
           message SubMessage {
             string subfield = 1;
           }
           string message_id = 1; // Mapped to URL path.
           int64 revision = 2;    // Mapped to URL query parameter `revision`.
           SubMessage sub = 3;    // Mapped to URL query parameter `sub.subfield`.
         }

     This enables a HTTP JSON to RPC mapping as below:

     HTTP | gRPC
     -----|-----
     `GET /v1/messages/123456?revision=2&sub.subfield=foo` |
     `GetMessage(message_id: "123456" revision: 2 sub: SubMessage(subfield:
     "foo"))`

     Note that fields which are mapped to URL query parameters must have a
     primitive type or a repeated primitive type or a non-repeated message type.
     In the case of a repeated type, the parameter can be repeated in the URL
     as `...?param=A&param=B`. In the case of a message type, each field of the
     message is mapped to a separate parameter, such as
     `...?foo.a=A&foo.b=B&foo.c=C`.

     For HTTP methods that allow a request body, the `body` field
     specifies the mapping. Consider a REST update method on the
     message resource collection:

         service Messaging {
           rpc UpdateMessage(UpdateMessageRequest) returns (Message) {
             option (google.api.http) = {
               patch: "/v1/messages/{message_id}"
               body: "message"
             };
           }
         }
         message UpdateMessageRequest {
           string message_id = 1; // mapped to the URL
           Message message = 2;   // mapped to the body
         }

     The following HTTP JSON to RPC mapping is enabled, where the
     representation of the JSON in the request body is determined by
     protos JSON encoding:

     HTTP | gRPC
     -----|-----
     `PATCH /v1/messages/123456 { "text": "Hi!" }` | `UpdateMessage(message_id:
     "123456" message { text: "Hi!" })`

     The special name `*` can be used in the body mapping to define that
     every field not bound by the path template should be mapped to the
     request body.  This enables the following alternative definition of
     the update method:

         service Messaging {
           rpc UpdateMessage(Message) returns (Message) {
             option (google.api.http) = {
               patch: "/v1/messages/{message_id}"
               body: "*"
             };
           }
         }
         message Message {
           string message_id = 1;
           string text = 2;
         }


     The following HTTP JSON to RPC mapping is enabled:

     HTTP | gRPC
     -----|-----
     `PATCH /v1/messages/123456 { "text": "Hi!" }` | `UpdateMessage(message_id:
     "123456" text: "Hi!")`

     Note that when using `*` in the body mapping, it is not possible to
     have HTTP parameters, as all fields not bound by the path end in
     the body. This makes this option more rarely used in practice when
     defining REST APIs. The common usage of `*` is in custom methods
     which don't use the URL at all for transferring data.

     It is possible to define multiple HTTP methods for one RPC by using
     the `additional_bindings` option. Example:

         service Messaging {
           rpc GetMessage(GetMessageRequest) returns (Message) {
             option (google.api.http) = {
               get: "/v1/messages/{message_id}"
               additional_bindings {
                 get: "/v1/users/{user_id}/messages/{message_id}"
               }
             };
           }
         }
         message GetMessageRequest {
           string message_id = 1;
           string user_id = 2;
         }

     This enables the following two alternative HTTP JSON to RPC mappings:

     HTTP | gRPC
     -----|-----
     `GET /v1/messages/123456` | `GetMessage(message_id: "123456")`
     `GET /v1/users/me/messages/123456` | `GetMessage(user_id: "me" message_id:
     "123456")`

     ## Rules for HTTP mapping

     1. Leaf request fields (recursive expansion nested messages in the request
        message) are classified into three categories:
        - Fields referred by the path template. They are passed via the URL path.
        - Fields referred by the [HttpRule.body][google.api.HttpRule.body]. They are passed via the HTTP
          request body.
        - All other fields are passed via the URL query parameters, and the
          parameter name is the field path in the request message. A repeated
          field can be represented as multiple query parameters under the same
          name.
      2. If [HttpRule.body][google.api.HttpRule.body] is "*", there is no URL query parameter, all fields
         are passed via URL path and HTTP request body.
      3. If [HttpRule.body][google.api.HttpRule.body] is omitted, there is no HTTP request body, all
         fields are passed via URL path and URL query parameters.

     ### Path template syntax

         Template = "/" Segments [ Verb ] ;
         Segments = Segment { "/" Segment } ;
         Segment  = "*" | "**" | LITERAL | Variable ;
         Variable = "{" FieldPath [ "=" Segments ] "}" ;
         FieldPath = IDENT { "." IDENT } ;
         Verb     = ":" LITERAL ;

     The syntax `*` matches a single URL path segment. The syntax `**` matches
     zero or more URL path segments, which must be the last part of the URL path
     except the `Verb`.

     The syntax `Variable` matches part of the URL path as specified by its
     template. A variable template must not contain other variables. If a variable
     matches a single path segment, its template may be omitted, e.g. `{var}`
     is equivalent to `{var=*}`.

     The syntax `LITERAL` matches literal text in the URL path. If the `LITERAL`
     contains any reserved character, such characters should be percent-encoded
     before the matching.

     If a variable contains exactly one path segment, such as `"{var}"` or
     `"{var=*}"`, when such a variable is expanded into a URL path on the client
     side, all characters except `[-_.~0-9a-zA-Z]` are percent-encoded. The
     server side does the reverse decoding. Such variables show up in the
     [Discovery
     Document](https://developers.google.com/discovery/v1/reference/apis) as
     `{var}`.

     If a variable contains multiple path segments, such as `"{var=foo/*}"`
     or `"{var=**}"`, when such a variable is expanded into a URL path on the
     client side, all characters except `[-_.~/0-9a-zA-Z]` are percent-encoded.
     The server side does the reverse decoding, except "%2F" and "%2f" are left
     unchanged. Such variables show up in the
     [Discovery
     Document](https://developers.google.com/discovery/v1/reference/apis) as
     `{+var}`.

     ## Using gRPC API Service Configuration

     gRPC API Service Configuration (service config) is a configuration language
     for configuring a gRPC service to become a user-facing product. The
     service config is simply the YAML representation of the `google.api.Service`
     proto message.

     As an alternative to annotating your proto file, you can configure gRPC
     transcoding in your service config YAML files. You do this by specifying a
     `HttpRule` that maps the gRPC method to a REST endpoint, achieving the same
     effect as the proto annotation. This can be particularly useful if you
     have a proto that is reused in multiple services. Note that any transcoding
     specified in the service config will override any matching transcoding
     configuration in the proto.

     Example:

         http:
           rules:
             # Selects a gRPC method and applies HttpRule to it.
             - selector: example.v1.Messaging.GetMessage
               get: /v1/messages/{message_id}/{sub.subfield}

     ## Special notes

     When gRPC Transcoding is used to map a gRPC to JSON REST endpoints, the
     proto to JSON conversion must follow the [proto3
     specification](https://developers.google.com/protocol-buffers/docs/proto3#json).

     While the single segment variable follows the semantics of
     [RFC 6570](https://tools.ietf.org/html/rfc6570) Section 3.2.2 Simple String
     Expansion, the multi segment variable **does not** follow RFC 6570 Section
     3.2.3 Reserved Expansion. The reason is that the Reserved Expansion
     does not expand special characters like `?` and `#`, which would lead
     to invalid URLs. As the result, gRPC Transcoding uses a custom encoding
     for multi segment variables.

     The path variables **must not** refer to any repeated or mapped field,
     because client libraries are not capable of handling such variable expansion.

     The path variables **must not** capture the leading "/" character. The reason
     is that the most common use case "{var}" does not capture the leading "/"
     character. For consistency, all path variables must share the same behavior.

     Repeated message fields must not be mapped to URL query parameters, because
     no client library can support such complicated mapping.

     If an API needs to use a JSON array for request or response body, it can map
     the request or response body to a repeated field. However, some gRPC
     Transcoding implementations may not support this feature.
lifecycle: null
__proto:
    package: google.api
    targetfile: http.proto
    imports: []
    options:
        cc_enable_arenas: "true"
        go_package: google.golang.org/genproto/googleapis/api/annotations;annotations
        java_multiple_files: "true"
        java_outer_classname: HttpProto
        java_package: com.google.api
        objc_class_prefix: GAPI
fields:
    selector:
        type: string
        description: |-
            Selects a method to which this rule applies.

             Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
        __proto:
            number: 1
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: label.HttpRule.selector
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    get:
        type: string
        description: |-
            Maps to HTTP GET. Used for listing and getting information about
             resources.
        __proto:
            number: 2
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: label.HttpRule.get
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    put:
        type: string
        description: Maps to HTTP PUT. Used for replacing a resource.
        __proto:
            number: 3
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: label.HttpRule.put
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    post:
        type: string
        description: Maps to HTTP POST. Used for creating a resource or performing an action.
        __proto:
            number: 4
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: label.HttpRule.post
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    delete:
        type: string
        description: Maps to HTTP DELETE. Used for deleting a resource.
        __proto:
            number: 5
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: label.HttpRule.delete
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    patch:
        type: string
        description: Maps to HTTP PATCH. Used for updating a resource.
        __proto:
            number: 6
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: label.HttpRule.patch
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    custom:
        type: google.api.CustomHttpPattern
        description: |-
            The custom pattern is used for specifying an HTTP method that is not
             included in the `pattern` field, such as HEAD, or "*" to leave the
             HTTP method unspecified for this rule. The wild-card rule is useful
             for services that provide content to Web (HTML) clients.
        __proto:
            number: 8
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: label.HttpRule.custom
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    body:
        type: string
        description: |-
            The name of the request field whose value is mapped to the HTTP request
             body, or `*` for mapping all request fields not captured by the path
             pattern to the HTTP body, or omitted for not having any HTTP request body.

             NOTE: the referred field must be present at the top-level of the request
             message type.
        __proto:
            number: 7
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: label.HttpRule.body
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    response_body:
        type: string
        description: |-
            Optional. The name of the response field whose value is mapped to the HTTP
             response body. When omitted, the entire response message will be used
             as the HTTP response body.

             NOTE: The referred field must be present at the top-level of the response
             message type.
        __proto:
            number: 12
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: label.HttpRule.response_body
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    additional_bindings:
        type: google.api.HttpRule
        description: |-
            Additional HTTP bindings for the selector. Nested bindings must
             not contain an `additional_bindings` field themselves (that is,
             the nesting may only be one level deep).
        __proto:
            number: 11
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: label.HttpRule.additional_bindings
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
