<script lang="ts">
    import { loadProducts } from "$lib/store";
    import { UpdateCategoryById } from "$wails/go/backend/App";
    import type { backend } from "$wails/go/models";
    import type { ModalProps } from "@skeletonlabs/skeleton/utilities/Modal/Modal.svelte";

    export let parent: ModalProps;
    export let category: backend.Category;
    async function updateCategory(e: Event) {
        e.preventDefault();
        await UpdateCategoryById(category.id, category);
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
    <h2>Modificar Categoría</h2>
    <div>
        <label class="label">
            <span>Nombre de la categoría:</span>
            <input
                type="text"
                bind:value={category.name}
                placeholder="limpieza..."
                class="input"
            />
        </label>
        <label class="label">
            <span>Porcentaje de la ganancia:</span>
            <input
                type="number"
                placeholder="30..."
                class="input"
                bind:value={category.profit_percent}
            />
        </label>
    </div>
    <button
        type="button"
        on:click={updateCategory}
        class="w-full btn btn- variant-filled-secondary">Modificar</button
    >
</form>
