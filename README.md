# Caster - WIP

Caster is a language with contemporary syntax and transpiler to generate Verilog.

TODOs:

* Better name
* Make it usable
* Document at least TODOs

Example code is like this.

    module m(p input, r output) {
        stage s1 {
            r <= p
            goto s1
        }
    }

This uses participle to parse the input.
https://github.com/alecthomas/participle
