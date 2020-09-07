{
 "name": "big_decimal",
 "type": "BigDecimal",
 "description": "A BigDecimal is defined by two values: an arbitrary precision integer and a 32-bit integer scale. The value of the BigDecimal is defined to be unscaledValue*10^{-scale}.",
 "__proto": {
  "package": "furo",
  "targetfile": "bigdecimal.proto",
  "imports": [],
  "options": null
 },
 "fields": {
  "display_name": {
   "type": "string",
   "description": "String representation of BigDecimal entity",
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
  "scale": {
   "type": "int32",
   "description": "If zero or positive, the scale is the number of digits to the right of the decimal point. If negative, the unscaled value of the number is multiplied by ten to the power of the negation of the scale. For example, a scale of -3 means the unscaled value is multiplied by 1000.",
   "__proto": {
    "number": 2,
    "oneof": ""
   },
   "__ui": null,
   "meta": {
    "default": "",
    "hint": "",
    "label": "",
    "options": null,
    "readonly": false,
    "repeated": false,
    "typespecific": null
   },
   "constraints": {}
  },
  "int_val": {
   "type": "int64",
   "description": "The integer value of the BigDecimal",
   "__proto": {
    "number": 3,
    "oneof": ""
   },
   "__ui": null,
   "meta": {
    "default": "",
    "hint": "",
    "label": "",
    "options": null,
    "readonly": false,
    "repeated": false,
    "typespecific": null
   },
   "constraints": {}
  }
 }
}