export interface Customer {
  id: number;
  name: string;
  phone: string;
}

const STORAGE_KEY = "laundry_customers";

export const getCustomers = (): Customer[] => {
  return JSON.parse(localStorage.getItem(STORAGE_KEY) || "[]");
};

export const saveCustomers = (customers: Customer[]) => {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(customers));
};

export const addCustomer = (customer: Customer) => {
  const customers = getCustomers();
  customers.push(customer);
  saveCustomers(customers);
};
