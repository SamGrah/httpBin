import { requestRecords, RequestProperties } from './model.data';

export async function addRequestToBin(binId: string, requestProps: RequestProperties) {
  return await requestRecords.updateOne(
    {'bin': binId},
    {$push: {'requests': requestProps}} 
  )
}