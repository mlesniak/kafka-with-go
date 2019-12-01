[![Go Report Card](https://goreportcard.com/badge/github.com/mlesniak/kafka-with-go)](https://goreportcard.com/report/github.com/mlesniak/kafka-with-go)
[![Build Status](https://github.com/mlesniak/kafka-with-go/workflows/Go/badge.svg)](https://github.com/mlesniak/kafka-with-go/actions?query=workflow%3AGo)
[![Code of Conduct](https://img.shields.io/badge/%E2%9D%A4-code%20of%20conduct-orange.svg?style=flat)](CODE_OF_CONDUCT.md)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://raw.githubusercontent.com/mlesniak/kafka-with-go/master/LICENSE)

# Overview

This is a small playground example for using Apache Kafka with Go. Initially, I used Shopify's samara library but switched
to Confluents library due to easier use and more examples. Currently, you can create and list topics and inject random
messages  of pre-defined length and size as well as consume them (see below).

## Build

The underlying library needs `librdkafka`, hence it has to be installed using

    brew install librdkafka
    
(For linux, see the `Dockerfile`) . Afterwards, a simple

    go build
    
builds the binary.

## Usage  

To **create a topic** specify the topic and the number of partitions; note that the replication factor is always 1

    kafka-with-go -create -topic demo -partitions 4
    
To **list all topics**, call 

    kafka-with-go -list
    
which shows all (including internal) topics and their number of partitions.

To **inject messages into kafka**, define the topic, the number of messages and their length

    kafka-with-go -produce -number 10000 -length 10 -topic demo
    
and to **consume messages**, define the topic and the group id

    kafka-with-go -consume -topic demo -group 1                

There are optional additional flags you can specify, e.g.

- `tick` to define how often status message are printed
- `broker` to set the broker (default is `localhost:9092`)
- `number` in the consumer to stop receiving messages after the predefined number of messages
  
## Docker

The (multi-stage) `Dockerfile` builds a container without external dependencies, which can then be used in arbitrary
environments, e.g. in docker compose or kubernetes; a corresponding `docker-compose.yml` is provided. Use it as follows: 


## License

As always, the source code is licensed under [Apache license 2.0](https://raw.githubusercontent.com/mlesniak/kafka-with-go/master/LICENSE).
