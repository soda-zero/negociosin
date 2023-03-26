<script lang="ts">
    import { onMount } from "svelte";
    import {
        modalStore,
        type ModalComponent,
        type ModalSettings,
    } from "@skeletonlabs/skeleton";
    import { loadProviders, providerStore } from "$lib/store";
    import type { backend } from "$wails/go/models";
    import { DeleteProviderById } from "$wails/go/backend/App";
    import EditIcon from "./icons/EditIcon.svelte";
    import TrashIcon from "./icons/TrashIcon.svelte";
    import UpdateProviderModal from "./modals/UpdateProviderModal.svelte";
    import CreateProviderModal from "./modals/CreateProviderModal.svelte";
    import PlusIcon from "./icons/PlusIcon.svelte";

    let searchQuery: string = "";
    let providers: backend.Provider[] = [];
    $: providers = $providerStore;
    let filteredProviders: backend.Provider[] = [];

    async function initializeData() {
        filteredProviders = providers;
        await loadProviders();
    }
    onMount(async () => {
        await initializeData();
    });

    function createProvider() {
        const c: ModalComponent = {
            ref: CreateProviderModal,
        };
        const d: ModalSettings = {
            buttonTextCancel: "X",
            type: "component",
            component: c,
            response: async () => {
                await loadProviders();
                filterProviders();
            },
        };
        modalStore.trigger(d);
    }
    async function deleteProvider(id: number) {
        const confirm: ModalSettings = {
            type: "confirm",
            title: "Eliminar proveedor",
            body: "Estás seguro que deseas eliminar el proveedor?",
            buttonTextCancel: "Cancelar",
            buttonTextConfirm: "Confirmar",
            response: async (confirmed: boolean) => {
                if (confirmed) {
                    await DeleteProviderById(id);
                    await loadProviders();
                    filterProviders();
                }
            },
        };
        modalStore.trigger(confirm);
    }

    function editProvider(provider: backend.Provider) {
        const c: ModalComponent = {
            ref: UpdateProviderModal,
            props: { provider },
        };
        const d: ModalSettings = {
            buttonTextCancel: "X",
            type: "component",
            component: c,
            title: "Edit Provider",
            response: async () => {
                await loadProviders();
                filterProviders();
            },
        };
        modalStore.trigger(d);
    }

    function filterProviders() {
        if (!providers) {
            providers = [];
        }
        filteredProviders = providers.filter((provider) => {
            return (
                filterByName(provider) ||
                filterByPhoneNumber(provider) ||
                filterByUpdatedAt(provider)
            );
        });
    }

    function filterByName(provider: backend.Provider) {
        return provider.name.toLowerCase().includes(searchQuery.toLowerCase());
    }

    function filterByPhoneNumber(provider: backend.Provider) {
        return provider.phone_number
            .toLowerCase()
            .includes(searchQuery.toLowerCase());
    }

    function filterByUpdatedAt(provider: backend.Provider) {
        const updatedAt = new Date(provider.updated_at);
        return updatedAt
            .toLocaleString()
            .toLowerCase()
            .includes(searchQuery.toLowerCase());
    }
</script>

<button class="btn  variant-filled-primary" on:click={createProvider}
    >Agregar Proveedor <PlusIcon /></button
>
<div class="input-group input-group-divider grid-cols-[auto_1fr_auto]">
    <div class="input-group-shim">Productos</div>
    <input
        type="search"
        placeholder="Buscar..."
        bind:value={searchQuery}
        on:input={filterProviders}
    />
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
            {#if filteredProviders}
                {#each filteredProviders as provider}
                    <tr>
                        <td>{provider.id}</td>
                        <td class="capitalize">{provider.name.toLowerCase()}</td
                        >
                        <td>{provider.phone_number}</td>
                        <td>{new Date(provider.updated_at).toLocaleString()}</td
                        >
                        <td
                            ><button
                                class=" bg-initial hover:text-success-400"
                                on:click={() => editProvider(provider)}
                                ><EditIcon /></button
                            >

                            <button
                                class=" bg-initial hover:text-error-400"
                                on:click={() => deleteProvider(provider.id)}
                                ><TrashIcon /></button
                            ></td
                        >
                    </tr>
                {/each}
            {/if}
        </tbody>
    </table>
</div>
