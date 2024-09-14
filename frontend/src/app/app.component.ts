import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { ButtonComponent } from './components/button/button.component';
import { Order, OrderService } from './services/orders.service';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    RouterOutlet,
    ButtonComponent,
  ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
})
export class AppComponent {
  title = 'frontend';
  url = `http://localhost:8080/ping`
  orderService = new OrderService(this.url)

  async ngOnInit() {
    let orders: Order[] = await this.orderService.getOrders()
  }
}
