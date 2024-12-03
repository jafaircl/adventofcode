import { isSafeArray } from './part1';

/**
 * If an array of numbers would be safe by dropping one number, treat it as safe
 */
export function isSafeArrayWithOneDrop(arr: number[]): boolean {
  // If the array is already safe, return true
  if (isSafeArray(arr)) {
    return true;
  }
  // Iterate through the array and drop one number at a time. If the array is safe after dropping
  // a number, return true.
  for (let i = 0; i < arr.length; i++) {
    const newArr = [...arr];
    newArr.splice(i, 1);
    if (isSafeArray(newArr)) {
      return true;
    }
  }
  return false;
}
