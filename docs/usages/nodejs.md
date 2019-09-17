# Using the Node.js event library

## Installation

You can install the generated Node.js event library locally with

```
npm install PATH
```

## Package organization

Each event in your specification will be added as a submodule in the package.
For example, if you have two events, `basic` and `advanced`, in your generator
specifcation, your package structure look as follows:

```
mypackage/
    Basic/
        ...
    Advanced/
        ...
    index.js
    package.json
    README.md
    ...
```

The events themselves will be parsed into a Node.js class. Additionally,
each `object` type attribute will also be parsed into a Node.js class.
For example, if the `advanced` event has two `object` type attributes,
`name` and `address`, the package structure looks as follows:

```
mypackage/
    Basic/
        # Event class
        Basic.js
    Advanced/
        # Event class
        Advanced.js
        # Data classes
        Name.js
        Address.js
    # Other files
    ...
```

When you import the event library above, it will return an object as follows:

```javascript
// pkg = require('mypackage')
pkg = {
    Basic: {
        Event: [Function Basic]
        DataClasses: {}
    },
    Advanced: {
        Event: [Function Event]
        DataClasses: {
            Name: [Function Name],
            Address: [Function Address]
        }
    }
}
```

To use the `advanced` event, refer to the function in the required package.
Initialize it with an object of attribute names and values (if no default
values are provided). Your event library will validate the input values
as specified.

```javascript
event = new pkg.Advanced.Event({
    id: 'my-id',
    source: 'my-source'
})
```

Similarly, to use the data classes:

```javascript
name = new pkg.Advanced.DataClasses.Name()
address = new pkg.Advanced.DataClasses.Address()
```

**Important**: All the class names are cast in camel case.

## Using the JSON binding

To serialize an event into a JSON string with your event library,
call the `toJSON` method on an event instance:

```javascript
event = new pkg.Advanced.Event()
jsonStr = event.toJSON()
```

To deserialize an event from a JSON string with your event library,
call the `from_JSON` method on an event class:

```javascript
event = pkg.Advanced.Event.fromJSON(json_str)
```

## Using the HTTP binding

To send an event via HTTP, call the `sendHTTP` method. This method takes
two parameters, `url` and `mode`, the latter of which accepts two values,
`structured` (using the structured mapping mode, default) and `binary`
(using the binary mapping mode).

Note: See [Transport Bindings](/cloudevents-generator/bindings/overview)
for more information about bindings and mapping modes.

```javascript
// pkg = require('mypackage')
event = pkg.Advanced.Event()
event.sendHTTP('http://my-endpoint.com')
# or
# event.sendHTTP('http://my-endpoint.com', 'binary')
```

In the background your event library uses [`request`](https://github.com/request/request)
(synchronized) to send the HTTP request. This command returns a
Promise with `response` and `body`.

To receive an event via HTTP, extract headers and payload (body) from your HTTP
framework or application and pass it to the `receiveHTTP` method in the event
class. The headers must be an object and the payload (body) must be a
UTF-8 encoded string.

```javascript
event = pkg.Advanced.Event.receiveHTTP(headers, payload)
```

**Note**: If you are using the binary mode and some of your attribute names are
not in lower case, double check that the headers object you pass
switch cases automatically. HTTP headers are case-insensitive, and your event
library queries the headers object using the names exactly
as you provides.