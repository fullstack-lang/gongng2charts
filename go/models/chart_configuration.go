package models

type ChartConfiguration struct {
	Name      string
	Datasets  []*Dataset
	Labels    []*Label
	ChartType ChartType

	// width and height of the canvas object <canvas baseChart [width]="width" [height]="height"
	Width  int
	Heigth int
}
