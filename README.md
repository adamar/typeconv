# typeconv


A simple library for handling all the different type conversions in Golang (for built-in types)




## Why?

I got sick of trying to figure out how to do more complicated conversions between types and just wanted a simple library for it, or even just a reference



## Interface


The intention is every function should have an identical function signature
```

func {SOURCE_TYPE}To{DESTINATION_TYPE}(input) (output, error) {

```

ie. 
```

func Int32ToFloat32(input int32) (float32, error) {

```

even if the underlying functions dont return an error, the typeconv function will follow this convention to keep everything aligned


