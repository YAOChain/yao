package client

import (
	"github.com/YAOChain/yao/x/distribution/client/cli"
	"github.com/YAOChain/yao/x/distribution/client/rest"
	govclient "github.com/YAOChain/yao/x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
