package models

type ChartType string

// values for EnumType
const (
	LINE       ChartType = "line"
	BAR        ChartType = "bar"
	RADAR      ChartType = "radar"
	PIE        ChartType = "pie"
	POLAR_AREA ChartType = "polarArea"
	DOUGHNUT   ChartType = "doughnut"
)
