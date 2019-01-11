import { Component, OnInit } from '@angular/core';
import { Item } from '../item/item';
import { ItemService } from '../item.service';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.css']
})

export class ListComponent implements OnInit {
  items: Item[];

  constructor(private itemService: ItemService) { }

  ngOnInit() {
    this.getItems();
  }

  getItems(): void {
    this.itemService.getItems()
      .subscribe(items => this.items = items);
  }

  add(name: string, message: string): void {
    name = name.trim();
    message = message.trim();
    if (!name) { return; }
    this.itemService.addItem({ name, message } as Item)
      .subscribe(item => {
        this.items.push(item);
      });
  }

  delete(item: Item): void {
    this.items = this.items.filter(i => i !== item);
    this.itemService.deleteItem(item).subscribe();
  }
}
