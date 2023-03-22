<script lang="ts">
    import { onDestroy } from "svelte";
    import type { ModalProps } from "@skeletonlabs/skeleton/utilities/Modal/Modal.svelte";
    import { backend } from "$wails/go/models";
    import { CreateProduct } from "$wails/go/backend/App";
    import { categoryStore, loadProducts, providerStore } from "$lib/store";

    export let parent: ModalProps;

    let product = new backend.Product();
    async function createProduct(e: Event) {
        e.preventDefault();
        await CreateProduct(product);
        await loadProducts();
        parent.onClose();
    }

    let categoryList: backend.Category[] = [];
    const unsubscribe = categoryStore.subscribe((categories) => {
        categoryList = categories;
    });
    let providerList: backend.Provider[] = [];
    const unsubs = providerStore.subscribe((providers) => {
        providerList = providers;
    });

    onDestroy(() => {
        unsubscribe();
        unsubs();
    });
</script>

<form class="flex flex-col gap-2 card p-4 w-modal-slim shadow-xl space-y-4 ">
    <div class="flex justify-end">
        <button
            type="button"
            class="btn-icon btn-icon-sm {parent.buttonNeutral} "
            on:click={parent.onClose}>{parent.buttonTextCancel}</button
        >
    </div>
    <h2>Agregar Producto</h2>
    <div class="flex gap-4">
        <div>
            <label class="label">
                <span>Nombre</span>
                <input
                    type="text"
                    bind:value={product.name}
                    placeholder="alfajor..."
                    class="input"
                />
            </label>
            <label class="label">
                <span>Proveedor</span>
                <select bind:value={product.provider_id} class="select">
                    <option value="" disabled selected
                        >Selecciona una categoría</option
                    >
                    {#each providerList as provider}
                        <option value={provider.id}>{provider.name}</option>
                    {/each}</select
                >
            </label>

            <label class="label">
                <span>Categoría</span>
                <select bind:value={product.category_id} class="select">
                    <option value="" disabled selected
                        >Selecciona una categoría</option
                    >
                    {#each categoryList as category}
                        <option value={category.id}>{category.name}</option>
                    {/each}</select
                >
            </label>
        </div>
        <div>
            <label class="label">
                <span>Precio de costo(No unitario)</span>
                <input
                    type="number"
                    placeholder="100..."
                    class="input"
                    bind:value={product.original_cost_price}
                    step="0.001"
                />
            </label>
            <label class="label">
                <span>Cantidad</span>
                <input
                    type="number"
                    placeholder="1,2,3..."
                    class="input"
                    bind:value={product.quantity}
                />
            </label>
            <label class="label">
                <span>Imp. Interno</span>
                <input
                    type="number"
                    placeholder="1,2,3..."
                    class="input"
                    bind:value={product.internal_tax}
                    step="0.001"
                />
            </label>
            <label class="label">
                <span>Iva</span>
                <input
                    type="number"
                    placeholder="21..."
                    class="input"
                    bind:value={product.iva}
                    step="0.001"
                />
            </label>
        </div>
    </div>
    <button
        on:click={createProduct}
        type="button"
        class="w-full btn btn- variant-filled-secondary">Agregar</button
    >
</form>
