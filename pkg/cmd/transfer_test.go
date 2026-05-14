// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/yoshi-ai-dev/yoshi-cli/internal/mocktest"
)

func TestTransfersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"transfers", "create",
			"--amount", "1",
			"--from-id", "x",
			"--method", "bank_transfer",
			"--to-id", "x",
			"--currency-code", "USD",
			"--description", "description",
			"--request-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: 1\n" +
			"from_id: x\n" +
			"method: bank_transfer\n" +
			"to_id: x\n" +
			"currency_code: USD\n" +
			"description: description\n" +
			"request_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"transfers", "create",
		)
	})
}
