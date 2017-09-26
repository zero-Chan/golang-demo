package main

import (
	"fmt"

	"sunteng/commons/confutil"

	"heimdall/conf"
	api "heimdall/models/adx/acelink"
)

func main() {
	err := Init()
	if err != nil {
		fmt.Println("Init fail. error:", err)
		//		return
	}

	// add
	err = Add()
	if err != nil {
		fmt.Println("Add api error:", err)
		//		return
	}

	// upload
	err = Update()
	if err != nil {
		fmt.Println("Upload api error:", err)
		//		return
	}

	// get
	err = Get()
	if err != nil {
		fmt.Println("Get api error:", err)
		//		return
	}
}

func Init() error {
	advconf := conf.AdxTplConf{}
	err := confutil.LoadExtendConf("../files/adx/acelink.json", &advconf)
	if err != nil {
		return err
	}

	api.InitCfg(advconf)

	return nil
}

func Add() (err error) {
	// PC Banner
	pms := api.Creative{
		//		CreativeId:
		AdType:      1,                      // pc banner
		CreativeUrl: "http://www.baidu.com", // 改这里
		TargetUrl:   "http://www.baidu.com", // 改这里
		LandingPage: "http://www.baidu.com", // 改这里
		MonitorUrls: []string{
			"http://rtb.cc.com/vw?info=%%EXT%%&wp=%%WINPRICE%%", // 改这里
		},
		Height:          500,
		Width:           500,
		CreativeTradeId: 2,
		AdvertiserId:    101, // 改这里
	}

	cts := make([]api.Creative, 0)
	cts = append(cts, pms)

	cli := api.NewCreativeAddClient()

	result, err := cli.Add(cts)
	if err != nil {
		return err
	}

	fmt.Println("add result:", result)

	return
}

func Update() (err error) {
	// PC Banner
	pms := api.Creative{
		CreativeId:  6308278,
		AdType:      1,                      // pc banner
		CreativeUrl: "http://www.baidu.com", // 改这里
		TargetUrl:   "http://www.baidu.com", // 改这里
		LandingPage: "http://www.baidu.com", // 改这里
		MonitorUrls: []string{
			"http://rtb.cc.com/vw?info=%%EXT%%&wp=%%WINPRICE%%", // 改这里
		},
		Height:          500,
		Width:           500,
		CreativeTradeId: 2,
		AdvertiserId:    101, // 改这里
	}

	cts := make([]api.Creative, 0)
	cts = append(cts, pms)

	cli := api.NewCreativeUpdateClient()

	result, err := cli.Update(cts)
	if err != nil {
		return err
	}

	fmt.Println("add result:", result)

	return
}

func Get() (err error) {
	return
}
