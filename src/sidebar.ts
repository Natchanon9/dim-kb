// Sidebar for dim-kb. Each group autogenerates from its folder;
// page order inside a group = `sidebar.order` in frontmatter.
import type { StarlightUserConfig } from '@astrojs/starlight/types';

type SidebarItem = NonNullable<StarlightUserConfig['sidebar']>[number];

export const goSidebar: SidebarItem[] = [
	{
		"label": "Getting Started",
		"translations": {
			"th": "เริ่มต้น"
		},
		"collapsed": true,
		"items": [
			{
				"autogenerate": {
					"directory": "go/intro"
				}
			}
		]
	},
	{
		"label": "Basic Syntax",
		"translations": {
			"th": "Syntax พื้นฐาน"
		},
		"collapsed": true,
		"items": [
			{
				"autogenerate": {
					"directory": "go/basics"
				}
			}
		]
	},
	{
		"label": "Control Flow",
		"translations": {
			"th": "Control Flow"
		},
		"collapsed": true,
		"items": [
			{
				"autogenerate": {
					"directory": "go/control"
				}
			}
		]
	},
	{
		"label": "Functions",
		"translations": {
			"th": "Functions"
		},
		"collapsed": true,
		"items": [
			{
				"autogenerate": {
					"directory": "go/functions"
				}
			}
		]
	},
	{
		"label": "Slices & Maps",
		"translations": {
			"th": "Slices และ Maps"
		},
		"collapsed": true,
		"items": [
			{
				"autogenerate": {
					"directory": "go/collections"
				}
			}
		]
	},
	{
		"label": "Strings",
		"translations": {
			"th": "Strings"
		},
		"collapsed": true,
		"items": [
			{
				"autogenerate": {
					"directory": "go/strings"
				}
			}
		]
	},
	{
		"label": "Structs & Interfaces",
		"translations": {
			"th": "Structs และ Interfaces"
		},
		"collapsed": true,
		"items": [
			{
				"autogenerate": {
					"directory": "go/types"
				}
			}
		]
	},
	{
		"label": "Pointers",
		"translations": {
			"th": "Pointers"
		},
		"collapsed": true,
		"items": [
			{
				"autogenerate": {
					"directory": "go/pointers"
				}
			}
		]
	},
	{
		"label": "Error Handling",
		"translations": {
			"th": "Error Handling"
		},
		"collapsed": true,
		"items": [
			{
				"autogenerate": {
					"directory": "go/errors"
				}
			}
		]
	},
	{
		"label": "Packages & Modules",
		"translations": {
			"th": "Packages และ Modules"
		},
		"collapsed": true,
		"items": [
			{
				"autogenerate": {
					"directory": "go/modules"
				}
			}
		]
	},
	{
		"label": "Concurrency",
		"translations": {
			"th": "Concurrency"
		},
		"collapsed": true,
		"items": [
			{
				"autogenerate": {
					"directory": "go/concurrency"
				}
			}
		]
	},
	{
		"label": "Generics",
		"translations": {
			"th": "Generics"
		},
		"collapsed": true,
		"items": [
			{
				"autogenerate": {
					"directory": "go/generics"
				}
			}
		]
	},
	{
		"label": "HTTP & APIs",
		"translations": {
			"th": "HTTP และ API"
		},
		"collapsed": true,
		"items": [
			{
				"autogenerate": {
					"directory": "go/web"
				}
			}
		]
	},
	{
		"label": "Database",
		"translations": {
			"th": "Database"
		},
		"collapsed": true,
		"items": [
			{
				"autogenerate": {
					"directory": "go/database"
				}
			}
		]
	},
	{
		"label": "Architecture",
		"translations": {
			"th": "Architecture"
		},
		"collapsed": true,
		"items": [
			{
				"autogenerate": {
					"directory": "go/architecture"
				}
			}
		]
	},
	{
		"label": "Testing",
		"translations": {
			"th": "Testing"
		},
		"collapsed": true,
		"items": [
			{
				"autogenerate": {
					"directory": "go/testing"
				}
			}
		]
	},
	{
		"label": "Deploy & Tooling",
		"translations": {
			"th": "Deploy และ Tooling"
		},
		"collapsed": true,
		"items": [
			{
				"autogenerate": {
					"directory": "go/deploy"
				}
			}
		]
	},
	{
		"label": "Reference",
		"translations": {
			"th": "Reference"
		},
		"collapsed": true,
		"items": [
			{
				"autogenerate": {
					"directory": "go/reference"
				}
			}
		]
	}
];
