import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent {

  // choices for the top radio button
  view = 'Chart view'
  default = 'Default view'
  diagrams = 'Diagrams view'
  meta = 'Meta view'
  chart = "Chart view"
  views: string[] = [this.chart, this.default, this.diagrams, this.meta];
}
