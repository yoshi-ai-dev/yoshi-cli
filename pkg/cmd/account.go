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

var accountsList = cli.Command{
	Name:    "list",
	Usage:   "List linked financial accounts with display metadata and public lifecycle\nstatus.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "hidden",
			Usage:     `Allowed values: "true", "false".`,
			QueryPath: "hidden",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     `Allowed values: "open", "closed", "all".`,
			Default:   "open",
			QueryPath: "status",
		},
	},
	Action:          handleAccountsList,
	HideHelpCommand: true,
}

var accountsCreateRealEstate = cli.Command{
	Name:    "create-real-estate",
	Usage:   "Create a real estate property account directly from an address. If\nselected_value is omitted, Yoshi estimates the property value with web search\nbefore creating the account.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "address-line1",
			Required: true,
			BodyPath: "address_line1",
		},
		&requestflag.Flag[string]{
			Name:     "account-name",
			BodyPath: "account_name",
		},
		&requestflag.Flag[string]{
			Name:     "address-city",
			BodyPath: "address_city",
		},
		&requestflag.Flag[string]{
			Name:     "address-state",
			BodyPath: "address_state",
		},
		&requestflag.Flag[string]{
			Name:     "address-zip-code",
			BodyPath: "address_zip_code",
		},
		&requestflag.Flag[float64]{
			Name:     "selected-value",
			BodyPath: "selected_value",
		},
	},
	Action:          handleAccountsCreateRealEstate,
	HideHelpCommand: true,
}

func handleAccountsList(ctx context.Context, cmd *cli.Command) error {
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

	params := yoshi.AccountListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Accounts.List(ctx, params, options...)
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
		Title:          "accounts list",
		Transform:      transform,
	})
}

func handleAccountsCreateRealEstate(ctx context.Context, cmd *cli.Command) error {
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

	params := yoshi.AccountNewRealEstateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Accounts.NewRealEstate(ctx, params, options...)
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
		Title:          "accounts create-real-estate",
		Transform:      transform,
	})
}
