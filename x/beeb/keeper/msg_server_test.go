package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/Benzbeeb/beeb/x/beeb/types"
    "github.com/Benzbeeb/beeb/x/beeb/keeper"
    keepertest "github.com/Benzbeeb/beeb/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BeebKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
