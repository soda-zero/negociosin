<script lang="ts">
    import { validateForm } from "$lib/helper";
    import { ProviderSchema, type Provider } from "$models/ProviderSchema";
    import { CreateProvider } from "$wails/go/backend/App";
    import { backend } from "$wails/go/models";
    import type { ModalProps } from "@skeletonlabs/skeleton/utilities/Modal/Modal.svelte";

    export let parent: ModalProps;

    let provider = new backend.Provider();
    let errors: Provider = {};
    async function createProvider() {
        const { errors: validationErrors, valid } = validateForm(
            ProviderSchema,
            provider
        );
        if (!valid) {
            errors = validationErrors;
            return;
        }
        await CreateProvider(provider);
        parent.onClose();
    }
</script>

<form
    class="flex flex-col gap-2 card p-4 w-modal-slim shadow-xl space-y-4 "
    on:submit|preventDefault={createProvider}
>
    <div class="flex justify-end">
        <button
            type="button"
            class="btn-icon btn-icon-sm {parent.buttonNeutral} "
            on:click={parent.onClose}>{parent.buttonTextCancel}</button
        >
    </div>
    <h2>Modificar Proveedor</h2>
    <div>
        <label
            class={`${errors.name ? "text-error-500 flex flex-col" : ""} label`}
        >
            <span>Nombre del proveedor</span>
            <input
                type="text"
                bind:value={provider.name}
                placeholder="arcor..."
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
                errors.phone_number ? "text-error-500 flex flex-col" : ""
            } label`}
        >
            <span>Número de telefono del proveedor</span>
            <input
                type="text"
                placeholder="3462-666999..."
                class={`${errors.phone_number ? "input-error" : "input"}`}
                bind:value={provider.phone_number}
                step="0.001"
            />
            {#if errors.phone_number}
                <p class="text-error-500 font-bold  w-fit !text-sm">
                    {errors.phone_number}
                </p>
            {/if}
        </label>
    </div>
    <button type="submit" class="w-full btn btn- variant-filled-secondary"
        >Modificar</button
    >
</form>
