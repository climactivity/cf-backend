export interface Participation {
  collectionId: string;
  collectionName: string;
  created: string;
  date: string;
  id: string;
  note: string;
  option: string;
  owner: string;
  updated: string;
  weeknumber: number;
}

export interface Week {
  FridayReminderSent: boolean;
  SaturdayReminderSent: boolean;
  collectionId: string;
  collectionName: string;
  completion: Participation;
  created: string;
  id: string;
  owner: string;
  state: string;
  updated: string;
  weeknumber: number;
  year: number;
}

export interface UserRecord {
  avatar: string;
  collectionId: string;
  collectionName: string;
  created: string;
  email: string;
  emailVisibility: boolean;
  id: string;
  name: string;
  region: string;
  updated: string;
  username: string;
  verified: boolean;
  notification_setup: boolean;
  notification_push: boolean;
  notification_email: boolean;
}