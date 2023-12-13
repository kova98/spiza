/** @type {import('./$types').PageLoad} */
export async function load({ parent }) {
	const { restaurant } = await parent();
	return { restaurant };
}
