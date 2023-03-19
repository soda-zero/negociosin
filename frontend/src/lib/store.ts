import { GetAllProducts } from "../../wailsjs/go/backend/App";
import { writable } from "svelte/store";

export const productsStore = writable([]);

export async function loadProducts() {
  const products = await GetAllProducts();
  productsStore.set(products);
}
