<script lang="ts">
	import '../app.css';
	import { page } from '$app/stores';

	const isLoginPage = $derived($page.url.pathname === '/');
</script>

{#if isLoginPage}
	<slot />
{:else}
	<div class="shell">
		<nav class="sidebar">
			<div class="logo">
				<span class="logo-mark">H</span>
				<span class="logo-text">Hijackr</span>
			</div>
			<ul class="nav-links">
				<li><a href="/dashboard" class:active={$page.url.pathname === '/dashboard'}>Dashboard</a></li>
				<li><a href="/machines" class:active={$page.url.pathname.startsWith('/machines')}>Machines</a></li>
				<li><a href="/billing" class:active={$page.url.pathname === '/billing'}>Billing</a></li>
			</ul>
			<div class="sidebar-footer">
				<a href="https://hijackr.io" target="_blank" rel="noopener">hijackr.io ↗</a>
			</div>
		</nav>
		<main class="content">
			<slot />
		</main>
	</div>
{/if}

<style>
	.shell {
		display: flex;
		height: 100vh;
		overflow: hidden;
	}

	.sidebar {
		width: 220px;
		flex-shrink: 0;
		background: var(--surface);
		border-right: 1px solid var(--border);
		display: flex;
		flex-direction: column;
		padding: 24px 16px;
		gap: 32px;
	}

	.logo {
		display: flex;
		align-items: center;
		gap: 10px;
		padding: 0 8px;
	}

	.logo-mark {
		width: 28px;
		height: 28px;
		background: var(--accent);
		border-radius: 6px;
		display: flex;
		align-items: center;
		justify-content: center;
		font-weight: 700;
		font-size: 14px;
		color: white;
	}

	.logo-text {
		font-weight: 600;
		font-size: 16px;
	}

	.nav-links {
		list-style: none;
		display: flex;
		flex-direction: column;
		gap: 4px;
		flex: 1;
	}

	.nav-links a {
		display: block;
		padding: 8px 12px;
		border-radius: var(--radius);
		color: var(--text-muted);
		font-size: 14px;
		transition: all 0.15s;
	}

	.nav-links a:hover,
	.nav-links a.active {
		background: rgba(99, 102, 241, 0.12);
		color: var(--text);
	}

	.nav-links a.active {
		color: var(--accent);
	}

	.sidebar-footer {
		font-size: 12px;
		color: var(--text-muted);
		padding: 0 8px;
	}

	.content {
		flex: 1;
		overflow-y: auto;
		padding: 40px;
	}
</style>