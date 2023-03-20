<script lang="ts">
    import { loadProducts, providerStore } from "../lib/store";
    import { onDestroy } from "svelte";
    import type { Provider } from "src/models/products";
    import { UpdateByProvider } from "../../wailsjs/go/backend/App";

    let providerList: Provider[] = [];
    const unsubscribe = providerStore.subscribe((providers) => {
        providerList = providers;
    });
    let provider = "" 
    let percentage = 0;
    async function updateByProvider(e) {
        e.preventDefault()
        await UpdateByProvider(percentage, provider);
        await loadProducts();
    }

    onDestroy(() => {
        unsubscribe();
    });
</script>
<div>{provider}</div>
<form on:submit={updateByProvider}>
    <select class="select" bind:value={provider} >
        <option value="" selected> Seleccionar</option>
        {#each providerList as provider}
            <option value={provider.name}>{provider.name}</option>
        {/each}
    </select>
    <input class="input" type="number" bind:value={percentage} />
    <button class="variant-filled-secondary btn">Aumentar {percentage}%</button>
</form>
