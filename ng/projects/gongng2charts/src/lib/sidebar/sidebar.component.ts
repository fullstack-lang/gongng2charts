import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { BehaviorSubject, Subscription } from 'rxjs';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbService } from '../commitnb.service'
import { GongstructSelectionService } from '../gongstruct-selection.service'

// insertion point for per struct import code
import { ChartConfigurationService } from '../chartconfiguration.service'
import { getChartConfigurationUniqueID } from '../front-repo.service'
import { DataPointService } from '../datapoint.service'
import { getDataPointUniqueID } from '../front-repo.service'
import { DatasetService } from '../dataset.service'
import { getDatasetUniqueID } from '../front-repo.service'
import { LabelService } from '../label.service'
import { getLabelUniqueID } from '../front-repo.service'

/**
 * Types of a GongNode / GongFlatNode
 */
export enum GongNodeType {
  STRUCT = "STRUCT",
  INSTANCE = "INSTANCE",
  ONE__ZERO_ONE_ASSOCIATION = 'ONE__ZERO_ONE_ASSOCIATION',
  ONE__ZERO_MANY_ASSOCIATION = 'ONE__ZERO_MANY_ASSOCIATION',
}

/**
 * GongNode is the "data" node
 */
interface GongNode {
  name: string; // if STRUCT, the name of the struct, if INSTANCE the name of the instance
  children: GongNode[];
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


/** 
 * GongFlatNode is the dynamic visual node with expandable and level information
 * */
interface GongFlatNode {
  expandable: boolean;
  name: string;
  level: number;
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


@Component({
  selector: 'app-gongng2charts-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css'],
})
export class SidebarComponent implements OnInit {

  /**
  * _transformer generated a displayed node from a data node
  *
  * @param node input data noe
  * @param level input level
  *
  * @returns an ExampleFlatNode
  */
  private _transformer = (node: GongNode, level: number) => {
    return {

      /**
      * in javascript, The !! ensures the resulting type is a boolean (true or false).
      *
      * !!node.children will evaluate to true is the variable is defined
      */
      expandable: !!node.children && node.children.length > 0,
      name: node.name,
      level: level,
      type: node.type,
      structName: node.structName,
      associationField: node.associationField,
      associatedStructName: node.associatedStructName,
      id: node.id,
      uniqueIdPerStack: node.uniqueIdPerStack,
    }
  }

  /**
   * treeControl is passed as the paramter treeControl in the "mat-tree" selector
   *
   * Flat tree control. Able to expand/collapse a subtree recursively for flattened tree.
   *
   * Construct with flat tree data node functions getLevel and isExpandable.
  constructor(
    getLevel: (dataNode: T) => number,
    isExpandable: (dataNode: T) => boolean, 
    options?: FlatTreeControlOptions<T, K> | undefined);
   */
  treeControl = new FlatTreeControl<GongFlatNode>(
    node => node.level,
    node => node.expandable
  );

  /**
   * from mat-tree documentation
   *
   * Tree flattener to convert a normal type of node to node with children & level information.
   */
  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children
  );

  /**
   * data is the other paramter to the "mat-tree" selector
   * 
   * strangely, the dataSource declaration has to follow the treeFlattener declaration
   */
  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  /**
   * hasChild is used by the selector for expandable nodes
   * 
   *  <mat-tree-node *matTreeNodeDef="let node;when: hasChild" matTreeNodePadding>
   * 
   * @param _ 
   * @param node 
   */
  hasChild = (_: number, node: GongFlatNode) => node.expandable;

  // front repo
  frontRepo: FrontRepo = new (FrontRepo)
  commitNb: number = 0

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  // SelectedStructChanged is the behavior subject that will emit
  // the selected gong struct whose table has to be displayed in the table outlet
  SelectedStructChanged: BehaviorSubject<string> = new BehaviorSubject("");

  subscription: Subscription = new Subscription

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private commitNbService: CommitNbService,
    private gongstructSelectionService: GongstructSelectionService,

