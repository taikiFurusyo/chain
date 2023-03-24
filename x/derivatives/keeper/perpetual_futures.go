package keeper

import (
	"errors"
	"fmt"

	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifiTypes "github.com/UnUniFi/chain/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

// fixme: it has not been tested
// todo:rename GetCurrentPrice to GetCurrentUsdPrice
func (k Keeper) GetCurrentPrice(ctx sdk.Context, denom string) (sdk.Dec, error) {
	ticker, err := k.pricefeedKeeper.GetTicker(ctx, denom)
	if err != nil {
		return sdk.Dec{}, err
	}
	rate, err := k.GetPrice(ctx, ticker, "usd")
	if err != nil {
		return sdk.Dec{}, err
	}
	return rate.Price, nil
}

func (k Keeper) GetPairUsdPrice(ctx sdk.Context, base, quote string) (sdk.Dec, sdk.Dec, error) {
	baseUsdPrice, err := k.GetCurrentPrice(ctx, base)
	if err != nil {
		return sdk.Dec{}, sdk.Dec{}, err
	}
	quoteUsdPrice, err := k.GetCurrentPrice(ctx, quote)
	if err != nil {
		return sdk.Dec{}, sdk.Dec{}, err
	}
	return baseUsdPrice, quoteUsdPrice, nil
}

func (k Keeper) GetPairUsdPriceFromMarket(ctx sdk.Context, market types.Market) (sdk.Dec, sdk.Dec, error) {
	return k.GetPairUsdPrice(ctx, market.BaseDenom, market.QuoteDenom)
}

func (k Keeper) OpenPerpetualFuturesPosition(ctx sdk.Context, positionId string, sender ununifiTypes.StringAccAddress, margin sdk.Coin, market types.Market, positionInstance types.PerpetualFuturesPositionInstance) (*types.Position, error) {
	openedBaseRate, err := k.GetCurrentPrice(ctx, market.BaseDenom)
	if err != nil {
		return nil, err
	}

	openedQuoteRate, err := k.GetCurrentPrice(ctx, market.QuoteDenom)
	if err != nil {
		return nil, err
	}
	any, err := codecTypes.NewAnyWithValue(&positionInstance)
	if err != nil {
		return nil, err
	}

	position := types.Position{
		Id:               positionId,
		Market:           market,
		Address:          sender,
		OpenedAt:         ctx.BlockTime(),
		OpenedHeight:     uint64(ctx.BlockHeight()),
		OpenedBaseRate:   openedBaseRate,
		OpenedQuoteRate:  openedQuoteRate,
		PositionInstance: *any,
		RemainingMargin:  margin,
	}

	if position.Validate() != nil {
		return nil, types.ErrorMarginNotEnough
	}

	params := k.GetParams(ctx)
	if position.NeedLiquidation(params.PerpetualFutures.MarginMaintenanceRate, position.OpenedBaseRate, position.OpenedQuoteRate) {
		return nil, types.ErrorInvalidPositionParams
	}

	switch positionInstance.PositionType {
	case types.PositionType_LONG:
		k.AddPerpetualFuturesNetPositionOfMarket(ctx, market, positionInstance.Size_)
		break
	case types.PositionType_SHORT:
		k.SubPerpetualFuturesNetPositionOfMarket(ctx, market, positionInstance.Size_)
		break
	case types.PositionType_POSITION_UNKNOWN:
		return nil, fmt.Errorf("unknown position type")
	}

	ctx.EventManager().EmitTypedEvent(&types.EventPerpetualFuturesPositionOpened{
		Sender:     sender.AccAddress().String(),
		PositionId: positionId,
	})

	return &position, nil
}

func (k Keeper) ClosePerpetualFuturesPosition(ctx sdk.Context, position types.PerpetualFuturesPosition) error {
	params := k.GetParams(ctx)
	commissionRate := params.PerpetualFutures.CommissionRate
	feeAmountDec := position.PositionInstance.Size_.Mul(commissionRate)
	tradeAmount := position.PositionInstance.Size_.Sub(feeAmountDec)
	feeAmount := feeAmountDec.RoundInt()

	baseUsdPrice, err := k.GetCurrentPrice(ctx, position.Market.BaseDenom)
	if err != nil {
		return err
	}
	quoteUsdPrice, err := k.GetCurrentPrice(ctx, position.Market.QuoteDenom)
	if err != nil {
		return err
	}

	// TODO: this is wrong. refer to Issue#407
	if !feeAmount.IsZero() {
		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, position.Address.AccAddress(), types.ModuleName, sdk.Coins{sdk.NewCoin(position.Market.BaseDenom, feeAmount)})
		if err != nil {
			return err
		}
	}

	returningAmount, lossToLP := position.CalcReturningAmountAtClose(baseUsdPrice, quoteUsdPrice)

	if !(lossToLP.IsNil()) {
		// TODO: emit event to tell how much loss is taken by liquidity provider.
	}

	returningCoin := sdk.NewCoin(position.RemainingMargin.Denom, returningAmount)
	if returningCoin.IsPositive() {
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, position.Address.AccAddress(), sdk.Coins{returningCoin}); err != nil {
			return err
		}
	}

	ctx.EventManager().EmitTypedEvent(&types.EventPerpetualFuturesPositionClosed{
		Sender:      position.Address.AccAddress().String(),
		PositionId:  position.Id,
		FeeAmount:   feeAmount.String(),
		TradeAmount: tradeAmount.String(),
	})

	return nil
}

