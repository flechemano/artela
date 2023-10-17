package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"

	"github.com/artela-network/evm/vm"
	tmrpctypes "github.com/cometbft/cometbft/rpc/core/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/artela-network/artela/ethereum/rpc/ethapi"
	rpctypes "github.com/artela-network/artela/ethereum/rpc/types"
	"github.com/artela-network/artela/x/evm/txs"
	evmtypes "github.com/artela-network/artela/x/evm/types"
)

// Blockchain API

func (b *backend) SetHead(_ uint64) {
	b.logger.Error("SetHead is not implemented")
}

func (b *backend) HeaderByNumber(_ context.Context, number rpc.BlockNumber) (*ethtypes.Header, error) {
	resBlock, err := b.CosmosBlockByNumber(number)
	if err != nil {
		return nil, err
	}

	if resBlock == nil {
		return nil, fmt.Errorf("block not found for height %d", number)
	}

	blockRes, err := b.CosmosBlockResultByNumber(&resBlock.Block.Height)
	if err != nil {
		return nil, fmt.Errorf("block result not found for height %d", resBlock.Block.Height)
	}

	bloom, err := b.blockBloom(blockRes)
	if err != nil {
		b.logger.Debug("HeaderByNumber BlockBloom failed", "height", resBlock.Block.Height)
	}

	baseFee, err := b.BaseFee(blockRes)
	if err != nil {
		// handle the error for pruned node.
		b.logger.Error("failed to fetch Base Fee from prunned block. Check node prunning configuration", "height", resBlock.Block.Height, "error", err)
	}

	ethHeader := rpctypes.EthHeaderFromTendermint(resBlock.Block.Header, bloom, baseFee)
	return ethHeader, nil
}

func (b *backend) HeaderByHash(_ context.Context, hash common.Hash) (*ethtypes.Header, error) {
	return nil, nil
}

func (b *backend) HeaderByNumberOrHash(ctx context.Context,
	blockNrOrHash rpc.BlockNumberOrHash,
) (*ethtypes.Header, error) {
	return nil, errors.New("HeaderByNumberOrHash is not implemented")
}

func (b *backend) CurrentHeader() *ethtypes.Header {
	block, err := b.BlockByNumber(context.Background(), rpc.LatestBlockNumber)
	if err != nil || block == nil {
		b.logger.Error("geet CurrentHeader failed, error or block is nil", "error", err)
		return nil
	}
	return block.Header()
}

func (b *backend) CurrentBlock() *rpctypes.Block {
	block, err := b.ArtBlockByNumber(context.Background(), rpc.LatestBlockNumber)
	if err != nil {
		b.logger.Error("get CurrentBlock failed", "error", err)
		return nil
	}
	return block
}

func (b *backend) BlockByNumber(_ context.Context, number rpc.BlockNumber) (*ethtypes.Block, error) {
	block, err := b.ArtBlockByNumber(context.Background(), number)
	if err != nil {
		return nil, err
	}
	return block.EthBlock(), nil
}

func (b *backend) ArtBlockByNumber(_ context.Context, number rpc.BlockNumber) (*rpctypes.Block, error) {
	resBlock, err := b.CosmosBlockByNumber(number)
	if err != nil || resBlock == nil {
		return nil, fmt.Errorf("query block failed, block number %d, %w", number, err)
	}

	blockRes, err := b.CosmosBlockResultByNumber(&resBlock.Block.Height)
	if err != nil {
		return nil, fmt.Errorf("block result not found for height %d", resBlock.Block.Height)
	}

	return b.BlockFromCosmosBlock(resBlock, blockRes)
}

func (b *backend) BlockByHash(_ context.Context, hash common.Hash) (*rpctypes.Block, error) {
	resBlock, err := b.CosmosBlockByHash(hash)
	if err != nil || resBlock == nil {
		return nil, fmt.Errorf("failed to get block by hash %s, %w", hash.Hex(), err)
	}

	blockRes, err := b.CosmosBlockResultByNumber(&resBlock.Block.Height)
	if err != nil {
		return nil, fmt.Errorf("block result not found for height %d", resBlock.Block.Height)
	}

	return b.BlockFromCosmosBlock(resBlock, blockRes)
}

