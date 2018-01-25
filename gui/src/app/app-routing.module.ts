import { NgModule } from '@angular/core';
import { Routes, RouterModule, RouterLinkActive } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { PostComponent } from './post/post.component';
import { NewPostComponent } from './new-post/new-post.component';

const routes: Routes = [
  {
    path: '',
    component: HomeComponent
  },
  {
    path: 'post/:id',
    component: PostComponent
  },
  {
    path: 'new-post',
    component: NewPostComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
