import { Component, OnInit } from '@angular/core';
import { NewPost } from '../new-post';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-new-post',
  templateUrl: './new-post.component.html',
  styleUrls: ['./new-post.component.scss']
})

export class NewPostComponent implements OnInit {

  model = new NewPost(
    "fred",
    "Add post content here ..."
  )

  newPostLink: string

  constructor(private http: HttpClient) { }

  ngOnInit() {
  }

  sendPost() {
    console.log("sending data")
    console.log(this.model)
    this.http.post("/api/newpost", 
      this.model).subscribe(
      data => {
        console.log(data)
        let response: any = data 
        this.newPostLink = `/post/` + response.newID
      }
    )
  }

}
