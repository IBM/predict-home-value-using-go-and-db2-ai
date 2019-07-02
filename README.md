## DB2WML

[![](https://img.shields.io/badge/IBM%20Cloud-powered-blue.svg)](https://bluemix.net)
![Platform](https://img.shields.io/badge/platform-go-lightgrey.svg?style=flat)

## Table of Contents

* [Summary](#summary)
* [Requirements](#requirements)
* [Configuration](#configuration)
* [Run](#run)

<a name="summary"></a>
## Summary

The Web basic starter contains an opinionated set of files for web serving:

- `public/index.html`
- `public/404.html`
- `public/500.html`


All of your `dep` dependencies are stored inside of `Gopkg.toml`.

## Requirements
#### Local Development Setup (optional)

- Install [Go](https://golang.org/dl/)
- Install [dep](https://github.com/golang/dep)

### IBM Cloud Developer Tools (optional)

[IBM Cloud Developer Tools](https://cloud.ibm.com/docs/cli/index.html#overview) simplifies the building, running, and deployment of your application from you local environment to the cloud in containerized environments.

1. Install [IBM Cloud Developer Tools](https://cloud.ibm.com/docs/cli/index.html#step1) on your machine  
2. Install the dev plugin: `ibmcloud plugin install dev`

#### cli-config.yml

The `cli-config.yml` contains the commands that are used by IBM Cloud Developer Tools.  If needed, you can update these to reflect how you want to run your project:
* `test-cmd`: The command to execute tests for the code in the tools container (i.e. `go test ./...`)

* `build-cmd-debug`: The command to build the code and docker image for `DEBUG` mode (i.e. `go build` to ensure that the application compiles cleanly)

* `debug-cmd`: The command to execute debug of the code in the tools container using [delve](https://github.com/derekparker/delve) (i.e. `dlv debug --headless --listen=0.0.0.0:8181`)

### IBM Cloud DevOps (optional)

[![Create Toolchain](https://cloud.ibm.com/devops/graphics/create_toolchain_button.png)](https://cloud.ibm.com/devops/setup/deploy/)

[IBM Cloud DevOps](https://cloud.ibm.com/devops/getting-started) services provides toolchains as a set of tool integrations that support development, deployment, and operations tasks inside IBM Cloud. The **Create Toolchain** button creates a DevOps toolchain and acts as a single-click deploy to IBM Cloud including provisioning all required services.

## Run

#### Configuration

This project contains IBM Cloud specific files that are used to deploy the application as part of an IBM Cloud DevOps flow. The `.bluemix` directory contains files used to define the IBM Cloud toolchain and pipeline for your application.

Credentials are either taken from the VCAP_SERVICES or Kubernetes environment variablea if in IBM Cloud, or from a config file if running locally or on VSIs.

More information about configuration best practices and abstraction of environments can be found in the IBM Cloud [Go Programming Guide](https://cloud.ibm.com/docs/go/configuration.html#configuration).

### Using IBM Cloud Developer Tools

 IBM Cloud Developer Tools makes it easy to compile and run your application if you do not have all of the tools installed on your computer yet. Your application will be compiled with Docker containers. To compile and run your app, run:

```bash
ibmcloud dev build
ibmcloud dev run
```

### Using your local environment

In order for Go applications to run locally, they must be placed in the correct file path. The application must exist in `$GOPATH/src/db2wml`

To run your application locally:

```bash
dep ensure
go run server.go
```

Once the Go toolchain has been installed, you can compile a Go project with:

```bash
go install
```

Your sources will be compiled to your `$GOPATH/bin` directory.

### Application Endpoints

Your application is running at: `http://localhost:8080` in your browser.

- Health endpoint: `/health`
