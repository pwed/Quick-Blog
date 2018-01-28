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

  PostContent = 'Loading ...';

  ngOnInit(): void {
    const url = '/api/post/' + this.route.snapshot.params.id;
    this.http.get(url).subscribe(
      data => {
        const post: any = data;
        this.PostContent = post.Body.HTML;
      }
    );
  }
}
