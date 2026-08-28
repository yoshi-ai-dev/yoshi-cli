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

var benchmarksReplay = requestflag.WithInnerFlags(cli.Command{
	Name:    "replay",
	Usage:   "Replay arbitrary supported benchmark tickers against a cash-flow chronology and\nreturn benchmark value series, return metrics, alpha inputs, and data-quality\nflags.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "allocation",
			Required: true,
			BodyPath: "allocations",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "cash-flow",
			Required: true,
			BodyPath: "cash_flows",
		},
		&requestflag.Flag[float64]{
			Name:     "actual-end-value",
			Usage:    "Optional actual ending portfolio value used to calculate dollar and percent alpha.",
			BodyPath: "actual_end_value",
		},
		&requestflag.Flag[string]{
			Name:     "end-date",
			BodyPath: "end_date",
		},
		&requestflag.Flag[string]{
			Name:     "start-date",
			BodyPath: "start_date",
		},
	},
	Action:          handleBenchmarksReplay,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"allocation": {
		&requestflag.InnerFlag[string]{
			Name:       "allocation.ticker",
			InnerField: "ticker",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "allocation.weight",
			InnerField: "weight",
		},
	},
	"cash-flow": {
		&requestflag.InnerFlag[float64]{
			Name:       "cash-flow.amount",
			Usage:      "External cash flow amount. Contributions are positive; withdrawals are negative.",
			InnerField: "amount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "cash-flow.date",
			InnerField: "date",
		},
	},
})

func handleBenchmarksReplay(ctx context.Context, cmd *cli.Command) error {
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

	params := yoshi.BenchmarkReplayParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Benchmarks.Replay(ctx, params, options...)
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
		Title:          "benchmarks replay",
		Transform:      transform,
	})
}
