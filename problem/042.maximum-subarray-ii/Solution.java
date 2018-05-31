import java.util.List;
import java.util.ArrayList;

public class Solution {
    /*
     * Dynamic planning
     * @param nums: A list of integers
     * @return: An integer denotes the sum of max two non-overlapping subarrays
     */
    public int maxTwoSubArrays(List<Integer> nums) {
        if(nums.size()==0){
            return 0;
        }
        if(nums.size()==2){
            return nums.get(0)+nums.get(1);
        }
        // 从左往右计算最大和
        int[] list1 = new int[nums.size()];
        // 从右往左计算最大和
        int[] list2 = new int[nums.size()];
        int maxLeft = 0;
        int maxRight = 0;
        for(int i=0;i<nums.size();i++){
            maxLeft = max(maxLeft+nums.get(i),nums.get(i));
            list1[i] = maxLeft;
            maxRight =  max(maxRight+nums.get(nums.size()-1-i),nums.get(nums.size()-1-i));
            list2[nums.size()-1-i] =maxRight;
        }
        int result=nums.get(0)+nums.get(nums.size()-1);
        // 计算两个数组不重合子数组的最大值
        for(int i=0;i<nums.size()-1;i++){
            for(int j=i+1;j<nums.size();j++){
                if(list1[i]+list2[j]>result){
                    result = list1[i] + list2[j];
                }
            }
        }
        return result;
    }

    /*
     * Divide and conquer
     * @param nums: A list of integers
     * @return: An integer denotes the sum of max two non-overlapping subarrays
     */
    public int maxTwoSubArrays_2(List<Integer> nums) {
        if(nums.size()==0){
            return 0;
        }
        if(nums.size()==2){
            return nums.get(0)+nums.get(1);
        }
        int result=nums.get(0)+nums.get(nums.size()-1);
        for(int i=0;i<nums.size()-1;i++){
            int maxLeft=maxSubarray(nums,0,i);
            int maxRight=maxSubarray(nums,i+1,nums.size()-1);
            result=max(result,maxLeft+maxRight);
        }
        return result;
    }

    public int maxSubarray(List<Integer> nums,int left,int right){
        int result=nums.get(left);
        if (left==right){
            return result;
        }
        int maxNum=nums.get(left);
        for(int i=left+1;i<=right;i++){
            maxNum=max(maxNum+nums.get(i),nums.get(i));
            if(maxNum>result){
                result=maxNum;
            }
        }
        return result;
    }

    public int max(int a,int b){
        if(a>b){
            return a;
        }
        return b;
    }
}