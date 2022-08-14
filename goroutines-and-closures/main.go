// A Closure is a function taken together with an environment.
// The function is typically an anonymous function defined inside
// another function. The environment is the lexical scope of the
// enclosing function (very basic idea of a lexical scope of a
// function would be the scope that exists between the function's braces.)

// func g() {
//     i := 0
//     f := func() { // anonymous function
//         fmt.Println("f called")
//     }
// }

// within the body of function f in above case, the variables of both f anf g are accessible.
// However it is the scope of g that forms the environment part of the closure
package main

import "fmt"

func NaturalNumbers() func() int {
	i := 0
	f := func() int { // f is the function part of closure
		i++
		return i
	}
	return f
}

func main() {
	n := NaturalNumbers() // here n stores the name of function (which is f)
	fmt.Println(n())      // 1
	fmt.Println(n())      // 2
	fmt.Println(n())      // 3

	m := NaturalNumbers()
	fmt.Println(m()) // 1
	fmt.Println(n()) // 4
	fmt.Println(m()) // 2

}

// In above definition, NaturalNumbers has an inner function f which NaturalNumbers returns.
// Inside f, variable i defined within the scope of NaturalNumbers is being accessed.

// Now n is a closure. It is a function (defined by f) which also has an associated
// environment (scope of NaturalNumbers).
// In case of n, the environment part only contains one variable: i

// As evident from above output, each time n is called, it increments i.
// i starts at 0, and each call to n executes i++.

// The value of i is retained between calls. That is, the environment,
// being a part of closure, persists.

// Calling NaturalNumbers again would create and return a new function.
// This would initialize a new i within NaturalNumbers.
// Which means that the newly returned function forms another closure having the
// same part for function (still f) but a brand new environment (a newly initialized i).

// NOTE: Usually when the function returns, the local variables goes
//       out of scope. so what happens in case of closures, the outer
// 	  function has returned, but the runtime is clever enough to
// 	  see that the reference to the local variable i is still being
// 	  held by the inner function, it moves it from stack to heap,
// 	  so that inner function still has access to variable even after
// 	  the enclosing function returns.
