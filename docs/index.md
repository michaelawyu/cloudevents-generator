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

## Documentation

See the topics below to learn more about CloudEvents Generator:

* [Installation](/cloudevents-generator/install)
* [Getting started](/cloudevents-generator/getting_started)
* [CLI](/cloudevents-generator/cli)
* [Specification](/cloudevents-generator/specs)
    * [Populating attributes automatically](/cloudevents-generator/auto)
* [Bindings](/cloudevents-generator/bindings/overview)
    * [JSON](/cloudevents-generator/bindings/overview#JSON)
    * [HTTP](/cloudevents-generator/bindings/overview#HTTP)
* Using generated event libraries
    * [Name Formatting](/cloudevents-generator/usages/name)
    * [Python](/cloudevents-generator/usages/python)
    * [Node.js](/cloudevents-generator/usages/nodejs)
