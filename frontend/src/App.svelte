<script lang="ts">
  import { AppRail, AppRailTile, LightSwitch, Modal } from "@skeletonlabs/skeleton";
    import { writable, type Writable } from "svelte/store";
    import {Database } from "../wailsjs/go/backend/App"
  import Table from "./components/Table.svelte";
  const storeValue: Writable<number> = writable(0);

  async function connectDb() {
    return await Database()
  }
  connectDb()
  console.log(connectDb)
</script>

<main class=" flex">
  <Modal />
    <AppRail selected={storeValue} height='h-screen'>
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
