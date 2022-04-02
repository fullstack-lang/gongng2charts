// insertion point sub template for components imports 
  import { ChartsTableComponent } from './charts-table/charts-table.component'
  import { ChartSortingComponent } from './chart-sorting/chart-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfChartsComponents: Map<string, any> = new Map([["ChartsTableComponent", ChartsTableComponent],])
  export const MapOfChartSortingComponents: Map<string, any> = new Map([["ChartSortingComponent", ChartSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Chart", MapOfChartsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Chart", MapOfChartSortingComponents],
    ]
  )
