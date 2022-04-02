import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { ChartConfigurationDB } from '../chartconfiguration-db'
import { ChartConfigurationService } from '../chartconfiguration.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface chartconfigurationDummyElement {
}

const ELEMENT_DATA: chartconfigurationDummyElement[] = [
];

@Component({
	selector: 'app-chartconfiguration-presentation',
	templateUrl: './chartconfiguration-presentation.component.html',
	styleUrls: ['./chartconfiguration-presentation.component.css'],
})
export class ChartConfigurationPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	chartconfiguration: ChartConfigurationDB = new (ChartConfigurationDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private chartconfigurationService: ChartConfigurationService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getChartConfiguration();

		// observable for changes in 
		this.chartconfigurationService.ChartConfigurationServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getChartConfiguration()
				}
			}
		)
	}

	getChartConfiguration(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.chartconfiguration = this.frontRepo.ChartConfigurations.get(id)!

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
				github_com_fullstack_lang_gongng2charts_go_editor: ["github_com_fullstack_lang_gongng2charts_go-" + "chartconfiguration-detail", ID]
			}
		}]);
	}
}
