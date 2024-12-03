import { radixSort } from '../common/radix-sort';

/**
 * Given two columns of numbers, sort them and calculate the distance between each pair of numbers.
 */
export function calculateDistance(col1: number[], col2: number[]): number {
  const sortedCol1 = radixSort(col1);
  const sortedCol2 = radixSort(col2);
  let distance = 0;
  for (let i = 0; i < col1.length; i++) {
    distance += Math.abs(sortedCol1[i] - sortedCol2[i]);
  }
  return distance;
}
