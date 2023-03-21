import {
  GetAllCategories,
  GetAllProducts,
  GetAllProviders,
} from "../../wailsjs/go/backend/App";
import { writable } from "svelte/store";

export const productsStore = writable([]);
export const providerStore = writable([]);
export const categoryStore = writable([]);
export async function loadProducts() {
  const products = await GetAllProducts();
  const providers = await GetAllProviders();
  const categories = await GetAllCategories();
  productsStore.set(products);
  providerStore.set(providers);
  categoryStore.set(categories);
}
