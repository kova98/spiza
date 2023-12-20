/** @type {import('./$types').PageLoad} */
export async function load({ parent }) {
	const { restaurant, socket } = await parent();
	return { restaurant, socket };
}
