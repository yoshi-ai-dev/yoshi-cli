// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/yoshi-ai-dev/yoshi-cli/internal/mocktest"
)

func TestGoalsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"goals", "create",
			"--goal-type", "cash",
			"--name", "x",
			"--target-amount", "0",
			"--linked-account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--priority", "0",
			"--status", "active",
			"--target-date", "7321-69-10",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"goal_type: cash\n" +
			"name: x\n" +
			"target_amount: 0\n" +
			"linked_account_ids:\n" +
			"  - 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"priority: 0\n" +
			"status: active\n" +
			"target_date: '7321-69-10'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"goals", "create",
		)
	})
}

func TestGoalsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"goals", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--current-value", "0",
			"--goal-type", "cash",
			"--linked-account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--name", "x",
			"--priority", "0",
			"--status", "active",
			"--target-amount", "0",
			"--target-date", "7321-69-10",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"current_value: 0\n" +
			"goal_type: cash\n" +
			"linked_account_ids:\n" +
			"  - 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"name: x\n" +
			"priority: 0\n" +
			"status: active\n" +
			"target_amount: 0\n" +
			"target_date: '7321-69-10'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"goals", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestGoalsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"goals", "list",
			"--status", "active",
		)
	})
}

func TestGoalsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"goals", "delete",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
