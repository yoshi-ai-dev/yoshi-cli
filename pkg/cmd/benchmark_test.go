// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/yoshi-ai-dev/yoshi-cli/internal/mocktest"
	"github.com/yoshi-ai-dev/yoshi-cli/internal/requestflag"
)

func TestBenchmarksReplay(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"benchmarks", "replay",
			"--allocation", "{ticker: ticker, weight: 1}",
			"--cash-flow", "{amount: 0, date: '7321-69-10'}",
			"--actual-end-value", "0",
			"--end-date", "7321-69-10",
			"--start-date", "7321-69-10",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(benchmarksReplay)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"benchmarks", "replay",
			"--allocation.ticker", "ticker",
			"--allocation.weight", "1",
			"--cash-flow.amount", "0",
			"--cash-flow.date", "7321-69-10",
			"--actual-end-value", "0",
			"--end-date", "7321-69-10",
			"--start-date", "7321-69-10",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"allocations:\n" +
			"  - ticker: ticker\n" +
			"    weight: 1\n" +
			"cash_flows:\n" +
			"  - amount: 0\n" +
			"    date: '7321-69-10'\n" +
			"actual_end_value: 0\n" +
			"end_date: '7321-69-10'\n" +
			"start_date: '7321-69-10'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"benchmarks", "replay",
		)
	})
}
