// InMetrics represents the profit/loss amount in the metrics asset of the market.
// In the most cases, it means it's in "usd".
// And IMPORTANTLY, it means it's not calcualted in micro case.

package types

import (
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/cosmos/gogoproto/proto"
)

type PositionInstance interface {
	proto.Message
}

type Positions []Position

func (m Position) IsValid(params Params, AvailableAssetInPoolByDenom sdk.Coin) error {
	if !m.IsValidMarginAsset() {
		return ErrMarginAssetNotValid
	}

	// check the least requirement for the margin
	if !m.RemainingMargin.Amount.IsPositive() {
		return ErrNegativeMargin
	}

	pfPosition, err := NewPerpetualFuturesPositionFromPosition(m)
	if err != nil {
		return err
	}

	if !pfPosition.PositionInstance.IsValidLeverage(params.PerpetualFutures.MaxLeverage) {
		return ErrInvalidLeverage
	}

	if !pfPosition.IsValidPositionSize(params.PoolParams.QuoteTicker) {
		return ErrInvalidPositionSize
	}

	if AvailableAssetInPoolByDenom.Amount.LT(pfPosition.PositionInstance.SizeInDenomExponent(OneMillionInt)) {
		return ErrInsufficientPoolFund
	}

	return nil
}

// Margin asset must be one of the market assets.
func (m Position) IsValidMarginAsset() bool {
	return (m.Market.BaseDenom == m.RemainingMargin.Denom || m.Market.QuoteDenom == m.RemainingMargin.Denom)
}

func (m PerpetualFuturesPosition) IsValidPositionSize(quoteTicker string) bool {
	// check position size validity
	baseMetricsRate := NewMetricsRateType(quoteTicker, m.Market.BaseDenom, m.OpenedBaseRate)
	quoteMetricsRate := NewMetricsRateType(quoteTicker, m.Market.QuoteDenom, m.OpenedQuoteRate)
	marginMaintenanceRate := m.MarginMaintenanceRate(baseMetricsRate, quoteMetricsRate)

	return !marginMaintenanceRate.LT(sdk.OneDec())
}

func (m PerpetualFuturesPositionInstance) IsValidLeverage(maxLeverage uint32) bool {
	return m.Leverage > 0 && m.Leverage <= maxLeverage
}

func UnpackPositionInstance(positionAny types.Any) (PositionInstance, error) {
	position := UnpackPerpetualFuturesPositionInstance(positionAny)
	if position != nil {
		return position, nil
	}

	position = UnpackPerpetualOptionsPosition(positionAny)
	if position != nil {
		return position, nil
	}

	return nil, ErrInvalidPositionInstance
}

func MustUnpackPositionInstance(positionAny types.Any) (PositionInstance, error) {
	position, err := UnpackPositionInstance(positionAny)
	if err != nil {
		return nil, err
	}
	return position, nil
}

func (m Position) NeedLiquidation(MarginMaintenanceRate sdk.Dec, currentBaseMetricsRate, currentQuoteMetricsRate MetricsRateType) bool {
	ins, err := UnpackPositionInstance(m.PositionInstance)
	if err != nil {
		return false
	}

	switch positionInstance := ins.(type) {
	case *PerpetualFuturesPositionInstance:
		perpetualFuturesPosition := NewPerpetualFuturesPosition(m, *positionInstance)
		return perpetualFuturesPosition.NeedLiquidation(MarginMaintenanceRate, currentBaseMetricsRate, currentQuoteMetricsRate)
	case *PerpetualOptionsPositionInstance:
		// todo implement
		return false
	default:
		return false
	}
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

func NewPerpetualFuturesPositionFromPosition(position Position) (PerpetualFuturesPosition, error) {
	ins, err := UnpackPositionInstance(position.PositionInstance)
	if err != nil {
		return PerpetualFuturesPosition{}, err
	}
	switch positionInstance := ins.(type) {
	case *PerpetualFuturesPositionInstance:
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
			PositionInstance: *positionInstance,
		}, nil
	default:
		return PerpetualFuturesPosition{}, ErrInvalidPositionInstance
	}
}

func (m PerpetualFuturesPosition) NeedLiquidation(minMarginMaintenanceRate sdk.Dec, currentBaseMetricsRate, currentQuoteMetricsRate MetricsRateType) bool {
	marginMaintenanceRate := m.MarginMaintenanceRate(currentBaseMetricsRate, currentQuoteMetricsRate)
	if marginMaintenanceRate.LTE(minMarginMaintenanceRate) {
		return true
	} else {
		return false
	}
}

func (m PerpetualFuturesPosition) OpenedPairRate() sdk.Dec {
	return m.OpenedBaseRate.Quo(m.OpenedQuoteRate)
}

// todo make test
func (m PerpetualFuturesPosition) EvaluatePosition(currentBaseMetricsRate MetricsRateType) sdk.Dec {
	// FIXME: Don't use OneMillionInt directly. issue #476
	return currentBaseMetricsRate.Amount.Amount.Mul(sdk.NewDecFromInt(m.PositionInstance.SizeInDenomExponent(OneMillionInt)))
}

