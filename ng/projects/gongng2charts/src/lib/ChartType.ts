// generated from ng_file_enum.ts.go
export enum ChartType {
	// insertion point	
	BAR = "bar",
	DOUGHNUT = "doughnut",
	LINE = "line",
	PIE = "pie",
	POLAR_AREA = "polarArea",
	RADAR = "radar",
}

export interface ChartTypeSelect {
	value: string;
	viewValue: string;
}

export const ChartTypeList: ChartTypeSelect[] = [ // insertion point	
	{ value: ChartType.BAR, viewValue: "bar" },
	{ value: ChartType.DOUGHNUT, viewValue: "doughnut" },
	{ value: ChartType.LINE, viewValue: "line" },
	{ value: ChartType.PIE, viewValue: "pie" },
	{ value: ChartType.POLAR_AREA, viewValue: "polarArea" },
	{ value: ChartType.RADAR, viewValue: "radar" },
];
