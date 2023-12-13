/** @type {import('./$types').LayoutLoad} */
export async function load({ fetch, params }) {
	const apiRoot = 'http://127.0.0.1:5002/api/restaurant';
	const response = await fetch(apiRoot + '/' + params.id);
	const restaurant = await response.json();
	return { restaurant };
}
