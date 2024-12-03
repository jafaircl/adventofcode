import { arrs } from './data';
import { isSafeArray } from './part1';

describe('Day 2 - Part 1', () => {
  it('isSafeArray', () => {
    expect(isSafeArray([1, 2, 3, 4, 5])).toBe(true);
    expect(isSafeArray([5, 4, 3, 2, 1])).toBe(true);

    expect(isSafeArray([7, 6, 4, 2, 1])).toBe(true);
    expect(isSafeArray([1, 2, 7, 8, 9])).toBe(false);
    expect(isSafeArray([9, 7, 6, 2, 1])).toBe(false);
    expect(isSafeArray([1, 3, 2, 4, 5])).toBe(false);
    expect(isSafeArray([8, 6, 4, 4, 1])).toBe(false);
    expect(isSafeArray([1, 3, 6, 7, 9])).toBe(true);
  });

  it('solution', () => {
    let safeCount = 0;
    for (const row of arrs) {
      if (isSafeArray(row)) {
        safeCount++;
      }
    }
    expect(safeCount).toBe(321);
  });
});
