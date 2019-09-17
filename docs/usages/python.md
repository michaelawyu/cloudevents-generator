# Using the Python event library

## Installation

You can install the generated Python event library locally with

```
pip install -e PATH
```

## Package organization

Each event in your specification will be added as a submodule in the package.
For example, if you have two events, `basic` and `advanced`, in your generator
specifcation, your package structure look as follows:

```
mypackage/
    Basic/
        __init__.py
        ...
    Advanced/
        __init__.py
        ...
    setup.py
    requirements.txt
    __init__.py
    README.md
    ...
```

The events themselves will be parsed into a Python class. Additionally,
each `object` type attribute will also be parsed into a Python class.
For example, if the `advanced` event has two `object` type attributes,
`name` and `address`, the package structure looks as follows:

```
mypackage/
    Basic/
        __init__.py
        # Event class
        Basic.py
    Advanced/
        __init__.py
        # Event class
        Advanced.py
        # Data classes
        Name.py
        Address.py
    # Other files
    ...
```

All the event classes are visible at the package level. To use the `advanced`
event class, simply import it from the package:

```python
from mypackage import Advanced
```

Initialize the imported class with attribute names and values (if not default
values are provided). Your event library will validate the input values as
specified.

```python
event = Advanced(id = 'my-id', source='my-source')
```

All the data classes live in their respective submodules:

```python
from mypackage.Advanced import Name, Address
```

Initialize them with their respective names and values as well.

**Important**: All the class names are cast in camel case.

## Using the JSON binding

To serialize an event into a JSON string with your event library,
call the `to_JSON` method on an event instance:

```python
event = Advanced()
json_str = event.to_JSON()
```

To deserialize an event from a JSON string with your event library,
call the `from_JSON` method on an event class:

```python
event = Advanced.from_JSON(json_str)
```

## Using the HTTP binding

To send an event via HTTP, call the `send_http` method. This method takes
two parameters, `url` and `mode`, where the latter accepts two values,
`structured` (using the structured mapping mode, default) and `binary`
(using the binary mapping mode).

Note: See [Transport Bindings](/cloudevents-generator/bindings/overview)
for more information about bindings and mapping modes.

```python
event = Advanced()
event.send_http(url='http://my-endpoint.com')
# or
# event.send_http(url='http://my-endpoint.com', mode='binary')
```

In the background your event library uses `requests` (synchronized) to send
the HTTP request. This command returns a [`requests` response object](https://2.python-requests.org/en/master/user/quickstart/#response-content).

To receive an event via HTTP, extract headers and payload (body) from your HTTP
framework or application and pass it to the `receive_http` method in the event
class. The headers must be a `dikt`-like object and the payload (body) must be a
UTF-8 encoded string.

```python
event = Advanced.receive_http(headers=headers, payload=body)
```

**Note**: If you are using the binary mode and some of your attribute names are
not in lower case, double check that the headers `dikt`-like object you pass
switch cases automatically. HTTP headers are case-insensitive, and your event
library queries the headers `dikt`-like object using the names exactly
as you provides.