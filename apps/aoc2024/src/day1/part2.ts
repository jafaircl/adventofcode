/**
 * Given 2 columns of numbers, calculate a similarity score by adding up the numbers in the first column
 * after multiplying them by the number of times that number occurs in the second column.
 */
export function calculateSimilarity(col1: number[], col2: number[]) {
  // First, iterate through the second column to create a map of numbers to their occurences.
  const col2Map = new Map<number, number>();
  for (const num of col2) {
    col2Map.set(num, (col2Map.get(num) || 0) + 1);
  }

  // Now, iterate through the first column and calculate the similarity score.
  let similarity = 0;
  for (const num of col1) {
    similarity += num * (col2Map.get(num) || 0);
  }
  return similarity;
}
