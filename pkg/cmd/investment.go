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

var investmentsHoldingHistory = cli.Command{
	Name:    "holding-history",
	Usage:   "List historical account holding snapshots for investment accounts.",
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
			Usage:     "Inclusive as-of date upper bound",
			QueryPath: "end_date",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Items per page (1-100, default 50)",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "origin",
			Usage:     `Allowed values: "from_source", "reconstructed", "trade".`,
			QueryPath: "origin",
		},
		&requestflag.Flag[string]{
			Name:      "security-id",
			Usage:     "Filter by security ID",
			QueryPath: "security_id",
		},
		&requestflag.Flag[string]{
			Name:      "start-date",
			Usage:     "Inclusive as-of date lower bound",
			QueryPath: "start_date",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleInvestmentsHoldingHistory,
	HideHelpCommand: true,
}

var investmentsHoldings = cli.Command{
	Name:    "holdings",
	Usage:   "List current investment holdings per account and security with identifiers and\nfee metadata.",
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
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Items per page (1-100, default 50)",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleInvestmentsHoldings,
	HideHelpCommand: true,
}

var investmentsPerformance = cli.Command{
	Name:    "performance",
	Usage:   "Get investment portfolio performance, TWR series, realized gains, unrealized\ngains, and income.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			QueryPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:      "period",
			Usage:     `Allowed values: "1w", "1m", "3m", "ytd", "1y", "all".`,
			Default:   "all",
			QueryPath: "period",
		},
	},
	Action:          handleInvestmentsPerformance,
	HideHelpCommand: true,
}

var investmentsTaxLots = cli.Command{
	Name:    "tax-lots",
	Usage:   "List normalized current/open tax lots with cost basis, holding period,\nunrealized gain/loss, and coverage metadata.",
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
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Items per page (1-100, default 50)",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "security-id",
			Usage:     "Filter by security ID",
			QueryPath: "security_id",
		},
		&requestflag.Flag[string]{
			Name:      "symbol",
			Usage:     "Filter by ticker symbol",
			QueryPath: "symbol",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleInvestmentsTaxLots,
	HideHelpCommand: true,
}

var investmentsTransactions = cli.Command{
	Name:    "transactions",
	Usage:   "List investment transactions with security identifiers and explicit fee fields.",
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
			Usage:     "Inclusive transaction date upper bound",
			QueryPath: "end_date",
		},
		&requestflag.Flag[bool]{
			Name:      "fees-only",
			Usage:     "Return only explicit fee transactions or rows with fees",
			QueryPath: "fees_only",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Items per page (1-100, default 50)",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "security-id",
			Usage:     "Filter by security ID",
			QueryPath: "security_id",
		},
		&requestflag.Flag[string]{
			Name:      "start-date",
			Usage:     "Inclusive transaction date lower bound",
			QueryPath: "start_date",
		},
		&requestflag.Flag[string]{
			Name:      "subtype",
			Usage:     `Allowed values: "account fee", "adjustment", "assignment", "buy", "buy to cover", "contribution", "deposit", "distribution", "dividend", "dividend reinvestment", "exercise", "expire", "fund fee", "interest", "interest receivable", "interest reinvestment", "legal fee", "loan payment", "long-term capital gain", "long-term capital gain reinvestment", "management fee", "margin expense", "merger", "miscellaneous fee", "non-qualified dividend", "non-resident tax", "pending credit", "pending debit", "qualified dividend", "rebalance", "return of principal", "request", "sell", "sell short", "send", "short-term capital gain", "short-term capital gain reinvestment", "spin off", "split", "stock distribution", "tax", "tax withheld", "trade", "transfer", "transfer fee", "trust fee", "unqualified gain", "withdrawal".`,
			QueryPath: "subtype",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     `Allowed values: "buy", "sell", "cancel", "cash", "fee", "transfer".`,
			QueryPath: "type",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleInvestmentsTransactions,
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

func handleInvestmentsHoldingHistory(ctx context.Context, cmd *cli.Command) error {
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

	params := yoshi.InvestmentHoldingHistoryParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Investments.HoldingHistory(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "investments holding-history",
			Transform:      transform,
		})
	} else {
		iter := client.Investments.HoldingHistoryAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "investments holding-history",
			Transform:      transform,
		})
	}
}

func handleInvestmentsHoldings(ctx context.Context, cmd *cli.Command) error {
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

	params := yoshi.InvestmentHoldingsParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Investments.Holdings(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "investments holdings",
			Transform:      transform,
		})
	} else {
		iter := client.Investments.HoldingsAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "investments holdings",
			Transform:      transform,
		})
	}
}

func handleInvestmentsPerformance(ctx context.Context, cmd *cli.Command) error {
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

	params := yoshi.InvestmentPerformanceParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Investments.Performance(ctx, params, options...)
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
		Title:          "investments performance",
		Transform:      transform,
	})
}

func handleInvestmentsTaxLots(ctx context.Context, cmd *cli.Command) error {
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

	params := yoshi.InvestmentTaxLotsParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Investments.TaxLots(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "investments tax-lots",
			Transform:      transform,
		})
	} else {
		iter := client.Investments.TaxLotsAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "investments tax-lots",
			Transform:      transform,
		})
	}
}

func handleInvestmentsTransactions(ctx context.Context, cmd *cli.Command) error {
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

	params := yoshi.InvestmentTransactionsParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Investments.Transactions(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "investments transactions",
			Transform:      transform,
		})
	} else {
		iter := client.Investments.TransactionsAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "investments transactions",
			Transform:      transform,
		})
	}
}
