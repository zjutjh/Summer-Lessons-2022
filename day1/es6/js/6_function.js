function foo(x = 1, y) {
  console.log(x, y);
}

foo(1, 2); // 1 2
foo(2, 2); // 2 2
foo(2); // 2 undefined
foo((x = 2)); // 2 undefined
foo((y = 2)); // 2 undefined
