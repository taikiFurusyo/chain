package types_test

import (
	"fmt"
	"testing"
	time "time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

// make position.NeedLiquidationPerpetualFutures test
func TestPosition_NeedLiquidationPerpetualFutures(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("ununifi1a8jcsmla6heu99ldtazc27dna4qcd4jygsthx6")
	testCases := []struct {
		name          string
		position      types.PerpetualFuturesPosition
		minMarginRate sdk.Dec
		closedRate    []sdk.Dec //first is base rate, second is quote rate
		exp           bool
	}{
		{
			name: "not_change_rate_is_not_need_liquidation",
			position: types.PerpetualFuturesPosition{
				Id:      "0",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedAt:        time.Now().UTC(),
				OpenedHeight:    1,
				OpenedBaseRate:  sdk.MustNewDecFromStr("100"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("100"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("10"),
					Leverage:     5,
				},
			},
			minMarginRate: sdk.MustNewDecFromStr("0.5"),
			closedRate: []sdk.Dec{
				sdk.MustNewDecFromStr("100"),
				sdk.MustNewDecFromStr("100"),
			},
			exp: false,
		},
		{
			name: "down_rate_is_need_liquidation",
			position: types.PerpetualFuturesPosition{
				Id:      "0",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedAt:        time.Now().UTC(),
				OpenedHeight:    1,
				OpenedBaseRate:  sdk.MustNewDecFromStr("100"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("100"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     1,
				},
			},
			minMarginRate: sdk.MustNewDecFromStr("0.5"),
			closedRate: []sdk.Dec{
				sdk.MustNewDecFromStr("1"),
				sdk.MustNewDecFromStr("100"),
			},
			exp: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.position.NeedLiquidation(tc.minMarginRate, tc.closedRate[0], tc.closedRate[1])
			if tc.exp != result {
				t.Error(tc, "expected %v, got %v", tc.exp, result)
			}
		})
	}
}

type CurrencyRate struct {
	name string
	rate sdk.Dec
}

func TestPosition_MarginMaintenanceRate(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("ununifi1a8jcsmla6heu99ldtazc27dna4qcd4jygsthx6")

	testCases := []struct {
		name     string
		position types.PerpetualFuturesPosition
		Rate     []CurrencyRate
		exp      sdk.Dec
	}{
		{
			name: "long position not change rate",
			position: types.PerpetualFuturesPosition{
				Id:      "0",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "ubtc",
				},
				OpenedAt:     time.Now().UTC(),
				OpenedHeight: 1,
				// atom/usd = 0.4
				OpenedBaseRate:  sdk.MustNewDecFromStr("400"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("400"),
				// In the case of Long, BaseDenom is RemainingMargin.
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("5"),
					Leverage:     5,
				},
			},
			// not change rate
			Rate: []CurrencyRate{
				{
					name: "uatom/usd",
					rate: sdk.MustNewDecFromStr("400"),
				},
				{
					name: "ubtc/usd",
					rate: sdk.MustNewDecFromStr("400"),
				},
			},
			exp: sdk.MustNewDecFromStr("1"),
		},
		{
			name: "long position down 10%",
			position: types.PerpetualFuturesPosition{
				Id:      "1",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "ubtc",
				},
				OpenedAt:        time.Now().UTC(),
				OpenedHeight:    1,
				OpenedBaseRate:  sdk.MustNewDecFromStr("100"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("100"),
				// In the case of Long, BaseDenom is RemainingMargin.
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(100)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     1,
				},
			},
			// down 10%
			Rate: []CurrencyRate{
				{
					name: "uatom/usd",
					rate: sdk.MustNewDecFromStr("90"),
				},
				{
					name: "ubtc/usd",
					rate: sdk.MustNewDecFromStr("100"),
				},
			},
			exp: sdk.MustNewDecFromStr("99.888888888888888889"),
		},
		{
			name: "long position up 10%",
			position: types.PerpetualFuturesPosition{
				Id:      "3",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "ubtc",
				},
				OpenedAt:        time.Now().UTC(),
				OpenedHeight:    1,
				OpenedBaseRate:  sdk.MustNewDecFromStr("100"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("100"),
				// In the case of Long, BaseDenom is RemainingMargin.
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(100)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("5"),
					Leverage:     5,
				},
			},
			Rate: []CurrencyRate{
				{
					name: "uatom/usd",
					rate: sdk.MustNewDecFromStr("110"),
				},
				{
					// up 10%
					name: "ubtc/usd",
					rate: sdk.MustNewDecFromStr("100"),
				},
			},
			exp: sdk.MustNewDecFromStr("102.272727272727272727"),
		},
		{
			name: "short position not change rate",
			position: types.PerpetualFuturesPosition{
				Id:      "0",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "ubtc",
				},
				OpenedAt:        time.Now().UTC(),
				OpenedHeight:    1,
				OpenedBaseRate:  sdk.MustNewDecFromStr("400"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("400"),
				// In the case of Long, BaseDenom is RemainingMargin.
				RemainingMargin: sdk.NewCoin("ubtc", sdk.NewInt(1)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("5"),
					Leverage:     5,
				},
			},
			// not change rate
			Rate: []CurrencyRate{
				{
					name: "uatom/usd",
					rate: sdk.MustNewDecFromStr("400"),
				},
				{
					name: "ubtc/usd",
					rate: sdk.MustNewDecFromStr("400"),
				},
			},
			exp: sdk.MustNewDecFromStr("1"),
		},
		{
			name: "short position down 10%",
			position: types.PerpetualFuturesPosition{
				Id:      "1",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "ubtc",
				},
				OpenedAt:        time.Now().UTC(),
				OpenedHeight:    1,
				OpenedBaseRate:  sdk.MustNewDecFromStr("100"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("100"),
				// In the case of Long, BaseDenom is RemainingMargin.
				RemainingMargin: sdk.NewCoin("ubtc", sdk.NewInt(100)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("5"),
					Leverage:     5,
				},
			},
			// down 10%
			Rate: []CurrencyRate{
				{
					name: "uatom/usd",
					rate: sdk.MustNewDecFromStr("90"),
				},
				{
					name: "ubtc/usd",
					rate: sdk.MustNewDecFromStr("100"),
				},
			},
			exp: sdk.MustNewDecFromStr("113.888888888888888889"),
		},
		{
			name: "short position up 10%",
			position: types.PerpetualFuturesPosition{
				Id:      "3",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "ubtc",
				},
				OpenedAt:        time.Now().UTC(),
				OpenedHeight:    1,
				OpenedBaseRate:  sdk.MustNewDecFromStr("100"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("100"),
				// In the case of Long, BaseDenom is RemainingMargin.
				RemainingMargin: sdk.NewCoin("ubtc", sdk.NewInt(100)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("5"),
					Leverage:     5,
				},
			},
			Rate: []CurrencyRate{
				{
					name: "uatom/usd",
					rate: sdk.MustNewDecFromStr("110"),
				},
				{
					// up 10%
					name: "ubtc/usd",
					rate: sdk.MustNewDecFromStr("100"),
				},
			},
			exp: sdk.MustNewDecFromStr("88.636363636363636364"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.position.MarginMaintenanceRate(tc.Rate[0].rate, tc.Rate[1].rate)
			if !tc.exp.Equal(result) {
				t.Error(tc, "expected %v, got %v", tc.exp, result)
			}
		})
	}
}

