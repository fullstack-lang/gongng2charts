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

	"github.com/fullstack-lang/gongdoc/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Vertice_sql sql.NullBool
var dummy_Vertice_time time.Duration
var dummy_Vertice_sort sort.Float64Slice

// VerticeAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model verticeAPI
type VerticeAPI struct {
	gorm.Model

	models.Vertice

	// encoding of pointers
	VerticePointersEnconding
}

// VerticePointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type VerticePointersEnconding struct {
	// insertion for pointer fields encoding declaration
}

// VerticeDB describes a vertice in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model verticeDB
type VerticeDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field verticeDB.X
	X_Data sql.NullFloat64

	// Declation for basic field verticeDB.Y
	Y_Data sql.NullFloat64

	// Declation for basic field verticeDB.Name
	Name_Data sql.NullString
	// encoding of pointers
	VerticePointersEnconding
}

// VerticeDBs arrays verticeDBs
// swagger:response verticeDBsResponse
type VerticeDBs []VerticeDB

// VerticeDBResponse provides response
// swagger:response verticeDBResponse
type VerticeDBResponse struct {
	VerticeDB
}

// VerticeWOP is a Vertice without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type VerticeWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	X float64 `xlsx:"1"`

	Y float64 `xlsx:"2"`

	Name string `xlsx:"3"`
	// insertion for WOP pointer fields
}

var Vertice_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"X",
	"Y",
	"Name",
}

type BackRepoVerticeStruct struct {
	// stores VerticeDB according to their gorm ID
	Map_VerticeDBID_VerticeDB *map[uint]*VerticeDB

	// stores VerticeDB ID according to Vertice address
	Map_VerticePtr_VerticeDBID *map[*models.Vertice]uint

	// stores Vertice according to their gorm ID
	Map_VerticeDBID_VerticePtr *map[uint]*models.Vertice

	db *gorm.DB
}

func (backRepoVertice *BackRepoVerticeStruct) GetDB() *gorm.DB {
	return backRepoVertice.db
}

// GetVerticeDBFromVerticePtr is a handy function to access the back repo instance from the stage instance
func (backRepoVertice *BackRepoVerticeStruct) GetVerticeDBFromVerticePtr(vertice *models.Vertice) (verticeDB *VerticeDB) {
	id := (*backRepoVertice.Map_VerticePtr_VerticeDBID)[vertice]
	verticeDB = (*backRepoVertice.Map_VerticeDBID_VerticeDB)[id]
	return
}

