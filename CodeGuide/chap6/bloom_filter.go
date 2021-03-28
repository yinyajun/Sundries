package main

/*
100亿条黑名单url，每个url占64B
url过滤，容忍万分之一以下的误判，使用空间不要超过30GB
如果使用map，所需要的空间为64B*100*10^8 = 640GB

bloom filter代表一个集合，精确判断一个元素是否在集合中
*/