func NormalToMicroInt(amount sdk.Dec) sdk.Int {
	return amount.Mul(sdk.MustNewDecFromStr(OneMillionString)).TruncateInt()
}

func NormalToMicroDec(amount sdk.Dec) sdk.Dec {
	return amount.Mul(sdk.MustNewDecFromStr(OneMillionString))
}

// todo make test
func (m Positions) EvaluatePositions(posType PositionType, quoteTicker string, getCurrentPriceF func(denom string) (sdk.Dec, error)) (sdk.Dec, error) {
	usdMap := map[string]sdk.Dec{}
	result := sdk.ZeroDec()
	for _, position := range m {
		ins, err := UnpackPositionInstance(position.PositionInstance)
		if err != nil {
			return sdk.ZeroDec(), err
		}

		if _, ok := usdMap[position.Market.BaseDenom]; !ok {
			price, err := getCurrentPriceF(position.Market.BaseDenom)
			if err != nil {
				return sdk.ZeroDec(), err
			}
			usdMap[position.Market.BaseDenom] = price
		}

		switch positionInstance := ins.(type) {
		case *PerpetualFuturesPositionInstance:
			perpetualFuturesPosition := NewPerpetualFuturesPosition(position, *positionInstance)
			if perpetualFuturesPosition.PositionInstance.PositionType != posType {
				continue
			}

			metricsRate := NewMetricsRateType(quoteTicker, position.Market.BaseDenom, usdMap[position.Market.BaseDenom])

			result = result.Add(perpetualFuturesPosition.EvaluatePosition(metricsRate))
		case *PerpetualOptionsPositionInstance:
			return sdk.ZeroDec(), ErrNotImplemented
		default:
			continue
		}
	}
	return result, nil
}

func (m Positions) EvaluateLongPositions(quoteTicker string, getCurrentPriceF func(denom string) (sdk.Dec, error)) (sdk.Dec, error) {
	value, err := m.EvaluatePositions(PositionType_LONG, quoteTicker, getCurrentPriceF)
	if err != nil {
		return sdk.ZeroDec(), err
	}
	return value, nil
}

func (m Positions) EvaluateShortPositions(quoteTicker string, getCurrentPriceF func(denom string) (sdk.Dec, error)) (sdk.Dec, error) {
	value, err := m.EvaluatePositions(PositionType_SHORT, quoteTicker, getCurrentPriceF)
	if err != nil {
		return sdk.ZeroDec(), err
	}
	return value, nil
}

func (positionInstance PerpetualFuturesPositionInstance) MarginRequirement(currencyRate sdk.Dec) sdk.Int {
	// FIXME: Don't use OneMillionInt directly. issue #476
	return sdk.NewDecFromInt(positionInstance.SizeInDenomExponent(OneMillionInt)).Mul(currencyRate).Quo(sdk.NewDec(int64(positionInstance.Leverage))).TruncateInt()
}

func (m PerpetualFuturesPosition) RequiredMarginInQuote(baseQuoteRate sdk.Dec) sdk.Int {
	// 必要証拠金(quote単位) = 現在のbase/quoteレート * ポジションサイズ(base単位) ÷ レバレッジ
	return m.PositionInstance.MarginRequirement(baseQuoteRate)
}

func (m PerpetualFuturesPosition) RequiredMarginInBase() sdk.Int {
	// 必要証拠金(base単位) = ポジションサイズ(base単位) ÷ レバレッジ // レートでの変動なし
	return m.PositionInstance.MarginRequirement(sdk.MustNewDecFromStr("1"))
}

func (m PerpetualFuturesPosition) RequiredMarginInMetrics(baseMetricsRate, quoteMetricsRate MetricsRateType) sdk.Dec {
	// 必要証拠金(USD単位) = 必要証拠金(quote単位) * 現在のquote/USDレート
	//                    = 必要証拠金(base単位) * 現在のbase/USDレート
	if m.RemainingMargin.Denom == m.Market.QuoteDenom {
		baseQuoteRate := baseMetricsRate.Amount.Amount.Quo(quoteMetricsRate.Amount.Amount)
		return sdk.NewDecFromInt(m.RequiredMarginInQuote(baseQuoteRate)).Mul(quoteMetricsRate.Amount.Amount)
	} else {
		return sdk.NewDecFromInt(m.RequiredMarginInBase()).Mul(baseMetricsRate.Amount.Amount)
	}
}

