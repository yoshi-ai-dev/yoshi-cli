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

var transactionsList = cli.Command{
	Name:    "list",
	Usage:   "Sync transactions for the authenticated user with cursor-based pagination and\noptional date backfill filters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Filter by account ID",
			QueryPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque cursor from a previous response",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "end-date",
			Usage:     "Inclusive posted-date upper bound",
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
			Usage:     "Inclusive posted-date lower bound",
			QueryPath: "start_date",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleTransactionsList,
	HideHelpCommand: true,
}

var transactionsChanges = cli.Command{
	Name:    "changes",
	Usage:   "List transaction upserts and removals ordered by change time. Use this for\nincremental sync; use /transactions for current-state backfills.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Filter by account ID",
			QueryPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque cursor from a previous response",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "end-date",
			Usage:     "Inclusive posted-date upper bound",
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
			Usage:     "Inclusive posted-date lower bound",
			QueryPath: "start_date",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleTransactionsChanges,
	HideHelpCommand: true,
}

func handleTransactionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := yoshi.TransactionListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Transactions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "transactions list",
			Transform:      transform,
		})
	} else {
		iter := client.Transactions.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "transactions list",
			Transform:      transform,
		})
	}
}

func handleTransactionsChanges(ctx context.Context, cmd *cli.Command) error {
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

	params := yoshi.TransactionChangesParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Transactions.Changes(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "transactions changes",
			Transform:      transform,
		})
	} else {
		iter := client.Transactions.ChangesAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "transactions changes",
			Transform:      transform,
		})
	}
}
