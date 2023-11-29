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




Problem 2: Add Two Numbers

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:
Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]


Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.

Solution: 

Approach
1. Run a loop to traverse list1 and list2. 
2. Calculate the total by adding values of corresponding nodes from list1, list2 and the value of carry. If node is null, then add only the values we have. 
3. Update carry and total. 
4. Create a new node in the list3 with value as the total and make the list point to it. 
5. Move list3 to the next node (list3.Next).
6. If there is a carry left after the loop exits, create an additional node for it and append it to the list. 
7. Return the head of the resulting linked list. 

Time complexity is O(n)
Space complexity is O(n)
===n is maximum lengths of l1 and l2===
