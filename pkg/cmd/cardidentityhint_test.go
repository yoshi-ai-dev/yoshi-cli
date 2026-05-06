// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/yoshi-ai-dev/yoshi-cli/internal/mocktest"
)

func TestCardIdentityHintsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-identity-hints", "list",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--status", "open",
		)
	})
}
