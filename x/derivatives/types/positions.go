package types

import (
	"fmt"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

type PositionInstance interface {
	proto.Message
}

type Positions []Position

func UnpackPositionInstance(positionAny types.Any) (PositionInstance, error) {
	position := UnpackPerpetualFuturesPositionInstance(positionAny)
	if position != nil {
		return position, nil
	}

	position = UnpackPerpetualOptionsPosition(positionAny)
	if position != nil {
		return position, nil
	}

	return nil, fmt.Errorf("this Any doesn't have PositionInstance value")
}

func MustUnpackPositionInstance(positionAny types.Any) PositionInstance {
	position, err := UnpackPositionInstance(positionAny)
	if err != nil {
		panic(err)
	}
	return position
}

func (m Position) NeedLiquidation(MarginMaintenanceRate, currentBaseUsdRate, currentQuoteUsdRate, currentMarginUsdRate sdk.Dec) bool {
	ins, err := UnpackPositionInstance(m.PositionInstance)
	if err != nil {
		return false
	}

	switch positionInstance := ins.(type) {
	case *PerpetualFuturesPositionInstance:
		perpetualFuturesPosition := NewPerpetualFuturesPosition(m, *positionInstance)
		return perpetualFuturesPosition.NeedLiquidation(MarginMaintenanceRate, currentBaseUsdRate, currentQuoteUsdRate, currentMarginUsdRate)
		break
	case *PerpetualOptionsPositionInstance:
		panic("not implemented")
		break
	default:
		panic("not implemented")
	}
	return false
}

func NewPerpetualFuturesPosition(position Position, ins PerpetualFuturesPositionInstance) PerpetualFuturesPosition {
	return PerpetualFuturesPosition{
		Id:               position.Id,
		Market:           position.Market,
		Address:          position.Address,
		OpenedAt:         position.OpenedAt,
		OpenedBaseRate:   position.OpenedBaseRate,
		OpenedQuoteRate:  position.OpenedQuoteRate,
		OpenedHeight:     position.OpenedHeight,
		RemainingMargin:  position.RemainingMargin,
		LastLeviedAt:     position.LastLeviedAt,
		PositionInstance: ins,
	}
}

func (m PerpetualFuturesPosition) NeedLiquidation(minMarginMaintenanceRate, currentBaseUsdRate, currentQuoteUsdRate, currentMarginUsdRate sdk.Dec) bool {
	marginMaintenanceRate := m.MarginMaintenanceRate(currentBaseUsdRate, currentQuoteUsdRate, currentMarginUsdRate)
	if marginMaintenanceRate.LT(minMarginMaintenanceRate) {
		return true
	} else {
		return false
	}
}

// todo make test
func (m PerpetualFuturesPosition) EffectiveMargin(currentBaseUsdRate, currentQuoteUsdRate, currentMarginUsdRate sdk.Dec) sdk.Dec {
	effectiveMargin := sdk.NewDecFromInt(m.RemainingMargin.Amount).Mul(currentBaseUsdRate.Quo(currentMarginUsdRate))

	revenue := m.CalcProfit(currentBaseUsdRate.Quo(currentQuoteUsdRate))
	if revenue.RevenueType == RevenueType_PROFIT {
		effectiveMargin = effectiveMargin.Add(sdk.NewDecFromInt(revenue.Amount.Amount))
	} else {
		effectiveMargin = effectiveMargin.Sub(sdk.NewDecFromInt(revenue.Amount.Amount))
	}
	return effectiveMargin
}

func (m PerpetualFuturesPosition) MarginMaintenanceRate(currentBaseUsdRate, currentQuoteUsdRate, currentMarginUsdRate sdk.Dec) sdk.Dec {
	marginRequirement := m.PositionInstance.MarginRequirement(currentBaseUsdRate.Quo(currentMarginUsdRate))
	effectiveMargin := m.EffectiveMargin(currentBaseUsdRate, currentQuoteUsdRate, currentMarginUsdRate)

	marginMaintenanceRate := effectiveMargin.Quo(marginRequirement)
	return marginMaintenanceRate
}

func (m PerpetualFuturesPosition) CalcProfitAndLoss(closedRate sdk.Dec) math.Int {
	sub := closedRate.Sub(m.OpenedBaseRate)
	if m.PositionInstance.PositionType == PositionType_SHORT {
		sub = sub.Neg()
	}

	resultDec := sub.Mul(m.PositionInstance.Size_)

	// profit must be calculated in remaining margin denom
	if m.RemainingMargin.Denom != m.Market.QuoteDenom {
		resultDec = resultDec.Quo(closedRate)
	}

	// make it micro unit by multiplying 1000000
	// this means it assumes the price difference is calculated in normal unit, not micro unit.
	// e.g. In ubtc/uusdc market, the market price of ubtc is actually in BTC unit.
	// And, the position size follows the market price unit.
	actualResultAmount := resultDec.Mul(sdk.MustNewDecFromStr("1000000")).RoundInt()

	return actualResultAmount
}

func (m PerpetualFuturesPosition) CalcReturningAmountAtClose(closedRate sdk.Dec) math.Int {
	principal := m.RemainingMargin.Amount
	pnlAmount := m.CalcProfitAndLoss(closedRate)
	returningAmount := principal.Add(pnlAmount)

	return returningAmount
}

// todo make test
func (m Positions) EvaluatePositions(posType PositionType, getCurrentPriceF func(denom string) (sdk.Dec, error)) sdk.Dec {
	usdMap := map[string]sdk.Dec{}
	result := sdk.ZeroDec()
	for _, position := range m {
		ins, err := UnpackPositionInstance(position.PositionInstance)
		if err != nil {
			panic(err)
		}

		if _, ok := usdMap[position.Market.BaseDenom]; !ok {
			price, err := getCurrentPriceF(position.Market.BaseDenom)
			if err != nil {
				panic(err)
			}
			usdMap[position.Market.BaseDenom] = price
		}

		switch positionInstance := ins.(type) {
		case *PerpetualFuturesPositionInstance:
			perpetualFuturesPosition := NewPerpetualFuturesPosition(position, *positionInstance)
			if perpetualFuturesPosition.PositionInstance.PositionType != posType {
				continue
			}
			result = result.Add(perpetualFuturesPosition.EvaluatePosition(usdMap[position.Market.BaseDenom]))
			break
		case *PerpetualOptionsPositionInstance:
			panic("not implemented")
		default:
			continue
		}
	}
	return result
}

func (m Positions) EvaluateLongPositions(getCurrentPriceF func(denom string) (sdk.Dec, error)) sdk.Dec {
	return m.EvaluatePositions(PositionType_LONG, getCurrentPriceF)
}

func (m Positions) EvaluateShortPositions(getCurrentPriceF func(denom string) (sdk.Dec, error)) sdk.Dec {
	return m.EvaluatePositions(PositionType_SHORT, getCurrentPriceF)
}
