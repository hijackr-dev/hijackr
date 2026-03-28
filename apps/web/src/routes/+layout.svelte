<script lang="ts">
	import '../app.css';
	import { page } from '$app/stores';

	const nav = [
		{ label: 'Products', href: '/#products' },
		{ label: 'Provr', href: '/provr' },
		{ label: 'About', href: '/about' }
	];
</script>

<svelte:head>
	<meta name="description" content="Hijackr — an integrated post-production operating system." />
</svelte:head>

<!-- Navigation -->
<header class="fixed top-0 left-0 right-0 z-50 border-b border-surface-700/50 bg-surface-900/80 backdrop-blur-md">
	<div class="mx-auto flex max-w-7xl items-center justify-between px-6 py-4">
		<a href="/" class="flex items-center gap-2">
			<span class="text-xl font-bold tracking-tight text-white">Hijackr</span>
		</a>
		<nav class="hidden items-center gap-8 md:flex">
			{#each nav as item}
				<a
					href={item.href}
					class="text-sm font-medium text-surface-300 transition-colors hover:text-white"
					class:text-white={$page.url.pathname === item.href}
				>
					{item.label}
				</a>
			{/each}
		</nav>
		<div class="flex items-center gap-3">
			<a
				href="https://app.hijackr.io"
				class="rounded-lg bg-primary-500 px-4 py-2 text-sm font-semibold text-white transition-colors hover:bg-primary-400"
			>
				Sign in
			</a>
		</div>
	</div>
</header>

<!-- Page content -->
<main class="min-h-screen bg-surface-900 pt-16 text-white">
	<slot />
</main>

<!-- Footer -->
<footer class="border-t border-surface-700/50 bg-surface-900 px-6 py-12">
	<div class="mx-auto max-w-7xl">
		<div class="flex flex-col items-start justify-between gap-8 md:flex-row">
			<div>
				<p class="text-lg font-bold text-white">Hijackr</p>
				<p class="mt-1 text-sm text-surface-400">Hijack your workflow. Own the result.</p>
				<p class="mt-4 text-xs text-surface-500">© {new Date().getFullYear()} Hijackr Ltd. All rights reserved.</p>
			</div>
			<div class="grid grid-cols-2 gap-8 text-sm md:grid-cols-3">
				<div>
					<p class="mb-3 font-semibold text-surface-300">Products</p>
					<ul class="space-y-2 text-surface-400">
						<li><a href="/offloadr" class="hover:text-white">Offloadr</a></li>
						<li><a href="/scrollr" class="hover:text-white">Scrollr</a></li>
						<li><a href="/provr" class="hover:text-white">Provr</a></li>
						<li><a href="/#products" class="hover:text-white">All products →</a></li>
					</ul>
				</div>
				<div>
					<p class="mb-3 font-semibold text-surface-300">Company</p>
					<ul class="space-y-2 text-surface-400">
						<li><a href="/about" class="hover:text-white">About</a></li>
						<li><a href="https://github.com/hijackr-dev" class="hover:text-white">GitHub</a></li>
					</ul>
				</div>
				<div>
					<p class="mb-3 font-semibold text-surface-300">Legal</p>
					<ul class="space-y-2 text-surface-400">
						<li><a href="/privacy" class="hover:text-white">Privacy</a></li>
						<li><a href="/terms" class="hover:text-white">Terms</a></li>
					</ul>
				</div>
			</div>
		</div>
	</div>
</footer>