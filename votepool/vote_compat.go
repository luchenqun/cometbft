package votepool

import (
	"errors"
	"time"
)

const (
	eventHashLen = 32
	pubKeyLen    = 128
	signatureLen = 64
)

// Vote is kept as a minimal data carrier so downstream clients and mocks keep
// compiling after the full votepool subsystem removal.
type Vote struct {
	PubKey    []byte    `json:"pub_key"`
	Signature []byte    `json:"signature"`
	EventType EventType `json:"event_type"`
	EventHash []byte    `json:"event_hash"`

	expireAt time.Time
}

func NewVote(pubKey, signature []byte, eventType uint8, eventHash []byte) *Vote {
	return &Vote{
		PubKey:    pubKey,
		Signature: signature,
		EventType: EventType(eventType),
		EventHash: eventHash,
	}
}

func (v *Vote) Key() string {
	return string(v.EventHash[:]) + string(v.PubKey[:])
}

func (v *Vote) ValidateBasic() error {
	if len(v.EventHash) != eventHashLen {
		return errors.New("invalid event hash")
	}
	if v.EventType != ToBscCrossChainEvent &&
		v.EventType != FromBscCrossChainEvent &&
		v.EventType != FromOpCrossChainEvent &&
		v.EventType != ToOpCrossChainEvent &&
		v.EventType != DataAvailabilityChallengeEvent {
		return errors.New("invalid event type")
	}
	if len(v.PubKey) != pubKeyLen {
		return errors.New("invalid public key")
	}
	if len(v.Signature) != signatureLen {
		return errors.New("invalid signature")
	}
	return nil
}
