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

var securitiesPriceHistoryList = cli.Command{
	Name:    "list",
	Usage:   "List daily OHLCV/NAV price history for a supported security symbol.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "symbol",
			Usage:     "Ticker symbol, e.g. AAPL",
			Required:  true,
			PathParam: "symbol",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque cursor from a previous response",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "end-date",
			Usage:     "Inclusive price-date upper bound",
			QueryPath: "end_date",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Items per page (1-100, default 50)",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "start-date",
			Usage:     "Inclusive price-date lower bound",
			QueryPath: "start_date",
		},
	},
	Action:          handleSecuritiesPriceHistoryList,
	HideHelpCommand: true,
}

func handleSecuritiesPriceHistoryList(ctx context.Context, cmd *cli.Command) error {
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

	params := yoshi.SecurityPriceHistoryListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Securities.PriceHistory.List(
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
		Title:          "securities:price-history list",
		Transform:      transform,
	})
}
