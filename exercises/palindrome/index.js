// --- Directions
// Given a string, return true if the string is a palindrome
// or false if it is not.  Palindromes are strings that
// form the same word if it is reversed. *Do* include spaces
// and punctuation in determining if the string is a palindrome.
// --- Examples:
//   palindrome("abba") === true
//   palindrome("abcdefg") === false

// My solution is using two pointer
function palindrome(str) {
  let firstPtr = 0
  let secondPtr = str.length - 1
  while (firstPtr < secondPtr) {
    if (str[firstPtr] != str[secondPtr]) {
      return false
    }
    firstPtr++
    secondPtr--
  }
  return true
}

// Second is using array method 
// function palindrome(str) {
// const reversed = str.split("").reverse().join("")
// return str === reversed
// }

module.exports = palindrome;
