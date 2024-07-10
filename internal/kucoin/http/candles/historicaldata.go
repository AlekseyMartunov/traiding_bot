package kucoincandles

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	entity "tradingbot/internal/kucoin/entity"
	kucoinerrors "tradingbot/internal/kucoin/errors"
)

func (cm *CandlesManager) GetHistoricalData(
	pair string,
	period entity.CandlePeriod,
	from, to time.Time) ([]*entity.Candle, error) {

	url := strings.Join([]string{cm.cfg.GetBaseEndpoint(), endpoint}, "")
	resp, err := cm.client.R().
		SetQueryParam("type", string(period)).
		SetQueryParam("symbol", pair).
		SetQueryParam("startAt", fmt.Sprint(from.Unix())).
		SetQueryParam("endAt", fmt.Sprint(to.Unix())).
		Get(url)

	if err != nil {
		cm.log.Error(err.Error())
		return nil, err
	}

	if resp.StatusCode() != 200 {
		cm.log.Error(fmt.Sprintf("body: %s, code: %d", resp.String(), resp.StatusCode()))
		return nil, kucoinerrors.ErrStatusCodeIsNot200
	}

	j := candlesJSON{}

	err = json.Unmarshal(resp.Body(), &j)
	if err != nil {
		fmt.Println(err)
	}

	return j.toBaseEntity()
}
