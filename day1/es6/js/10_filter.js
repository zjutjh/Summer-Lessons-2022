const arr = [1, 2, 3, 4, 5, 6];

const after = arr.filter((item) => {
  if (item % 2 === 1) return true;
  // return item % 2 === 1;
});

// const after = arr.filter((item) => item % 2 === 1);

console.log(after);
