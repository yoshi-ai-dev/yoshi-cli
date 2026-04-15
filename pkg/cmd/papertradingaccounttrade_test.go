// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/yoshi-ai-dev/yoshi-cli/internal/mocktest"
)

func TestPaperTradingAccountsTradesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"paper-trading:accounts:trades", "create",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--side", "buy",
			"--symbol", "x",
			"--notional", "1",
			"--quantity", "1",
			"--skip-approval=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"side: buy\n" +
			"symbol: x\n" +
			"notional: 1\n" +
			"quantity: 1\n" +
			"skip_approval: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"paper-trading:accounts:trades", "create",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPaperTradingAccountsTradesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"paper-trading:accounts:trades", "list",
			"--max-items", "10",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--cursor", "cursor",
			"--limit", "1",
		)
	})
}
