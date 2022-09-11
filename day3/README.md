# Summer-Lessons-2022

2022年暑期授课代码样例

## Day3
Vue.js 前端开发实战之课程表的实现(二)

[@FinleyGe](https://github.com/FinleyGe)

### TODO list:
- 实现apis
- 登录/注册表单
  - 实现组件 k-form, k-input
- 增删改查课程表
  - 课程表数据结构修改
  - 课程表组件修改
  - 实现组件 k-time-selector, k-select

## Vue

### v-if v.s v-show
V-if false -> 不会渲染 DOM
V-show false -> display:none;


## V-model:
V-model 双向绑定, 语法糖.
v-model :
value, @input

## JS Sort 字典序
1,10,11,12,13,14,15,16,2,3,4,5,6,7...

## number & Number
Number -> NumberConstructor

## `...` 拓展运算符, 解包

## `!` `?`
`!` != , 取反, 
TS: 强制类型断言
`?.` ES6 [可选链](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Operators/Optional_chaining) 

## API

API Interface
Apifox 软件 格式化 API文档 测试API, mock

**Axios**

> Ajax: Asynchronous JavaScript and XML 
> 实际上就是与后端的通信方式.
> 底层实现上, Axios 和 ajax 略有不同. 

**RESTful API**
Representational state transfer 表征状态转移

是一种API设计规范.

参考:
https://www.ruanyifeng.com/blog/2014/05/restful_api.html

Method: 
- GET
- POST
- PUT(传入所有数据)
- PATCH(传入部分数据)
- DELETE

## Template Setup

- props -> defineprops;
- emits -> defineemits;
- no return

# 其他

## 学习
程序员: keep learning
前端: 

`查`

`问`

学到了东西 -> 做 (产出)

Blog 总结

# 大作业
- 两个人组队: 建议一个前端, 一个后端. 也可以不组队.
- 日历: 
  - 基本要求: curd
  - 结合功能实现, 代码规范性等综合评分.
