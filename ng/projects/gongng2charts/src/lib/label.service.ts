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

import { LabelDB } from './label-db';

// insertion point for imports
import { ChartConfigurationDB } from './chartconfiguration-db'

@Injectable({
  providedIn: 'root'
})
export class LabelService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  LabelServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private labelsUrl: string

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
    this.labelsUrl = origin + '/api/github.com/fullstack-lang/gongng2charts/go/v1/labels';
  }

  /** GET labels from the server */
  getLabels(): Observable<LabelDB[]> {
    return this.http.get<LabelDB[]>(this.labelsUrl)
      .pipe(
        tap(_ => this.log('fetched labels')),
        catchError(this.handleError<LabelDB[]>('getLabels', []))
      );
  }

  /** GET label by id. Will 404 if id not found */
  getLabel(id: number): Observable<LabelDB> {
    const url = `${this.labelsUrl}/${id}`;
    return this.http.get<LabelDB>(url).pipe(
      tap(_ => this.log(`fetched label id=${id}`)),
      catchError(this.handleError<LabelDB>(`getLabel id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new label to the server */
  postLabel(labeldb: LabelDB): Observable<LabelDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let _ChartConfiguration_Labels_reverse = labeldb.ChartConfiguration_Labels_reverse
    labeldb.ChartConfiguration_Labels_reverse = new ChartConfigurationDB

    return this.http.post<LabelDB>(this.labelsUrl, labeldb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        labeldb.ChartConfiguration_Labels_reverse = _ChartConfiguration_Labels_reverse
        this.log(`posted labeldb id=${labeldb.ID}`)
      }),
      catchError(this.handleError<LabelDB>('postLabel'))
    );
  }

  /** DELETE: delete the labeldb from the server */
  deleteLabel(labeldb: LabelDB | number): Observable<LabelDB> {
    const id = typeof labeldb === 'number' ? labeldb : labeldb.ID;
    const url = `${this.labelsUrl}/${id}`;

    return this.http.delete<LabelDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted labeldb id=${id}`)),
      catchError(this.handleError<LabelDB>('deleteLabel'))
    );
  }

  /** PUT: update the labeldb on the server */
  updateLabel(labeldb: LabelDB): Observable<LabelDB> {
    const id = typeof labeldb === 'number' ? labeldb : labeldb.ID;
    const url = `${this.labelsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let _ChartConfiguration_Labels_reverse = labeldb.ChartConfiguration_Labels_reverse
    labeldb.ChartConfiguration_Labels_reverse = new ChartConfigurationDB

    return this.http.put<LabelDB>(url, labeldb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        labeldb.ChartConfiguration_Labels_reverse = _ChartConfiguration_Labels_reverse
        this.log(`updated labeldb id=${labeldb.ID}`)
      }),
      catchError(this.handleError<LabelDB>('updateLabel'))
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