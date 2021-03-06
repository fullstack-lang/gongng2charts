import { Component, Input, OnInit, ViewChild } from '@angular/core';
import { Router } from '@angular/router'

import { ChartConfiguration, ChartEvent, ChartType } from 'chart.js';
import { BaseChartDirective } from 'ng2-charts';

import { Observable, timer } from 'rxjs';

import * as gongng2charts from 'gongng2charts'
@Component({
  selector: 'lib-gongng2charts-chart',
  templateUrl: './gongng2charts-chart.component.html',
  styleUrls: ['./gongng2charts-chart.component.css']
})
export class Gongng2chartsChartComponent implements OnInit {

  // name of the chart
  @Input() chartName: string = ""

  width = 600
  height = 600


  /**
 * the component is refreshed when modification are performed in the back repo 
 * 
 * the checkCommitNbTimer polls the commit number of the back repo
 * if the commit number has increased, it pulls the front repo and redraw the diagram
 */
  checkCommitNbTimer: Observable<number> = timer(500, 500);
  lastCommitNb = -1
  lastPushFromFrontNb = -1
  currTime: number = 0

  public lineChartData: ChartConfiguration['data'] = {
    datasets: [],
    labels: []
  };

  public lineChartOptions: ChartConfiguration['options'] = {
    elements: {
      line: {
        tension: 0.5
      }
    },
    scales: {
      // We use this empty structure as a placeholder for dynamic theming.
      x: {},
      'y-axis-0':
      {
        position: 'left',
      },
      'y-axis-1': {
        position: 'right',
        grid: {
          color: 'rgba(255,0,0,0.3)',
        },
        ticks: {
          color: 'red'
        }
      }
    },
  }

  public lineChartType: ChartType = 'line';

  @ViewChild(BaseChartDirective) chart?: BaseChartDirective;

  // events
  public chartClicked({ event, active }: { event?: ChartEvent, active?: {}[] }): void {
    console.log(event, active);
  }

  public chartHovered({ event, active }: { event?: ChartEvent, active?: {}[] }): void {
    console.log(event, active);
  }

  frontRepo: gongng2charts.FrontRepo = new gongng2charts.FrontRepo

  constructor(
    private gongng2chartsCommitNbService: gongng2charts.CommitNbService,
    private gongng2chartsPushFromFrontNbService: gongng2charts.PushFromFrontNbService,
    private frontRepoService: gongng2charts.FrontRepoService,
    private router: Router,
  ) {

    // important to distinguish between different chartName
    this.router.routeReuseStrategy.shouldReuseRoute = () => false;
  }

  ngOnInit(): void {

    this.checkCommitNbTimer.subscribe(
      currTime => {
        this.currTime = currTime

        // see above for the explanation
        this.gongng2chartsCommitNbService.getCommitNb().subscribe(
          commitNb => {
            if (this.lastCommitNb < commitNb) {

              console.log("last commit nb " + this.lastCommitNb + " new: " + commitNb)
              this.refresh()
              this.lastCommitNb = commitNb
            }
          }
        )

        // see above for the explanation
        this.gongng2chartsPushFromFrontNbService.getPushFromFrontNb().subscribe(
          pushFromFrontNb => {
            if (this.lastPushFromFrontNb < pushFromFrontNb) {

              console.log("last commit nb " + this.lastPushFromFrontNb + " new: " + pushFromFrontNb)
              this.refresh()
              this.lastPushFromFrontNb = pushFromFrontNb
            }
          }
        )
      }
    )

    this.refresh()
  }

  chartConfig: gongng2charts.ChartConfigurationDB = new gongng2charts.ChartConfigurationDB
  refresh(): void {

    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo returned for map " + this.chartName)

        let matchFound: boolean = false
        for (let chartConfigIdx = 0; chartConfigIdx < this.frontRepo.ChartConfigurations_array.length; chartConfigIdx++) {

          if (this.frontRepo.ChartConfigurations_array[chartConfigIdx].Name == this.chartName) {
            this.chartConfig = this.frontRepo.ChartConfigurations_array[chartConfigIdx]
            matchFound = true
          }
        }

        if (!matchFound) {
          return
        }

        // reset
        this.lineChartData.datasets = []
        this.lineChartData.labels = []

        for (let i = 0; i < this.chartConfig.Datasets!.length; i++) {

          let dataset = this.chartConfig.Datasets![i]
          // let datapoints = new Array<number>()
          let datapoints: number[] = []

          for (let j = 0; j < dataset.DataPoints!.length; j++) {
            let datapoint = dataset.DataPoints![j]
            datapoints.push(datapoint.Value)
          }

          this.lineChartData.datasets.push(
            {
              // data: [65, 59, 80, 81, 56, 55, 40],
              data: datapoints,
              label: dataset.Label,
              backgroundColor: 'rgba(148,159,177,0.2)',
              borderColor: 'rgba(148,159,177,1)',
              pointBackgroundColor: 'rgba(148,159,177,1)',
              pointBorderColor: '#fff',
              pointHoverBackgroundColor: '#fff',
              pointHoverBorderColor: 'rgba(148,159,177,0.8)',
              fill: 'origin',
            },
          )
        }

        for (let i = 0; i < this.chartConfig.Labels!.length; i++) {
          let label = this.chartConfig.Labels![i]
          this.lineChartData.labels.push(label.Name)
        }

        // set chart type
        this.lineChartType = this.chartConfig.ChartType as ChartType

        // set width & height
        this.width = this.chartConfig.Width
        this.height = this.chartConfig.Heigth

        console.log("finished rendering")
        this.chart?.update();
      }
    )
  }
}
