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

* [Installation](/cloud-events-generator/install)
* [Getting started](/cloud-events-generator/getting_started)
* [CLI](/cloud-events-generator/cli)
* [Specification](/cloud-events-generator/specs)
    * [Populating attributes automatically](/cloud-events-generator/auto)
* [Bindings](/cloud-events-generator/bindings/overview)
    * [JSON](/cloud-events-generator/bindings/overview#JSON)
    * [HTTP](/cloud-events-generator/bindings/overview#HTTP)
* Using generated event libraries
    * [Name Formatting](/cloud-events-generator/usages/name)
    * [Python](/cloud-events-generator/usages/python)
    * [Node.js](/cloud-events-generator/usages/nodejs)
