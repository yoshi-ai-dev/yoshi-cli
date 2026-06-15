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

var spendingRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get the spending breakdown by category for a period. Spending counts outflow\ntransactions excluding internal transfers.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Filter by account ID",
			QueryPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:      "group-by",
			Usage:     "Category granularity to group by",
			Default:   "tier1",
			QueryPath: "group_by",
		},
		&requestflag.Flag[string]{
			Name:      "period",
			Usage:     "Trailing window for the breakdown",
			Default:   "1m",
			QueryPath: "period",
		},
		&requestflag.Flag[int64]{
			Name:      "top-n",
			Usage:     "Limit to the top N groups",
			QueryPath: "top_n",
		},
	},
	Action:          handleSpendingRetrieve,
	HideHelpCommand: true,
}

func handleSpendingRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := yoshi.SpendingGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Spending.Get(ctx, params, options...)
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
		Title:          "spending retrieve",
		Transform:      transform,
	})
}
