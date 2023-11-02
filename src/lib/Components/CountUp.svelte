<script lang="ts">
	import { onMount } from 'svelte';

	export let number = 0;
	export let initialValue = 0;

	let value = 0;

	const MAX_DURATION = 3500;
	let duration = MAX_DURATION;

    export const updateNumber = (newNum: number) => {
        _update(newNum, value);
    }

    const _update = (newVal: number, oldVal: number) => {
        value = newVal;

		let delta = newVal - oldVal;
		duration = Math.min(delta * 100, MAX_DURATION); 
    }

	onMount(() => _update(number, initialValue));
</script>

<div class="text-center font-mono">
	<span style="--num: {value}; --duration: {duration}ms;" class="counter text-end font-bold" />
</div>

<style lang="scss">
	@property --num {
		syntax: '<integer>';
		initial-value: 0;
		inherits: false;
	}

	.counter {
		--duration: 5s;
		transition: --num var(--duration) ease-out;
		counter-set: num var(--num);
	}

	.counter::after {
		content: counter(num);
	}

</style>
