/** @type {import('./$types').LayoutLoad} */
export async function load({ fetch, params }) {
	const apiRoot = '://127.0.0.1:5002/api/restaurant/' + params.id;
	const response = await fetch('http' + apiRoot);
	const restaurant = await response.json();
	return { restaurant, socket: null };
}
