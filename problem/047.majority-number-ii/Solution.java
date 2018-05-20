import java.util.List;
import java.util.Map;
import java.util.HashMap;

public class Solution {
    /*
     * @param nums: a list of integers
     * @return: The majority number that occurs more than 1/3
     */
    public int majorityNumber(List<Integer> nums) {
        if (nums.size() == 0) {
            return -1;
        }
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.size(); i++) {
            int num = nums.get(i);
            if (map.containsKey(num)) {
                int count = map.get(num) + 1;
                if (count > nums.size() / 3) {
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