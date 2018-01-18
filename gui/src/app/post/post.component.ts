import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.scss']
})
export class PostComponent implements OnInit {

  constructor(private http: HttpClient) { }

  ngOnInit(): void {
    // let url = "http://localhost:3000/api/post/1";
    this.http.get("http://localhost:3000/api/post/1").subscribe(
      data => {
        console.log(data)
      }
    )
  }
}
