<script lang="ts">
    import CategoryTable from "$components/CategoryTable.svelte";
    import ProductTable from "$components/ProductTable.svelte";
    import ProviderTable from "$components/ProviderTable.svelte";
    import { loadCategory, loadProducts, loadProviders } from "$lib/store";
    import { Database } from "$wails/go/backend/App";
    import {
        Modal,
    } from "@skeletonlabs/skeleton";
    import { onMount } from "svelte";
    import { writable, type Writable } from "svelte/store";
    async function ConnectDB() {
        return await Database();
    }
    ConnectDB();

    let items = [
        { label: "Productos", value: 1, component: ProductTable },
        { label: "Categorías", value: 2, component: CategoryTable },
        { label: "Proveedores", value: 3, component: ProviderTable },
    ];
    $: selectedTab = $storeValue;
    const storeValue: Writable<number> = writable(1);

    async function handleClick(value: number): Promise<void> {
        storeValue.set(value);
        await loadProducts();
    }
    onMount(async () => {
        await loadProducts();
        await loadCategory();
        await loadProviders();
    });
</script>

<main class="mx-auto">
    <Modal />

    <div>
        <div class="flex flex-row">
            <div
                class="h-screen flex flex-col gap-2 bg-surface-100 p-1 min-w-[180px]"
            >
                {#each items as item}
                    <button
                        type="button"
                        class={`btn ${
                            selectedTab === item.value
                                ? "variant-filled-primary"
                                : ""
                        }`}
                        on:click={() => handleClick(item.value)}
                    >
                        {item.label}
                    </button>
                {/each}
            </div>
            <div class="my-1 mx-2 w-full flex flex-col gap-2 overflow-auto">
                {#each items as item}
                    {#if item.value === selectedTab}
                        <svelte:component this={item.component} />
                    {/if}
                {/each}
            </div>
        </div>
    </div>
</main>
