// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
	"github.com/yoshi-ai-dev/yoshi-cli/internal/autocomplete"
	"github.com/yoshi-ai-dev/yoshi-cli/internal/requestflag"
)

var (
	Command            *cli.Command
	CommandErrorBuffer bytes.Buffer
)

func init() {
	Command = &cli.Command{
		Name:      "yoshi",
		Usage:     "CLI for the yoshi API",
		Suggest:   true,
		Version:   Version,
		ErrWriter: &CommandErrorBuffer,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
				Validator: func(baseURL string) error {
					return ValidateBaseURL(baseURL, "--base-url")
				},
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
			&cli.BoolFlag{
				Name:    "raw-output",
				Aliases: []string{"r"},
				Usage:   "If the result is a string, print it without JSON quotes. This can be useful for making output transforms talk to non-JSON-based systems.",
			},
			&requestflag.Flag[string]{
				Name:    "api-key",
				Sources: cli.EnvVars("YOSHI_API_KEY"),
			},
			&cli.StringFlag{
				Name:  "environment",
				Usage: "Set the environment for API requests",
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "accounts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&accountsList,
					&accountsCreateRealEstate,
				},
			},
			{
				Name:     "accounts:balance-series",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&accountsBalanceSeriesList,
				},
			},
			{
				Name:     "transactions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&transactionsList,
					&transactionsChanges,
				},
			},
			{
				Name:     "card-identity-hints",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cardIdentityHintsList,
				},
			},
			{
				Name:     "scores",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&scoresList,
				},
			},
			{
				Name:     "goals",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&goalsCreate,
					&goalsUpdate,
					&goalsList,
					&goalsDelete,
				},
			},
			{
				Name:     "recurring",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&recurringList,
				},
			},
			{
				Name:     "benefits",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&benefitsExpiring,
					&benefitsSummary,
				},
			},
			{
				Name:     "investments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&investmentsList,
					&investmentsHoldingHistory,
					&investmentsHoldings,
					&investmentsPerformance,
					&investmentsTransactions,
				},
			},
			{
				Name:     "benchmarks",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&benchmarksReplay,
				},
			},
			{
				Name:     "trades",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&tradesCreate,
				},
			},
			{
				Name:     "transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&transfersCreate,
				},
			},
			{
				Name:     "income",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&incomeRetrieve,
				},
			},
			{
				Name:     "spending",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&spendingRetrieve,
				},
			},
			{
				Name:     "net-worth",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&netWorthHistory,
				},
			},
			{
				Name:     "credit-debt",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&creditDebtRetrieve,
				},
			},
			{
				Name:     "automations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&automationsList,
				},
			},
			{
				Name:     "briefs",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&briefsRetrieve,
					&briefsList,
				},
			},
			{
				Name:     "securities",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&securitiesRetrieve,
					&securitiesSearch,
				},
			},
			{
				Name:     "securities:options",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&securitiesOptionsChain,
				},
			},
			{
				Name:     "securities:price-history",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&securitiesPriceHistoryList,
				},
			},
			{
				Name:     "me",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&meRetrieve,
					&meSummary,
				},
			},
			{
				Name:     "paper-trading:accounts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&paperTradingAccountsCreate,
					&paperTradingAccountsList,
				},
			},
			{
				Name:     "paper-trading:accounts:holdings",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&paperTradingAccountsHoldingsList,
				},
			},
			{
				Name:     "paper-trading:accounts:trades",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&paperTradingAccountsTradesCreate,
					&paperTradingAccountsTradesList,
				},
			},
			{
				Name:     "webhooks:endpoints",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&webhooksEndpointsCreate,
					&webhooksEndpointsRetrieve,
					&webhooksEndpointsUpdate,
					&webhooksEndpointsList,
					&webhooksEndpointsDelete,
					&webhooksEndpointsRotate,
					&webhooksEndpointsTest,
				},
			},
			{
				Name:     "webhooks:deliveries",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&webhooksDeliveriesList,
				},
			},
			{
				Name:     "webhooks:events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&webhooksEventsList,
				},
			},
			{
				Name:     "webhooks:portal",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&webhooksPortalRetrieve,
				},
			},
			{
				Name:     "approvals",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&approvalsRetrieve,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "yoshi @manpages [-o yoshi.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
		},
		HideHelpCommand: true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "yoshi.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "yoshi.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
