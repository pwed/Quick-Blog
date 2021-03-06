import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {

  constructor(private http: HttpClient, private route: ActivatedRoute) { }

  PostCount = '';
  PostList: string[];

  ngOnInit(): void {
    const url = '/api/posts/' + 50;
    this.http.get(url).subscribe(
      data => {
        const ids: any = data;
        this.PostCount = ids.length;
        this.PostList = new Array(ids.length);
        let pos = 0;
        ids.forEach(element => {
          this.PostList[pos] = ids[pos];
          pos++;
        });
      }
    );
  }

}
