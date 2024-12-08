// --- Directions
// Given a string, return a new string with the reversed
// order of characters
// --- Examples
//   reverse('apple') === 'leppa'
//   reverse('hello') === 'olleh'
//   reverse('Greetings!') === '!sgniteerG'

// First solution, Using array method
// const reverse = (str) => str.split("").reverse().join("");

// Second solution will be using For loop to concatenate the string 
function reverse(str) {
  let result = ""
  for (let i = str.length - 1; i >= 0; i--) {
    result += str[i]
    debugger;
  }
  return result
};

reverse("abcd")

module.exports = reverse;
