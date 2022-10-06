// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { DataPointDB } from '../datapoint-db'
import { DataPointService } from '../datapoint.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { DatasetDB } from '../dataset-db'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// DataPointDetailComponent is initilizaed from different routes
// DataPointDetailComponentState detail different cases 
enum DataPointDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Dataset_DataPoints_SET,
}

@Component({
	selector: 'app-datapoint-detail',
	templateUrl: './datapoint-detail.component.html',
	styleUrls: ['./datapoint-detail.component.css'],
})
export class DataPointDetailComponent implements OnInit {

	// insertion point for declarations

	// the DataPointDB of interest
	datapoint: DataPointDB = new DataPointDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: DataPointDetailComponentState = DataPointDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	constructor(
		private datapointService: DataPointService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {

		// compute state
		this.id = +this.route.snapshot.paramMap.get('id')!;
		this.originStruct = this.route.snapshot.paramMap.get('originStruct')!;
		this.originStructFieldName = this.route.snapshot.paramMap.get('originStructFieldName')!;

		const association = this.route.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = DataPointDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = DataPointDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "DataPoints":
						// console.log("DataPoint" + " is instanciated with back pointer to instance " + this.id + " Dataset association DataPoints")
						this.state = DataPointDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Dataset_DataPoints_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getDataPoint()

		// observable for changes in structs
		this.datapointService.DataPointServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getDataPoint()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getDataPoint(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case DataPointDetailComponentState.CREATE_INSTANCE:
						this.datapoint = new (DataPointDB)
						break;
					case DataPointDetailComponentState.UPDATE_INSTANCE:
						let datapoint = frontRepo.DataPoints.get(this.id)
						console.assert(datapoint != undefined, "missing datapoint with id:" + this.id)
						this.datapoint = datapoint!
						break;
					// insertion point for init of association field
					case DataPointDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Dataset_DataPoints_SET:
						this.datapoint = new (DataPointDB)
						this.datapoint.Dataset_DataPoints_reverse = frontRepo.Datasets.get(this.id)!
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.datapoint.Dataset_DataPoints_reverse != undefined) {
			if (this.datapoint.Dataset_DataPointsDBID == undefined) {
				this.datapoint.Dataset_DataPointsDBID = new NullInt64
			}
			this.datapoint.Dataset_DataPointsDBID.Int64 = this.datapoint.Dataset_DataPoints_reverse.ID
			this.datapoint.Dataset_DataPointsDBID.Valid = true
			if (this.datapoint.Dataset_DataPointsDBID_Index == undefined) {
				this.datapoint.Dataset_DataPointsDBID_Index = new NullInt64
			}
			this.datapoint.Dataset_DataPointsDBID_Index.Valid = true
			this.datapoint.Dataset_DataPoints_reverse = new DatasetDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case DataPointDetailComponentState.UPDATE_INSTANCE:
				this.datapointService.updateDataPoint(this.datapoint)
					.subscribe(datapoint => {
						this.datapointService.DataPointServiceChanged.next("update")
					});
				break;
			default:
				this.datapointService.postDataPoint(this.datapoint).subscribe(datapoint => {
					this.datapointService.DataPointServiceChanged.next("post")
					this.datapoint = new (DataPointDB) // reset fields
				});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string, selectionMode: string,
		sourceField: string, intermediateStructField: string, nextAssociatedStruct: string) {

		console.log("mode " + selectionMode)

		const dialogConfig = new MatDialogConfig();

		let dialogData = new DialogData();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		if (selectionMode == SelectionMode.ONE_MANY_ASSOCIATION_MODE) {

			dialogData.ID = this.datapoint.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(AssociatedStruct).get(
					AssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}
		if (selectionMode == SelectionMode.MANY_MANY_ASSOCIATION_MODE) {
			dialogData.ID = this.datapoint.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "DataPoint"
			dialogData.SourceField = sourceField

			// set up the intermediate struct
			dialogData.IntermediateStruct = AssociatedStruct
			dialogData.IntermediateStructField = intermediateStructField

			// set up the end struct
			dialogData.NextAssociationStruct = nextAssociatedStruct

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(nextAssociatedStruct).get(
					nextAssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}

	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.datapoint.ID,
			ReversePointer: reverseField,
			OrderingMode: true,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfSortingComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'SortingComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
		});
	}

	fillUpNameIfEmpty(event: { value: { Name: string; }; }) {
		if (this.datapoint.Name == "") {
			this.datapoint.Name = event.value.Name
		}
	}

	toggleTextArea(fieldName: string) {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			let displayAsTextArea = this.mapFields_displayAsTextArea.get(fieldName)
			this.mapFields_displayAsTextArea.set(fieldName, !displayAsTextArea)
		} else {
			this.mapFields_displayAsTextArea.set(fieldName, true)
		}
	}

	isATextArea(fieldName: string): boolean {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			return this.mapFields_displayAsTextArea.get(fieldName)!
		} else {
			return false
		}
	}

	compareObjects(o1: any, o2: any) {
		if (o1?.ID == o2?.ID) {
			return true;
		}
		else {
			return false
		}
	}
}
