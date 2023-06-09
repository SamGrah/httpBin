import { addRequestToBin, RequestProperties } from "../data/index.data";

export async function logRequest(binId: string, requestProps: RequestProperties) {
  return await addRequestToBin(binId, requestProps)
}
