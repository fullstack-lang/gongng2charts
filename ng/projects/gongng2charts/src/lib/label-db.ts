// insertion point for imports
import { ChartConfigurationDB } from './chartconfiguration-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LabelDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	ChartConfiguration_LabelsDBID: NullInt64 = new NullInt64
	ChartConfiguration_LabelsDBID_Index: NullInt64  = new NullInt64 // store the index of the label instance in ChartConfiguration.Labels
	ChartConfiguration_Labels_reverse?: ChartConfigurationDB 

}