    // insertion point for per struct service declaration
    private chartconfigurationService: ChartConfigurationService,
    private datapointService: DataPointService,
    private datasetService: DatasetService,
    private labelService: LabelService,
  ) { }

  ngOnDestroy() {
    // prevent memory leak when component destroyed
    this.subscription.unsubscribe();
  }

  ngOnInit(): void {

    this.subscription = this.gongstructSelectionService.gongtructSelected$.subscribe(
      gongstructName => {
        // console.log("sidebar gongstruct selected " + gongstructName)

        this.setTableRouterOutlet(gongstructName.toLowerCase() + "s")
      });

    this.refresh()

    this.SelectedStructChanged.subscribe(
      selectedStruct => {
        this.setTableRouterOutlet(selectedStruct)
      }
    )

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.chartconfigurationService.ChartConfigurationServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.datapointService.DataPointServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.datasetService.DatasetServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.labelService.LabelServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
  }

  refresh(): void {
    this.frontRepoService.pull().subscribe(frontRepo => {
      this.frontRepo = frontRepo

      // use of a Gödel number to uniquely identfy nodes : 2 * node.id + 3 * node.level
      let memoryOfExpandedNodes = new Map<number, boolean>()
      let nonInstanceNodeId = 1

      this.treeControl.dataNodes?.forEach(
        node => {
          if (this.treeControl.isExpanded(node)) {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, true)
          } else {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, false)
          }
        }
      )

      // reset the gong node tree
      this.gongNodeTree = new Array<GongNode>();

      // insertion point for per struct tree construction
      /**
      * fill up the ChartConfiguration part of the mat tree
      */
      let chartconfigurationGongNodeStruct: GongNode = {
        name: "ChartConfiguration",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "ChartConfiguration",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(chartconfigurationGongNodeStruct)

      this.frontRepo.ChartConfigurations_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.ChartConfigurations_array.forEach(
        chartconfigurationDB => {
          let chartconfigurationGongNodeInstance: GongNode = {
            name: chartconfigurationDB.Name,
            type: GongNodeType.INSTANCE,
            id: chartconfigurationDB.ID,
            uniqueIdPerStack: getChartConfigurationUniqueID(chartconfigurationDB.ID),
            structName: "ChartConfiguration",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          chartconfigurationGongNodeStruct.children!.push(chartconfigurationGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Datasets
          */
          let DatasetsGongNodeAssociation: GongNode = {
            name: "(Dataset) Datasets",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: chartconfigurationDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "ChartConfiguration",
            associationField: "Datasets",
            associatedStructName: "Dataset",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          chartconfigurationGongNodeInstance.children.push(DatasetsGongNodeAssociation)

          chartconfigurationDB.Datasets?.forEach(datasetDB => {
            let datasetNode: GongNode = {
              name: datasetDB.Name,
              type: GongNodeType.INSTANCE,
              id: datasetDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getChartConfigurationUniqueID(chartconfigurationDB.ID)
                + 11 * getDatasetUniqueID(datasetDB.ID),
              structName: "Dataset",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            DatasetsGongNodeAssociation.children.push(datasetNode)
          })

          /**
          * let append a node for the slide of pointer Labels
          */
          let LabelsGongNodeAssociation: GongNode = {
            name: "(Label) Labels",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: chartconfigurationDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "ChartConfiguration",
            associationField: "Labels",
            associatedStructName: "Label",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          chartconfigurationGongNodeInstance.children.push(LabelsGongNodeAssociation)

          chartconfigurationDB.Labels?.forEach(labelDB => {
            let labelNode: GongNode = {
              name: labelDB.Name,
              type: GongNodeType.INSTANCE,
              id: labelDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getChartConfigurationUniqueID(chartconfigurationDB.ID)
                + 11 * getLabelUniqueID(labelDB.ID),
              structName: "Label",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            LabelsGongNodeAssociation.children.push(labelNode)
          })

        }
      )

      /**
      * fill up the DataPoint part of the mat tree
      */
      let datapointGongNodeStruct: GongNode = {
        name: "DataPoint",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "DataPoint",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(datapointGongNodeStruct)

      this.frontRepo.DataPoints_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.DataPoints_array.forEach(
        datapointDB => {
          let datapointGongNodeInstance: GongNode = {
            name: datapointDB.Name,
            type: GongNodeType.INSTANCE,
            id: datapointDB.ID,
            uniqueIdPerStack: getDataPointUniqueID(datapointDB.ID),
            structName: "DataPoint",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          datapointGongNodeStruct.children!.push(datapointGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Dataset part of the mat tree
      */
      let datasetGongNodeStruct: GongNode = {
        name: "Dataset",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Dataset",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(datasetGongNodeStruct)

      this.frontRepo.Datasets_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Datasets_array.forEach(
        datasetDB => {
          let datasetGongNodeInstance: GongNode = {
            name: datasetDB.Name,
            type: GongNodeType.INSTANCE,
            id: datasetDB.ID,
            uniqueIdPerStack: getDatasetUniqueID(datasetDB.ID),
            structName: "Dataset",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          datasetGongNodeStruct.children!.push(datasetGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer DataPoints
          */
          let DataPointsGongNodeAssociation: GongNode = {
            name: "(DataPoint) DataPoints",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: datasetDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Dataset",
            associationField: "DataPoints",
            associatedStructName: "DataPoint",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          datasetGongNodeInstance.children.push(DataPointsGongNodeAssociation)

          datasetDB.DataPoints?.forEach(datapointDB => {
            let datapointNode: GongNode = {
              name: datapointDB.Name,
              type: GongNodeType.INSTANCE,
              id: datapointDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getDatasetUniqueID(datasetDB.ID)
                + 11 * getDataPointUniqueID(datapointDB.ID),
              structName: "DataPoint",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            DataPointsGongNodeAssociation.children.push(datapointNode)
          })

        }
      )

      /**
      * fill up the Label part of the mat tree
      */
      let labelGongNodeStruct: GongNode = {
        name: "Label",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Label",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(labelGongNodeStruct)

      this.frontRepo.Labels_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Labels_array.forEach(
        labelDB => {
          let labelGongNodeInstance: GongNode = {
            name: labelDB.Name,
            type: GongNodeType.INSTANCE,
            id: labelDB.ID,
            uniqueIdPerStack: getLabelUniqueID(labelDB.ID),
            structName: "Label",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          labelGongNodeStruct.children!.push(labelGongNodeInstance)

          // insertion point for per field code
        }
      )


      this.dataSource.data = this.gongNodeTree

      // expand nodes that were exapanded before
      this.treeControl.dataNodes?.forEach(
        node => {
          if (memoryOfExpandedNodes.get(node.uniqueIdPerStack)) {
            this.treeControl.expand(node)
          }
        }
      )
    });

    // fetch the number of commits
    this.commitNbService.getCommitNb().subscribe(
      commitNb => {
        this.commitNb = commitNb
      }
    )
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongng2charts_go_table: ["github_com_fullstack_lang_gongng2charts_go-" + path]
      }
    }]);
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutletFromTree(path: string, type: GongNodeType, structName: string, id: number) {

    if (type == GongNodeType.STRUCT) {
      this.router.navigate([{
        outlets: {
          github_com_fullstack_lang_gongng2charts_go_table: ["github_com_fullstack_lang_gongng2charts_go-" + path.toLowerCase()]
        }
      }]);
    }

    if (type == GongNodeType.INSTANCE) {
      this.router.navigate([{
        outlets: {
          github_com_fullstack_lang_gongng2charts_go_presentation: ["github_com_fullstack_lang_gongng2charts_go-" + structName.toLowerCase() + "-presentation", id]
        }
      }]);
    }
  }

  setEditorRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongng2charts_go_editor: ["github_com_fullstack_lang_gongng2charts_go-" + path.toLowerCase()]
      }
    }]);
  }

  setEditorSpecialRouterOutlet(node: GongFlatNode) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongng2charts_go_editor: ["github_com_fullstack_lang_gongng2charts_go-" + node.associatedStructName.toLowerCase() + "-adder", node.id, node.structName, node.associationField]
      }
    }]);
  }
}
