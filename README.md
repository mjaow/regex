# regex

自己写的正则表达式引擎

- [参照Russ Cox的paper实现的thompson NFA正则引擎](https://github.com/mjaow/regex/blob/master/src/regex/nfa.go) 
- [基于子集构造法实现的DFA正则引擎](https://github.com/mjaow/regex/blob/master/src/regex/dfa.go) 

NFA和DFA性能对比

- [NFA vs DFA](https://github.com/mjaow/regex/blob/master/doc/perf.md)

收集各种版本正则表达式的实现

Russ Cox

- [Regular Expression Matching Can Be Simple And Fast](https://swtch.com/~rsc/regexp/regexp1.html)
- [Regular Expression Matching: the Virtual Machine Approach](https://swtch.com/~rsc/regexp/regexp2.html)

Denis Kyashif

- [Implementing a Regular Expression Engine](https://deniskyashif.com/implementing-a-regular-expression-engine)

Ken Thompson

- [Regular Expression Search Algorithm](https://www.fing.edu.uy/inco/cursos/intropln/material/p419-thompson.pdf)

Rob Pike

- [A Regular Expression Matcher](http://www.cs.princeton.edu/courses/archive/spr09/cos333/beautiful.html)
