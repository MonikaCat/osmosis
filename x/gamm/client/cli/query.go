package cli

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/c-osmosis/osmosis/x/gamm/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group gamm queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdPool(),
		GetCmdPools(),
		GetCmdPoolParams(),
		GetCmdTotalShare(),
		GetCmdRecords(),
		GetCmdSpotPrice(),
		GetCmdEstimateSwapExactAmountIn(),
		GetCmdEstimateSwapExactAmountOut(),
	)

	return cmd
}

// GetCmdPool returns pool
func GetCmdPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pool <poolID>",
		Short: "Query pool",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query pool.
Example:
$ %s query gamm pool 1
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			poolID, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.Pool(cmd.Context(), &types.QueryPoolRequest{
				PoolId: uint64(poolID),
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdPools return pools
func GetCmdPools() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pools",
		Short: "Query pools",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query pools.
Example:
$ %s query gamm pools
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Pools(cmd.Context(), &types.QueryPoolsRequest{
				Pagination: nil,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdPoolParams return pool params
func GetCmdPoolParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pool-params <poolID>",
		Short: "Query pool-params",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query pool-params.
Example:
$ %s query gamm pool-params 1
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			poolID, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.PoolParams(cmd.Context(), &types.QueryPoolParamsRequest{
				PoolId: uint64(poolID),
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdTotalShare return total share
func GetCmdTotalShare() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-share <poolID>",
		Short: "Query total-share",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query total-share.
Example:
$ %s query gamm total-share 1
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			poolID, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.TotalShare(cmd.Context(), &types.QueryTotalShareRequest{
				PoolId: uint64(poolID),
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdRecords return records
func GetCmdRecords() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "records <poolID>",
		Short: "Query records",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query records.
Example:
$ %s query gamm records 1
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			poolID, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.Records(cmd.Context(), &types.QueryRecordsRequest{
				PoolId: uint64(poolID),
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdSpotPrice returns spot price
func GetCmdSpotPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "spot-price <poolID> <tokenInDenom> <tokenOutDenom>",
		Short: "Query spot-price",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query spot-price.
Example:
$ %s query gamm spot-price 1 stake stake2
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			poolID, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.SpotPrice(cmd.Context(), &types.QuerySpotPriceRequest{
				PoolId:        uint64(poolID),
				TokenInDenom:  args[1],
				TokenOutDenom: args[2],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdEstimateSwapExactAmountIn returns estimation of output coin when amount of x token input
func GetCmdEstimateSwapExactAmountIn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate-swap-exact-amount-in <poolID> <sender> <tokenIn>",
		Short: "Query estimate-swap-exact-amount-in",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query estimate-swap-exact-amount-in.
Example:
$ %s query gamm estimate-swap-exact-amount-in 1 osm11vmx8jtggpd9u7qr0t8vxclycz85u925sazglr7 stake --swap-route-pool-ids=2 --swap-route-amounts=100stake2 --swap-route-pool-ids=3 --swap-route-amounts=100stake
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			poolID, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			routes, err := swapAmountInRoutes(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.EstimateSwapExactAmountIn(cmd.Context(), &types.QuerySwapExactAmountInRequest{
				Sender:  args[1],        // TODO: where sender is used?
				PoolId:  uint64(poolID), // TODO: is this poolId used?
				TokenIn: args[2],
				Routes:  routes,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().AddFlagSet(FlagSetQuerySwapRoutes())
	flags.AddQueryFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(FlagSwapRoutePoolIds)
	_ = cmd.MarkFlagRequired(FlagSwapRouteAmounts)

	return cmd
}

// GetCmdEstimateSwapExactAmountOut returns estimation of input coin to get exact amount of x token output
func GetCmdEstimateSwapExactAmountOut() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate-swap-exact-amount-out <poolID> <sender> <tokenOut>",
		Short: "Query estimate-swap-exact-amount-out",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query estimate-swap-exact-amount-out.
Example:
$ %s query gamm estimate-swap-exact-amount-out 1 osm11vmx8jtggpd9u7qr0t8vxclycz85u925sazglr7 stake --swap-route-pool-ids=2 --swap-route-amounts=100stake2 --swap-route-pool-ids=3 --swap-route-amounts=100stake
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			poolID, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			swapRoutePoolIds, err := cmd.Flags().GetStringArray(FlagSwapRoutePoolIds)
			if err != nil {
				return err
			}

			swapRouteAmounts, err := cmd.Flags().GetStringArray(FlagSwapRouteAmounts)
			if err != nil {
				return err
			}

			if len(swapRoutePoolIds) != len(swapRouteAmounts) {
				return errors.New("swap route pool ids and amounts mismatch")
			}

			routes := []types.SwapAmountOutRoute{}
			for index, poolIDStr := range swapRoutePoolIds {
				pID, err := strconv.Atoi(poolIDStr)
				if err != nil {
					return err
				}
				routes = append(routes, types.SwapAmountOutRoute{
					PoolId:       uint64(pID),
					TokenInDenom: swapRouteAmounts[index],
				})
			}

			res, err := queryClient.EstimateSwapExactAmountOut(cmd.Context(), &types.QuerySwapExactAmountOutRequest{
				Sender:   args[1],        // TODO: where sender is used?
				PoolId:   uint64(poolID), // TODO: is this poolId used?
				Routes:   routes,
				TokenOut: args[2],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().AddFlagSet(FlagSetQuerySwapRoutes())
	flags.AddQueryFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(FlagSwapRoutePoolIds)
	_ = cmd.MarkFlagRequired(FlagSwapRouteAmounts)

	return cmd
}
