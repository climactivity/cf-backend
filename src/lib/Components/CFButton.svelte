<script lang="ts">
	export let secondary = false;
	export let ghost = false;
	export let transparent = false;
	export let share = false;
	export let rounded = false;
	export let disabled = false;
	export let onclick: Function = () => console.warn('No onclick handler defined!');
</script>

<button
	on:click={onclick()}
	{disabled}
	class="
    px-8
    py-4
    transition-all
    hover:duration-250
    active:duration-100
    font-semibold
    font-serif
    {rounded ? 'rounded-full' : 'rounded-lg'}
    {disabled
		? 'bg-neutral-100 text-neutral-400 hover:bg-neutral-100'
		: ghost
		? 'bg-transparent outline outline-1 text-neutral-900'
		: transparent
		? 'bg-transparent text-neutral-900 underline'
		: secondary
		? 'bg-secondary-200 hover:bg-secondary-300 text-secondary-900'
		: 'bg-primary-200 hover:bg-primary-300 text-primary-900'}
    {disabled
		? ''
		: transparent
		? 'active:text-opacity-80'
		: 'active:text-opacity-80 active:scale-[0.98] active:translate-y-0.5'}
    "
>
	{#if share}
		<div class="flex flex-row justify-center gap-4">
			<svg
				width="24"
				height="24"
				viewBox="0 0 24 24"
				fill="none"
				xmlns="http://www.w3.org/2000/svg"
			>
				<g id="Icon/Outline/share">
					<path
						id="Icon"
						d="M8.68387 13.3419C8.88616 12.9381 9 12.4824 9 12C9 11.5176 8.88616 11.0619 8.68387 10.6581M8.68387 13.3419C8.19134 14.3251 7.17449 15 6 15C4.34315 15 3 13.6569 3 12C3 10.3431 4.34315 9 6 9C7.17449 9 8.19134 9.67492 8.68387 10.6581M8.68387 13.3419L15.3161 16.6581M8.68387 10.6581L15.3161 7.34193M15.3161 7.34193C15.8087 8.32508 16.8255 9 18 9C19.6569 9 21 7.65685 21 6C21 4.34315 19.6569 3 18 3C16.3431 3 15 4.34315 15 6C15 6.48237 15.1138 6.93815 15.3161 7.34193ZM15.3161 16.6581C15.1138 17.0619 15 17.5176 15 18C15 19.6569 16.3431 21 18 21C19.6569 21 21 19.6569 21 18C21 16.3431 19.6569 15 18 15C16.8255 15 15.8087 15.6749 15.3161 16.6581Z"
						stroke="currentcolor"
						stroke-width="2"
						stroke-linecap="round"
						stroke-linejoin="round"
					/>
				</g>
			</svg>

			<slot />
		</div>
	{:else}
		<slot />
	{/if}
</button>

<style lang="scss">
</style>
