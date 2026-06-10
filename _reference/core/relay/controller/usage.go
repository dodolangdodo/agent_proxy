package controller

import "github.com/labring/aiproxy/core/model"

type RequestUsage struct {
	Usage   model.Usage
	Context model.UsageContext
}

func NewRequestUsage(usage model.Usage) RequestUsage {
	return RequestUsage{Usage: usage}
}
