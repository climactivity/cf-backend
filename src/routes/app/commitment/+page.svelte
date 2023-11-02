<script>
	import CfHeading from '$lib/Components/CFHeading.svelte';
	import ClimateFriday from '$lib/Components/ClimateFriday.svelte';
</script>

<!-- TODO make the background more like the mockup -->
<h1 class="text-4xl font-semibold font-serif -rotate-12 pt-8 pb-16">Das Commitment <ClimateFriday/></h1>

<div>
	Freitags zeigen wir Flagge für Tempo und Gerechtigkeit beim Klimaschutz. Wir fordern dies von den
	Regierungen und sind selber dazu bereit. Deshalb

	<ul class="list-disc">
		<li>zeigen wir am Freitag die Warming Stripes,</li>
		<li>
			fordern wir eine sozial gerechte Klimaschutzpolitik im Einklang mit dem Klimaschutz-Abkommen
			von Paris
		</li>
		<li>ermitteln und verkleinern wir unseren eigenen Fußabdruck.</li>
	</ul>
</div>

<style lang="scss">
	li {
		@apply ml-8 my-4
	}
</style>