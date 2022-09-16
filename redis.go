// package main

// import (
// 	"context"
// 	"fmt"

// 	"github.com/go-redis/redis/v8"
// 	"github.com/google/uuid"
// )

// func main() {

// 	var (
// 		uuID = uuid.NewString()
// 		ctx  = context.Background()
// 	)

// 	opt, err := redis.ParseURL("redis://:@localhost:6379/0")
// 	rdb := redis.NewClient(opt)

// 	for i := 0; i < 20; i++ {
// 		// command to add fake events
// 		rdb.RPush(ctx, "lms_events", []byte(uuID))
// 	}

// 	fmt.Println("Added 20 values\n")

// 	switch {
// 	case err == redis.Nil:
// 		fmt.Println("key does not exist")
// 	case err != nil:
// 		fmt.Println(err)
// 	}

// 	// command to range on the key
// 	lr, err := rdb.LRange(ctx, "lms_events", 0, 100000).Result()
// 	fmt.Println("range:\n", lr)

// 	// command to get a specific value
// 	sv, _ := rdb.LIndex(ctx, "lms_events", 3).Result()
// 	fmt.Println("\nvalue:\n", sv)

// 	// command to remove specific value retrieved previously
// 	rm, _ := rdb.LRem(ctx, "lms_events", 1, sv).Result()
// 	fmt.Println("\nremove: '1 means ok'\n", rm)

// 	// command to remove a range list
// 	rl, _ := rdb.LPopCount(ctx, "lms_events", 9).Result()
// 	fmt.Println("\nremoveand return 9 values from oldest to newest \n", rl)

// 	// command to get the len of the key after removal
// 	llen, err := rdb.LLen(ctx, "lms_events").Result()
// 	fmt.Println("\nlen:\n", llen)

// 	// command to clean all database
// 	// rdb.FlushAll(ctx)
// }
