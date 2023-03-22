<script lang="ts">
    import { loadProducts, providerStore } from "../lib/store";
    import { onDestroy } from "svelte";
    import { UpdateByProvider } from "../../wailsjs/go/backend/App";
    import type { backend } from "wailsjs/go/models";
    import type { ModalProps } from "@skeletonlabs/skeleton/utilities/Modal/Modal.svelte";

    export let parent: ModalProps;

    let providerList: backend.Provider[] = [];
    const unsubscribe = providerStore.subscribe((providers) => {
        providerList = providers;
    });
    function resetForm() {
        provider = undefined;
        percentage = 0;
    }
    let provider = undefined;
    let percentage = 0;
    async function updateByProvider(e: Event) {
        e.preventDefault();
        await UpdateByProvider(percentage, provider);
        await loadProducts();
        resetForm();
        parent.onClose();
    }

    onDestroy(() => {
        unsubscribe();
    });
</script>

<form
    on:submit={updateByProvider}
    class="flex flex-col gap-2 card p-4 w-modal shadow-xl space-y-4 "
>
    <div class="flex justify-end">
        <button
            class="btn-icon btn-icon-sm {parent.buttonNeutral} "
            on:click={parent.onClose}>{parent.buttonTextCancel}</button
        >
    </div>
    <h2>Modificar precio a todos los productos con el mismo Proveedor</h2>
    <div class="flex flex-col gap-2">
        <label class="label">
            <span>Proveedor</span>
            <select class="select" bind:value={provider}>
                <option value="disabled" selected disabled> Seleccionar</option>
                {#if providerList != null}
                    {#each providerList as provider}
                        <option value={provider.name}>{provider.name}</option>
                    {/each}
                {/if}
            </select>
        </label>
        <label class="label">
            <span>Porcentaje</span>
            <input class="input" type="number" bind:value={percentage} />
        </label>
    </div>
    <button class="w-full variant-filled-secondary btn"
        >Aumentar precio en {percentage | 0}%</button
    >
</form>
