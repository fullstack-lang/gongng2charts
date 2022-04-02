import { NgModule } from '@angular/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatSortModule } from '@angular/material/sort'
import { MatPaginatorModule } from '@angular/material/paginator'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatExpansionModule } from '@angular/material/expansion';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatTreeModule } from '@angular/material/tree';
import { DragDropModule } from '@angular/cdk/drag-drop';

import { AngularSplitModule, SplitComponent } from 'angular-split';

import {
	NgxMatDatetimePickerModule,
	NgxMatNativeDateModule,
	NgxMatTimepickerModule
} from '@angular-material-components/datetime-picker';

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent } from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';

// insertion point for imports 
import { ChartConfigurationsTableComponent } from './chartconfigurations-table/chartconfigurations-table.component'
import { ChartConfigurationSortingComponent } from './chartconfiguration-sorting/chartconfiguration-sorting.component'
import { ChartConfigurationDetailComponent } from './chartconfiguration-detail/chartconfiguration-detail.component'
import { ChartConfigurationPresentationComponent } from './chartconfiguration-presentation/chartconfiguration-presentation.component'

import { DataPointsTableComponent } from './datapoints-table/datapoints-table.component'
import { DataPointSortingComponent } from './datapoint-sorting/datapoint-sorting.component'
import { DataPointDetailComponent } from './datapoint-detail/datapoint-detail.component'
import { DataPointPresentationComponent } from './datapoint-presentation/datapoint-presentation.component'

import { DatasetsTableComponent } from './datasets-table/datasets-table.component'
import { DatasetSortingComponent } from './dataset-sorting/dataset-sorting.component'
import { DatasetDetailComponent } from './dataset-detail/dataset-detail.component'
import { DatasetPresentationComponent } from './dataset-presentation/dataset-presentation.component'

import { LabelsTableComponent } from './labels-table/labels-table.component'
import { LabelSortingComponent } from './label-sorting/label-sorting.component'
import { LabelDetailComponent } from './label-detail/label-detail.component'
import { LabelPresentationComponent } from './label-presentation/label-presentation.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		ChartConfigurationsTableComponent,
		ChartConfigurationSortingComponent,
		ChartConfigurationDetailComponent,
		ChartConfigurationPresentationComponent,

		DataPointsTableComponent,
		DataPointSortingComponent,
		DataPointDetailComponent,
		DataPointPresentationComponent,

		DatasetsTableComponent,
		DatasetSortingComponent,
		DatasetDetailComponent,
		DatasetPresentationComponent,

		LabelsTableComponent,
		LabelSortingComponent,
		LabelDetailComponent,
		LabelPresentationComponent,


		SplitterComponent,
		SidebarComponent
	],
	imports: [
		FormsModule,
		ReactiveFormsModule,
		CommonModule,
		RouterModule,

		AppRoutingModule,

		MatSliderModule,
		MatSelectModule,
		MatFormFieldModule,
		MatInputModule,
		MatDatepickerModule,
		MatTableModule,
		MatSortModule,
		MatPaginatorModule,
		MatCheckboxModule,
		MatButtonModule,
		MatIconModule,
		MatToolbarModule,
		MatExpansionModule,
		MatListModule,
		MatDialogModule,
		MatGridListModule,
		MatTreeModule,
		DragDropModule,

		NgxMatDatetimePickerModule,
		NgxMatNativeDateModule,
		NgxMatTimepickerModule,

		AngularSplitModule,
	],
	exports: [
		// insertion point for declarations 
		ChartConfigurationsTableComponent,
		ChartConfigurationSortingComponent,
		ChartConfigurationDetailComponent,
		ChartConfigurationPresentationComponent,

		DataPointsTableComponent,
		DataPointSortingComponent,
		DataPointDetailComponent,
		DataPointPresentationComponent,

		DatasetsTableComponent,
		DatasetSortingComponent,
		DatasetDetailComponent,
		DatasetPresentationComponent,

		LabelsTableComponent,
		LabelSortingComponent,
		LabelDetailComponent,
		LabelPresentationComponent,


		SplitterComponent,
		SidebarComponent,

	],
	providers: [
		{
			provide: MatDialogRef,
			useValue: {}
		},
	],
})
export class Gongng2chartsModule { }
