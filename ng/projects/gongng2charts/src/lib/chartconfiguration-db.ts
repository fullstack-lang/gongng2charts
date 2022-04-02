// insertion point for imports
import { DatasetDB } from './dataset-db'
import { LabelDB } from './label-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ChartConfigurationDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	ChartType: string = ""

	// insertion point for other declarations
	Datasets?: Array<DatasetDB>
	Labels?: Array<LabelDB>
}
