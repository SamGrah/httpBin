import {fetchRequests} from "../data/index.data";

export async function fetchBinRequests(binId: string) {
  return await fetchRequests(binId)
}