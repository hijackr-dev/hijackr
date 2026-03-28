import type { Config } from 'tailwindcss';

const config: Config = {
	darkMode: 'class',
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			fontFamily: {
				sans: ['Inter', 'system-ui', 'sans-serif'],
				mono: ['JetBrains Mono', 'monospace']
			},
			colors: {
				// Hijackr design tokens — replace with brand colours when ready
				primary: {
					400: '#818cf8',
					500: '#6366f1',
					900: '#1e1b4b'
				},
				surface: {
					300: '#d1d5db',
					400: '#9ca3af',
					500: '#6b7280',
					600: '#4b5563',
					700: '#374151',
					800: '#1f2937',
					900: '#111827'
				}
			}
		}
	},
	plugins: []
};

export default config;