# Populating attributes automatically

This document discusses how to set up an auto-populated attribute using
CloudEvents Generator.

The event library prepared by CloudEvents Generator can populate a **`string`**
attribute automatically **if no other values are provided**. To create such an
attribute, you must

* Use one of the supported `format`s in the attribute specification
* Set the `auto` field to `true` in the attribute specification

CloudEvents Generator will ignore all the other optional fields in the
attribute specification, except for `description`, if `auto` and `format`
are present and valid.

## Supported formats

At this moment you may use one of the following `format`s:

* `UUIDv4`: [Universally Unique Identifier/Version 4](https://en.wikipedia.org/wiki/Universally_unique_identifier)
* `RFC3339`: [Date and Time on the Internet: Timestamp](https://tools.ietf.org/html/rfc3339)

The following example describes an auto-populated attribute `data` of the
`UUIDv4` format:

```yaml
attributes:
    data:
        type: string
        format: UUIDv4
        auto: true
```
