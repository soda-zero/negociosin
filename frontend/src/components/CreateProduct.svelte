<script lang="ts">
    import { onDestroy } from "svelte";
    import { CreateProduct } from "../../wailsjs/go/backend/App";
    import { categoryStore, loadProducts } from "../lib/store";
    import { backend } from "../../wailsjs/go/models";
    import type { ModalProps } from "@skeletonlabs/skeleton/utilities/Modal/Modal.svelte";

    export let parent: ModalProps;

    function resetForm() {
        product.name = "";
        product.provider = "";
        product.cost_price = undefined;
        product.category_id = undefined;
    }

    let product = new backend.Product();
    async function createProduct(e: Event) {
        e.preventDefault();
        await CreateProduct(product);
        await loadProducts();
        resetForm();
        parent.onClose()
    }

    let categoryList: backend.Category[] = [];
    const unsubscribe = categoryStore.subscribe((categories) => {
        categoryList = categories;
    });

    onDestroy(() => unsubscribe);
</script>

<form
    on:submit={createProduct}
    class="flex flex-col gap-2 card p-4 w-modal-slim shadow-xl space-y-4 "
>
    <div class="flex justify-end">
        <button
            class="btn-icon btn-icon-sm {parent.buttonNeutral} "
            on:click={parent.onClose}>{parent.buttonTextCancel}</button
        >
    </div>
    <h2>Agregar Producto</h2>
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
            <span>Precio de costo</span>
            <input
                type="number"
                placeholder="100..."
                class="input"
                bind:value={product.cost_price}
            />
        </label>

        <label class="label">
            <span>Proveedor</span>
            <input
                type="text"
                placeholder="arcor..."
                bind:value={product.provider}
                class="input"
            />
        </label>

        <label class="label">
            <span>Categoría</span>
            <select bind:value={product.category_id} class="select">
                <option value="" disabled selected
                    >Selecciona una categoría</option
                >
                >
                {#each categoryList as category}
                    <option value={category.id}>{category.name}</option>
                {/each}</select
            >
        </label>
    </div>
    <button class="w-full btn btn- variant-filled-secondary">Agregar</button>
</form>
