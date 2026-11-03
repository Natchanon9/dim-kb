import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';
import { goSidebar } from './src/sidebar';

// GitHub Pages: https://natchanon9.github.io/dim-kb/
export default defineConfig({
	site: 'https://natchanon9.github.io',
	base: '/dim-kb',
	integrations: [
		starlight({
			title: 'dim-kb',
			description: 'Personal knowledge base — Go, Rust, DevOps',
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
			customCss: ['./src/styles/custom.css'],
			head: [
				{ tag: 'link', attrs: { rel: 'preconnect', href: 'https://fonts.googleapis.com' } },
				{ tag: 'link', attrs: { rel: 'preconnect', href: 'https://fonts.gstatic.com', crossorigin: '' } },
				{
					tag: 'link',
					attrs: {
						rel: 'stylesheet',
						href: 'https://fonts.googleapis.com/css2?family=IBM+Plex+Sans+Thai:wght@400;500;600;700&display=swap',
					},
				},
			],
			sidebar: [
				{
					label: 'Go',
					items: [{ label: 'Overview', translations: { th: 'ภาพรวม' }, slug: 'go' }, ...goSidebar],
				},
				{ label: 'Rust', collapsed: true, items: [{ label: 'Overview', translations: { th: 'ภาพรวม' }, slug: 'rust' }] },
				{ label: 'DevOps', collapsed: true, items: [{ label: 'Overview', translations: { th: 'ภาพรวม' }, slug: 'devops' }] },
			],
		}),
	],
});
