<script>
	import { Capacitor, CapacitorHttp } from '@capacitor/core';
	import BlogPost from './BlogPost.svelte';
	import BlogSekelton from './BlogSkeleton.svelte';
	import { fade } from 'svelte/transition';
	import CfSectionHeading from './CFSectionHeading.svelte';

	const postsEndpoint = "https://climactivity.de/wp-json/wp/v2/posts?per_page=1&_embed"

	const fetchPostsGeneric = fetch(postsEndpoint)
		.then((res) => res.json())
		.catch((error) => console.error(error));

		const fetchPostsCapacitor = () => {
		console.log('using capactior http...');
		const options = {
			url: postsEndpoint
		};

		return CapacitorHttp.get(options)
			.then((res) => res.data)
			.catch((error) => console.error('Error loading blog with Capacitor/http:', error));
	};

	const fetchPosts = Capacitor.isNativePlatform() ? fetchPostsCapacitor() : fetchPostsGeneric;
</script>

<CfSectionHeading>Neuigkeiten</CfSectionHeading>
<div class="">
	{#await fetchPosts}
		<BlogSekelton />
	{:then posts}
	<div class="w-80" transition:fade={{delay: 250}}>

		{#each posts as post}
				<BlogPost {post} />
		{/each}
	</div>

	{:catch error}
		<div>Error {error}</div>
	{/await}
</div>
