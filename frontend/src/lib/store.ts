import { GetAllProducts, GetProviders } from "../../wailsjs/go/backend/App";
import { writable } from "svelte/store";

export const productsStore = writable([]);
export const providerStore = writable([]);

export async function loadProducts() {
  const products = await GetAllProducts();
  const providers = await GetProviders();
  productsStore.set(products);
  providerStore.set(providers);
}
