<script lang="ts">
	import { mapToRange } from '$lib/Util';
	import { onMount } from 'svelte';
	export let adaptAngle = false;
	export let opacity = 1.0;
	export let angle = -12;

	const aspectRationToAngle = (ratio: number) => {
		angle = mapToRange(ratio, 0.5, 2.0, -12, -5);
	};

	/* Do horrible things with the window aspect ratio to figure out if we should 
	change the angle of the bg gradient to not hide the right hand side of the image */
	onMount(() => {
		if (adaptAngle) {
			onresize = (_event) => {
				if (!adaptAngle) return;
				aspectRationToAngle(window.innerWidth / window.innerHeight);
			};
		}
	});
</script>

<div
	style="--opacity: {opacity}; --angle: {angle}deg"
	class="absolute left-0 top-0 right-0 ws-gradient ws-gradient-top -z-50 w-full"
/>

<style lang="scss">
	.ws-gradient {
		opacity: var(--opacity, 1);
		background-blend-mode: overlay;
		background-size: 100%;
		background-repeat: inherit;
		--background-color: rgba(255, 255, 255, 1);
		height: 20rem;
	}
	.ws-gradient-top {
		background-image: url('/warming_stripes.png'),
			linear-gradient(var(--angle), var(--background-color) 50%, rgba(131, 58, 180, 0) 100%);
	}

	.ws-gradient-bottom {
		background-image: url('/warming_stripes.png'),
			linear-gradient(
				calc(180deg + var(--angle)),
				var(--background-color) 50%,
				rgba(131, 58, 180, 0) 100%
			);
	}
</style>
