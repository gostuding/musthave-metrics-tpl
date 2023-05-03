package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type gauge struct {
	name  string
	value float64
}

type counter struct {
	name  string
	value int64
}

type MSWriter interface {
	AddMetric(string) (int, error)
	// Print() string
}

type MemStorage struct {
	gauges   []gauge
	counters []counter
}

func NewMemStorage() *MemStorage {
	return new(MemStorage)
}

func (ms *MemStorage) AddMetric(path string) (int, error) {
	items := strings.Split(path, "/")
	if len(items) >= 5 {
		items = items[2:5]
	}
	if len(items) != 3 {
		fmt.Printf("Metric parse error. Len not equal to 3. %v\n", items)
		return http.StatusBadRequest, errors.New("metric parse error. Checks url path and repeat")
	}
	if items[0] == "gauge" {
		val, err := strconv.ParseFloat(items[2], 64)
		if err != nil {
			return http.StatusBadRequest, err
		}
		ms.addGauge(items[1], val)
	} else if items[0] == "counter" {
		val, err := strconv.ParseInt(items[2], 10, 64)
		if err != nil {
			return http.StatusBadRequest, err
		}
		ms.addCounter(items[1], val)
	} else {
		fmt.Printf("Metric type incorrect. Type is: %s\n", items[0])
		return http.StatusBadRequest, errors.New("metric type incorrect. Availible types are: guage or counter")
	}
	return http.StatusOK, nil
}

func (ms MemStorage) Print() string {
	body := "==== MemoryStorage ====\n"
	for index, value := range ms.gauges {
		body += fmt.Sprintf("Gauges: %d, name: '%s', value: '%f'\n", index, value.name, value.value)
	}
	for index, value := range ms.counters {
		body += fmt.Sprintf("Counter: %d, name: '%s', value: '%d'\n", index, value.name, value.value)
	}
	return body
}

func (ms *MemStorage) addGauge(name string, value float64) {
	for index, item := range ms.gauges {
		if item.name == name {
			ms.gauges[index].value = value
			return
		}
	}
	ms.gauges = append(ms.gauges, gauge{name: name, value: value})
}

func (ms *MemStorage) addCounter(name string, value int64) {
	for index, item := range ms.counters {
		if item.name == name {
			ms.counters[index].value += value
			return
		}
	}
	ms.counters = append(ms.counters, counter{name: name, value: value})
}
