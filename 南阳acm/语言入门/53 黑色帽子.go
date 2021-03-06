/*
黑色帽子
时间限制：1000 ms  |  内存限制：65535 KB
难度：1
描述
        最近发现了一个搞笑的游戏，不过目前还没玩过。一个舞会上，每个人头上都戴着一顶帽子，帽子只有黑白两种，黑的至少有一顶。每个人都能看到别人帽子的颜色，可是看不见自己的。主持人先让大家看看别人头上戴的是什么帽子，然后关灯，如果有人认为自己戴的的黑色帽子，就打自己一个耳光（，都很自觉，而且不许打别人的哦），开灯，关灯，开灯……。因为都不想打自己耳光，所以不确定的情况下都不会打自己的，现在有n顶黑色帽子，第几次关灯才会听到有人打自己耳光？
输入
第一行只有一个整数m（m<=100000),表示测试数据组数。
接下来的m行，每行有一个整数n(n<=100000000)，表示黑色帽子的顶数。
输出
输出第几次关灯能听到耳光声，每组输出占一行。
样例输入
1
2
样例输出
2

分析：来自南阳acm讨论区复制
他们一开始只知道至少有一顶黑帽子，是不知道具体有几只黑帽子的。
假设n=1.第一次开灯，带走黑帽子的那个人看其他的人都是白帽子，但是他知道至少有一顶黑帽子，那么他就确认只有他带了黑帽子。所以第一次关灯后他会自己打自己。
假设n=2.假设带有黑帽子的为甲和乙。第一次开灯，甲看到只有乙一个人带走黑帽子，乙看到只有甲一个人带走黑帽子，但是由于他们都知道至少有一顶黑帽子，所以不确定自己带的也是黑帽子，所以第一次关灯后不会自己打自己。接着第二次开灯，因为甲和乙看到第一次关灯后没有一个人自己打自己，所以他们知道黑帽子至少有2顶，但是由于甲知道乙带了黑帽子，乙知道甲带了黑帽子，但是黑帽子至少有2顶了，而且其他人又都是白帽子，那么甲知道自己带的是黑帽子，乙也知道自己带的是黑帽子，第二次关灯后就会听到2个响声。之后都是这样以此类推
 */
package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m)
	for i := 0;i < m;i++ {
		fmt.Scan(&n)
		fmt.Println(n)
	}
}