# Transport Bindings

Transport bindings help you pass events between devices, apps, and services.
CloudEvents Generator at this moment supports two transport bindings:
JSON and HTTP. To use any of them, pass it as an argument to the CLI.

## JSON

Event library prepared by CloudEvents Generator can deserialize an event
to a JSON string and serialize an event from a JSON string. For specifics,
see [Usages](/cloud-events-generator/specs).

## HTTP

Event library prepared by CloudEvents Generator can also send and receive events
via HTTP. It supports two of the three mapping modes dictated by [Cloud
Events HTTP Transport Binding Specification](https://github.com/cloudevents/spec/blob/v0.3/http-transport-binding.md):
`structured` and `binary`.

### `structured`

In this mapping mode, your event library first deserializes your **entire**
event into a JSON string and pass it in the HTTP body with the `content-type`
header set to `application/cloudevents+json; charset=UTF-8`. This is the
default mapping mode the event library uses.

On the receiving end, pass the headers (in the form of a dict, object, or
equivalent in your preferred programming language) and the body (a UTF-8 encoded
string) to your event library. The library will reconstruct the event
automatically.

For specifics, see [Usages](/cloud-events-generator/specs).

### `binary`

In this mapping mode, your event library passes only the `data` attribute of
your event (if present) in the HTTP body with all the other attributes,
except for `datacontenttype`, as HTTP headers (with the `ce-` prefix).
`datacontenttype` will be mapped as the `content-type` header.

If your `data` attribute is of a primitive type (`string`, `integer`, `number`,
or `boolean`), your event library sends it in its **string** form. Otherwise,
it will be deserialized into a JSON string in the HTTP body.

On the receiving end, similarly, pass the headers (in the form of a dict, object, or
equivalent in your preferred programming language) and the body (a UTF-8 encoded
string) to your event library. The library will reconstruct the event
automatically.

For specifics, see [Usages](/cloud-events-generator/specs).

**Important**: With the `binary` mode, if you do not have a `datacontenttype`
attribute in your event, the `content-type` header will be left empty. Some
HTTP frameworks/applications depend on this header to parse the HTTP body;
you can always, of course, ask for the body in its raw from and cast it to
a UTF-8 string manually.

**Important**: HTTP headers are case-insensitive. Check your attribute names
before using the binary mode.
