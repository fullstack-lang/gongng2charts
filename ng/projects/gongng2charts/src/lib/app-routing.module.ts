import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { ChartsTableComponent } from './charts-table/charts-table.component'
import { ChartDetailComponent } from './chart-detail/chart-detail.component'
import { ChartPresentationComponent } from './chart-presentation/chart-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gongng2charts_go-charts', component: ChartsTableComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_table' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-chart-adder', component: ChartDetailComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_editor' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-chart-adder/:id/:originStruct/:originStructFieldName', component: ChartDetailComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_editor' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-chart-detail/:id', component: ChartDetailComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_editor' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-chart-presentation/:id', component: ChartPresentationComponent, outlet: 'github_com_fullstack_lang_gongng2charts_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongng2charts_go-chart-presentation-special/:id', component: ChartPresentationComponent, outlet: 'github_com_fullstack_lang_gongng2charts_gochartpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
