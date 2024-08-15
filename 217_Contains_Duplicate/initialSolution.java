import java.util.HashMap;
import java.util.Map;

class Solution {
    public boolean containsDuplicate(int[] nums) {
        HashMap<Integer, Integer> solutionMap = new HashMap<>();

        for (int num : nums) {
            solutionMap.put(num, solutionMap.getOrDefault(num, 0) + 1);

            // The line above does what these 6 lines do
            // if (solutionMap.containsKey(num)) {
            // solutionMap.replace(num, solutionMap.get(num) + 1);
            // }
            // if (!solutionMap.containsKey(num)) {
            // solutionMap.put(num, 1);
            // }

        }
        for (Map.Entry<Integer, Integer> entry : solutionMap.entrySet()) {
            if (entry.getValue() > 1) {
                return true;
            }
        }
        return false;
    }
}
