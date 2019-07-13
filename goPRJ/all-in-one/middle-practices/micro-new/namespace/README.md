# Namespace Service

This is the Namespace service

Generated with

```
micro new github.com/micro-in-cn/all-in-one/middle-practices/micro-new/namespace --namespace=mu.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: mu.micro.srv.namespace
- Type: srv
- Alias: namespace

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./namespace-srv
```

Build a docker image
```
make docker
```