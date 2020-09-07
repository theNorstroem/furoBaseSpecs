{
 "name": "money",
 "type": "Money",
 "description": "Represents an amount of money with its currency type. https://github.com/googleapis/googleapis/blob/master/google/money.proto",
 "__proto": {
  "package": "google.type",
  "targetfile": "money.proto",
  "imports": [],
  "options": null
 },
 "fields": {
  "display_name": {
   "type": "string",
   "description": "String representation of money entity",
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
  "currency_code": {
   "type": "string",
   "description": "The 3-letter currency code defined in ISO 4217.",
   "__proto": {
    "number": 2,
    "oneof": ""
   },
   "__ui": null,
   "meta": {
    "default": "",
    "hint": "",
    "label": "Währungscode",
    "options": null,
    "readonly": false,
    "repeated": false,
    "typespecific": null
   },
   "constraints": {}
  },
  "units": {
   "type": "int64",
   "description": "The whole units of the amount.",
   "__proto": {
    "number": 3,
    "oneof": ""
   },
   "__ui": null,
   "meta": {
    "default": "",
    "hint": "",
    "label": "Ganzahliger Währungsbetrag",
    "options": null,
    "readonly": false,
    "repeated": false,
    "typespecific": null
   },
   "constraints": {}
  },
  "nanos": {
   "type": "int64",
   "description": "Number of nano (10^-9) units of the amount. For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.",
   "__proto": {
    "number": 4,
    "oneof": ""
   },
   "__ui": null,
   "meta": {
    "default": "",
    "hint": "",
    "label": "Nanos",
    "options": null,
    "readonly": false,
    "repeated": false,
    "typespecific": null
   },
   "constraints": {}
  }
 }
}