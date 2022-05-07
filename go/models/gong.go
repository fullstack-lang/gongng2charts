// generated by ModelGongFileTemplate
package models

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
)

// swagger:ignore
type __void struct{}

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	GetFields() (res []string)
	GetFieldStringValue(fieldName string) (res string)
}

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct { // insertion point for definition of arrays registering instances
	ChartConfigurations           map[*ChartConfiguration]struct{}
	ChartConfigurations_mapString map[string]*ChartConfiguration

	DataPoints           map[*DataPoint]struct{}
	DataPoints_mapString map[string]*DataPoint

	Datasets           map[*Dataset]struct{}
	Datasets_mapString map[string]*Dataset

	Labels           map[*Label]struct{}
	Labels_mapString map[string]*Label

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitChartConfiguration(chartconfiguration *ChartConfiguration)
	CheckoutChartConfiguration(chartconfiguration *ChartConfiguration)
	CommitDataPoint(datapoint *DataPoint)
	CheckoutDataPoint(datapoint *DataPoint)
	CommitDataset(dataset *Dataset)
	CheckoutDataset(dataset *Dataset)
	CommitLabel(label *Label)
	CheckoutLabel(label *Label)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

// swagger:ignore instructs the gong compiler (gongc) to avoid this particular struct
var Stage StageStruct = StageStruct{ // insertion point for array initiatialisation
	ChartConfigurations:           make(map[*ChartConfiguration]struct{}),
	ChartConfigurations_mapString: make(map[string]*ChartConfiguration),

	DataPoints:           make(map[*DataPoint]struct{}),
	DataPoints_mapString: make(map[string]*DataPoint),

	Datasets:           make(map[*Dataset]struct{}),
	Datasets_mapString: make(map[string]*Dataset),

	Labels:           make(map[*Label]struct{}),
	Labels_mapString: make(map[string]*Label),

	// end of insertion point
	Map_GongStructName_InstancesNb: make(map[string]int),
}

func (stage *StageStruct) Commit() {
	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["ChartConfiguration"] = len(stage.ChartConfigurations)
	stage.Map_GongStructName_InstancesNb["DataPoint"] = len(stage.DataPoints)
	stage.Map_GongStructName_InstancesNb["Dataset"] = len(stage.Datasets)
	stage.Map_GongStructName_InstancesNb["Label"] = len(stage.Labels)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
func (stage *StageStruct) getChartConfigurationOrderedStructWithNameField() []*ChartConfiguration {
	// have alphabetical order generation
	chartconfigurationOrdered := []*ChartConfiguration{}
	for chartconfiguration := range stage.ChartConfigurations {
		chartconfigurationOrdered = append(chartconfigurationOrdered, chartconfiguration)
	}
	sort.Slice(chartconfigurationOrdered[:], func(i, j int) bool {
		return chartconfigurationOrdered[i].Name < chartconfigurationOrdered[j].Name
	})
	return chartconfigurationOrdered
}

// Stage puts chartconfiguration to the model stage
func (chartconfiguration *ChartConfiguration) Stage() *ChartConfiguration {
	Stage.ChartConfigurations[chartconfiguration] = __member
	Stage.ChartConfigurations_mapString[chartconfiguration.Name] = chartconfiguration

	return chartconfiguration
}

// Unstage removes chartconfiguration off the model stage
func (chartconfiguration *ChartConfiguration) Unstage() *ChartConfiguration {
	delete(Stage.ChartConfigurations, chartconfiguration)
	delete(Stage.ChartConfigurations_mapString, chartconfiguration.Name)
	return chartconfiguration
}

// commit chartconfiguration to the back repo (if it is already staged)
func (chartconfiguration *ChartConfiguration) Commit() *ChartConfiguration {
	if _, ok := Stage.ChartConfigurations[chartconfiguration]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitChartConfiguration(chartconfiguration)
		}
	}
	return chartconfiguration
}

// Checkout chartconfiguration to the back repo (if it is already staged)
func (chartconfiguration *ChartConfiguration) Checkout() *ChartConfiguration {
	if _, ok := Stage.ChartConfigurations[chartconfiguration]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutChartConfiguration(chartconfiguration)
		}
	}
	return chartconfiguration
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of chartconfiguration to the model stage
func (chartconfiguration *ChartConfiguration) StageCopy() *ChartConfiguration {
	_chartconfiguration := new(ChartConfiguration)
	*_chartconfiguration = *chartconfiguration
	_chartconfiguration.Stage()
	return _chartconfiguration
}

