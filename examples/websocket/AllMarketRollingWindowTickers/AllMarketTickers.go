package main

import (
	"fmt"

	binance_connector "github.com/capsis-finance/capsis-binance-client-go"
)

func main() {
	WsAllMarketTickersExample()
}

func WsAllMarketTickersExample() {
	websocketStreamClient := binance_connector.NewWebsocketStreamClient(false)
	wsAllMarketTickersHandler := func(event binance_connector.WsAllMarketRollingWindowTickersStatEvent) {

		for _, e := range event {
			if e.Symbol == "BTCUSDT" {
				fmt.Println(binance_connector.PrettyPrint(e))
			}
		}
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, err := websocketStreamClient.WsAllMarketRollingWindowTickersStatServe("1h", wsAllMarketTickersHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneCh
}
