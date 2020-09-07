{
 "name": "fieldoption",
 "type": "Fieldoption",
 "description": "Metas for a field",
 "__proto": {
  "package": "furo",
  "targetfile": "meta.proto",
  "imports": [
   "google/protobuf/any.proto"
  ],
  "options": null
 },
 "fields": {
  "list": {
   "type": "google.protobuf.Any",
   "description": "a list with options, use furo.optionitem or your own",
   "__proto": {
    "number": 1,
    "oneof": ""
   },
   "__ui": {
    "component": "",
    "flags": [
     "full",
     "condensed"
    ],
    "no_init": false,
    "no_skip": false
   },
   "meta": {
    "default": "",
    "hint": "",
    "label": "",
    "options": {},
    "readonly": false,
    "repeated": true,
    "typespecific": null
   },
   "constraints": null
  },
  "flags": {
   "type": "string",
   "description": "Add flags for your field. This can be something like \"searchable\". \n//The flags can be used by generators, ui components,...\n",
   "__proto": {
    "number": 2,
    "oneof": ""
   },
   "__ui": {
    "component": "",
    "flags": [],
    "no_init": false,
    "no_skip": false
   },
   "meta": {
    "default": "",
    "hint": "optional flags",
    "label": "flags",
    "options": {},
    "readonly": false,
    "repeated": true,
    "typespecific": null
   },
   "constraints": null
  }
 }
}