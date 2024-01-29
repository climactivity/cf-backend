<script>
	import '../app.css';
</script>

<div class="app">
	<slot />
</div>

<style lang="scss">
	body {
	/* Other styles above... */
	}
	.app {
		display: flex;
		flex-direction: column;
		min-height: 100vh;
		padding: env(safe-area-inset-top) env(safe-area-inset-right) env(safe-area-inset-bottom) env(safe-area-inset-left);

		@apply text-neutral-900;
	}
</style>
