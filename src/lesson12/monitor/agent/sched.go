package main

import (
	"lesson12/monitor/common"
	"time"
)

type MetricFunc func() []*common.Metric

type Sched struct {
	ch *common.Metric
}

func NewSched(ch chan *common.Metric) *Sched {

}

func (s *Sched) AddMetric(collecter MetricFunc, step time.Duration) {

}
