package main

// v = len(n)/ 2
// start compare first element x last element on the chain  v = len(n)/ 2 , if i > j swap,
//then divide again v = len(n)/ 2 , and continue until be 1 "space" then compare one by one and if swap then compare
// with the left value if it is bigger does nothing, if it is not swap,
// func ShellSort() []int     {}

//compara 2 elements if i > j , swap and continue iteration until not find i > j
// func BubbleSort() []int    {}

// use the current minimun (smallest) and first element,  with current item and when current item <  smallest , smallest = item , swap them and continue
// func SelectionSort() []int {}

//look each and compare if is > to  its left until not one is bigger , if it is swap , if its not its sorted
// func InsertionSort() []int {}

//recursively , divide and conquer , divide array / 2 until there is only individual itens then merge back sorting until have one array
// func MergeSort() []int     {}

// recursive , choose i pivot (after sorted -> correct position , items to the left are smaller and to the right are larger) and partition the array
// find partition position por pivot -  increment i until i <= pivot then decrement j until > pivot then swap then if i < j and continue next iteration until all elements
// to the left are smaller then pivot and to the right are bigger , then recursively quick sort the 2 remaining lists
// func QuickSort() []int     {}

// heap = ordered binary tree , max heap = parent node > child , heapify=  build max heap when one or more elements are already sorted
//  find the largest item ( max heap ) and then swap if the last element , and then remove from the tree and consider it sorted
// call heapify again recursively
// func HeapSort() []int      {}

// when you have short range of numbers and repeating numbers , create a new array with the total max number that appears and then fill each index
// example [3,4,7,4,1]
//    [1,0,1,2,0,0,1] max = 7
// first count how much each number index occurs , then sum index value with predecessor value , [1,0,1,2,0,0,1] = 1,1,2,4,4,4,5 then
// create anothter empty array using initial valueswith index with previos array then find the index , decrement index array value
// [1 ,3, 4, ,4 ,7]

// func CountingSort() []int  {}
// find max value of the buckets ( value ) [3,2,4,7,6,5] = would be 7  , separete each value to an appropriate bucket

//  1 , 2 , 3,  4  ,5 ,6  ,7
//      2   3  4  5   6   7
// put every back
// func BucketSort() []int    {}
//

//find the len of array and create a bucket with that  [ 3,4,1,4,30,8] = len 6
// find the biggest number and use as default the number of digits in this example 30 = 2 , make all other have 2 as well
// 03, 04, 01, 04, 30 ,08
// find the last number and fill the buckets like that
// 0   1   2   3   4  5
// 30  01      03  04
//                 04
// take them back to the array in order like that
// [30,01,03,04,04]
// now sort using 2 digit from right
// 01       30
// 03
//04
//04
// now put back
// [01,03,04,04,30]
// func RadixSort() []int     {}

// < 64 elements -> binary insert sort then merge , if more than 64 use galloping
// func TimSort() []int       {}

// func main() {
// 	BubbleSort()
// }
