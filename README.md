# Jet

## Introduction

Jet is a dynamically-typed interpreted object-_orchestration_ language inspired by Go, PHP, and Python.

## Language Features

### Entrypoint

Jet uses the `main` method as its default entrypoint.

`.jet` is the default file extension for Jet files.

### Comments

### Scope

A variable declared in :
* File scope is global
* Descriptor scope is accessible only to that Descriptor and Objects it Describes
* Object scope is accessible only to that Object and blocks that Object has been instantiated in.
* Function scope is isolated to code within the function
* For/If scope is isolated to that for/if loop

### Types

Jet supports the following types:

* `int`
* `float`
* `boolean`
* `string`
* `array`
* `map`

Under the hood, arrays and maps are effectively identical. This allows for a unified set of methods for access and manipulation.

* `(map, value)->add || add(map, value)` - adds the value to the map with the lowest available int starting at `0`
* `(map, key)->remove || remove(map, key)` - removes the specified key and associated value from the map
* `map->unique || unique(map)` - returns a copy of the map with any duplicate **values** removed

Consider the following:
```
myArray = ["a", "b", "c", "d"]
```

In Jet, this is actually a map:
```
[0: "a", 1: "b", 2: "c", 3: "d"]
```

This allows simple declaration, predictable ordering when iterating, and the ability to dynamically change values with ease.

As Jet is dynamically typed, maps/arrays do not care about having mixed value or key types.

### Returning

Jet does not have the return keyword, but instead uses the pass through syntax: `->`

Methods in Jet do not have explicit expectations in regard to return values, so numerous arguments can be returned by encapsulating them within parenthesis like so: `(x, y)->`

### Methods

#### 1. Declaration

Methods are declared with the `meth` keyword and code blocks are defined with braces `{}`.

Arguments are declared after a colon (`:`) and are comma seperated. When an unknown number of arguments are required, the vararg suffix (`*`) can be used. Additional arguments will be compiled into a map and can be accessed via the argument name that precedes the `*` token. 

Methods can be declared in any of the following formats.

```
meth myFunction {}

meth myFunction: arg1, arg2 {}

meth myFunction: arg1, arg* {}

myClosure = meth: x, y { (x + y)-> }

myClosure = meth {}
```

#### 2. Calling Methods

If a method has no parameters, parenthesis `()` can be omitted.
```
myMethod
myObject.MyMethod
myMethod(x, y)
myObject.MyMethod()
```
```
a, _ = myMethod  // Returns 
_, b = myMethod
c = myMethod
```

#### 3. Non-Declarative Argument Parsing

Just as `->` is used to return, values can be passed directly into functions to create function chains as follows:
```
meth Foo: array { (a, b, c)-> }
meth Bar: args* { (d, e)-> }
meth Baz: arg1, arg2 { (string)-> }

myString  = Foo(myArray)->Bar->Baz

a, b, c   = Foo(myArray)
d, e      = Bar(a, b, c)
myString2 = Baz(d, e)

if myString == myString2 {
    // Evaluates to true.
} 
```

This can be particularly powerful when using varargs as the return values of the first function (if there are multiple) will be produced as an array. One of 3 things will then occur:
1. The number of values being passed through is identical to the number of arguments taken, and arguments will be assigned to the next methods parameters in return order.
   1. Returned values that already are arrays will not be broken down as a pass through.
2. The number of values being passed through is not identical to the expected number of parameters, however the last/only argument is a vararg. Values that cannot be assigned to the single value arg, will be passed to the vararg.
3. The number of values being passed through is not identical to the expected number of parameters and a panic will occur.

### Descriptors

Descriptors are Jet's response to Classes. A descriptor is used to **describe** the functionality of a given object.

Both descriptors and their properties should be capitalised. Properties not declared as a descriptor argument can be labelled as constant.

```
describe Vehicle: Seats {
    const Material = "Metal"
    Wheels = 4
}

describe Jet: Name, TopSpeed {
    meth speedBoost { (TopSpeed * 2)-> }

    meth canFly { (true)-> }
}
```

### Objects

Objects are orchestrated from descriptors. Arguments are inherited from the descriptors in the order they are assigned to the object.

Inherited methods can be overloaded by the object. New methods can also be added to the object, allowing utilisation of properties from across descriptors.

Properties not included as a descriptor argument can be updated later. Accessing properties within the object requires the `Descriptor.Property` format.

New properties cannot be added in runtime code, and constants cannot be updated. Attempting to do so will cause Jet to panic.

```
object FighterJet: Vehicle, Jet {
    overload speedBoost { (Jet.TopSpeed * 4)-> }
    
    meth describeFighterJet {
        ("{Jet.Name} is made of {Vehicle.Material} has {Vehicle.Wheels}")->print
    }
}

myFighter = FighterJet(2, "Falcon", 100)
3->myFighter.Wheels
```

## Language Objectives

* [ ] Jet uses a common entrypoint; `main` will always be used to initialise a program.
* [ ] Inheritance is "shallow". Object-types can be defined and their default methods defined and implemented, however one object-type cannot inherit from another. They must be orchestrated together.
* [ ] Attributes and child methods are accessed via `.` syntax
* [ ] Values are piped into and out of functions with the "spoon" syntax: `()->`

## Resources

[High Level Principles / Tokenization](https://www.freecodecamp.org/news/the-programming-language-pipeline-91d3f449c919/)

[Abstract Syntax Trees](https://www.twilio.com/blog/abstract-syntax-trees)

[Golang AST Package](https://tech.ingrid.com/introduction-ast-golang/)