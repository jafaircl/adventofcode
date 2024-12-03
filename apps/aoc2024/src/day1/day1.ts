// Need a good sorting algorithm for this one. Radix sort seems fun.

function getDigit(num: number, place: number) {
  return Math.floor(Math.abs(num) / Math.pow(10, place)) % 10;
}

function digitCount(num: number) {
  if (num === 0) return 1;
  return Math.floor(Math.log10(Math.abs(num))) + 1;
}

function mostDigits(nums: number[]) {
  let maxDigits = 0;
  for (let i = 0; i < nums.length; i++) {
    maxDigits = Math.max(maxDigits, digitCount(nums[i]));
  }
  return maxDigits;
}

export function radixSort(arr: number[]): number[] {
  let result: number[] = [...arr];
  const maxDigits = mostDigits(arr);
  for (let k = 0; k < maxDigits; k++) {
    const buckets: number[][] = Array.from({ length: 10 }, () => []);
    for (let i = 0; i < result.length; i++) {
      const digit = getDigit(result[i], k);
      buckets[digit].push(result[i]);
    }
    result = ([] as number[]).concat(...buckets);
  }
  return result;
}

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
