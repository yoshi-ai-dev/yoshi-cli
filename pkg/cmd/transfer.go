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

var transfersCreate = cli.Command{
	Name:    "create",
	Usage:   "Create an approval-backed bank transfer (ACH). Omitted method defaults to\nbank_transfer.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "amount",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "from-id",
			Required: true,
			BodyPath: "from_id",
		},
		&requestflag.Flag[string]{
			Name:     "to-id",
			Required: true,
			BodyPath: "to_id",
		},
		&requestflag.Flag[string]{
			Name:     "currency-code",
			Usage:    `Allowed values: "USD".`,
			Default:  "USD",
			BodyPath: "currency_code",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "method",
			Usage:    `Allowed values: "bank_transfer".`,
			Default:  "bank_transfer",
			BodyPath: "method",
		},
		&requestflag.Flag[string]{
			Name:     "request-id",
			BodyPath: "request_id",
		},
	},
	Action:          handleTransfersCreate,
	HideHelpCommand: true,
}

func handleTransfersCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := yoshi.TransferNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Transfers.New(ctx, params, options...)
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
		Title:          "transfers create",
		Transform:      transform,
	})
}
