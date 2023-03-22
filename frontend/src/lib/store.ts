import {
  GetAllCategories,
  GetAllProducts,
  GetAllProviders,
} from "$wails/go/backend/App";
import { writable } from "svelte/store";

export const productsStore = writable([]);
export const providerStore = writable([]);
export const categoryStore = writable([]);
export async function loadProducts() {
  const products = await GetAllProducts();
  const providers = await GetAllProviders();
  const categories = await GetAllCategories();
  providerStore.set(providers);
  categoryStore.set(categories);
  productsStore.set(products);
}
