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
var dummy_Dataset_sql sql.NullBool
var dummy_Dataset_time time.Duration
var dummy_Dataset_sort sort.Float64Slice

// DatasetAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model datasetAPI
type DatasetAPI struct {
	gorm.Model

	models.Dataset

	// encoding of pointers
	DatasetPointersEnconding
}

// DatasetPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type DatasetPointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// Implementation of a reverse ID for field ChartConfiguration{}.Datasets []*Dataset
	ChartConfiguration_DatasetsDBID sql.NullInt64

	// implementation of the index of the withing the slice
	ChartConfiguration_DatasetsDBID_Index sql.NullInt64
}

// DatasetDB describes a dataset in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model datasetDB
type DatasetDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field datasetDB.Name {{BasicKind}} (to be completed)
	Name_Data sql.NullString
	// encoding of pointers
	DatasetPointersEnconding
}

// DatasetDBs arrays datasetDBs
// swagger:response datasetDBsResponse
type DatasetDBs []DatasetDB

// DatasetDBResponse provides response
// swagger:response datasetDBResponse
type DatasetDBResponse struct {
	DatasetDB
}

// DatasetWOP is a Dataset without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type DatasetWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Dataset_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoDatasetStruct struct {
	// stores DatasetDB according to their gorm ID
	Map_DatasetDBID_DatasetDB *map[uint]*DatasetDB

	// stores DatasetDB ID according to Dataset address
	Map_DatasetPtr_DatasetDBID *map[*models.Dataset]uint

	// stores Dataset according to their gorm ID
	Map_DatasetDBID_DatasetPtr *map[uint]*models.Dataset

	db *gorm.DB
}

func (backRepoDataset *BackRepoDatasetStruct) GetDB() *gorm.DB {
	return backRepoDataset.db
}

// GetDatasetDBFromDatasetPtr is a handy function to access the back repo instance from the stage instance
func (backRepoDataset *BackRepoDatasetStruct) GetDatasetDBFromDatasetPtr(dataset *models.Dataset) (datasetDB *DatasetDB) {
	id := (*backRepoDataset.Map_DatasetPtr_DatasetDBID)[dataset]
	datasetDB = (*backRepoDataset.Map_DatasetDBID_DatasetDB)[id]
	return
}

// BackRepoDataset.Init set up the BackRepo of the Dataset
func (backRepoDataset *BackRepoDatasetStruct) Init(db *gorm.DB) (Error error) {

	if backRepoDataset.Map_DatasetDBID_DatasetPtr != nil {
		err := errors.New("In Init, backRepoDataset.Map_DatasetDBID_DatasetPtr should be nil")
		return err
	}

	if backRepoDataset.Map_DatasetDBID_DatasetDB != nil {
		err := errors.New("In Init, backRepoDataset.Map_DatasetDBID_DatasetDB should be nil")
		return err
	}

	if backRepoDataset.Map_DatasetPtr_DatasetDBID != nil {
		err := errors.New("In Init, backRepoDataset.Map_DatasetPtr_DatasetDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.Dataset, 0)
	backRepoDataset.Map_DatasetDBID_DatasetPtr = &tmp

	tmpDB := make(map[uint]*DatasetDB, 0)
	backRepoDataset.Map_DatasetDBID_DatasetDB = &tmpDB

	tmpID := make(map[*models.Dataset]uint, 0)
	backRepoDataset.Map_DatasetPtr_DatasetDBID = &tmpID

	backRepoDataset.db = db
	return
}

// BackRepoDataset.CommitPhaseOne commits all staged instances of Dataset to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoDataset *BackRepoDatasetStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for dataset := range stage.Datasets {
		backRepoDataset.CommitPhaseOneInstance(dataset)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, dataset := range *backRepoDataset.Map_DatasetDBID_DatasetPtr {
		if _, ok := stage.Datasets[dataset]; !ok {
			backRepoDataset.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoDataset.CommitDeleteInstance commits deletion of Dataset to the BackRepo
func (backRepoDataset *BackRepoDatasetStruct) CommitDeleteInstance(id uint) (Error error) {

	dataset := (*backRepoDataset.Map_DatasetDBID_DatasetPtr)[id]

	// dataset is not staged anymore, remove datasetDB
	datasetDB := (*backRepoDataset.Map_DatasetDBID_DatasetDB)[id]
	query := backRepoDataset.db.Unscoped().Delete(&datasetDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoDataset.Map_DatasetPtr_DatasetDBID), dataset)
	delete((*backRepoDataset.Map_DatasetDBID_DatasetPtr), id)
	delete((*backRepoDataset.Map_DatasetDBID_DatasetDB), id)

	return
}

// BackRepoDataset.CommitPhaseOneInstance commits dataset staged instances of Dataset to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoDataset *BackRepoDatasetStruct) CommitPhaseOneInstance(dataset *models.Dataset) (Error error) {

	// check if the dataset is not commited yet
	if _, ok := (*backRepoDataset.Map_DatasetPtr_DatasetDBID)[dataset]; ok {
		return
	}

	// initiate dataset
	var datasetDB DatasetDB
	datasetDB.CopyBasicFieldsFromDataset(dataset)

	query := backRepoDataset.db.Create(&datasetDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoDataset.Map_DatasetPtr_DatasetDBID)[dataset] = datasetDB.ID
	(*backRepoDataset.Map_DatasetDBID_DatasetPtr)[datasetDB.ID] = dataset
	(*backRepoDataset.Map_DatasetDBID_DatasetDB)[datasetDB.ID] = &datasetDB

	return
}

// BackRepoDataset.CommitPhaseTwo commits all staged instances of Dataset to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDataset *BackRepoDatasetStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, dataset := range *backRepoDataset.Map_DatasetDBID_DatasetPtr {
		backRepoDataset.CommitPhaseTwoInstance(backRepo, idx, dataset)
	}

	return
}

