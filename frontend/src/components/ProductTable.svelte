<script lang="ts">
    import {
        Accordion,
        AccordionItem,
        modalStore,
        type ModalComponent,
        type ModalSettings,
    } from "@skeletonlabs/skeleton";
    import { loadProducts, productsStore } from "$lib/store";
    import type { backend } from "$wails/go/models";
    import { DeleteProductById } from "$wails/go/backend/App";
    import EditIcon from "./icons/EditIcon.svelte";
    import TrashIcon from "./icons/TrashIcon.svelte";
    import UpdateProductModal from "./modals/UpdateProductModal.svelte";
    import FilterIcon from "./icons/FilterIcon.svelte";
    import TagIcon from "./icons/TagIcon.svelte";
    import TruckIcon from "./icons/TruckIcon.svelte";
    import MagnifyingGlass from "./icons/MagnifyingGlass.svelte";
    import { onMount } from "svelte";
    import CreateProductModal from "./modals/CreateProductModal.svelte";
    import PlusIcon from "./icons/PlusIcon.svelte";

    let searchQuery = "";
    let products: backend.ProductWithCategoryAndProvider[] = [];
    $: products = $productsStore;

    let providers: string[] = [];
    let categories: string[] = [];
    let filteredProducts: backend.ProductWithCategoryAndProvider[] = [];

    function updateCategoriesAndProviders() {
        if (products) {
            providers = [
                ...new Set(products.map((product) => product.provider)),
            ];
            categories = [
                ...new Set(products.map((product) => product.category_name)),
            ];
        }
        if (!products) {
            providers = [];
            categories = [];
        }
    }
    async function initializeData() {
        await loadProducts();
        filteredProducts = products;
        updateCategoriesAndProviders();
    }

    onMount(async () => {
        await initializeData();
    });

    function roundPrice(num: number) {
        const roundedNum = Math.round(num * 2) / 2;
        return roundedNum;
    }

    function createProduct() {
        const c: ModalComponent = {
            ref: CreateProductModal,
        };
        const d: ModalSettings = {
            buttonTextCancel: "X",
            type: "component",
            component: c,
            response: async () => {
                await loadProducts();
                updateCategoriesAndProviders();
                filterProducts();
            },
        };
        modalStore.trigger(d);
    }

    function deleteProduct(id: number) {
        const confirm: ModalSettings = {
            type: "confirm",
            title: "Eliminar producto",
            body: "Estás seguro que deseas eliminar el producto?",
            buttonTextCancel: "Cancelar",
            buttonTextConfirm: "Confirmar",
            response: async (confirmed: boolean) => {
                if (confirmed) {
                    await DeleteProductById(id);
                    await loadProducts();
                    updateCategoriesAndProviders();
                    filterProducts();
                }
            },
        };
        modalStore.trigger(confirm);
    }
    function editProduct(product: backend.ProductWithCategoryAndProvider) {
        const c: ModalComponent = {
            ref: UpdateProductModal,
            props: { product },
        };
        const d: ModalSettings = {
            buttonTextCancel: "X",
            type: "component",
            component: c,
            title: "Edit Product",
            response: async () => {
                await loadProducts();
                await updateCategoriesAndProviders();
                filterProducts();
            },
        };
        modalStore.trigger(d);
    }

    let selectedProviders: string[] = [];
    let selectedCategories: string[] = [];

    function toggleSelectedProvider(provider: string) {
        const index = selectedProviders.indexOf(provider);
        if (index !== -1) {
            selectedProviders.splice(index, 1);
        } else {
            selectedProviders.push(provider);
        }
        filterProducts();
    }

    function toggleSelectedCategory(category: string) {
        const index = selectedCategories.indexOf(category);
        if (index !== -1) {
            selectedCategories.splice(index, 1);
        } else {
            selectedCategories.push(category);
        }
        filterProducts();
    }
    function filterProducts() {
        if (!products) {
            products = [];
        }

        filteredProducts = products.filter((product) => {
            const selectedProvider =
                selectedProviders.length === 0 ||
                selectedProviders.includes(product.provider);
            const selectedCategory =
                selectedCategories.length === 0 ||
                selectedCategories.includes(product.category_name);

            return (
                (!selectedProviders.length || selectedProvider) &&
                (!selectedCategories.length || selectedCategory) &&
                (filterByName(product) ||
                    filterByProvider(product) ||
                    filterByCategory(product) ||
                    filterByPrice(product) ||
                    filterByUpdatedAt(product))
            );
        });
    }
    function filterByName(product: backend.ProductWithCategoryAndProvider) {
        return product.name.toLowerCase().includes(searchQuery.toLowerCase());
    }

    function filterByProvider(product: backend.ProductWithCategoryAndProvider) {
        return product.provider
            .toLowerCase()
            .includes(searchQuery.toLowerCase());
    }

    function filterByCategory(product: backend.ProductWithCategoryAndProvider) {
        return product.category_name
            .toLowerCase()
            .includes(searchQuery.toLowerCase());
    }

    function filterByPrice(product: backend.ProductWithCategoryAndProvider) {
        return (
            product.unit_cost_price
                .toString()
                .toLowerCase()
                .includes(searchQuery.toLowerCase()) ||
            product.unit_sell_price
                .toString()
                .toLowerCase()
                .includes(searchQuery.toLowerCase())
        );
    }

    function filterByUpdatedAt(
        product: backend.ProductWithCategoryAndProvider
    ) {
        const updatedAt = new Date(product.updated_at);
        return updatedAt
            .toLocaleString()
            .toLowerCase()
            .includes(searchQuery.toLowerCase());
    }
