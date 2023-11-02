<script>
// @ts-nocheck

	import CfButton from "$lib/Components/CFButton.svelte";
import CountUp from "$lib/Components/CountUp.svelte";

let updateNumber

</script>
<CountUp bind:updateNumber number={100}/>

<CfButton onclick={() => updateNumber && updateNumber(120)}>
    +++
</CfButton>

<CfButton onclick={() => updateNumber && updateNumber(150)}>
    +++
</CfButton>