func (b *backend) BlockByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*rpctypes.Block, error) {
	return nil, errors.New("BlockByNumberOrHash is not implemented")
}

func (b *backend) StateAndHeaderByNumber(
	ctx context.Context, number rpc.BlockNumber,
) (*state.StateDB, *ethtypes.Header, error) {
	return nil, nil, errors.New("StateAndHeaderByNumber is not implemented")
}

func (b *backend) StateAndHeaderByNumberOrHash(
	ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash,
) (*state.StateDB, *ethtypes.Header, error) {
	return nil, nil, errors.New("invalid arguments; neither block nor hash specified")
}

func (b *backend) PendingBlockAndReceipts() (*ethtypes.Block, types.Receipts) {
	b.logger.Error("PendingBlockAndReceipts is not implemented")
	return nil, nil
}

// GetReceipts get receipts by block hash
func (b *backend) GetReceipts(_ context.Context, hash common.Hash) (types.Receipts, error) {
	return nil, errors.New("GetReceipts is not implemented")
}

func (b *backend) GetTd(_ context.Context, hash common.Hash) *big.Int {
	b.logger.Error("GetTd is not implemented")
	return nil
}

func (b *backend) GetEVM(ctx context.Context, msg *core.Message, state *state.StateDB,
	header *ethtypes.Header, vmConfig *vm.Config, _ *vm.BlockContext,
) (*vm.EVM, func() error) {
	return nil, func() error {
		return errors.New("GetEVM is not impelemted")
	}
}

func (b *backend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	return b.scope.Track(b.chainFeed.Subscribe(ch))
}

func (b *backend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	b.logger.Debug("called eth.rpc.backend.SubscribeChainHeadEvent", "ch", ch)
	return b.scope.Track(b.chainHeadFeed.Subscribe(ch))
}

func (b *backend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	b.logger.Debug("called eth.rpc.backend.SubscribeChainSideEvent", "ch", ch)
	return b.scope.Track(b.chainSideFeed.Subscribe(ch))
}

func (b *backend) CosmosBlockByHash(blockHash common.Hash) (*tmrpctypes.ResultBlock, error) {
	resBlock, err := b.clientCtx.Client.BlockByHash(b.ctx, blockHash.Bytes())
	if err != nil {
		return nil, err
	}

	if resBlock.Block == nil {
		return nil, fmt.Errorf("failed to query block for hash: %s", blockHash.Hex())
	}

	return resBlock, nil
}

func (b *backend) CosmosBlockByNumber(blockNum rpc.BlockNumber) (*tmrpctypes.ResultBlock, error) {
	height := blockNum.Int64()
	if height <= 0 {
		// fetch the latest block number from the app state, more accurate than the tendermint block store state.
		n, err := b.BlockNumber()
		if err != nil {
			return nil, err
		}
		height = int64(n) //#nosec G701 -- checked for int overflow already
	}
	resBlock, err := b.clientCtx.Client.Block(b.ctx, &height)
	if err != nil {
		return nil, err
	}

	if resBlock.Block == nil {
		return nil, fmt.Errorf("failed to query block for blockNum: %d", blockNum.Int64())
	}

	return resBlock, nil
}

// BlockNumberFromTendermint returns the BlockNumber from BlockNumberOrHash
func (b *backend) blockNumberFromCosmos(blockNrOrHash rpc.BlockNumberOrHash) (rpc.BlockNumber, error) {
	switch {
	case blockNrOrHash.BlockHash == nil && blockNrOrHash.BlockNumber == nil:
		return rpc.EarliestBlockNumber, fmt.Errorf("types BlockHash and BlockNumber cannot be both nil")
	case blockNrOrHash.BlockHash != nil:
		resBlock, err := b.CosmosBlockByHash(*blockNrOrHash.BlockHash)
		if err != nil || resBlock.Block == nil {
			return rpc.EarliestBlockNumber, err
		}
		return rpc.BlockNumber(resBlock.Block.Height), nil
	case blockNrOrHash.BlockNumber != nil:
		if *blockNrOrHash.BlockNumber == rpc.LatestBlockNumber {
			currentHeight := b.CurrentHeader().Number
			return rpc.BlockNumber(currentHeight.Int64()), nil
		}
		return *blockNrOrHash.BlockNumber, nil
	default:
		return rpc.EarliestBlockNumber, nil
	}
}

