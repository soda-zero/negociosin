<script lang="ts">
    import { categoryStore, loadCategory, loadProducts } from "$lib/store";
    import { DeleteCategoryById } from "$wails/go/backend/App";
    import type { backend } from "$wails/go/models";
    import { onMount } from "svelte";
    import TrashIcon from "$components/icons/TrashIcon.svelte";
    import {
        modalStore,
        type ModalComponent,
        type ModalSettings,
    } from "@skeletonlabs/skeleton";
    import UpdateCategoryModal from "./modals/UpdateCategoryModal.svelte";
    import EditIcon from "./icons/EditIcon.svelte";
    import CreateCategoryModal from "./modals/CreateCategoryModal.svelte";
    import PlusIcon from "./icons/PlusIcon.svelte";

    let searchQuery: string = "";
    let categories: backend.Category[] = [];
    $: categories = $categoryStore;
    let filteredCategories = [];

    async function initalizeData() {
        await loadCategory();
        filteredCategories = categories;
    }
    onMount(async () => {
        await initalizeData();
    });

    function createCategory() {
        const c: ModalComponent = {
            ref: CreateCategoryModal,
        };
        const d: ModalSettings = {
            buttonTextCancel: "X",
            type: "component",
            component: c,
            response: async () => {
                await loadCategory();
                filterCategories();
            },
        };
        modalStore.trigger(d);
    }
    async function deleteCategory(id: number) {
        const confirm: ModalSettings = {
            type: "confirm",
            title: "Eliminar categoría",
            body: "Estás seguro que deseas eliminar la categoría?",
            buttonTextCancel: "Cancelar",
            buttonTextConfirm: "Confirmar",
            response: async (confirmed: boolean) => {
                if (confirmed) {
                    await DeleteCategoryById(id);
                    await loadCategory();
                    filterCategories();
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
            response: async () => {
                await loadCategory();
                filterCategories();
            },
        };
        modalStore.trigger(d);
    }

    function filterCategories() {
        if (!categories) {
            categories = [];
        }
        filteredCategories = categories.filter((category) => {
            return (
                filterByName(category) ||
                filterByProfitPercent(category) ||
                filterByUpdatedAt(category)
            );
        });
    }

    function filterByName(category: backend.Category) {
        return category.name.toLowerCase().includes(searchQuery.toLowerCase());
    }
    function filterByProfitPercent(category: backend.Category) {
        return category.profit_percent
            .toString()
            .toLowerCase()
            .includes(searchQuery.toLowerCase());
    }
    function filterByUpdatedAt(category: backend.Category) {
        const updatedAt = new Date(category.updated_at);
        return updatedAt
            .toLocaleString()
            .toLowerCase()
            .includes(searchQuery.toLowerCase());
    }
</script>

<button class="btn  variant-filled-primary" on:click={createCategory}
    >Agregar Categoría <PlusIcon /></button
>

<div class="input-group input-group-divider grid-cols-[auto_1fr_auto]">
    <div class="input-group-shim">Categorías</div>
    <input
        type="search"
        placeholder="Buscar..."
        bind:value={searchQuery}
        on:input={filterCategories}
    />
</div>

<div class="table-container shadow-xl">
    <table class="table table-hover">
        <thead>
            <tr>
                <th>Cod</th>
                <th>Nombre categoría</th>
                <th>Porcentaje ganancia</th>
                <th>Actualizado</th>
                <th />
            </tr>
        </thead>
        <tbody>
            {#if filteredCategories}
                {#each filteredCategories as category}
                    <tr>
                        <td>{category.id}</td>
                        <td class="capitalize">{category.name}</td>
                        <td>{category.profit_percent}%</td>
                        <td>{new Date(category.updated_at).toLocaleString()}</td
                        >
                        <td>
                            <button
                                class=" bg-initial hover:text-success-400"
                                on:click={() => EditCategory(category)}
                                ><EditIcon /></button
                            >

                            <button
                                class=" bg-initial hover:text-error-400"
                                on:click={() => deleteCategory(category.id)}
                                ><TrashIcon /></button
                            ></td
                        >
                    </tr>
                {/each}
            {/if}
        </tbody>
    </table>
</div>