// make PerpetualFuturesPosition.CalcProfitAndLoss test
func TestPerpetualFuturesPosition_CalcProfitAndLoss(t *testing.T) {
	testCases := []struct {
		name        string
		position    types.PerpetualFuturesPosition
		closedRates []sdk.Dec
		exp         math.Int
	}{
		{
			name: "Long position profit",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("100"),
					Leverage:     5,
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("11.1"),
				sdk.MustNewDecFromStr("11.1"),
			},
			exp: sdk.NewInt(110000000),
		},
		{
			name: "Long position loss",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("100"),
					Leverage:     5,
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("9.1"),
				sdk.MustNewDecFromStr("9.1"),
			},
			exp: sdk.NewInt(-90000000),
		},
		{
			name: "Short position profit",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("100"),
					Leverage:     5,
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("9.1"),
				sdk.MustNewDecFromStr("9.1"),
			},
			exp: sdk.NewInt(90000000),
		},
		{
			name: "Short position loss",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("100"),
					Leverage:     5,
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("12.1"),
				sdk.MustNewDecFromStr("12.1"),
			},
			exp: sdk.NewInt(-210000000),
		},
		{
			name: "Profit Long position in base denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     10,
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("20"),
				sdk.MustNewDecFromStr("20"),
			},
			exp: sdk.NewInt(500000),
		},
		{
			name: "Profit Short position in base denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("20"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     20,
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("10"),
				sdk.MustNewDecFromStr("10"),
			},
			exp: sdk.NewInt(1000000),
		},
		{
			name: "Loss Long position in base denom margin ",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("20"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     20,
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("10"),
				sdk.MustNewDecFromStr("10"),
			},
			exp: sdk.NewInt(-1000000),
		},
		{
			name: "Loss Short position in base denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     10,
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("20"),
				sdk.MustNewDecFromStr("20"),
			},
			exp: sdk.NewInt(-500000),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resultDec := tc.position.ProfitAndLossInMetrics(tc.closedRates[0], tc.closedRates[1])
			result := types.MicroToNormalDenom(resultDec)
			fmt.Println(result)
			if !tc.exp.Equal(result) {
				t.Error(tc, "expected %v, got %v", tc.exp, result)
			}
		})
	}
}

