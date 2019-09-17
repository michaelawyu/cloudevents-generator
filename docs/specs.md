# Generator Specification

This document discusses the structure (format) of CloudEvents Generator
specification.

A CloudEvents Generator specification consists of two parts, `events` and
`metadata`:

* `events` includes the schemas of your events, where each event may
have a variety of attributes of different types
* `metadata` includes supplementary information for package preparation, such
as the name of your package, its version, and more.

You may write CloudEvents Generator specification in
[JSON](https://www.json.org/) or [YAML](https://yaml.org/).

## `events`

`events` is a mapping (hash/dictionary) that takes the names of events as keys
and the specification of events as values. Each event may have two
fields, `attributes` and `required`:

* `attributes` is a mapping that takes the names of attributes as keys and the
specification of attributes as values.
* `required` is an array (list) that specifies required attributes

Note: For compatibility reasons, CloudEvents Generator attempts to format
names of events and attributes in [lower camel case](https://en.wikipedia.org/wiki/Camel_case).
This formatting works exclusively with in-memory structures; the library
will automatically cast the names back to their original forms when used with
a binding.

For example, a event named `BasicEvent` with [only required attributes
from [CloudEvents Specification](https://github.com/cloudevents/spec/blob/v0.3/spec.md#required-attributes)
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
            # Every event should have the following attributes
            - id
            - source
            - type
            - specversion
```

Each attribute specification must have a field, `type`, which specifies
the data type of the attribute. CloudEvents Generator supports the following
types:

* `string`: a string of text
* `integer`: an integer
* `number`: a floating-point number
* `boolean`: a boolean
* `object`: a mapping of keys and values
* `array`: a collection of `string`, `integer`, `number`, `boolean`, or `object`

Attribute of a specific `type` may have other required or optional fields
for CloudEvents Generator to use.

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
| `auto`  | Optional | If set to `true`, the attribute will be populated automatically in accordance with the `format` field at the time of event creation. See [Auto Population](/auto) for more information. |
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

For another example, here is a `string` attribute `data` that takes only values
`foo` and `bar`, with the `enum` field specified:

```yaml
attributes:
    data:
        type: string
        enum:
            - foo
            - bar
        description: a string attribute that takes only value foo and bar
```

**Important**: CloudEvents Generator will ignore all the other optional fields,
except for `description`, if `enum` is present.

### `integer` type attributes

The specification of an `integer` type attribute may have the following fields
in addition to `type`:

| Field        | Type     | Note   |
|--------------|----------|--------------|
| `maximum`  | Optional | The maximum of the integer. |
| `minimum`  | Optional | The minmum of the integer. |
| `exclusiveMaximum` | Optional | If set to `true`, uses exclusive maximum (<) instead of the default inclusive maximum (<=). |
| `exclusiveMinimum`  | Optional | If set to `true`, uses exclusive minimum (>) instead of the default inclusive minimum (>=). |
| `description`  | Optional | The description of the attribute. |
| `enum`  | Optional | An array (list) of integer values the attribute can take. |
| `default`  | Optional | The default value of the attribute. |

The following example describes an `integer` attribute `data` with a maximum of
10 (inclusive), a minimum of 3 (exclusive), and a default value of `7`:

```yaml
attributes:
    data:
        type: integer
        maximum: 10
        minimum: 3
        exclusiveMinimum: true
        default: 7
        description: an integer attribute that is <= 10 and > 3.
```

**Important**: Similar to `string` attributes, CloudEvents Generator will
ignore all the other optional fields in the specification of an `integer`
attribute, except for `description`, if `enum` is present.

### `number` type attributes

The specification of a `number` type attribute may have the following fields
in addition to `type`:

| Field        | Type     | Note   |
|--------------|----------|--------------|
| `maximum`  | Optional | The maximum of the number. |
| `minimum`  | Optional | The minmum of the number. |
| `exclusiveMaximum` | Optional | If set to `true`, uses exclusive maximum (<) instead of the default inclusive maximum (<=). |
| `exclusiveMinimum`  | Optional | If set to `true`, uses exclusive minimum (>) instead of the default inclusive minimum (>=). |
| `description`  | Optional | The description of the attribute. |
| `enum`  | Optional | An array (list) of number values the attribute can take. |
| `default`  | Optional | The default value of the attribute. |

The following example describes a `number` attribute `data` with a maximum of
10.5 (inclusive), a minimum of 3.2 (exclusive), and a default value of `7.1`:

```yaml
attributes:
    data:
        type: number
        maximum: 10.5
        minimum: 3.2
        exclusiveMinimum: true
        default: 7.1
        description: a number attribute that is <= 10 and > 3.
```

**Important**: Similar to `string` attributes, CloudEvents Generator will
ignore all the other optional fields in the specification of a `number`
attribute, except for `description`, if `enum` is present.

### `boolean` type attributes

The specification of a `boolean` type attribute may have the following fields
in addition to `type`:

| Field        | Type     | Note   |
|--------------|----------|--------------|
| `description`  | Optional | The description of the attribute. |

The following example describes a `boolean` attribute:

```yaml
attributes:
    data:
        type: boolean
        description: a boolean attribute
```

### `object` type attributes

The specification of a `number` type attribute may have the following fields
in addition to `type`:

| Field        | Type     | Note   |
|--------------|----------|--------------|
| `properties`  | **Required** | The properties (attributes) of this attribute. |
| `required`  | Optional | An array (list) of names of required properties (attributes) of this attribute. |
| `description`  | Optional | The description of the attribute. |

The `properties` of an `object` attribute may include one or more `string`,
`integer`, `number`, `boolean`, and `object` attributes.

The follwing example describes an `object` type attribute `data` that features
the names (first and last) of an individual:

```yaml
attributes:
    data:
        type: object
        properties:
            # A string attribute
            firstName:
                type: string
                maxLength: 20
                minLength: 1
            # Another string attribute
            lastName:
                type: string
                maxLength: 20
                minLength: 1
        required:
            - firstName
            - lastName
        description: the names of an individual
```

As mentioned earlier, you may use nested `object`s as well:

```yaml
attributes:
    data:
        type: object
        properties:
            name:
                type: object
                properties:
                    firstName:
                        type: string
                    lastName:
                        type: string
            age:
                type: integer
        description: personal info of an individual
```

### `array` type attributes

The specification of an `array` type attribute may have the following fields
in addition to `type`:

| Field        | Type     | Note   |
|--------------|----------|--------------|
| `items`  | **Required** | The specification of array items. |
| `maxItems`  | Optional | The maximum number of items this array attribute may have. |
| `minItems`  | Optional | The minimum number of items this array attribute may have. |
| `description`  | Optional | The description of the attribute. |

The `items` of an `array` type attribute may include **exactly one** `string`,
`integer`, `number`, `boolean`, or `object` attribute without the namee.
**Nested arrays are not supported**.

The follwing example describes an `array` type attribute `data` that features
1-3 `object` type items:

```yaml
attributes:
    data:
        type: array
        items:
            # Note that no name is required
            type: object
            properties:
                productId:
                    type: string
                count:
                    type: integer
                unitPrice:
                    type: number
        maxItems: 3
        minItems: 1
        description: an array of orders
```

## `metadata`

`metadata` is a mapping that specifies the name of your package, its version,
and more. It may have the following fields:

| Field        | Type     | Note   |
|--------------|----------|--------------|
| `packageName`  | Optional | The name of your package. If left empty, CloudEvents Generator uses the name `mypackage`. |
| `version`  | Optional | The version of your package. If left empty, CloudEvents Generator uses `0.0.1`. |
| `description`  | Optional | The description of the package. |
| `contact`  | Optional | The contact of your package. |
| `url`  | Optional | The url of the package. |

## Conformity to CloudEvents specs

At this moment, CloudEvents Generator checks only if each of your event
has the required fields CloudEvents Specification dictates. The generator will
add a default specification if one of the required fields is not present. You
can learn more about the default specification here: [YAML](https://github.com/michaelawyu/cloudevents-generator/blob/master/specs/default.yaml)
/[JSON](https://github.com/michaelawyu/cloudevents-generator/blob/master/specs/default.json).
In other words,  CloudEvents Generator does not perform additional validation
regarding comformity to CloudEvents Specificatiion; it is strongly recommended
that you check with the specificatiion yourself before preparing your event
library with CloudEvents Generator.

## What's next

[Transport Bindings](/cloudevents-generator/bindings/overview)
