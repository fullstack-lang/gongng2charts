// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { GongBasicFieldDB } from './gongbasicfield-db';

// insertion point for imports
import { GongEnumDB } from './gongenum-db'
import { GongStructDB } from './gongstruct-db'

@Injectable({
  providedIn: 'root'
})
export class GongBasicFieldService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  GongBasicFieldServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private gongbasicfieldsUrl: string

  constructor(
    private http: HttpClient,
    private location: Location,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.gongbasicfieldsUrl = origin + '/api/github.com/fullstack-lang/gong/go/v1/gongbasicfields';
  }

  /** GET gongbasicfields from the server */
  getGongBasicFields(): Observable<GongBasicFieldDB[]> {
    return this.http.get<GongBasicFieldDB[]>(this.gongbasicfieldsUrl)
      .pipe(
        tap(_ => this.log('fetched gongbasicfields')),
        catchError(this.handleError<GongBasicFieldDB[]>('getGongBasicFields', []))
      );
  }

  /** GET gongbasicfield by id. Will 404 if id not found */
  getGongBasicField(id: number): Observable<GongBasicFieldDB> {
    const url = `${this.gongbasicfieldsUrl}/${id}`;
    return this.http.get<GongBasicFieldDB>(url).pipe(
      tap(_ => this.log(`fetched gongbasicfield id=${id}`)),
      catchError(this.handleError<GongBasicFieldDB>(`getGongBasicField id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new gongbasicfield to the server */
  postGongBasicField(gongbasicfielddb: GongBasicFieldDB): Observable<GongBasicFieldDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    gongbasicfielddb.GongEnum = new GongEnumDB
    let _GongStruct_GongBasicFields_reverse = gongbasicfielddb.GongStruct_GongBasicFields_reverse
    gongbasicfielddb.GongStruct_GongBasicFields_reverse = new GongStructDB

    return this.http.post<GongBasicFieldDB>(this.gongbasicfieldsUrl, gongbasicfielddb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        gongbasicfielddb.GongStruct_GongBasicFields_reverse = _GongStruct_GongBasicFields_reverse
        this.log(`posted gongbasicfielddb id=${gongbasicfielddb.ID}`)
      }),
      catchError(this.handleError<GongBasicFieldDB>('postGongBasicField'))
    );
  }

  /** DELETE: delete the gongbasicfielddb from the server */
  deleteGongBasicField(gongbasicfielddb: GongBasicFieldDB | number): Observable<GongBasicFieldDB> {
    const id = typeof gongbasicfielddb === 'number' ? gongbasicfielddb : gongbasicfielddb.ID;
    const url = `${this.gongbasicfieldsUrl}/${id}`;

    return this.http.delete<GongBasicFieldDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted gongbasicfielddb id=${id}`)),
      catchError(this.handleError<GongBasicFieldDB>('deleteGongBasicField'))
    );
  }

  /** PUT: update the gongbasicfielddb on the server */
  updateGongBasicField(gongbasicfielddb: GongBasicFieldDB): Observable<GongBasicFieldDB> {
    const id = typeof gongbasicfielddb === 'number' ? gongbasicfielddb : gongbasicfielddb.ID;
    const url = `${this.gongbasicfieldsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    gongbasicfielddb.GongEnum = new GongEnumDB
    let _GongStruct_GongBasicFields_reverse = gongbasicfielddb.GongStruct_GongBasicFields_reverse
    gongbasicfielddb.GongStruct_GongBasicFields_reverse = new GongStructDB

    return this.http.put<GongBasicFieldDB>(url, gongbasicfielddb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        gongbasicfielddb.GongStruct_GongBasicFields_reverse = _GongStruct_GongBasicFields_reverse
        this.log(`updated gongbasicfielddb id=${gongbasicfielddb.ID}`)
      }),
      catchError(this.handleError<GongBasicFieldDB>('updateGongBasicField'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {

  }
}
