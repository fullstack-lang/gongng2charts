import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Gongng2chartsChartComponent } from './gongng2charts-chart.component';

describe('Gongng2chartsChartComponent', () => {
  let component: Gongng2chartsChartComponent;
  let fixture: ComponentFixture<Gongng2chartsChartComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ Gongng2chartsChartComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(Gongng2chartsChartComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
