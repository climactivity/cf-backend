<script>
	import CfSectionHeading from "./CFSectionHeading.svelte";

</script>

<div>
    <CfSectionHeading><slot name="heading"/></CfSectionHeading>
    <slot/>
</div>