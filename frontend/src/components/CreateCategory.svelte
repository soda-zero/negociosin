<script lang="ts">
    import { CreateCategory } from "../../wailsjs/go/backend/App";
    import { loadProducts } from "../lib/store";
    import { backend } from "../../wailsjs/go/models";
    import type { ModalProps } from "@skeletonlabs/skeleton/utilities/Modal/Modal.svelte";

    export let parent: ModalProps;

    function resetForm() {
        category.name = undefined;
        category.profit_percent = undefined;
    }

    let category = new backend.Category();
    async function createCategory(e: Event) {
        e.preventDefault();
        await CreateCategory(category);
        await loadProducts();
        resetForm();
        parent.onClose();
    }
</script>

<form
    on:submit={createCategory}
    class="flex flex-col gap-2 card p-4 w-modal-slim shadow-xl space-y-4 "
>
    <div class="flex justify-end">
        <button
            class="btn-icon btn-icon-sm {parent.buttonNeutral} "
            on:click={parent.onClose}>{parent.buttonTextCancel}</button
        >
    </div>
    <h2>Agregar Categoría</h2>
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
    <button class="w-full btn btn- variant-filled-secondary">Agregar</button>
</form>
