import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet],
  template: `
    <h1>
      Home!
    </h1>
  `,
  styles: ``
})
export class AppComponent {
  title = 'frontend';
}
