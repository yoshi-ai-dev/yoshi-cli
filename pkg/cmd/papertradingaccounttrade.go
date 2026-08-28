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

var paperTradingAccountsTradesCreate = cli.Command{
	Name:    "create",
	Usage:   "Place a buy or sell trade. Requires user approval in the Yoshi web app before\nexecution.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Required:  true,
			PathParam: "accountId",
		},
		&requestflag.Flag[string]{
			Name:     "side",
			Usage:    `Allowed values: "buy", "sell".`,
			Required: true,
			BodyPath: "side",
		},
		&requestflag.Flag[string]{
			Name:     "symbol",
			Required: true,
			BodyPath: "symbol",
		},
		&requestflag.Flag[float64]{
			Name:     "notional",
			BodyPath: "notional",
		},
		&requestflag.Flag[float64]{
			Name:     "quantity",
			BodyPath: "quantity",
		},
		&requestflag.Flag[bool]{
			Name:     "skip-approval",
			Usage:    "Skip the approval flow and execute immediately. Not supported yet — reserved for future use.",
			Default:  false,
			BodyPath: "skip_approval",
		},
	},
	Action:          handlePaperTradingAccountsTradesCreate,
	HideHelpCommand: true,
}

var paperTradingAccountsTradesList = cli.Command{
	Name:    "list",
	Usage:   "List trade history for a Test Drive account with cursor-based pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Required:  true,
			PathParam: "accountId",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque cursor from a previous response",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Items per page (1-100, default 50)",
			Default:   50,
			QueryPath: "limit",
		},
	},
	Action:          handlePaperTradingAccountsTradesList,
	HideHelpCommand: true,
}

func handlePaperTradingAccountsTradesCreate(ctx context.Context, cmd *cli.Command) error {
	client := yoshi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := yoshi.PaperTradingAccountTradeNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PaperTrading.Accounts.Trades.New(
		ctx,
		cmd.Value("account-id").(string),
		params,
		options...,
	)
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
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "paper-trading:accounts:trades create",
		Transform:      transform,
	})
}

func handlePaperTradingAccountsTradesList(ctx context.Context, cmd *cli.Command) error {
	client := yoshi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := yoshi.PaperTradingAccountTradeListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PaperTrading.Accounts.Trades.List(
		ctx,
		cmd.Value("account-id").(string),
		params,
		options...,
	)
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
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "paper-trading:accounts:trades list",
		Transform:      transform,
	})
}