// StageAndCommit appends chartconfiguration to the model stage and commit to the orm repo
func (chartconfiguration *ChartConfiguration) StageAndCommit() *ChartConfiguration {
	chartconfiguration.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMChartConfiguration(chartconfiguration)
	}
	return chartconfiguration
}

// DeleteStageAndCommit appends chartconfiguration to the model stage and commit to the orm repo
func (chartconfiguration *ChartConfiguration) DeleteStageAndCommit() *ChartConfiguration {
	chartconfiguration.Unstage()
	DeleteORMChartConfiguration(chartconfiguration)
	return chartconfiguration
}

// StageCopyAndCommit appends a copy of chartconfiguration to the model stage and commit to the orm repo
func (chartconfiguration *ChartConfiguration) StageCopyAndCommit() *ChartConfiguration {
	_chartconfiguration := new(ChartConfiguration)
	*_chartconfiguration = *chartconfiguration
	_chartconfiguration.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMChartConfiguration(chartconfiguration)
	}
	return _chartconfiguration
}

// CreateORMChartConfiguration enables dynamic staging of a ChartConfiguration instance
func CreateORMChartConfiguration(chartconfiguration *ChartConfiguration) {
	chartconfiguration.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMChartConfiguration(chartconfiguration)
	}
}

// DeleteORMChartConfiguration enables dynamic staging of a ChartConfiguration instance
func DeleteORMChartConfiguration(chartconfiguration *ChartConfiguration) {
	chartconfiguration.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMChartConfiguration(chartconfiguration)
	}
}

// for satisfaction of GongStruct interface
func (chartconfiguration *ChartConfiguration) GetName() (res string) {
	return chartconfiguration.Name
}

func (chartconfiguration *ChartConfiguration) GetFields() (res []string) {
	// list of fields
	res = []string{"Name", "Datasets", "Labels", "ChartType", "Width", "Heigth"}
	return
}

func (chartconfiguration *ChartConfiguration) GetFieldStringValue(fieldName string) (res string) {
	switch fieldName {
	// string value of fields
	case "Name":
		res = chartconfiguration.Name
	case "Datasets":
		for idx, __instance__ := range chartconfiguration.Datasets {
			if idx > 0 {
				res += "\n"
			}
			res += __instance__.Name
		}
	case "Labels":
		for idx, __instance__ := range chartconfiguration.Labels {
			if idx > 0 {
				res += "\n"
			}
			res += __instance__.Name
		}
	case "ChartType":
		res = chartconfiguration.ChartType.ToCodeString()
	case "Width":
		res = fmt.Sprintf("%d", chartconfiguration.Width)
	case "Heigth":
		res = fmt.Sprintf("%d", chartconfiguration.Heigth)
	}
	return
}

func (stage *StageStruct) getDataPointOrderedStructWithNameField() []*DataPoint {
	// have alphabetical order generation
	datapointOrdered := []*DataPoint{}
	for datapoint := range stage.DataPoints {
		datapointOrdered = append(datapointOrdered, datapoint)
	}
	sort.Slice(datapointOrdered[:], func(i, j int) bool {
		return datapointOrdered[i].Name < datapointOrdered[j].Name
	})
	return datapointOrdered
}

// Stage puts datapoint to the model stage
func (datapoint *DataPoint) Stage() *DataPoint {
	Stage.DataPoints[datapoint] = __member
	Stage.DataPoints_mapString[datapoint.Name] = datapoint

	return datapoint
}

// Unstage removes datapoint off the model stage
func (datapoint *DataPoint) Unstage() *DataPoint {
	delete(Stage.DataPoints, datapoint)
	delete(Stage.DataPoints_mapString, datapoint.Name)
	return datapoint
}

// commit datapoint to the back repo (if it is already staged)
func (datapoint *DataPoint) Commit() *DataPoint {
	if _, ok := Stage.DataPoints[datapoint]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitDataPoint(datapoint)
		}
	}
	return datapoint
}

// Checkout datapoint to the back repo (if it is already staged)
func (datapoint *DataPoint) Checkout() *DataPoint {
	if _, ok := Stage.DataPoints[datapoint]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutDataPoint(datapoint)
		}
	}
	return datapoint
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of datapoint to the model stage
func (datapoint *DataPoint) StageCopy() *DataPoint {
	_datapoint := new(DataPoint)
	*_datapoint = *datapoint
	_datapoint.Stage()
	return _datapoint
}

