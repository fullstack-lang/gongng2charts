// generated by gong
import { Component, OnInit, AfterViewInit, ViewChild, Inject, Optional } from '@angular/core';
import { BehaviorSubject } from 'rxjs'
import { MatSort } from '@angular/material/sort';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData, FrontRepoService, FrontRepo, SelectionMode } from '../front-repo.service'
import { NullInt64 } from '../null-int64'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { DatasetDB } from '../dataset-db'
import { DatasetService } from '../dataset.service'

// insertion point for additional imports

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-datasetstable',
  templateUrl: './datasets-table.component.html',
  styleUrls: ['./datasets-table.component.css'],
})
export class DatasetsTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Dataset instances
  selection: SelectionModel<DatasetDB> = new (SelectionModel)
  initialSelection = new Array<DatasetDB>()

  // the data source for the table
  datasets: DatasetDB[] = []
  matTableDataSource: MatTableDataSource<DatasetDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.datasets
  frontRepo: FrontRepo = new (FrontRepo)

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  // for sorting & pagination
  @ViewChild(MatSort)
  sort: MatSort | undefined
  @ViewChild(MatPaginator)
  paginator: MatPaginator | undefined;

  ngAfterViewInit() {

    // enable sorting on all fields (including pointers and reverse pointer)
    this.matTableDataSource.sortingDataAccessor = (datasetDB: DatasetDB, property: string) => {
      switch (property) {
        case 'ID':
          return datasetDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return datasetDB.Name;

        case 'ChartConfiguration_Datasets':
          return this.frontRepo.ChartConfigurations.get(datasetDB.ChartConfiguration_DatasetsDBID.Int64)!.Name;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (datasetDB: DatasetDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the datasetDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += datasetDB.Name.toLowerCase()
      if (datasetDB.ChartConfiguration_DatasetsDBID.Int64 != 0) {
        mergedContent += this.frontRepo.ChartConfigurations.get(datasetDB.ChartConfiguration_DatasetsDBID.Int64)!.Name.toLowerCase()
      }


      let isSelected = mergedContent.includes(filter.toLowerCase())
      return isSelected
    };

    this.matTableDataSource.sort = this.sort!
    this.matTableDataSource.paginator = this.paginator!
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.matTableDataSource.filter = filterValue.trim().toLowerCase();
  }

  constructor(
    private datasetService: DatasetService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of dataset instances
    public dialogRef: MatDialogRef<DatasetsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {

    // compute mode
    if (dialogData == undefined) {
      this.mode = TableComponentMode.DISPLAY_MODE
    } else {
      switch (dialogData.SelectionMode) {
        case SelectionMode.ONE_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.ONE_MANY_ASSOCIATION_MODE
          break
        case SelectionMode.MANY_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.MANY_MANY_ASSOCIATION_MODE
          break
        default:
      }
    }

    // observable for changes in structs
    this.datasetService.DatasetServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getDatasets()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "ChartConfiguration_Datasets",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "ChartConfiguration_Datasets",
      ]
      this.selection = new SelectionModel<DatasetDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getDatasets()
    this.matTableDataSource = new MatTableDataSource(this.datasets)
  }

  getDatasets(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.datasets = this.frontRepo.Datasets_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries
        
        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let dataset of this.datasets) {
            let ID = this.dialogData.ID
            let revPointer = dataset[this.dialogData.ReversePointer as keyof DatasetDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(dataset)
            }
            this.selection = new SelectionModel<DatasetDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, DatasetDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as DatasetDB[]
          for (let associationInstance of sourceField) {
            let dataset = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as DatasetDB
            this.initialSelection.push(dataset)
          }

          this.selection = new SelectionModel<DatasetDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.datasets
      }
    )
  }

  // newDataset initiate a new dataset
  // create a new Dataset objet
  newDataset() {
  }

  deleteDataset(datasetID: number, dataset: DatasetDB) {
    // list of datasets is truncated of dataset before the delete
    this.datasets = this.datasets.filter(h => h !== dataset);

    this.datasetService.deleteDataset(datasetID).subscribe(
      dataset => {
        this.datasetService.DatasetServiceChanged.next("delete")
      }
    );
  }

  editDataset(datasetID: number, dataset: DatasetDB) {

  }

  // display dataset in router
  displayDatasetInRouter(datasetID: number) {
    this.router.navigate(["github_com_fullstack_lang_gongng2charts_go-" + "dataset-display", datasetID])
  }

  // set editor outlet
  setEditorRouterOutlet(datasetID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongng2charts_go_editor: ["github_com_fullstack_lang_gongng2charts_go-" + "dataset-detail", datasetID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(datasetID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongng2charts_go_presentation: ["github_com_fullstack_lang_gongng2charts_go-" + "dataset-presentation", datasetID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.datasets.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.datasets.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<DatasetDB>()

      // reset all initial selection of dataset that belong to dataset
      for (let dataset of this.initialSelection) {
        let index = dataset[this.dialogData.ReversePointer as keyof DatasetDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(dataset)

      }

      // from selection, set dataset that belong to dataset
      for (let dataset of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = dataset[this.dialogData.ReversePointer as keyof DatasetDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(dataset)
      }


      // update all dataset (only update selection & initial selection)
      for (let dataset of toUpdate) {
        this.datasetService.updateDataset(dataset)
          .subscribe(dataset => {
            this.datasetService.DatasetServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, DatasetDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedDataset = new Set<number>()
      for (let dataset of this.initialSelection) {
        if (this.selection.selected.includes(dataset)) {
          // console.log("dataset " + dataset.Name + " is still selected")
        } else {
          console.log("dataset " + dataset.Name + " has been unselected")
          unselectedDataset.add(dataset.ID)
          console.log("is unselected " + unselectedDataset.has(dataset.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let dataset = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as DatasetDB
      if (unselectedDataset.has(dataset.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<DatasetDB>) = new Array<DatasetDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          dataset => {
            if (!this.initialSelection.includes(dataset)) {
              // console.log("dataset " + dataset.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + dataset.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = dataset.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = dataset.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("dataset " + dataset.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<DatasetDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
