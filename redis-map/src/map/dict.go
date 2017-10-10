package main

import (
	"flag"
	"fmt"
	"time"

	"sunteng/commons/confutil/light"

	gn_adx "gungnir/model/adx"
)

func main() {
	flag.Parse()

	err := light.LoadIndex()
	if err != nil {
		fmt.Println(err)
		return
	}

	gn_adx.InitDict([]int64{20066}, gn_adx.CATEGORY)

	fmt.Printf("\n\n\n===== Start: get value in golang map =====\n")
	N := len(gn_adx.Dict.AdvCate[20066].Adx)
	fmt.Printf("N = [%d]\n", N)

	// random get key
	randomnums := []int{
		10, 50, 100, // %
	}

	for _, rnum := range randomnums {
		getn := int64(N * rnum / 100)
		keys := make([]int64, getn)
		// get map is random

		var i int64
		for key, _ := range gn_adx.Dict.AdvCate[20066].Adx {
			keys[i] = key
			i++
			if i == getn {
				break
			}
		}

		start := time.Now()
		for _, key := range keys {
			gn_adx.Dict.AdvCate.GetBid(20066, key)
		}
		end := time.Now()

		fmt.Printf("when N=[%d], getn=[%d], time=[%s]\n", N, getn, end.Sub(start).String())
	}

}
