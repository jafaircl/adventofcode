import { radixSort } from './radix-sort';

describe('radixSort', () => {
  it('radixSort', () => {
    // Test the radixSort function
    expect(radixSort([1, 2, 3, 4, 5])).toEqual([1, 2, 3, 4, 5]);
    expect(radixSort([5, 4, 3, 2, 1])).toEqual([1, 2, 3, 4, 5]);
    expect(radixSort([5, 4, 3, 2, 1, 0])).toEqual([0, 1, 2, 3, 4, 5]);
    expect(radixSort([1, 33, 444, 0, 3, 2])).toEqual([0, 1, 2, 3, 33, 444]);
    expect(radixSort([1, 33, 444, 0, 3, 2, 12345, 123456])).toEqual([
      0, 1, 2, 3, 33, 444, 12345, 123456,
    ]);
    expect(
      radixSort([1, 33, 444, 0, 3, 2, 12345, 123456, -1, -33, -444])
    ).toEqual([-444, -33, -1, 0, 1, 2, 3, 33, 444, 12345, 123456]);
  });
});
