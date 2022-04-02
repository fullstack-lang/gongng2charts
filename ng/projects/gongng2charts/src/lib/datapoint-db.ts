// insertion point for imports
import { DatasetDB } from './dataset-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DataPointDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: number = 0

	// insertion point for other declarations
	Dataset_DataPointsDBID: NullInt64 = new NullInt64
	Dataset_DataPointsDBID_Index: NullInt64  = new NullInt64 // store the index of the datapoint instance in Dataset.DataPoints
	Dataset_DataPoints_reverse?: DatasetDB 

}
