package main

import "fmt"

// Solution overview:
// We can construct an array of contiguous differences
// they  generally follow the following patterns
// 1. + + + + : sorted array
// 2. + + - + : contiguously swapped elements
// 3. - - + + : reversed subsegment
// 4. + - + - : randomly swapped elements

//  for 1. we return "yes"
// for 2. check that the implied swap would make the array sorted.If so return yes, else return no
//  for 3. We check that there is only one chain of - in the diffArray and return "yes" if not "no"
//  for 4. we check that the implied swap would make the array sorted. If so "yes" else "no"

type pair struct {
	element int32
	first int32
	second int32
}
// result in a sorted arrray
func  main(){
	arr := []int32{4104,8529,49984,54956,63034,82534,84473,86411,92941,95929,108831,894947,125082,137123,137276,142534,149840,154703,174744,180537,207563,221088,223069,231982,249517,252211,255192,260283,261543,262406,270616,274600,274709,283838,289532,295589,310856,314991,322201,339198,343271,383392,385869,389367,403468,441925,444543,454300,455366,469896,478627,479055,484516,499114,512738,543943,552836,560153,578730,579688,591631,594436,606033,613146,621500,627475,631582,643754,658309,666435,667186,671190,674741,685292,702340,705383,722375,722776,726812,748441,790023,795574,797416,813164,813248,827778,839998,843708,851728,857147,860454,861956,864994,868755,116375,911042,912634,914500,920825,979477}
	almostSorted(arr)
}

func almostSorted(arr []int32) {
	if (len(arr)==1){
		fmt.Println("already sorted")
		return
	} 
	diffArr := []pair{}
	index  := 1
	for ;index < len(arr); index ++ {
		entry := pair{element:arr[index] - arr[index-1], first:arr[index], second:arr[index-1]}
		diffArr = append(diffArr, entry)
	}
	negCount, ind, lastInd := countNegatives(diffArr)
	if (negCount==0){
		fmt.Println("yes")
	} else if (negCount == 1){
		val := diffArr[ind]
		if (len(diffArr)==1){
			fmt.Println("yes")
			fmt.Println("swap", 1, 2)
		}else if (ind ==0) {
			if (val.second < diffArr[ind+1].first){
				fmt.Println("yes")
				fmt.Println("swap", ind+1, ind+2)
			} else {
				fmt.Println("no")
			}
		}else if (val.first >= diffArr[ind-1].second){
			fmt.Println("yes")
			fmt.Println("swap",  ind+1, ind+2)
		} else {
			fmt.Println("no")
		}
	} else if (negCount ==2){
		val := diffArr[ind]
		if(diffArr[ind+1].element <0){
			fmt.Println("yes")
			fmt.Println("reverse", ind+1, ind +3)
		} else if (diffArr[ind+1].element >=0){
			next := diffArr[lastInd]
			if (next.first > val.first){
				fmt.Println("no")
			} else {
				fmt.Println("yes")
				fmt.Println("swap",  ind+1, lastInd+2)
			}
		}
	} else {
		if (lastInd - ind +1 != negCount){
			fmt.Println("no")
		} else {
			fmt.Println("yes")
			fmt.Println("reverse", ind+1,lastInd+2)
		}
	}
}

func countNegatives(arr []pair) (int32, int32, int32){
	var count int32 = 0
	var index int32 = -1
	var lastIndex int32 = -1
	for i, v := range arr {
		if (v.element <0){
			count ++
			lastIndex = int32(i)
			if (index == -1){
				index = int32(i)
			}
		}
	}
	return count, index, lastIndex
}