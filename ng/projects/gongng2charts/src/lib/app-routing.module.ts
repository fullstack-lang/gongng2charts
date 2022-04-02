import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { ChartsTableComponent } from './charts-table/charts-table.component'
import { ChartDetailComponent } from './chart-detail/chart-detail.component'
import { ChartPresentationComponent } from './chart-presentation/chart-presentation.component'

import { ChartConfigurationsTableComponent } from './chartconfigurations-table/chartconfigurations-table.component'
import { ChartConfigurationDetailComponent } from './chartconfiguration-detail/chartconfiguration-detail.component'
import { ChartConfigurationPresentationComponent } from './chartconfiguration-presentation/chartconfiguration-presentation.component'

import { DataPointsTableComponent } from './datapoints-table/datapoints-table.component'
import { DataPointDetailComponent } from './datapoint-detail/datapoint-detail.component'
import { DataPointPresentationComponent } from './datapoint-presentation/datapoint-presentation.component'

import { DatasetsTableComponent } from './datasets-table/datasets-table.component'
import { DatasetDetailComponent } from './dataset-detail/dataset-detail.component'
import { DatasetPresentationComponent } from './dataset-presentation/dataset-presentation.component'

import { LabelsTableComponent } from './labels-table/labels-table.component'
import { LabelDetailComponent } from './label-detail/label-detail.component'
import { LabelPresentationComponent } from './label-presentation/label-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gongng2charts_go-charts', component: ChartsTableComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_table' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-chart-adder', component: ChartDetailComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_editor' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-chart-adder/:id/:originStruct/:originStructFieldName', component: ChartDetailComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_editor' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-chart-detail/:id', component: ChartDetailComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_editor' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-chart-presentation/:id', component: ChartPresentationComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-chart-presentation-special/:id', component: ChartPresentationComponent, outlet: 'github_com_fullstack_lang_gongng2charts_gochartpres' },

	{ path: 'github_com_fullstack_lang_gongng2charts_go-chartconfigurations', component: ChartConfigurationsTableComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_table' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-chartconfiguration-adder', component: ChartConfigurationDetailComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_editor' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-chartconfiguration-adder/:id/:originStruct/:originStructFieldName', component: ChartConfigurationDetailComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_editor' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-chartconfiguration-detail/:id', component: ChartConfigurationDetailComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_editor' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-chartconfiguration-presentation/:id', component: ChartConfigurationPresentationComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-chartconfiguration-presentation-special/:id', component: ChartConfigurationPresentationComponent, outlet: 'github_com_fullstack_lang_gongng2charts_gochartconfigurationpres' },

	{ path: 'github_com_fullstack_lang_gongng2charts_go-datapoints', component: DataPointsTableComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_table' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-datapoint-adder', component: DataPointDetailComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_editor' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-datapoint-adder/:id/:originStruct/:originStructFieldName', component: DataPointDetailComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_editor' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-datapoint-detail/:id', component: DataPointDetailComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_editor' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-datapoint-presentation/:id', component: DataPointPresentationComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-datapoint-presentation-special/:id', component: DataPointPresentationComponent, outlet: 'github_com_fullstack_lang_gongng2charts_godatapointpres' },

	{ path: 'github_com_fullstack_lang_gongng2charts_go-datasets', component: DatasetsTableComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_table' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-dataset-adder', component: DatasetDetailComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_editor' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-dataset-adder/:id/:originStruct/:originStructFieldName', component: DatasetDetailComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_editor' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-dataset-detail/:id', component: DatasetDetailComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_editor' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-dataset-presentation/:id', component: DatasetPresentationComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-dataset-presentation-special/:id', component: DatasetPresentationComponent, outlet: 'github_com_fullstack_lang_gongng2charts_godatasetpres' },

	{ path: 'github_com_fullstack_lang_gongng2charts_go-labels', component: LabelsTableComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_table' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-label-adder', component: LabelDetailComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_editor' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-label-adder/:id/:originStruct/:originStructFieldName', component: LabelDetailComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_editor' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-label-detail/:id', component: LabelDetailComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_editor' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-label-presentation/:id', component: LabelPresentationComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-label-presentation-special/:id', component: LabelPresentationComponent, outlet: 'github_com_fullstack_lang_gongng2charts_golabelpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
