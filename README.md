## SM - AWS Session Manager CLI: manage your EC2 instances in style!

The SM tool provides a command line interface for interacting with your AWS EC2 instances.
The aim of this project is to make navigation and connection easier.
your EC2 instances using Session Manager, which is a safer option than SSH.

## Requirements

You must install and configure the AWS CLI on your system because the SM uses the .aws/config file to authenticate with AWS.

## Installation

SM is available on Linux, macOS and Windows platforms.
Binary files for these operating systems are available as tarballs on the [release page](https://github.com/filipponova/sm/releases).

## Building From Source

 SM is currently using Golang v1.24.x or above.
 In order to build SM from source you must:

```bash
git clone https://github.com/filipponova/sm.git
cd sm
go mod tidy
go build -o sm
```

## Usage

```bash
 sm list --profile profileName
```