// StageAndCommit appends datapoint to the model stage and commit to the orm repo
func (datapoint *DataPoint) StageAndCommit() *DataPoint {
	datapoint.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMDataPoint(datapoint)
	}
	return datapoint
}

// DeleteStageAndCommit appends datapoint to the model stage and commit to the orm repo
func (datapoint *DataPoint) DeleteStageAndCommit() *DataPoint {
	datapoint.Unstage()
	DeleteORMDataPoint(datapoint)
	return datapoint
}

// StageCopyAndCommit appends a copy of datapoint to the model stage and commit to the orm repo
func (datapoint *DataPoint) StageCopyAndCommit() *DataPoint {
	_datapoint := new(DataPoint)
	*_datapoint = *datapoint
	_datapoint.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMDataPoint(datapoint)
	}
	return _datapoint
}

// CreateORMDataPoint enables dynamic staging of a DataPoint instance
func CreateORMDataPoint(datapoint *DataPoint) {
	datapoint.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMDataPoint(datapoint)
	}
}

// DeleteORMDataPoint enables dynamic staging of a DataPoint instance
func DeleteORMDataPoint(datapoint *DataPoint) {
	datapoint.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMDataPoint(datapoint)
	}
}

// for satisfaction of GongStruct interface
func (datapoint *DataPoint) GetName() (res string) {
	return datapoint.Name
}

func (datapoint *DataPoint) GetFields() (res []string) {
	// list of fields
	res = []string{"Name", "Value"}
	return
}

func (datapoint *DataPoint) GetFieldStringValue(fieldName string) (res string) {
	switch fieldName {
	// string value of fields
	case "Name":
		res = datapoint.Name
	case "Value":
		res = fmt.Sprintf("%f", datapoint.Value)
	}
	return
}

func (stage *StageStruct) getDatasetOrderedStructWithNameField() []*Dataset {
	// have alphabetical order generation
	datasetOrdered := []*Dataset{}
	for dataset := range stage.Datasets {
		datasetOrdered = append(datasetOrdered, dataset)
	}
	sort.Slice(datasetOrdered[:], func(i, j int) bool {
		return datasetOrdered[i].Name < datasetOrdered[j].Name
	})
	return datasetOrdered
}

// Stage puts dataset to the model stage
func (dataset *Dataset) Stage() *Dataset {
	Stage.Datasets[dataset] = __member
	Stage.Datasets_mapString[dataset.Name] = dataset

	return dataset
}

// Unstage removes dataset off the model stage
func (dataset *Dataset) Unstage() *Dataset {
	delete(Stage.Datasets, dataset)
	delete(Stage.Datasets_mapString, dataset.Name)
	return dataset
}

// commit dataset to the back repo (if it is already staged)
func (dataset *Dataset) Commit() *Dataset {
	if _, ok := Stage.Datasets[dataset]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitDataset(dataset)
		}
	}
	return dataset
}

// Checkout dataset to the back repo (if it is already staged)
func (dataset *Dataset) Checkout() *Dataset {
	if _, ok := Stage.Datasets[dataset]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutDataset(dataset)
		}
	}
	return dataset
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of dataset to the model stage
func (dataset *Dataset) StageCopy() *Dataset {
	_dataset := new(Dataset)
	*_dataset = *dataset
	_dataset.Stage()
	return _dataset
}

// StageAndCommit appends dataset to the model stage and commit to the orm repo
func (dataset *Dataset) StageAndCommit() *Dataset {
	dataset.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMDataset(dataset)
	}
	return dataset
}

// DeleteStageAndCommit appends dataset to the model stage and commit to the orm repo
func (dataset *Dataset) DeleteStageAndCommit() *Dataset {
	dataset.Unstage()
	DeleteORMDataset(dataset)
	return dataset
}

// StageCopyAndCommit appends a copy of dataset to the model stage and commit to the orm repo
func (dataset *Dataset) StageCopyAndCommit() *Dataset {
	_dataset := new(Dataset)
	*_dataset = *dataset
	_dataset.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMDataset(dataset)
	}
	return _dataset
}

// CreateORMDataset enables dynamic staging of a Dataset instance
func CreateORMDataset(dataset *Dataset) {
	dataset.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMDataset(dataset)
	}
}

// DeleteORMDataset enables dynamic staging of a Dataset instance
func DeleteORMDataset(dataset *Dataset) {
	dataset.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMDataset(dataset)
	}
}

// for satisfaction of GongStruct interface
func (dataset *Dataset) GetName() (res string) {
	return dataset.Name
}

