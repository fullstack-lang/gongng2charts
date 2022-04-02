import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { ChartConfigurationDB } from './chartconfiguration-db'
import { ChartConfigurationService } from './chartconfiguration.service'

import { DataPointDB } from './datapoint-db'
import { DataPointService } from './datapoint.service'

import { DatasetDB } from './dataset-db'
import { DatasetService } from './dataset.service'

import { LabelDB } from './label-db'
import { LabelService } from './label.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  ChartConfigurations_array = new Array<ChartConfigurationDB>(); // array of repo instances
  ChartConfigurations = new Map<number, ChartConfigurationDB>(); // map of repo instances
  ChartConfigurations_batch = new Map<number, ChartConfigurationDB>(); // same but only in last GET (for finding repo instances to delete)
  DataPoints_array = new Array<DataPointDB>(); // array of repo instances
  DataPoints = new Map<number, DataPointDB>(); // map of repo instances
  DataPoints_batch = new Map<number, DataPointDB>(); // same but only in last GET (for finding repo instances to delete)
  Datasets_array = new Array<DatasetDB>(); // array of repo instances
  Datasets = new Map<number, DatasetDB>(); // map of repo instances
  Datasets_batch = new Map<number, DatasetDB>(); // same but only in last GET (for finding repo instances to delete)
  Labels_array = new Array<LabelDB>(); // array of repo instances
  Labels = new Map<number, LabelDB>(); // map of repo instances
  Labels_batch = new Map<number, LabelDB>(); // same but only in last GET (for finding repo instances to delete)
}

//
// Store of all instances of the stack
//
export const FrontRepoSingloton = new (FrontRepo)

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
  ID: number = 0 // ID of the calling instance

  // the reverse pointer is the name of the generated field on the destination
  // struct of the ONE-MANY association
  ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean = false // if true, this is for ordering items

  // there are different selection mode : ONE_MANY or MANY_MANY
  SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

  // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
  //
  // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
  // 
  // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
  // at the end of the ONE-MANY association
  SourceStruct: string = ""  // The "Aclass"
  SourceField: string = "" // the "AnarrayofbUse"
  IntermediateStruct: string = "" // the "AclassBclassUse" 
  IntermediateStructField: string = "" // the "Bclass" as field
  NextAssociationStruct: string = "" // the "Bclass"
}

