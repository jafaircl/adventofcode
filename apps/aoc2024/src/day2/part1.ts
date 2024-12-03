/**
 * Determine whether an array is "safe". An array is "safe" if the numbers are
 * increasing or decreasing by between 1 and 3
 */
export function isSafeArray(arr: number[]): boolean {
  let previousValue: number | null = null;
  // Assume true to start
  let isIncreasing = true;
  for (let i = 0; i < arr.length; i++) {
    const curr = arr[i];
    // On the first pass, just set the previous value
    if (previousValue === null) {
      previousValue = curr;
      continue;
    }
    // Determine whether the array is increasing or decreasing
    if (i === 1) {
      isIncreasing = curr > previousValue;
    } else {
      // If the array is increasing make sure the current value is greater than the previous
      // value. If the array is decreasing, make sure the current value is less than the previous
      // value. If the number is not increasing or decreasing, return false.
      if (
        (isIncreasing && curr < previousValue) ||
        (!isIncreasing && curr > previousValue)
      ) {
        return false;
      }
    }
    // If the numbers differ by more than 3, return false
    const diff = Math.abs(curr - previousValue);
    if (diff > 3 || diff < 1) {
      return false;
    }
    previousValue = curr;
  }
  return true;
}
