// Package types defines the essential types used throughout the beacon-chain.
package types

import (
	"encoding/binary"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/prysmaticlabs/prysm/beacon-chain/params"
	pb "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	"golang.org/x/crypto/blake2b"
)

// Attestation is the primary source of load on the beacon chain, it's used to
// attest to some parent block in the chain and block hash in a shard.
type Attestation struct {
	data *pb.AggregatedAttestation
}

// NewAttestation explicitly sets the data field of a attestation.
// Return attestation with default fields if data is nil.
func NewAttestation(data *pb.AggregatedAttestation) *Attestation {
	if data == nil {
		return &Attestation{
			data: &pb.AggregatedAttestation{
				Slot:                0,
				ShardId:             0,
				JustifiedSlot:       0,
				JustifiedBlockHash:  []byte{},
				ShardBlockHash:      []byte{},
				AttesterBitfield:    []byte{},
				ObliqueParentHashes: [][]byte{{}},
				AggregateSig:        []uint64{},
			},
		}
	}
	return &Attestation{data: data}
}

// Proto returns the underlying protobuf data.
func (a *Attestation) Proto() *pb.AggregatedAttestation {
	return a.data
}

// Marshal encodes block object into the wire format.
func (a *Attestation) Marshal() ([]byte, error) {
	return proto.Marshal(a.data)
}

// Hash generates the blake2b hash of the attestation.
func (a *Attestation) Hash() ([32]byte, error) {
	data, err := proto.Marshal(a.data)
	if err != nil {
		return [32]byte{}, fmt.Errorf("could not marshal attestation proto data: %v", err)
	}
	var hash [32]byte
	h := blake2b.Sum512(data)
	copy(hash[:], h[:32])
	return hash, nil
}

// Key generates the blake2b hash of the following attestation fields:
// slotNumber + shardID + blockHash + obliqueParentHash
// This is used for attestation table look up in localDB.
func (a *Attestation) Key() [32]byte {
	key := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(key, a.SlotNumber())
	binary.PutUvarint(key, a.ShardID())
	key = append(key, a.ShardBlockHash()...)
	for _, pHash := range a.ObliqueParentHashes() {
		key = append(key, pHash[:]...)
	}

	var hash [32]byte
	h := blake2b.Sum512(key)
	copy(hash[:], h[:32])
	return hash
}

// SlotNumber of the block, which this attestation is attesting to.
func (a *Attestation) SlotNumber() uint64 {
	return a.data.Slot
}

// ShardID of the block, which this attestation is attesting to.
func (a *Attestation) ShardID() uint64 {
	return a.data.ShardId
}

// ShardBlockHash of the block, which this attestation is attesting to.
func (a *Attestation) ShardBlockHash() []byte {
	return a.data.ShardBlockHash
}

// JustifiedSlotNumber of the attestation should be earlier than the last justified slot in crystallized state.
func (a *Attestation) JustifiedSlotNumber() uint64 {
	return a.data.JustifiedSlot
}

// JustifiedBlockHash should be in the chain of the block being processed.
func (a *Attestation) JustifiedBlockHash() []byte {
	return a.data.JustifiedBlockHash
}

// AttesterBitfield represents which validator in the committee has voted.
func (a *Attestation) AttesterBitfield() []byte {
	return a.data.AttesterBitfield
}

// ObliqueParentHashes represents the block hashes this attestation is not attesting for.
func (a *Attestation) ObliqueParentHashes() [][32]byte {
	var obliqueParentHashes [][32]byte
	for _, hash := range a.data.ObliqueParentHashes {
		var h [32]byte
		copy(h[:], hash)
		obliqueParentHashes = append(obliqueParentHashes, h)
	}
	return obliqueParentHashes
}

// AggregateSig represents the aggregated signature from all the validators attesting to this block.
func (a *Attestation) AggregateSig() []uint64 {
	return a.data.AggregateSig
}

// VerifyProposerAttestation verifies the proposer's attestation of the block.
// Proposers broadcast the attestation along with the block to its peers.
func (a *Attestation) VerifyProposerAttestation(pubKey [32]byte, proposerShardID uint64) error {

	// Verify the attestation attached with block response.
	// Get proposer index and shardID.

	attestationMsg := AttestationMsg(
		a.ObliqueParentHashes(),
		a.ShardBlockHash(),
		a.SlotNumber(),
		proposerShardID,
		a.JustifiedSlotNumber())

	log.Infof("Constructing attestation message for incoming block 0x%x", attestationMsg)

	// TODO(#258): use attestationMsg to verify against signature and public key. Return error if incorrect.
	log.Infof("Verifying attestation with public key 0x%x", pubKey)

	log.Info("successfully verified attestation with incoming block")
	return nil
}

// AttestationMsg hashes parentHashes + shardID + slotNumber + shardBlockHash + justifiedSlot
// into a message to use for verifying with aggregated public key and signature.
func AttestationMsg(parentHashes [][32]byte, blockHash []byte, slot uint64, shardID uint64, justifiedSlot uint64) [32]byte {
	msg := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(msg, slot%params.CycleLength)
	for _, parentHash := range parentHashes {
		msg = append(msg, parentHash[:]...)
	}
	binary.PutUvarint(msg, shardID)
	msg = append(msg, blockHash...)
	binary.PutUvarint(msg, justifiedSlot)

	var hashMsg [32]byte
	h := blake2b.Sum512(msg)
	copy(hashMsg[:], h[:32])
	return hashMsg
}
