package main

import (
	"github.com/gostuding/musthave-metrics-tpl/cmd/agent/metrics"
)

func main() {
	GetFlags()
	// fmt.Println("Args values: ", SendAddress, UpdateTime, SendTime)
	Storage := new(metrics.MetricsStorage)
	// запуск функции для бесконечного цикла отправки сообщений
	DoAgent(Storage)
}
