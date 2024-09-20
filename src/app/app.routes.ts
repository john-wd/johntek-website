import { Routes } from '@angular/router';
import { AboutComponent } from './about/about.component';
import { HomepageComponent } from './homepage/homepage.component';

export const routes: Routes = [
  { path: "about", component: AboutComponent },
  { path: "home", component: HomepageComponent },
  // { path: "", redirectTo: "/home" },
];
