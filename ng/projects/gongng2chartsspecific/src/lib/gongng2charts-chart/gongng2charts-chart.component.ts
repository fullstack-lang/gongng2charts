import { Component, OnInit, ViewChild } from '@angular/core';

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
    datasets: [
      {
        data: [65, 59, 80, 81, 56, 55, 40],
        label: 'Series A',
        backgroundColor: 'rgba(148,159,177,0.2)',
        borderColor: 'rgba(148,159,177,1)',
        pointBackgroundColor: 'rgba(148,159,177,1)',
        pointBorderColor: '#fff',
        pointHoverBackgroundColor: '#fff',
        pointHoverBorderColor: 'rgba(148,159,177,0.8)',
        fill: 'origin',
      },
      {
        data: [28, 48, 40, 19, 86, 27, 90],
        label: 'Series B',
        backgroundColor: 'rgba(77,83,96,0.2)',
        borderColor: 'rgba(77,83,96,1)',
        pointBackgroundColor: 'rgba(77,83,96,1)',
        pointBorderColor: '#fff',
        pointHoverBackgroundColor: '#fff',
        pointHoverBorderColor: 'rgba(77,83,96,1)',
        fill: 'origin',
      },
      {
        data: [180, 480, 770, 90, 1000, 270, 400],
        label: 'Series C',
        yAxisID: 'y-axis-1',
        backgroundColor: 'rgba(255,0,0,0.3)',
        borderColor: 'red',
        pointBackgroundColor: 'rgba(148,159,177,1)',
        pointBorderColor: '#fff',
        pointHoverBackgroundColor: '#fff',
        pointHoverBorderColor: 'rgba(148,159,177,0.8)',
        fill: 'origin',
      }
    ],
    labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July']
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


  constructor(
    private gongng2chartsCommitNbService: gongng2charts.CommitNbService,
    private gongng2chartsPushFromFrontNbService: gongng2charts.PushFromFrontNbService,
  ) { }

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
  }

  refresh(): void {
    // // get the singloton
    // this.gongmarkdownMarkdownContentService.getMarkdownContents().subscribe(
    //   markdownContentDBs => {
    //     if (markdownContentDBs.length == 1) {
    //       this.markdownContentDB = markdownContentDBs[0]
    //     } else {
    //       console.log("wrong number of markdown content " + markdownContentDBs.length)
    //     }
    //   }
    // )
  }
}
