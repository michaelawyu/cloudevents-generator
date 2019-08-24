# Specification

This document discusses the structure (format) of Cloud Events Generator
specification.

A Cloud Events Generator specification consists of two parts, `events` and
`metadata`:

* `events` includes the schemas of your Cloud Events, where each event may
have a variety of attributes of different types
* `metadata` includes supplementary information for package preparation, such
as the name of your package, its version, and more.

You may write Cloud Events Generator specification in
[JSON](https://www.json.org/) or [YAML](https://yaml.org/).

## `events`

`events` is a mapping (hash/dictionary) that takes the names of events as keys
and the specification of events as values. Each event may have two
fields, `attributes` and `required`:

* `attributes` is a mapping that takes the names of attributes as keys and the
specification of attributes as values.
* `required` is an array (list) that specifies required attributes

For example, a Cloud Event named `BasicEvent` with [only required attributes
from Cloud Events Specification](https://github.com/cloudevents/spec/blob/v0.3/spec.md#required-attributes)
may be described as follows:

```yaml
events:
    BasicEvent:
        attributes:
            id:
                ... # Specification of the id attribute
            source:
                ... # Specification of the source attribute
            type:
                ... # Specification of the type attribute
            specversion:
                ... # Specification of the specversion attribute
        required:
            # Every Cloud Event should have the following attributes
            - id
            - source
            - type
            - specversion
```

Each attribute specification must have a field, `type`, which specifies
the data type of the attribute. Cloud Events Generator supports the following
types:

* `string`: a string of text
* `integer`: an integer
* `number`: a floating-point number
* `boolean`: a boolean
* `object`: a mapping of keys and values
* `array`: a collection of `string`, `integer`, `number`, `boolean`, or `object`

Attribute of a specific `type` may have other required or optional fields
for Cloud Events Generator to use.

### `string` type attributes

The specification of a `string` type attribute may have the following fields
in addition to `type`:

| Field        | Type     | Note   |
|--------------|----------|--------------|
| `maxLength`  | Optional | The maximum length of the string. |
| `minLength`  | Optional | The minmum length of the string. |
| `pattern` | Optional | A regular expression that matches the string. |
| `enum`  | Optional | An array (list) of string values the attribute can take. |
| `description`  | Optional | The description of the attribute. |
| `format`  | Optional | The format of the attribute. |
| `auto`  | Optional | If set to `true`, the attribute will be populated automatically in accordance to the `format` field at the time of event creation. See [Auto Population](/auto) for more information. |
| `default`  | Optional | The default value of the attribute. |

The following example describes an lowercase letter only `string` attribute
`data` with a maximum length of `5`, a minimum length of `1` and a default
value of `test`:

```yaml
attributes:
    data:
        type: string
        maxLength: 5
        minLength: 1
        pattern: '^[a-z]*$'
        default: 'test'
        description: an lowercase letter only string attribute with a variable length between 1 and 5
```

For another example, here is a `string` attribute `data` that can only take values
`foo` and `bar`, using the `enum` attribute:

```yaml
attributes:
    data:
        type: string
        enum:
            - foo
            - bar
        description: a string attribute that takes only value foo and bar
```

Important: **Cloud Events Generator will ignore all the other optional fields,
except for `description`, if `enum` is present.**

### `integer` type attributes

### `number` type attributes

### `boolean` type attributes

### `object` type attributes

### `array` type attributes
