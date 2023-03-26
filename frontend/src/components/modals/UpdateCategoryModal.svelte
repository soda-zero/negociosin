<script lang="ts">
    import { validateForm } from "$lib/helper";
    import { CategorySchema, type Category } from "$models/Category";
    import { UpdateCategoryById } from "$wails/go/backend/App";
    import type { backend } from "$wails/go/models";
    import type { ModalProps } from "@skeletonlabs/skeleton/utilities/Modal/Modal.svelte";

    export let parent: ModalProps;
    export let category: backend.Category;

    let errors: Category = {};

    async function updateCategory() {
        const { errors: validationErrors, valid } = validateForm(
            CategorySchema,
            category
        );
        if (!valid) {
            errors = validationErrors;
            return;
        }
        await UpdateCategoryById(category.id, category);
        parent.onClose();
    }
</script>

<form
    class="flex flex-col gap-2 card p-4 w-modal-slim shadow-xl space-y-4"
    on:submit|preventDefault={updateCategory}
>
    <div class="flex justify-end">
        <button
            type="button"
            class="btn-icon btn-icon-sm {parent.buttonNeutral} "
            on:click={parent.onClose}>{parent.buttonTextCancel}</button
        >
    </div>
    <h2>Modificar Categoría</h2>
    <div>
        <label
            class={`${errors.name ? "text-error-500 flex flex-col" : ""} label`}
        >
            <span>Nombre de la categoría</span>
            <input
                type="text"
                bind:value={category.name}
                placeholder="limpieza..."
                class={`${errors.name ? "input-error" : "input"}`}
            />
            {#if errors.name}
                <p class="text-error-500 font-bold  w-fit !text-sm">
                    {errors.name}
                </p>
            {/if}
        </label>

        <label
            class={`${
                errors.profit_percent ? "text-error-500 flex flex-col" : ""
            } label`}
        >
            <span>Porcentaje de la ganancia</span>
            <input
                type="number"
                placeholder="100..."
                class={`${errors.profit_percent ? "input-error" : "input"}`}
                bind:value={category.profit_percent}
                step="0.001"
            />
            {#if errors.profit_percent}
                <p class="text-error-500 font-bold  w-fit !text-sm">
                    {errors.profit_percent}
                </p>
            {/if}
        </label>
    </div>
    <button type="submit" class="w-full btn btn- variant-filled-secondary"
        >Modificar</button
    >
</form>
