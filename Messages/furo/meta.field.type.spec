{
 "name": "metafield",
 "type": "MetaField",
 "description": "fields of meta info",
 "__proto": {
  "package": "furo",
  "targetfile": "meta.proto",
  "imports": [],
  "options": null
 },
 "fields": {
  "meta": {
   "type": "furo.FieldMeta",
   "description": "meta informatioxn of a field",
   "__proto": {
    "number": 1,
    "oneof": ""
   },
   "__ui": null,
   "meta": null,
   "constraints": null
  },
  "constraints": {
   "type": "map\u003cstring,furo.FieldConstraint\u003e",
   "description": "constraints for a field",
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