func (dataset *Dataset) GetFields() (res []string) {
	// list of fields
	res = []string{"Name", "DataPoints", "Label"}
	return
}

func (dataset *Dataset) GetFieldStringValue(fieldName string) (res string) {
	switch fieldName {
	// string value of fields
	case "Name":
		res = dataset.Name
	case "DataPoints":
		for idx, __instance__ := range dataset.DataPoints {
			if idx > 0 {
				res += "\n"
			}
			res += __instance__.Name
		}
	case "Label":
		res = dataset.Label
	}
	return
}

func (stage *StageStruct) getLabelOrderedStructWithNameField() []*Label {
	// have alphabetical order generation
	labelOrdered := []*Label{}
	for label := range stage.Labels {
		labelOrdered = append(labelOrdered, label)
	}
	sort.Slice(labelOrdered[:], func(i, j int) bool {
		return labelOrdered[i].Name < labelOrdered[j].Name
	})
	return labelOrdered
}

// Stage puts label to the model stage
func (label *Label) Stage() *Label {
	Stage.Labels[label] = __member
	Stage.Labels_mapString[label.Name] = label

	return label
}

// Unstage removes label off the model stage
func (label *Label) Unstage() *Label {
	delete(Stage.Labels, label)
	delete(Stage.Labels_mapString, label.Name)
	return label
}

// commit label to the back repo (if it is already staged)
func (label *Label) Commit() *Label {
	if _, ok := Stage.Labels[label]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitLabel(label)
		}
	}
	return label
}

// Checkout label to the back repo (if it is already staged)
func (label *Label) Checkout() *Label {
	if _, ok := Stage.Labels[label]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutLabel(label)
		}
	}
	return label
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of label to the model stage
func (label *Label) StageCopy() *Label {
	_label := new(Label)
	*_label = *label
	_label.Stage()
	return _label
}

// StageAndCommit appends label to the model stage and commit to the orm repo
func (label *Label) StageAndCommit() *Label {
	label.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMLabel(label)
	}
	return label
}

// DeleteStageAndCommit appends label to the model stage and commit to the orm repo
func (label *Label) DeleteStageAndCommit() *Label {
	label.Unstage()
	DeleteORMLabel(label)
	return label
}

// StageCopyAndCommit appends a copy of label to the model stage and commit to the orm repo
func (label *Label) StageCopyAndCommit() *Label {
	_label := new(Label)
	*_label = *label
	_label.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMLabel(label)
	}
	return _label
}

// CreateORMLabel enables dynamic staging of a Label instance
func CreateORMLabel(label *Label) {
	label.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMLabel(label)
	}
}

// DeleteORMLabel enables dynamic staging of a Label instance
func DeleteORMLabel(label *Label) {
	label.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMLabel(label)
	}
}

// for satisfaction of GongStruct interface
func (label *Label) GetName() (res string) {
	return label.Name
}

func (label *Label) GetFields() (res []string) {
	// list of fields
	res = []string{"Name"}
	return
}

func (label *Label) GetFieldStringValue(fieldName string) (res string) {
	switch fieldName {
	// string value of fields
	case "Name":
		res = label.Name
	}
	return
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMChartConfiguration(ChartConfiguration *ChartConfiguration)
	CreateORMDataPoint(DataPoint *DataPoint)
	CreateORMDataset(Dataset *Dataset)
	CreateORMLabel(Label *Label)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMChartConfiguration(ChartConfiguration *ChartConfiguration)
	DeleteORMDataPoint(DataPoint *DataPoint)
	DeleteORMDataset(Dataset *Dataset)
	DeleteORMLabel(Label *Label)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.ChartConfigurations = make(map[*ChartConfiguration]struct{})
	stage.ChartConfigurations_mapString = make(map[string]*ChartConfiguration)

	stage.DataPoints = make(map[*DataPoint]struct{})
	stage.DataPoints_mapString = make(map[string]*DataPoint)

	stage.Datasets = make(map[*Dataset]struct{})
	stage.Datasets_mapString = make(map[string]*Dataset)

	stage.Labels = make(map[*Label]struct{})
	stage.Labels_mapString = make(map[string]*Label)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.ChartConfigurations = nil
	stage.ChartConfigurations_mapString = nil

	stage.DataPoints = nil
	stage.DataPoints_mapString = nil

	stage.Datasets = nil
	stage.Datasets_mapString = nil

	stage.Labels = nil
	stage.Labels_mapString = nil

}

