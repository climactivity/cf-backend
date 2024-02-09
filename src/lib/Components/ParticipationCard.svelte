<script lang="ts">
	import { optionToString } from '$lib/Util';
	import { DateTime } from 'luxon';
	import type { Participation } from '$lib/types/collections';

	export let participation: Participation;
	let dateLabel: string = DateTime.fromSQL(participation.date).toFormat('dd.MM.YY');
</script>

<div style="grid-template-columns: 5rem 1fr;" class="grid grid-flow-col items-center py-2 linething relative">
	<div class="date-label rounded-full bg-primary-300 text-white font-bold px-2 py-1 text-center">
		{dateLabel}
	</div>
	<div class="pl-6 h-12 flex flex-col justify-center">
		{#if participation.note}
			{participation.note}
		{:else}
			Ich habe mit {optionToString(participation.option)} teilgenommen
		{/if}
	</div>
</div>

<!-- 
<div>
    <div>Option: {participation.option}</div>
    <div>Note: {participation.note}</div>
    <div>Weeknumber: {participation.weeknumber}</div>
    <div>date: {participation.date}</div>


</div> -->

<style lang="scss">
	.date-label {
		font-family: JetBrains Mono;
		font-weight: 700;
	}
    .linething::before {
        position: absolute;
        content: '';
        top: 0;
        bottom: -1rem;
        left:2.4rem;
        width: 0.2rem;
        background-color: black;
        z-index: -1;
    }
</style>
