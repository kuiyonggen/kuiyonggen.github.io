# Annotations

You can annotate functions and values with the `fx.Annotate` function
before passing them to
`fx.Provide`, `fx.Supply`, `fx.Invoke`, `fx.Decorate`, or `fx.Replace`.

This allows you to re-use a plain Go function to do the following
without manually wrapping the function to use
[parameter](#parameter-objects) or [result](#result-objects) objects.

- [feed values to a value group](#with-annotated-functions)
- [consume values from a value group](#with-annotated-functions)

<!-- TODO: named values and optional dependencies in the list above -->

## Annotating a function

**Prerequisites**

A function that:

- does not accept a [parameter object](#parameter-objects), when
  annotating with `fx.ParamTags`.
- does not return a [result object](#result-objects) when annotating
  with `fx.ResultTags`.

**Steps**

1. Given a function that you're passing to
   `fx.Provide`, `fx.Invoke`, or `fx.Decorate`,

   ```go
       fx.Provide(
         NewHTTPClient,
       ),
   ```

2. Wrap the function with `fx.Annotate`.

   ```go
       fx.Provide(
         fx.Annotate(
           NewHTTPClient,
         ),
       ),
   ```

3. Inside `fx.Annotate`, pass in your annotations.

   ```go
       fx.Provide(
         fx.Annotate(
           NewHTTPClient,
           fx.ResultTags(`name:"client"`),
         ),
       ),
   ```

   This annotation tags the result of the function with a name.

**Related resources**

- [fx.Annotation](https://pkg.go.dev/go.uber.org/fx#Annotation)
  holds a list of all supported annotations.

## Casting structs to interfaces

You can use function annotations to cast a struct value returned by a function
into an interface consumed by another function.

**Prerequisites**

1. A function that produces a struct or pointer value.

   ```go
   func NewHTTPClient(Config) (*http.Client, error) {
   ```

2. A function that consumes the result of the producer.

   ```go
   func NewGitHubClient(client *http.Client) *github.Client {
   ```

3. Both functions are provided to the Fx application.

   ```go
       fx.Provide(
         NewHTTPClient,
         NewGitHubClient,
       ),
   ```

**Steps**

1. Declare an interface that matches the API of the produced `*http.Client`.

   ```go
   type HTTPClient interface {
     Do(*http.Request) (*http.Response, error)
   }

   // This is a compile-time check that verifies
   // that our interface matches the API of http.Client.
   var _ HTTPClient = (*http.Client)(nil)
   ```

2. Change the consumer to accept the interface instead of the struct.

   ```go
   func NewGitHubClient(client HTTPClient) *github.Client {
   ```

3. Finally, annotate the producer with `fx.As` to state
   that it produces an interface value.

   ```go
       fx.Provide(
         fx.Annotate(
           NewHTTPClient,
           fx.As(new(HTTPClient)),
         ),
         NewGitHubClient,
       ),
   ```

With this change,

- the annotated function now only puts the interface into the container
- the producer's API remains unchanged
- the consumer is decoupled from the implementation and independently testable
