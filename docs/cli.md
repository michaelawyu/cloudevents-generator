# CLI

CloudEvents Generator CLI provides the following commands:

| Command      | Note     |
|--------------|----------|
| `version`, `ver` | View the version number. |
| `generate`, `gen` | Generate an event library. |
| `help`, `h` | View the help message. |

## `generate` subcommand

The `generate` subcommand uses the following command:

| Flag        | Type     | Note    |
|--------------|----------|--------------|
| `--input PATH`, `-i PATH` | Required | The `PATH` to the input CloudEvents Generator specification. |
| `--output PATH`, `-o PATH` | Required | The `PATH` where the generated event library is saved. |
| `--language LANG`, `-lang LANG` | Required | The language to use. |
| `--binding BIND`, `-bind BIND` | Optional | The transport binding to use. If not specified, JSON binding is enabled. |
| `--verbose`, `-v` | Optional | Enables verbose logging. |

## What's next

[CloudEvents Generator Specification](/cloudevents-generator/specs)