// const readline = require('readline').createInterface({
//   input: process.stdin,
//   output: process.stdout,
// });
// readline.question('Enter a name: ', (name) => {
//   const reversedName = name.split('').reverse().join('');
//   if (name === reversedName) {
//     console.log(`${name} is a palindrome`);
//   } else {
//     console.log(`${name} is not a palindrome`);
//   }
//   readline.close();
// });


// function fastestIsPalindrome(str) {
//   var len = Math.floor(str.length/2);
  
//   for (var i = 0; i < len; i++)
//     if (str[i] !== str[str.length - i - 1])
//       return false;
//   return true;
// }
// console.log(fastestIsPalindrome('chico'));



// const list1 = [3,4,6,3,2,1];
// const list2 = [1,4,5,6];

// // // sort lists removing duplicates and sorting them
// // const sortedList1 = list1.sort().filter((item, index, array) => {
// //   return array.indexOf(item) === index;
// // });
// // const sortedList2 = list2.sort().filter((item, index, array) => {
// //   return array.indexOf(item) === index;
// // }
// // );
// // console.log(sortedList1,sortedList2)



// // merge sorted lists and remove duplicates
// const mergedList = list1.concat(list2).filter((item, index, array) => {
//   return array.indexOf(item) === index;
// }).sort();
// // reduce array merged 
// const newlist = mergedList.reduce((a, b) =>  a + b, 0);
// console.log(newlist)
// console.log(mergedList);

