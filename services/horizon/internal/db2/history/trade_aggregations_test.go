package history

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimitTimeRangeNoChange(t *testing.T) {
	q := TradeAggregationsQ{
		resolution: 60000,
		offset:     0,
		startTime:  0,
		endTime:    0,
	}

	res, err := q.LimitTimeRange()
	assert.NoError(t, err)
	assert.Equal(t, int64(0), int64(res.startTime), "startTime should be equal")
	assert.Equal(t, int64(0), int64(res.endTime), "endTime should be equal")

	q = TradeAggregationsQ{
		resolution: 60000,
		offset:     0,
		startTime:  1512764500000,
		endTime:    1512775500000,
	}

	res, err = q.LimitTimeRange()
	assert.NoError(t, err)
	assert.Equal(t, int64(1512764500000), int64(res.startTime), "startTime should be equal")
	assert.Equal(t, int64(1512775500000), int64(res.endTime), "endTime should be equal")

	q = TradeAggregationsQ{
		resolution: 3600000,
		offset:     0,
		startTime:  1512689100000,
		endTime:    1512775500000,
	}

	res, err = q.LimitTimeRange()
	assert.NoError(t, err)
	assert.Equal(t, int64(1512689100000), int64(res.startTime), "startTime should be equal")
	assert.Equal(t, int64(1512775500000), int64(res.endTime), "endTime should be equal")
}

func TestLimitTimeRangeSetStartTime(t *testing.T) {
	q := TradeAggregationsQ{
		resolution: 60000,
		offset:     0,
		startTime:  1512689100000,
		endTime:    1512775500000,
	}

	res, err := q.LimitTimeRange()
	assert.NoError(t, err)
	assert.Equal(t, int64(1512763500000), int64(res.startTime), "startTime should be equal")
	assert.Equal(t, int64(1512775500000), int64(res.endTime), "endTime should be equal")

	q = TradeAggregationsQ{
		resolution: 3600000,
		offset:     0,
		startTime:  0,
		endTime:    1512775500000,
	}

	res, err = q.LimitTimeRange()
	assert.NoError(t, err)
	assert.Equal(t, int64(1512055500000), int64(res.startTime), "startTime should be equal")
	assert.Equal(t, int64(1512775500000), int64(res.endTime), "endTime should be equal")
}

func TestLimitTimeRangeWithOffset(t *testing.T) {
	q := TradeAggregationsQ{
		resolution: 3600000,
		offset:     3600000,
		startTime:  0,
		endTime:    1512775500000,
	}

	res, err := q.LimitTimeRange()
	assert.NoError(t, err)
	assert.Equal(t, int64(1512051900000), int64(res.startTime), "startTime should be equal")
	assert.Equal(t, int64(1512775500000), int64(res.endTime), "endTime should be equal")
}
