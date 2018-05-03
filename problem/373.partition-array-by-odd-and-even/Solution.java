import java.util.ArrayList;
import java.util.List;

public class Solution {
    /*
     * @param nums: an array of integers
     * @return: nothing
     */
    public void partitionArray(int[] nums) {
        quick(nums,0,nums.length-1);
    }

    public void quick(int[] nums,int left,int right){
        int i = left;
        int j = right;
        while(i<j){
            while(i<j&&nums[j]%2==0){
                j--;
            }
            while(i<j&&nums[i]%2==1){
                i++;
            }
            if(i<j){
                int temp=nums[i];
                nums[i]=nums[j];
                nums[j]=temp;
            }
        }
    }

    public static void main(String[] args) {
        int[] nums = new int[]{1,2,3,4};
        Solution solution = new Solution();
        solution.partitionArray(nums);
        for(int i=0;i<nums.length;i++){
            System.out.print(nums[i]+",");
        }
        System.out.println();
    }
}