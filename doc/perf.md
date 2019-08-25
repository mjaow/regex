## DFA performance

- DFA:case [a(b|c)kk:abkk] is ok and cost 617 ns
- DFA:case [abcdefg:abcdefg] is ok and cost 812 ns
- DFA:case [a(b|c)*:abbbbbbbbbbbbbbbbbbb] is ok and cost 1184 ns
- DFA:case [(a|b)*a:ababababab] is ok and cost 812 ns
- DFA:case [(a|b)*a:aaaaaaaaba] is ok and cost 756 ns
- DFA:case [(a|b)*a:aaaaaabac] is ok and cost 705 ns
- DFA:case [a(b|c)*d:abccbcccd] is ok and cost 781 ns
- DFA:case [a(b|c)*d:abccbcccde] is ok and cost 856 ns
- DFA:case [a(b|c)+d:acd] is ok and cost 434 ns
- DFA:case [a(b|c)+d:ad] is ok and cost 343 ns
- DFA:case [a(b|c)+d:abbbbd] is ok and cost 548 ns
- DFA:case [a(b|c)?d:acd] is ok and cost 409 ns
- DFA:case [a(b|c)?d:accd] is ok and cost 359 ns
- DFA:case [a(b|c)?d:ad] is ok and cost 238 ns
- DFA:case [(a*)(a*)(a*)(a*)(a*):] is ok and cost 72 ns
- DFA:case [(a*)(a*)(a*)(a*)(a*):a] is ok and cost 187 ns
- DFA:case [(a*)(a*)(a*)(a*)(a*):aaaaaaa] is ok and cost 534 ns
- DFA:case [(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*):aba] is ok and cost 635 ns
- DFA:case [:] is ok and cost 75 ns
- DFA:case [    :] is ok and cost 68 ns
- DFA:case [    :    ] is ok and cost 610 ns

## NFA performance

- NFA:case [a(b|c)kk:abkk] is ok and cost 940 ns
- NFA:case [abcdefg:abcdefg] is ok and cost 641 ns
- NFA:case [a(b|c)*:abbbbbbbbbbbbbbbbbbb] is ok and cost 5872 ns
- NFA:case [(a|b)*a:ababababab] is ok and cost 3510 ns
- NFA:case [(a|b)*a:aaaaaaaaba] is ok and cost 2933 ns
- NFA:case [(a|b)*a:aaaaaabac] is ok and cost 2508 ns
- NFA:case [a(b|c)*d:abccbcccd] is ok and cost 2298 ns
- NFA:case [a(b|c)*d:abccbcccde] is ok and cost 2281 ns
- NFA:case [a(b|c)+d:acd] is ok and cost 633 ns
- NFA:case [a(b|c)+d:ad] is ok and cost 311 ns
- NFA:case [a(b|c)+d:abbbbd] is ok and cost 1466 ns
- NFA:case [a(b|c)?d:acd] is ok and cost 583 ns
- NFA:case [a(b|c)?d:accd] is ok and cost 511 ns
- NFA:case [a(b|c)?d:ad] is ok and cost 492 ns
- NFA:case [(a*)(a*)(a*)(a*)(a*):] is ok and cost 825 ns
- NFA:case [(a*)(a*)(a*)(a*)(a*):a] is ok and cost 2360 ns
- NFA:case [(a*)(a*)(a*)(a*)(a*):aaaaaaa] is ok and cost 98717 ns
- NFA:case [(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*):aba] is ok and cost 1168793 ns
- NFA:case [:] is ok and cost 79 ns
- NFA:case [    :] is ok and cost 205 ns
- NFA:case [    :    ] is ok and cost 541 ns
