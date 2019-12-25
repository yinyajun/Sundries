import java.util.LinkedList;
import java.util.List;
import java.util.TreeMap;

public class Solution {
    //N个元素中选出前M个元素，如果用排序，复杂度为O（NlgN）
    //如果使用优先队列，复杂度为O（NlgM）
    //使用优先队列，维护前M个元素
    //使用最小堆的优先队列，每次遍历一个新元素的时候(它大于队首)，将最小的元素出队，那么最小堆维护的就是前M个元素
    //等效的，使用最大堆的优先队列，可以将小的元素设置为高优先级

    //最大堆 维护前M个
    //最小堆 维护后M个

    private class Freq implements Comparable<Freq> {
        public int e, freq;

        public Freq(int e, int freq) {
            this.e = e;
            this.freq = freq;
        }

        @Override
        public int compareTo(Freq o) {
            if (this.freq < o.freq)
                return 1;
            else if (this.freq > o.freq)
                return -1;
            else
                return 0;
        }
    }


    public List<Integer> topKFreq(int[] nums, int k) {
        // 用Map得到频次
        TreeMap<Integer, Integer> map = new TreeMap<>();
        for (int num : nums) {
            if (map.containsKey(num))
                map.put(num, map.get(num) + 1);
            else
                map.put(num, 1);
        }

        // 用Priority Queue得到前K个
        PriorityQueue<Freq> pq = new PriorityQueue<>();
        for(int key:map.keySet()){
            if (pq.getSize()<k)
                pq.enqueue(new Freq(key, map.get(key)));
            else{
                if(map.get(key)>pq.getFront().freq){
                    pq.dequeue();
                    pq.enqueue(new Freq(key,map.get(key)));
                }
            }
        }

        // 用list得到返回结果
        LinkedList<Integer> ret = new LinkedList<>();
        while (!pq.isEmpty()) {
            ret.add(pq.dequeue().e);
        }

        return ret;
    }

}




