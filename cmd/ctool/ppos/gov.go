package ppos

import (
	"errors"

	"gopkg.in/urfave/cli.v1"

	"github.com/AlayaNetwork/Alaya-Go/common"
)

var (
	GovCmd = cli.Command{
		Name:  "gov",
		Usage: "use for gov func",
		Subcommands: []cli.Command{
			getProposalCmd,
			getTallyResultCmd,
			listProposalCmd,
			getActiveVersionCmd,
			getGovernParamValueCmd,
			getAccuVerifiersCountCmd,
			listGovernParamCmd,
		},
	}
	getProposalCmd = cli.Command{
		Name:   "getProposal",
		Usage:  "2100,get proposal,parameter:proposalID",
		Before: netCheck,
		Action: getProposal,
		Flags:  []cli.Flag{rpcUrlFlag, addressHRPFlag, proposalIDFlag, jsonFlag},
	}
	getTallyResultCmd = cli.Command{
		Name:   "getTallyResult",
		Usage:  "2101,get tally result,parameter:proposalID",
		Before: netCheck,
		Action: getTallyResult,
		Flags:  []cli.Flag{rpcUrlFlag, addressHRPFlag, proposalIDFlag, jsonFlag},
	}
	listProposalCmd = cli.Command{
		Name:   "listProposal",
		Usage:  "2102,list proposal",
		Before: netCheck,
		Action: listProposal,
		Flags:  []cli.Flag{rpcUrlFlag, addressHRPFlag, jsonFlag},
	}
	getActiveVersionCmd = cli.Command{
		Name:   "getActiveVersion",
		Usage:  "2103,query the effective version of the  chain",
		Before: netCheck,
		Action: getActiveVersion,
		Flags:  []cli.Flag{rpcUrlFlag, addressHRPFlag, jsonFlag},
	}
	getGovernParamValueCmd = cli.Command{
		Name:   "getGovernParamValue",
		Usage:  "2104,query the governance parameter value of the current block height,parameter:module,name",
		Before: netCheck,
		Action: getGovernParamValue,
		Flags:  []cli.Flag{rpcUrlFlag, addressHRPFlag, moduleFlag, nameFlag, jsonFlag},
	}
	getAccuVerifiersCountCmd = cli.Command{
		Name:   "getAccuVerifiersCount",
		Usage:  "2105,query the cumulative number of votes available for a proposal,parameter:proposalID,blockHash",
		Before: netCheck,
		Action: getAccuVerifiersCount,
		Flags:  []cli.Flag{rpcUrlFlag, addressHRPFlag, proposalIDFlag, blockHashFlag, jsonFlag},
	}
	listGovernParamCmd = cli.Command{
		Name:   "listGovernParam",
		Usage:  "2106,query the list of governance parameters,parameter:module",
		Before: netCheck,
		Action: listGovernParam,
		Flags:  []cli.Flag{rpcUrlFlag, addressHRPFlag, moduleFlag, jsonFlag},
	}
	proposalIDFlag = cli.StringFlag{
		Name:  "proposalID",
		Usage: "proposalID",
	}
	moduleFlag = cli.StringFlag{
		Name:  "module",
		Usage: "module",
	}
	nameFlag = cli.StringFlag{
		Name:  "name",
		Usage: "name",
	}
	blockHashFlag = cli.StringFlag{
		Name:  "blockHash",
		Usage: "blockHash",
	}
)

func getProposal(c *cli.Context) error {
	proposalIDstring := c.String(proposalIDFlag.Name)
	if proposalIDstring == "" {
		return errors.New("proposalID not set")
	}
	proposalID := common.HexToHash(proposalIDstring)

	return query(c, 2100, proposalID)
}

func getTallyResult(c *cli.Context) error {
	proposalIDstring := c.String(proposalIDFlag.Name)
	if proposalIDstring == "" {
		return errors.New("param proposalID not set")
	}
	proposalID := common.HexToHash(proposalIDstring)

	return query(c, 2101, proposalID)
}

func listProposal(c *cli.Context) error {
	return query(c, 2102)
}

func getActiveVersion(c *cli.Context) error {
	return query(c, 2103)
}

func getGovernParamValue(c *cli.Context) error {
	module := c.String(moduleFlag.Name)
	if module == "" {
		return errors.New("param module not set")
	}
	name := c.String(nameFlag.Name)
	if name == "" {
		return errors.New("param name not set")
	}
	return query(c, 2104, module, name)
}

func getAccuVerifiersCount(c *cli.Context) error {
	proposalIDstring := c.String(proposalIDFlag.Name)
	if proposalIDstring == "" {
		return errors.New("param proposalID not set")
	}
	blockHash := c.String(blockHashFlag.Name)
	if blockHash == "" {
		return errors.New("param block hash not set")
	}
	return query(c, 2105, common.HexToHash(proposalIDstring), common.HexToHash(blockHash))
}

func listGovernParam(c *cli.Context) error {
	module := c.String(moduleFlag.Name)
	return query(c, 2106, module)
}
