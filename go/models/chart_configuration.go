package models

type ChartConfiguration struct {
	Name      string
	Datasets  []*Dataset
	Labels    []*Label
	ChartType ChartType
}
