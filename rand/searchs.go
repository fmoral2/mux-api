
//binary search tree ( BST ) is a binary tree where you need to have at least 2 children node for each ( except last ) and all left values are less
//and right values are greater
// [2,3,5,7,9,10,12]  target = 9
// find avg index by getting first and last index, then compare if the value is greater then the index middle value,
// if it is cut the left part of array if it is less cut the right part , find average again until find the value in the middle
//value is the target.
// value of i is increasing exponentially, then when i is out of bound in the array, apply binary search (x, min(i/2) , high(1,n-1) )
//usefull for a very large array or infinite
// func exponentialSearch(arr []int, x int) int {
// }

//the array should be sorted and uniformly distributed
//   position =  low (first value) +  x-a(low) / a[h] - a[l]  x  (h-l)
// x= 9
// [3 , 5 ,7, 9, 11,13] ------   low = 0 , h= 5,  a[l] = 3 , a[h] = 13
//   pos = 0 + (9-1) / (13-1) x 5-0
// 8/12 * 5 = 3.333333
// func interpolationSearch(arr []int, x int) int {
// }

// initialize the array in array[0]= 0  and array[1]= 1
// i comeca no 2 ,   i = 2 ; i <= n ; i++
// array[i] = array[i-1] + array[i-2]
// func fibonaccinumbers(n int){}

//https://github.com/graphoarty/python-dsa/tree/master/Algorithms/Search
// func fibonacciSearch(arr []int, x int) int {
// }

// use a jump range to go find the value and in the last block just do a linear search
// func jumpSearch(arr []int, x int) int {
// }

// compare with all elements starting from the left
// func linearSearch(arr []int, x int) int {
// }

//find min element index , start applying low and high pointers and divide by 2 ,value of  index pivot must be  < then pivot - , return the that min index.
//  if pivot > low && > high move the low to pivot +1 , else move to pivot-1

// func modifiedBinarySearch(arr []int, x int) int {
// }

// divide array by 3 , finding first midpoint then second, check if the value is in which portion e find it
// func ternarySearch(arr []int, x int) int {
// }

