/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			animation: {
				'modal-show': 'slideup 1s ease-out'
			},
			keyframes: {
				slideup: {
					'0%': { transform: 'translate(0,100%) scale(1.8)' },
					'100%': { transform: 'translate(0,0) scale(1.0)' }
				}
			},
			colors: {
				'primary-900': '#081928',
				'primary-800': '#0F3457',
				'primary-700': '#10518D',
				'primary-500': '#095FAE',
				'primary-300': '#70A8DC',
				'primary-200': '#AECDEB',
				'primary-100': '#D4E3F0',
				'primary-50': '#081928',
				'secondary-900': '#08281F',
				'secondary-800': '#0B4837',
				'secondary-700': '#4A7E70',
				'secondary-500': '#4EAF94',
				'secondary-300': '#8DE1C9',
				'secondary-200': '#AAF4DF',
				'secondary-100': '#D4E3F0',
				'secondary-50': '#081928',
				'neutral-900': '#252627',
				'neutral-800': '#434343',
				'neutral-700': '#A4A8AC',
				'neutral-500': '#B4BBC2',
				'neutral-300': '#DADEE1',
				'neutral-200': '#E9ECEF',
				'neutral-100': '#F0F3F5',
				'neutral-50': '#F8F9F9',
				'accent-900': '#600A1E',
				'accent-800': '#840725',
				'accent-700': '#B20730',
				'accent-500': '#D81F4C',
				'accent-300': '#FF88A5',
				'accent-200': '#FFD0DB',
				'accent-100': '#FDEFF2',
				'success-900': '#145239',
				'success-800': '#197741',
				'success-700': '#259D58',
				'success-500': '#38C172',
				'success-300': '#74D99F',
				'success-200': '#A8EEC1',
				'success-100': '#E3FCEC',
				'warning-900': '#5C4813',
				'warning-800': '#8C6D1F',
				'warning-700': '#CAA53D',
				'warning-500': '#F4CA64',
				'warning-300': '#FAE29F',
				'warning-200': '#FDF3D7',
				'warning-100': '#FFFCF4',
				'error-900': '#601717',
				'error-800': '#891B1B',
				'error-700': '#B82020',
				'error-500': '#DC3030',
				'error-300': '#E36363',
				'error-200': '#F5AAAA',
				'error-100': '#FCE8E8'
			},
			fontFamily: {
				sans: ['Inter', 'sans-serif'],
				serif: ['Poppins', 'serif']
			}
		}
	},
	safelist: [
		'bg-success-100',
		'bg-success-200',
		'bg-success-300',
		'bg-success-500',
		'bg-success-700',
		'bg-success-800',
		'bg-success-900',
		'border-success-100',
		'border-success-200',
		'border-success-300',
		'border-success-500',
		'border-success-700',
		'border-success-800',
		'border-success-900',
		'text-success-100',
		'text-success-200',
		'text-success-300',
		'text-success-500',
		'text-success-700',
		'text-success-800',
		'text-success-900',
	],
	plugins: [require('daisyui'), require('@tailwindcss/typography')],
	daisyui: {
		themes: ['light'], // true: all themes | false: only light + dark | array: specific themes like this ["light", "dark", "cupcake"]
		//darkTheme: "dark", // name of one of the included themes for dark mode
		base: false, // applies background color and foreground color for root element by default
		styled: true, // include daisyUI colors and design decisions for all components
		utils: true, // adds responsive and modifier utility classes
		rtl: false, // rotate style direction from left-to-right to right-to-left. You also need to add dir="rtl" to your html tag and install `tailwindcss-flip` plugin for Tailwind CSS.
		prefix: '', // prefix for daisyUI classnames (components, modifiers and responsive class names. Not colors)
		logs: false // Shows info about daisyUI version and used config in the console when building your CSS
	}
};
