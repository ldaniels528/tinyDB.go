Tiny VM (Go variant)
======================

The purpose of this project is evaluate the developer experience offered by
Rust and Go by implementing the same small virtual machine in both languages.

You can find the [Rust variant here](https://github.com/ldaniels528/tiny_vm).

I'm just getting started but check back from time to time to see the progress.

#### Test input

```textmate
this\n is\n 1 \n\"way of the world\"\n - `x` + '%'
```

#### Test output

```textmate
=== RUN   TestMapF
2024/02/25 15:14:31 items: ["100","200","300"]
2024/02/25 15:14:31 newItems: [100,200,300]
--- PASS: TestMapF (0.00s)
=== RUN   TestToByte
--- PASS: TestToByte (0.00s)
=== RUN   TestParse
2024/02/25 15:14:31 token: {"text":"this","type":0,"start":0,"end":4,"lineNumber":1,"columnNumber":1}
2024/02/25 15:14:31 token: {"text":"is","type":0,"start":6,"end":8,"lineNumber":2,"columnNumber":4}
2024/02/25 15:14:31 token: {"text":"1","type":3,"start":10,"end":11,"lineNumber":3,"columnNumber":7}
2024/02/25 15:14:31 token: {"text":"\"way of the world\"","type":2,"start":13,"end":31,"lineNumber":4,"columnNumber":9}
2024/02/25 15:14:31 token: {"text":"-","type":4,"start":33,"end":34,"lineNumber":5,"columnNumber":28}
2024/02/25 15:14:31 token: {"text":"`x`","type":1,"start":35,"end":38,"lineNumber":5,"columnNumber":30}
2024/02/25 15:14:31 token: {"text":"+","type":4,"start":39,"end":40,"lineNumber":5,"columnNumber":34}
2024/02/25 15:14:31 token: {"text":"'%'","type":5,"start":41,"end":44,"lineNumber":5,"columnNumber":36}
--- PASS: TestParse (0.00s)
=== RUN   TestParseFully
2024/02/25 15:14:31 token[0]: |this| {"text":"this","type":0,"start":0,"end":4,"lineNumber":1,"columnNumber":1}
2024/02/25 15:14:31 token[1]: |is| {"text":"is","type":0,"start":6,"end":8,"lineNumber":2,"columnNumber":4}
2024/02/25 15:14:31 token[2]: |1| {"text":"1","type":3,"start":10,"end":11,"lineNumber":3,"columnNumber":7}
2024/02/25 15:14:31 token[3]: |way of the world| {"text":"\"way of the world\"","type":2,"start":13,"end":31,"lineNumber":4,"columnNumber":9}
2024/02/25 15:14:31 token[4]: |-| {"text":"-","type":4,"start":33,"end":34,"lineNumber":5,"columnNumber":28}
2024/02/25 15:14:31 token[5]: |x| {"text":"`x`","type":1,"start":35,"end":38,"lineNumber":5,"columnNumber":30}
2024/02/25 15:14:31 token[6]: |+| {"text":"+","type":4,"start":39,"end":40,"lineNumber":5,"columnNumber":34}
2024/02/25 15:14:31 token[7]: |%| {"text":"'%'","type":5,"start":41,"end":44,"lineNumber":5,"columnNumber":36}
--- PASS: TestParseFully (0.00s)
PASS
```