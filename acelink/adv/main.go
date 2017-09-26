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

func Add() error {
	at := api.Advertiser{
		AdvertiserId:     101,
		AdvertiserName:   "name1",                // 改这里
		AdvertiserIdsite: "http://www.baidu.com", // 改这里
		Qualifications:   make([]*api.Qualification, 0),
	}

	q1 := &api.Qualification{
		Name: "营业执照",
		Url:  "http://tieba.baidu.com", // 改这里
	}
	at.Qualifications = append(at.Qualifications, q1)

	// requests
	ats := make([]*api.Advertiser, 0)
	ats = append(ats, &at)

	// add
	cli := api.NewAdvertiserAddClient()
	fmt.Println("add URI:", cli.Url)

	result, err := cli.Add(ats)
	if err != nil {
		return err
	}

	fmt.Println("add result:", result)

	return nil
}

func Update() error {
	at := api.Advertiser{
		AdvertiserId:     101,
		AdvertiserName:   "name1",                // 改这里
		AdvertiserIdsite: "http://www.baidu.com", // 改这里
		Qualifications:   make([]*api.Qualification, 0),
	}

	q1 := &api.Qualification{
		Name: "营业执照",
		Url:  "http://tieba.baidu.com", // 改这里
	}
	at.Qualifications = append(at.Qualifications, q1)

	// requests
	ats := make([]*api.Advertiser, 0)
	ats = append(ats, &at)

	// update
	cli := api.NewAdvertiserUpdateClient()
	fmt.Println("update URI:", cli.Url)

	result, err := cli.Update(ats)
	if err != nil {
		return err
	}

	fmt.Println("update result:", result)

	return nil
}

func Get() error {
	advIds := []int64{
		101,
	}

	// get
	cli := api.NewAdvertiserGetClient()
	fmt.Println("get URI:", cli.Url)

	result, err := cli.Get(advIds)
	if err != nil {
		return err
	}

	fmt.Println("get result:", result)

	return nil
}
