import javafx.util.Pair;

import java.util.*;

public class Case347 {

    public int[] topKFrequent(int[] nums, int k) {
        //k:nums[i]  v:count
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            map.put(nums[i], map.getOrDefault(nums[i], 0) + 1);
        }

        //key:count,value:nums[i]
        PriorityQueue<Pair<Integer, Integer>> queue = new PriorityQueue<>(Comparator.comparingInt(Pair::getKey));
        map.forEach((key, v) -> {
            if (queue.size() == k) {
                if (queue.peek().getKey() < v) {
                    queue.poll();
                    queue.offer(new Pair<>(v, key));
                }
            } else {
                queue.offer(new Pair<>(v, key));
            }

        });
        int[] ret = new int[k];
        for (int i = 0; i < k; ++i) {
            ret[i] = queue.poll().getValue();
        }
        return ret;
    }

    public static void main(String[] args) {
        int[] nums = {3,0,1,0};
        new Case347().topKFrequent(nums, 1);
    }
}
