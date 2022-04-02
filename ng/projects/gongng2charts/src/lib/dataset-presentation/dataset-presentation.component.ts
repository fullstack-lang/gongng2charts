import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { DatasetDB } from '../dataset-db'
import { DatasetService } from '../dataset.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface datasetDummyElement {
}

const ELEMENT_DATA: datasetDummyElement[] = [
];

@Component({
	selector: 'app-dataset-presentation',
	templateUrl: './dataset-presentation.component.html',
	styleUrls: ['./dataset-presentation.component.css'],
})
export class DatasetPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	dataset: DatasetDB = new (DatasetDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private datasetService: DatasetService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getDataset();

		// observable for changes in 
		this.datasetService.DatasetServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getDataset()
				}
			}
		)
	}

	getDataset(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.dataset = this.frontRepo.Datasets.get(id)!

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
				github_com_fullstack_lang_gongng2charts_go_editor: ["github_com_fullstack_lang_gongng2charts_go-" + "dataset-detail", ID]
			}
		}]);
	}
}
