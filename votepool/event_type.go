package votepool

// EventType defines the legacy vote categories kept for downstream
// compatibility only.
type EventType uint8

const (
	ToBscCrossChainEvent EventType = 1
	FromBscCrossChainEvent EventType = 2
	DataAvailabilityChallengeEvent EventType = 3
	ToOpCrossChainEvent EventType = 4
	FromOpCrossChainEvent EventType = 5
)
