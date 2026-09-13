package utils

import "github.com/go-echarts/go-echarts/v2/opts"

func FloatsToLineData(values []float32) []opts.LineData {
	items := make([]opts.LineData, len(values))
	for i, v := range values {
		items[i] = opts.LineData{Value: v}
	}
	return items
}
