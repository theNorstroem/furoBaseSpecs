{
 "name": "fieldmeta",
 "type": "FieldMeta",
 "description": "Metas for a field",
 "__proto": {
  "package": "furo",
  "targetfile": "meta.proto",
  "imports": [
   "google/protobuf/any.proto"
  ],
  "options": {}
 },
 "fields": {
  "label": {
   "type": "string",
   "description": "The label",
   "__proto": {
    "number": 1,
    "oneof": ""
   },
   "__ui": null,
   "meta": {
    "default": "",
    "hint": "Also used for input-fields",
    "label": "Label",
    "options": null,
    "readonly": false,
    "repeated": false,
    "typespecific": null
   },
   "constraints": null
  },
  "hint": {
   "type": "string",
   "description": "A hint",
   "__proto": {
    "number": 2,
    "oneof": ""
   },
   "__ui": null,
   "meta": {
    "default": "",
    "hint": "Also used for input-fields",
    "label": "Hint",
    "options": null,
    "readonly": false,
    "repeated": false,
    "typespecific": null
   },
   "constraints": null
  },
  "default": {
   "type": "string",
   "description": "The default value as JSON string",
   "__proto": {
    "number": 3,
    "oneof": ""
   },
   "__ui": null,
   "meta": {
    "default": "",
    "hint": "",
    "label": "Default value",
    "options": null,
    "readonly": false,
    "repeated": false,
    "typespecific": null
   },
   "constraints": null
  },
  "readonly": {
   "type": "bool",
   "description": "readonly",
   "__proto": {
    "number": 4,
    "oneof": ""
   },
   "__ui": null,
   "meta": {
    "default": "",
    "hint": "",
    "label": "readonly",
    "options": null,
    "readonly": false,
    "repeated": false,
    "typespecific": null
   },
   "constraints": null
  },
  "repeated": {
   "type": "bool",
   "description": "repeated",
   "__proto": {
    "number": 5,
    "oneof": ""
   },
   "__ui": null,
   "meta": {
    "default": "",
    "hint": "",
    "label": "repeated",
    "options": null,
    "readonly": false,
    "repeated": false,
    "typespecific": null
   },
   "constraints": null
  },
  "options": {
   "type": "furo.Fieldoption",
   "description": "Fieldoptions",
   "__proto": {
    "number": 6,
    "oneof": ""
   },
   "__ui": null,
   "meta": {
    "default": "",
    "hint": "",
    "label": "options",
    "options": null,
    "readonly": false,
    "repeated": false,
    "typespecific": null
   },
   "constraints": null
  },
  "typespecific": {
   "type": "google.protobuf.Any",
   "description": "Put in type specific metas for your fields here",
   "__proto": {
    "number": 7,
    "oneof": ""
   },
   "__ui": {
    "component": "",
    "flags": null,
    "no_init": false,
    "no_skip": false
   },
   "meta": {
    "default": "",
    "hint": "",
    "label": "typespecific meta",
    "options": null,
    "readonly": false,
    "repeated": false,
    "typespecific": null
   },
   "constraints": null
  }
 }
}