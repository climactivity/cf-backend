export type PBError = {code: string, message: string}
export type ValidationErrors = {[key: string]: PBError} | null; 

