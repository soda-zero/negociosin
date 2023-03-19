<script lang="ts">
    import { loadProducts, productsStore } from "../lib/store";
    import { onDestroy, onMount } from "svelte";
    import type { Product } from "src/models/products";

    let productsList: Product[] = [];
    let filteredProductsList: Product[] = [];
    let searchQuery: string = "";
    const unsubscribe = productsStore.subscribe((products) => {
        productsList = products;
        filterProducts();
    });
    async function LoadProducts() {
        await loadProducts();
    }
    onMount(() => {
        LoadProducts();
    });

    onDestroy(() => {
        unsubscribe();
    });
    function filterProducts() {
        filteredProductsList = productsList.filter((product) => {
            const updatedAt = new Date(product.updated_at);
            return (
                product.name
                    .toLowerCase()
                    .includes(searchQuery.toLowerCase()) ||
                product.provider
                    .toLowerCase()
                    .includes(searchQuery.toLowerCase()) ||
                product.cost_price
                    .toString()
                    .toLowerCase()
                    .includes(searchQuery.toLowerCase()) ||
                updatedAt
                    .toLocaleString()
                    .toLowerCase()
                    .includes(searchQuery.toLowerCase())
            );
        });
    }
</script>

<div class="input-group input-group-divider grid-cols-[auto_1fr_auto]">
    <div class="input-group-shim">Productos</div>
    <input
        type="search"
        placeholder="Buscar..."
        bind:value={searchQuery}
        on:input={filterProducts}
    />
    <button class="variant-filled-secondary" on:click={filterProducts}
        >Buscar</button
    >
</div>

<div class="table-container">
    <table class="table table-hover">
        <thead>
            <tr>
                <th>Cod</th>
                <th>Nombre</th>
                <th>Actualizado</th>
                <th>Proveedor</th>
                <th>Precio Costo</th>
                <th>Precio Venta</th>
            </tr>
        </thead>
        <tbody>
            {#each filteredProductsList as product}
                <tr>
                    <td>{product.id}</td>
                    <td>{product.name}</td>
                    <td>{new Date(product.updated_at).toLocaleString()}</td>
                    <td>{product.provider}</td>
                    <td>{product.cost_price}</td>
                    <td>{(product.cost_price * 1.15).toFixed(2)} </td></tr
                >
            {/each}
        </tbody>
    </table>
</div>
