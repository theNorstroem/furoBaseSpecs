{
  "name": "metafield",
  "type": "MetaField",
  "description": "fields of meta info",
  "__proto": {
    "package": "furo",
    "imports": [
    ],
    "targetfile": "meta.proto"
  },
  "fields": {
    "meta": {
      "description": "meta informatioxn of a field",
      "type": "furo.FieldMeta",
      "__proto": {
        "number": 1
      }
    },
    "constraints": {
      "description": "constraints for a field",
      "type": "map<string,furo.FieldConstraint>",
      "__proto": {
        "number": 2
      }
    }
  }
}