// CalcReturningAmountAtClose calculates the amount of the principal and the profit/loss at the close of the position.
func (m PerpetualFuturesPosition) CalcReturningAmountAtClose(baseMetricsRate, quoteMetricsRate MetricsRateType, tradingFee sdk.Int) (returningAmount math.Int, lossToLP math.Int) {
	principal := m.RemainingMargin.Amount
	// pnlAmountInMetrics represents the profit/loss amount in the metrics asset of the market.
	// In the most cases, it means it's in "usd".
	// AND, MORE IMPORTANTLY,
	// it's not calculated on a micro level. So, it has to be modified to micro level by multiplying
	// one million to represent the returning amount.
	pnlAmount := m.ProfitAndLoss(baseMetricsRate, quoteMetricsRate)

	returningAmount = principal.Add(pnlAmount)

	// If loss is over the margin, it means liquidity provider takes the loss.
	if returningAmount.IsNegative() {
		lossToLP = returningAmount
		returningAmount = sdk.ZeroInt()
	} else {
		lossToLP = sdk.ZeroInt()
	}

	// Subtract the trading fee.
	if returningAmount.LT(tradingFee) {
		// Return 0 returning amount and the trading fee subtracted by the returning amount as LossToLP
		return sdk.ZeroInt(), tradingFee.Sub(returningAmount)
	}

	returningAmount = returningAmount.Sub(tradingFee)

	return returningAmount, lossToLP
}

// ProfitAndLoss returns the profit/loss amount in the margin denom
func (m PerpetualFuturesPosition) ProfitAndLoss(baseMetricsRate, quoteMetricsRate MetricsRateType) sdk.Int {
	pnlAmountInMetrics := m.ProfitAndLossInMetrics(baseMetricsRate, quoteMetricsRate)

	// Make it be calculated in the corresponding asset as the principal.
	var pnlAmount sdk.Dec
	if m.RemainingMargin.Denom == m.Market.BaseDenom {
		pnlAmount = pnlAmountInMetrics.Quo(baseMetricsRate.Amount.Amount)
	} else {
		pnlAmount = pnlAmountInMetrics.Quo(quoteMetricsRate.Amount.Amount)
	}

	return pnlAmount.TruncateInt()
}

func (m PerpetualFuturesPosition) ProfitAndLossInQuote(baseMetricsRate, quoteMetricsRate MetricsRateType) sdk.Dec {
	// 損益(quote単位) = (longなら*1,shortなら*-1) * (現在のbase/quoteレート - ポジション開設時base/quoteレート) * ポジションサイズ(base単位)
	baseQuoteRate := baseMetricsRate.Amount.Amount.Quo(quoteMetricsRate.Amount.Amount)
	// FIXME: Don't use OneMillionInt derectly. issue #476
	profitAndLoss := baseQuoteRate.Sub(m.OpenedPairRate()).Mul(sdk.NewDecFromInt(m.PositionInstance.SizeInDenomExponent(OneMillionInt)))
	if m.PositionInstance.PositionType == PositionType_LONG {
		return profitAndLoss
	} else {
		return profitAndLoss.Neg()
	}
}

func (m PerpetualFuturesPosition) ProfitAndLossInMetrics(baseMetricsRate, quoteMetricsRate MetricsRateType) sdk.Dec {
	// 損益(USD単位) = 損益(quote単位) * 現在のquote/USDレート
	return m.ProfitAndLossInQuote(baseMetricsRate, quoteMetricsRate).Mul(quoteMetricsRate.Amount.Amount)
}

// position size takes 0 decimal although price takes 6 decimal (micro unit)
func (m PerpetualFuturesPosition) MarginMaintenanceRate(baseMetricsRate, quoteMetricsRate MetricsRateType) sdk.Dec {
	// 証拠金維持率 = 有効証拠金(USD単位) ÷ 必要証拠金(USD単位)
	return m.EffectiveMarginInMetrics(baseMetricsRate, quoteMetricsRate).Quo(m.RequiredMarginInMetrics(baseMetricsRate, quoteMetricsRate))
}

func (m PerpetualFuturesPosition) RemainingMarginInMetrics(baseMetricsRate, quoteMetricsRate MetricsRateType) sdk.Dec {
	// 残存証拠金(USD単位) = 残存証拠金(base単位) * 現在のbase/USDレート
	//                    = 残存証拠金(quote単位) * 現在のquote/USDレート
	remainingMarginAmountInDec := sdk.NewDecFromInt(m.RemainingMargin.Amount)
	if m.RemainingMargin.Denom == m.Market.BaseDenom {
		return remainingMarginAmountInDec.Mul(baseMetricsRate.Amount.Amount)
	} else if m.RemainingMargin.Denom == m.Market.QuoteDenom {
		return remainingMarginAmountInDec.Mul(quoteMetricsRate.Amount.Amount)
	} else {
		// not supported denom
		return sdk.ZeroDec()
	}
}

func (m PerpetualFuturesPosition) EffectiveMarginInMetrics(baseMetricsRate, quoteMetricsRate MetricsRateType) sdk.Dec {
	// 有効証拠金(USD単位) = 残存証拠金(USD単位) + 損益(USD単位)
	return m.RemainingMarginInMetrics(baseMetricsRate, quoteMetricsRate).Add(m.ProfitAndLossInMetrics(baseMetricsRate, quoteMetricsRate))
}

func NewMetricsRateType(unit string, denom string, amount sdk.Dec) MetricsRateType {
	return MetricsRateType{
		MetricsUnit: unit,
		Amount:      sdk.NewDecCoinFromDec(denom, amount),
	}
}

type MetricsRateType struct {
	MetricsUnit string
	Amount      sdk.DecCoin
}
