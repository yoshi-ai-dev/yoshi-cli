// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/yoshi-ai-dev/yoshi-cli/internal/mocktest"
)

func TestInvestmentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"investments", "list",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--category", "category",
			"--sort-by", "symbol",
			"--sort-dir", "asc",
		)
	})
}

func TestInvestmentsHoldingHistory(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"investments", "holding-history",
			"--max-items", "10",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--cursor", "cursor",
			"--end-date", "7321-69-10",
			"--limit", "1",
			"--origin", "from_source",
			"--security-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--start-date", "7321-69-10",
		)
	})
}

func TestInvestmentsHoldings(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"investments", "holdings",
			"--max-items", "10",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--cursor", "cursor",
			"--limit", "1",
		)
	})
}

func TestInvestmentsPerformance(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"investments", "performance",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--period", "1w",
		)
	})
}

func TestInvestmentsTaxLots(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"investments", "tax-lots",
			"--max-items", "10",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--cursor", "cursor",
			"--limit", "1",
			"--security-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--symbol", "x",
		)
	})
}

func TestInvestmentsTransactions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"investments", "transactions",
			"--max-items", "10",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--cursor", "cursor",
			"--end-date", "7321-69-10",
			"--fees-only=true",
			"--limit", "1",
			"--security-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--start-date", "7321-69-10",
			"--subtype", "account fee",
			"--type", "buy",
		)
	})
}
