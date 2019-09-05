# Getting started with Cloud Events Generator

## Introduction

This tutorial help you get a general understanding of how Cloud Events
Generator works by preparing an event library in Python with Cloud Events
Generator.

It takes approximately 5 minutes to complete the tutorial.

Click **Start** to continue.

## Setup

If you are running this tutorial in Cloud Shell, you may skip this step. Click
**Next** to continue.

Otherwise, pick the compiled executable of your platform:

Linux (x64): ./bin/cloud-events-generator-linux-amd64
Linux (x86): ./bin/cloud-events-generator-linux-386
macOS (x64): ./bin/cloud-events-generator-darwin-amd64
macOS (x86): ./bin/cloud-events-generator-darwin-386
Windows (x64): ./bin/cloud-events-generator-windows-amd64.exe
Windows (x86): ./bin/cloud-events-generator-linux-386.exe

## Generating an event library

Run the command below to prepare an event library in Python

```
./bin/cloud-events-generator-linux-amd64 generate \
                                         --input ./examples/basic/events.yaml
                                         --output ./genfiles
                                         --language python
```

**Important**: Use the executable of your platform instead if you are not
running this tutorial in Cloud Shell.

Click **Next** to continue.

## Installing the generated library and creating an event

Install the package with

```
pip3 install -e genfiles/
```

Start Python:

```
python3
```

And run the following Python script:

```
from mypackage import Basic

# Attributes id and time will be auto-populated
# If not specified, attributes source, type, and specversion use their respective default values
event = Basic(data = 'Hello World!')
event.to_JSON()
```

You should see your Cloud Event output in the JSON format.

Click **Next** to continue.

## Congratulations

<walkthrough-conclusion-trophy></walkthrough-conclusion-trophy>

You have finished the tutorial.
