// Copyright 2019 The go-filestorm Authors
// This file is part of the go-filestorm library.
//
// The go-filestorm library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-filestorm library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-filestorm library. If not, see <http://www.gnu.org/licenses/>.

package pbft

import (
	"fmt"

	"github.com/filestorm/go-filestorm/common"
	"github.com/filestorm/go-filestorm/consensus"
	"github.com/filestorm/go-filestorm/core/types"
	"github.com/filestorm/go-filestorm/rpc"
)

// API is a user facing RPC API to allow controlling the signer and voting
// mechanisms of the practical Byzantine fault tolerance scheme.
type API struct {
	chain consensus.ChainReader
	pbft  *Pbft
}

// GetBlockStatusByNumber retrieves the state snapshot at a given block.
func (api *API) GetBlockStatusByNumber(number *rpc.BlockNumber) (*Snapshot, error) {
	// Retrieve the requested block number (or current if none requested)
	var header *types.Header
	if number == nil || *number == rpc.LatestBlockNumber {
		header = api.chain.CurrentHeader()
	} else {
		header = api.chain.GetHeaderByNumber(uint64(number.Int64()))
	}
	// Ensure we have an actually valid block and return its snapshot
	if header == nil {
		return nil, errUnknownBlock
	}
	return api.pbft.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil)
}

// GetBlockStatusByHash retrieves the state snapshot at a given block.
func (api *API) GetBlockStatusByHash(hash common.Hash) (*Snapshot, error) {
	header := api.chain.GetHeaderByHash(hash)
	if header == nil {
		return nil, errUnknownBlock
	}
	return api.pbft.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil)
}

// GetValidatorsByNumber retrieves the list of authorized signers at the specified block.
func (api *API) GetValidatorsByNumber(number *rpc.BlockNumber) ([]common.Address, error) {
	// Retrieve the requested block number (or current if none requested)
	var header *types.Header
	if number == nil || *number == rpc.LatestBlockNumber {
		header = api.chain.CurrentHeader()
	} else {
		header = api.chain.GetHeaderByNumber(uint64(number.Int64()))
	}
	// Ensure we have an actually valid block and return the signers from its snapshot
	if header == nil {
		return nil, errUnknownBlock
	}
	snap, err := api.pbft.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil)
	if err != nil {
		return nil, err
	}
	return snap.signers(), nil
}

// GetValidatorsByHash retrieves the list of authorized signers at the specified block.
func (api *API) GetValidatorsByHash(hash common.Hash) ([]common.Address, error) {
	header := api.chain.GetHeaderByHash(hash)
	if header == nil {
		return nil, errUnknownBlock
	}
	snap, err := api.pbft.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil)
	if err != nil {
		return nil, err
	}
	return snap.signers(), nil
}

// Votes returns the current votes the node tries to uphold and vote on.
func (api *API) Votes() map[common.Address]bool {
	api.pbft.lock.RLock()
	defer api.pbft.lock.RUnlock()

	proposals := make(map[common.Address]bool)
	for address, auth := range api.pbft.proposals {
		proposals[address] = auth
	}
	return proposals
}

// Vote injects a new authorization vote that the signer will attempt to
// push through.
func (api *API) Vote(address common.Address, auth bool) {
	api.pbft.lock.Lock()
	defer api.pbft.lock.Unlock()

	api.pbft.proposals[address] = auth
}

// DropVote drops a currently running proposal, stopping the signer from casting
// further votes (either for or against).
func (api *API) DropVote(address common.Address) {
	api.pbft.lock.Lock()
	defer api.pbft.lock.Unlock()

	delete(api.pbft.proposals, address)
}

type status struct {
	InturnPercent float64                `json:"inturnPercent"`
	SigningStatus map[common.Address]int `json:"sealerActivity"`
	NumBlocks     uint64                 `json:"numBlocks"`
}

// BlockStatus returns the status of the last N blocks,
// - the number of active signers,
// - the number of signers,
// - the percentage of in-turn blocks
func (api *API) BlockStatus() (*status, error) {
	var (
		numBlocks = uint64(64)
		header    = api.chain.CurrentHeader()
		diff      = uint64(0)
		optimals  = 0
	)
	snap, err := api.pbft.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil)
	if err != nil {
		return nil, err
	}
	var (
		signers = snap.signers()
		end     = header.Number.Uint64()
		start   = end - numBlocks
	)
	if numBlocks > end {
		start = 1
		numBlocks = end - start
	}
	signStatus := make(map[common.Address]int)
	for _, s := range signers {
		signStatus[s] = 0
	}
	for n := start; n < end; n++ {
		h := api.chain.GetHeaderByNumber(n)
		if h == nil {
			return nil, fmt.Errorf("missing block %d", n)
		}
		if h.Difficulty.Cmp(diffInTurn) == 0 {
			optimals++
		}
		diff += h.Difficulty.Uint64()
		sealer, err := api.pbft.Author(h)
		if err != nil {
			return nil, err
		}
		signStatus[sealer]++
	}
	return &status{
		InturnPercent: float64((100 * optimals)) / float64(numBlocks),
		SigningStatus: signStatus,
		NumBlocks:     numBlocks,
	}, nil
}