</script>

<button class="btn  variant-filled-primary" on:click={createProduct}
    >Agregar Producto <PlusIcon /></button
>
<div class=" card ">
    <Accordion>
        <AccordionItem>
            <svelte:fragment slot="lead"><FilterIcon /></svelte:fragment>
            <svelte:fragment slot="summary">Filtrar</svelte:fragment>
            <svelte:fragment slot="content">
                <AccordionItem>
                    <svelte:fragment slot="lead"><TruckIcon /></svelte:fragment>
                    <svelte:fragment slot="summary">Proveedores</svelte:fragment
                    >
                    <svelte:fragment slot="content">
                        <div class="flex gap-2">
                            {#if providers}
                                {#each providers as provider}
                                    <label>
                                        <input
                                            type="checkbox"
                                            value={provider}
                                            on:change={() =>
                                                toggleSelectedProvider(
                                                    provider
                                                )}
                                        />
                                        {provider}
                                    </label>
                                {/each}
                            {/if}
                        </div>
                    </svelte:fragment>
                </AccordionItem>
                <AccordionItem>
                    <svelte:fragment slot="lead"><TagIcon /></svelte:fragment>
                    <svelte:fragment slot="summary">Categoría</svelte:fragment>
                    <svelte:fragment slot="content">
                        <div class="flex gap-2">
                            {#if categories}
                                {#each categories as category}
                                    <label>
                                        <input
                                            type="checkbox"
                                            value={category}
                                            on:change={() =>
                                                toggleSelectedCategory(
                                                    category
                                                )}
                                        />
                                        {category}
                                    </label>
                                {/each}
                            {/if}
                        </div>
                    </svelte:fragment>
                </AccordionItem>
            </svelte:fragment>
        </AccordionItem>
    </Accordion>
</div>
<div class="input-group input-group-divider grid-cols-[auto_1fr_auto]">
    <div class="input-group-shim flex gap-2">
        <MagnifyingGlass /> Productos
    </div>
    <input
        type="search"
        placeholder="Buscar..."
        bind:value={searchQuery}
        on:input={filterProducts}
    />
</div>

<div class="table-container shadow-xl">
    <table class="table table-hover ">
        <thead>
            <tr>
                <th>Cod</th>
                <th>Nombre</th>
                <th>Precio Costo Unitario</th>
                <th>Precio Venta Unitario</th>
                <th>Proveedor</th>
                <th>Categoría</th>
                <th>Actualizado</th>
                <th />
            </tr>
        </thead>
        <tbody>
            {#if filteredProducts}
                {#each filteredProducts as product}
                    <tr>
                        <td>{product.id}</td>
                        <td class="capitalize">{product.name.toLowerCase()}</td>
                        <td>{product.unit_cost_price}</td>
                        <td
                            >{`${roundPrice(product.unit_sell_price)}${" "}(${
                                product.category_profit_percent
                            }%)`}
                        </td>
                        <td class="capitalize"
                            >{product.provider.toLowerCase()}</td
                        >
                        <td class="capitalize"
                            >{product.category_name.toLowerCase()}
                        </td>
                        <td>{new Date(product.updated_at).toLocaleString()}</td>
                        <td class="flex"
                            ><button
                                class=" bg-initial hover:text-success-400"
                                on:click={() => editProduct(product)}
                                ><EditIcon /></button
                            >

                            <button
                                class=" bg-initial hover:text-error-400"
                                on:click={() => deleteProduct(product.id)}
                                ><TrashIcon /></button
                            ></td
                        >
                    </tr>
                {/each}
            {/if}
        </tbody>
    </table>
</div>

