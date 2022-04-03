package main

import (
	"fmt"
	"math/rand"

	gongng2charts_models "github.com/fullstack-lang/gongng2charts/go/models"
)

const NbPoints = 5
const NbDatasets = 2

func GenerateChart() {
	GenerateChart1()
	GenerateChart2()
	gongng2charts_models.Stage.Commit()
}
func GenerateChart1() {

	chartConfig := (&gongng2charts_models.ChartConfiguration{Name: "Chart 1"}).Stage()
	chartConfig.ChartType = gongng2charts_models.LINE

	for idx_dataset := 0; idx_dataset < NbDatasets; idx_dataset = idx_dataset + 1 {
		dataset := (&gongng2charts_models.Dataset{Name: fmt.Sprintf("Dataset %d", 0)}).Stage()
		chartConfig.Datasets = append(chartConfig.Datasets, dataset)

		for idx := 0; idx < NbPoints; idx = idx + 1 {
			datapoint := (&gongng2charts_models.DataPoint{Name: fmt.Sprintf("Month %d", idx)}).Stage()
			datapoint.Value = rand.Float64() * 100.0
			dataset.DataPoints = append(dataset.DataPoints, datapoint)
		}

		dataset.Label = fmt.Sprintf("set %d", idx_dataset)
	}

	for idx := 0; idx < NbPoints; idx = idx + 1 {
		label := (&gongng2charts_models.Label{Name: fmt.Sprintf("Month %d", idx)}).Stage()
		chartConfig.Labels = append(chartConfig.Labels, label)
	}
}

func GenerateChart2() {

	chartConfig := (&gongng2charts_models.ChartConfiguration{Name: "Chart 2"}).Stage()
	chartConfig.ChartType = gongng2charts_models.BAR

	for idx_dataset := 0; idx_dataset < NbDatasets; idx_dataset = idx_dataset + 1 {
		dataset := (&gongng2charts_models.Dataset{Name: fmt.Sprintf("Dataset 2 %d", 0)}).Stage()
		chartConfig.Datasets = append(chartConfig.Datasets, dataset)

		for idx := 0; idx < NbPoints; idx = idx + 1 {
			datapoint := (&gongng2charts_models.DataPoint{Name: fmt.Sprintf("Day %d", idx)}).Stage()
			datapoint.Value = rand.Float64() * 30.0
			dataset.DataPoints = append(dataset.DataPoints, datapoint)
		}

		dataset.Label = fmt.Sprintf("set %d", idx_dataset)
	}

	for idx := 0; idx < NbPoints; idx = idx + 1 {
		label := (&gongng2charts_models.Label{Name: fmt.Sprintf("Day %d", idx)}).Stage()
		chartConfig.Labels = append(chartConfig.Labels, label)
	}
}
