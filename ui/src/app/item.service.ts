import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';
import { MessageService } from './message.service';
import { Item } from './item/item';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable({providedIn: 'root'})
export class ItemService {

  constructor(
    private http: HttpClient,
    private messageService: MessageService
  ) { }

  private apiPrefixl = 'api'

  private log(message: string) {
    this.messageService.add(`ItemService: ${message}`);
  }

  /**
   * GET items
   */
  getItems(): Observable<Item[]> {
    return this.http.get<Item[]>(`${this.apiPrefixl}/items`)
    .pipe(
      tap(_ => this.log('fetched Items')),
      catchError(this.handleError('getItems', []))
    );
  }

  getItem(id: number): Observable<Item> {
    const url = `${this.apiPrefixl}/getItem/${id}`;
    return this.http.get<Item>(url)
    .pipe(
      tap(_ => this.log(`fetched item id=${id}`)),
      catchError(this.handleError<Item>(`getItem id=${id}`))
    );
 }

  updateItem (item: Item): Observable<any> {
    return this.http.put(`${this.apiPrefixl}/updateItem`, item, httpOptions)
      .pipe(
        tap(_ => this.log(`updated item id=${item.id}`)),
        catchError(this.handleError<any>('updateItem'))
      );
  }

  /** POST: add a new hero to the server */
  addItem (item: Item): Observable<Item> {
    return this.http.post<Item>(`${this.apiPrefixl}/addItem`, item, httpOptions).pipe(
      tap((item: Item) => this.log(`added hero w/ id=${item.id}`)),
      catchError(this.handleError<Item>('addItem'))
    );
  }

  /** DELETE: delete the hero from the server */
  deleteItem (item: Item | number): Observable<Item> {
    const id = typeof item === 'number' ? item : item.id;
    const url = `${this.apiPrefixl}/deleteItem/${id}`;

    return this.http.delete<Item>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted item id=${id}`)),
      catchError(this.handleError<Item>('deleteItem'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T> (operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead
  
      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);
  
      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }
}
