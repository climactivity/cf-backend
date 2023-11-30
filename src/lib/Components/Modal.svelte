<script lang="ts">
	import { onMount } from "svelte";

	export let isModalOpen = false;
	let modalElement: HTMLElement;

	const _updateModal = () => {
		if (isModalOpen) {
			modalElement.setAttribute('open', '');
		} else {
			modalElement.removeAttribute('open');
		}
	};

	export const hide = () => {
		isModalOpen = false;
		_updateModal();
	};

	export const show = () => {
		isModalOpen = true;
		_updateModal();
	};

	export const toggleModal = () => {
		isModalOpen = !isModalOpen;
		_updateModal();
	};

    onMount(() => _updateModal())
</script>

<div class="modal items-end transition-all" bind:this={modalElement}>
	<div class="modal-box rounded-b-none grr  w-[91vw]">
		<button class="fixed right-0 top-0 pt-2 px-2" on:click={hide}>
			<svg class="h-11 w-11" viewBox="0 0 100 100">
				<circle
					class="fill-gray-100 focus:fill-gray-300 active:fill-gray-300 hover:fill-gray-300"
					cx="50"
					cy="50"
					r="50"
				/>
				<line
					class="stroke-gray-400 focus:stroke-gray-950 active:stroke-gray-950 hover:stroke-gray-950"
					x1="30"
					y1="30"
					x2="70"
					y2="70"
					stroke-width="8px"
					stroke-linecap="round"
				/>
				<line
					class="stroke-gray-400 focus:stroke-gray-950 active:stroke-gray-950 hover:stroke-gray-950"
					x1="30"
					y1="70"
					x2="70"
					y2="30"
					stroke-width="8px"
					stroke-linecap="round"
				/>
			</svg>
		</button>

		<div class="pt-10 pb-20">
			<slot />
		</div>
	</div>
	<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
	<!-- svelte-ignore a11y-label-has-associated-control -->
	<!-- svelte-ignore a11y-click-events-have-key-events -->
	<label class="modal-backdrop" on:click={hide}>Close</label>
</div>

<style lang="scss">
    .grr {
        --tw-scale-x: 1;
        --tw-scale-y: 1;
        --tw-translate-y: 100vh;
        @apply transition-transform duration-300
    }

</style>
