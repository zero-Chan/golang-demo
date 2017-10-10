package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"gopkg.in/redis.v3"
)

var (
	count *int = flag.Int("count", 10, "set count for dict.")
)

func init() {
	flag.Parse()
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := make(map[int][]int)

	fmt.Printf("dict amount [%d]\n", *count)
	for i := 0; i < *count; i++ {
		key := r.Intn(20000)
		for {
			if _, ok := data[key]; !ok {
				break
			}
			key = r.Intn(20000)
		}

		slin := r.Intn(5)
		for slin == 0 {
			slin = r.Intn(5)
		}

		val := RandomSlice(slin)
		data[key] = val
	}

	fmt.Printf("data: %+v\n", data)

	// write file
	err := SaveFile(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// write redis
	err = SaveRedis(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func RandomSlice(n int) []int {
	sli := make([]int, n)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		sli[i] = r.Intn(20000)
	}

	return sli
}

func SaveFile(data map[int][]int) (err error) {
	filename := "./resources/dict/category_dict_20066.txt"

	os.Remove(filename)
	fp, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer fp.Close()

	datastr := make(map[string]string)
	for key, sli := range data {
		var slistr string
		for _, n := range sli {
			if len(slistr) != 0 {
				curstr := fmt.Sprintf("%s,%d", slistr, n)
				slistr = curstr
			} else {
				slistr = fmt.Sprintf("%d", n)
			}
		}
		datastr[fmt.Sprintf("%d", key)] = slistr
	}

	// write
	for key, val := range datastr {
		line := fmt.Sprintf("%s %s\n", key, val)
		_, err = fp.Write([]byte(line))
		if err != nil {
			return
		}
	}

	return
}

func SaveRedis(data map[int][]int) (err error) {
	cli := redis.NewClient(&redis.Options{
		Addr: "localhost:10379",
	})

	cmd := cli.Ping()
	err = cmd.Err()
	if err != nil {
		return
	}

	intcmd := cli.Del("category_dict_20066_adxbsw", "category_dict_20066_bswadx")
	if intcmd.Err() != nil {
		err = intcmd.Err()
		return
	}

	adxbsw := make(map[string]string)
	bswadx := make(map[string]string)
	for key, sli := range data {
		var slistr string
		for _, n := range sli {
			if len(slistr) != 0 {
				curstr := fmt.Sprintf("%s,%d", slistr, n)
				slistr = curstr
			} else {
				slistr = fmt.Sprintf("%d", n)
			}

			// set bswadx
			adxslistr, ok := bswadx[fmt.Sprintf("%d", n)]
			if !ok {
				adxslistr = fmt.Sprintf("%d", key)
			} else {
				curstr := fmt.Sprintf("%s,%d", adxslistr, key)
				adxslistr = curstr
			}
			bswadx[fmt.Sprintf("%d", n)] = adxslistr
		}

		// set adxbsw
		adxbsw[fmt.Sprintf("%d", key)] = slistr
	}

	// adxbsw map to slice
	adxbswslice := make([]string, 0)
	for key, val := range adxbsw {
		adxbswslice = append(adxbswslice, key, val)
	}

	// bswadx map to slice
	bswadxslice := make([]string, 0)
	for key, val := range bswadx {
		bswadxslice = append(bswadxslice, key, val)
	}

	// hmset
	if len(adxbswslice) > 2 {
		k1 := adxbswslice[0]
		v1 := adxbswslice[1]

		statuscmd := cli.HMSet("category_dict_20066_adxbsw", k1, v1, adxbswslice[2:]...)
		if statuscmd.Err() != nil {
			err = statuscmd.Err()
			return
		}
	}

	if len(bswadxslice) > 2 {
		k1 := bswadxslice[0]
		v1 := bswadxslice[1]
		statuscmd := cli.HMSet("category_dict_20066_bswadx", k1, v1, bswadxslice[2:]...)
		if statuscmd.Err() != nil {
			err = statuscmd.Err()
			return
		}
	}

	return
}
