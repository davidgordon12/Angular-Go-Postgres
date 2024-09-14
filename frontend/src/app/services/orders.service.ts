export interface Order {
  Id: Number,
  Created_on: String,
  Price: Number,
  Category: String,
}

export class OrderService {
  url: String;

  constructor(url: String) { this.url = url }

  public async getOrders(): Promise<Order[]> {
    let response: Response = await fetch(`${this.url}`)
    return await response.json() as Order[]
  }
}
