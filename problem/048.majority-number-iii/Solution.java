import java.util.List;
import java.util.Map;
import java.util.HashMap;

public class Solution {
    /**
     * @param nums: A list of integers
     * @param k: An integer
     * @return: The majority number
     */
    public int majorityNumber(List<Integer> nums, int k) {
        if (nums.size() == 0) {
            return -1;
        }
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.size(); i++) {
            int num = nums.get(i);
            if (map.containsKey(num)) {
                int count = map.get(num) + 1;
                if (count > nums.size() / k) {
                    return num;
                }
                map.put(num, count);
            } else {
                map.put(num, 1);
            }
        }
        return nums.get(0);
    }
}