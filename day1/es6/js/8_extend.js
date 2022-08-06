const a = [1, 2, 3];
const b = [3, 4, 5];
const c = a.concat(b);

console.log(c);

const d = [...a, ...b];
console.log(d);

const e = new Set([...a, ...b]);
console.log(e);
console.log([...e]);

const obj1 = {
  a: 3,
  c: 5,
};
const obj2 = {
  a: 2,
  b: 2,
};
const obj3 = { ...obj1, ...obj2 };
console.log(obj3);