// BackRepoDataset.CommitPhaseTwoInstance commits {{structname }} of models.Dataset to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDataset *BackRepoDatasetStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, dataset *models.Dataset) (Error error) {

	// fetch matching datasetDB
	if datasetDB, ok := (*backRepoDataset.Map_DatasetDBID_DatasetDB)[idx]; ok {

		datasetDB.CopyBasicFieldsFromDataset(dataset)

		// insertion point for translating pointers encodings into actual pointers
		// This loop encodes the slice of pointers dataset.DataPoints into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, datapointAssocEnd := range dataset.DataPoints {

			// get the back repo instance at the association end
			datapointAssocEnd_DB :=
				backRepo.BackRepoDataPoint.GetDataPointDBFromDataPointPtr(datapointAssocEnd)

			// encode reverse pointer in the association end back repo instance
			datapointAssocEnd_DB.Dataset_DataPointsDBID.Int64 = int64(datasetDB.ID)
			datapointAssocEnd_DB.Dataset_DataPointsDBID.Valid = true
			datapointAssocEnd_DB.Dataset_DataPointsDBID_Index.Int64 = int64(idx)
			datapointAssocEnd_DB.Dataset_DataPointsDBID_Index.Valid = true
			if q := backRepoDataset.db.Save(datapointAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		query := backRepoDataset.db.Save(&datasetDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Dataset intance %s", dataset.Name))
		return err
	}

	return
}

// BackRepoDataset.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoDataset *BackRepoDatasetStruct) CheckoutPhaseOne() (Error error) {

	datasetDBArray := make([]DatasetDB, 0)
	query := backRepoDataset.db.Find(&datasetDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	datasetInstancesToBeRemovedFromTheStage := make(map[*models.Dataset]struct{})
	for key, value := range models.Stage.Datasets {
		datasetInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, datasetDB := range datasetDBArray {
		backRepoDataset.CheckoutPhaseOneInstance(&datasetDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		dataset, ok := (*backRepoDataset.Map_DatasetDBID_DatasetPtr)[datasetDB.ID]
		if ok {
			delete(datasetInstancesToBeRemovedFromTheStage, dataset)
		}
	}

	// remove from stage and back repo's 3 maps all datasets that are not in the checkout
	for dataset := range datasetInstancesToBeRemovedFromTheStage {
		dataset.Unstage()

		// remove instance from the back repo 3 maps
		datasetID := (*backRepoDataset.Map_DatasetPtr_DatasetDBID)[dataset]
		delete((*backRepoDataset.Map_DatasetPtr_DatasetDBID), dataset)
		delete((*backRepoDataset.Map_DatasetDBID_DatasetDB), datasetID)
		delete((*backRepoDataset.Map_DatasetDBID_DatasetPtr), datasetID)
	}

	return
}

// CheckoutPhaseOneInstance takes a datasetDB that has been found in the DB, updates the backRepo and stages the
// models version of the datasetDB
func (backRepoDataset *BackRepoDatasetStruct) CheckoutPhaseOneInstance(datasetDB *DatasetDB) (Error error) {

	dataset, ok := (*backRepoDataset.Map_DatasetDBID_DatasetPtr)[datasetDB.ID]
	if !ok {
		dataset = new(models.Dataset)

		(*backRepoDataset.Map_DatasetDBID_DatasetPtr)[datasetDB.ID] = dataset
		(*backRepoDataset.Map_DatasetPtr_DatasetDBID)[dataset] = datasetDB.ID

		// append model store with the new element
		dataset.Name = datasetDB.Name_Data.String
		dataset.Stage()
	}
	datasetDB.CopyBasicFieldsToDataset(dataset)

	// preserve pointer to datasetDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_DatasetDBID_DatasetDB)[datasetDB hold variable pointers
	datasetDB_Data := *datasetDB
	preservedPtrToDataset := &datasetDB_Data
	(*backRepoDataset.Map_DatasetDBID_DatasetDB)[datasetDB.ID] = preservedPtrToDataset

	return
}

// BackRepoDataset.CheckoutPhaseTwo Checkouts all staged instances of Dataset to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDataset *BackRepoDatasetStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, datasetDB := range *backRepoDataset.Map_DatasetDBID_DatasetDB {
		backRepoDataset.CheckoutPhaseTwoInstance(backRepo, datasetDB)
	}
	return
}

// BackRepoDataset.CheckoutPhaseTwoInstance Checkouts staged instances of Dataset to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDataset *BackRepoDatasetStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, datasetDB *DatasetDB) (Error error) {

	dataset := (*backRepoDataset.Map_DatasetDBID_DatasetPtr)[datasetDB.ID]
	_ = dataset // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// This loop redeem dataset.DataPoints in the stage from the encode in the back repo
	// It parses all DataPointDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	dataset.DataPoints = dataset.DataPoints[:0]
	// 2. loop all instances in the type in the association end
	for _, datapointDB_AssocEnd := range *backRepo.BackRepoDataPoint.Map_DataPointDBID_DataPointDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if datapointDB_AssocEnd.Dataset_DataPointsDBID.Int64 == int64(datasetDB.ID) {
			// 4. fetch the associated instance in the stage
			datapoint_AssocEnd := (*backRepo.BackRepoDataPoint.Map_DataPointDBID_DataPointPtr)[datapointDB_AssocEnd.ID]
			// 5. append it the association slice
			dataset.DataPoints = append(dataset.DataPoints, datapoint_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(dataset.DataPoints, func(i, j int) bool {
		datapointDB_i_ID := (*backRepo.BackRepoDataPoint.Map_DataPointPtr_DataPointDBID)[dataset.DataPoints[i]]
		datapointDB_j_ID := (*backRepo.BackRepoDataPoint.Map_DataPointPtr_DataPointDBID)[dataset.DataPoints[j]]

		datapointDB_i := (*backRepo.BackRepoDataPoint.Map_DataPointDBID_DataPointDB)[datapointDB_i_ID]
		datapointDB_j := (*backRepo.BackRepoDataPoint.Map_DataPointDBID_DataPointDB)[datapointDB_j_ID]

		return datapointDB_i.Dataset_DataPointsDBID_Index.Int64 < datapointDB_j.Dataset_DataPointsDBID_Index.Int64
	})

	return
}

// CommitDataset allows commit of a single dataset (if already staged)
func (backRepo *BackRepoStruct) CommitDataset(dataset *models.Dataset) {
	backRepo.BackRepoDataset.CommitPhaseOneInstance(dataset)
	if id, ok := (*backRepo.BackRepoDataset.Map_DatasetPtr_DatasetDBID)[dataset]; ok {
		backRepo.BackRepoDataset.CommitPhaseTwoInstance(backRepo, id, dataset)
	}
}

// CommitDataset allows checkout of a single dataset (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutDataset(dataset *models.Dataset) {
	// check if the dataset is staged
	if _, ok := (*backRepo.BackRepoDataset.Map_DatasetPtr_DatasetDBID)[dataset]; ok {

		if id, ok := (*backRepo.BackRepoDataset.Map_DatasetPtr_DatasetDBID)[dataset]; ok {
			var datasetDB DatasetDB
			datasetDB.ID = id

			if err := backRepo.BackRepoDataset.db.First(&datasetDB, id).Error; err != nil {
				log.Panicln("CheckoutDataset : Problem with getting object with id:", id)
			}
			backRepo.BackRepoDataset.CheckoutPhaseOneInstance(&datasetDB)
			backRepo.BackRepoDataset.CheckoutPhaseTwoInstance(backRepo, &datasetDB)
		}
	}
}

// CopyBasicFieldsFromDataset
func (datasetDB *DatasetDB) CopyBasicFieldsFromDataset(dataset *models.Dataset) {
	// insertion point for fields commit

	datasetDB.Name_Data.String = dataset.Name
	datasetDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromDatasetWOP
func (datasetDB *DatasetDB) CopyBasicFieldsFromDatasetWOP(dataset *DatasetWOP) {
	// insertion point for fields commit

	datasetDB.Name_Data.String = dataset.Name
	datasetDB.Name_Data.Valid = true
}

// CopyBasicFieldsToDataset
func (datasetDB *DatasetDB) CopyBasicFieldsToDataset(dataset *models.Dataset) {
	// insertion point for checkout of basic fields (back repo to stage)
	dataset.Name = datasetDB.Name_Data.String
}

// CopyBasicFieldsToDatasetWOP
func (datasetDB *DatasetDB) CopyBasicFieldsToDatasetWOP(dataset *DatasetWOP) {
	dataset.ID = int(datasetDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	dataset.Name = datasetDB.Name_Data.String
}

// Backup generates a json file from a slice of all DatasetDB instances in the backrepo
func (backRepoDataset *BackRepoDatasetStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "DatasetDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*DatasetDB, 0)
	for _, datasetDB := range *backRepoDataset.Map_DatasetDBID_DatasetDB {
		forBackup = append(forBackup, datasetDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Dataset ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Dataset file", err.Error())
	}
}

// Backup generates a json file from a slice of all DatasetDB instances in the backrepo
func (backRepoDataset *BackRepoDatasetStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*DatasetDB, 0)
	for _, datasetDB := range *backRepoDataset.Map_DatasetDBID_DatasetDB {
		forBackup = append(forBackup, datasetDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Dataset")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Dataset_Fields, -1)
	for _, datasetDB := range forBackup {

		var datasetWOP DatasetWOP
		datasetDB.CopyBasicFieldsToDatasetWOP(&datasetWOP)

		row := sh.AddRow()
		row.WriteStruct(&datasetWOP, -1)
	}
}

// RestoreXL from the "Dataset" sheet all DatasetDB instances
func (backRepoDataset *BackRepoDatasetStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoDatasetid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Dataset"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoDataset.rowVisitorDataset)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoDataset *BackRepoDatasetStruct) rowVisitorDataset(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var datasetWOP DatasetWOP
		row.ReadStruct(&datasetWOP)

		// add the unmarshalled struct to the stage
		datasetDB := new(DatasetDB)
		datasetDB.CopyBasicFieldsFromDatasetWOP(&datasetWOP)

		datasetDB_ID_atBackupTime := datasetDB.ID
		datasetDB.ID = 0
		query := backRepoDataset.db.Create(datasetDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoDataset.Map_DatasetDBID_DatasetDB)[datasetDB.ID] = datasetDB
		BackRepoDatasetid_atBckpTime_newID[datasetDB_ID_atBackupTime] = datasetDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "DatasetDB.json" in dirPath that stores an array
// of DatasetDB and stores it in the database
// the map BackRepoDatasetid_atBckpTime_newID is updated accordingly
func (backRepoDataset *BackRepoDatasetStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoDatasetid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "DatasetDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Dataset file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*DatasetDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_DatasetDBID_DatasetDB
	for _, datasetDB := range forRestore {

		datasetDB_ID_atBackupTime := datasetDB.ID
		datasetDB.ID = 0
		query := backRepoDataset.db.Create(datasetDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoDataset.Map_DatasetDBID_DatasetDB)[datasetDB.ID] = datasetDB
		BackRepoDatasetid_atBckpTime_newID[datasetDB_ID_atBackupTime] = datasetDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Dataset file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Dataset>id_atBckpTime_newID
// to compute new index
func (backRepoDataset *BackRepoDatasetStruct) RestorePhaseTwo() {

	for _, datasetDB := range *backRepoDataset.Map_DatasetDBID_DatasetDB {

		// next line of code is to avert unused variable compilation error
		_ = datasetDB

		// insertion point for reindexing pointers encoding
		// This reindex dataset.Datasets
		if datasetDB.ChartConfiguration_DatasetsDBID.Int64 != 0 {
			datasetDB.ChartConfiguration_DatasetsDBID.Int64 =
				int64(BackRepoChartConfigurationid_atBckpTime_newID[uint(datasetDB.ChartConfiguration_DatasetsDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoDataset.db.Model(datasetDB).Updates(*datasetDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoDatasetid_atBckpTime_newID map[uint]uint
