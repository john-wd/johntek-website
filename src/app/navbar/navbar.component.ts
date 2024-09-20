import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { RouterModule } from '@angular/router';
import { MenubarModule } from 'primeng/menubar';


@Component({
  selector: 'app-navbar',
  standalone: true,
  imports: [CommonModule, MenubarModule, RouterModule],
  templateUrl: './navbar.component.html',
  styleUrl: './navbar.component.scss'
})
export class NavbarComponent {
  // constructor(private router: Router) { }
  items = [
    {
      label: "Home",
      route: "/home"
    },
    {
      label: "Projects",
      route: "/home",
      fragment: "projects"
    },
    { label: "About", route: "/about" }
  ]
}
