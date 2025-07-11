export interface Order {
  id: number;
  customerName: string;
  status: "Dalam Proses" | "selesai" | "Siap Diambil";
  price: number;
}

const STORAGE_KEY = "laundry_orders";

export const getOrders = (): Order[] => {
  return JSON.parse(localStorage.getItem(STORAGE_KEY) || "[]");
};

export const saveOrders = (orders: Order[]) => {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(orders));
};

export const addOrder = (order: Order) => {
  const orders = getOrders();
  orders.push(order);
  saveOrders(orders);
};

export const updateOrder = (updatedOrder: Order) => {
  let orders = getOrders();
  orders = orders.map((order) =>
    order.id === updatedOrder.id ? updatedOrder : order,
  );
  saveOrders(orders);
};

export const deleteOrder = (id: number) => {
  const orders = getOrders().filter((order) => order.id != id);
  saveOrders(orders);
};
