# Consuming value groups

To consume a value group of type `T`,
you have to tag a `[]T` dependency with `group:"$name"`
where `$name` is the name of the value group.

You can do this in one of the following ways:

- with parameter objects
- with annotated functions

## With parameter objects

You can use [parameter objects](#parameter-objects)
to tag a slice parameter of a function as a value group.

**Prerequisites**

1. A function that consumes a parameter object.

   ```go
   type Params struct {
     fx.In

     // ...
   }

   func New(p Params) (Result, error) {
     // ...
   ```

2. This function is provided to the Fx application.

   ```go
     fx.Provide(New),
   ```

**Steps**

1. Add a new **exported** field to the parameter object
   with the type `[]T`, where `T` is the kind of value in the value group.
   Tag this field with the name of the value group.

   ```go
   type Params struct {
     fx.In

     // ...
     Watchers []Watcher `group:"watchers"`
   }
   ```

2. Consume this slice in the function that takes this parameter object.

   ```go
   func New(p Params) (Result, error) {
     // ...
     for _, w := range p.Watchers {
       // ...
   ```

::: warning

**Do not** rely on the order of values inside the slice.
The order is randomized.

:::

## With annotated functions

You can use [annotations](#annotations)
to consume a value group from an existing function.

**Prerequisites**

1. A function that accepts a slice of the kind of values in the group.

   ```go
   func NewEmitter(watchers []Watcher) (*Emitter, error) {
   ```

2. The function is provided to the Fx application.

   ```go
     fx.Provide(
       NewEmitter,
     ),
   ```

**Steps**

1. Wrap the function passed into `fx.Provide` with `fx.Annotate`.

   ```go
     fx.Provide(
       fx.Annotate(
         NewEmitter,
       ),
     ),
   ```

2. Annotate this function to state that its slice parameter is a value group.

   ```go
       fx.Annotate(
         NewEmitter,
         fx.ParamTags(`group:"watchers"`),
       ),
   ```

3. Consume this slice in the function.

   ```go
   func NewEmitter(watchers []Watcher) (*Emitter, error) {
     for _, w := range watchers {
       // ...
   ```

::: tip Tip: Functions can accept variadic arguments

You can consume a value group from a function
that accepts variadic arguments instead of a slice.

```go
func EmitterFrom(watchers ...Watcher) (*Emitter, error) {
  return &Emitter{ws: watchers}, nil
}
```

Annotate the variadic argument like it was a slice to do this.

```go
    fx.Annotate(
      EmitterFrom,
      fx.ParamTags(`group:"watchers"`),
    ),
```

:::
