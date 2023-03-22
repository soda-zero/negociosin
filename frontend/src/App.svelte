<script lang="ts">
    import CategoryTable from "$components/CategoryTable.svelte";
    import PlusIcon from "$components/icons/PlusIcon.svelte";
    import CreateCategoryModal from "$components/modals/CreateCategoryModal.svelte";
    import CreateProductModal from "$components/modals/CreateProductModal.svelte";
    import CreateProviderModal from "$components/modals/CreateProviderModal.svelte";
    import ProductActions from "$components/ProductActions.svelte";
    import ProductTable from "$components/ProductTable.svelte";
    import ProviderTable from "$components/ProviderTable.svelte";
    import { Database } from "$wails/go/backend/App";
    import {
        AppRail,
        AppRailTile,
        AppShell,
        Modal,
        modalStore,
        type ModalComponent,
        type ModalSettings,
    } from "@skeletonlabs/skeleton";
    import { writable, type Writable } from "svelte/store";
    async function ConnectDB() {
        return await Database();
    }
    ConnectDB();

    function modalComponentActions(): void {
        const c: ModalComponent = {
            ref: ProductActions,
        };
        const d: ModalSettings = {
            buttonTextCancel: "X",
            type: "component",
            component: c,
            title: "Custom Form Component",
        };
        modalStore.trigger(d);
    }
    function modalCreateProduct(): void {
        const c: ModalComponent = {
            ref: CreateProductModal,
        };
        const d: ModalSettings = {
            buttonTextCancel: "X",
            type: "component",
            component: c,
            title: "Create Product Modal Component",
        };
        modalStore.trigger(d);
    }
    function modalCreateCategory(): void {
        const c: ModalComponent = {
            ref: CreateCategoryModal,
        };
        const d: ModalSettings = {
            buttonTextCancel: "X",
            type: "component",
            component: c,
            title: "Create Category Modal Component",
        };
        modalStore.trigger(d);
    }
    function modalCreateProvider(): void {
        const c: ModalComponent = {
            ref: CreateProviderModal,
        };
        const d: ModalSettings = {
            buttonTextCancel: "X",
            type: "component",
            component: c,
            title: "Create Provider Modal Component",
        };
        modalStore.trigger(d);
    }
    let items = [
        { label: "Productos", value: 1, component: ProductTable },
        { label: "Categorías", value: 2, component: CategoryTable },
        { label: "Proveedores", value: 3, component: ProviderTable },
    ];
    $: selectedTab = $storeValue;
    const storeValue: Writable<number> = writable(1);

    function handleClick(value: number): void {
        storeValue.set(value);
    }
</script>

<main class="mx-auto">
    <Modal />

    <div>
        <AppShell>
            <svelte:fragment slot="sidebarLeft">
                <AppRail
                    selected={storeValue}
                    width="w-full"
                    background="bg-surface-200"
                    height="h-screen"
                >
                    {#each items as item}
                        <AppRailTile
                            value={item.value}
                            on:click={() => handleClick(item.value)}
                            ><p class="mx-2">{item.label}</p></AppRailTile
                        >
                    {/each}
                </AppRail>
            </svelte:fragment>
            <div class="mx-2 flex flex-col gap-2">
                {#each items as item}
                    {#if item.value === selectedTab}
                        {#if item.value === 1}
                            <div>
                                <button
                                    class="btn  variant-filled-primary"
                                    on:click={modalComponentActions}
                                    >Modificar por Proveedor</button
                                >
                                <button
                                    class="btn  variant-filled-primary"
                                    on:click={modalCreateProduct}
                                    >Agregar Producto <PlusIcon /></button
                                >
                                <button
                                    class="btn  variant-filled-primary"
                                    on:click={modalCreateCategory}
                                    >Agregar Categoria <PlusIcon /></button
                                >
                                <button
                                    class="btn  variant-filled-primary"
                                    on:click={modalCreateProvider}
                                    >Agregar Proveedor<PlusIcon /></button
                                >
                            </div>
                        {/if}

                        <svelte:component this={item.component} />
                    {/if}
                {/each}
            </div>
        </AppShell>
    </div>
</main>
