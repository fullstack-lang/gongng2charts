import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Gongng2chartsspecificComponent } from './gongng2chartsspecific.component';

describe('Gongng2chartsspecificComponent', () => {
  let component: Gongng2chartsspecificComponent;
  let fixture: ComponentFixture<Gongng2chartsspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ Gongng2chartsspecificComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(Gongng2chartsspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