// BackRepoVertice.Init set up the BackRepo of the Vertice
func (backRepoVertice *BackRepoVerticeStruct) Init(db *gorm.DB) (Error error) {

	if backRepoVertice.Map_VerticeDBID_VerticePtr != nil {
		err := errors.New("In Init, backRepoVertice.Map_VerticeDBID_VerticePtr should be nil")
		return err
	}

	if backRepoVertice.Map_VerticeDBID_VerticeDB != nil {
		err := errors.New("In Init, backRepoVertice.Map_VerticeDBID_VerticeDB should be nil")
		return err
	}

	if backRepoVertice.Map_VerticePtr_VerticeDBID != nil {
		err := errors.New("In Init, backRepoVertice.Map_VerticePtr_VerticeDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.Vertice, 0)
	backRepoVertice.Map_VerticeDBID_VerticePtr = &tmp

	tmpDB := make(map[uint]*VerticeDB, 0)
	backRepoVertice.Map_VerticeDBID_VerticeDB = &tmpDB

	tmpID := make(map[*models.Vertice]uint, 0)
	backRepoVertice.Map_VerticePtr_VerticeDBID = &tmpID

	backRepoVertice.db = db
	return
}

// BackRepoVertice.CommitPhaseOne commits all staged instances of Vertice to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoVertice *BackRepoVerticeStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for vertice := range stage.Vertices {
		backRepoVertice.CommitPhaseOneInstance(vertice)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, vertice := range *backRepoVertice.Map_VerticeDBID_VerticePtr {
		if _, ok := stage.Vertices[vertice]; !ok {
			backRepoVertice.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoVertice.CommitDeleteInstance commits deletion of Vertice to the BackRepo
func (backRepoVertice *BackRepoVerticeStruct) CommitDeleteInstance(id uint) (Error error) {

	vertice := (*backRepoVertice.Map_VerticeDBID_VerticePtr)[id]

	// vertice is not staged anymore, remove verticeDB
	verticeDB := (*backRepoVertice.Map_VerticeDBID_VerticeDB)[id]
	query := backRepoVertice.db.Unscoped().Delete(&verticeDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoVertice.Map_VerticePtr_VerticeDBID), vertice)
	delete((*backRepoVertice.Map_VerticeDBID_VerticePtr), id)
	delete((*backRepoVertice.Map_VerticeDBID_VerticeDB), id)

	return
}

// BackRepoVertice.CommitPhaseOneInstance commits vertice staged instances of Vertice to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoVertice *BackRepoVerticeStruct) CommitPhaseOneInstance(vertice *models.Vertice) (Error error) {

	// check if the vertice is not commited yet
	if _, ok := (*backRepoVertice.Map_VerticePtr_VerticeDBID)[vertice]; ok {
		return
	}

	// initiate vertice
	var verticeDB VerticeDB
	verticeDB.CopyBasicFieldsFromVertice(vertice)

	query := backRepoVertice.db.Create(&verticeDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoVertice.Map_VerticePtr_VerticeDBID)[vertice] = verticeDB.ID
	(*backRepoVertice.Map_VerticeDBID_VerticePtr)[verticeDB.ID] = vertice
	(*backRepoVertice.Map_VerticeDBID_VerticeDB)[verticeDB.ID] = &verticeDB

	return
}

// BackRepoVertice.CommitPhaseTwo commits all staged instances of Vertice to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoVertice *BackRepoVerticeStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, vertice := range *backRepoVertice.Map_VerticeDBID_VerticePtr {
		backRepoVertice.CommitPhaseTwoInstance(backRepo, idx, vertice)
	}

	return
}

// BackRepoVertice.CommitPhaseTwoInstance commits {{structname }} of models.Vertice to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoVertice *BackRepoVerticeStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, vertice *models.Vertice) (Error error) {

	// fetch matching verticeDB
	if verticeDB, ok := (*backRepoVertice.Map_VerticeDBID_VerticeDB)[idx]; ok {

		verticeDB.CopyBasicFieldsFromVertice(vertice)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoVertice.db.Save(&verticeDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Vertice intance %s", vertice.Name))
		return err
	}

	return
}

