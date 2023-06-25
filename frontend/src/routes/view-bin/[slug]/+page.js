// @ts-nocheck
import axios from 'axios';
import { error } from '@sveltejs/kit';
import { binRequestHistory } from '../../../store';

const hostname = 'http://localhost:8080';

/** @type {import('./$types').PageLoad} */
export async function load({ params }) {
  const binId = params.slug;

  try {
    const binRequests = (await axios.get(`${hostname}/bin-contents/${binId}`)).data.Data;
    binRequestHistory.set(binRequests.binContents);
  } catch (err) {
    binRequestHistory.set([]);
    throw error(err.response.status, err.response.message);
  }

  return {
    binId,
    hostname
  }
}