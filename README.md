# MA job test task

Job Interview test for ma

```
Процессу на stdin приходят строки, содержащие интересующие нас URL.
Каждый такой URL нужно дернуть и посчитать кол-во вхождений строки "Go".
В конце работы приложение выводит на экран общее кол-во найденных строк
"Go" во всех переданных URL, например:

$ echo -e 'https://golang.org\nhttps://golang.org' | go run 1.go
Count for https://golang.org: 9
Count for https://golang.org: 9
Total: 18

Введенный URL должен начать обрабатываться сразу после вычитывания и
параллельно с вычитыванием следующего. URL должны обрабатываться
параллельно, но не более k=5 одновременно. Обработчики url-ов не должны
порождать лишних горутин, т.е. если k=1000 а обрабатываемых URL-ов нет,
не должно создаваться 1000 горутин.
Нужно обойтись без глобальных переменных и использовать только
стандартные библиотеки.

```

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
