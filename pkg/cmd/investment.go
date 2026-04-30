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

var investmentsList = cli.Command{
	Name:    "list",
	Usage:   "Get investment holdings grouped by asset class.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			QueryPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:      "category",
			Usage:     "Filter by asset class",
			QueryPath: "category",
		},
		&requestflag.Flag[string]{
			Name:      "sort-by",
			Usage:     `Allowed values: "symbol", "last_price", "day_change", "total_gain_loss", "current_value", "quantity", "cost_basis_total".`,
			QueryPath: "sort_by",
		},
		&requestflag.Flag[string]{
			Name:      "sort-dir",
			Usage:     `Allowed values: "asc", "desc".`,
			QueryPath: "sort_dir",
		},
	},
	Action:          handleInvestmentsList,
	HideHelpCommand: true,
}

func handleInvestmentsList(ctx context.Context, cmd *cli.Command) error {
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

	params := yoshi.InvestmentListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Investments.List(ctx, params, options...)
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
		Title:          "investments list",
		Transform:      transform,
	})
}
