import { requestRecords } from "./model.data";

export async function createNewBin(binId: string) {
  return await requestRecords.create({bin: binId, requests: []})
}