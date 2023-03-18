<script lang="ts">
    import {
        AppRail,
        AppRailTile,
        LightSwitch,
        Modal,
    } from "@skeletonlabs/skeleton";
    import { writable, type Writable } from "svelte/store";
    import { Database, GetInfo } from "../wailsjs/go/backend/App";
    import Table from "./components/Table.svelte";
    const storeValue: Writable<number> = writable(0);

    let info = "";
    async function connectDb() {
        return await Database();
    }
    async function Getdb() {
        info = await GetInfo();
    }

    connectDb();
    Getdb();
</script>

<main class=" flex">
    <Modal />
    <div>{info}</div>
    <AppRail selected={storeValue} height="h-screen">
        <div class="flex items-center justify-center p-2">
            <LightSwitch />
        </div>
        <AppRailTile label="Tile 1" value={0}>(icon)</AppRailTile>
        <AppRailTile label="Tile 2" value={1}>(icon)</AppRailTile>
        <AppRailTile label="Tile 3" value={2}>(icon)</AppRailTile>
    </AppRail>
    <div class="flex gap-4  w-full  m-2">
        <Table />
    </div>
</main>
