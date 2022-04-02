// insertion point for imports
import { DataPointDB } from './datapoint-db'
import { ChartConfigurationDB } from './chartconfiguration-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DatasetDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Label: string = ""

	// insertion point for other declarations
	DataPoints?: Array<DataPointDB>
	ChartConfiguration_DatasetsDBID: NullInt64 = new NullInt64
	ChartConfiguration_DatasetsDBID_Index: NullInt64  = new NullInt64 // store the index of the dataset instance in ChartConfiguration.Datasets
	ChartConfiguration_Datasets_reverse?: ChartConfigurationDB 

}
