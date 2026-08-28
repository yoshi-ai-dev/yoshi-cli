// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
	"github.com/yoshi-ai-dev/yoshi-cli/internal/apiquery"
	"github.com/yoshi-ai-dev/yoshi-cli/internal/requestflag"
	"github.com/yoshi-ai-dev/yoshi-go"
	"github.com/yoshi-ai-dev/yoshi-go/option"
)

var goalsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new financial goal (cash savings, investment, or debt payoff).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "goal-type",
			Usage:    `Allowed values: "cash", "investment", "debt_payoff".`,
			Required: true,
			BodyPath: "goal_type",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[*float64]{
			Name:     "target-amount",
			Required: true,
			BodyPath: "target_amount",
		},
		&requestflag.Flag[[]string]{
			Name:     "linked-account-id",
			BodyPath: "linked_account_ids",
		},
		&requestflag.Flag[*int64]{
			Name:     "priority",
			BodyPath: "priority",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    `Allowed values: "active", "paused", "completed", "cancelled".`,
			BodyPath: "status",
		},
		&requestflag.Flag[string]{
			Name:     "target-date",
			BodyPath: "target_date",
		},
	},
	Action:          handleGoalsCreate,
	HideHelpCommand: true,
}

var goalsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update an existing goal. Only provided fields change; `target_date: null` clears\nthe date.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[*float64]{
			Name:     "current-value",
			BodyPath: "current_value",
		},
		&requestflag.Flag[string]{
			Name:     "goal-type",
			Usage:    `Allowed values: "cash", "investment", "debt_payoff".`,
			BodyPath: "goal_type",
		},
		&requestflag.Flag[[]string]{
			Name:     "linked-account-id",
			BodyPath: "linked_account_ids",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[*int64]{
			Name:     "priority",
			BodyPath: "priority",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    `Allowed values: "active", "paused", "completed", "cancelled".`,
			BodyPath: "status",
		},
		&requestflag.Flag[*float64]{
			Name:     "target-amount",
			BodyPath: "target_amount",
		},
		&requestflag.Flag[*string]{
			Name:     "target-date",
			BodyPath: "target_date",
		},
	},
	Action:          handleGoalsUpdate,
	HideHelpCommand: true,
}

var goalsList = cli.Command{
	Name:    "list",
	Usage:   "List financial goals with optional status filter.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     `Allowed values: "active", "paused", "completed", "cancelled".`,
			QueryPath: "status",
		},
	},
	Action:          handleGoalsList,
	HideHelpCommand: true,
}

var goalsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an existing goal. System-managed goals cannot be deleted.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleGoalsDelete,
	HideHelpCommand: true,
}

func handleGoalsCreate(ctx context.Context, cmd *cli.Command) error {
	client := yoshi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := yoshi.GoalNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Goals.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "goals create",
		Transform:      transform,
	})
}

func handleGoalsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := yoshi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := yoshi.GoalUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Goals.Update(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "goals update",
		Transform:      transform,
	})
}

func handleGoalsList(ctx context.Context, cmd *cli.Command) error {
	client := yoshi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := yoshi.GoalListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Goals.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "goals list",
		Transform:      transform,
	})
}

func handleGoalsDelete(ctx context.Context, cmd *cli.Command) error {
	client := yoshi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Goals.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "goals delete",
		Transform:      transform,
	})
}
