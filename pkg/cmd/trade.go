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

var tradesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create an approval-backed paper or live brokerage trade. The account source\ndetermines the action type.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Required: true,
			BodyPath: "account_id",
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
	Action:          handleTradesCreate,
	HideHelpCommand: true,
}

func handleTradesCreate(ctx context.Context, cmd *cli.Command) error {
	client := yoshi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := yoshi.TradeNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Trades.New(ctx, params, options...)
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
		Title:          "trades create",
		Transform:      transform,
	})
}
