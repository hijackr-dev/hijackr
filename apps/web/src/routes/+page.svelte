<script lang="ts">
	import { products } from '$lib/products';

	const statusLabel: Record<string, string> = {
		live: 'Live',
		beta: 'Beta',
		'coming-soon': 'Coming Soon'
	};

	const statusClass: Record<string, string> = {
		live: 'bg-success-500/20 text-success-400',
		beta: 'bg-warning-500/20 text-warning-400',
		'coming-soon': 'bg-surface-700 text-surface-400'
	};
</script>

<svelte:head>
	<title>Hijackr — Hijack your workflow. Own the result.</title>
	<meta
		name="description"
		content="Hijackr is an integrated post-production operating system — a unified suite of professional tools for the media and entertainment industry."
	/>
</svelte:head>

<!-- Hero -->
<section class="relative overflow-hidden px-6 py-32 text-center">
	<div class="absolute inset-0 bg-gradient-to-b from-primary-900/20 to-transparent"></div>
	<div class="relative mx-auto max-w-4xl">
		<p class="mb-4 text-sm font-semibold uppercase tracking-widest text-primary-400">
			The integrated post-production operating system
		</p>
		<h1 class="mb-6 text-5xl font-extrabold leading-tight tracking-tight text-white md:text-7xl">
			Hijack your workflow.<br />Own the result.
		</h1>
		<p class="mx-auto mb-10 max-w-2xl text-lg text-surface-300">
			A unified suite of professional tools for the media and entertainment industry — built for
			on-set, post-production, and the cloud.
		</p>
		<div class="flex flex-col items-center justify-center gap-4 sm:flex-row">
			<a
				href="https://app.hijackr.io"
				class="rounded-xl bg-primary-500 px-8 py-3 text-base font-semibold text-white transition-colors hover:bg-primary-400"
			>
				Get started free
			</a>
			<a
				href="#products"
				class="rounded-xl border border-surface-600 px-8 py-3 text-base font-semibold text-surface-300 transition-colors hover:border-surface-400 hover:text-white"
			>
				Explore products
			</a>
		</div>
	</div>
</section>

<!-- Products grid -->
<section id="products" class="px-6 py-24">
	<div class="mx-auto max-w-7xl">
		<div class="mb-12 text-center">
			<h2 class="text-3xl font-bold text-white">The Hijackr Ecosystem</h2>
			<p class="mt-3 text-surface-400">
				{products.length} tools. One platform. Every stage of production.
			</p>
		</div>
		<div class="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
			{#each products as product}
				<a
					href="/{product.slug}"
					class="group flex flex-col rounded-2xl border border-surface-700/50 bg-surface-800/50 p-6 transition-all hover:border-primary-500/50 hover:bg-surface-800"
				>
					<div class="mb-4 flex items-start justify-between">
						<span class="text-xs font-medium text-surface-500">{product.category}</span>
						<span
							class="rounded-full px-2 py-0.5 text-xs font-medium {statusClass[product.status]}"
						>
							{statusLabel[product.status]}
						</span>
					</div>
					<h3 class="mb-1 text-lg font-bold text-white group-hover:text-primary-400 transition-colors">
						{product.name}
					</h3>
					<p class="mb-4 text-sm text-surface-400">{product.strapline}</p>
					<div class="mt-auto">
						<span class="text-xs font-medium text-primary-500 group-hover:text-primary-400">
							Learn more →
						</span>
					</div>
				</a>
			{/each}
		</div>
	</div>
</section>

<!-- CTA -->
<section class="px-6 py-24">
	<div class="mx-auto max-w-3xl rounded-3xl border border-primary-500/20 bg-primary-900/20 p-12 text-center">
		<h2 class="mb-4 text-3xl font-bold text-white">Ready to hijack your workflow?</h2>
		<p class="mb-8 text-surface-300">
			Join the Hijackr community and get early access to the tools that are redefining
			post-production.
		</p>
		<a
			href="https://app.hijackr.io"
			class="inline-block rounded-xl bg-primary-500 px-8 py-3 text-base font-semibold text-white transition-colors hover:bg-primary-400"
		>
			Create your account
		</a>
	</div>
</section>