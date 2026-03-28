// API client for hijackr-api
// All requests go to the API_BASE URL with the session token in the Authorization header.

export const API_BASE =
	typeof window !== 'undefined' && window.location.hostname === 'localhost'
		? 'http://localhost:8080'
		: 'https://api.hijackr.io';

function getToken(): string | null {
	if (typeof localStorage === 'undefined') return null;
	return localStorage.getItem('session_token');
}

async function request<T>(path: string, options: RequestInit = {}): Promise<T> {
	const token = getToken();
	const headers: Record<string, string> = {
		'Content-Type': 'application/json',
		...(options.headers as Record<string, string>)
	};
	if (token) {
		headers['Authorization'] = `Bearer ${token}`;
	}

	const res = await fetch(`${API_BASE}${path}`, { ...options, headers });

	if (res.status === 401) {
		// Session expired — clear and redirect to login
		localStorage.removeItem('session_token');
		window.location.href = '/';
		throw new Error('Unauthorized');
	}

	if (!res.ok) {
		const body = await res.json().catch(() => ({}));
		throw new Error(body.error || `HTTP ${res.status}`);
	}

	return res.json();
}

// ── Auth ──────────────────────────────────────────────────────────────────────

export async function requestMagicLink(email: string): Promise<{ message: string; magic_link?: string }> {
	return request('/v1/auth/magic-link', {
		method: 'POST',
		body: JSON.stringify({ email })
	});
}

export async function verifyMagicLink(token: string): Promise<{ session_token: string; expires_at: string }> {
	return request(`/v1/auth/verify?token=${encodeURIComponent(token)}`);
}

// ── Account ───────────────────────────────────────────────────────────────────

export interface Customer {
	id: string;
	email: string;
	created_at: string;
}

export async function getMe(): Promise<Customer> {
	return request('/v1/me');
}

// ── Licences ──────────────────────────────────────────────────────────────────

export interface Licence {
	id: string;
	product: string;
	tier: string;
	licence_key: string;
	status: string;
	machine_limit: number;
	expires_at: string | null;
	created_at: string;
}

export async function getLicences(): Promise<{ licences: Licence[] }> {
	return request('/v1/licences');
}

// ── Machines ──────────────────────────────────────────────────────────────────

export interface Machine {
	id: string;
	machine_id: string;
	machine_name: string;
	activated_at: string;
	last_seen_at: string;
}

export async function getMachines(licenceId: string): Promise<{ machines: Machine[] }> {
	return request(`/v1/licences/${licenceId}/machines`);
}

export async function deactivateMachine(licenceId: string, machineId: string): Promise<void> {
	return request(`/v1/licences/${licenceId}/machines/${machineId}`, { method: 'DELETE' });
}

// ── Billing ───────────────────────────────────────────────────────────────────

export async function createBillingPortalSession(): Promise<{ url: string }> {
	return request('/v1/billing/portal', { method: 'POST' });
}