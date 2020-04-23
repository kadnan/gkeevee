# gKeeVee - A simple key-value store DB in Go
`gKeeVee` is a simple file based key-value store DB written in Go language. It uses [MessagePack](https://msgpack.org/index.html) to store the compressed data into a file.

## What's new in gKeeVee?

Infact nothing. I made this package as a part of my journey learning Golang. There are many other packages available like **Badger** which provides more facilities. 

`gKeeVee` is a very simple db that only accepts `strings` as value. It is by design. If you want to store values like `int`, `float` etc, you will have to cast them to strings first and then store. For complex structures like `map` or `struct` you can convert them into a JSON string and store them as value.