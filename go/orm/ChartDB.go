// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gongng2charts/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Chart_sql sql.NullBool
var dummy_Chart_time time.Duration
var dummy_Chart_sort sort.Float64Slice

// ChartAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model chartAPI
type ChartAPI struct {
	gorm.Model

	models.Chart

	// encoding of pointers
	ChartPointersEnconding
}

// ChartPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type ChartPointersEnconding struct {
	// insertion for pointer fields encoding declaration
}

// ChartDB describes a chart in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model chartDB
type ChartDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field chartDB.Name {{BasicKind}} (to be completed)
	Name_Data sql.NullString
	// encoding of pointers
	ChartPointersEnconding
}

// ChartDBs arrays chartDBs
// swagger:response chartDBsResponse
type ChartDBs []ChartDB

// ChartDBResponse provides response
// swagger:response chartDBResponse
type ChartDBResponse struct {
	ChartDB
}

// ChartWOP is a Chart without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type ChartWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Chart_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoChartStruct struct {
	// stores ChartDB according to their gorm ID
	Map_ChartDBID_ChartDB *map[uint]*ChartDB

	// stores ChartDB ID according to Chart address
	Map_ChartPtr_ChartDBID *map[*models.Chart]uint

	// stores Chart according to their gorm ID
	Map_ChartDBID_ChartPtr *map[uint]*models.Chart

	db *gorm.DB
}

func (backRepoChart *BackRepoChartStruct) GetDB() *gorm.DB {
	return backRepoChart.db
}

// GetChartDBFromChartPtr is a handy function to access the back repo instance from the stage instance
func (backRepoChart *BackRepoChartStruct) GetChartDBFromChartPtr(chart *models.Chart) (chartDB *ChartDB) {
	id := (*backRepoChart.Map_ChartPtr_ChartDBID)[chart]
	chartDB = (*backRepoChart.Map_ChartDBID_ChartDB)[id]
	return
}

// BackRepoChart.Init set up the BackRepo of the Chart
func (backRepoChart *BackRepoChartStruct) Init(db *gorm.DB) (Error error) {

	if backRepoChart.Map_ChartDBID_ChartPtr != nil {
		err := errors.New("In Init, backRepoChart.Map_ChartDBID_ChartPtr should be nil")
		return err
	}

	if backRepoChart.Map_ChartDBID_ChartDB != nil {
		err := errors.New("In Init, backRepoChart.Map_ChartDBID_ChartDB should be nil")
		return err
	}

	if backRepoChart.Map_ChartPtr_ChartDBID != nil {
		err := errors.New("In Init, backRepoChart.Map_ChartPtr_ChartDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.Chart, 0)
	backRepoChart.Map_ChartDBID_ChartPtr = &tmp

	tmpDB := make(map[uint]*ChartDB, 0)
	backRepoChart.Map_ChartDBID_ChartDB = &tmpDB

	tmpID := make(map[*models.Chart]uint, 0)
	backRepoChart.Map_ChartPtr_ChartDBID = &tmpID

	backRepoChart.db = db
	return
}

// BackRepoChart.CommitPhaseOne commits all staged instances of Chart to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoChart *BackRepoChartStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for chart := range stage.Charts {
		backRepoChart.CommitPhaseOneInstance(chart)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, chart := range *backRepoChart.Map_ChartDBID_ChartPtr {
		if _, ok := stage.Charts[chart]; !ok {
			backRepoChart.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoChart.CommitDeleteInstance commits deletion of Chart to the BackRepo
func (backRepoChart *BackRepoChartStruct) CommitDeleteInstance(id uint) (Error error) {

	chart := (*backRepoChart.Map_ChartDBID_ChartPtr)[id]

	// chart is not staged anymore, remove chartDB
	chartDB := (*backRepoChart.Map_ChartDBID_ChartDB)[id]
	query := backRepoChart.db.Unscoped().Delete(&chartDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoChart.Map_ChartPtr_ChartDBID), chart)
	delete((*backRepoChart.Map_ChartDBID_ChartPtr), id)
	delete((*backRepoChart.Map_ChartDBID_ChartDB), id)

	return
}

// BackRepoChart.CommitPhaseOneInstance commits chart staged instances of Chart to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoChart *BackRepoChartStruct) CommitPhaseOneInstance(chart *models.Chart) (Error error) {

	// check if the chart is not commited yet
	if _, ok := (*backRepoChart.Map_ChartPtr_ChartDBID)[chart]; ok {
		return
	}

	// initiate chart
	var chartDB ChartDB
	chartDB.CopyBasicFieldsFromChart(chart)

	query := backRepoChart.db.Create(&chartDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoChart.Map_ChartPtr_ChartDBID)[chart] = chartDB.ID
	(*backRepoChart.Map_ChartDBID_ChartPtr)[chartDB.ID] = chart
	(*backRepoChart.Map_ChartDBID_ChartDB)[chartDB.ID] = &chartDB

	return
}

// BackRepoChart.CommitPhaseTwo commits all staged instances of Chart to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoChart *BackRepoChartStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, chart := range *backRepoChart.Map_ChartDBID_ChartPtr {
		backRepoChart.CommitPhaseTwoInstance(backRepo, idx, chart)
	}

	return
}

