<script lang="ts">
    import {
        CreateProduct,
        Database,
        GetAllProducts,
    } from "../wailsjs/go/backend/App";
    import type { Product } from "./models/products";

    let products = [];
    let prod: Product = { provider: "", cost_price: 0, name: "" };
    async function getProducts() {
        products = await GetAllProducts();
    }
    async function createProduct() {
        await CreateProduct(prod);
        products = await GetAllProducts();
        prod = {name: "", cost_price: 0, provider: ""}
    }
    async function connectDb() {
        return await Database();
    }
    connectDb();
</script>

<main class="  flex-col flex">
    <button on:click={getProducts}>Click to Fetch</button>

    <h3>Create Productos</h3>

    <div class="flex flex-col">
        <label for="name">Nombre:</label>
        <input
            autocomplete="off"
            bind:value={prod.name}
            class="input"
            id="name"
            type="text"
        />
        <label for="provider">Provider:</label>
        <input
            autocomplete="off"
            bind:value={prod.provider}
            class="input"
            id="provider"
            type="text"
        />
        <label for="precio">Precio</label>
        <input
            autocomplete="off"
            bind:value={prod.cost_price}
            class="input"
            id="prceio"
            type="number"
        />
        <button on:click={createProduct}>AddProd</button>
    </div>
    <!---->
    {#each products as product}<div>{product.name}</div>{/each}
    <!-- <AppRail selected={storeValue} height="h-screen"> -->
    <!--     <div class="flex items-center justify-center p-2"> -->
    <!--         <LightSwitch /> -->
    <!--     </div> -->
    <!--     <AppRailTile label="Tile 1" value={0}>(icon)</AppRailTile> -->
    <!--     <AppRailTile label="Tile 2" value={1}>(icon)</AppRailTile> -->
    <!--     <AppRailTile label="Tile 3" value={2}>(icon)</AppRailTile> -->
    <!-- </AppRail> -->
    <!-- <div class="flex gap-4  w-full  m-2"> -->
    <!--     <Table /> -->
    <!-- </div> -->
</main>
