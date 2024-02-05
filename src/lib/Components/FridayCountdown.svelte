<script lang="ts">
	import ClimateFriday from './ClimateFriday.svelte';
	import { isFriday, timeToNextFriday } from '$lib/Services/TimerStore';

	const pad = (x: number, digits = 2) => {
		return x.toString().padStart(digits, '0');
	};

	let days: string, hours: string, minutes: string, seconds: string;
	$: days = pad($timeToNextFriday.days || 0, 0);
	$: hours = pad($timeToNextFriday.hours || 0);
	$: minutes = pad($timeToNextFriday.minutes || 0);
	$: seconds = pad(Math.floor($timeToNextFriday.seconds || 0));
</script>


{#if !isFriday()}
	<div class="text-center text-lg gap-2 flex flex-col flex-1">
		<div>
			Nur noch
		</div>
		<div class="mx-auto font-bold font-mono countdown2 flex-row flex justify-around align-middle w-full">
			<div class="flex flex-col"><span>{days}</span> Tage</div>
			:
			<div class="flex flex-col"><span>{hours}</span> Stunden</div>
			:
			<div class="flex flex-col"><span>{minutes}</span> Minuten</div>
			:
			<div class="flex flex-col"><span>{seconds}</span> Sekunden
			</div>
		</div>
		<div>
			bis zum nächsten
			<ClimateFriday />
		</div>
	</div>
{/if}

<style>
    .countdown2 > span {

    }
</style>