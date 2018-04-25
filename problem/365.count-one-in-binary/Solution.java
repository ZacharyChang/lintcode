public class Solution {
    /*
     * num和num-1进行与运算，可以使原来num中的最后一个1变为0
     * @param num: An integer
     * @return: An integer
     */
    public int countOnes(int num) {
        int result = 0;
        while(num!=0){
            num = num & (num - 1);
            result++;
        }
        return result;
    }

    public static void main(String[] args){
        System.out.println(new Solution().countOnes(32) == 1);
        System.out.println(new Solution().countOnes(5) == 2);
        System.out.println(new Solution().countOnes(1023) == 10);
        System.out.println(new Solution().countOnes(-1) == 32);
        System.out.println("OK");
    }
};