<script lang="ts">
    import { loadProducts, productsStore } from "../lib/store";
    import { onDestroy, onMount } from "svelte";
    import type { backend } from "../../wailsjs/go/models";
    import { DeleteProductById } from "../../wailsjs/go/backend/App";
    import TrashIcon from "./TrashIcon.svelte";

    function roundPrice(num: number) {
        const roundedNum = Math.round(num * 2) / 2;
        return roundedNum;
    }
    let productsList: backend.ProductWithCategory[] = [];
    let filteredProductsList: backend.ProductWithCategory[] = [];
    let searchQuery: string = "";
    const unsubscribe = productsStore.subscribe((products) => {
        productsList = products;
        if (productsList !== null) {
            filterProducts();
        }
    });

    onMount(() => {
        loadProducts();
    });

    onDestroy(() => {
        unsubscribe();
    });
    async function DeleteProduct(id: number) {
        await DeleteProductById(id);
        await loadProducts();
    }

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
                    .includes(searchQuery.toLowerCase()) ||
                product.sell_price
                    .toString()
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

<div class="table-container shadow-xl">
    <table class="table table-hover ">
        <thead>
            <tr>
                <th>Cod</th>
                <th>Nombre</th>
                <th>Actualizado</th>
                <th>Proveedor</th>
                <th>Precio Costo</th>
                <th>Precio Venta</th>
                <th>Categoría</th>
                <th />
            </tr>
        </thead>
        <tbody>
            {#each filteredProductsList as product}
                <tr>
                    <td>{product.id}</td>
                    <td class="capitalize">{product.name.toLowerCase()}</td>
                    <td>{new Date(product.updated_at).toLocaleString()}</td>
                    <td class="capitalize">{product.provider.toLowerCase()}</td>
                    <td>{roundPrice(product.cost_price)}</td>
                    <td
                        >{`${roundPrice(product.sell_price)}${" "}(${
                            product.category_profit_percent
                        }%)`}
                    </td>
                    <td class="capitalize"
                        >{product.category_name.toLowerCase()}
                    </td>
                    <td
                        ><button
                            class=" bg-initial hover:text-error-400"
                            on:click={() => DeleteProduct(product.id)}
                            ><TrashIcon /></button
                        ></td
                    >
                </tr>
            {/each}
        </tbody>
    </table>
</div>
