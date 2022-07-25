package distance_between_bus_stops

/*
1184. 公交站间的距离

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/distance-between-bus-stops/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	path1 := 0
	path2 := 0
	if destination < start {
		start, destination = destination, start
	}
	for i := 0; i < len(distance); i++ {
		if i >= start && i < destination {
			path1 = path1 + distance[i]
		}
		path2 = path2 + distance[i]
	}

	return min(path2-path1, path1)
}
