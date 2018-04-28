public class Solution {
    /**
     * @param matrix: matrix, a list of lists of integers
     * @param target: An integer
     * @return: a boolean, indicate whether matrix contains target
     */
    public boolean searchMatrix(int[][] matrix, int target) {
        if (matrix.length == 0) {
            return false;
        }

        int start = 0;
        int end = matrix.length - 1;

        while (start < (end - 1)) {
            int mid = (start + end) / 2;

            if (matrix[mid][0] == target) {
                return true;
            } else if (matrix[mid][0] > target) {
                end = mid;
            } else if (matrix[mid][0] < target) {
                start = mid;
            }
        }

        return binarySearch(matrix[start], target) || binarySearch(matrix[end], target);
    }

    public boolean binarySearch(int[] array, int target) {
        if (array.length == 0) {
            return false;
        }

        int start = 0;
        int end = array.length - 1;

        while (start < (end - 1)) {
            int mid = (start + end) / 2;

            if (array[mid] == target) {
                return true;
            } else if (array[mid] > target) {
                end = mid;
            } else if (array[mid] < target) {
                start = mid;
            }
        }

        return (array[start] == target) || (array[end] == target);
    }
}
