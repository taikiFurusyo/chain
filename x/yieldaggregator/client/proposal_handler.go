package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/UnUniFi/chain/x/yieldaggregator/client/cli"
)

var ProposalAddStrategyHandler = govclient.NewProposalHandler(cli.NewSubmitProposalAddStrategyTxCmd)
