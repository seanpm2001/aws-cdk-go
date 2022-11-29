package awssqs


// Whether the FIFO queue throughput quota applies to the entire queue or per message group.
// Experimental.
type FifoThroughputLimit string

const (
	// Throughput quota applies per queue.
	// Experimental.
	FifoThroughputLimit_PER_QUEUE FifoThroughputLimit = "PER_QUEUE"
	// Throughput quota applies per message group id.
	// Experimental.
	FifoThroughputLimit_PER_MESSAGE_GROUP_ID FifoThroughputLimit = "PER_MESSAGE_GROUP_ID"
)

