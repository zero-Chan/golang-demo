package main

import (
	"flag"
	"fmt"
	"time"

	"sunteng/commons/confutil/light"

	gn_adx "gungnir/model/adx"

	"gopkg.in/redis.v3"
)

func main() {
	flag.Parse()

	err := light.LoadIndex()
	if err != nil {
		fmt.Println(err)
		return
	}

	gn_adx.InitDict([]int64{20066}, gn_adx.CATEGORY)

	cli := redis.NewClient(&redis.Options{
		Addr: "localhost:10379",
	})

	cmd := cli.Ping()
	err = cmd.Err()
	if err != nil {
		return
	}

	fmt.Printf("\n\n\n===== Start: get value in golang map =====\n")
	N := len(gn_adx.Dict.AdvCate[20066].Adx)
	fmt.Printf("N = [%d]\n", N)

	// random get key
	randomnums := []int{
		10, 50, 100, // %
	}

	for _, rnum := range randomnums {
		getn := int64(N * rnum / 100)
		keys := make([]string, getn)
		// get map is random

		var i int64
		for key, _ := range gn_adx.Dict.AdvCate[20066].Adx {
			keys[i] = fmt.Sprintf("%d", key)
			i++
			if i == getn {
				break
			}
		}

		start := time.Now()
		for _, key := range keys {
			cli.HGet("category_dict_20066_adxbsw", key)
		}
		end := time.Now()

		fmt.Printf("when N=[%d], getn=[%d], time=[%s]\n", N, getn, end.Sub(start).String())
	}

}
