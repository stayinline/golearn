import java.util.Deque;
import java.util.LinkedList;

public class Case239 {


    public int[] maxSlidingWindow(int[] nums, int k) {
        if (nums.length == 0 || k == 0) {
            return new int[0];
        }
        Deque<Integer> deque = new LinkedList<>();
        int[] res = new int[nums.length - k + 1];
        //i=1-k;是为了初始化处理第一个窗口中的所有值，
        // 所以，更新队列、最大值的时候需判断i>0
        for (int j = 0, i = 1 - k; j < nums.length; i++, j++) {
            // 删除 deque 中对应的 nums[i-1]，
            // 因为队列中头部的这个nums[i - 1]元素已经记录过了，
            // 窗口向后滑动的时候，应该删掉
            if (i > 0 && deque.peekFirst() == nums[i - 1]) {
                deque.removeFirst();
            }
            // 保持 deque 递减
            while (!deque.isEmpty() && deque.peekLast() < nums[j]) {
                deque.removeLast();
            }
            deque.addLast(nums[j]);
            // 记录窗口最大值
            if (i >= 0) {
                res[i] = deque.peekFirst();
            }
        }
        return res;
    }

}
