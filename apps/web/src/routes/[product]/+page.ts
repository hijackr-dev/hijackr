import { getProduct } from '$lib/products';
import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = ({ params }) => {
	const product = getProduct(params.product);
	if (!product) {
		throw error(404, `Product "${params.product}" not found`);
	}
	return { product };
};