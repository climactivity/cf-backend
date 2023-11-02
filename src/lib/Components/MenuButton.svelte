<script lang="ts">
	import { menuOpen } from '$lib/Services/MenuState';

	const toggleMenu = () => {
		menuOpen.set(!$menuOpen);
	};
</script>

<button aria-controls="primary-navigation" class="w-12 h-12 z-20" on:click={toggleMenu}>
	<svg
		viewBox="0 0 26.458333 26.458333"
		xmlns="http://www.w3.org/2000/svg"
		class=" w-12 h-12 transition-all duration-500 {$menuOpen ? 'expanded' : 'collapsed'}"
	>
		<g id="layer1">
			<path
				id="bottom2"
				class="line2"
				d="M 3.7732683,18.978041 H 22.68511 c 0,0 5.067961,-0.01076 2.149154,4.017304 -1.814915,2.504651 -4.323456,-0.80244 -4.323456,-0.80244 L 6.5428366,6.4729796"
			/>
			<path id="middle2" class="line2" d="M 3.7732681,13.229167 H 22.685065" />
			<path
				id="top2"
				class="line2"
				d="M 3.7732683,7.4104353 H 22.68511 c 0,0 5.067961,0.010758 2.149154,-4.0173039 -1.814915,-2.5046506 -4.323456,0.8024401 -4.323456,0.8024401 L 6.5428366,19.915497"
			/>
			<path
				id="bottom"
				class="line"
				d="M 3.7732683,18.978041 H 22.68511 c 0,0 5.067961,-0.01076 2.149154,4.017304 -1.814915,2.504651 -4.323456,-0.80244 -4.323456,-0.80244 L 6.5428366,6.4729796"
			/>
			<path id="middle" class="line" d="M 3.7732681,13.229167 H 22.685065" />
			<path
				id="top"
				class="line"
				d="M 3.7732683,7.4104353 H 22.68511 c 0,0 5.067961,0.010758 2.149154,-4.0173039 -1.814915,-2.5046506 -4.323456,0.8024401 -4.323456,0.8024401 L 6.5428366,19.915497"
			/>
		</g>
	</svg>
</button>

<style lang="scss">
	.line {
		fill: none;
		stroke: white;
		stroke-width: 2px;
		stroke-linecap: round;
		stroke-linejoin: miter;
		stroke-miterlimit: 4;
		stroke-dasharray: 52;
		stroke-opacity: 1;
		outline: #e40045;
	}

	.line2 {
		fill: none;
		stroke: black;
		stroke-width: 3px;
		stroke-linecap: round;
		stroke-linejoin: miter;
		stroke-miterlimit: 4;
		stroke-dasharray: 52;
		stroke-opacity: 1;
		outline: #e40045;
	}

	svg {
		stroke: #000;
	}

	.collapsed {
		stroke: #000;

		stroke-dashoffset: 136.9;
	}
	.expanded {
		stroke: #e40045;
		stroke-dashoffset: 71.1;
	}
</style>
