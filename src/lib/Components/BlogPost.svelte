<script lang="ts">
	import type { Post } from '$lib/types/post';
	import { DateTime } from 'luxon';
	import { Browser } from '@capacitor/browser';

	export let post: Post;

	const openInappBrowser = async () => {
		await Browser.open({ url: post.link });
	};

	const getDateString = (date: string) => {
		return DateTime.fromISO(date).setLocale('de').toLocaleString(DateTime.DATE_FULL);
	};

	// i should put this on npm
	const estimateReadingTime = (artice: string) => {
		const wpm = 250; // chosen by fair dice roll
		const words = artice.trim().split(/\s+/).length;

		return `~${Math.ceil(words / wpm)} min Lesezeit`;
	};
</script>

<!-- svelte-ignore empty-block -->
<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div class="blog-post mt-2 bg-gray-50 rounded-sm shadow-md " on:click={openInappBrowser}>
	{#if post.featured_media !== 0}
		<img src={post._embedded['wp:featuredmedia'][0].source_url} alt="Post img" class="w-24" />
	{:else}
		<img src="/images/not-found.webp" alt="Post img" class="w-24" />
	{/if}
	<div class="p-2 h-32">
		<h1 class="title break-words ">
			{@html post.title.rendered.slice(0, 50)}
			{#if post.title.rendered.length >= 50}
				<span class="">...</span>
			{/if}
		</h1>

		<div class="flex flex-col text-sm text-gray-500">
			<span>{estimateReadingTime(post.content.rendered)}</span>
			<span>{getDateString(post.date)}</span>
		</div>
		<!-- 
			<p>
				{@html post.excerpt.rendered.slice(0, 75)}
				{#if post.excerpt.rendered.length >= 75}
					<span class="">...</span>
				{/if}
			</p> -->
	</div>
</div>

<style>
	.blog-post {
		cursor: pointer;
		display: flex;
		align-items: center;
		width: 328px;
		height: 119px;
	}
	img {
		width: 156px;
		height: 119px;
	}
</style>