func (k Keeper) ReportLiquidationNeededPerpetualFuturesPosition(ctx sdk.Context, rewardRecipient ununifiTypes.StringAccAddress, position types.PerpetualFuturesPosition) error {
	params := k.GetParams(ctx)

	currentBaseUsdRate, currentQuoteUsdRate, err := k.GetPairUsdPriceFromMarket(ctx, position.Market)
	if err != nil {
		panic(err)
	}

	if position.NeedLiquidation(params.PerpetualFutures.MarginMaintenanceRate, currentBaseUsdRate, currentQuoteUsdRate) {
		k.ClosePerpetualFuturesPosition(ctx, position)

		rewardAmount := sdk.NewDecFromInt(position.RemainingMargin.Amount).Mul(params.PoolParams.ReportLiquidationRewardRate).RoundInt()
		reward := sdk.NewCoins(sdk.NewCoin(position.RemainingMargin.Denom, rewardAmount))
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, rewardRecipient.AccAddress(), reward)
		if err != nil {
			return err
		}

		ctx.EventManager().EmitTypedEvent(&types.EventPerpetualFuturesPositionLiquidated{
			RewardRecipient: rewardRecipient.AccAddress().String(),
			PositionId:      position.Id,
			RemainingMargin: position.RemainingMargin.String(),
			RewardAmount:    rewardAmount.String(),
		})
		return nil
	}

	return errors.New("no liquidation needed")
}

func (k Keeper) ReportLevyPeriodPerpetualFuturesPosition(ctx sdk.Context, rewardRecipient ununifiTypes.StringAccAddress, position types.Position, positionInstance types.PerpetualFuturesPositionInstance) error {
	params := k.GetParams(ctx)

	netPosition := k.GetPositionSizeOfNetPositionOfMarket(ctx, position.Market)

	imaginaryFundingRate := netPosition.Mul(params.PerpetualFutures.ImaginaryFundingRateProportionalCoefficient)
	imaginaryFundingFee := sdk.NewDecFromInt(position.RemainingMargin.Amount).Mul(imaginaryFundingRate).RoundInt()
	commissionFee := sdk.NewDecFromInt(position.RemainingMargin.Amount).Mul(params.PerpetualFutures.CommissionRate).RoundInt()

	if imaginaryFundingRate.IsNegative() {
		if positionInstance.PositionType == types.PositionType_SHORT {
			position.RemainingMargin.Amount = position.RemainingMargin.Amount.Sub(imaginaryFundingFee)
		} else {
			position.RemainingMargin.Amount = position.RemainingMargin.Amount.Add(imaginaryFundingFee.Sub(commissionFee))
		}
	} else {
		if positionInstance.PositionType == types.PositionType_LONG {
			position.RemainingMargin.Amount = position.RemainingMargin.Amount.Sub(imaginaryFundingFee)
		} else {
			position.RemainingMargin.Amount = position.RemainingMargin.Amount.Add(imaginaryFundingFee.Sub(commissionFee))
		}
	}
	position.LastLeviedAt = ctx.BlockTime()

	rewardAmount := sdk.NewDecFromInt(commissionFee).Mul(params.PoolParams.ReportLevyPeriodRewardRate).RoundInt()
	reward := sdk.NewCoins(sdk.NewCoin(position.RemainingMargin.Denom, rewardAmount))
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, rewardRecipient.AccAddress(), reward)
	if err != nil {
		return err
	}

	k.SetPosition(ctx, position)

	ctx.EventManager().EmitTypedEvent(&types.EventPerpetualFuturesPositionLevied{
		RewardRecipient: rewardRecipient.AccAddress().String(),
		PositionId:      position.Id,
		RemainingMargin: position.RemainingMargin.String(),
		RewardAmount:    rewardAmount.String(),
	})

	return nil
}

