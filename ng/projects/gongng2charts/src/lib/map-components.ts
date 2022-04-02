// insertion point sub template for components imports 
  import { ChartsTableComponent } from './charts-table/charts-table.component'
  import { ChartSortingComponent } from './chart-sorting/chart-sorting.component'
  import { ChartConfigurationsTableComponent } from './chartconfigurations-table/chartconfigurations-table.component'
  import { ChartConfigurationSortingComponent } from './chartconfiguration-sorting/chartconfiguration-sorting.component'
  import { DataPointsTableComponent } from './datapoints-table/datapoints-table.component'
  import { DataPointSortingComponent } from './datapoint-sorting/datapoint-sorting.component'
  import { DatasetsTableComponent } from './datasets-table/datasets-table.component'
  import { DatasetSortingComponent } from './dataset-sorting/dataset-sorting.component'
  import { LabelsTableComponent } from './labels-table/labels-table.component'
  import { LabelSortingComponent } from './label-sorting/label-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfChartsComponents: Map<string, any> = new Map([["ChartsTableComponent", ChartsTableComponent],])
  export const MapOfChartSortingComponents: Map<string, any> = new Map([["ChartSortingComponent", ChartSortingComponent],])
  export const MapOfChartConfigurationsComponents: Map<string, any> = new Map([["ChartConfigurationsTableComponent", ChartConfigurationsTableComponent],])
  export const MapOfChartConfigurationSortingComponents: Map<string, any> = new Map([["ChartConfigurationSortingComponent", ChartConfigurationSortingComponent],])
  export const MapOfDataPointsComponents: Map<string, any> = new Map([["DataPointsTableComponent", DataPointsTableComponent],])
  export const MapOfDataPointSortingComponents: Map<string, any> = new Map([["DataPointSortingComponent", DataPointSortingComponent],])
  export const MapOfDatasetsComponents: Map<string, any> = new Map([["DatasetsTableComponent", DatasetsTableComponent],])
  export const MapOfDatasetSortingComponents: Map<string, any> = new Map([["DatasetSortingComponent", DatasetSortingComponent],])
  export const MapOfLabelsComponents: Map<string, any> = new Map([["LabelsTableComponent", LabelsTableComponent],])
  export const MapOfLabelSortingComponents: Map<string, any> = new Map([["LabelSortingComponent", LabelSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Chart", MapOfChartsComponents],
      ["ChartConfiguration", MapOfChartConfigurationsComponents],
      ["DataPoint", MapOfDataPointsComponents],
      ["Dataset", MapOfDatasetsComponents],
      ["Label", MapOfLabelsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Chart", MapOfChartSortingComponents],
      ["ChartConfiguration", MapOfChartConfigurationSortingComponents],
      ["DataPoint", MapOfDataPointSortingComponents],
      ["Dataset", MapOfDatasetSortingComponents],
      ["Label", MapOfLabelSortingComponents],
    ]
  )
