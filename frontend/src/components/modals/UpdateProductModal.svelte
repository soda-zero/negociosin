<script lang="ts">
    import { categoryStore, loadProducts, providerStore } from "$lib/store";
    import { UpdateProductById } from "$wails/go/backend/App";
    import type { backend } from "$wails/go/models";
    import type { ModalProps } from "@skeletonlabs/skeleton/utilities/Modal/Modal.svelte";
    import { onDestroy } from "svelte";
    export let parent: ModalProps;
    export let product: backend.ProductWithCategoryAndProvider;

    let categoryList: backend.Category[] = [];
    let providerList: backend.Provider[] = [];

    const unsubscribeCategory = categoryStore.subscribe((categories) => {
        categoryList = categories;
    });

    const unsubscribeProvider = providerStore.subscribe((providers) => {
        providerList = providers;
    });

    async function updateProduct() {
        await UpdateProductById(product.id, product);
        await loadProducts();
        parent.onClose();
    }

    onDestroy(() => {
        unsubscribeCategory();
        unsubscribeProvider();
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
    <h2>Modificar Producto</h2>
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
                        >Selecciona un proveedor</option
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
                    step="0.001"
                    bind:value={product.original_cost_price}
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
                    step="0.001"
                    bind:value={product.internal_tax}
                />
            </label>
            <label class="label">
                <span>Iva</span>
                <input
                    type="number"
                    placeholder="21..."
                    class="input"
                    step="0.001"
                    bind:value={product.iva}
                />
            </label>
        </div>
    </div>
    <button
        type="button"
        on:click={updateProduct}
        class="w-full btn btn- variant-filled-secondary">Modificar</button
    >
</form>
