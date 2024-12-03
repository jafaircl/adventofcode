import { arrs } from './data';
import { isSafeArrayWithOneDrop } from './part2';

describe('Day 2 - Part 2', () => {
  it('isSafeArrayWithOneDrop', () => {
    expect(isSafeArrayWithOneDrop([1, 2, 3, 4, 5])).toBe(true);
    expect(isSafeArrayWithOneDrop([5, 4, 3, 2, 1])).toBe(true);
    expect(isSafeArrayWithOneDrop([5, 5, 3, 2, 1])).toBe(true);
    expect(isSafeArrayWithOneDrop([5, 5, 5, 2, 1])).toBe(false);

    expect(isSafeArrayWithOneDrop([7, 6, 4, 2, 1])).toBe(true);
    expect(isSafeArrayWithOneDrop([1, 2, 7, 8, 9])).toBe(false);
    expect(isSafeArrayWithOneDrop([9, 7, 6, 2, 1])).toBe(false);
    expect(isSafeArrayWithOneDrop([1, 3, 2, 4, 5])).toBe(true);
    expect(isSafeArrayWithOneDrop([8, 6, 4, 4, 1])).toBe(true);
    expect(isSafeArrayWithOneDrop([1, 3, 6, 7, 9])).toBe(true);
  });

  it('solution', () => {
    let safeCount = 0;
    for (const row of arrs) {
      if (isSafeArrayWithOneDrop(row)) {
        safeCount++;
      }
    }
    expect(safeCount).toBe(386);
  });
});
