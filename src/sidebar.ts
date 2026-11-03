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

// Mirrors algomaster's Go course 1:1 (same sections, same order, lessons as pages; no exercises).
const goItems: SidebarItem[] = [
	overview('programming/go'),
	section('Introduction to Go', 'แนะนำ Go', 'go/introduction'),
	section('Basic Syntax', 'Syntax พื้นฐาน', 'go/basic-syntax'),
	section('Control Flow', 'Control Flow', 'go/control-flow'),
	section('Functions', 'Functions', 'go/functions'),
	section('Arrays & Slices', 'Arrays และ Slices', 'go/arrays-slices'),
	section('Maps', 'Maps', 'go/maps'),
	section('Strings', 'Strings', 'go/strings'),
	section('Structs', 'Structs', 'go/structs'),
	section('Methods & Interfaces', 'Methods และ Interfaces', 'go/methods-interfaces'),
	section('Pointers', 'Pointers', 'go/pointers'),
	section('Error Handling', 'Error Handling', 'go/error-handling'),
	section('Packages & Modules', 'Packages และ Modules', 'go/packages-modules'),
	section('Goroutines', 'Goroutines', 'go/goroutines'),
	section('Channels', 'Channels', 'go/channels'),
	section('Concurrency Patterns', 'Concurrency Patterns', 'go/concurrency-patterns'),
	section('Context Package', 'Context Package', 'go/context'),
	section('Generics', 'Generics', 'go/generics'),
	section('File I/O', 'File I/O', 'go/file-io'),
	section('JSON & Encoding', 'JSON และ Encoding', 'go/json-encoding'),
	section('Testing', 'Testing', 'go/testing'),
	section('Networking & HTTP', 'Networking และ HTTP', 'go/networking-http'),
	section('Database Access', 'Database Access', 'go/database-access'),
	section('Reflection', 'Reflection', 'go/reflection'),
	section('Unsafe & cgo', 'Unsafe และ cgo', 'go/unsafe-cgo'),
	section('Standard Library', 'Standard Library', 'go/standard-library'),
	section('Modern Go', 'Modern Go', 'go/modern-go'),
	section('Best Practices', 'Best Practices', 'go/best-practices'),
	{
		// Not in algomaster: backend chapters from mikelopster GoAPI Essential
		label: 'Backend Practice',
		translations: { th: 'ฝึกทำ Backend' },
		items: [
			section('Fiber', 'Fiber', 'go/fiber'),
			section('GORM', 'GORM', 'go/gorm'),
			section('Architecture', 'Architecture', 'go/architecture'),
			section('Deploy', 'Deploy', 'go/deploy'),
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
