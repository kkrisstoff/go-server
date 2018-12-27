import { Component, OnInit } from '@angular/core';
import { Item } from './item'

@Component({
  selector: 'app-item',
  templateUrl: './item.component.html',
  styleUrls: ['./item.component.css']
})
export class ItemComponent implements OnInit {
  item: Item = {
    id: 1,
    name: 'Item\'s Name',
    message: 'Some message'
  };
  
  constructor() { }

  ngOnInit() {
  }

}
