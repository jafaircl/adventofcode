import { col1, col2 } from './data';
import { calculateSimilarity } from './part2';

describe('Day 1 - Part 2', () => {
  it('calculateSimilarity', () => {
    // Test the calculateSimilarity function
    expect(calculateSimilarity([3, 4, 2, 1, 3, 3], [4, 3, 5, 3, 9, 3])).toBe(
      31
    );
  });

  it('solution', () => {
    expect(calculateSimilarity(col1, col2)).toBe(26800609);
  });
});
