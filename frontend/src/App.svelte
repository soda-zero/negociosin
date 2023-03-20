<script lang="ts">
    import { Database } from "../wailsjs/go/backend/App";
    import CreateProduct from "./components/CreateProduct.svelte";
    import ProductActions from "./components/ProductActions.svelte";
    import ProductTable from "./components/ProductTable.svelte";
    import { AppRail, AppRailTile, AppShell } from "@skeletonlabs/skeleton";
    async function ConnectDB() {
        return await Database();
    }
    ConnectDB();
</script>

<main class="mx-auto">
    <div>
        <AppShell>
            <svelte:fragment slot="sidebarLeft">
                <AppRail>
                    <AppRailTile label="Tile 1" value={0}>(icon)</AppRailTile>
                    <AppRailTile label="Tile 2" value={1}>(icon)</AppRailTile>
                    <AppRailTile label="Tile 3" value={2}>(icon)</AppRailTile>
                </AppRail>
            </svelte:fragment>
            <slot />
            <ProductActions />
            <ProductTable />
            <CreateProduct />
        </AppShell>
    </div>
</main>
