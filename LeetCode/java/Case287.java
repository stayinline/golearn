import sun.jvm.hotspot.utilities.BitMap;

import java.util.BitSet;

public class Case287 {

    public int findDuplicate(int[] nums) {
        BitMap bitmap = new BitMap(nums.length);
        for (int i = 0; i < nums.length; i++) {
            boolean isExit = bitmap.at(nums[i]);
            if (isExit) {
                return nums[i];
            }
            bitmap.atPut(nums[i], true);
        }
        return 0;
    }
}
