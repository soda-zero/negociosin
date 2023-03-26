import { writable } from "svelte/store";
import { GetAllProducts, GetAllProviders, GetAllCategories } from "$wails/go/backend/App";

// Cache the data to avoid making unnecessary API requests
let cachedProducts = [];
let cachedProviders = [];
let cachedCategories = [];

export const productsStore = writable(cachedProducts);
export const providerStore = writable(cachedProviders);
export const categoryStore = writable(cachedCategories);

// Handle errors when fetching data
export async function loadProducts() {
  try {
    const products = await GetAllProducts();
    cachedProducts = products;
    productsStore.set(products);
  } catch (error) {
    console.error("Error fetching products:", error);
    // Optionally, you could add some UI feedback to indicate the error to the user
  }
}

export async function loadProviders() {
  try {
    const providers = await GetAllProviders();
    cachedProviders = providers;
    providerStore.set(providers);
  } catch (error) {
    console.error("Error fetching providers:", error);
    // Optionally, you could add some UI feedback to indicate the error to the user
  }
}

export async function loadCategory() {
  try {
    const categories = await GetAllCategories();
    cachedCategories = categories;
    categoryStore.set(categories);
  } catch (error) {
    console.error("Error fetching categories:", error);
    // Optionally, you could add some UI feedback to indicate the error to the user
  }
}
