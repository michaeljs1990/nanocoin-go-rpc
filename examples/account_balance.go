package main

import (
	"fmt"

	nanocoinrpc "github.com/michaeljs1990/nanocoin-go-rpc"
)

func main() {
	nano, err := nanocoinrpc.NewClient("http://[::1]:7076")
	if err != nil {
		fmt.Println(err)
		return
	}

	req := nanocoinrpc.AccountBalanceRequest{
		Account: "nano_37ortkby6k68z8tkk8g63ndbp8wjbmofhn56oyxb4rm6s3x51pkpiwcnpgmq",
	}

	resp, err := nano.Nano.AccountBalance(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", resp)

	historyReq := nanocoinrpc.AccountHistoryRequest{
		Account: "nano_37ortkby6k68z8tkk8g63ndbp8wjbmofhn56oyxb4rm6s3x51pkpiwcnpgmq",
		Count:   "1",
	}

	historyResp, err := nano.Nano.AccountHistory(historyReq)
	if err != nil {
		fmt.Println("historyResp", err)
		return
	}

	fmt.Printf("%+v\n", historyResp)
}