// BackRepoChart.CommitPhaseTwoInstance commits {{structname }} of models.Chart to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoChart *BackRepoChartStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, chart *models.Chart) (Error error) {

	// fetch matching chartDB
	if chartDB, ok := (*backRepoChart.Map_ChartDBID_ChartDB)[idx]; ok {

		chartDB.CopyBasicFieldsFromChart(chart)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoChart.db.Save(&chartDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Chart intance %s", chart.Name))
		return err
	}

	return
}

// BackRepoChart.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoChart *BackRepoChartStruct) CheckoutPhaseOne() (Error error) {

	chartDBArray := make([]ChartDB, 0)
	query := backRepoChart.db.Find(&chartDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	chartInstancesToBeRemovedFromTheStage := make(map[*models.Chart]struct{})
	for key, value := range models.Stage.Charts {
		chartInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, chartDB := range chartDBArray {
		backRepoChart.CheckoutPhaseOneInstance(&chartDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		chart, ok := (*backRepoChart.Map_ChartDBID_ChartPtr)[chartDB.ID]
		if ok {
			delete(chartInstancesToBeRemovedFromTheStage, chart)
		}
	}

	// remove from stage and back repo's 3 maps all charts that are not in the checkout
	for chart := range chartInstancesToBeRemovedFromTheStage {
		chart.Unstage()

		// remove instance from the back repo 3 maps
		chartID := (*backRepoChart.Map_ChartPtr_ChartDBID)[chart]
		delete((*backRepoChart.Map_ChartPtr_ChartDBID), chart)
		delete((*backRepoChart.Map_ChartDBID_ChartDB), chartID)
		delete((*backRepoChart.Map_ChartDBID_ChartPtr), chartID)
	}

	return
}

// CheckoutPhaseOneInstance takes a chartDB that has been found in the DB, updates the backRepo and stages the
// models version of the chartDB
func (backRepoChart *BackRepoChartStruct) CheckoutPhaseOneInstance(chartDB *ChartDB) (Error error) {

	chart, ok := (*backRepoChart.Map_ChartDBID_ChartPtr)[chartDB.ID]
	if !ok {
		chart = new(models.Chart)

		(*backRepoChart.Map_ChartDBID_ChartPtr)[chartDB.ID] = chart
		(*backRepoChart.Map_ChartPtr_ChartDBID)[chart] = chartDB.ID

		// append model store with the new element
		chart.Name = chartDB.Name_Data.String
		chart.Stage()
	}
	chartDB.CopyBasicFieldsToChart(chart)

	// preserve pointer to chartDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_ChartDBID_ChartDB)[chartDB hold variable pointers
	chartDB_Data := *chartDB
	preservedPtrToChart := &chartDB_Data
	(*backRepoChart.Map_ChartDBID_ChartDB)[chartDB.ID] = preservedPtrToChart

	return
}

// BackRepoChart.CheckoutPhaseTwo Checkouts all staged instances of Chart to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoChart *BackRepoChartStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, chartDB := range *backRepoChart.Map_ChartDBID_ChartDB {
		backRepoChart.CheckoutPhaseTwoInstance(backRepo, chartDB)
	}
	return
}

// BackRepoChart.CheckoutPhaseTwoInstance Checkouts staged instances of Chart to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoChart *BackRepoChartStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, chartDB *ChartDB) (Error error) {

	chart := (*backRepoChart.Map_ChartDBID_ChartPtr)[chartDB.ID]
	_ = chart // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitChart allows commit of a single chart (if already staged)
func (backRepo *BackRepoStruct) CommitChart(chart *models.Chart) {
	backRepo.BackRepoChart.CommitPhaseOneInstance(chart)
	if id, ok := (*backRepo.BackRepoChart.Map_ChartPtr_ChartDBID)[chart]; ok {
		backRepo.BackRepoChart.CommitPhaseTwoInstance(backRepo, id, chart)
	}
}

