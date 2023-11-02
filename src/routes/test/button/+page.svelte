<script>
	import CfButton from '$lib/Components/CFButton.svelte';
</script>

<div class="form-control px-6 py-4 gap-4">
	<CfButton onclick={() => alert('hi')}>Primary</CfButton>
	<CfButton onclick={() => {}} share>Primary Share</CfButton>
	<CfButton rounded>Primary Rounded</CfButton>
	<CfButton share rounded>Primary Rounded Share</CfButton>
	<CfButton disabled onclick={() => alert('hi')}>Primary Disabled</CfButton>
	<CfButton disabled share>Primary Disabled Share</CfButton>
	<CfButton disabled rounded>Primary Disabled Rounded</CfButton>
	<CfButton disabled share rounded>Primary Disabled Rounded Share</CfButton>
	<CfButton secondary>Secondary</CfButton>
	<CfButton secondary share>Secondary Share</CfButton>
	<CfButton secondary rounded>Secondary Rounded</CfButton>
	<CfButton secondary rounded share>Secondary Rounded Share</CfButton>
	<CfButton secondary disabled>Secondary Disabled</CfButton>
	<CfButton secondary disabled share>Secondary Disabled Share</CfButton>
	<CfButton secondary disabled rounded>Secondary Disabled Rounded</CfButton>
	<CfButton secondary disabled rounded share>Secondary Disabled Rounded Share</CfButton>
	<CfButton ghost>Ghost</CfButton>
	<CfButton transparent>Transparent</CfButton>
</div>
