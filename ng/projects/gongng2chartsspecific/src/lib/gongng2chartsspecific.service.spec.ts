import { TestBed } from '@angular/core/testing';

import { Gongng2chartsspecificService } from './gongng2chartsspecific.service';

describe('Gongng2chartsspecificService', () => {
  let service: Gongng2chartsspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(Gongng2chartsspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
