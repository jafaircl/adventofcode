import { col1, col2 } from './data';
import { calculateDistance } from './part1';

describe('Day 1 - Part 1', () => {
  it('calculateDistance', () => {
    expect(calculateDistance([3, 4, 2, 1, 3, 3], [4, 3, 5, 3, 9, 3])).toBe(11);
  });

  it('solution', () => {
    expect(calculateDistance(col1, col2)).toBe(1530215);
  });
});
