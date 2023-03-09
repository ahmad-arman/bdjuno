package local

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ccvprovidertypes "github.com/cosmos/interchain-security/x/ccv/provider/types"
	ccvprovidersource "github.com/forbole/bdjuno/v4/modules/ccv/provider/source"
	"github.com/forbole/juno/v4/node/local"
)

var (
	_ ccvprovidersource.Source = &Source{}
)

// Source implements ccvprovidersource.Source using a local node
type Source struct {
	*local.Source
	querier ccvprovidertypes.QueryServer
}

// NewSource implements a new Source instance
func NewSource(source *local.Source, querier ccvprovidertypes.QueryServer) *Source {
	return &Source{
		Source:  source,
		querier: querier,
	}
}

// GetAllConsumerChains implements ccvprovidersource.Source
func (s Source) GetAllConsumerChains(height int64) ([]*ccvprovidertypes.Chain, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return []*ccvprovidertypes.Chain{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.querier.QueryConsumerChains(sdk.WrapSDKContext(ctx), &ccvprovidertypes.QueryConsumerChainsRequest{})
	if err != nil {
		return []*ccvprovidertypes.Chain{}, nil
	}

	return res.Chains, nil

}

// GetConsumerChainStarts implements ccvprovidersource.Source
func (s Source) GetConsumerChainStarts(height int64) (*ccvprovidertypes.ConsumerAdditionProposals, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nil, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.querier.QueryConsumerChainStarts(sdk.WrapSDKContext(ctx), &ccvprovidertypes.QueryConsumerChainStartProposalsRequest{})
	if err != nil {
		return nil, err
	}

	return res.Proposals, nil

}

// GetConsumerChainStops implements ccvprovidersource.Source
func (s Source) GetConsumerChainStops(height int64) (*ccvprovidertypes.ConsumerRemovalProposals, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nil, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.querier.QueryConsumerChainStops(sdk.WrapSDKContext(ctx), &ccvprovidertypes.QueryConsumerChainStopProposalsRequest{})
	if err != nil {
		return nil, err
	}

	return res.Proposals, nil

}