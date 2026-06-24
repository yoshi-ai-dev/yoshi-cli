// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/yoshi-ai-dev/yoshi-cli/internal/mocktest"
)

func TestSecuritiesPriceHistoryList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"securities:price-history", "list",
			"--symbol", "symbol",
			"--cursor", "cursor",
			"--end-date", "7321-69-10",
			"--limit", "1",
			"--start-date", "7321-69-10",
		)
	})
}
