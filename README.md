# CloudEvents Generator

CloudEvents Generator makes it easier to produce, consume, and collaborate
on [CloudEvents](https://cloudevents.io/). It allows you to
generate a event library of your own using a JSON/YAML specification, which
helps:

* Write and read events idiomatically in the programming language of
your choice with additional support for structured data payload
* Write and read events easier with IDEs
* Verify event attributes and fields automatically
* Send and receive events using one of the supported bindings
* Collaborate on event-driven systems using shared event schemas
* (TO-DO) Perform common tasks in event-driven systems, such as enforcing
exactly-once delivery with persistent storage and signature
preparation/verification, automatically

**Important: This is a (very) early build of CloudEvents Generator with a
limited set of features and potentially some bugs/DX problems.**

In other words, instead of composing a event manually in your
preferred language using composite data structures:

```python
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

CloudEvents Generator helps you achieve this:

```python
# Python snippet
from my_event_package import Order
from my_event_package.Order import Data

# Event library prepared by CloudEvents Generator can populate fields automatically
event = Order()
# Event library prepared by CloudEvents Generator can check input values automatically,
# e.g. making sure that the count in the data payload never goes over 999
event.data = Data(productId='someProduct', count=3, unitPrice=15)

event_str = event.to_JSON()
```

The generated library, along with the schema of your events, also help
every producer and consumer of your events conform to the same standard.

## Installation

CloudEvents Generator is written in Go. To use the package:

1. Download the compiled executables of your platform.

| Platform     | Architecture | Link     | 
|--------------|--------------|----------|
| Linux | x64 | [Download](https://github.com/michaelawyu/cloud-events-generator/raw/master/bin/cloud-events-generator-linux-amd64) |
| Linux | x86 | [Download](https://github.com/michaelawyu/cloud-events-generator/raw/master/bin/cloud-events-generator-linux-386) |
| macOS (Darwin) | x64 | [Download](https://github.com/michaelawyu/cloud-events-generator/raw/master/bin/cloud-events-generator-darwin-amd64) |
| macOS (Darwin) | x86 | [Download](https://github.com/michaelawyu/cloud-events-generator/raw/master/bin/cloud-events-generator-darwin-386) |
| Windows | x64 | [Download](https://github.com/michaelawyu/cloud-events-generator/raw/master/bin/cloud-events-generator-windows-amd64) |
| Windows | x86 | [Download](https://github.com/michaelawyu/cloud-events-generator/raw/master/bin/cloud-events-generator-windows-amd64) |

2. (Linux or macOS) Make the downloaded file executable:

```
chmod +x [YOUR-FILE-NAME]
```

If your platform is not listed above or you prefer building the package from
stratch, clone the project from GitHub and run `go build ./src/` (Go 1.12 required).

## Supported languages and bindings

At this early stage, CloudEvents Generator supports the following languages:

* Python (3+)
* Node.js (8+)

And the following bindings:

* JSON
* HTTP (binary and structured mode only)

## Getting started

**Note: [You can learn more about CloudEvents specification here.](https://github.com/cloudevents/spec/blob/v0.3/spec.md)**

**Note: The following example uses Python 3. For guides and tutorials in other
languages, see [CloudEvents Generator Documentation](https://michaelawyu.github.io/cloud-events-generator/)**

CloudEvents Generator takes a specification in JSON or YAML format as input,
which consists of the schemas of your events and some metadata. The
specification of the earlier example, for instance, is as follows:

```yaml
events:
    order:
        # id, source, specversion, and type are required attributes for every
        # event. If not specified in the schema, CloudEvents Generator
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
    packageName: my_event_package
```

Save the file as `eventSchema.yaml` in the same folder as the CloudEvents
Generator you downloaded earlier. Run CloudEvents Generator to prepare the
package:

```bash
./cloudevents-generator generate --input eventSchema.yaml --output ./genfiles/ --language python --binding JSON
```

A Python package will be generated in `/genfiles`. Install the package with

```bash
pip install -e genfiles/
```

The package is now ready for use. To create an `order` event, for example,
run the Python script below:

```python
from my_event_package import Order
from my_event_package.Order import Data

event = Order()
event.data = Data(productId='myProduct', count=3, unitPrice=20)
event.to_JSON()
```

## Documentation

See [CloudEvents Generator Documentation](https://michaelawyu.github.io/cloud-events-generator/).

## Note

CloudEvents Generator is licensed under Apache 2.0.

CloudEvents Generator uses a subset of the [OpenAPI data model specification syntax](https://github.com/OAI/OpenAPI-Specification)
in its event specification syntax. It also uses some modified templates from
[OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator). OpenAPI
specification and OpenAPI generator are both licensed under Apache 2.0.
