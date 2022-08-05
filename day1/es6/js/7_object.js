let obj;
const x = 1;
const y = 2;

obj = { x, y };
//obj = {x: x, y: y}
console.log(obj);

obj = {
  foo1: function () {
    return "hello1";
  },

  foo2() {
    return "hello2";
  },
};
console.log(obj.foo1());

obj.foo3 = function () {
  return "hello3";
};
console.log(obj.foo3());

obj["f" + "oo4"] = function () {
  return "hello4";
};
console.log(obj.foo4());
