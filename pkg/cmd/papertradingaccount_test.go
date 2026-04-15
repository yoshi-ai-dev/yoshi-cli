// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/yoshi-ai-dev/yoshi-cli/internal/mocktest"
)

func TestPaperTradingAccountsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"paper-trading:accounts", "create",
			"--name", "x",
			"--skip-approval=true",
			"--starting-cash-balance", "321669910225",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: x\n" +
			"skip_approval: true\n" +
			"starting_cash_balance: '321669910225'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"paper-trading:accounts", "create",
		)
	})
}

func TestPaperTradingAccountsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"paper-trading:accounts", "list",
		)
	})
}
