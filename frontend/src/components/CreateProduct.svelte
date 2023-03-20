<script lang="ts">
    import type { Product } from "src/models/products";
    import { CreateProduct } from "../../wailsjs/go/backend/App";
    import { loadProducts } from "../lib/store";

    let product: Product = { name: "", provider: "", cost_price: 0 };
    async function createProduct(e) {
        e.preventDefault();
        await CreateProduct(product);
        await loadProducts();
    }
</script>

<form
    class="input-group input-group-divider grid-cols-auto"
    on:submit={createProduct}
>
    <div class="input-group-shim">Agregar productos.</div>
    <input type="text" bind:value={product.name} placeholder="Nombre..." />
    <input
        type="number"
        placeholder="Precio de costo..."
        bind:value={product.cost_price}
    />
    <input
        type="text"
        placeholder="Proveedor..."
        bind:value={product.provider}
    />
    <button class="variant-filled-secondary">Agregar</button>
</form>
