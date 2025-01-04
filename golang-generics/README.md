# GOLANG GENERICS

Generics in Go have been available since version 1.18. With generics, we can modify
data types as desired without having to specify the data type again in a function or
convert the data type.

## Type Parameters
The creation of type parameters starts with creating pair brackets [], where we can include more than one type parameter inside.

```Go
func Generic[T1, T2]() {

}
```

## Type Constraints
Di golang wajib juga untuk menentukan type datanya dalam Generic[T1, T2]