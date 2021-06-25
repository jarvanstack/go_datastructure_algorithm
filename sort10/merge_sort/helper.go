package main

func start(old_array []int)(new_array []int,err error) {
	if old_array==nil || len(old_array)<=1 {
		return nil,err
	}
	//1.未排序的数组
	//fmt.Printf("old_array=%#v\n", old_array)
	//开始排序
	new_array = divide(old_array)
	return new_array, nil
}

//分
func divide(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	middle := len(array) / 2
	left_array := array[:middle]
	right_array := array[middle:]
	//合并.
	return merge(divide(left_array), divide(right_array))
}

//合并
//排序在合并时候进行.
func merge(left_array, right_array []int) []int {
	result_merge := make([]int, 0)
	//左边的长度始终要小于等于右边，知道将左边的元素全部 merge.
	for len(left_array) > 0 {
		if len(right_array) == 0 || left_array[0] < right_array[0] {
			result_merge = append(result_merge, left_array[0])
			left_array = left_array[1:]
		} else {
			result_merge = append(result_merge, right_array[0])
			right_array = right_array[1:]
		}
	}
	//merge 右边的的元素
	for len(right_array) > 0 {
		result_merge = append(result_merge, right_array[0])
		right_array = right_array[1:]
	}
	return result_merge
}
