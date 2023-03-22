<script lang="ts">
    import { loadProducts } from "$lib/store";
    import { UpdateProviderById } from "$wails/go/backend/App";
    import type { backend } from "$wails/go/models";
    import type { ModalProps } from "@skeletonlabs/skeleton/utilities/Modal/Modal.svelte";

    export let parent: ModalProps;
    export let provider: backend.Provider;
    async function updateProvider(e: Event) {
        e.preventDefault();
        await UpdateProviderById(provider.id, provider);
        await loadProducts();
        parent.onClose();
    }
</script>

<form class="flex flex-col gap-2 card p-4 w-modal-slim shadow-xl space-y-4 ">
    <div class="flex justify-end">
        <button
            type="button"
            class="btn-icon btn-icon-sm {parent.buttonNeutral} "
            on:click={parent.onClose}>{parent.buttonTextCancel}</button
        >
    </div>
    <h2>Modificar Proveedor</h2>
    <div>
        <label class="label">
            <span>Nombre del proveedor:</span>
            <input
                type="text"
                bind:value={provider.name}
                placeholder="Arcor..."
                class="input"
            />
        </label>
        <label class="label">
            <span>Número de telefono:</span>
            <input
                type="number"
                placeholder="3462-505050..."
                class="input"
                bind:value={provider.phone_number}
            />
        </label>
    </div>
    <button
        type="button"
        on:click={updateProvider}
        class="w-full btn btn- variant-filled-secondary">Modificar</button
    >
</form>
