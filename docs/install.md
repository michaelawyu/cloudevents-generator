# Installation

1. Download the compiled executables of your platform.

    | Platform     | Architecture | Link     | 
    |--------------|--------------|----------|
    | Linux | x64 | [Download](https://github.com/michaelawyu/cloudevents-generator/raw/master/bin/cloudevents-generator-linux-amd64) |
    | Linux | x86 | [Download](https://github.com/michaelawyu/cloudevents-generator/raw/master/bin/cloudevents-generator-linux-386) |
    | macOS (Darwin) | x64 | [Download](https://github.com/michaelawyu/cloudevents-generator/raw/master/bin/cloudevents-generator-darwin-amd64) |
    | macOS (Darwin) | x86 | [Download](https://github.com/michaelawyu/cloudevents-generator/raw/master/bin/cloudevents-generator-darwin-386) |
    | Windows | x64 | [Download](https://github.com/michaelawyu/cloudevents-generator/raw/master/bin/cloudevents-generator-windows-amd64) |
    | Windows | x86 | [Download](https://github.com/michaelawyu/cloudevents-generator/raw/master/bin/cloudevents-generator-windows-amd64) |

2. (Linux or macOS) Make the downloaded file executable:

    ```
    chmod +x [YOUR-FILE-NAME]
    ```

If your platform is not listed above or you prefer building the package from
stratch, clone the project from GitHub and run `go build ./src/` (Go 1.12 required).

## What's next

[Getting started with CloudEvents Generator](/cloudevents-generator/getting_started)