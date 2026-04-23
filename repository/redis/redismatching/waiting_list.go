package redismatching

import (
	"context"
	"fmt"
	"github.com/mobin-alz/gameapp/entity"
	"github.com/mobin-alz/gameapp/pkg/richerror"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	WaitingListPrefix = "waitinglist"
)

func (d DB) AddToWaitingList(userID uint, category entity.Category) error {
	const op = "redismatching.AddToWaitingList"

	zSetKey := fmt.Sprintf("%s:%s", WaitingListPrefix, category)
	_, err := d.adapter.Client().ZAdd(context.Background(), zSetKey, redis.Z{
		Score:  float64(time.Now().UnixMicro()),
		Member: fmt.Sprintf("%d", userID),
	}).Result()

	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}
