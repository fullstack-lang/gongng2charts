import { NgModule } from '@angular/core';
import { Gongng2chartsspecificComponent } from './gongng2chartsspecific.component';
import { Gongng2chartsChartComponent } from './gongng2charts-chart/gongng2charts-chart.component';

import { NgChartsModule } from 'ng2-charts';

@NgModule({
  declarations: [
    Gongng2chartsspecificComponent,
    Gongng2chartsChartComponent
  ],
  imports: [
    NgChartsModule,
  ],
  exports: [
    Gongng2chartsspecificComponent,
    Gongng2chartsChartComponent
  ]
})
export class Gongng2chartsspecificModule { }
