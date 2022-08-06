// 数组解构
let [a, b, c] = [1, 2, "cxnb"];
// let [, b, c] = [1, 2, "cxnb"];
// let [a, b] = [1, 2, "cxnb"];

console.log(a, b, c);

/*
let a = 1;
let b = 2;
let c = "cxnb"; 
*/

// 对象 (Object) 解构
let person = {
  name: "cx",
  sex: {
    child: "male",
    adult: "unknow",
  },
};

const person_sex1 = person.sex.adult;
console.log(person_sex1);

const { sex } = person; // 变量名要和属性名相同
// const { sex: sex } = person;
// const { sex } = person || {};

const { adult: person_sex2 } = sex; // 自定义变量名
console.log(person_sex2);
