# Getting started

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
./cloud-events-generator generate --input eventSchema.yaml --output ./genfiles/ --language python --binding JSON
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

## What's next

[Using CloudEvents Generator CLI](/cloud-events-generator/clis)