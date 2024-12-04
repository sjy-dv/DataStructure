package datastructure_test

import (
	"fmt"
	"testing"
)

func bubbleSort(list []int) []int {

	for i := len(list) - 1; i >= 0; i-- {
		for j := range i {
			if list[j] > list[j+1] {
				temp := list[j]
				list[j] = list[j+1]
				list[j+1] = temp
			}
		}
	}
	return list
}

func selectionSort(list []int) []int {
	for i := range len(list) - 1 {
		selectIndex := i
		for j := selectIndex; j < len(list); j++ {
			if list[j] < list[selectIndex] {
				selectIndex = j
			}
		}
		if i != selectIndex {
			temp := list[i]
			list[i] = list[selectIndex]
			list[selectIndex] = temp
		}
	}
	return list
}

func insertionSort(list []int) []int {
	// if list 0 or 1 not working also not needs to sort
	for i := 1; i < len(list); i++ {
		temp := list[i]
		j := i - 1
		//[5 3 2 1 4]
		// temp = 3
		// list[0] = 3
		// list[1] = 5 <- ok

		// if i = 3
		// temp = 1
		// iter list[2] = 2 (swap j--)
		// iter list[1] = 3 (swap j--)
		// iter list[0] = 5 (swap j--)
		// list [1 5 3 2 4]
		// j is index pos, not smaller than zero
		for j > -1 && temp < list[j] {
			list[j+1] = list[j]
			list[j] = temp
			j--
		}
	}
	return list
}

func merge(listA, listB []int) []int {
	combined := []int{}
	i := 0
	j := 0
	for i < len(listA) && j < len(listB) {
		if listA[i] < listB[j] {
			combined = append(combined, listA[i])
			i++
		} else {
			combined = append(combined, listB[j])
			j++
		}
	}
	for i < len(listA) {
		combined = append(combined, listA[i])
		i++
	}
	for j < len(listB) {
		combined = append(combined, listB[j])
		j++
	}
	return combined
}

func mergeSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	middleIndex := int(len(list) / 2)
	left := mergeSort(list[:middleIndex])
	right := mergeSort(list[middleIndex:])
	return merge(left, right)
}

func pivot(list []int, pIndex, eIndex int) int {
	swapIndex := pIndex
	for i := pIndex + 1; i < eIndex+1; i++ {
		if list[i] < list[pIndex] {
			swapIndex += 1
			swap(list, swapIndex, i)
		}
	}
	swap(list, pIndex, swapIndex)
	return swapIndex
}

func quickSort(list []int) []int {
	return quickSortHelper(list, 0, len(list)-1)
}

func quickSortHelper(list []int, left, right int) []int {
	if left < right {
		pivotIndex := pivot(list, left, right)
		quickSortHelper(list, left, pivotIndex-1)
		quickSortHelper(list, pivotIndex+1, right)
	}
	return list
}

func swap(list []int, a, b int) {
	list[a], list[b] = list[b], list[a]
}

func TestBubbleSor(t *testing.T) {
	testList := []int{15, 74, 241, 643, 1, 3}
	fmt.Println(bubbleSort(testList))
}

func TestSelectionSort(t *testing.T) {
	testList := []int{15, 74, 241, 643, 1, 3}
	fmt.Println(selectionSort(testList))
}

func TestInsertionSort(t *testing.T) {
	testList := []int{15, 74, 241, 643, 1, 3}
	fmt.Println(insertionSort(testList))
}

func TestMergeSort(t *testing.T) {
	testList := []int{15, 74, 241, 643, 1, 3}
	fmt.Println(mergeSort(testList))
}

func TestQuickSort(t *testing.T) {
	testList := []int{15, 74, 241, 643, 1, 3}
	fmt.Println(quickSort(testList))
}
