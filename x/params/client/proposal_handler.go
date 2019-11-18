package client

import (
	govclient "github.com/YAOChain/yao/x/gov/client"
	"github.com/YAOChain/yao/x/params/client/cli"
	"github.com/YAOChain/yao/x/params/client/rest"
)

// param change proposal handler
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
