// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { LabelDB } from '../label-db'
import { LabelService } from '../label.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { ChartConfigurationDB } from '../chartconfiguration-db'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// LabelDetailComponent is initilizaed from different routes
// LabelDetailComponentState detail different cases 
enum LabelDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_ChartConfiguration_Labels_SET,
}

@Component({
	selector: 'app-label-detail',
	templateUrl: './label-detail.component.html',
	styleUrls: ['./label-detail.component.css'],
})
export class LabelDetailComponent implements OnInit {

	// insertion point for declarations

	// the LabelDB of interest
	label: LabelDB = new LabelDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: LabelDetailComponentState = LabelDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	constructor(
		private labelService: LabelService,
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
			this.state = LabelDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = LabelDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "Labels":
						// console.log("Label" + " is instanciated with back pointer to instance " + this.id + " ChartConfiguration association Labels")
						this.state = LabelDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_ChartConfiguration_Labels_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getLabel()

		// observable for changes in structs
		this.labelService.LabelServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getLabel()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getLabel(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case LabelDetailComponentState.CREATE_INSTANCE:
						this.label = new (LabelDB)
						break;
					case LabelDetailComponentState.UPDATE_INSTANCE:
						let label = frontRepo.Labels.get(this.id)
						console.assert(label != undefined, "missing label with id:" + this.id)
						this.label = label!
						break;
					// insertion point for init of association field
					case LabelDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_ChartConfiguration_Labels_SET:
						this.label = new (LabelDB)
						this.label.ChartConfiguration_Labels_reverse = frontRepo.ChartConfigurations.get(this.id)!
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
		if (this.label.ChartConfiguration_Labels_reverse != undefined) {
			if (this.label.ChartConfiguration_LabelsDBID == undefined) {
				this.label.ChartConfiguration_LabelsDBID = new NullInt64
			}
			this.label.ChartConfiguration_LabelsDBID.Int64 = this.label.ChartConfiguration_Labels_reverse.ID
			this.label.ChartConfiguration_LabelsDBID.Valid = true
			if (this.label.ChartConfiguration_LabelsDBID_Index == undefined) {
				this.label.ChartConfiguration_LabelsDBID_Index = new NullInt64
			}
			this.label.ChartConfiguration_LabelsDBID_Index.Valid = true
			this.label.ChartConfiguration_Labels_reverse = new ChartConfigurationDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case LabelDetailComponentState.UPDATE_INSTANCE:
				this.labelService.updateLabel(this.label)
					.subscribe(label => {
						this.labelService.LabelServiceChanged.next("update")
					});
				break;
			default:
				this.labelService.postLabel(this.label).subscribe(label => {
					this.labelService.LabelServiceChanged.next("post")
					this.label = new (LabelDB) // reset fields
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

			dialogData.ID = this.label.ID!
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
			dialogData.ID = this.label.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "Label"
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
			ID: this.label.ID,
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
		if (this.label.Name == "") {
			this.label.Name = event.value.Name
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
