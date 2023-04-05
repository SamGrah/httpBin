import { requestRecords } from './model.data';

export async function fetchRequests(binId: string) {
  const response = await requestRecords.findOne({'bin': binId})
  return response.requests
}