// move this to another file
export const mapToRange = (x: number,  fromMin= 0.0, fromMax= 1.0, toMin = 0.0, toMax= 1.0) => {
    return toMin + (x - fromMin) * (toMax - toMin) /(fromMax - fromMin)
}
export const clamp = (x: number, min = 0.0, max = 1.0) => {
    return Math.max( min, Math.min(x, max));
}

	// stolen from https://svelte.dev/repl/a1c9e4be23d746ddba6b7d62d4aad875?version=3.35.0
export const slugify = (str = '') => str.toLowerCase().replace(/ /g, '-').replace(/\./g, '');

export const optionToString = (option: string) => {
    const options: {[key: string]: string} = {
        "banner": "Banner", 
        "window_foil": "Fensterfolie",
        "button": "Button",
        "sticker": "Sticker",
        "flag": "Fähnchen"
    }
    if (option == undefined || option == null || option == "") {
        return ""
    } 
    if (!Object.keys(options).includes(option)) {
        return option
    }
    return options[option] 
}