// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
	"github.com/yoshi-ai-dev/yoshi-cli/internal/apiquery"
	"github.com/yoshi-ai-dev/yoshi-cli/internal/requestflag"
	"github.com/yoshi-ai-dev/yoshi-go"
	"github.com/yoshi-ai-dev/yoshi-go/option"
)

var paperTradingAccountsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a paper trading account. Requires user approval in the Yoshi web app\nbefore the account is created.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[bool]{
			Name:     "skip-approval",
			Usage:    "Skip the approval flow and execute immediately. Not supported yet — reserved for future use.",
			Default:  false,
			BodyPath: "skip_approval",
		},
		&requestflag.Flag[string]{
			Name:     "starting-cash-balance",
			Default:  "100000",
			BodyPath: "starting_cash_balance",
		},
	},
	Action:          handlePaperTradingAccountsCreate,
	HideHelpCommand: true,
}

var paperTradingAccountsList = cli.Command{
	Name:            "list",
	Usage:           "List the user's paper trading accounts with current balances.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handlePaperTradingAccountsList,
	HideHelpCommand: true,
}

func handlePaperTradingAccountsCreate(ctx context.Context, cmd *cli.Command) error {
	client := yoshi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := yoshi.PaperTradingAccountNewParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PaperTrading.Accounts.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		Title:          "paper-trading:accounts create",
		Transform:      transform,
	})
}

func handlePaperTradingAccountsList(ctx context.Context, cmd *cli.Command) error {
	client := yoshi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PaperTrading.Accounts.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		Title:          "paper-trading:accounts list",
		Transform:      transform,
	})
}
