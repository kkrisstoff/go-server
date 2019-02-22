import { Component, OnInit, Input } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Location } from '@angular/common';
import { Item } from '../item/item';
import { ItemService } from '../item.service'

@Component({
  selector: 'app-item-detail',
  templateUrl: './item-detail.component.html',
  styleUrls: ['./item-detail.component.css']
})
export class ItemDetailComponent implements OnInit {
  @Input() item: Item;
  
  constructor(
    private route: ActivatedRoute,
    private itemService: ItemService,
    private location: Location
  ) { }

  ngOnInit(): void {
    this.getHero();
  }

  getHero(): void {
    const id = +this.route.snapshot.paramMap.get('id');
    console.log(id);
    this.itemService.getItem(id)
      .subscribe(item => this.item = item);
  }

  save(): void {
    this.itemService.updateItem(this.item)
      .subscribe(() => this.goBack());
  }

  goBack(): void {
    this.location.back();
  }
}
