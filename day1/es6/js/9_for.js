const arr = [12, 1, 5, 2, 1];

for (let i = 0; i < arr.length; i++) {
  console.log(arr[i]);
  // if (i > 2) break;
}

arr.forEach((item, index) => {
  item++;
  console.log(index, item);
});
console.log(arr);

const arr2 = arr.map((item) => item + 1);
console.log(arr2);

const objArr = [{ a: 1 }, { a: 2 }];
objArr.forEach((item) => (item["b"] = 2));
console.log(objArr);
