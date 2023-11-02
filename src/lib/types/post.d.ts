/* eslint-disable @typescript-eslint/no-explicit-any */
export interface Post {
	id: number;
	date: string;
	date_gmt: string;
	guid: Guid;
	modified: string;
	modified_gmt: string;
	slug: string;
	status: string;
	type: string;
	link: string;
	title: Guid;
	content: Content;
	excerpt: Content;
	author: number;
	featured_media: number;
	comment_status: string;
	ping_status: string;
	sticky: boolean;
	template: string;
	format: string;
	meta: Meta;
	categories: number[];
	tags: any[];
	_links: Links;
	_embedded: any;
}

interface Links {
	self: Self[];
	collection: Self[];
	about: Self[];
	author: Author[];
	replies: Author[];
	'version-history': Versionhistory[];
	'predecessor-version': Predecessorversion[];
	'wp:attachment': Self[];
	'wp:term': Wpterm[];
	curies: Cury[];
}

interface Cury {
	name: string;
	href: string;
	templated: boolean;
}

interface Wpterm {
	taxonomy: string;
	embeddable: boolean;
	href: string;
}

interface Predecessorversion {
	id: number;
	href: string;
}

interface Versionhistory {
	count: number;
	href: string;
}

interface Author {
	embeddable: boolean;
	href: string;
}

interface Self {
	href: string;
}

interface Meta {
	_EventAllDay: boolean;
	_EventTimezone: string;
	_EventStartDate: string;
	_EventEndDate: string;
	_EventStartDateUTC: string;
	_EventEndDateUTC: string;
	_EventShowMap: boolean;
	_EventShowMapLink: boolean;
	_EventURL: string;
	_EventCost: string;
	_EventCostDescription: string;
	_EventCurrencySymbol: string;
	_EventCurrencyCode: string;
	_EventCurrencyPosition: string;
	_EventDateTimeSeparator: string;
	_EventTimeRangeSeparator: string;
	_EventOrganizerID: any[];
	_EventVenueID: number;
	_OrganizerEmail: string;
	_OrganizerPhone: string;
	_OrganizerWebsite: string;
	_VenueAddress: string;
	_VenueCity: string;
	_VenueCountry: string;
	_VenueProvince: string;
	_VenueZip: string;
	_VenuePhone: string;
	_VenueURL: string;
	_VenueStateProvince: string;
	_VenueLat: string;
	_VenueLng: string;
}

interface Content {
	rendered: string;
	protected: boolean;
}

interface Guid {
	rendered: string;
}
