# Amazon Simple Queue Service Construct Library

Amazon Simple Queue Service (SQS) is a fully managed message queuing service that
enables you to decouple and scale microservices, distributed systems, and serverless
applications. SQS eliminates the complexity and overhead associated with managing and
operating message oriented middleware, and empowers developers to focus on differentiating work.
Using SQS, you can send, store, and receive messages between software components at any volume,
without losing messages or requiring other services to be available.

## Installation

Import to your project:

```go
import sqs "github.com/aws/aws-cdk-go/awscdk"
```

## Basic usage

Here's how to add a basic queue to your application:

```go
sqs.NewQueue(this, jsii.String("Queue"))
```

## Encryption

If you want to encrypt the queue contents, set the `encryption` property. You can have
the messages encrypted with a key that SQS manages for you, or a key that you
can manage yourself.

```go
// Use managed key
// Use managed key
sqs.NewQueue(this, jsii.String("Queue"), &queueProps{
	encryption: sqs.queueEncryption_KMS_MANAGED,
})

// Use custom key
myKey := kms.NewKey(this, jsii.String("Key"))

sqs.NewQueue(this, jsii.String("Queue"), &queueProps{
	encryption: sqs.*queueEncryption_KMS,
	encryptionMasterKey: myKey,
})
```

## First-In-First-Out (FIFO) queues

FIFO queues give guarantees on the order in which messages are dequeued, and have additional
features in order to help guarantee exactly-once processing. For more information, see
the SQS manual. Note that FIFO queues are not available in all AWS regions.

A queue can be made a FIFO queue by either setting `fifo: true`, giving it a name which ends
in `".fifo"`, or by enabling a FIFO specific feature such as: content-based deduplication,
deduplication scope or fifo throughput limit.