// CalcReturningAmountAtClose test
func TestCalcReturningAmountAtClose(t *testing.T) {
	testCases := []struct {
		name           string
		position       types.PerpetualFuturesPosition
		closedPairRate sdk.Dec
		exp            math.Int
	}{
		{
			name: "Profit Long position in quote denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     10,
				},
			},
			closedPairRate: sdk.MustNewDecFromStr("20"),
			exp:            sdk.NewInt(11000000),
		},
		{
			name: "Profit Short position in quote denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("20"),
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     20,
				},
			},
			closedPairRate: sdk.MustNewDecFromStr("10"),
			exp:            sdk.NewInt(11000000),
		},
		{
			name: "Loss Long position in quote denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     10,
				},
			},
			closedPairRate: sdk.MustNewDecFromStr("9"),
			exp:            sdk.NewInt(0),
		},
		{
			name: "Loss Short position in quote denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     1,
				},
			},
			closedPairRate: sdk.MustNewDecFromStr("10.5"),
			exp:            sdk.NewInt(500000),
		},
		{
			name: "Profit Long position in base denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     1,
				},
			},
			closedPairRate: sdk.MustNewDecFromStr("20"),
			exp:            sdk.NewInt(1500000),
		},
		{
			name: "Profit Short position in base denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("20"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     1,
				},
			},
			closedPairRate: sdk.MustNewDecFromStr("10"),
			exp:            sdk.NewInt(2000000),
		},
		{
			name: "Loss Long position in base denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     10,
				},
			},
			closedPairRate: sdk.MustNewDecFromStr("9"),
			exp:            sdk.NewInt(888889),
		},
		{
			name: "Loss Short position in base denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     1,
				},
			},
			closedPairRate: sdk.MustNewDecFromStr("11"),
			exp:            sdk.NewInt(909091),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			returningAmount, lossToLP := tc.position.CalcReturningAmountAtClose(tc.closedPairRate, tc.closedPairRate)
			fmt.Println(returningAmount, lossToLP)
			if !tc.exp.Equal(returningAmount) {
				t.Error(tc, "expected %v, got %v", tc.exp, returningAmount)
			}
		})
	}
}