func (b *backend) BlockNumber() (hexutil.Uint64, error) {
	// do any grpc query, ignore the response and use the returned block height
	var header metadata.MD
	_, err := b.queryClient.Params(b.ctx, &txs.QueryParamsRequest{}, grpc.Header(&header))
	if err != nil {
		return hexutil.Uint64(0), err
	}

	blockHeightHeader := header.Get(grpctypes.GRPCBlockHeightHeader)
	if headerLen := len(blockHeightHeader); headerLen != 1 {
		return 0, fmt.Errorf("unexpected '%s' gRPC header length; got %d, expected: %d", grpctypes.GRPCBlockHeightHeader, headerLen, 1)
	}

	height, err := strconv.ParseUint(blockHeightHeader[0], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse block height: %w", err)
	}

	if height > math.MaxInt64 {
		return 0, fmt.Errorf("block height %d is greater than max uint64", height)
	}

	return hexutil.Uint64(height), nil
}

func (b *backend) BlockTimeByNumber(blockNum int64) (uint64, error) {
	resBlock, err := b.clientCtx.Client.Block(b.ctx, &blockNum)
	if err != nil {
		return 0, err
	}
	return uint64(resBlock.Block.Time.Unix()), nil
}

func (b *backend) CosmosBlockResultByNumber(height *int64) (*tmrpctypes.ResultBlockResults, error) {
	return b.clientCtx.Client.BlockResults(b.ctx, height)
}

func (b *backend) GetCode(address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	blockNum, err := b.blockNumberFromCosmos(blockNrOrHash)
	if err != nil {
		return nil, err
	}

	req := &txs.QueryCodeRequest{
		Address: address.String(),
	}

	res, err := b.queryClient.Code(rpctypes.ContextWithHeight(blockNum.Int64()), req)
	if err != nil {
		return nil, err
	}

	return res.Code, nil
}

// BlockBloom query block bloom filter from block results
func (b *backend) blockBloom(blockRes *tmrpctypes.ResultBlockResults) (ethtypes.Bloom, error) {
	for _, event := range blockRes.EndBlockEvents {
		if event.Type != evmtypes.EventTypeBlockBloom {
			continue
		}

		for _, attr := range event.Attributes {
			if attr.Key == evmtypes.AttributeKeyEthereumBloom {
				return ethtypes.BytesToBloom([]byte(attr.Value)), nil
			}
		}
	}
	return ethtypes.Bloom{}, errors.New("block bloom event is not found")
}

func (b *backend) BlockFromCosmosBlock(resBlock *tmrpctypes.ResultBlock, blockRes *tmrpctypes.ResultBlockResults) (*rpctypes.Block, error) {
	block := resBlock.Block
	height := block.Height
	bloom, err := b.blockBloom(blockRes)
	if err != nil {
		b.logger.Debug("HeaderByNumber BlockBloom failed", "height", height)
	}

	baseFee, err := b.BaseFee(blockRes)
	if err != nil {
		b.logger.Error("failed to fetch Base Fee from prunned block. Check node prunning configuration", "height", height, "error", err)
	}

	ethHeader := rpctypes.EthHeaderFromTendermint(block.Header, bloom, baseFee)
	msgs := b.EthMsgsFromCosmosBlock(resBlock, blockRes)

	txs := make([]*ethtypes.Transaction, len(msgs))
	for i, ethMsg := range msgs {
		txs[i] = ethMsg.AsTransaction()
	}

	gasUsed := uint64(0)
	for _, txsResult := range blockRes.TxsResults {
		// workaround for cosmos-sdk bug. https://github.com/cosmos/cosmos-sdk/issues/10832
		if ShouldIgnoreGasUsed(txsResult) {
			// block gas limit has exceeded, other txs must have failed with same reason.
			break
		}
		gasUsed += uint64(txsResult.GetGasUsed())
	}
	ethHeader.GasUsed = gasUsed

	gasLimit, err := rpctypes.BlockMaxGasFromConsensusParams(context.Background(), b.clientCtx, block.Height)
	if err != nil {
		b.logger.Error("failed to query consensus params", "error", err.Error())
	}
	ethHeader.GasLimit = uint64(gasLimit)

	blockHash := common.BytesToHash(block.Hash().Bytes())
	receipts, err := b.GetReceipts(context.Background(), blockHash)
	if err != nil {
		b.logger.Debug("failed to fetch receipts for block number %d", height)
	}

	ethBlock := ethtypes.NewBlock(ethHeader, txs, nil, receipts, trie.NewStackTrie(nil))
	res := rpctypes.EthBlockToBlock(ethBlock)
	res.SetHash(blockHash)
	return res, nil
}