// BackRepoVertice.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoVertice *BackRepoVerticeStruct) CheckoutPhaseOne() (Error error) {

	verticeDBArray := make([]VerticeDB, 0)
	query := backRepoVertice.db.Find(&verticeDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	verticeInstancesToBeRemovedFromTheStage := make(map[*models.Vertice]any)
	for key, value := range models.Stage.Vertices {
		verticeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, verticeDB := range verticeDBArray {
		backRepoVertice.CheckoutPhaseOneInstance(&verticeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		vertice, ok := (*backRepoVertice.Map_VerticeDBID_VerticePtr)[verticeDB.ID]
		if ok {
			delete(verticeInstancesToBeRemovedFromTheStage, vertice)
		}
	}

	// remove from stage and back repo's 3 maps all vertices that are not in the checkout
	for vertice := range verticeInstancesToBeRemovedFromTheStage {
		vertice.Unstage()

		// remove instance from the back repo 3 maps
		verticeID := (*backRepoVertice.Map_VerticePtr_VerticeDBID)[vertice]
		delete((*backRepoVertice.Map_VerticePtr_VerticeDBID), vertice)
		delete((*backRepoVertice.Map_VerticeDBID_VerticeDB), verticeID)
		delete((*backRepoVertice.Map_VerticeDBID_VerticePtr), verticeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a verticeDB that has been found in the DB, updates the backRepo and stages the
// models version of the verticeDB
func (backRepoVertice *BackRepoVerticeStruct) CheckoutPhaseOneInstance(verticeDB *VerticeDB) (Error error) {

	vertice, ok := (*backRepoVertice.Map_VerticeDBID_VerticePtr)[verticeDB.ID]
	if !ok {
		vertice = new(models.Vertice)

		(*backRepoVertice.Map_VerticeDBID_VerticePtr)[verticeDB.ID] = vertice
		(*backRepoVertice.Map_VerticePtr_VerticeDBID)[vertice] = verticeDB.ID

		// append model store with the new element
		vertice.Name = verticeDB.Name_Data.String
		vertice.Stage()
	}
	verticeDB.CopyBasicFieldsToVertice(vertice)

	// preserve pointer to verticeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_VerticeDBID_VerticeDB)[verticeDB hold variable pointers
	verticeDB_Data := *verticeDB
	preservedPtrToVertice := &verticeDB_Data
	(*backRepoVertice.Map_VerticeDBID_VerticeDB)[verticeDB.ID] = preservedPtrToVertice

	return
}

// BackRepoVertice.CheckoutPhaseTwo Checkouts all staged instances of Vertice to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoVertice *BackRepoVerticeStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, verticeDB := range *backRepoVertice.Map_VerticeDBID_VerticeDB {
		backRepoVertice.CheckoutPhaseTwoInstance(backRepo, verticeDB)
	}
	return
}

// BackRepoVertice.CheckoutPhaseTwoInstance Checkouts staged instances of Vertice to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoVertice *BackRepoVerticeStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, verticeDB *VerticeDB) (Error error) {

	vertice := (*backRepoVertice.Map_VerticeDBID_VerticePtr)[verticeDB.ID]
	_ = vertice // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitVertice allows commit of a single vertice (if already staged)
func (backRepo *BackRepoStruct) CommitVertice(vertice *models.Vertice) {
	backRepo.BackRepoVertice.CommitPhaseOneInstance(vertice)
	if id, ok := (*backRepo.BackRepoVertice.Map_VerticePtr_VerticeDBID)[vertice]; ok {
		backRepo.BackRepoVertice.CommitPhaseTwoInstance(backRepo, id, vertice)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitVertice allows checkout of a single vertice (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutVertice(vertice *models.Vertice) {
	// check if the vertice is staged
	if _, ok := (*backRepo.BackRepoVertice.Map_VerticePtr_VerticeDBID)[vertice]; ok {

		if id, ok := (*backRepo.BackRepoVertice.Map_VerticePtr_VerticeDBID)[vertice]; ok {
			var verticeDB VerticeDB
			verticeDB.ID = id

			if err := backRepo.BackRepoVertice.db.First(&verticeDB, id).Error; err != nil {
				log.Panicln("CheckoutVertice : Problem with getting object with id:", id)
			}
			backRepo.BackRepoVertice.CheckoutPhaseOneInstance(&verticeDB)
			backRepo.BackRepoVertice.CheckoutPhaseTwoInstance(backRepo, &verticeDB)
		}
	}
}

// CopyBasicFieldsFromVertice
func (verticeDB *VerticeDB) CopyBasicFieldsFromVertice(vertice *models.Vertice) {
	// insertion point for fields commit

	verticeDB.X_Data.Float64 = vertice.X
	verticeDB.X_Data.Valid = true

	verticeDB.Y_Data.Float64 = vertice.Y
	verticeDB.Y_Data.Valid = true

	verticeDB.Name_Data.String = vertice.Name
	verticeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromVerticeWOP
func (verticeDB *VerticeDB) CopyBasicFieldsFromVerticeWOP(vertice *VerticeWOP) {
	// insertion point for fields commit

	verticeDB.X_Data.Float64 = vertice.X
	verticeDB.X_Data.Valid = true

	verticeDB.Y_Data.Float64 = vertice.Y
	verticeDB.Y_Data.Valid = true

	verticeDB.Name_Data.String = vertice.Name
	verticeDB.Name_Data.Valid = true
}

// CopyBasicFieldsToVertice
func (verticeDB *VerticeDB) CopyBasicFieldsToVertice(vertice *models.Vertice) {
	// insertion point for checkout of basic fields (back repo to stage)
	vertice.X = verticeDB.X_Data.Float64
	vertice.Y = verticeDB.Y_Data.Float64
	vertice.Name = verticeDB.Name_Data.String
}

// CopyBasicFieldsToVerticeWOP
func (verticeDB *VerticeDB) CopyBasicFieldsToVerticeWOP(vertice *VerticeWOP) {
	vertice.ID = int(verticeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	vertice.X = verticeDB.X_Data.Float64
	vertice.Y = verticeDB.Y_Data.Float64
	vertice.Name = verticeDB.Name_Data.String
}

// Backup generates a json file from a slice of all VerticeDB instances in the backrepo
func (backRepoVertice *BackRepoVerticeStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "VerticeDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*VerticeDB, 0)
	for _, verticeDB := range *backRepoVertice.Map_VerticeDBID_VerticeDB {
		forBackup = append(forBackup, verticeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Vertice ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Vertice file", err.Error())
	}
}

// Backup generates a json file from a slice of all VerticeDB instances in the backrepo
func (backRepoVertice *BackRepoVerticeStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*VerticeDB, 0)
	for _, verticeDB := range *backRepoVertice.Map_VerticeDBID_VerticeDB {
		forBackup = append(forBackup, verticeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Vertice")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Vertice_Fields, -1)
	for _, verticeDB := range forBackup {

		var verticeWOP VerticeWOP
		verticeDB.CopyBasicFieldsToVerticeWOP(&verticeWOP)

		row := sh.AddRow()
		row.WriteStruct(&verticeWOP, -1)
	}
}

// RestoreXL from the "Vertice" sheet all VerticeDB instances
func (backRepoVertice *BackRepoVerticeStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoVerticeid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Vertice"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoVertice.rowVisitorVertice)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoVertice *BackRepoVerticeStruct) rowVisitorVertice(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var verticeWOP VerticeWOP
		row.ReadStruct(&verticeWOP)

		// add the unmarshalled struct to the stage
		verticeDB := new(VerticeDB)
		verticeDB.CopyBasicFieldsFromVerticeWOP(&verticeWOP)

		verticeDB_ID_atBackupTime := verticeDB.ID
		verticeDB.ID = 0
		query := backRepoVertice.db.Create(verticeDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoVertice.Map_VerticeDBID_VerticeDB)[verticeDB.ID] = verticeDB
		BackRepoVerticeid_atBckpTime_newID[verticeDB_ID_atBackupTime] = verticeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "VerticeDB.json" in dirPath that stores an array
// of VerticeDB and stores it in the database
// the map BackRepoVerticeid_atBckpTime_newID is updated accordingly
func (backRepoVertice *BackRepoVerticeStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoVerticeid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "VerticeDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Vertice file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*VerticeDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_VerticeDBID_VerticeDB
	for _, verticeDB := range forRestore {

		verticeDB_ID_atBackupTime := verticeDB.ID
		verticeDB.ID = 0
		query := backRepoVertice.db.Create(verticeDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoVertice.Map_VerticeDBID_VerticeDB)[verticeDB.ID] = verticeDB
		BackRepoVerticeid_atBckpTime_newID[verticeDB_ID_atBackupTime] = verticeDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Vertice file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Vertice>id_atBckpTime_newID
// to compute new index
func (backRepoVertice *BackRepoVerticeStruct) RestorePhaseTwo() {

	for _, verticeDB := range *backRepoVertice.Map_VerticeDBID_VerticeDB {

		// next line of code is to avert unused variable compilation error
		_ = verticeDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoVertice.db.Model(verticeDB).Updates(*verticeDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoVerticeid_atBckpTime_newID map[uint]uint
