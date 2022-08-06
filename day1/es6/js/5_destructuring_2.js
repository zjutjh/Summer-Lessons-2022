// 函数参数解构

const fruit = {
  apple: true,
  peach: false,
  grape: true,
};

const animal = ["pdx", "sd"];

function add([{ apple, grape }, visitor]) {
  console.log(apple, grape);
  console.log(visitor[0]);

  /* if (apple) console.log("apple is good");
  else console.log("apple is bad"); */

  // 模板字符串
  // console.log(apple ? "apple is good" : "apple is bad");
  // console.log("apple is " + apple ? "good" : "bad"); // wrong
  // console.log(`apple is ${apple ? "good" : "bad"}`);
}

add([fruit, animal]);
