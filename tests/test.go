package main

import (
	"encoding/json"
	"fmt"
	oumlagosdk "oumla-go-sdk"
	"oumla-go-sdk/types"
)

func main() {
	sdk := oumlagosdk.NewOumla(oumlagosdk.Config{
		BaseURL: "https://sandbox.oumla.com/api/v1",
		ApiKey:  "oumla_3ZKFM9QFFqa1f1tqHVtF3EYY",
	})
	fmt.Println("sdk initialized")
	_, err := sdk.Wallets.Generate(types.NetworkTBTC, "test-user3")
	if err != nil {
		panic(err)
	}
	address, err := sdk.Addresses.Generate(types.NetworkTBTC, "test-user3", "")
	if err != nil {
		js, _ := json.Marshal(err)
		fmt.Println(string(js))
		panic(err)
	}
	// fmt.Println(resp.Data[0])
	// js, err := json.Marshal(resp)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println(address)
}