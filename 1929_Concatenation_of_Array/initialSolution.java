
class Solution {
    public int[] getConcatenation(int[] nums) {
        int lengthOfNums = nums.length;
        int[] result = new int[lengthOfNums + lengthOfNums];

        System.arraycopy(nums, 0, result, 0, lengthOfNums);
        System.arraycopy(nums, 0, result, lengthOfNums, lengthOfNums);
        return result;

    }
}