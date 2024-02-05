<script lang="ts">

	import { fade, fly } from 'svelte/transition';

	export let value: string;
	export let update: (value) => Promise<string>;
	export let id: string;
	export let name: string;
	export let isPass = false;
	export let placeholder: string;
	let editing = false;
	let initial_value: string;

	function cancel() {
		editing = false;
		value = initial_value;
	}

	function edit() {
		editing = true;
		initial_value = value;
		if (isPass) {
			value = '';
		}
	}

	async function save() {
		try {
			value = await update(value);
			initial_value = value;
			editing = false;
			success = true;
			setTimeout(() => success = false, 5000);
		} catch (e) {
			error = e as string;
		}

	}

	let error: string;
	let success = false;

	export let successMessage = 'Gespeichert';
</script>

<fieldset>
	<label class="m-2 mb-1 flex-row flex gap-4" for={id}>
		<span class="label-text">{name}</span>
		{#if success}
			<span transition:fade class="label-text text-success-700">{successMessage}</span>
		{/if}
	</label>

	{#if error}
		<label class="label" for={id}>
			<span class="label-text">{error}</span>
		</label>
	{/if}

	{#if isPass}
		<input
			type="password"
			placeholder={placeholder}
			class="input input-bordered w-[16rem]"
			id={id}
			bind:value={value}
			disabled={!editing}
		/>
	{:else}
		<input
			type="text"
			placeholder={placeholder}
			class="input input-bordered w-[16rem]"
			id={id}
			bind:value={value}
			disabled={!editing}
		/>
	{/if}

	{#if editing}
		<button in:fly={{delay:200, x: -10}} out:fly={{ x: -10, duration: 200}} on:click={cancel}>
			<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
					 class="w-6 h-6">
				<path stroke-linecap="round" stroke-linejoin="round" d="M6 18 18 6M6 6l12 12" />
			</svg>
		</button>
		<button in:fly={{delay:200, x: -40}} out:fly={{ x: -40, duration: 200}} on:click={save}>
			<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
					 class="w-6 h-6">
				<path stroke-linecap="round" stroke-linejoin="round" d="m4.5 12.75 6 6 9-13.5" />
			</svg>
		</button>
	{:else}
		<button in:fly={{delay:200, x: 10}} out:fly={{ x: 10, duration: 200}} on:click={edit}>
			<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
					 class="w-6 h-6">
				<path stroke-linecap="round" stroke-linejoin="round"
							d="m16.862 4.487 1.687-1.688a1.875 1.875 0 1 1 2.652 2.652L6.832 19.82a4.5 4.5 0 0 1-1.897 1.13l-2.685.8.8-2.685a4.5 4.5 0 0 1 1.13-1.897L16.863 4.487Zm0 0L19.5 7.125" />
			</svg>
		</button>
	{/if}


</fieldset>