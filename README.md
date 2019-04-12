# CloudWatch to Slack AWS Lambda Func

This is a basic [Lambda](https://aws.amazon.com/lambda/) function which receives a [CloudWatch Event](https://github.com/aws/aws-lambda-go/blob/master/events/cloudwatch_events.go) and posts the details to a Slack channel.

It was build with Go 1.12.

## Getting started

To install Go, see the [install instructions](https://golang.org/doc/install#install).

This Lambda uses [Modules](https://github.com/golang/go/wiki/Modules) which were introduced in 1.11.

### Fetching the modules

To fetch the modules, run:

```cmd
go get ./...
```

### Building the executable

We want to target a Linux environment for the Lambda's OS, so we need to prepend our build command with `GOOS=linux`. 

To build, run:

```
GOOS=linux go build -o build/main main.go
```

### Packaging for Upload

Lambda wants a .zip archive

### Makefile

The above build or packaging steps have been simplified in the [Makefile](Makefile) by running `make build` or `make pack`.
