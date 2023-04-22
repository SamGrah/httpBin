import axios from 'axios';
import { error } from '@sveltejs/kit';
import { binRequestHistory } from '../../../store';

const hostname = 'http://localhost:8080';

/** @type {import('./$types').PageLoad} */
export function load({ params }) {
  const binId = params.slug;

  (async function () {
    const {data} = await axios.get(`${hostname}/bin-contents/${binId}`);
    console.log(data)
    binRequestHistory.set(data.Data.binContents);
  })()

  return {
    binId,
    hostname
  }
}