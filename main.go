package main

import (
	rd "64HW/redis"
	"context"
	"fmt"
	"log"
)

var ctx = context.Background()

func main() {
	rdb := rd.NewClient()
	defer rdb.Close()

	err := rdb.Set(ctx, "mykey", "myvalue", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	val, err := rdb.Get(ctx, "mykey").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mykey:", val)

	err = rdb.Set(ctx, "counter", 0, 0).Err()
	if err != nil {
		log.Fatal(err)
	}
	err = rdb.Incr(ctx, "counter").Err()
	if err != nil {
		log.Fatal(err)
	}
	counter, err := rdb.Get(ctx, "counter").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("counter:", counter)

	err = rdb.HSet(ctx, "user:1", map[string]interface{}{
		"id":    1,
		"name":  "shaxboz",
		"email": "shaxbozakramovic@gmail.com",
	}).Err()
	if err != nil {
		log.Fatal(err)
	}

	err = rdb.HMSet(ctx, "user:1", map[string]interface{}{
		"name":  "shaxboz",
		"email": "shaxbozakramovic@gmail.com",
	}).Err()
	if err != nil {
		log.Fatal(err)
	}

	user, err := rdb.HGetAll(ctx, "user:1").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user:1:", user)

	err = rdb.RPush(ctx, "mylist", "element1").Err()
	if err != nil {
		log.Fatal(err)
	}
	err = rdb.RPush(ctx, "mylist", "element2").Err()
	if err != nil {
		log.Fatal(err)
	}

	elements, err := rdb.LRange(ctx, "mylist", 0, -1).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mylist:", elements)

	err = rdb.LRem(ctx, "mylist", 0, "element1").Err()
	if err != nil {
		log.Fatal(err)
	}
	updatedList, err := rdb.LRange(ctx, "mylist", 0, -1).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated mylist:", updatedList)

	go func() {
		pubsub := rdb.Subscribe(ctx, "mychannel")
		_, err := pubsub.Receive(ctx)
		if err != nil {
			log.Fatal(err)
		}

		ch := pubsub.Channel()
		for msg := range ch {
			fmt.Println("Received message:", msg.Payload)
		}
	}()

	err = rdb.Publish(ctx, "mychannel", "hello world").Err()
	if err != nil {
		log.Fatal(err)
	}

}
