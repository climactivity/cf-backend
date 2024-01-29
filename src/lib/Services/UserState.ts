import {writable} from "svelte/store";

import { currentUser, pb } from "./PocketbaseWrapper";
import type {RecordModel} from "pocketbase";

export const participationCount = writable(0);
export const streakCount = writable(0);

export const notificationsSetup = writable({ notification_setup: false, notification_push: false, notification_email: false});

let unsubscribe: () => void; 
let _user: Partial<UserRecord>;

const updateCounters = async () => {
    if (_user === null) return;  

    const _resParticipations = await pb.collection('participations_count').getOne(_user.id || "");
    const _resStreak = await pb.collection('streak_count').getOne(_user.id || "");

    participationCount.set(_resParticipations.count); 
    streakCount.set(_resStreak.length);
}

const updateUserState = async (user) => {
    user = await pb.collection('users').getOne(_user.id || "");

    console.log("notificationsSetup", user)
    notificationsSetup.set(
      {
          notification_setup: user.notification_setup,
          notification_push: user.notification_push,
          notification_email: user.notification_email
      }

    )
}

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const participationsUpdate: (params: {action: string, record: string}) => void = ({ action }) => {
    console.log('updated!', action);
    if(action === "create" || action === "delete") {
        updateCounters(); 
    }
};

const userUpdate: (params: {action: string, record: string}) => Promise<void> = async ({ action }) => {
    console.log('updated user!', action);
    //currentUserRecord.set(await fetchUserRecordFromAuthModel())
};


currentUser.subscribe( async user => {
    if (user === null) return;  
    _user = user; 
    await updateCounters();
    await updateUserState(user);
    unsubscribe = await pb.collection('participations').subscribe("*", participationsUpdate);

    //await pb.collection('users').subscribe(pb.authStore.model?.id, userUpdate);

    console.log("Store updated!");
})


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
}

