// Sidebar topics for Dim Kb (starlight-sidebar-topics): one topic per language, each with its own sidebar.
// Each section autogenerates from its folder; page order inside a section = `sidebar.order` in frontmatter.
import type { StarlightUserConfig } from '@astrojs/starlight/types';

type SidebarItem = NonNullable<StarlightUserConfig['sidebar']>[number];

/** A collapsed section that lists every page in `programming/<dir>`. */
function section(label: string, th: string, dir: string): SidebarItem {
	return { label, translations: { th }, collapsed: true, items: [{ autogenerate: { directory: `programming/${dir}` } }] };
}

function overview(slug: string): SidebarItem {
	return { label: 'Overview', translations: { th: 'ภาพรวม' }, slug };
}

const goItems: SidebarItem[] = [
	overview('programming/go'),
	{
		label: 'Fundamentals',
		translations: { th: 'พื้นฐาน' },
		items: [
			section('Getting Started', 'เริ่มต้น', 'go/intro'),
			section('Basic Syntax', 'Syntax พื้นฐาน', 'go/basics'),
			section('Control Flow', 'Control Flow', 'go/control'),
			section('Functions', 'Functions', 'go/functions'),
			section('Slices & Maps', 'Slices และ Maps', 'go/collections'),
			section('Strings', 'Strings', 'go/strings'),
			section('Structs & Interfaces', 'Structs และ Interfaces', 'go/types'),
			section('Pointers', 'Pointers', 'go/pointers'),
			section('Error Handling', 'Error Handling', 'go/errors'),
			section('Packages & Modules', 'Packages และ Modules', 'go/modules'),
		],
	},
	{
		label: 'Concurrency',
		translations: { th: 'Concurrency' },
		items: [{ autogenerate: { directory: 'programming/go/concurrency' } }],
	},
	{
		label: 'Advanced',
		translations: { th: 'ขั้นสูง' },
		items: [section('Generics', 'Generics', 'go/generics')],
	},
	{
		label: 'Ecosystem',
		translations: { th: 'Ecosystem' },
		items: [
			section('HTTP & APIs', 'HTTP และ API', 'go/web'),
			section('Database', 'Database', 'go/database'),
			section('Architecture', 'Architecture', 'go/architecture'),
			section('Testing', 'Testing', 'go/testing'),
			section('Deploy & Tooling', 'Deploy และ Tooling', 'go/deploy'),
		],
	},
	{
		label: 'Reference',
		translations: { th: 'Reference' },
		items: [{ autogenerate: { directory: 'programming/go/reference' } }],
	},
];

export const programmingTopics = [
	{ label: 'Go', link: '/programming/go/', icon: 'seti:go', items: goItems },
	{ label: 'Rust', link: '/programming/rust/', icon: 'seti:rust', items: [overview('programming/rust')] },
	{ label: 'TypeScript', link: '/programming/typescript/', icon: 'seti:typescript', items: [overview('programming/typescript')] },
];
