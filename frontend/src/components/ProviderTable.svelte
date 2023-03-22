<script lang="ts">
    import { onDestroy, onMount } from "svelte";
    import {
        modalStore,
        type ModalComponent,
        type ModalSettings,
    } from "@skeletonlabs/skeleton";
    import { loadProducts, providerStore } from "$lib/store";
    import type { backend } from "$wails/go/models";
    import { DeleteProviderById } from "$wails/go/backend/App";
    import EditIcon from "./icons/EditIcon.svelte";
    import TrashIcon from "./icons/TrashIcon.svelte";
    import UpdateProviderModal from "./modals/UpdateProviderModal.svelte";

    let providerList: backend.Provider[] = [];
    let filteredProviderList: backend.Provider[] = [];
    let searchQuery: string = "";
    const unsubscribe = providerStore.subscribe((provider) => {
        providerList = provider;
        if (providerList !== null) {
            filterProviders();
        }
    });

    async function DeleteProvider(id: number) {
        const confirm: ModalSettings = {
            type: "confirm",
            title: "Eliminar proveedor",
            body: "Estás seguro que deseas eliminar el proveedor?",
            buttonTextCancel: "Cancelar",
            buttonTextConfirm: "Confirmar",
            response: async (confirmed: boolean) => {
                if (confirmed) {
                    await DeleteProviderById(id);
                    const index = providerList.findIndex(
                        (provider) => provider.id === id
                    );
                    if (index !== -1) {
                        providerList.splice(index, 1);
                        filterProviders();
                    }
                    await loadProducts();
                }
            },
        };
        modalStore.trigger(confirm);
    }

    function EditProvider(provider: backend.Provider) {
        const c: ModalComponent = {
            ref: UpdateProviderModal,
            props: { provider },
        };
        const d: ModalSettings = {
            buttonTextCancel: "X",
            type: "component",
            component: c,
            title: "Edit Provider",
        };
        modalStore.trigger(d);
    }

    function filterProviders() {
        filteredProviderList = providerList.filter((provider) => {
            const updatedAt = new Date(provider.updated_at);
            return (
                provider.name
                    .toLowerCase()
                    .includes(searchQuery.toLowerCase()) ||
                provider.phone_number
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
    <div class="input-group-shim">Productos</div>
    <input
        type="search"
        placeholder="Buscar..."
        bind:value={searchQuery}
        on:input={filterProviders}
    />
    <button class="variant-filled-secondary" on:click={filterProviders}
        >Buscar</button
    >
</div>

<div class="table-container shadow-xl">
    <table class="table table-hover ">
        <thead>
            <tr>
                <th>Cod</th>
                <th>Nombre</th>
                <th>Número de telefono</th>
                <th>Actualizado</th>
                <th />
            </tr>
        </thead>
        <tbody>
            {#each filteredProviderList as provider}
                <tr>
                    <td>{provider.id}</td>
                    <td class="capitalize">{provider.name.toLowerCase()}</td>
                    <td>{provider.phone_number}</td>
                    <td>{new Date(provider.updated_at).toLocaleString()}</td>
                    <td
                        ><button
                            class=" bg-initial hover:text-success-400"
                            on:click={() => EditProvider(provider)}
                            ><EditIcon /></button
                        >

                        <button
                            class=" bg-initial hover:text-error-400"
                            on:click={() => DeleteProvider(provider.id)}
                            ><TrashIcon /></button
                        ></td
                    >
                </tr>
            {/each}
        </tbody>
    </table>
</div>
