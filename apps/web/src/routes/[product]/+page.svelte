<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
	$: product = data.product;

	const statusLabel: Record<string, string> = {
		live: 'Live',
		beta: 'Beta',
		'coming-soon': 'Coming Soon'
	};

	const statusClass: Record<string, string> = {
		live: 'bg-success-500/20 text-success-300 border border-success-500/30',
		beta: 'bg-warning-500/20 text-warning-300 border border-warning-500/30',
		'coming-soon': 'bg-surface-700/50 text-surface-400 border border-surface-600'
	};

	const waveLabel: Record<number, string> = {
		1: 'Wave 1 — Foundation',
		2: 'Wave 2 — M&E Toolkit',
		3: 'Wave 3 — Collaboration',
		4: 'Wave 4 — Enterprise'
	};
</script>

<svelte:head>
	<title>{product.name} — {product.strapline} | Hijackr</title>
	<meta name="description" content={product.description} />
</svelte:head>

<!-- Hero -->
<section class="relative overflow-hidden px-6 py-28">
	<div class="absolute inset-0 bg-gradient-to-b from-primary-900/10 to-transparent"></div>
	<div class="relative mx-auto max-w-4xl">
		<div class="mb-6 flex flex-wrap items-center gap-3">
			<a href="/" class="text-sm text-surface-500 hover:text-surface-300">Hijackr</a>
			<span class="text-surface-600">/</span>
			<span class="text-sm text-surface-400">{product.name}</span>
			<span class="rounded-full px-3 py-1 text-xs font-medium {statusClass[product.status]}">
				{statusLabel[product.status]}
			</span>
			<span class="rounded-full border border-surface-700 px-3 py-1 text-xs text-surface-500">
				{waveLabel[product.wave]}
			</span>
		</div>

		<p class="mb-3 text-sm font-semibold uppercase tracking-widest text-primary-400">
			{product.category}
		</p>
		<h1 class="mb-4 text-5xl font-extrabold tracking-tight text-white md:text-6xl">
			{product.name}
		</h1>
		<p class="mb-6 text-2xl font-light text-surface-300">{product.strapline}</p>
		<p class="mb-10 max-w-2xl text-lg text-surface-400">{product.description}</p>

		<div class="flex flex-wrap gap-4">
			{#if product.status === 'live' || product.status === 'beta'}
				<a
					href="https://app.hijackr.io/{product.slug}"
					class="rounded-xl bg-primary-500 px-8 py-3 text-base font-semibold text-white transition-colors hover:bg-primary-400"
				>
					{product.status === 'beta' ? 'Join the beta' : 'Get started'}
				</a>
			{:else}
				<a
					href="https://app.hijackr.io"
					class="rounded-xl bg-primary-500 px-8 py-3 text-base font-semibold text-white transition-colors hover:bg-primary-400"
				>
					Join the waitlist
				</a>
			{/if}
			<a
				href={product.github}
				target="_blank"
				rel="noopener noreferrer"
				class="rounded-xl border border-surface-600 px-8 py-3 text-base font-semibold text-surface-300 transition-colors hover:border-surface-400 hover:text-white"
			>
				GitHub →
			</a>
		</div>
	</div>
</section>

<!-- Features -->
<section class="px-6 py-20">
	<div class="mx-auto max-w-4xl">
		<h2 class="mb-12 text-2xl font-bold text-white">Key Features</h2>
		<div class="grid grid-cols-1 gap-6 md:grid-cols-3">
			{#each product.features as feature}
				<div class="rounded-2xl border border-surface-700/50 bg-surface-800/50 p-6">
					<h3 class="mb-2 font-semibold text-white">{feature.title}</h3>
					<p class="text-sm text-surface-400">{feature.description}</p>
				</div>
			{/each}
		</div>
	</div>
</section>

<!-- Audience + Competitors -->
<section class="px-6 py-12">
	<div class="mx-auto grid max-w-4xl grid-cols-1 gap-8 md:grid-cols-2">
		<div class="rounded-2xl border border-surface-700/50 bg-surface-800/50 p-6">
			<h3 class="mb-4 font-semibold text-white">Built for</h3>
			<ul class="space-y-2">
				{#each product.audience as role}
					<li class="flex items-center gap-2 text-sm text-surface-300">
						<span class="h-1.5 w-1.5 rounded-full bg-primary-500"></span>
						{role}
					</li>
				{/each}
			</ul>
		</div>
		{#if product.competitors.length > 0}
			<div class="rounded-2xl border border-surface-700/50 bg-surface-800/50 p-6">
				<h3 class="mb-4 font-semibold text-white">Replaces</h3>
				<ul class="space-y-2">
					{#each product.competitors as competitor}
						<li class="flex items-center gap-2 text-sm text-surface-400">
							<span class="h-1.5 w-1.5 rounded-full bg-surface-600"></span>
							{competitor}
						</li>
					{/each}
				</ul>
			</div>
		{/if}
	</div>
</section>

<!-- Back to all products -->
<section class="px-6 py-16 text-center">
	<a href="/#products" class="text-sm text-surface-500 hover:text-surface-300">
		← Back to all products
	</a>
</section>