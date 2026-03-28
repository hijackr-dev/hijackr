export interface Product {
	slug: string;
	name: string;
	domain: string;
	category: string;
	strapline: string;
	description: string;
	audience: string[];
	features: { title: string; description: string }[];
	competitors: string[];
	wave: 1 | 2 | 3 | 4;
	status: 'live' | 'beta' | 'coming-soon';
	github: string;
}

export const products: Product[] = [
	{
		slug: 'offloadr',
		name: 'Offloadr',
		domain: 'offloadr.io',
		category: 'Automatic Archive',
		strapline: 'Your media, safely offloaded.',
		description:
			'Offloadr is a professional media offload tool for on-set data wrangling — verified, checksum-backed copies to multiple destinations simultaneously.',
		audience: ['DITs', 'Assistant Editors', 'Archivists', 'Post-Production Facilities'],
		features: [
			{
				title: 'Watch-Folder Automation',
				description: 'Drag files into a folder to begin a verified archive process automatically.'
			},
			{
				title: 'Cloud & LTO Integration',
				description: 'Archive to any destination: AWS, Backblaze, Google Cloud, and LTO tape.'
			},
			{
				title: 'Checksum Verification',
				description: 'Guarantees every file is a perfect copy with cryptographic hash checks.'
			}
		],
		competitors: ['Hedge / Offshoot', 'Pomfort Silverstack', 'ShotPut Pro'],
		wave: 2,
		status: 'beta',
		github: 'https://github.com/hijackr-dev/offloadr'
	},
	{
		slug: 'scrollr',
		name: 'Scrollr',
		domain: 'scrollr.io',
		category: 'Film End Roller Titles',
		strapline: 'The final word in professional credits.',
		description:
			'Scrollr automates the creation of end credit rollers by importing credit lists from spreadsheets and rendering high-resolution video files with legally compliant templates.',
		audience: ['Editors', 'Online Editors', 'Finishing Artists', 'Title Designers'],
		features: [
			{
				title: 'Spreadsheet Import',
				description: 'Import credit lists directly from Google Sheets or Excel (CSV).'
			},
			{
				title: 'Legally Compliant Templates',
				description: 'Pre-built templates for guild and union requirements.'
			},
			{
				title: 'High-Resolution Export',
				description: 'Renders up to 8K with alpha channels for perfect compositing.'
			}
		],
		competitors: ['Endcrawl', 'Avid/Premiere/FCP Title Tools', 'In-house Graphics Departments'],
		wave: 1,
		status: 'beta',
		github: 'https://github.com/hijackr-dev/scrollr'
	},
	{
		slug: 'provr',
		name: 'Provr',
		domain: 'provr.media',
		category: 'Open Standard',
		strapline: 'Cryptographic media provenance for the next century.',
		description:
			'Provr is an open standard for media integrity — replacing ASC-MHL v2 with a modern, cloud-native, cryptographically rigorous framework built for multi-petabyte workflows.',
		audience: ['DITs', 'Post-Production Facilities', 'Studios', 'Camera Manufacturers'],
		features: [
			{
				title: 'BLAKE3 + Merkle DAG',
				description: 'Parallelised cryptographic hashing with O(log N) verification.'
			},
			{
				title: 'Zero-Egress Cloud Verification',
				description: 'Verify petabytes of S3 data server-side with no download cost.'
			},
			{
				title: 'Post-Quantum Signatures',
				description: 'Hybrid Ed25519 + SLH-DSA (FIPS 205) for century-scale archival security.'
			}
		],
		competitors: ['ASC-MHL v2'],
		wave: 2,
		status: 'coming-soon',
		github: 'https://github.com/hijackr-dev/provr'
	},
	{
		slug: 'organisr',
		name: 'Organisr',
		domain: 'organisr.io',
		category: 'Project Organisation',
		strapline: 'Start every project perfectly organised.',
		description:
			'Organisr automates the creation of standardised folder structures for new projects. Create unlimited custom templates and spin up a perfectly organised project directory in seconds.',
		audience: [
			'Production Coordinators',
			'Post-Production Supervisors',
			'Assistant Editors',
			'Freelance Creatives'
		],
		features: [
			{
				title: 'Template Library',
				description: 'Create and customise unlimited folder structure templates.'
			},
			{
				title: 'One-Click Projects',
				description: 'Spin up a complete, organised project directory instantly.'
			},
			{
				title: 'Hijackr Cloud Integration',
				description: 'Create new projects directly from the Hijackr dashboard.'
			}
		],
		competitors: ['Post Haste'],
		wave: 2,
		status: 'coming-soon',
		github: 'https://github.com/hijackr-dev/organisr'
	},
	{
		slug: 'catalogr',
		name: 'Catalogr',
		domain: 'catalogr.io',
		category: 'Disk Cataloguing',
		strapline: 'Your entire archive, instantly searchable.',
		description:
			'Catalogr creates a searchable, offline index of all files on any connected drive and generates high-quality video thumbnails for easy identification.',
		audience: [
			'DITs',
			'Assistant Editors',
			'Archivists',
			'Post-Production Facilities'
		],
		features: [
			{
				title: 'Offline Drive Index',
				description: 'Catalog any drive type (HDD, SSD, LTO, NAS) without it being connected.'
			},
			{
				title: 'Video Thumbnails',
				description: 'High-quality thumbnails for instant visual identification.'
			},
			{
				title: 'Offloadr Integration',
				description: 'Automatically catalog drives after a backup is complete.'
			}
		],
		competitors: ['DiskCatalogMaker', 'NeoFinder', 'WinCatalog'],
		wave: 2,
		status: 'coming-soon',
		github: 'https://github.com/hijackr-dev/catalogr'
	},
	{
		slug: 'forwardr',
		name: 'Forwardr',
		domain: 'forwardr.io',
		category: 'File Sharing',
		strapline: 'Deliver anything, anywhere. Securely.',
		description:
			'Forwardr is a professional-grade file transfer solution built for media production — no file size limits, end-to-end encryption, and custom-branded download pages.',
		audience: ['Production Teams', 'DITs', 'Editors', 'VFX Artists', 'Sound Mixers'],
		features: [
			{
				title: 'No File Size Limits',
				description: 'Built to handle massive camera original files and exports.'
			},
			{
				title: 'End-to-End Encryption',
				description: 'Password protection and secure links as standard.'
			},
			{
				title: 'Custom Branded Pages',
				description: 'Present deliveries professionally with your own logo and branding.'
			}
		],
		competitors: ['MASV', 'WeTransfer', 'IBM Aspera'],
		wave: 2,
		status: 'coming-soon',
		github: 'https://github.com/hijackr-dev/forwardr'
	},
	{
		slug: 'replayr',
		name: 'Replayr',
		domain: 'replayr.io',
		category: 'Team Video Review',
		strapline: 'Creative review, streamlined.',
		description:
			'Replayr is a video review and approval platform with frame-accurate, time-stamped comments, on-screen drawing tools, and side-by-side version comparison.',
		audience: ['Directors', 'Producers', 'Agency Creatives', 'Clients', 'Editors'],
		features: [
			{
				title: 'Frame-Accurate Comments',
				description: 'Time-stamped comments and on-screen drawing for clear feedback.'
			},
			{
				title: 'Version Comparison',
				description: 'Compare different cuts side-by-side with version stacking.'
			},
			{
				title: 'Secure Review Links',
				description: 'Share with password protection and expiry dates.'
			}
		],
		competitors: ['Frame.io (Adobe)', 'PIX (X2X Creative)', 'Vimeo Review Tools'],
		wave: 3,
		status: 'coming-soon',
		github: 'https://github.com/hijackr-dev/replayr'
	},
	{
		slug: 'transcodr',
		name: 'Transcodr',
		domain: 'transcodr.io',
		category: 'Transcoding Engine',
		strapline: 'Any format, any time.',
		description:
			'Transcodr is a GPU-accelerated transcoding engine with full support for professional codecs and an API for deep workflow integration.',
		audience: ['Assistant Editors', 'DITs', 'Media Managers', 'Encoding Technicians'],
		features: [
			{
				title: 'GPU-Accelerated Engine',
				description: 'Maximum speed for creating dailies and proxy files.'
			},
			{
				title: 'Pro Codec Support',
				description: 'Full support for ProRes, DNxHD/HR, H.264/H.265, and more.'
			},
			{
				title: 'Workflow API',
				description: 'Plug the engine into larger media pipelines programmatically.'
			}
		],
		competitors: ['Adobe Media Encoder', 'DaVinci Resolve', 'Telestream Vantage'],
		wave: 4,
		status: 'coming-soon',
		github: 'https://github.com/hijackr-dev/transcodr'
	},
	{
		slug: 'examinr',
		name: 'Examinr',
		domain: 'examinr.io',
		category: 'Automatic Quality Control',
		strapline: 'Every pixel, perfected. Automatically.',
		description:
			'Examinr is an automated quality control tool that scans media files for technical errors and generates detailed, time-stamped PDF reports.',
		audience: ['QC Operators', 'Post-Production Facilities', 'Broadcasters', 'VFX Teams'],
		features: [
			{
				title: 'Automated Error Detection',
				description: 'Scans for dead pixels, macroblocking, freeze frames, and more.'
			},
			{
				title: 'Audio & Loudness Checks',
				description: 'Verifies LUFS loudness, phase, and detects audio dropouts.'
			},
			{
				title: 'PDF Report Generation',
				description: 'Detailed, time-stamped reports with thumbnails for all flagged issues.'
			}
		],
		competitors: ['Interra Systems Baton', 'Telestream Vidchecker', 'QScan'],
		wave: 3,
		status: 'coming-soon',
		github: 'https://github.com/hijackr-dev/examinr'
	},
	{
		slug: 'shuttlr',
		name: 'Shuttlr',
		domain: 'shuttlr.io',
		category: 'Storage Collaboration',
		strapline: "Your team's assets, perfectly in sync.",
		description:
			'Shuttlr solves remote editing by mounting cloud storage as a local drive — edit directly from the cloud with minimal latency and real-time file locking.',
		audience: ['Post-Production Supervisors', 'Remote Editors', 'Broadcasters'],
		features: [
			{
				title: 'Mount Cloud as Local Drive',
				description: 'Edit directly from the cloud with no latency.'
			},
			{
				title: 'Real-Time File Locking',
				description: 'Prevents conflicts when multiple users work on the same project.'
			},
			{
				title: 'Granular Permissions',
				description: 'Control who can view, download, or edit specific folders and files.'
			}
		],
		competitors: ['LucidLink', 'Dropbox / Google Drive', 'Hedge Postlab'],
		wave: 4,
		status: 'coming-soon',
		github: 'https://github.com/hijackr-dev/shuttlr'
	},
	{
		slug: 'versionr',
		name: 'Versionr',
		domain: 'versionr.io',
		category: 'Deliverables Creation',
		strapline: 'Master every deliverable.',
		description:
			'Versionr automates the creation of complex, standardised deliverables from a single master file — broadcast specs, social formats, DCPs, and more.',
		audience: ['Assistant Editors', 'Finishing Artists', 'Post Producers', 'Social Media Managers'],
		features: [
			{
				title: 'Template-Based Creation',
				description: 'Build templates for any deliverable spec (Instagram, Broadcast HD, DCP).'
			},
			{
				title: 'Automatic Slates & Watermarks',
				description: 'Apply spec-compliant slates, clocks, and watermarks automatically.'
			},
			{
				title: 'Batch Processing',
				description: 'Export dozens of different versions from a single master file simultaneously.'
			}
		],
		competitors: ['Adobe Media Encoder', 'DaVinci Resolve Render Page', 'Apple Compressor'],
		wave: 4,
		status: 'coming-soon',
		github: 'https://github.com/hijackr-dev/versionr'
	},
	{
		slug: 'crewr',
		name: 'Crewr',
		domain: 'crewr.io',
		category: 'Team Communication',
		strapline: 'The communication hub for production.',
		description:
			'Crewr is a real-time messaging platform built specifically for film and television production, with deep integration into the Hijackr suite.',
		audience: ['Producers', 'Coordinators', 'Editors', 'VFX Artists', 'Sound Mixers'],
		features: [
			{
				title: 'Project-Based Channels',
				description: 'Persistent, organised channels keep conversations structured.'
			},
			{
				title: 'Direct Messaging',
				description: 'One-on-one communication across the crew.'
			},
			{
				title: 'Hijackr Integration',
				description: 'Automated notifications from Replayr, Forwardr, Offloadr, and more.'
			}
		],
		competitors: ['Slack', 'Microsoft Teams'],
		wave: 3,
		status: 'coming-soon',
		github: 'https://github.com/hijackr-dev/crewr'
	},
	{
		slug: 'resourcr',
		name: 'Resourcr',
		domain: 'resourcr.io',
		category: 'Resource Scheduling',
		strapline: 'Your production, scheduled.',
		description:
			'Resourcr is a cloud-based scheduling and resource management platform built specifically for the media and entertainment industry.',
		audience: [
			'Production Managers',
			'Post-Production Supervisors',
			'Resource Managers',
			'Coordinators'
		],
		features: [
			{
				title: 'M&E Scheduling',
				description: 'Book human resources, equipment, and facilities in one place.'
			},
			{
				title: 'Centralised Resource Pool',
				description: 'Manage all resources from a single cloud-based platform.'
			},
			{
				title: 'Budget Tracking',
				description: 'Monitor utilisation, track timesheets, and manage project budgets.'
			}
		],
		competitors: ['Resource Guru', 'Asana', 'Monday.com'],
		wave: 3,
		status: 'coming-soon',
		github: 'https://github.com/hijackr-dev/resourcr'
	},
	{
		slug: 'renamr',
		name: 'Renamr',
		domain: 'renamr.io',
		category: 'File Renaming',
		strapline: 'The right name, every time.',
		description:
			'Renamr is a powerful batch renaming utility that lets you build complex, customisable renaming templates based on file metadata.',
		audience: ['DITs', 'Assistant Editors', 'VFX Artists', 'Photographers'],
		features: [
			{
				title: 'Metadata-Driven Templates',
				description: 'Build renaming rules from EXIF, timecode, camera data, and custom fields.'
			},
			{
				title: 'Live Preview',
				description: 'See exactly what files will be renamed before committing.'
			},
			{
				title: 'Batch Processing',
				description: 'Rename thousands of files in seconds.'
			}
		],
		competitors: ['A Better Finder Rename', 'Adobe Bridge', 'Bulk Rename Utility'],
		wave: 2,
		status: 'coming-soon',
		github: 'https://github.com/hijackr-dev/renamr'
	},
	{
		slug: 'conformr',
		name: 'Conformr',
		domain: 'conformr.io',
		category: 'Online Conform',
		strapline: 'Collect everything. Miss nothing.',
		description:
			'Conformr automates the online conform stage of post-production — ingest an EDL, AAF, or XML and collect all the corresponding high-resolution source files.',
		audience: ['Assistant Editors', 'Online Editors'],
		features: [
			{
				title: 'EDL/AAF/XML Ingestion',
				description: 'Parse edit decision lists from any major NLE.'
			},
			{
				title: 'Automatic File Collection',
				description: 'Locate and consolidate all referenced high-resolution source files.'
			},
			{
				title: 'Discrepancy Detection',
				description: 'AI-assisted cross-referencing flags metadata mismatches before they cause problems.'
			}
		],
		competitors: ['DaVinci Resolve conform tools'],
		wave: 3,
		status: 'coming-soon',
		github: 'https://github.com/hijackr-dev/conformr'
	},
	{
		slug: 'verifyr',
		name: 'Verifyr',
		domain: 'verifyr.io',
		category: 'Hash Verification',
		strapline: 'Trust, but verify.',
		description:
			'Verifyr is a lightweight verification tool for checking the integrity of media manifests — supporting both Provr (.provr) and legacy ASC-MHL v2 (.mhl) formats.',
		audience: ['DITs', 'Assistant Editors', 'Post-Production Facilities'],
		features: [
			{
				title: 'Provr Support',
				description: 'Verify .provr manifests against the Provr Media Manifest Specification v1.0.'
			},
			{
				title: 'ASC-MHL Support',
				description: 'Verify legacy .mhl chain files for backward compatibility.'
			},
			{
				title: 'Drag-and-Drop',
				description: 'Drop a manifest file to instantly verify the accompanying media.'
			}
		],
		competitors: ['ASC-MHL tools'],
		wave: 2,
		status: 'coming-soon',
		github: 'https://github.com/hijackr-dev/verifyr'
	},
	{
		slug: 'reformattr',
		name: 'Reformattr',
		domain: 'reformattr.io',
		category: 'File Repair',
		strapline: 'Find the flaw. Fix the file.',
		description:
			'Reformattr is a surgical file repair tool that detects bit-flips and silent corruption in media files, and repairs only the damaged bytes.',
		audience: ['DITs', 'Archivists', 'Post-Production Facilities'],
		features: [
			{
				title: 'Bit-Flip Detection',
				description: 'Identifies silent corruption at the byte level using cryptographic verification.'
			},
			{
				title: 'Surgical Repair',
				description: 'Requests and patches only the damaged chunk — not the entire file.'
			},
			{
				title: 'Provr Integration',
				description: 'Works natively with Provr manifests to locate and repair corrupted chunks.'
			}
		],
		competitors: [],
		wave: 2,
		status: 'coming-soon',
		github: 'https://github.com/hijackr-dev/reformattr'
	},
	{
		slug: 'relay',
		name: 'Relay',
		domain: 'relay.hijackr.io',
		category: 'Expert Network',
		strapline: 'The answer you need, from the expert who knows.',
		description:
			'Relay is a marketplace for asynchronous micro-consulting — get specific, expert answers to production questions by paying a fee set by the expert.',
		audience: ['Production Professionals', 'M&E Experts'],
		features: [
			{
				title: 'Expert Marketplace',
				description: 'Browse and book recognised industry experts by specialism.'
			},
			{
				title: 'Asynchronous Q&A',
				description: 'Submit questions and receive detailed video or written responses.'
			},
			{
				title: 'M&E Focused',
				description: 'Launched exclusively within the Hijackr community.'
			}
		],
		competitors: ['Cameo', 'Clarity.fm'],
		wave: 3,
		status: 'coming-soon',
		github: 'https://github.com/hijackr-dev/relay'
	},
	{
		slug: 'frameworkr',
		name: 'Frameworkr',
		domain: 'frameworkr.io',
		category: 'Business Procedures',
		strapline: 'The definitive source for operational best practices.',
		description:
			'Frameworkr is a public, version-controlled repository for business documents and procedures — applying the software development workflow to business paperwork.',
		audience: ['Production Companies', 'Post-Houses', 'Agencies'],
		features: [
			{
				title: 'Version-Controlled Documents',
				description: 'Fork, branch, and merge business procedures like code.'
			},
			{
				title: 'Public Template Library',
				description: 'Discover and adopt best-practice templates from the community.'
			},
			{
				title: 'Private Repositories',
				description: 'Teams pay for private repos to manage internal procedures securely.'
			}
		],
		competitors: ['Notion', 'Confluence'],
		wave: 3,
		status: 'coming-soon',
		github: 'https://github.com/hijackr-dev/frameworkr'
	},
	{
		slug: 'chronicler',
		name: 'Chronicler',
		domain: 'chronicler.io',
		category: "Writers' Version Control",
		strapline: 'Every draft, preserved. Every change, understood.',
		description:
			"Chronicler is a version control system designed specifically for complex writing projects — scripts, academic papers, and books. It combines a premium writing tool with the collaborative power of GitHub.",
		audience: ['Screenwriters', 'Novelists', 'Academics', "Writers' Rooms"],
		features: [
			{
				title: 'Prose-Native Diff',
				description: 'See exactly what changed between drafts in human-readable form.'
			},
			{
				title: 'Distraction-Free Writing',
				description: 'A clean, focused writing environment built for long-form work.'
			},
			{
				title: 'Collaborative Branching',
				description: "Writers' rooms can work on parallel versions and merge changes."
			}
		],
		competitors: ['Final Draft', 'Highland 2', 'Scrivener'],
		wave: 4,
		status: 'coming-soon',
		github: 'https://github.com/hijackr-dev/chronicler'
	}
];

export function getProduct(slug: string): Product | undefined {
	return products.find((p) => p.slug === slug);
}

export const waveLabels: Record<number, string> = {
	1: 'Wave 1 — Foundation',
	2: 'Wave 2 — M&E Toolkit',
	3: 'Wave 3 — Collaboration',
	4: 'Wave 4 — Enterprise'
};