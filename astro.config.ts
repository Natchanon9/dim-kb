import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';
import starlightSidebarTopics from 'starlight-sidebar-topics';
import { programmingTopics } from './src/sidebar';

// GitHub Pages: https://natchanon9.github.io/dim-kb/
export default defineConfig({
	site: 'https://natchanon9.github.io',
	base: '/dim-kb',
	integrations: [
		starlight({
			title: 'Dim Kb',
			description: 'Personal knowledge base: Go, Rust, DevOps',
			favicon: '/favicon.svg',
			defaultLocale: 'root',
			locales: {
				// English is the source language (served at /), Thai is a translation (/th/)
				root: { label: 'English', lang: 'en' },
				th: { label: 'ไทย', lang: 'th' },
			},
			social: [{ icon: 'github', label: 'GitHub', href: 'https://github.com/Natchanon9/dim-kb' }],
			editLink: { baseUrl: 'https://github.com/Natchanon9/dim-kb/edit/main/' },
			lastUpdated: true,
			customCss: [
				// Self-hosted fonts (design option C): Inter for Latin, Noto Sans Thai for Thai, JetBrains Mono for code
				'@fontsource-variable/inter',
				'@fontsource-variable/noto-sans-thai',
				'@fontsource-variable/jetbrains-mono',
				'./src/styles/custom.css',
			],
			expressiveCode: {
				styleOverrides: { borderRadius: '0.75rem', codePaddingBlock: '1rem', codePaddingInline: '1.25rem' },
			},
			plugins: [starlightSidebarTopics(programmingTopics)],
		}),
	],
});
