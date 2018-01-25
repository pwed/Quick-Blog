import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.scss']
})
export class PostComponent implements OnInit {

  constructor(private http: HttpClient, private route: ActivatedRoute) { }

  PostContent: string = "Loading ..."

  ngOnInit(): void {
    let url = "http://localhost:3000/api/post/" + this.route.snapshot.params.id;
    this.http.get(url).subscribe(
      data => {
        let post: any = data;
        this.PostContent = post.Body;
      }
    )
  }
}
