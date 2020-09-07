{
 "name": "string_option_property",
 "type": "StringOptionProperty",
 "description": "String type to use in property",
 "__proto": {
  "package": "furo",
  "targetfile": "property.proto",
  "imports": [],
  "options": null
 },
 "fields": {
  "display_name": {
   "type": "string",
   "description": "String representation of val",
   "__proto": {
    "number": 1,
    "oneof": ""
   },
   "__ui": null,
   "meta": {
    "default": "",
    "hint": "",
    "label": "",
    "options": null,
    "readonly": true,
    "repeated": false,
    "typespecific": null
   },
   "constraints": {}
  },
  "id": {
   "type": "string",
   "description": "The value, Id is used to make working with data-inputs easier",
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