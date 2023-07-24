# Declare a variable
var name type = expression 
# Declare three ints 
var i, j, k int 
# Declare three different types 
var b, f, s = true, 4.5, "four"
# Declare vars from function call returning multiple vars 
var f, err = os.Open(file)
# Short variable declaration: only used within functions 
# var declaration is used if we want a type that's different from initializer 
# short variable declaration must declare at least a new variable, and it "assigns" to already existing vars 
name := expression 

# In contrast to stack varibles in language like C/C++, a local variable returned from a function is retained even after the function call!!
# But the returned value is distinc for each function call 
# This is valid:
var p = f()
func f() *int () {
    v := 1 
    return &v
}

# new returns a unnamed variable of the given type
# Note new is a predeclared function, not a keyword and thus can be overridden in a scope 
p := new(int) 

# Variable memory allocation in determined by the compiler, either on stack or in heap, depending on if the variables are still 
# reachable/unreachable after the function returns. If a variable is still reachable, we say the variable escape from the function 
# escape variables have an extra memory allocation, which impacts performance. 

# Tuple assignment: allows multiple variables to be assigned at once. Note all right-hand variables are evaluated first before they are assigned to the left variables 
func fib( n int ) int {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        x, y = y, x+y
    }
    return x
}

# type keyword is similar to typedef in some languages
# two values of different named types can not be compared even though their underlying types are the same 

# To export an identifer to outside the package, name it with upper-case letter 

# package initialization function is used to do some initialization at package level by defining a init function, which is similar to other defined functions except it can not be called 

# Go offers four data types 
# 1. basic types: numbers, strings, booleans 
# 2. aggregate types: arrays, structs 
# 3. reference types: pointers, slices, maps, functions, channels 
# 4. interface type 


# strings in Go are IMMUTABLE
# raw string literals are defined use backquotes instead of double quotes 
# byte slices are mmutable and strings can be converted to byte slices 

# the Go terminology 'rune' means the Unicode code point, it's int32 as Go use int32 to encode unicode code points

# when declaring groups of consts, the values can be omitted all but for the first 
const (
    a = 1
    b 
    c = 2
    d
)
declares: a = 1, b = 1, c = 2, d = 2
# The above doesn't support enum in c/C++, where the following item has incremented values 
# Go provides const generator itoa to provide this utility, i.e., it defines the relationship can be a multipe of the previous item 

# Go provides fixed-sized containers (array) and variable-lenght containers (slices)
# [25]int   an array of 25 ints
# []string  a slice of strings 
# arrays are comparable using "==" or "!=", however, slices are generally not comparable (except compare against nil), which means we need to do comparison ourselves 
# exception is that the standard library provides comparison for byte slices
# slice len being zero doesn't imply it is nil. 
# s = []int{} has len zero, but it is NOT nil 
# the make function can be used to create a slice 
# make([]T, len) or make([]T, len, cap) where cap is equal to len 

# In Go, map is a reference to a hashtable, and is written as map[k]v, where k and v are the types of keys and values 
# we can use make to create a map, e.g., ages := make(map[string]int)
# To differentiate the cases where a key don't exist from a key assuming default zero value (depending on the value type), 
# we could use a return status by indexing the key 
# e.g., if age, ok := ages["nonexistkey"]; !ok { }, handling not existing case 

# when declaring a struct, it could not define a field of struct of itself
# but it can define a member of pointer of struct being declared 

# struct with no fields is called empty struct 
# struct are comparable based on each field being comparable 
# Go's struct embedding allows us to define anonymous fields and to access the fields of anonymous members directly 

# Converting a Go data structure to json form is called "marshaling" and only exported fields (i.e., capitalized names) can be marshaled 
# using json.Marshall 
# Inversely, converting a Json to a Go data structure is called "unmarshaling" and done by json.Unmarshal. 
# A field tag is a string of metadata associated at compile time with the field of a struct:
# Year int `json:"released"`
# Color bool `json:"color,omitempty"`

# Anonymous functions can be defined where they are used and they are defined as  
# usual but without giving the function name. 
# Anonymous functions have access to the entire lexical environment 
# When an anonymouse function requires recursion, we must first declare a variable and then assign the anonymous function to that variable 

# A variadic function is one that can be called with varying numbers of arguments and it is declared by preceding an ellipsis "..." to the type 
# of the final parameter, e.g., "func sum(vals ...int) int" defines a sum function over a varying number of numbers 
# To invoke a variadic function when the arguments are already in a slice: place an ellipsis after the final argument:
# values := []int{1,2,3,4}
# fmt.Prinfln(sum(values...))

# Go provides the defer statement that defers the execution of a function to the time whenever and however the containing function completes 
# execution. It is typically used to deal with resource release so there is no resource leakage regardless how the function exits 
# It can be used to pair lock/unlock, open/close, connection/disconnection, on-entry/on-exit functions 
# deferred function are executed in reverse order 

# Method: 
# A method is declared in a similar way to a function, except with an extra parameter appears before the function name. 
# The parameter attaches the function to the type of that parameter 
# Named types and pointers to them are the only types that may appear in a receiver function 
# Method declarations are not permitted on named types that are themselves pointer types 
# calling methods can be tricky and the following are supported by Go: 
# Both receiver argument and receiver parameter have the same type, i.e., eight both are type T or type *T 
# The receiver argument is a variable of type T, and the receiver parameter has type *T
# The receiver argument has type *T, and the receiver parameter has type T
# If all the methods of a named type T have a receiver type of T itself (not *T), it is safe to copy instances of that type 
# calling any of its method necessarily makes a copy 
# However, if any method has a pointer receiver, should avoid copying instances of T

# Method value vs Method expression 
# Method value has the form of p.method(), where p is an instance. It can be used to call a method 
# Method expression has the form of T.method(), where p is a type.  It must be called with its first parameter being type T as the receiver.

# Go provides the concept of "interface" to define a contract between a function 
# parameter and it's argument. As long as the passed concrete type has the interface 
# define, it can be passed to the function which takes the interface 