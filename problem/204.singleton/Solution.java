class Solution {
    private static Solution instance;
    /**
     * @return: The same instance of this class every time
     */
    public static Solution getInstance() {
        if (Solution.instance == null){
            Solution.instance = new Solution();
        }
        return Solution.instance;
    }
};