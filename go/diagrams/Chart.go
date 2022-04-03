package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongng2charts/go/models"
)

var Chart uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.ChartConfiguration{}),
			Position: &uml.Position{
				X: 54.000000,
				Y: 75.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Links: []*uml.Link{
				{
					Field: models.ChartConfiguration{}.Datasets,
					Middlevertice: &uml.Vertice{
						X: 333.500000,
						Y: 183.500000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
				{
					Field: models.ChartConfiguration{}.Labels,
					Middlevertice: &uml.Vertice{
						X: 360.500000,
						Y: 104.000000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.ChartConfiguration{}.ChartType,
				},
				{
					Field: models.ChartConfiguration{}.Name,
				},
			},
		},
		{
			Struct: &(models.DataPoint{}),
			Position: &uml.Position{
				X: 430.000000,
				Y: 270.000000,
			},
			Width:  240.000000,
			Heigth: 93.000000,
			Fields: []*uml.Field{
				{
					Field: models.DataPoint{}.Name,
				},
				{
					Field: models.DataPoint{}.Value,
				},
			},
		},
		{
			Struct: &(models.Dataset{}),
			Position: &uml.Position{
				X: 440.000000,
				Y: 160.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Links: []*uml.Link{
				{
					Field: models.Dataset{}.DataPoints,
					Middlevertice: &uml.Vertice{
						X: 825.000000,
						Y: 304.000000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Dataset{}.Label,
				},
				{
					Field: models.Dataset{}.Name,
				},
			},
		},
		{
			Struct: &(models.Label{}),
			Position: &uml.Position{
				X: 447.000000,
				Y: 75.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Fields: []*uml.Field{
				{
					Field: models.Label{}.Name,
				},
			},
		},
	},
}
