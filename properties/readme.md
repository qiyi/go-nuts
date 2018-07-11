# properties 

Go properties package support properties file parse and write.

To load properties from file:
```go
p, err := properties.LoadFile("filename")
if err != nil {
	return err
}
if v1, ok := p.Get("k1"); ok {
	fmt.Println("Get k1=" + v1)
}
```

The Properties struct will not keep the order of keys, if you need the sorted keys including comments and blank lines, use the Parser:
```go
file, err := os.Open(name)
if err != nil {
	return err
}
defer file.Close()
p, err := properties.NewParser(file)
if err != nil {
	return err
}
for {
	n, err := parser.Next()
	if err != nil {
		return nil, err
	}
	if n == nil {
		break
	}
	if pn, ok := n.(*PropertyNode); ok {
		fmt.Println(pn.Key + "=" + pn.Value)
	}
}
```

To write properties to file:
```go
err := p.StoreFile("filename")
if err != nil {
	fmt.Printlt("Store properties failed.")
}
```

The Writer supports writing comments and blanks:
```go
file, err := os.Create(name)
if err != nil {
	return err
}
defer file.Close()
writer, err := properties.NewWriter(file)
if err != nil {
	return err
}
writer.WriteComment("comments...")
writer.WriteProperty("key", "value")
``` 
