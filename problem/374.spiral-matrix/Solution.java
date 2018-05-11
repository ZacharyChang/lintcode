import java.util.List;
import java.util.ArrayList;

public class Solution {
    /**
     * @param matrix: a matrix of m x n elements
     * @return: an integer list
     */
    public List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> result = new ArrayList<>();
        if (matrix==null){
            return result;
        }
        int m = matrix.length;
        if (m == 0) {
            return result;
        }
        int n = matrix[0].length;
        if (m == 1) {
            for (int i = 0; i < n; i++) {
                result.add(matrix[0][i]);
            }
            return result;
        }
        
        if (n == 1) {
            for (int i = 0; i < m; i++) {
                result.add(matrix[i][0]);
            }
            return result;
        }
        for (int i = 0; i < n; i++) {
            result.add(matrix[0][i]);
        }
        for (int i = 1; i < m - 1; i++) {
            result.add(matrix[i][n - 1]);
        }
        for (int i = n - 1; i >= 0; i--) {
            result.add(matrix[m - 1][i]);
        }
        for (int i = m - 2; i >= 1; i--) {
            result.add(matrix[i][0]);
        }
        if (m - 2 >= 1 && n - 2 >= 1) {
            int[][] nextMatrix = new int[m - 2][n - 2];
            for (int i = 1; i < m - 1; i++) {
                for (int j = 1; j < n - 1; j++) {
                    nextMatrix[i - 1][j - 1] = matrix[i][j];
                }
            }
            result.addAll(spiralOrder(nextMatrix));
        }
        return result;
    }
}