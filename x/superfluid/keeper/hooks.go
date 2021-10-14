package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	epochstypes "github.com/osmosis-labs/osmosis/x/epochs/types"
	lockuptypes "github.com/osmosis-labs/osmosis/x/lockup/types"
	"github.com/osmosis-labs/osmosis/x/superfluid/types"
)

func (k Keeper) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
}

func (k Keeper) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
	params := k.GetParams(ctx)
	if epochIdentifier == params.RefreshEpochIdentifier {
		for _, asset := range k.GetAllSuperfluidAssets(ctx) {
			// TODO: should include unlocking asset as well
			// TODO: should we enable all the locks for specific lp token
			// or only locks that people want to participiate in superfluid staking within those locks?
			totalAmt := k.lk.GetPeriodLocksAccumulation(ctx, lockuptypes.QueryCondition{
				LockQueryType: lockuptypes.ByDuration,
				Denom:         asset.Denom,
				Duration:      time.Second,
			})
			k.SetSuperfluidAssetInfo(ctx, types.SuperfluidAssetInfo{
				Denom:                      asset.Denom,
				TotalStakedAmount:          totalAmt,
				RiskAdjustedOsmoEquivalent: k.GetRiskAdjustedOsmoValue(ctx, asset, totalAmt),
			})
		}

		for _, asset := range k.GetAllSuperfluidAssets(ctx) {
			twap := sdk.NewDec(1)
			if asset.AssetType == types.SuperfluidAssetTypeLPShare {
				// TODO: should get twap price from gamm module and use the price
				// potential calculation rule
				// LP_token_Osmo_equivalent = OSMO_amount_on_pool / LP_token_supply
			} else if asset.AssetType == types.SuperfluidAssetTypeNative {
				// TODO: should get twap price from gamm module and use the price
				// which pool should it use to calculate native token price?
			}
			k.SetEpochOsmoEquivalentTWAP(ctx, epochNumber, asset.Denom, twap)
		}
	}
}

// ___________________________________________________________________________________________________

// Hooks wrapper struct for incentives keeper
type Hooks struct {
	k Keeper
}

var _ epochstypes.EpochHooks = Hooks{}

// Return the wrapper struct
func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// epochs hooks
func (h Hooks) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
	h.k.BeforeEpochStart(ctx, epochIdentifier, epochNumber)
}

func (h Hooks) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
	h.k.AfterEpochEnd(ctx, epochIdentifier, epochNumber)
}
