Problems from LeetCode

Problem 1: Two Sum

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.


Solution: 

Approach
1. Run two loops.
2. Outer loop runs from 0 to n-2, and inner loop runs from i+1 to n-1.
3. In each iteration check if the numbers of inner loop and outer loop add up to the target sum.
4. If nums[outerIndex] + nums[innerIndex] is equal to target, return {outerIndex, InnerIndex}. Else, continue.
5. Repeat.

Time complexity is O(n^2)

Space complexity is O(1)
