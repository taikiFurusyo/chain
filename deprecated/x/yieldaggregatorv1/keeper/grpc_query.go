package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/deprecated/x/yieldaggregatorv1/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) AssetManagementAccount(c context.Context, req *types.QueryAssetManagementAccountRequest) (*types.QueryAssetManagementAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	amAcc := k.GetAssetManagementAccount(ctx, req.Id)
	targets := k.GetAssetManagementTargetsOfAccount(ctx, req.Id)
	return &types.QueryAssetManagementAccountResponse{
		Account: types.AssetManagementAccountInfo{
			Id:                     amAcc.Id,
			Name:                   amAcc.Name,
			AssetManagementTargets: targets,
		},
	}, nil
}

func (k Keeper) AllAssetManagementAccounts(c context.Context, req *types.QueryAllAssetManagementAccountsRequest) (*types.QueryAllAssetManagementAccountsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	amAccs := k.GetAllAssetManagementAccounts(ctx)
	amAccInfos := []types.AssetManagementAccountInfo{}
	for _, amAcc := range amAccs {
		targets := k.GetAssetManagementTargetsOfAccount(ctx, amAcc.Id)
		amAccInfos = append(amAccInfos, types.AssetManagementAccountInfo{
			Id:                     amAcc.Id,
			Name:                   amAcc.Name,
			AssetManagementTargets: targets,
		})
	}
	return &types.QueryAllAssetManagementAccountsResponse{
		Accounts: amAccInfos,
	}, nil
}

func (k Keeper) UserInfo(c context.Context, req *types.QueryUserInfoRequest) (*types.QueryUserInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	orders := k.GetFarmingOrdersOfAddress(ctx, addr)

	deposits := k.GetUserDeposit(ctx, addr)
	return &types.QueryUserInfoResponse{
		UserInfo: types.QueryUserInfo{
			Amount:        deposits,
			FarmingOrders: orders,
			FarmedCounter: 0,
		},
	}, nil
}

func (k Keeper) AllFarmingUnits(c context.Context, req *types.QueryAllFarmingUnitsRequest) (*types.QueryAllFarmingUnitsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryAllFarmingUnitsResponse{
		Units: k.GetAllFarmingUnits(ctx),
	}, nil
}

func (k Keeper) DailyRewardPercents(c context.Context, req *types.QueryDailyRewardPercentsRequest) (*types.QueryDailyRewardPercentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryDailyRewardPercentsResponse{
		DailyPercents: k.GetAllDailyRewardPercents(ctx),
	}, nil
}
