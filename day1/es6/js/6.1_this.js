const obj = {
  a: 1,
  f1() {
    console.log("hello");
    console.log(this);
  },

  // 箭头函数
  f2: () => {
    console.log("hello");
    console.log(this);
  },
};

obj.f1();
obj.f2();
