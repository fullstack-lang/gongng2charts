package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *ChartConfiguration:
		if stage.OnAfterChartConfigurationCreateCallback != nil {
			stage.OnAfterChartConfigurationCreateCallback.OnAfterCreate(stage, target)
		}
	case *DataPoint:
		if stage.OnAfterDataPointCreateCallback != nil {
			stage.OnAfterDataPointCreateCallback.OnAfterCreate(stage, target)
		}
	case *Dataset:
		if stage.OnAfterDatasetCreateCallback != nil {
			stage.OnAfterDatasetCreateCallback.OnAfterCreate(stage, target)
		}
	case *Label:
		if stage.OnAfterLabelCreateCallback != nil {
			stage.OnAfterLabelCreateCallback.OnAfterCreate(stage, target)
		}
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *ChartConfiguration:
		newTarget := any(new).(*ChartConfiguration)
		if stage.OnAfterChartConfigurationUpdateCallback != nil {
			stage.OnAfterChartConfigurationUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DataPoint:
		newTarget := any(new).(*DataPoint)
		if stage.OnAfterDataPointUpdateCallback != nil {
			stage.OnAfterDataPointUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Dataset:
		newTarget := any(new).(*Dataset)
		if stage.OnAfterDatasetUpdateCallback != nil {
			stage.OnAfterDatasetUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Label:
		newTarget := any(new).(*Label)
		if stage.OnAfterLabelUpdateCallback != nil {
			stage.OnAfterLabelUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *ChartConfiguration:
		if stage.OnAfterChartConfigurationDeleteCallback != nil {
			staged := any(staged).(*ChartConfiguration)
			stage.OnAfterChartConfigurationDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DataPoint:
		if stage.OnAfterDataPointDeleteCallback != nil {
			staged := any(staged).(*DataPoint)
			stage.OnAfterDataPointDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Dataset:
		if stage.OnAfterDatasetDeleteCallback != nil {
			staged := any(staged).(*Dataset)
			stage.OnAfterDatasetDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Label:
		if stage.OnAfterLabelDeleteCallback != nil {
			staged := any(staged).(*Label)
			stage.OnAfterLabelDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *ChartConfiguration:
		if stage.OnAfterChartConfigurationReadCallback != nil {
			stage.OnAfterChartConfigurationReadCallback.OnAfterRead(stage, target)
		}
	case *DataPoint:
		if stage.OnAfterDataPointReadCallback != nil {
			stage.OnAfterDataPointReadCallback.OnAfterRead(stage, target)
		}
	case *Dataset:
		if stage.OnAfterDatasetReadCallback != nil {
			stage.OnAfterDatasetReadCallback.OnAfterRead(stage, target)
		}
	case *Label:
		if stage.OnAfterLabelReadCallback != nil {
			stage.OnAfterLabelReadCallback.OnAfterRead(stage, target)
		}
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *ChartConfiguration:
		stage.OnAfterChartConfigurationUpdateCallback = any(callback).(OnAfterUpdateInterface[ChartConfiguration])
	
	case *DataPoint:
		stage.OnAfterDataPointUpdateCallback = any(callback).(OnAfterUpdateInterface[DataPoint])
	
	case *Dataset:
		stage.OnAfterDatasetUpdateCallback = any(callback).(OnAfterUpdateInterface[Dataset])
	
	case *Label:
		stage.OnAfterLabelUpdateCallback = any(callback).(OnAfterUpdateInterface[Label])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *ChartConfiguration:
		stage.OnAfterChartConfigurationCreateCallback = any(callback).(OnAfterCreateInterface[ChartConfiguration])
	
	case *DataPoint:
		stage.OnAfterDataPointCreateCallback = any(callback).(OnAfterCreateInterface[DataPoint])
	
	case *Dataset:
		stage.OnAfterDatasetCreateCallback = any(callback).(OnAfterCreateInterface[Dataset])
	
	case *Label:
		stage.OnAfterLabelCreateCallback = any(callback).(OnAfterCreateInterface[Label])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *ChartConfiguration:
		stage.OnAfterChartConfigurationDeleteCallback = any(callback).(OnAfterDeleteInterface[ChartConfiguration])
	
	case *DataPoint:
		stage.OnAfterDataPointDeleteCallback = any(callback).(OnAfterDeleteInterface[DataPoint])
	
	case *Dataset:
		stage.OnAfterDatasetDeleteCallback = any(callback).(OnAfterDeleteInterface[Dataset])
	
	case *Label:
		stage.OnAfterLabelDeleteCallback = any(callback).(OnAfterDeleteInterface[Label])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *ChartConfiguration:
		stage.OnAfterChartConfigurationReadCallback = any(callback).(OnAfterReadInterface[ChartConfiguration])
	
	case *DataPoint:
		stage.OnAfterDataPointReadCallback = any(callback).(OnAfterReadInterface[DataPoint])
	
	case *Dataset:
		stage.OnAfterDatasetReadCallback = any(callback).(OnAfterReadInterface[Dataset])
	
	case *Label:
		stage.OnAfterLabelReadCallback = any(callback).(OnAfterReadInterface[Label])
	
	}
}
