import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { DataPointDB } from '../datapoint-db'
import { DataPointService } from '../datapoint.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface datapointDummyElement {
}

const ELEMENT_DATA: datapointDummyElement[] = [
];

@Component({
	selector: 'app-datapoint-presentation',
	templateUrl: './datapoint-presentation.component.html',
	styleUrls: ['./datapoint-presentation.component.css'],
})
export class DataPointPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	datapoint: DataPointDB = new (DataPointDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private datapointService: DataPointService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getDataPoint();

		// observable for changes in 
		this.datapointService.DataPointServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getDataPoint()
				}
			}
		)
	}

	getDataPoint(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.datapoint = this.frontRepo.DataPoints.get(id)!

				// insertion point for recovery of durations
				// insertion point for recovery of enum tint
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongng2charts_go_presentation: ["github_com_fullstack_lang_gongng2charts_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongng2charts_go_editor: ["github_com_fullstack_lang_gongng2charts_go-" + "datapoint-detail", ID]
			}
		}]);
	}
}
