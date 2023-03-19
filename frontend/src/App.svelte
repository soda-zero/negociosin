<script lang="ts">
    import { Database } from "../wailsjs/go/backend/App";
    import ProductTable from "./components/ProductTable.svelte";
    import {
        AppBar,
        AppRail,
        AppRailTile,
        AppShell,
        LightSwitch,
    } from "@skeletonlabs/skeleton";
    async function ConnectDB() {
        return await Database();
    }
    ConnectDB();
</script>

<main class="mx-auto">
    <AppShell>
        <svelte:fragment slot="header">
            <AppBar>
                <svelte:fragment slot="lead">{" "}</svelte:fragment>
                <svelte:fragment slot="trail"><LightSwitch /></svelte:fragment>
            </AppBar>
        </svelte:fragment>
        <svelte:fragment slot="sidebarLeft">
            <AppRail>
                <AppRailTile label="Tile 1" value={0}>(icon)</AppRailTile>
                <AppRailTile label="Tile 2" value={1}>(icon)</AppRailTile>
                <AppRailTile label="Tile 3" value={2}>(icon)</AppRailTile>
            </AppRail>
        </svelte:fragment>
        <slot />
        <ProductTable />
    </AppShell>
</main>
