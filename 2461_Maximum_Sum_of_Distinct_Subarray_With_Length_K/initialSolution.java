import java.util.HashSet;

class Solution {
    public long maximumSubarraySum(int[] nums, int k) {
        long maxSum = 0;
        long currentSum = 0;
        HashSet<Integer> windowSet = new HashSet<>();

        int left = 0;
        for (int right = 0; right < nums.length; right++) {
            while (windowSet.contains(nums[right])) {
                windowSet.remove(nums[left]);
                currentSum -= nums[left];
                left++;
            }

            windowSet.add(nums[right]);
            currentSum += nums[right];

            if (right - left + 1 == k) {
                maxSum = Math.max(maxSum, currentSum);
                windowSet.remove(nums[left]);
                currentSum -= nums[left];
                left++;
            }
        }

        return maxSum;
    }
}
