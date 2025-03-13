package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"math/rand"
	"os"
)

func gaugeChart() *charts.Gauge {
	gauge := charts.NewGauge()
	gauge.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "CPU Usage"}))
	gauge.AddSeries("Usage", []opts.GaugeData{
		{Name: "Usage", Value: rand.Intn(100)},
	})
	return gauge
}

func gridChart() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Server Load"}))

	servers := []string{"Server-1", "Server-2", "Server-3", "Server-4", "Server-5"}
	values := []opts.BarData{}
	for _, _ = range servers {
		values = append(values, opts.BarData{Value: rand.Intn(100)})
	}
	bar.SetXAxis(servers).AddSeries("Load", values)
	return bar
}

func histogramChart() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Latency Distribution"}))

	latencyBuckets := []string{"0-50ms", "50-100ms", "100-200ms", "200-500ms", "500+ms"}
	values := []opts.BarData{}
	for _, _ = range latencyBuckets {
		values = append(values, opts.BarData{Value: rand.Intn(50)})
	}

	bar.SetXAxis(latencyBuckets).AddSeries("Requests", values)
	return bar
}

func drillDownChart() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Disk Usage"}))

	data := []opts.PieData{
		{Name: "Used", Value: rand.Intn(500)},
		{Name: "Free", Value: rand.Intn(500)},
	}

	pie.AddSeries("Usage", data)
	return pie
}

func main() {
	page := components.NewPage()
	page.AddCharts(
		gaugeChart(),
		gridChart(),
		drillDownChart(),
	)

	f, _ := os.Create("charts.html")
	page.Render(f)
}
