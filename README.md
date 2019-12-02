[![Go Report Card](https://goreportcard.com/badge/github.com/mlesniak/kafka-with-go)](https://goreportcard.com/report/github.com/mlesniak/kafka-with-go)
[![Build Status](https://github.com/mlesniak/kafka-with-go/workflows/Go/badge.svg)](https://github.com/mlesniak/kafka-with-go/actions?query=workflow%3AGo)
[![Code of Conduct](https://img.shields.io/badge/%E2%9D%A4-code%20of%20conduct-orange.svg?style=flat)](CODE_OF_CONDUCT.md)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://raw.githubusercontent.com/mlesniak/kafka-with-go/master/LICENSE)

# Overview

This is a small playground example for using Apache Kafka with Go. Initially, I used [Shopify's samara](https://github.com/Shopify/sarama)
library but switched
to [Confluent's kafka](https://github.com/confluentinc/confluent-kafka-go) library due to easier use and more examples. Note that my final opinion has not been
fully set, yet.
 
 Currently, you can create and list topics and inject random
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

    # Once at the beginning
    docker-compose build
    
    # Everytime...
    docker-compose up
    docker-compose exec go ash
    
and use the aforementioned `kafka-with-go` commands in the provided shell. 

## Full example

    # Create topic
    $ kafka-with-go -create -topic github -partitions 4
    2019/12/01 13:37:16 Starting topic creation
    2019/12/01 13:37:16 Creating topic github with 4 partitions
    2019/12/01 13:37:16 Finished topic creation
    
    # Add 1k messages with length 128bytes each; rather slow on my old machine ¯\_(ツ)_/¯ ... 
    $ kafka-with-go -produce -topic github -number 1000 -length 128 -tick 100
    2019/12/01 13:38:05 Starting producer
    2019/12/01 13:38:05 Sending message 100/1000
    2019/12/01 13:38:06 Sending message 200/1000
    2019/12/01 13:38:06 Sending message 300/1000
    2019/12/01 13:38:07 Sending message 400/1000
    2019/12/01 13:38:07 Sending message 500/1000
    2019/12/01 13:38:07 Sending message 600/1000
    2019/12/01 13:38:08 Sending message 700/1000
    2019/12/01 13:38:08 Sending message 800/1000
    2019/12/01 13:38:09 Sending message 900/1000
    2019/12/01 13:38:09 Sending message 1000/1000
    2019/12/01 13:38:09 Finished producer
    
    # Consume 1k messages
    $ kafka-with-go -consume -topic github -number 1000 -tick 100 -group 1
    2019/12/01 13:39:06 Starting consumer
    2019/12/01 13:39:10 Received 100 messages (100 new)
    2019/12/01 13:39:10 Received 200 messages (100 new)
    2019/12/01 13:39:10 Received 300 messages (100 new)
    2019/12/01 13:39:10 Received 400 messages (100 new)
    2019/12/01 13:39:10 Received 500 messages (100 new)
    2019/12/01 13:39:10 Received 600 messages (100 new)
    2019/12/01 13:39:10 Received 700 messages (100 new)
    2019/12/01 13:39:10 Received 800 messages (100 new)
    2019/12/01 13:39:10 Received 900 messages (100 new)
    2019/12/01 13:39:10 Received 1000 messages (100 new)
    2019/12/01 13:39:10 Consumer read 1000 entries
    2019/12/01 13:39:10 Finished consumer   

## License

As always, the source code is licensed under [Apache license 2.0](https://raw.githubusercontent.com/mlesniak/kafka-with-go/master/LICENSE).
