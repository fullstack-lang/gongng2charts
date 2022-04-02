package main

import (
	"fmt"
	"math/rand"

	gongng2charts_models "github.com/fullstack-lang/gongng2charts/go/models"
)

const NbPoints = 20
const NbDatasets = 4

func GenerateChart() {

	chartConfig := (&gongng2charts_models.ChartConfiguration{Name: "Dummy"}).Stage()
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

	gongng2charts_models.Stage.Commit()
}
