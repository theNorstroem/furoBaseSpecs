{
 "name": "any",
 "type": "Any",
 "description": "Any contains an arbitrary serialized protocol buffer message along with a\n// URL that describes the type of the serialized message. client uses type `ArrayBuffer` for the value field .  https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/any.proto",
 "__proto": {
  "package": "google.protobuf",
  "targetfile": "any.proto",
  "imports": [],
  "options": {}
 },
 "fields": {
  "type_url": {
   "type": "string",
   "description": "",
   "__proto": {
    "number": 1,
    "oneof": ""
   },
   "__ui": null,
   "meta": null,
   "constraints": null
  },
  "value": {
   "type": "bytes",
   "description": "",
   "__proto": {
    "number": 2,
    "oneof": ""
   },
   "__ui": null,
   "meta": null,
   "constraints": null
  }
 }
}