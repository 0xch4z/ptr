# ptr

### Why is this a thing?

There is no function to create a reference to a literal value in Go's built-ins
or it's standard library. Before generics, I found myself copy-pasting the same
`stringPtr(string *string`, `intPtr(int) *int`, etc. function declarations in
every code base I started.

### Usage

#### `ptr.To`

```
instance, err := client.UpdateInstance(ctx, instanceID, &linodego.InstanceUpdateOptions{
  Tags: ptr.To([]string{"env=production", "team=pals"}),
})
// ...
```

