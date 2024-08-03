// @ts-nocheck
import { PUBLIC_API_HOSTNAME } from '$env/static/public'
import { error } from '@sveltejs/kit';

export const load = async ({ params }) => {
  const binId = () => params.slug;
  const hostname = () => `${PUBLIC_API_HOSTNAME}`;

  const fetchBinContents = async () => {
    const uri = () => `http://${hostname()}/bin-contents/${binId()}`;
    try {
      const res = await fetch(uri());
      console.log(`request made to ${uri} for bin contents`);
      const binRequests = (await res.json()).Data;
      console.log('binRequests', binRequests);
      return binRequests.binContents;
    } catch (err) {
      console.log(err)
      // throw error(err.response.status, err.response.message);
      throw error(404, 'Bin not found');
    }
  }

  return {
    binId: binId(),
    binContents: fetchBinContents(),
    hostname: hostname(),
  }
}