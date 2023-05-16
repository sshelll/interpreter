# Part 9

基于 Part 8 中的计算器，实现简单的 Pascal 语言解释器。



## Pascal 语法

``` pascal
BEGIN
    BEGIN
        number := 2;
        a := number;
        b := 10 * a + 10 * number / 4;
        c := a - - b
    END;
    x := 11;
END.
```

1. **一段 Pascal 代码（program），是由一个`复合语句`后紧跟一个`.`组成的：**

``` pascal
BEGIN {复合语句} END.
```

因此这里我们抽象一个 Grammar：

``` text
program = compound_statement DOT
```

2. **复合语句（compound_statement）是一个由 `BEGIN END` 包裹而成的代码块，内部有 0 或多条 Pascal 语句，并且在复合语句内部，除了最后一条语句，其余的语句都应该以`;`结尾：**

```pascal
// 内部包含 0 条语句
BEGIN END

// 内部包含两条赋值语句，第一条以 ‘;’ 结尾
BEGIN a := 5; x := 11 END

// 内部包含两条赋值语句，两条都以 ‘;’ 结尾
BEGIN a := 5; x := 11; END

// 内部嵌套了‘复合语句’的复合语句
BEGIN 
    BEGIN 
        a := 5 
    END; 
    x := 11 
END
```

3. **语句列表（statement_list），也就是复合语句内部的语句（0 或多条）。**

4. **语句（statement），可以是复合语句，也可以是赋值语句，也可以是空语句。**

5. **赋值语句（assignment_statement），一个变量，跟随 `:=`，再跟随一个 `expr` 表达式：**

   ```pascal
   a := 11
   b := a + 9 - 5 * 2
   ```

6. **变量（variable），就是一个标识符，如上面的 `a` 和 `b`。**

7. **空语句（no_op_statement），不包含任何内容。用于语句列表的末尾，或者空的复合语句`BEGIN END`。**

8. **之前章节的 `expr` 、`term`、`factor` 等也需要进行更新来适应新的语法规则（处理变量）。**



综上我们可以总结出完整的语法如下：

```text
Grammars:
program              = compound_statement DOT
compound_statement   = BEGIN statement_list END
statement_list       = statement | statement SEMI statement_list
statment             = compound_statement | assignment_statement | empty
assignment_statement = variable ASSIGN expr
empty                = 
expr                 = term ((PLUS | MINUS) term)*
term                 = factor ((MUL | DIV) factor)*
factor               = PLUS factor |
											 MINUS factor |
                       INTEGER |
                       LPAREN expr RPAREN |
                       variable
variable             = ID
```



## 实现

### Lexer

主要有以下四个改动点：

1. 新增 Token 类型
2. 新增一个 `peek` 方法 —— 用于在不改变 `pos` 的情况下，预读取后续的字符处理可能的 token，目前是 `:=`
3. 新增一个 `ID` 方法——用于处理变量和保留字（`BEGIN`、`END`）
4. 重构 `getNextToken` 方法

### Parser

主要有以下三个改动点：

1. 新增 AST 节点，参照上述的上下文无关语法
2. 新增对应的解析方法，例如 `program` , `compound_statement` 等
3. 重构对应的 `parse` 和 `factor` 方法

### Interpreter

主要有以下两个改动点：

1. 新增对应的 visit 方法
2. 新增一个 `Global_Scope` 在结构体中，用于保存变量的值



## 总结

本节做的主要改动如下：

1. 添加新的语法规则
2. 添加新的 token，与其对应的词法分析方法，更新`get_next_token`方法
3. 添加新的抽象语法树节点
4. 为新的 AST 节点添加对应的语法分析方法
5. 为解释器添加新的 visit 方法，以遍历新的抽象语法树
6. 使用字典来作为符号表保存变量