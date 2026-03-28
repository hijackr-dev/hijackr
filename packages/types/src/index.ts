// ─── User & Auth ────────────────────────────────────────────────────────────

export interface User {
  id: string;
  email: string;
  name: string;
  createdAt: string;
}

// ─── Machines ────────────────────────────────────────────────────────────────

export type MachineStatus = 'online' | 'offline' | 'provisioning' | 'error';

export interface Machine {
  id: string;
  name: string;
  status: MachineStatus;
  region: string;
  ipAddress?: string;
  createdAt: string;
  lastSeenAt?: string;
}

// ─── Billing ─────────────────────────────────────────────────────────────────

export type PlanTier = 'free' | 'pro' | 'enterprise';

export interface Subscription {
  id: string;
  userId: string;
  plan: PlanTier;
  status: 'active' | 'cancelled' | 'past_due' | 'trialing';
  currentPeriodEnd: string;
  cancelAtPeriodEnd: boolean;
}

export interface Invoice {
  id: string;
  amount: number; // pence / cents
  currency: string;
  status: 'paid' | 'open' | 'void' | 'uncollectible';
  createdAt: string;
  hostedUrl?: string;
}

// ─── API Responses ────────────────────────────────────────────────────────────

export interface ApiError {
  code: string;
  message: string;
}

export interface PaginatedResponse<T> {
  data: T[];
  total: number;
  page: number;
  perPage: number;
}