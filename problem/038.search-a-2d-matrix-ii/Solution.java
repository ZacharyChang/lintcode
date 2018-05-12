public class Solution {
    /**
     * @param matrix: matrix, a list of lists of integers
     * @param target: An integer
     * @return: a boolean, indicate whether matrix contains target
     */
    public int searchMatrix(int[][] matrix, int target) {
        if (matrix.length == 0) {
            return 0;
        }

        int count = 0;
        int start = 0;
        int end = matrix.length - 1;
        int row = end;

        // 查找target可能出现的最后行(行首元素小于等于target)
        while (start < (end - 1)) {
            int mid = (start + end) / 2;

            if (matrix[mid][0] == target) {
                row = mid;
                break;
            } else if (matrix[mid][0] > target) {
                end = mid;
                row = mid;
            } else if (matrix[mid][0] < target) {
                start = mid;
            }
        }
        for(int i=0;i<=row;i++){
            if(binarySearch(matrix[i],target)){
                count++;
            }
        }
        return count;
    }

    public boolean binarySearch(int[] array, int target) {
        if (array.length == 0) {
            return false;
        }
        if (array[0]>target||array[array.length-1]<target){
            return false;
        }
        if (array[0]==target||array[array.length-1]==target){
            return true;
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

        return false;
    }

    public static void main(String[] args){
        int[][] array = new int[][]{
            {
                1, 3, 5, 7
            },
            {
                2,4,7,8
            },
            {
                3,5,9,10
            }
        };
        System.out.println(new Solution().searchMatrix(array,3)==2);
    }
}
