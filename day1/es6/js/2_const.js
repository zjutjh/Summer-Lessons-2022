const a = 1;
let b = 1;

// a = 1;
b = 2;

// 对象，数组，字符串 建议使用 const 声明
const obj = {
  number: 1,
  str: "",
  arr: [],
  subObj: {},
};
obj.number = 2;
console.log(obj);

const arr = [1, 2, 4];
arr[2] = 2;
console.log(arr);

const str = "hello";
console.log(str[0]);
str[0] = "2";
console.log(str);
