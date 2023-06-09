import { createNewBin, binIdPresentInDb } from "../data/index.data";
import os from "os";

const hostname = os.hostname();

const generateRandomBin = (): string => {
  let result = "";
  const chars = "abcdefghijklmnopqrstuvwxyx0123456789";
  for (let idx = 0; idx < 6; idx += 1) {
    result += chars.charAt(Math.floor(Math.random() * chars.length));
  }
  return result;
};

export const createUniqueBin = async () => {
  let binId: string; 
  while (true) {
    binId = generateRandomBin();
    const binIdIsPresent = await binIdPresentInDb(binId);
    if (!binIdIsPresent) break;
  }

  const newBin = await createNewBin(binId);
  return newBin 
};
