# MA job test task

Job Interview test for ma

## Usage
### Cmd parameters
```
Usage :
  -k int
        Concurrent jobs (default 5)
  -word string
        String to search count for (default "Go")
```

### Example usage
```
$ echo -e 'https://golang.org\nhttps://golang.org' | go run main.go
Count for https://golang.org: 9
Count for https://golang.org: 9
Total: 18
```

If you choose to run it in interactive mode - type "quit" to get total and exit.


## Built and Tested With

* [Go](https://golang.org/) - 1.10.3 darwin/amd64

## Authors

* **Andrey Yurchenkov** - *Initial work* - [painhardcore](https://github.com/painhardcore)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details