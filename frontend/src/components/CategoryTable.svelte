<script lang="ts">
    import { loadProducts, categoryStore } from "../lib/store";
    import { onDestroy, onMount } from "svelte";
    import type { backend } from "../../wailsjs/go/models";
    import { DeleteCategoryById } from "../../wailsjs/go/backend/App";
    import TrashIcon from "./TrashIcon.svelte";

    let categoryList: backend.Category[] = [];
    let filteredCategoryList: backend.Category[] = [];
    let searchQuery: string = "";
    const unsubscribe = categoryStore.subscribe((categories) => {
        categoryList = categories;
        if (categoryList !== null) {
            filterCategories();
        }
    });

    onMount(() => {
        loadProducts();
    });

    onDestroy(() => {
        unsubscribe();
    });
    async function DeleteProduct(id: number) {
        await DeleteCategoryById(id);
        await loadProducts();
    }

    function filterCategories() {
        filteredCategoryList = categoryList.filter((category) => {
            const updatedAt = new Date(category.updated_at);
            return (
                category.name
                    .toLowerCase()
                    .includes(searchQuery.toLowerCase()) ||
                category.profit_percent
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
    <div class="input-group-shim">Categorías</div>
    <input
        type="search"
        placeholder="Buscar..."
        bind:value={searchQuery}
        on:input={filterCategories}
    />
    <button class="variant-filled-secondary" on:click={filterCategories}
        >Buscar</button
    >
</div>

<div class="table-container">
    <table class="table table-hover">
        <thead>
            <tr>
                <th>Cod</th>
                <th>Nombre categoría</th>
                <th>Actualizado</th>
                <th>Porcentaje ganancia</th>
                <th />
            </tr>
        </thead>
        <tbody>
            {#each filteredCategoryList as category}
                <tr>
                    <td>{category.id}</td>
                    <td>{category.name}</td>
                    <td>{new Date(category.updated_at).toLocaleString()}</td>
                    <td>{category.profit_percent}%</td>
                    <td
                        ><button
                            class=" bg-initial hover:text-error-400"
                            on:click={() => DeleteProduct(category.id)}
                            ><TrashIcon /></button
                        ></td
                    >
                </tr>
            {/each}
        </tbody>
    </table>
</div>
