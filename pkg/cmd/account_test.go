// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/yoshi-ai-dev/yoshi-cli/internal/mocktest"
)

func TestAccountsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts", "list",
			"--hidden", "true",
			"--status", "open",
		)
	})
}

func TestAccountsCreateRealEstate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts", "create-real-estate",
			"--address-line1", "x",
			"--account-name", "x",
			"--address-city", "x",
			"--address-state", "x",
			"--address-zip-code", "x",
			"--selected-value", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"address_line1: x\n" +
			"account_name: x\n" +
			"address_city: x\n" +
			"address_state: x\n" +
			"address_zip_code: x\n" +
			"selected_value: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"accounts", "create-real-estate",
		)
	})
}
