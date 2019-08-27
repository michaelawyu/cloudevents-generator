# Cloud Events Generator

Cloud Events Generator makes it easier to produce, consume, and collaborate
on [Cloud Events](https://github.com/cloudevents/spec). It allows you to
generate a event library of your own using a JSON/YAML specification, which
helps:

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
limited set of features and potentially some bugs/DX problems.**

## Documentation

See the topics below to learn more about Cloud Events Generator:

* [Installation](/cloud-events-generator/installation)
* [Getting started](/cloud-events-generator/getting_started)
* [CLI](/cloud-events-generator/cli)
* [Specification](/cloud-events-generator/specs)
    * [Populating attributes automatically](/cloud-events-generator/auto)
* [Bindings](/cloud-events-generator/bindings)
    * [JSON](/cloud-events-generator/bindings#JSON)
    * [HTTP](/cloud-events-generator/bindings#HTTP)
* Using generated event libraries
    * [Name Formatting](/cloud-events-generator/usages/name.md)
    * [Python](/cloud-events-generator/usages/python)
    * [Node.js](/cloud-events-generator/usages/nodejs)
