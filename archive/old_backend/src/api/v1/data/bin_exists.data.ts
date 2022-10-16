import { requestRecords } from './model.data';

export async function binIdPresentInDb(binId: string) {
  return await requestRecords.exists({ bin: binId });
}
