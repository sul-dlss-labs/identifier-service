# DLSS Identifiers API for AWS Applications
This is a prototype that at the moment merely provides random DRUIDs.

# DLSS Identifiers API & Micro-service

## Guiding Questions around Identifier Service idea

1. Why Id Service instead of calling SURI directly?
    - Better permissions management (currently, SURI uses shared passwords across local codebases).
    - Separate the specification from the implementation / make connections more modular as SURI may (or may not?) change.
    - Set up a pattern for building APIs as AWS work continues
    - Capture better analytics for AWS querying SURI or future SDR components via these proxy APIs.
2. Identifier Service should be in AWS?
    - Yes, if we can support in production. Given TACO is in AWS, Identifier Service in AWS seems a good step for this.
3. Authentication for the Identifier Service
    - TBD: method and how we can improve this (but probably having authentication attached to a to-be-created AWS IAM
      Service Role)
    - Proposal of using API keys (à la GitHub) also proposed
    - Protection of SURI: Using SSL?

## Requirements

### For TACO / Hydrox:
- Create a new DRUID on a single call for single identifier action
- Support later additions (but no current implementation) for:
  - Request a list of DRUIDs
  - Use other namespaces for identifiers (not DRUIDs)
  - Do not support ability to create a new identifier namespace or identifier spec via this API (should occur in the SURI implementation)

### For APIs generally:
- Version the API
- Elegant solution for retries / connection issues when calling SURI


## Go Local Development Setup

1. Install go (grab binary from here or use `brew install go` on Mac OSX).
2. Setup your Go workspace (where your Go code, binaries, etc. are kept together. See some helpful documentation here on the Go concept of workspaces: https://github.com/alco/gostart#1-go-tool-is-only-compatible-with-code-that-resides-in-a-workspace.):
      ```bash
      $ mkdir -p ~/go
      $ export GOPATH=~/go
      $ export PATH=~/go/bin:$PATH
      $ cd ~/go
      ```
      Your Go code repositories will reside within `~/go/src/...` in the `$GOPATH`. Name these paths to avoid library clash, for example identifier-service Go code could be in `~/go/src/github.com/sul-dlss-labs/identifier-service`. This should be where your Github repository resides too.
3. In order to download the project code to `~/go/src/github.com/sul-dlss-labs/identifier-service`, from any directory in your ``$GOPATH`, run:
    ```bash
    $ go get github.com/sul-dlss-labs/identifier-service
    ```
4. Handle Go project dependencies with the Go `dep` package:
    * Install Go Dep via `brew install dep` then `brew upgrade dep` (if on Mac OSX).
    * If your project's `Gopkg.toml` and `Gopkg.lock` have not yet been populated, you need to add an inferred list of your dependencies by running `dep init`.
    * If your project has those file populated, then make sure your dependencies are synced via running `dep ensure`.
    * If you need to add a new dependency, run `dep ensure -add github.com/pkg/errors`. This should add the dependency and put the new dependency in your `Gopkg.*` files.

5. [Localstack](docs/localstack.md)

6. Configuration
    * You can override the default configuration by using environment variables. See the list of the environment variables in [config](config/config.go)
    * If you want to report failures to HoneyBadger set `HONEYBADGER_API_KEY=aaaaaaaa`

## Running the Go Code locally without a build

```shell
$ cd cmd/app
$ APP_PORT=1231 go run main.go
```

## Building the server binary

### Building for Docker
```shell
$ docker build -t identifier-service .
$ docker run -p 8080:8080 identifier-service
```

### Build for the local OS
This will create a binary in your path that you can then run the application with.

```shell
$ go build -o appd cmd/app/main.go
```

## Running the server

Now start the API server:
```shell
$ ./appd
```

Then you can interact with it using `curl`:
```shell
$ curl -X POST http://localhost:8080/v1/identifiers/druids
# => [{"id":"testing"}]
```

## Testing
The unit tests have no external dependencies and can be run like so:
```shell
$ go test -v ./... -short
```

The integration test depends on the binary running.  Once these conditions are met you can run the integration tests using:

```shell
$ go test test/integration_test.go
```


## API Code Structure

We use `go-swagger` to generate the API code within the `generated/` directory.

We connect the specification-generated API code to our own handlers defined with `handlers/`. Handlers are where we add our own logic for processing requests.

Our handlers and the generated API code is connected within `main.go`, which is the file to start the API.

### Install Go-Swagger

The API code is generated from `swagger.yml` using `go-swagger` library. As this is not used in the existing codebase anywhere currently, you'll need to install the `go-swagger` library before running these commands (commands for those using Mac OSX):

```shell
$ brew tap go-swagger/go-swagger
$ brew install go-swagger
$ brew upgrade go-swagger
```

This should give you the `swagger` binary command in your $GOPATH and allow you to manage versions better (TBD write this up). The version of your go-swagger binary is **0.13.0** (run `swagger version` to check this).

### Nota Bene on go-swagger install from source

If instead of brew, you decide to install go-swagger from source (i.e. `go get -u github.com/go-swagger/go-swagger/cmd/swagger`), you will be running the library at its Github `dev` branch. You will need to change into that library (`cd $GOPATH/src/github.com/go-swagger/go-swagger`) and checkout from Github the latest release (`git checkout tags/0.13.0`). Then you will need to run `go install github.com/go-swagger/go-swagger/cmd/swagger` to generate the go-swagger binary in your `$GOPATH/bin`.

### To generate the API code

There appears to be no best way to handle specification-based re-generation of the `generated/` API code, so the following steps are recommended:

```shell
$ git rm -rf generated/
$ mkdir generated
$ swagger generate server -t generated --exclude-main
```

### Non-generated code

Anything outside of `generated/` is our own code and should not be touched by a regeneration of the API code from the Swagger specification.

## SWAGGER Generated Documentation

If you want to generate documentation locally, you can run the following:

```shell
$ swagger serve swagger.json
```

This should prompt you to your web browser for the HTML generated docs.