export enum SelectionMode {
  ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
  MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient, // insertion point sub template 
    private chartconfigurationService: ChartConfigurationService,
    private datapointService: DataPointService,
    private datasetService: DatasetService,
    private labelService: LabelService,
  ) { }

  // postService provides a post function for each struct name
  postService(structName: string, instanceToBePosted: any) {
    let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
    let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

    servicePostFunction(instanceToBePosted).subscribe(
      instance => {
        let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("post")
      }
    );
  }

  // deleteService provides a delete function for each struct name
  deleteService(structName: string, instanceToBeDeleted: any) {
    let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
    let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

    serviceDeleteFunction(instanceToBeDeleted).subscribe(
      instance => {
        let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("delete")
      }
    );
  }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<ChartConfigurationDB[]>,
    Observable<DataPointDB[]>,
    Observable<DatasetDB[]>,
    Observable<LabelDB[]>,
  ] = [ // insertion point sub template 
      this.chartconfigurationService.getChartConfigurations(),
      this.datapointService.getDataPoints(),
      this.datasetService.getDatasets(),
      this.labelService.getLabels(),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            chartconfigurations_,
            datapoints_,
            datasets_,
            labels_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var chartconfigurations: ChartConfigurationDB[]
            chartconfigurations = chartconfigurations_ as ChartConfigurationDB[]
            var datapoints: DataPointDB[]
            datapoints = datapoints_ as DataPointDB[]
            var datasets: DatasetDB[]
            datasets = datasets_ as DatasetDB[]
            var labels: LabelDB[]
            labels = labels_ as LabelDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.ChartConfigurations_array = chartconfigurations

            // clear the map that counts ChartConfiguration in the GET
            FrontRepoSingloton.ChartConfigurations_batch.clear()

            chartconfigurations.forEach(
              chartconfiguration => {
                FrontRepoSingloton.ChartConfigurations.set(chartconfiguration.ID, chartconfiguration)
                FrontRepoSingloton.ChartConfigurations_batch.set(chartconfiguration.ID, chartconfiguration)
              }
            )

            // clear chartconfigurations that are absent from the batch
            FrontRepoSingloton.ChartConfigurations.forEach(
              chartconfiguration => {
                if (FrontRepoSingloton.ChartConfigurations_batch.get(chartconfiguration.ID) == undefined) {
                  FrontRepoSingloton.ChartConfigurations.delete(chartconfiguration.ID)
                }
              }
            )

            // sort ChartConfigurations_array array
            FrontRepoSingloton.ChartConfigurations_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.DataPoints_array = datapoints

            // clear the map that counts DataPoint in the GET
            FrontRepoSingloton.DataPoints_batch.clear()

            datapoints.forEach(
              datapoint => {
                FrontRepoSingloton.DataPoints.set(datapoint.ID, datapoint)
                FrontRepoSingloton.DataPoints_batch.set(datapoint.ID, datapoint)
              }
            )

            // clear datapoints that are absent from the batch
            FrontRepoSingloton.DataPoints.forEach(
              datapoint => {
                if (FrontRepoSingloton.DataPoints_batch.get(datapoint.ID) == undefined) {
                  FrontRepoSingloton.DataPoints.delete(datapoint.ID)
                }
              }
            )

            // sort DataPoints_array array
            FrontRepoSingloton.DataPoints_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Datasets_array = datasets

            // clear the map that counts Dataset in the GET
            FrontRepoSingloton.Datasets_batch.clear()

            datasets.forEach(
              dataset => {
                FrontRepoSingloton.Datasets.set(dataset.ID, dataset)
                FrontRepoSingloton.Datasets_batch.set(dataset.ID, dataset)
              }
            )

            // clear datasets that are absent from the batch
            FrontRepoSingloton.Datasets.forEach(
              dataset => {
                if (FrontRepoSingloton.Datasets_batch.get(dataset.ID) == undefined) {
                  FrontRepoSingloton.Datasets.delete(dataset.ID)
                }
              }
            )

            // sort Datasets_array array
            FrontRepoSingloton.Datasets_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Labels_array = labels

            // clear the map that counts Label in the GET
            FrontRepoSingloton.Labels_batch.clear()

            labels.forEach(
              label => {
                FrontRepoSingloton.Labels.set(label.ID, label)
                FrontRepoSingloton.Labels_batch.set(label.ID, label)
              }
            )

            // clear labels that are absent from the batch
            FrontRepoSingloton.Labels.forEach(
              label => {
                if (FrontRepoSingloton.Labels_batch.get(label.ID) == undefined) {
                  FrontRepoSingloton.Labels.delete(label.ID)
                }
              }
            )

            // sort Labels_array array
            FrontRepoSingloton.Labels_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            chartconfigurations.forEach(
              chartconfiguration => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            datapoints.forEach(
              datapoint => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Dataset.DataPoints redeeming
                {
                  let _dataset = FrontRepoSingloton.Datasets.get(datapoint.Dataset_DataPointsDBID.Int64)
                  if (_dataset) {
                    if (_dataset.DataPoints == undefined) {
                      _dataset.DataPoints = new Array<DataPointDB>()
                    }
                    _dataset.DataPoints.push(datapoint)
                    if (datapoint.Dataset_DataPoints_reverse == undefined) {
                      datapoint.Dataset_DataPoints_reverse = _dataset
                    }
                  }
                }
              }
            )
            datasets.forEach(
              dataset => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field ChartConfiguration.Datasets redeeming
                {
                  let _chartconfiguration = FrontRepoSingloton.ChartConfigurations.get(dataset.ChartConfiguration_DatasetsDBID.Int64)
                  if (_chartconfiguration) {
                    if (_chartconfiguration.Datasets == undefined) {
                      _chartconfiguration.Datasets = new Array<DatasetDB>()
                    }
                    _chartconfiguration.Datasets.push(dataset)
                    if (dataset.ChartConfiguration_Datasets_reverse == undefined) {
                      dataset.ChartConfiguration_Datasets_reverse = _chartconfiguration
                    }
                  }
                }
              }
            )
            labels.forEach(
              label => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field ChartConfiguration.Labels redeeming
                {
                  let _chartconfiguration = FrontRepoSingloton.ChartConfigurations.get(label.ChartConfiguration_LabelsDBID.Int64)
                  if (_chartconfiguration) {
                    if (_chartconfiguration.Labels == undefined) {
                      _chartconfiguration.Labels = new Array<LabelDB>()
                    }
                    _chartconfiguration.Labels.push(label)
                    if (label.ChartConfiguration_Labels_reverse == undefined) {
                      label.ChartConfiguration_Labels_reverse = _chartconfiguration
                    }
                  }
                }
              }
            )

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // ChartConfigurationPull performs a GET on ChartConfiguration of the stack and redeem association pointers 
  ChartConfigurationPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.chartconfigurationService.getChartConfigurations()
        ]).subscribe(
          ([ // insertion point sub template 
            chartconfigurations,
          ]) => {
            // init the array
            FrontRepoSingloton.ChartConfigurations_array = chartconfigurations

            // clear the map that counts ChartConfiguration in the GET
            FrontRepoSingloton.ChartConfigurations_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            chartconfigurations.forEach(
              chartconfiguration => {
                FrontRepoSingloton.ChartConfigurations.set(chartconfiguration.ID, chartconfiguration)
                FrontRepoSingloton.ChartConfigurations_batch.set(chartconfiguration.ID, chartconfiguration)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear chartconfigurations that are absent from the GET
            FrontRepoSingloton.ChartConfigurations.forEach(
              chartconfiguration => {
                if (FrontRepoSingloton.ChartConfigurations_batch.get(chartconfiguration.ID) == undefined) {
                  FrontRepoSingloton.ChartConfigurations.delete(chartconfiguration.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // DataPointPull performs a GET on DataPoint of the stack and redeem association pointers 
  DataPointPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.datapointService.getDataPoints()
        ]).subscribe(
          ([ // insertion point sub template 
            datapoints,
          ]) => {
            // init the array
            FrontRepoSingloton.DataPoints_array = datapoints

            // clear the map that counts DataPoint in the GET
            FrontRepoSingloton.DataPoints_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            datapoints.forEach(
              datapoint => {
                FrontRepoSingloton.DataPoints.set(datapoint.ID, datapoint)
                FrontRepoSingloton.DataPoints_batch.set(datapoint.ID, datapoint)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Dataset.DataPoints redeeming
                {
                  let _dataset = FrontRepoSingloton.Datasets.get(datapoint.Dataset_DataPointsDBID.Int64)
                  if (_dataset) {
                    if (_dataset.DataPoints == undefined) {
                      _dataset.DataPoints = new Array<DataPointDB>()
                    }
                    _dataset.DataPoints.push(datapoint)
                    if (datapoint.Dataset_DataPoints_reverse == undefined) {
                      datapoint.Dataset_DataPoints_reverse = _dataset
                    }
                  }
                }
              }
            )

            // clear datapoints that are absent from the GET
            FrontRepoSingloton.DataPoints.forEach(
              datapoint => {
                if (FrontRepoSingloton.DataPoints_batch.get(datapoint.ID) == undefined) {
                  FrontRepoSingloton.DataPoints.delete(datapoint.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // DatasetPull performs a GET on Dataset of the stack and redeem association pointers 
  DatasetPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.datasetService.getDatasets()
        ]).subscribe(
          ([ // insertion point sub template 
            datasets,
          ]) => {
            // init the array
            FrontRepoSingloton.Datasets_array = datasets

            // clear the map that counts Dataset in the GET
            FrontRepoSingloton.Datasets_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            datasets.forEach(
              dataset => {
                FrontRepoSingloton.Datasets.set(dataset.ID, dataset)
                FrontRepoSingloton.Datasets_batch.set(dataset.ID, dataset)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field ChartConfiguration.Datasets redeeming
                {
                  let _chartconfiguration = FrontRepoSingloton.ChartConfigurations.get(dataset.ChartConfiguration_DatasetsDBID.Int64)
                  if (_chartconfiguration) {
                    if (_chartconfiguration.Datasets == undefined) {
                      _chartconfiguration.Datasets = new Array<DatasetDB>()
                    }
                    _chartconfiguration.Datasets.push(dataset)
                    if (dataset.ChartConfiguration_Datasets_reverse == undefined) {
                      dataset.ChartConfiguration_Datasets_reverse = _chartconfiguration
                    }
                  }
                }
              }
            )

            // clear datasets that are absent from the GET
            FrontRepoSingloton.Datasets.forEach(
              dataset => {
                if (FrontRepoSingloton.Datasets_batch.get(dataset.ID) == undefined) {
                  FrontRepoSingloton.Datasets.delete(dataset.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // LabelPull performs a GET on Label of the stack and redeem association pointers 
  LabelPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.labelService.getLabels()
        ]).subscribe(
          ([ // insertion point sub template 
            labels,
          ]) => {
            // init the array
            FrontRepoSingloton.Labels_array = labels

            // clear the map that counts Label in the GET
            FrontRepoSingloton.Labels_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            labels.forEach(
              label => {
                FrontRepoSingloton.Labels.set(label.ID, label)
                FrontRepoSingloton.Labels_batch.set(label.ID, label)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field ChartConfiguration.Labels redeeming
                {
                  let _chartconfiguration = FrontRepoSingloton.ChartConfigurations.get(label.ChartConfiguration_LabelsDBID.Int64)
                  if (_chartconfiguration) {
                    if (_chartconfiguration.Labels == undefined) {
                      _chartconfiguration.Labels = new Array<LabelDB>()
                    }
                    _chartconfiguration.Labels.push(label)
                    if (label.ChartConfiguration_Labels_reverse == undefined) {
                      label.ChartConfiguration_Labels_reverse = _chartconfiguration
                    }
                  }
                }
              }
            )

            // clear labels that are absent from the GET
            FrontRepoSingloton.Labels.forEach(
              label => {
                if (FrontRepoSingloton.Labels_batch.get(label.ID) == undefined) {
                  FrontRepoSingloton.Labels.delete(label.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getChartConfigurationUniqueID(id: number): number {
  return 31 * id
}
export function getDataPointUniqueID(id: number): number {
  return 37 * id
}
export function getDatasetUniqueID(id: number): number {
  return 41 * id
}
export function getLabelUniqueID(id: number): number {
  return 43 * id
}
