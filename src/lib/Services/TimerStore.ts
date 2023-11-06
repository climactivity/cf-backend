import { DateTime } from 'luxon';
import { writable } from 'svelte/store';

const timerAccuracy = 1000; // how often to tick in ms

const getCountdownValue = () => {
    const now = DateTime.now();
    const friday = DateTime.fromObject({
        weekYear: now.weekYear,
        weekNumber: now.weekNumber,
        weekday: FRIDAY,
        hour: START_TIME
    });
    const _delta = friday.diff(now, ['days', 'hours', 'minutes', 'seconds']);
    return _delta.toObject(); 
};

const getDayOfWeek = () => {
    return DateTime.now().weekday;
}

setInterval(() => {
    timeToNextFriday.set(getCountdownValue());
    dayOfWeek.set(getDayOfWeek());
}, timerAccuracy);

export const isFriday = () => {
    //return false
    return getDayOfWeek() === FRIDAY;
}

export const FRIDAY = 5;
export const START_TIME = 0;
export const timeToNextFriday = writable(getCountdownValue());
export const dayOfWeek = writable(getDayOfWeek());