func (k Keeper) GetPerpetualFuturesNetPositionOfMarket(ctx sdk.Context, market types.Market) types.PerpetualFuturesNetPositionOfMarket {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.DenomNetPositionPerpetualFuturesKeyPrefix(market.BaseDenom, market.QuoteDenom))
	if bz == nil {
		return types.PerpetualFuturesNetPositionOfMarket{}
	}

	netPositionOfMarket := types.PerpetualFuturesNetPositionOfMarket{}
	k.cdc.MustUnmarshal(bz, &netPositionOfMarket)
	return netPositionOfMarket
}

func (k Keeper) GetPositionSizeOfNetPositionOfMarket(ctx sdk.Context, market types.Market) sdk.Dec {
	position := k.GetPerpetualFuturesNetPositionOfMarket(ctx, market)
	if position.PositionSize.IsNil() {
		return sdk.ZeroDec()
	}
	return position.PositionSize
}

func (k Keeper) GetAllPerpetualFuturesNetPositionOfMarket(ctx sdk.Context) []types.PerpetualFuturesNetPositionOfMarket {
	store := ctx.KVStore(k.storeKey)

	perpetualFuturesNetPositionOfMarkets := []types.PerpetualFuturesNetPositionOfMarket{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixPerpetualFutures))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		netPositionOfMarket := types.PerpetualFuturesNetPositionOfMarket{}
		k.cdc.MustUnmarshal(it.Value(), &netPositionOfMarket)

		perpetualFuturesNetPositionOfMarkets = append(
			perpetualFuturesNetPositionOfMarkets,
			netPositionOfMarket,
		)
	}
	return perpetualFuturesNetPositionOfMarkets
}

func (k Keeper) SetPerpetualFuturesNetPositionOfMarket(ctx sdk.Context, NetPositionOfMarket types.PerpetualFuturesNetPositionOfMarket) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&NetPositionOfMarket)

	store.Set(types.DenomNetPositionPerpetualFuturesKeyPrefix(NetPositionOfMarket.Market.BaseDenom, NetPositionOfMarket.Market.QuoteDenom), bz)
}

func (k Keeper) AddPerpetualFuturesNetPositionOfMarket(ctx sdk.Context, market types.Market, rhs sdk.Dec) {
	lhs := k.GetPositionSizeOfNetPositionOfMarket(ctx, market)
	result := lhs.Add(rhs)

	perpetualFuturesNetPositionOfMarket := types.NewPerpetualFuturesNetPositionOfMarket(market, result)
	k.SetPerpetualFuturesNetPositionOfMarket(ctx, perpetualFuturesNetPositionOfMarket)
}

func (k Keeper) SubPerpetualFuturesNetPositionOfMarket(ctx sdk.Context, market types.Market, rhs sdk.Dec) {
	lhs := k.GetPositionSizeOfNetPositionOfMarket(ctx, market)
	result := lhs.Sub(rhs)

	perpetualFuturesNetPositionOfMarket := types.NewPerpetualFuturesNetPositionOfMarket(market, result)
	k.SetPerpetualFuturesNetPositionOfMarket(ctx, perpetualFuturesNetPositionOfMarket)
}