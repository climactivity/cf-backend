import { pb } from '$lib/Services/PocketbaseWrapper';
import { writable } from 'svelte/store';
import type { Week } from '$lib/types/collections';

export const getCurrentWeek = async () => {
	try {
		return await pb.send('api/currentweek', {});
	} catch (e) {
		console.log('hmm');
		return undefined;
	}
};
export const postCurrentWeek = async () => {
	try {
		const week: Week | undefined = await pb.send('api/currentweek', {
			method: 'POST'
		});
		if (week) {
			currentWeek.set(week);
		}
		return week;
	} catch (e) {
		console.log('hmm');
		return undefined;
	}
};

export const currentWeek = writable<Week>(await getCurrentWeek());
// onMount(() => {
//
// 	currentWeekPromise = getCurrentWeek()
// 	currentWeekPromise.then(resolved => currentWeek = resolved)
// 	console.log(currentWeekPromise)
//
//
// })
export const justSetupWeek = writable(false);