func (b *backend) EthMsgsFromCosmosBlock(resBlock *tmrpctypes.ResultBlock, blockRes *tmrpctypes.ResultBlockResults) []*txs.MsgEthereumTx {
	var result []*txs.MsgEthereumTx
	block := resBlock.Block

	txResults := blockRes.TxsResults

	for i, tx := range block.Txs {
		// Check if tx exists on EVM by cross checking with blockResults:
		//  - Include unsuccessful tx that exceeds block gas limit
		//  - Exclude unsuccessful tx with any other error but ExceedBlockGasLimit
		if !rpctypes.TxSuccessOrExceedsBlockGasLimit(txResults[i]) {
			b.logger.Debug("invalid tx result code", "cosmos-hash", hexutil.Encode(tx.Hash()))
			continue
		}

		tx, err := b.clientCtx.TxConfig.TxDecoder()(tx)
		if err != nil {
			b.logger.Debug("failed to decode transaction in block", "height", block.Height, "error", err.Error())
			continue
		}

		for _, msg := range tx.GetMsgs() {
			ethMsg, ok := msg.(*txs.MsgEthereumTx)
			if !ok {
				continue
			}

			ethMsg.Hash = ethMsg.AsTransaction().Hash().Hex()
			result = append(result, ethMsg)
		}
	}

	return result
}

func (b *backend) DoCall(args ethapi.TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash) (*txs.MsgEthereumTxResponse, error) {
	blockNum, err := b.blockNumberFromCosmos(blockNrOrHash)
	if err != nil {
		return nil, err
	}

	bz, err := json.Marshal(&args)
	if err != nil {
		return nil, err
	}
	header, err := b.CosmosBlockByNumber(blockNum)
	if err != nil {
		// the error message imitates geth behavior
		return nil, errors.New("header not found")
	}

	req := txs.EthCallRequest{
		Args:            bz,
		GasCap:          b.RPCGasCap(),
		ProposerAddress: sdktypes.ConsAddress(header.Block.ProposerAddress),
		ChainId:         b.chainID.Int64(),
	}

	// From ContextWithHeight: if the provided height is 0,
	// it will return an empty context and the gRPC query will use
	// the latest block height for querying.
	ctx := rpctypes.ContextWithHeight(blockNum.Int64())
	timeout := b.RPCEVMTimeout()

	// Setup context so it may be canceled the call has completed
	// or, in case of unmetered gas, setup a context with a timeout.
	var cancel context.CancelFunc
	if timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout)
	} else {
		ctx, cancel = context.WithCancel(ctx)
	}

	// Make sure the context is canceled when the call has completed
	// this makes sure resources are cleaned up.
	defer cancel()

	res, err := b.queryClient.EthCall(ctx, &req)
	if err != nil {
		return nil, err
	}

	if res.Failed() {
		if res.VmError != vm.ErrExecutionReverted.Error() {
			return nil, status.Error(codes.Internal, res.VmError)
		}
		return nil, evmtypes.NewExecErrorWithReason(res.Ret)
	}

	return res, nil
}