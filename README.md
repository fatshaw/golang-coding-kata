# golang-coding-kata

1. bulls and cows
leetcode: https://leetcode.com/problems/bulls-and-cows/description/

2. re  
一个简单的函数编程练习
* literal(“abc”），匹配abc的的字符串
* sequence(literal(“abc”),literal(“def”)),匹配“abcdef”
* alterative(literal(“abc”),literal(“bcd”)),匹配abc或bcd
* 支持*和+的量词，贪心匹配；（not implemented）
* 支持any，表示匹配任意字符
* 支持oneof(“abc”),表示匹配其中任意字符