// CommitChart allows checkout of a single chart (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutChart(chart *models.Chart) {
	// check if the chart is staged
	if _, ok := (*backRepo.BackRepoChart.Map_ChartPtr_ChartDBID)[chart]; ok {

		if id, ok := (*backRepo.BackRepoChart.Map_ChartPtr_ChartDBID)[chart]; ok {
			var chartDB ChartDB
			chartDB.ID = id

			if err := backRepo.BackRepoChart.db.First(&chartDB, id).Error; err != nil {
				log.Panicln("CheckoutChart : Problem with getting object with id:", id)
			}
			backRepo.BackRepoChart.CheckoutPhaseOneInstance(&chartDB)
			backRepo.BackRepoChart.CheckoutPhaseTwoInstance(backRepo, &chartDB)
		}
	}
}

// CopyBasicFieldsFromChart
func (chartDB *ChartDB) CopyBasicFieldsFromChart(chart *models.Chart) {
	// insertion point for fields commit

	chartDB.Name_Data.String = chart.Name
	chartDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromChartWOP
func (chartDB *ChartDB) CopyBasicFieldsFromChartWOP(chart *ChartWOP) {
	// insertion point for fields commit

	chartDB.Name_Data.String = chart.Name
	chartDB.Name_Data.Valid = true
}

// CopyBasicFieldsToChart
func (chartDB *ChartDB) CopyBasicFieldsToChart(chart *models.Chart) {
	// insertion point for checkout of basic fields (back repo to stage)
	chart.Name = chartDB.Name_Data.String
}

// CopyBasicFieldsToChartWOP
func (chartDB *ChartDB) CopyBasicFieldsToChartWOP(chart *ChartWOP) {
	chart.ID = int(chartDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	chart.Name = chartDB.Name_Data.String
}

// Backup generates a json file from a slice of all ChartDB instances in the backrepo
func (backRepoChart *BackRepoChartStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "ChartDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ChartDB, 0)
	for _, chartDB := range *backRepoChart.Map_ChartDBID_ChartDB {
		forBackup = append(forBackup, chartDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Chart ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Chart file", err.Error())
	}
}

// Backup generates a json file from a slice of all ChartDB instances in the backrepo
func (backRepoChart *BackRepoChartStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ChartDB, 0)
	for _, chartDB := range *backRepoChart.Map_ChartDBID_ChartDB {
		forBackup = append(forBackup, chartDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Chart")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Chart_Fields, -1)
	for _, chartDB := range forBackup {

		var chartWOP ChartWOP
		chartDB.CopyBasicFieldsToChartWOP(&chartWOP)

		row := sh.AddRow()
		row.WriteStruct(&chartWOP, -1)
	}
}

// RestoreXL from the "Chart" sheet all ChartDB instances
func (backRepoChart *BackRepoChartStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoChartid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Chart"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoChart.rowVisitorChart)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoChart *BackRepoChartStruct) rowVisitorChart(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var chartWOP ChartWOP
		row.ReadStruct(&chartWOP)

		// add the unmarshalled struct to the stage
		chartDB := new(ChartDB)
		chartDB.CopyBasicFieldsFromChartWOP(&chartWOP)

		chartDB_ID_atBackupTime := chartDB.ID
		chartDB.ID = 0
		query := backRepoChart.db.Create(chartDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoChart.Map_ChartDBID_ChartDB)[chartDB.ID] = chartDB
		BackRepoChartid_atBckpTime_newID[chartDB_ID_atBackupTime] = chartDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "ChartDB.json" in dirPath that stores an array
// of ChartDB and stores it in the database
// the map BackRepoChartid_atBckpTime_newID is updated accordingly
func (backRepoChart *BackRepoChartStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoChartid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "ChartDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Chart file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*ChartDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_ChartDBID_ChartDB
	for _, chartDB := range forRestore {

		chartDB_ID_atBackupTime := chartDB.ID
		chartDB.ID = 0
		query := backRepoChart.db.Create(chartDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoChart.Map_ChartDBID_ChartDB)[chartDB.ID] = chartDB
		BackRepoChartid_atBckpTime_newID[chartDB_ID_atBackupTime] = chartDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Chart file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Chart>id_atBckpTime_newID
// to compute new index
func (backRepoChart *BackRepoChartStruct) RestorePhaseTwo() {

	for _, chartDB := range *backRepoChart.Map_ChartDBID_ChartDB {

		// next line of code is to avert unused variable compilation error
		_ = chartDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoChart.db.Model(chartDB).Updates(*chartDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoChartid_atBckpTime_newID map[uint]uint