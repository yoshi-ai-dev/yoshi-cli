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

var securitiesOptionsChain = cli.Command{
	Name:    "chain",
	Usage:   "Get a bounded delayed options chain slice for an equity symbol.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "symbol",
			Usage:     "Ticker symbol, e.g. AAPL",
			Required:  true,
			PathParam: "symbol",
		},
		&requestflag.Flag[string]{
			Name:      "expiration-date",
			Usage:     "Exact option expiration date as YYYY-MM-DD.",
			QueryPath: "expiration_date",
		},
		&requestflag.Flag[bool]{
			Name:      "include-greeks",
			Default:   false,
			QueryPath: "include_greeks",
		},
		&requestflag.Flag[int64]{
			Name:      "max-contracts",
			Usage:     "Maximum returned contracts. Defaults to 20, capped at 50.",
			Default:   20,
			QueryPath: "max_contracts",
		},
		&requestflag.Flag[string]{
			Name:      "moneyness",
			Usage:     `Allowed values: "all", "itm", "atm", "otm".`,
			Default:   "all",
			QueryPath: "moneyness",
		},
		&requestflag.Flag[int64]{
			Name:      "month",
			Usage:     "Expiration month.",
			QueryPath: "month",
		},
		&requestflag.Flag[string]{
			Name:      "option-type",
			Usage:     `Allowed values: "both", "calls", "puts".`,
			Default:   "both",
			QueryPath: "option_type",
		},
		&requestflag.Flag[*float64]{
			Name:      "strike-window-percent",
			Default:   requestflag.Ptr[float64](25),
			QueryPath: "strike_window_percent",
		},
		&requestflag.Flag[int64]{
			Name:      "year",
			Usage:     "Expiration year.",
			QueryPath: "year",
		},
	},
	Action:          handleSecuritiesOptionsChain,
	HideHelpCommand: true,
}

func handleSecuritiesOptionsChain(ctx context.Context, cmd *cli.Command) error {
	client := yoshi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("symbol") && len(unusedArgs) > 0 {
		cmd.Set("symbol", unusedArgs[0])
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

	params := yoshi.SecurityOptionChainParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Securities.Options.Chain(
		ctx,
		cmd.Value("symbol").(string),
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
		Title:          "securities:options chain",
		Transform:      transform,
	})
}
