import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { ChartDB } from '../chart-db'
import { ChartService } from '../chart.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface chartDummyElement {
}

const ELEMENT_DATA: chartDummyElement[] = [
];

@Component({
	selector: 'app-chart-presentation',
	templateUrl: './chart-presentation.component.html',
	styleUrls: ['./chart-presentation.component.css'],
})
export class ChartPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	chart: ChartDB = new (ChartDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private chartService: ChartService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getChart();

		// observable for changes in 
		this.chartService.ChartServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getChart()
				}
			}
		)
	}

	getChart(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.chart = this.frontRepo.Charts.get(id)!

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
				github_com_fullstack_lang_gongng2charts_go_editor: ["github_com_fullstack_lang_gongng2charts_go-" + "chart-detail", ID]
			}
		}]);
	}
}
