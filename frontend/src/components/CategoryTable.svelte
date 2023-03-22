<script lang="ts">
    import { categoryStore, loadProducts } from "$lib/store";
    import { DeleteCategoryById } from "$wails/go/backend/App";
    import type { backend } from "$wails/go/models";
    import { onDestroy, onMount } from "svelte";
    import TrashIcon from "$components/icons/TrashIcon.svelte";
    import {
        modalStore,
        type ModalComponent,
        type ModalSettings,
    } from "@skeletonlabs/skeleton";
    import UpdateCategoryModal from "./modals/UpdateCategoryModal.svelte";
    import EditIcon from "./icons/EditIcon.svelte";

    let categoryList: backend.Category[] = [];
    let filteredCategoryList: backend.Category[] = [];
    filteredCategoryList = categoryList || [];
    let searchQuery: string = "";
    const unsubscribe = categoryStore.subscribe((categories) => {
        categoryList = categories;
        if (categoryList !== null) {
            filterCategories();
        }
    });

    async function DeleteCategory(id: number) {
        const confirm: ModalSettings = {
            type: "confirm",
            title: "Eliminar categoría",
            body: "Estás seguro que deseas eliminar la categoría?",
            buttonTextCancel: "Cancelar",
            buttonTextConfirm: "Confirmar",
            response: async (confirmed: boolean) => {
                if (confirmed) {
                    await DeleteCategoryById(id);
                    const index = categoryList.findIndex(
                        (category) => category.id === id
                    );
                    if (index !== -1) {
                        categoryList.splice(index, 1);
                        filterCategories();
                    }
                    await loadProducts();
                }
            },
        };
        modalStore.trigger(confirm);
    }

    function EditCategory(category: backend.Category) {
        const c: ModalComponent = {
            ref: UpdateCategoryModal,
            props: { category },
        };
        const d: ModalSettings = {
            buttonTextCancel: "X",
            type: "component",
            component: c,
            title: "Edit Category",
        };
        modalStore.trigger(d);
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

    onMount(() => {
        loadProducts();
    });

    onDestroy(() => {
        unsubscribe();
    });
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
            {#if filteredCategoryList != null}
                {#each filteredCategoryList as category}
                    <tr>
                        <td>{category.id}</td>
                        <td class="capitalize">{category.name}</td>
                        <td>{new Date(category.updated_at).toLocaleString()}</td
                        >
                        <td>{category.profit_percent}%</td>
                        <td>
                            <button
                                class=" bg-initial hover:text-success-400"
                                on:click={() => EditCategory(category)}
                                ><EditIcon /></button
                            >

                            <button
                                class=" bg-initial hover:text-error-400"
                                on:click={() => DeleteCategory(category.id)}
                                ><TrashIcon /></button
                            ></td
                        >
                    </tr>
                {/each}
            {/if}
        </tbody>
    </table>
</div>
