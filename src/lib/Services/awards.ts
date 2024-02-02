export type Award = {title: string, reason: string, achieved: number}
export const streakAwards: {[key: number]: Award} = {
	2: {title: "2er Serie", reason: "zwei aufeinander folgende Teilnahmen", achieved: 2},
	5: {title: "5er Serie", reason: "fünf aufeinander folgende Teilnahmen", achieved: 5},
	10: {title: "10er Serie", reason: "10 aufeinander folgende Teilnahmen", achieved: 10},
	20: {title: "20er Serie", reason: "20 aufeinander folgende Teilnahmen", achieved: 20},
	50: {title: "50er Serie", reason: "50 aufeinander folgende Teilnahmen", achieved: 50},
	100: {title: "100er Serie", reason: "100 aufeinander folgende Teilnahmen", achieved: 100},
}

export const participationAwards: {[key: number]: Award} = {
	2: {title: "Frischling", reason: "insgesamt zwei Teilnahmen", achieved: 2},
	5: {title: "Talent", reason: "insgesamt fünf Teilnahmen", achieved: 5},
	10: {title: "Lehrling", reason: "insgesamt 10 Teilnahmen", achieved: 10},
	20: {title: "Geselle", reason: "insgesamt 20 Teilnahmen", achieved: 20},
	50: {title: "Meister", reason: "insgesamt 50 Teilnahmen", achieved: 50},
}

export const getNewAward = (n: number, type: "p" | "s"): Award => {
	const awardTable = type === "p" ? streakAwards : participationAwards;
	return awardTable[n]
}