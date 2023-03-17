<script lang="ts">
    import { Modal, modalStore, type ModalSettings } from '@skeletonlabs/skeleton';
    let tableArr = [
        {
            position: "Quilmes",
            name: "La PRoveedora",
            symbol: 25,
            weight: "24kg",
        },
    ];
    let totalWeight = "100kg";

const alert: ModalSettings = {
	type: 'alert',
	// Data
	title: 'Example Alert',
	body: 'This is an example modal.',
	image: 'https://i.imgur.com/WOgTG96.gif',
};
const prompt: ModalSettings = {
	type: 'prompt',
	// Data
	title: tableArr[0].position,
	body: 'Nuevo precio.',
	// Populates the input value and attributes
	value: tableArr[0].symbol,
	valueAttr: { type: 'text', minlength: 3, maxlength: 10, required: true },

	// Returns the updated response value
	response: (r: string) => console.log('response:', r),
};

</script>

<!-- Responsive Container (recommended) -->
<div class="table-container">
    <!-- Native Table Element -->
    <table class="table table-hover">
        <thead>
            <tr>
                <th>Producto</th>
                <th>Proveedor</th>
                <th>Precio</th>
                <th>Aumentar</th>
            </tr>
        </thead>
        <tbody>
            {#each tableArr as row, i}
                <tr>
                    <td>{row.position}</td>
                    <td>{row.name}</td>
                    <td>{row.symbol}$</td>
                    <td class="flex items-center">
                        <button on:click={() => modalStore.trigger(prompt)} class='btn variant-filled'>Modificar</button>
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
