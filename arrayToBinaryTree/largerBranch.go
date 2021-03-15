package main

import (
	math "math"
)

func largerBranch(arr []int) string {
	level := 0
	leftsum, rightsum := 0, 0
	for index := 0; index < len(arr); {
		if level == 0 {
			index, level = index+1, level+1
			continue
		}
		levelPow := int(math.Pow(2, float64(level)))
		for j := index; j < index+levelPow; j++ {
			if j >= len(arr) {
				break
			}
			if arr[j] == -1 {
				continue
			}
			if j-index <= (levelPow/2)-1 {
				leftsum += arr[j]
			} else if j-index > (levelPow/2)-1 {
				rightsum += arr[j]
			}
		}
		index, level = index+levelPow, level+1
	}

	//fmt.Printf("leftsum : %v, rightsum: %v\n", leftsum, rightsum)

	if leftsum > rightsum {
		return "Left"
	} else if rightsum > leftsum {
		return "Right"
	} else {
		return ""
	}
}
