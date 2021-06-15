import java.util.*;

public class Case215 {
    public int findKthLargest(int[] nums, int k) {

        Arrays.sort(nums);
        return nums[nums.length - k];
    }

    public int findKthLargest2(int[] nums, int k) {
        PriorityQueue<Integer> queue = new PriorityQueue<>();
        for (int i = 0; i < nums.length; i++) {
            if (i < k) {
                queue.add(nums[i]);
            } else {
                if (nums[i] > queue.peek()) {
                    queue.poll();
                    queue.add(nums[i]);
                }
            }
        }
        return queue.peek();
    }

    public static void main(String[] args) {
        int[] nums = {3, 2, 1, 5, 6, 4};
        int kthLargest = new Case215().findKthLargest2(nums, 2);
        System.out.println(kthLargest);
    }
}
