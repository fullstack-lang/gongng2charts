// generated by genORMTranslation.go
package orm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gorm.io/gorm"

	"github.com/fullstack-lang/gong/go/models"

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoGongBasicField BackRepoGongBasicFieldStruct

	BackRepoGongEnum BackRepoGongEnumStruct

	BackRepoGongEnumValue BackRepoGongEnumValueStruct

	BackRepoGongNote BackRepoGongNoteStruct

	BackRepoGongStruct BackRepoGongStructStruct

	BackRepoGongTimeField BackRepoGongTimeFieldStruct

	BackRepoModelPkg BackRepoModelPkgStruct

	BackRepoPointerToGongStructField BackRepoPointerToGongStructFieldStruct

	BackRepoSliceOfPointerToGongStructField BackRepoSliceOfPointerToGongStructFieldStruct

	CommitFromBackNb uint // this ng is updated at the BackRepo level but also at the BackRepo<GongStruct> level

	PushFromFrontNb uint // records increments from push from front
}

func (backRepo *BackRepoStruct) GetLastCommitFromBackNb() uint {
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitFromBackNb() uint {
	if models.Stage.OnInitCommitCallback != nil {
		models.Stage.OnInitCommitCallback.BeforeCommit(&models.Stage)
	}
	if models.Stage.OnInitCommitFromBackCallback != nil {
		models.Stage.OnInitCommitFromBackCallback.BeforeCommit(&models.Stage)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	if models.Stage.OnInitCommitCallback != nil {
		models.Stage.OnInitCommitCallback.BeforeCommit(&models.Stage)
	}
	if models.Stage.OnInitCommitFromFrontCallback != nil {
		models.Stage.OnInitCommitFromFrontCallback.BeforeCommit(&models.Stage)
	}
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitFromBackNb
}

// Init the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) init(db *gorm.DB) {
	// insertion point for per struct back repo declarations
	backRepo.BackRepoGongBasicField.Init(db)
	backRepo.BackRepoGongEnum.Init(db)
	backRepo.BackRepoGongEnumValue.Init(db)
	backRepo.BackRepoGongNote.Init(db)
	backRepo.BackRepoGongStruct.Init(db)
	backRepo.BackRepoGongTimeField.Init(db)
	backRepo.BackRepoModelPkg.Init(db)
	backRepo.BackRepoPointerToGongStructField.Init(db)
	backRepo.BackRepoSliceOfPointerToGongStructField.Init(db)

	models.Stage.BackRepo = backRepo
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoGongBasicField.CommitPhaseOne(stage)
	backRepo.BackRepoGongEnum.CommitPhaseOne(stage)
	backRepo.BackRepoGongEnumValue.CommitPhaseOne(stage)
	backRepo.BackRepoGongNote.CommitPhaseOne(stage)
	backRepo.BackRepoGongStruct.CommitPhaseOne(stage)
	backRepo.BackRepoGongTimeField.CommitPhaseOne(stage)
	backRepo.BackRepoModelPkg.CommitPhaseOne(stage)
	backRepo.BackRepoPointerToGongStructField.CommitPhaseOne(stage)
	backRepo.BackRepoSliceOfPointerToGongStructField.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoGongBasicField.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongEnum.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongEnumValue.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongNote.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongStruct.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongTimeField.CommitPhaseTwo(backRepo)
	backRepo.BackRepoModelPkg.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPointerToGongStructField.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSliceOfPointerToGongStructField.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoGongBasicField.CheckoutPhaseOne()
	backRepo.BackRepoGongEnum.CheckoutPhaseOne()
	backRepo.BackRepoGongEnumValue.CheckoutPhaseOne()
	backRepo.BackRepoGongNote.CheckoutPhaseOne()
	backRepo.BackRepoGongStruct.CheckoutPhaseOne()
	backRepo.BackRepoGongTimeField.CheckoutPhaseOne()
	backRepo.BackRepoModelPkg.CheckoutPhaseOne()
	backRepo.BackRepoPointerToGongStructField.CheckoutPhaseOne()
	backRepo.BackRepoSliceOfPointerToGongStructField.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoGongBasicField.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongEnum.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongEnumValue.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongNote.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongStruct.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongTimeField.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoModelPkg.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPointerToGongStructField.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSliceOfPointerToGongStructField.CheckoutPhaseTwo(backRepo)
}

var BackRepo BackRepoStruct

func GetLastCommitFromBackNb() uint {
	return BackRepo.GetLastCommitFromBackNb()
}

func GetLastPushFromFrontNb() uint {
	return BackRepo.GetLastPushFromFrontNb()
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoGongBasicField.Backup(dirPath)
	backRepo.BackRepoGongEnum.Backup(dirPath)
	backRepo.BackRepoGongEnumValue.Backup(dirPath)
	backRepo.BackRepoGongNote.Backup(dirPath)
	backRepo.BackRepoGongStruct.Backup(dirPath)
	backRepo.BackRepoGongTimeField.Backup(dirPath)
	backRepo.BackRepoModelPkg.Backup(dirPath)
	backRepo.BackRepoPointerToGongStructField.Backup(dirPath)
	backRepo.BackRepoSliceOfPointerToGongStructField.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoGongBasicField.BackupXL(file)
	backRepo.BackRepoGongEnum.BackupXL(file)
	backRepo.BackRepoGongEnumValue.BackupXL(file)
	backRepo.BackRepoGongNote.BackupXL(file)
	backRepo.BackRepoGongStruct.BackupXL(file)
	backRepo.BackRepoGongTimeField.BackupXL(file)
	backRepo.BackRepoModelPkg.BackupXL(file)
	backRepo.BackRepoPointerToGongStructField.BackupXL(file)
	backRepo.BackRepoSliceOfPointerToGongStructField.BackupXL(file)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(dirPath, "bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) Restore(stage *models.StageStruct, dirPath string) {
	models.Stage.Commit()
	models.Stage.Reset()
	models.Stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoGongBasicField.RestorePhaseOne(dirPath)
	backRepo.BackRepoGongEnum.RestorePhaseOne(dirPath)
	backRepo.BackRepoGongEnumValue.RestorePhaseOne(dirPath)
	backRepo.BackRepoGongNote.RestorePhaseOne(dirPath)
	backRepo.BackRepoGongStruct.RestorePhaseOne(dirPath)
	backRepo.BackRepoGongTimeField.RestorePhaseOne(dirPath)
	backRepo.BackRepoModelPkg.RestorePhaseOne(dirPath)
	backRepo.BackRepoPointerToGongStructField.RestorePhaseOne(dirPath)
	backRepo.BackRepoSliceOfPointerToGongStructField.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoGongBasicField.RestorePhaseTwo()
	backRepo.BackRepoGongEnum.RestorePhaseTwo()
	backRepo.BackRepoGongEnumValue.RestorePhaseTwo()
	backRepo.BackRepoGongNote.RestorePhaseTwo()
	backRepo.BackRepoGongStruct.RestorePhaseTwo()
	backRepo.BackRepoGongTimeField.RestorePhaseTwo()
	backRepo.BackRepoModelPkg.RestorePhaseTwo()
	backRepo.BackRepoPointerToGongStructField.RestorePhaseTwo()
	backRepo.BackRepoSliceOfPointerToGongStructField.RestorePhaseTwo()

	models.Stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {

	// clean the stage
	models.Stage.Reset()

	// commit the cleaned stage
	models.Stage.Commit()

	// open an existing file
	filename := filepath.Join(dirPath, "bckp.xlsx")
	file, err := xlsx.OpenFile(filename)

	if err != nil {
		log.Panic("Cannot read the XL file", err.Error())
	}

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoGongBasicField.RestoreXLPhaseOne(file)
	backRepo.BackRepoGongEnum.RestoreXLPhaseOne(file)
	backRepo.BackRepoGongEnumValue.RestoreXLPhaseOne(file)
	backRepo.BackRepoGongNote.RestoreXLPhaseOne(file)
	backRepo.BackRepoGongStruct.RestoreXLPhaseOne(file)
	backRepo.BackRepoGongTimeField.RestoreXLPhaseOne(file)
	backRepo.BackRepoModelPkg.RestoreXLPhaseOne(file)
	backRepo.BackRepoPointerToGongStructField.RestoreXLPhaseOne(file)
	backRepo.BackRepoSliceOfPointerToGongStructField.RestoreXLPhaseOne(file)

	// commit the restored stage
	models.Stage.Commit()
}
