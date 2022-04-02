package main

import (
	"fmt"
	"math/rand"

	gongng2charts_models "github.com/fullstack-lang/gongng2charts/go/models"
)

const NbPoints = 20

func GenerateChart() {

	chartConfig := (&gongng2charts_models.ChartConfiguration{Name: "Dummy"}).Stage()

	dataset := (&gongng2charts_models.Dataset{Name: fmt.Sprintf("Dataset %d", 0)}).Stage()
	chartConfig.Datasets = append(chartConfig.Datasets, dataset)

	for idx := 0; idx < NbPoints; idx = idx + 1 {
		label := (&gongng2charts_models.Label{Name: fmt.Sprintf("Month %d", idx)}).Stage()
		chartConfig.Labels = append(chartConfig.Labels, label)

		datapoint := (&gongng2charts_models.DataPoint{Name: fmt.Sprintf("Month %d", idx)}).Stage()
		datapoint.Value = rand.Float64() * 100.0
		dataset.DataPoints = append(dataset.DataPoints, datapoint)
	}

	gongng2charts_models.Stage.Commit()
}
