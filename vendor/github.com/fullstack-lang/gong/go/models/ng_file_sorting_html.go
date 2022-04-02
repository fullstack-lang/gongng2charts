package models

const NgSortingTemplateHTML = `<div cdkDropList class="example-list" (cdkDropListDropped)="drop($event)">
<div class="example-box" *ngFor="let {{structname}} of associated{{Structname}}s" cdkDrag>{{{{structname}}.Name}}</div>
</div>
<button class="table__save" color="primary" mat-raised-button (click)="save()">
Save
</button>`
