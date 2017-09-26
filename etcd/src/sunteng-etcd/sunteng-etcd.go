package main

import (
	"fmt"

	"sunteng/commons/confutil"
	"sunteng/commons/db/myetcd"
)

func main() {
	myetcd.Init(myetcd.EtcdConfig{
		Addrs: []confutil.DbConfig{
			confutil.DbConfig{
				NetBase: confutil.NetBase{
					Host: "127.0.0.1",
					Port: 2379,
				},
			},
		},
		Root: "/test",
	})
	cli := myetcd.NewClient()
	defer cli.Stop()

	// set value
	resp, err := cli.Slipper.Set("/test/x1", "x1v", 0)
	if err != nil {
		fmt.Println("etcd client set fail: ", err)
		return
	}

	fmt.Printf("etcd client set resp: %+v\n", resp)

	// get value
	resp, err = cli.Slipper.Get("/test/x1", true, true)
	if err != nil {
		fmt.Println("etcd client get fail: ", err)
		return
	}

	fmt.Printf("etcd client get: key: %s, value: %s\n", resp.Node.Key, resp.Node.Value)

	go cli.WatchRec("/test/x1")

	go cli.WatchAndCompare("/test/x1", "update", "x1v2", 10)

	cli.Slipper.Update("test/x1", "x1v2", 20)
	cli.Slipper.SetDir("/test/ooo", 4)
	cli.Slipper.SetDir("/test/ooo", 0)

}