const marshallRes = `package {{PackageName}}

import (
	"time"

	"{{ModelsPackageName}}"
)

func init() {
	var __Dummy_time_variable time.Time
	_ = __Dummy_time_variable
	InjectionGateway["{{databaseName}}"] = {{databaseName}}Injection
}

// {{databaseName}}Injection will stage objects of database "{{databaseName}}"
func {{databaseName}}Injection() {

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}

`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: "{{GeneratedFieldNameValue}}"}).Stage()`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

const StringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const NumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const PointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const SliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const TimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *StageStruct) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Println("filename of marshall output  is " + name)

	res := marshallRes
	res = strings.ReplaceAll(res, "{{databaseName}}", strings.ReplaceAll(path.Base(name), ".go", ""))
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	id := ""
	decl := ""
	setValueField := ""

	// insertion initialization of objects to stage
	map_ChartConfiguration_Identifiers := make(map[*ChartConfiguration]string)
	_ = map_ChartConfiguration_Identifiers

	chartconfigurationOrdered := []*ChartConfiguration{}
	for chartconfiguration := range stage.ChartConfigurations {
		chartconfigurationOrdered = append(chartconfigurationOrdered, chartconfiguration)
	}
	sort.Slice(chartconfigurationOrdered[:], func(i, j int) bool {
		return chartconfigurationOrdered[i].Name < chartconfigurationOrdered[j].Name
	})
	identifiersDecl += fmt.Sprintf("\n\n	// Declarations of staged instances of ChartConfiguration")
	for idx, chartconfiguration := range chartconfigurationOrdered {

		id = generatesIdentifier("ChartConfiguration", idx, chartconfiguration.Name)
		map_ChartConfiguration_Identifiers[chartconfiguration] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ChartConfiguration")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", chartconfiguration.Name)
		identifiersDecl += decl

		initializerStatements += fmt.Sprintf("\n\n	// ChartConfiguration %s values setup", chartconfiguration.Name)
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(chartconfiguration.Name))
		initializerStatements += setValueField

		if chartconfiguration.ChartType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ChartType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+chartconfiguration.ChartType.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", chartconfiguration.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Heigth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", chartconfiguration.Heigth))
		initializerStatements += setValueField

	}

	map_DataPoint_Identifiers := make(map[*DataPoint]string)
	_ = map_DataPoint_Identifiers

	datapointOrdered := []*DataPoint{}
	for datapoint := range stage.DataPoints {
		datapointOrdered = append(datapointOrdered, datapoint)
	}
	sort.Slice(datapointOrdered[:], func(i, j int) bool {
		return datapointOrdered[i].Name < datapointOrdered[j].Name
	})
	identifiersDecl += fmt.Sprintf("\n\n	// Declarations of staged instances of DataPoint")
	for idx, datapoint := range datapointOrdered {

		id = generatesIdentifier("DataPoint", idx, datapoint.Name)
		map_DataPoint_Identifiers[datapoint] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DataPoint")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datapoint.Name)
		identifiersDecl += decl

		initializerStatements += fmt.Sprintf("\n\n	// DataPoint %s values setup", datapoint.Name)
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datapoint.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", datapoint.Value))
		initializerStatements += setValueField

	}

	map_Dataset_Identifiers := make(map[*Dataset]string)
	_ = map_Dataset_Identifiers

	datasetOrdered := []*Dataset{}
	for dataset := range stage.Datasets {
		datasetOrdered = append(datasetOrdered, dataset)
	}
	sort.Slice(datasetOrdered[:], func(i, j int) bool {
		return datasetOrdered[i].Name < datasetOrdered[j].Name
	})
	identifiersDecl += fmt.Sprintf("\n\n	// Declarations of staged instances of Dataset")
	for idx, dataset := range datasetOrdered {

		id = generatesIdentifier("Dataset", idx, dataset.Name)
		map_Dataset_Identifiers[dataset] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Dataset")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", dataset.Name)
		identifiersDecl += decl

		initializerStatements += fmt.Sprintf("\n\n	// Dataset %s values setup", dataset.Name)
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(dataset.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Label")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(dataset.Label))
		initializerStatements += setValueField

	}

	map_Label_Identifiers := make(map[*Label]string)
	_ = map_Label_Identifiers

	labelOrdered := []*Label{}
	for label := range stage.Labels {
		labelOrdered = append(labelOrdered, label)
	}
	sort.Slice(labelOrdered[:], func(i, j int) bool {
		return labelOrdered[i].Name < labelOrdered[j].Name
	})
	identifiersDecl += fmt.Sprintf("\n\n	// Declarations of staged instances of Label")
	for idx, label := range labelOrdered {

		id = generatesIdentifier("Label", idx, label.Name)
		map_Label_Identifiers[label] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Label")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", label.Name)
		identifiersDecl += decl

		initializerStatements += fmt.Sprintf("\n\n	// Label %s values setup", label.Name)
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(label.Name))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	for idx, chartconfiguration := range chartconfigurationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ChartConfiguration", idx, chartconfiguration.Name)
		map_ChartConfiguration_Identifiers[chartconfiguration] = id

		// Initialisation of values
		for _, _dataset := range chartconfiguration.Datasets {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Datasets")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Dataset_Identifiers[_dataset])
			pointersInitializesStatements += setPointerField
		}

		for _, _label := range chartconfiguration.Labels {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Labels")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Label_Identifiers[_label])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, datapoint := range datapointOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DataPoint", idx, datapoint.Name)
		map_DataPoint_Identifiers[datapoint] = id

		// Initialisation of values
	}

	for idx, dataset := range datasetOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Dataset", idx, dataset.Name)
		map_Dataset_Identifiers[dataset] = id

		// Initialisation of values
		for _, _datapoint := range dataset.DataPoints {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DataPoints")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DataPoint_Identifiers[_datapoint])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, label := range labelOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Label", idx, label.Name)
		map_Label_Identifiers[label] = id

		// Initialisation of values
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	fmt.Fprintln(file, res)
}

// unique identifier per struct
func generatesIdentifier(gongStructName string, idx int, instanceName string) (identifier string) {

	identifier = instanceName
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(instanceName, "_")

	identifier = fmt.Sprintf("__%s__%06d_%s", gongStructName, idx, processedString)

	return
}

// insertion point of functions that provide maps for reverse associations
// generate function for reverse association maps of ChartConfiguration
func (stageStruct *StageStruct) CreateReverseMap_ChartConfiguration_Datasets() (res map[*Dataset]*ChartConfiguration) {
	res = make(map[*Dataset]*ChartConfiguration)

	for chartconfiguration := range stageStruct.ChartConfigurations {
		for _, dataset_ := range chartconfiguration.Datasets {
			res[dataset_] = chartconfiguration
		}
	}

	return
}

func (stageStruct *StageStruct) CreateReverseMap_ChartConfiguration_Labels() (res map[*Label]*ChartConfiguration) {
	res = make(map[*Label]*ChartConfiguration)

	for chartconfiguration := range stageStruct.ChartConfigurations {
		for _, label_ := range chartconfiguration.Labels {
			res[label_] = chartconfiguration
		}
	}

	return
}

// generate function for reverse association maps of DataPoint
// generate function for reverse association maps of Dataset
func (stageStruct *StageStruct) CreateReverseMap_Dataset_DataPoints() (res map[*DataPoint]*Dataset) {
	res = make(map[*DataPoint]*Dataset)

	for dataset := range stageStruct.Datasets {
		for _, datapoint_ := range dataset.DataPoints {
			res[datapoint_] = dataset
		}
	}

	return
}

// generate function for reverse association maps of Label

// insertion point of enum utility functions
// Utility function for ChartType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (charttype ChartType) ToString() (res string) {

	// migration of former implementation of enum
	switch charttype {
	// insertion code per enum code
	case BAR:
		res = "bar"
	case DOUGHNUT:
		res = "doughnut"
	case LINE:
		res = "line"
	case PIE:
		res = "pie"
	case POLAR_AREA:
		res = "polarArea"
	case RADAR:
		res = "radar"
	}
	return
}

func (charttype *ChartType) FromString(input string) {

	switch input {
	// insertion code per enum code
	case "bar":
		*charttype = BAR
	case "doughnut":
		*charttype = DOUGHNUT
	case "line":
		*charttype = LINE
	case "pie":
		*charttype = PIE
	case "polarArea":
		*charttype = POLAR_AREA
	case "radar":
		*charttype = RADAR
	}
}

func (charttype *ChartType) ToCodeString() (res string) {

	switch *charttype {
	// insertion code per enum code
	case BAR:
		res = "BAR"
	case DOUGHNUT:
		res = "DOUGHNUT"
	case LINE:
		res = "LINE"
	case PIE:
		res = "PIE"
	case POLAR_AREA:
		res = "POLAR_AREA"
	case RADAR:
		res = "RADAR"
	}
	return
}

