import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Solution {
    private Map<Integer, Integer> numberMap = new HashMap<>();

    /*
     * @param nums: a list of integers
     * 
     * @return: find a majority number
     */
    public int majorityNumber(List<Integer> nums) {
        for (int i = 0; i < nums.size(); i++) {
            int num = nums.get(i);
            if (numberMap.containsKey(num)) {
                int count = numberMap.get(num) + 1;
                if (count >= nums.size() / 2) {
                    return num;
                }
                numberMap.put(num, count);
            } else {
                numberMap.put(num, 1);
            }
        }
        return nums.get(0);
    }
}