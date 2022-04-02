import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { LabelDB } from '../label-db'
import { LabelService } from '../label.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface labelDummyElement {
}

const ELEMENT_DATA: labelDummyElement[] = [
];

@Component({
	selector: 'app-label-presentation',
	templateUrl: './label-presentation.component.html',
	styleUrls: ['./label-presentation.component.css'],
})
export class LabelPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	label: LabelDB = new (LabelDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private labelService: LabelService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getLabel();

		// observable for changes in 
		this.labelService.LabelServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getLabel()
				}
			}
		)
	}

	getLabel(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.label = this.frontRepo.Labels.get(id)!

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
				github_com_fullstack_lang_gongng2charts_go_editor: ["github_com_fullstack_lang_gongng2charts_go-" + "label-detail", ID]
			}
		}]);
	}
}
