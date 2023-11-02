<script lang="ts">
	import { menuOpen } from '$lib/Services/MenuState';
	import { cubicIn, cubicInOut, cubicOut } from 'svelte/easing';
	import { fade, fly, slide } from 'svelte/transition';

	const closeMenu = () => menuOpen.set(false);
</script>

{#if $menuOpen}
	<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
	<div
		class="fixed left-0 right-0 bottom-0 top-[5.25rem] bg-gray-900 bg-opacity-0 z-10"
		on:click={closeMenu}
		on:keypress={closeMenu}
		role="complementary"
	>
		<nav in:fly={{ x: -250, delay:75, duration:425, easing: cubicOut, opacity: 1 }} out:fly={{ x: -250, duration: 425, opacity: 1, easing: cubicIn}} class="h-full bg-neutral-50 shadow w-fit">
			<ul>
				<li class="text-2xl font-bold"><a href="/app/">Home</a></li>
				<li class="text-2xl font-bold"><a href="/app/history">Verlauf</a></li>
				<li class="text-2xl font-bold"><a href="/app/awards">Awards</a></li>
				<li class="text-2xl font-bold"><a href="/app/stats">Ergebnisse</a></li>
				<li class="text-2xl font-bold"><a href="/app/settings">Einstellungen</a></li>
                <hr class="my-2 ml-4 mr-6"/>
				<li class="text-lg"><a href="/app/commitment">Commitment</a></li>
                <li class="text-lg"><a href="/app/legal">Impressum</a></li>
			</ul>
		</nav>
	</div>
{/if}

<style lang="scss">
	ul {
		@apply gap-4;
	}

    li {
         @apply  font-serif ml-2 mr-4 px-2 py-4
    }
</style>
