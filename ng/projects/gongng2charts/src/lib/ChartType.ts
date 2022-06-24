// generated from ng_file_enum.ts.go
export enum ChartType {
	// insertion point	
	LINE = "line",
	BAR = "bar",
	RADAR = "radar",
	PIE = "pie",
	POLAR_AREA = "polarArea",
	DOUGHNUT = "doughnut",
}

export interface ChartTypeSelect {
	value: string;
	viewValue: string;
}

export const ChartTypeList: ChartTypeSelect[] = [ // insertion point	
	{ value: ChartType.LINE, viewValue: "line" },
	{ value: ChartType.BAR, viewValue: "bar" },
	{ value: ChartType.RADAR, viewValue: "radar" },
	{ value: ChartType.PIE, viewValue: "pie" },
	{ value: ChartType.POLAR_AREA, viewValue: "polarArea" },
	{ value: ChartType.DOUGHNUT, viewValue: "doughnut" },
];
