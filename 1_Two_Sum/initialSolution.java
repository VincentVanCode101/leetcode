import java.util.HashMap;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> alreadySeen = new HashMap<>();
        int[] resultIndexes = new int[2];
        for (int i = 0; i <= nums.length - 1; i++) {
            int restThatsNeeded = target - nums[i];
            if (alreadySeen.containsKey(restThatsNeeded)) {
                resultIndexes[0] = alreadySeen.get(restThatsNeeded);
                resultIndexes[1] = i;
                return resultIndexes;
            } else {
                alreadySeen.put(nums[i], i);
            }

        }
        return resultIndexes;
    }
}
