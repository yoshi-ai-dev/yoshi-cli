// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/yoshi-ai-dev/yoshi-cli/internal/mocktest"
)

func TestSecuritiesOptionsChain(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"securities:options", "chain",
			"--symbol", "symbol",
			"--expiration-date", "7321-69-10",
			"--include-greeks=true",
			"--max-contracts", "1",
			"--moneyness", "all",
			"--month", "1",
			"--option-type", "both",
			"--strike-window-percent", "0",
			"--year", "2000",
		)
	})
}
