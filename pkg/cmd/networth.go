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

var netWorthHistory = cli.Command{
	Name:    "history",
	Usage:   "Get the net worth trend over time, aggregated daily across all accounts.\nLiability accounts (credit, loan) contribute negatively.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "period",
			Usage:     "Trailing window for the trend (all is capped at 1 year)",
			Default:   "6m",
			QueryPath: "period",
		},
	},
	Action:          handleNetWorthHistory,
	HideHelpCommand: true,
}

func handleNetWorthHistory(ctx context.Context, cmd *cli.Command) error {
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

	params := yoshi.NetWorthHistoryParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.NetWorth.History(ctx, params, options...)
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
		Title:          "net-worth history",
		Transform:      transform,
	})
}
