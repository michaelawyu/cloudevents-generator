# Cloud Events Generator

Cloud Events Generator makes it easier to produce, consume, and collaborate
on Cloud Events. It allows you to generate a event library of your own using
a JSON/YAML specification, which helps:

* Write and read Cloud Events idiomatically in the programming language of
your choice with additional support for structured data payload
* Write and read Cloud Events easier with IDEs
* Verify Cloud Event attributes and fields automatically
* Send and receive Cloud Events using one of the supported bindings
* Collaborate on event-driven systems using shared event schemas
* (TO-DO) Perform common tasks in event-driven systems, such as enforcing
exactly-once delivery with persistent storage and signature
preparation/verification, automatically

**Important: This is a (very) early build of Cloud Events Generator with a
limited set of features and probably some number of bugs/DX problems.**

In other words, instead of writing code as follows:

```
# Python snippet
from datetime import datetime
import json
import uuid

event = {}
event['id'] = str(uuid.uuid4())
event['source'] = '/myservice/order'
event['specversion'] = '0.3'
event['type'] = 'com.example.order'
event['time'] = datetime.utcnow().isoformat('T') + 'Z'
event['data'] = {
    'productId': 'someProduct',
    'count': 3,
    'unitPrice': 15
}

event_str = json.dumps(event)
```

Cloud Events Generator does this:

```
# Python snippet
from my_event_package import Order
from my_event_package.Order import Data

# Event library prepared by Cloud Events Generator can populate fields automatically
event = Order()
# Event library prepared by Cloud Events Generator can check input values automatically,
# e.g. making sure that the count in the data payload never goes over 999
event.data = Data(productId='someProduct', count=3, unitPrice=15)

event_str = event.to_JSON()
```

## Installation

Cloud Events Generator is written in Go. To use the package:

1. Download the compiled executables of your platform.

| Platform     | Link     | 
|--------------|----------|
| Linux (x64)  | [Download](https://github.com/michaelawyu/cloud-events-generator/raw/master/bin/cloud-events-generator-linux-amd64) |
| Linux (x86)  | [Download](https://github.com/michaelawyu/cloud-events-generator/raw/master/bin/cloud-events-generator-linux-386) |
| macOS (Darwin) (x64) | [Download](https://github.com/michaelawyu/cloud-events-generator/raw/master/bin/cloud-events-generator-darwin-amd64) |
| macOS (Darwin) (x86) | [Download](https://github.com/michaelawyu/cloud-events-generator/raw/master/bin/cloud-events-generator-darwin-386) |
| Windows (x64) | [Download](https://github.com/michaelawyu/cloud-events-generator/raw/master/bin/cloud-events-generator-windows-amd64) |
| Windows (x86) | [Download](https://github.com/michaelawyu/cloud-events-generator/raw/master/bin/cloud-events-generator-windows-amd64) |

2. (Linux or macOS) Make the downloaded file executable:

```
chmod +x [YOUR-FILE-NAME]
```

If your platform is not listed above or you prefer building the package from
stratch, clone the project from GitHub and run `go build ./src/` (Go 1.12 required).

## Supported languages and bindings

At this early stage, Cloud Events Supports the following languages:

* Python (3+)
* Node.js (8+)

And the following bindings:

* JSON
* HTTP

## Getting started

**Note: [You can learn more about Cloud Events specification here.](https://github.com/cloudevents/spec/blob/v0.3/spec.md)**

**Note: The following example uses Python 3. For guides and tutorials in other
languages, see [Cloud Events Generator Documentation](https://michaelawyu.github.io/cloud-events-generator/)**

Cloud Events Generator takes a schema of your events in JSON or YAML as input.
The schema of the earlier example, is as follows:

```
events:
    order:
        # id, source, specversion, and type are required attributes for every
        # Cloud Event. If not specified in the schema, Cloud Events Generator
        # will apply a default specification automatically.
        source:
            type: string
            default: '/myservice/order'
        type:
            type: string
            default: 'com.example.order'
        data:
            type: object
            properties:
                productId:
                    type: string
                count:
                    type: integer
                    maximum: 999
                    minimum: 1
                unitPrice:
                    type: number
metadata:
    packageName: mypackage
```

Save the file as `eventSchema.yaml` in the same folder as the Cloud Events
Generator you downloaded earlier. Run Cloud Events Generator to prepare the
package:

```
./cloud-events-generate generate --input eventSchema.yaml --output ./genfiles/ --language python --binding JSON
```

A Python package will be generated in `/genfiles`. Install the package with

```
pip install -e genfiles/
```

The package is now ready for use. To create an `order` event, for example,
run the Python script below:

```
from mypackage import Order
from mypackage.Order import Data

event = Order()
event.data = Data(productId='myProduct', count=3, unitPrice=20)
event.toJSON()
```

## Documentation

See [Cloud Events Generator Documentation](https://michaelawyu.github.io/cloud-events-generator/).

## Note

Cloud Events Generator is licensed under Apache 2.0.

Cloud Events Generator uses a subset of [OpenAPI data model specification syntax](https://github.com/OAI/OpenAPI-Specification)
in its event specification syntax. It also uses some modified templates from
[OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator). OpenAPI
specification and OpenAPI generator are both licensed under Apache 2.0.
