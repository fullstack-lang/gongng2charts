import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent {

  // choices for the top radio button
  view = 'Chart 1 view'
  default = 'Default view'
  diagrams = 'Diagrams view'
  meta = 'Meta view'
  chart1 = "Chart 1 view"
  chart2 = "Chart 2 view"
  views: string[] = [this.chart1, this.chart2, this.default, this.diagrams, this.meta];
}
