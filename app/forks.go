package app

import (
	v3 "github.com/MonikaCat/osmosis/v6/app/upgrades/v3"
	v6 "github.com/MonikaCat/osmosis/v6/app/upgrades/v6"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlockForks is intended to be ran in
func BeginBlockForks(ctx sdk.Context, app *OsmosisApp) {
	switch ctx.BlockHeight() {
	case v3.UpgradeHeight:
		v3.RunForkLogic(ctx, app.GovKeeper, app.StakingKeeper)
	case v6.UpgradeHeight:
		v6.RunForkLogic(ctx)
	default:
		// do nothing
		return
	}
}
