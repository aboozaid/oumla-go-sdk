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
	address, err := sdk.Addresses.Generate(types.NetworkTBTC, "test-user3", "iHoIhxaktRkV/IoUhv3YTRTCv1bAGLXFbeuN6rH2yZTxCUcEwfRau6tVlzzVi8rDwwI+QMGs/101g+WZoUFMqUSsse52r2uGz/eVGz5T0swS+HIwZFVyp6UcaV+m1j1vtgaLlyDnSRBBGHkBW2FeUnc6Ol2jjGB+TiO5UvsQTKoZq36ghUwbjQNXgRV5+WOqdFv3KlPxH9AAWcm41Ps6B9EBT22NIGYfB184prnyxlNn5J3jSRV9Ovbp9fIrSDL/mJJL3kbackngn7SE4TTR38YPux5xM2tjfhL/tfxIaeoCIs0hqJu6CpL7g0YNgCwu")
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