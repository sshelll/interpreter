# Part 5

完善 Part 4 中的计算器，实现以下功能

- 支持处理四则运算

> Example: 
>
> 100 + 2 * 3、12 / 4 + 490 - 29 

## 定义运算符的优先级

| priority | op   | associativity |
| -------- | ---- | ------------- |
| 2        | * /  | left          |
| 1        | + -  | left          |

## 构造上下文无关语法

```text
expr   = term((PLUS|MINUS)term)*
term   = factor((MUL|DIV)factor)*
factor = INTEGER
```

## 构造语法图

```mermaid
graph LR
expr
S((start)) --> T1[term]
T1 --------> E((end))
T1 --> OP[op]
OP --> + --> T2[term]
OP --> - --> T2
T2 -.key path..-> OP
T2 --> E
```

```mermaid
graph LR
term
S((start)) --> T1[factor]
T1 --------> E((end))
T1 --> OP[op]
OP --> * --> T2[factor]
OP --> / --> T2
T2 -.key path..-> OP
T2 --> E
```

```mermaid
graph LR
factor
S((start)) --> I[INTEGER] --> E((end))
```

## 实现

参照 part4 中的代码转化规则。
