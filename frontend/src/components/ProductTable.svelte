<script lang="ts">
    import { onDestroy, onMount } from "svelte";
    import {
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

    function roundPrice(num: number) {
        const roundedNum = Math.round(num * 2) / 2;
        return roundedNum;
    }
    let productsList: backend.ProductWithCategoryAndProvider[] = [];
    let filteredProductsList: backend.ProductWithCategoryAndProvider[] = [];
    let searchQuery: string = "";
    const unsubscribe = productsStore.subscribe((products) => {
        productsList = products;
        if (productsList !== null) {
            filterProducts();
        }
    });

    async function DeleteProduct(id: number) {
        const confirm: ModalSettings = {
            type: "confirm",
            title: "Eliminar producto",
            body: "Estás seguro que deseas eliminar el producto?",
            buttonTextCancel: "Cancelar",
            buttonTextConfirm: "Confirmar",
            response: async (confirmed: boolean) => {
                if (confirmed) {
                    await DeleteProductById(id);
                    const index = productsList.findIndex(
                        (product) => product.id === id
                    );
                    if (index !== -1) {
                        productsList.splice(index, 1);
                        filterProducts();
                    }
                    await loadProducts();
                }
            },
        };
        modalStore.trigger(confirm);
    }

    function EditProduct(product: backend.ProductWithCategoryAndProvider) {
        const c: ModalComponent = {
            ref: UpdateProductModal,
            props: { product },
        };
        const d: ModalSettings = {
            buttonTextCancel: "X",
            type: "component",
            component: c,
            title: "Edit Product",
        };
        modalStore.trigger(d);
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
                product.unit_cost_price
                    .toString()
                    .toLowerCase()
                    .includes(searchQuery.toLowerCase()) ||
                updatedAt
                    .toLocaleString()
                    .toLowerCase()
                    .includes(searchQuery.toLowerCase()) ||
                product.unit_sell_price
                    .toString()
                    .toLowerCase()
                    .includes(searchQuery.toLowerCase()) ||
                product.category_name
                    .toLowerCase()
                    .includes(searchQuery.toLowerCase())
            );
        });
    }
    onMount(() => {
        loadProducts();
    });
    onDestroy(() => {
        unsubscribe();
    });
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
                <th>Precio Costo Unitario</th>
                <th>Precio Venta Unitario</th>
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
                    <td>{product.unit_cost_price}</td>
                    <td
                        >{`${roundPrice(product.unit_sell_price)}${" "}(${
                            product.category_profit_percent
                        }%)`}
                    </td>
                    <td class="capitalize"
                        >{product.category_name.toLowerCase()}
                    </td>
                    <td
                        ><button
                            class=" bg-initial hover:text-success-400"
                            on:click={() => EditProduct(product)}
                            ><EditIcon /></button
                        >

                        <button
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
