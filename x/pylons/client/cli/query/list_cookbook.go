package query

import (
	"fmt"

	"github.com/MikeSofaer/pylons/x/pylons/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

// ListCookbook queries the cookbooks
func ListCookbook(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list_cookbook",
		Short: "get all cookbooks for a user",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/list_cookbook/%s", queryRoute, cliCtx.GetFromAddress().String()), nil)
			if err != nil {
				return fmt.Errorf(err.Error())
			}

			var out types.CookbookList
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
