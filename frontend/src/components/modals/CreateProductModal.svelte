<script lang="ts">
    import { onMount } from "svelte";
    import type { ModalProps } from "@skeletonlabs/skeleton/utilities/Modal/Modal.svelte";
    import { backend } from "$wails/go/models";
    import { CreateProduct } from "$wails/go/backend/App";
    import {
        categoryStore,
        loadCategory,
        loadProducts,
        loadProviders,
        providerStore,
    } from "$lib/store";
    import { ProductSchema, type Product } from "../../models/Product";
    import { validateForm } from "$lib/helper";

    let categories: backend.Category[] = [];
    let providers: backend.Provider[] = [];

    async function initializeData() {
        await loadProducts();
        await loadCategory();
        await loadProviders();
        categories = $categoryStore;
        providers = $providerStore;
    }

    export let parent: ModalProps;
    let product = new backend.Product();
    let errors: Product = {};
    let ivaList: number[] = [10.5, 21];
    let selectedCategory: backend.Category = null;
    $: {
        if (product.category_id) {
            selectedCategory = categories.find(
                (category) => category.id === product.category_id
            );
        }
    }

    onMount(async () => {
        await initializeData();
    });

    async function createProduct() {
        const { errors: validationErrors, valid } = validateForm(
            ProductSchema,
            product
        );
        if (!valid) {
            errors = validationErrors;
            return;
        }
        await CreateProduct(product);
        parent.onClose();
    }
</script>

<form
    on:submit|preventDefault={createProduct}
    class="flex flex-col gap-2 card p-4 w-modal shadow-xl space-y-4 "
>
    <div class="flex justify-end">
        <button
            type="button"
            class="btn-icon btn-icon-sm {parent.buttonNeutral} "
            on:click={parent.onClose}>{parent.buttonTextCancel}</button
        >
    </div>
    <h2>Agregar Producto</h2>
    <div class="grid grid-cols-2 gap-4">
        <div>
            <label
                class={`${
                    errors.name ? "text-error-500 flex flex-col" : ""
                } label`}
            >
                <span>Nombre</span>
                <input
                    type="text"
                    bind:value={product.name}
                    placeholder="alfajor..."
                    class={`${errors.name ? "input-error" : "input"}`}
                />
                {#if errors.name}
                    <p class="text-error-500 font-bold  w-fit !text-sm">
                        {errors.name}
                    </p>
                {/if}
            </label>
            <label class="label">
                <span>Proveedor</span>
                <select bind:value={product.provider_id} class="select">
                    <option disabled value={null} selected
                        >Selecciona un proveedor</option
                    >
                    {#if providers}
                        {#each providers as provider}
                            <option value={provider.id}>{provider.name}</option>
                        {/each}
                    {/if}
                </select>
            </label>

            <label class="label">
                <span
                    >{`Categoría ${
                        selectedCategory
                            ? `(${selectedCategory.profit_percent}% ganacia)`
                            : ""
                    }`}</span
                >

                <select bind:value={product.category_id} class="select">
                    <span
                        >{`Categoría ${
                            selectedCategory ? selectedCategory.name : ""
                        }`}</span
                    >
                    <option value={null} disabled selected
                        >Selecciona una categoría</option
                    >
                    {#if categories}
                        {#each categories as category}
                            <option value={category.id}>{category.name}</option>
                        {/each}
                    {/if}
                </select>
            </label>
        </div>
        <div>
            <label
                class={`${
                    errors.original_cost_price
                        ? "text-error-500 flex flex-col"
                        : ""
                } label`}
            >
                <span>Precio de costo(No unitario)</span>
                <input
                    type="number"
                    placeholder="100..."
                    class={`${
                        errors.original_cost_price ? "input-error" : "input"
                    }`}
                    bind:value={product.original_cost_price}
                    step="0.001"
                />
                {#if errors.original_cost_price}
                    <p class="text-error-500 font-bold  w-fit !text-sm">
                        {errors.original_cost_price}
                    </p>
                {/if}
            </label>
            <label
                class={`${
                    errors.quantity ? "text-error-500 flex flex-col" : ""
                } label`}
            >
                <span>Cantidad</span>
                <input
                    type="number"
                    placeholder="1, 2, 3..."
                    class={`${errors.quantity ? "input-error" : "input"}`}
                    bind:value={product.quantity}
                    step="0.001"
                />
                {#if errors.quantity}
                    <p class="text-error-500 font-bold  w-fit !text-sm">
                        {errors.quantity}
                    </p>
                {/if}
            </label>
            <div />
            <div class="grid grid-cols-2 gap-2 ">
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
                    <span>IVA</span>
                    <select class="select" bind:value={product.iva}>
                        {#each ivaList as iva}
                            <option>{iva}</option>
                        {/each}
                    </select>
                </label>
            </div>
        </div>
    </div>
    <button type="submit" class="w-full btn btn- variant-filled-secondary"
        >Agregar</button
    >
</form>
