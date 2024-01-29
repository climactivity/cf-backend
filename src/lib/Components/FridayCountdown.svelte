<script lang="ts">
	import ClimateFriday from './ClimateFriday.svelte';
	import { isFriday, timeToNextFriday } from '$lib/Services/TimerStore';

	const pad = (x: number, digits = 2) => {
		return x.toString().padStart(digits,'0');
	}

    let days: string, hours: string, minutes: string, seconds: string;
    $: days = pad($timeToNextFriday.days || 0,0)
    $: hours = pad($timeToNextFriday.hours || 0)
    $: minutes = pad($timeToNextFriday.minutes || 0)
    $: seconds = pad(Math.floor($timeToNextFriday.seconds || 0))
</script>


{#if !isFriday()} 
<div class="text-center">
	<div class="mx-auto font-bold font-mono countdown2"> <span>{days}</span> : <span>{hours}</span> : <span>{minutes}</span> : <span>{seconds}</span></div>
	bis zum nächsten <ClimateFriday/>
</div>
{/if}

<style>
	.countdown2 > span {

	}
</style>