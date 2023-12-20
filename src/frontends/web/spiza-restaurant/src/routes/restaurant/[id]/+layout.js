/** @type {import('./$types').LayoutLoad} */
export async function load({ fetch, params }) {
	const apiRoot = `${import.meta.env.VITE_HTTP_ROOT}/api/restaurant/${params.id}`;
	const response = await fetch(apiRoot);
	const restaurant = await response.json();
	return { restaurant, socket: null };
}
