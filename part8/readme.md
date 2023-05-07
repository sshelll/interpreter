# Part 8

基于 Part 7 中的计算器，实现以下功能：

> 为 ‘+’ / ‘-’ 支持单目运算



## 单目运算

之前的 `+` / `-` 代表的是「相加」和「相减」，是双目运算符，也就是说这个符号左右都必须有值参与运算。

这里我们将为这两个符号支持单目运算，也就是只需要有右值即可。当然有无左值会影响该符号的含义，具体来说，在无左值的情况下：

- `+` 不改变操作数符号
-  `-` 将对其操作数取负值

也就是在原来的「加减法」的基础上，拓展了「正负数」的概念。

举例如下：

```text
expr = +-3
     = +(-(3))
     = 3

expr = 5---2
     = 5-(-(-(2)))
     = 5-(-(-2))
     = 5-(2)
     = 3
```

<u>***需要注意的是，单目运算符的优先级，比双目运算符要高。***</u>



## Grammar

> 由于单目运算符优先级更高，所以我们要从 factor 规则下手。

基于上述描述，更新规则如下：

```text
Original:
expr   = term((PLUS|MINUS)term)*
term   = factor((MUL|DIV)factor)*
factor = INTEGER|(LPAREN expr RPAREN)

Current:
expr   = term((PLUS|MINUS)term)*
term   = factor((MUL|DIV)factor)*
factor = (PLUS|MINUS)factor|INTEGER|(LPAREN expr RPAREN)
```



## 实现

我们在 part 7 中的 ```astNode``` 实现类中，新增一个单目运算符节点实现类：

```go
type astUnaryOp struct {
	Op    *Token
	Right astNode // 与双目运算符相比，少了一个左值
}
```

对应的需要修改 `parser` 中 `factor()` 函数以及在 `astVisitor` 中新增 `visitUnaryOp()` 函数。