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

import { UmlscDB } from './umlsc-db';

// insertion point for imports
import { DiagramPackageDB } from './diagrampackage-db'

@Injectable({
  providedIn: 'root'
})
export class UmlscService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  UmlscServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private umlscsUrl: string

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
    this.umlscsUrl = origin + '/api/github.com/fullstack-lang/gongdoc/go/v1/umlscs';
  }

  /** GET umlscs from the server */
  getUmlscs(): Observable<UmlscDB[]> {
    return this.http.get<UmlscDB[]>(this.umlscsUrl)
      .pipe(
        tap(_ => this.log('fetched umlscs')),
        catchError(this.handleError<UmlscDB[]>('getUmlscs', []))
      );
  }

  /** GET umlsc by id. Will 404 if id not found */
  getUmlsc(id: number): Observable<UmlscDB> {
    const url = `${this.umlscsUrl}/${id}`;
    return this.http.get<UmlscDB>(url).pipe(
      tap(_ => this.log(`fetched umlsc id=${id}`)),
      catchError(this.handleError<UmlscDB>(`getUmlsc id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new umlsc to the server */
  postUmlsc(umlscdb: UmlscDB): Observable<UmlscDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    umlscdb.States = []
    let _DiagramPackage_Umlscs_reverse = umlscdb.DiagramPackage_Umlscs_reverse
    umlscdb.DiagramPackage_Umlscs_reverse = new DiagramPackageDB

    return this.http.post<UmlscDB>(this.umlscsUrl, umlscdb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        umlscdb.DiagramPackage_Umlscs_reverse = _DiagramPackage_Umlscs_reverse
        this.log(`posted umlscdb id=${umlscdb.ID}`)
      }),
      catchError(this.handleError<UmlscDB>('postUmlsc'))
    );
  }

  /** DELETE: delete the umlscdb from the server */
  deleteUmlsc(umlscdb: UmlscDB | number): Observable<UmlscDB> {
    const id = typeof umlscdb === 'number' ? umlscdb : umlscdb.ID;
    const url = `${this.umlscsUrl}/${id}`;

    return this.http.delete<UmlscDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted umlscdb id=${id}`)),
      catchError(this.handleError<UmlscDB>('deleteUmlsc'))
    );
  }

  /** PUT: update the umlscdb on the server */
  updateUmlsc(umlscdb: UmlscDB): Observable<UmlscDB> {
    const id = typeof umlscdb === 'number' ? umlscdb : umlscdb.ID;
    const url = `${this.umlscsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    umlscdb.States = []
    let _DiagramPackage_Umlscs_reverse = umlscdb.DiagramPackage_Umlscs_reverse
    umlscdb.DiagramPackage_Umlscs_reverse = new DiagramPackageDB

    return this.http.put<UmlscDB>(url, umlscdb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        umlscdb.DiagramPackage_Umlscs_reverse = _DiagramPackage_Umlscs_reverse
        this.log(`updated umlscdb id=${umlscdb.ID}`)
      }),
      catchError(this.handleError<UmlscDB>('updateUmlsc'))
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
