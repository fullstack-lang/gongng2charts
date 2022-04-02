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
import { DataPointDB } from '../datapoint-db'
import { DataPointService } from '../datapoint.service'

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
  selector: 'app-datapointstable',
  templateUrl: './datapoints-table.component.html',
  styleUrls: ['./datapoints-table.component.css'],
})
export class DataPointsTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of DataPoint instances
  selection: SelectionModel<DataPointDB> = new (SelectionModel)
  initialSelection = new Array<DataPointDB>()

  // the data source for the table
  datapoints: DataPointDB[] = []
  matTableDataSource: MatTableDataSource<DataPointDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.datapoints
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
    this.matTableDataSource.sortingDataAccessor = (datapointDB: DataPointDB, property: string) => {
      switch (property) {
        case 'ID':
          return datapointDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return datapointDB.Name;

        case 'Value':
          return datapointDB.Value;

        case 'Dataset_DataPoints':
          return this.frontRepo.Datasets.get(datapointDB.Dataset_DataPointsDBID.Int64)!.Name;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (datapointDB: DataPointDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the datapointDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += datapointDB.Name.toLowerCase()
      mergedContent += datapointDB.Value.toString()
      if (datapointDB.Dataset_DataPointsDBID.Int64 != 0) {
        mergedContent += this.frontRepo.Datasets.get(datapointDB.Dataset_DataPointsDBID.Int64)!.Name.toLowerCase()
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
    private datapointService: DataPointService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of datapoint instances
    public dialogRef: MatDialogRef<DataPointsTableComponent>,
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
    this.datapointService.DataPointServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getDataPoints()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Value",
        "Dataset_DataPoints",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Value",
        "Dataset_DataPoints",
      ]
      this.selection = new SelectionModel<DataPointDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getDataPoints()
    this.matTableDataSource = new MatTableDataSource(this.datapoints)
  }

  getDataPoints(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.datapoints = this.frontRepo.DataPoints_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries
        
        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let datapoint of this.datapoints) {
            let ID = this.dialogData.ID
            let revPointer = datapoint[this.dialogData.ReversePointer as keyof DataPointDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(datapoint)
            }
            this.selection = new SelectionModel<DataPointDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, DataPointDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as DataPointDB[]
          for (let associationInstance of sourceField) {
            let datapoint = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as DataPointDB
            this.initialSelection.push(datapoint)
          }

          this.selection = new SelectionModel<DataPointDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.datapoints
      }
    )
  }

  // newDataPoint initiate a new datapoint
  // create a new DataPoint objet
  newDataPoint() {
  }

  deleteDataPoint(datapointID: number, datapoint: DataPointDB) {
    // list of datapoints is truncated of datapoint before the delete
    this.datapoints = this.datapoints.filter(h => h !== datapoint);

    this.datapointService.deleteDataPoint(datapointID).subscribe(
      datapoint => {
        this.datapointService.DataPointServiceChanged.next("delete")
      }
    );
  }

  editDataPoint(datapointID: number, datapoint: DataPointDB) {

  }

  // display datapoint in router
  displayDataPointInRouter(datapointID: number) {
    this.router.navigate(["github_com_fullstack_lang_gongng2charts_go-" + "datapoint-display", datapointID])
  }

  // set editor outlet
  setEditorRouterOutlet(datapointID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongng2charts_go_editor: ["github_com_fullstack_lang_gongng2charts_go-" + "datapoint-detail", datapointID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(datapointID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongng2charts_go_presentation: ["github_com_fullstack_lang_gongng2charts_go-" + "datapoint-presentation", datapointID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.datapoints.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.datapoints.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<DataPointDB>()

      // reset all initial selection of datapoint that belong to datapoint
      for (let datapoint of this.initialSelection) {
        let index = datapoint[this.dialogData.ReversePointer as keyof DataPointDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(datapoint)

      }

      // from selection, set datapoint that belong to datapoint
      for (let datapoint of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = datapoint[this.dialogData.ReversePointer as keyof DataPointDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(datapoint)
      }


      // update all datapoint (only update selection & initial selection)
      for (let datapoint of toUpdate) {
        this.datapointService.updateDataPoint(datapoint)
          .subscribe(datapoint => {
            this.datapointService.DataPointServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, DataPointDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedDataPoint = new Set<number>()
      for (let datapoint of this.initialSelection) {
        if (this.selection.selected.includes(datapoint)) {
          // console.log("datapoint " + datapoint.Name + " is still selected")
        } else {
          console.log("datapoint " + datapoint.Name + " has been unselected")
          unselectedDataPoint.add(datapoint.ID)
          console.log("is unselected " + unselectedDataPoint.has(datapoint.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let datapoint = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as DataPointDB
      if (unselectedDataPoint.has(datapoint.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<DataPointDB>) = new Array<DataPointDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          datapoint => {
            if (!this.initialSelection.includes(datapoint)) {
              // console.log("datapoint " + datapoint.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + datapoint.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = datapoint.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = datapoint.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("datapoint " + datapoint.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<DataPointDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}