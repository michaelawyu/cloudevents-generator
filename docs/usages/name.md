# Name formatting

At this moment, Cloud Events Generator does not check the format of the names
you give to attributes in your event. The package, however, will attempt to
cast every name to its lower camel case in-memory (in the generated event
library). Your event library to recast the name back to the original form
when using one of the bindings.

For best compatibility, it is strongly recommended that you use the
**lower-case-only, no-connector** case for attribute names.
