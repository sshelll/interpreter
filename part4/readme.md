# Part 4

基于 Part 3 中的计算器，实现以下功能

- 支持处理乘除法
- 移除对加减法的支持
- 重构代码实现 Lexer

> Example: 
>
> 2 * 6、490 / 4 / 2

虽然从功能层面看起来只是“换皮”，但重点在于使用下面的概念重构解释器。

## Grammar

>  为表达式 `7*4/2*3`定义如下语法规则：

``` text
expr = factor((MUL|DIV)factor)*
factor = INTEGER

每条规则左侧🫲为「非终止符」，右侧🫱为一系列「终止符」和「非终止符」组成的规则
上述语法中 MUL、DIV、INTEGER 为「终止符」，expr 和 factor 为「非终止符」
```

> 基于上述规则的一些简单推导过程：

``` text
# 3
expr = factor((MUL|DIV)factor)*
     = factor
     = INTEGER

# 3*7
expr = factor((MUL|DIV)factor)*
     = factor MUL factor
     = INTEGER MUL INTEGER
     
# 3*7/2
expr = factor((MUL|DIV)factor)*
     = factor MUL factor((MUL|DIV)factor)*
     = factor MUL factor DIV factor
     = INTEGER MUL INTEGER DIV INTEGER
```

语法到代码的转化规则：

1. 对于任意一条规则`R`，可以转化为一个同名函数`R()`。函数体亦遵循相同的原则
2. 或运算`(a1 | a2 | aN)` 翻译为`if-elif-else` 语句
3. 可选分组`(…)*`转化为`while`语句
4. 每个`Token(T)`，转化为`eat` 方法: `eat(T)`。如果`eat()` 与当前token 匹配，则从词法分析器获取下一个token 赋值给current_token 变量

## Summary

part 4 将 part 3 的代码做了一个重构，将单一的 `Interpreter` 拆分为 `Interpreter` + `Lexer`。

`Lexer`：用于处理 Token，主要暴露的方法是 `GetNextToken() *Token` ，从输入流中识别下一个 Token 的具体类型。

`Interpreter`：用于处理语法规则，根据不同的规则实现不同的函数，如本例当中实现了 `factor()` + `expr()` 两个方法来校验规则。（在校验的同时还可以进行计算）