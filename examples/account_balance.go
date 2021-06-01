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

	fmt.Printf("%+v\n\n", resp)

	historyReq := nanocoinrpc.AccountHistoryRequest{
		Account: "nano_37ortkby6k68z8tkk8g63ndbp8wjbmofhn56oyxb4rm6s3x51pkpiwcnpgmq",
		Count:   "10",
	}

	historyResp, err := nano.Nano.AccountHistory(historyReq)
	if err != nil {
		fmt.Println("historyResp", err)
		return
	}

	fmt.Printf("%+v\n\n", historyResp)

	infoReq := nanocoinrpc.AccountInfoRequest{
		Account:          "nano_3t6k35gi95xu6tergt6p69ck76ogmitsa8mnijtpxm9fkcm736xtoncuohr3",
		Representative:   "true",
		Weight:           "true",
		Pending:          "true",
		IncludeConfirmed: "true",
	}

	infoResp, err := nano.Nano.AccountInfo(infoReq)
	if err != nil {
		fmt.Println("infoResp", err)
		return
	}

	fmt.Printf("%+v\n\n", infoResp)

	balReq := nanocoinrpc.AccountsBalancesRequest{
		Accounts: []string{"nano_37ortkby6k68z8tkk8g63ndbp8wjbmofhn56oyxb4rm6s3x51pkpiwcnpgmq"},
	}

	balResp, err := nano.Nano.AccountsBalances(balReq)
	if err != nil {
		fmt.Println("balResp", err)
		return
	}

	fmt.Printf("%+v\n\n", balResp)

}
