import { Component, Input } from '@angular/core';
import { HlmButtonDirective } from '@spartan-ng/ui-button-helm';

@Component({
  selector: 'app-button',
  standalone: true,
  imports: [HlmButtonDirective],
  templateUrl: './button.component.html',
})
export class ButtonComponent {
  @Input()
  name: string = "Button";
}
