<script>
	import CFButton from '$lib/Components/CFButton.svelte';
	import { pb } from '$lib/Services/PocketbaseWrapper';

	const fun = () => {
		pb.send('/currentweek', {
			method: 'POST'
		});
	};

	const fun2 = () => {
		pb.send('/currentweek', {
			method: 'POST',
			body: {
				'remind': 'remind'
			}
		});
	};

</script>
<div class="p-8">
	<CFButton onclick={fun}> Dont Remind</CFButton>
</div>

<div class="p-8">
	<CFButton onclick={fun2}> Remind</CFButton>
</div>