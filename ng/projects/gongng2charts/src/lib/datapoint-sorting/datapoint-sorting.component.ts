// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { DataPointDB } from '../datapoint-db'
import { DataPointService } from '../datapoint.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-datapoint-sorting',
  templateUrl: './datapoint-sorting.component.html',
  styleUrls: ['./datapoint-sorting.component.css']
})
export class DataPointSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of DataPoint instances that are in the association
  associatedDataPoints = new Array<DataPointDB>();

  constructor(
    private datapointService: DataPointService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of datapoint instances
    public dialogRef: MatDialogRef<DataPointSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getDataPoints()
  }

  getDataPoints(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let datapoint of this.frontRepo.DataPoints_array) {
          let ID = this.dialogData.ID
          let revPointerID = datapoint[this.dialogData.ReversePointer as keyof DataPointDB] as unknown as NullInt64
          let revPointerID_Index = datapoint[this.dialogData.ReversePointer + "_Index" as keyof DataPointDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedDataPoints.push(datapoint)
          }
        }

        // sort associated datapoint according to order
        this.associatedDataPoints.sort((t1, t2) => {
          let t1_revPointerID_Index = t1[this.dialogData.ReversePointer + "_Index" as keyof typeof t1] as unknown as NullInt64
          let t2_revPointerID_Index = t2[this.dialogData.ReversePointer + "_Index" as keyof typeof t2] as unknown as NullInt64
          if (t1_revPointerID_Index && t2_revPointerID_Index) {
            if (t1_revPointerID_Index.Int64 > t2_revPointerID_Index.Int64) {
              return 1;
            }
            if (t1_revPointerID_Index.Int64 < t2_revPointerID_Index.Int64) {
              return -1;
            }
          }
          return 0;
        });
      }
    )
  }

  drop(event: CdkDragDrop<string[]>) {
    moveItemInArray(this.associatedDataPoints, event.previousIndex, event.currentIndex);

    // set the order of DataPoint instances
    let index = 0

    for (let datapoint of this.associatedDataPoints) {
      let revPointerID_Index = datapoint[this.dialogData.ReversePointer + "_Index" as keyof DataPointDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedDataPoints.forEach(
      datapoint => {
        this.datapointService.updateDataPoint(datapoint)
          .subscribe(datapoint => {
            this.datapointService.DataPointServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer +' done');
  }
}