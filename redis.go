package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer func(rdb *redis.Client) {
		err := rdb.Close()
		if err != nil {
			panic(err)
		}
	}(rdb)

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if errors.Is(err, redis.Nil) {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Custom Play with Redis
	zsetKey := "waitinglist:football"
	res, err := rdb.ZAdd(ctx, zsetKey, redis.Z{
		Score:  float64(time.Now().UnixMicro()),
		Member: "21",
	}).Result()

	if err != nil {
		fmt.Println("err :", err)
	} else {
		fmt.Println("res : ", res)
	}

	// Second Play (ZRANGE)
	list, err := rdb.ZRangeWithScores(ctx, zsetKey, 0, -1).Result()
	if err != nil {
		fmt.Println("err :", err)
	}
	fmt.Println("list: ", list)
	for _, item := range list {
		fmt.Printf("Member: %v (Type: %T), Score: %f\n", item.Member, item.Member, item.Score)

		mStr, ok := item.Member.(string)
		if ok && mStr == "21" {
			remRes, err := rdb.ZRem(ctx, zsetKey, item.Member).Result()
			if err != nil {
				fmt.Println("ZRem err :", err)
			} else {
				fmt.Println("Successfully removed! remRes : ", remRes)
			}
		}
	}
	fmt.Println(list)

}
