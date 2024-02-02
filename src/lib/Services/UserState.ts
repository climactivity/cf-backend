import { writable } from 'svelte/store';

import { currentUser, pb } from './PocketbaseWrapper';
import type { Participation, UserRecord } from '$lib/types/collections';
import type { RecordModel } from 'pocketbase';
import { DateTime } from 'luxon';

export const participationCount = writable(0);
export const streakCount = writable(0);

export const notificationsSetup = writable({
	notification_setup: false,
	notification_push: false,
	notification_email: false
});

export const currentParticipation = writable<RecordModel|null>(null);

let unsubscribe: () => void;
let _user: Partial<UserRecord | RecordModel>;

const updateCounters = async () => {
	if (_user === null) return;

	const _resParticipations = await pb.collection('participations_count').getOne(_user.id || '');
	const _resStreak = await pb.collection('streak_count').getOne(_user.id || '');

	participationCount.set(_resParticipations.count);
	streakCount.set(_resStreak.length);
};

const updateUserState = async (user: UserRecord) => {
	user = await pb.collection('users').getOne(_user.id || '');

	console.log('notificationsSetup', user);
	notificationsSetup.set(
		{
			notification_setup: user.notification_setup,
			notification_push: user.notification_push,
			notification_email: user.notification_email
		}
	);
};

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const participationsUpdate: (params: { action: string, record: string }) => void = ({ action }) => {
	console.log('updated!', action);
	if (action === 'create' || action === 'delete') {
		updateCounters();
	}

	updateCurrentParticipation()
};


const updateCurrentParticipation = async () => {
	try {
		const weeknumber = DateTime.now().weekNumber;
		currentParticipation.set(await pb.collection('participations').getFirstListItem<Participation>(pb.filter('weeknumber = {:weeknumber}', { weeknumber })));
	} catch (e) {
		currentParticipation.set(null)
	}

};


const userUpdate: (params: { action: string, record: string }) => Promise<void> = async ({ action }) => {
	console.log('updated user!', action);
	//currentUserRecord.set(await fetchUserRecordFromAuthModel())
};


currentUser.subscribe(async (user) => {
	if (user === null) return;
	_user = user;
	await updateCounters();
	await updateCurrentParticipation();
	await updateUserState(user);
	unsubscribe = await pb.collection('participations').subscribe('*', participationsUpdate);

	//await pb.collection('users').subscribe(pb.authStore.model?.id, userUpdate);

	console.log('Store updated!');
});


// const fetchUserRecordFromAuthModel = async () => {
//     try {
//         const rec =  await pb.collection("users").getOne(pb.authStore.model?.id);
//         currentUserRecord.set(rec)
//     } catch (e) {
//         console.log(e)
//         currentUserRecord.set(null)
//     }
//
// }
//
// export const currentUserRecord: Writable<RecordModel | null> = writable(null)

export const destroy = () => {
	if (unsubscribe) unsubscribe();
};

