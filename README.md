# Go Learning Jam Session
### [Exploring the basics](https://github.com/faringet/Go_Learning_Jam_Session/edit/master/README.md#exploring-the-basics-1)

* [Tests](https://github.com/faringet/Go_Learning_Jam_Session#tests)
  * [default testing package](https://github.com/faringet/Go_Learning_Jam_Session#default-testing-package)
  * [Go Convey](https://github.com/faringet/Go_Learning_Jam_Session#go-convey)
* [Synchronization primitives](https://github.com/faringet/Go_Learning_Jam_Session#synchronization-primitives)
  * [sync.Mutex/RWMutex](https://github.com/faringet/Go_Learning_Jam_Session#syncmutexrwmutex)
  * [sync.WaitGroup](https://github.com/faringet/Go_Learning_Jam_Session#syncwaitgroup)
  * [sync.Once](https://github.com/faringet/Go_Learning_Jam_Session#synconce)
* [Panics](https://github.com/faringet/Go_Learning_Jam_Session#panics)
* [Reading/Writing Files](https://github.com/faringet/Go_Learning_Jam_Session#readingwriting-files)
* [Shell Commands](https://github.com/faringet/Go_Learning_Jam_Session#shell-commands)
* [Context](https://github.com/faringet/Go_Learning_Jam_Session#context)
* [JSON](https://github.com/faringet/Go_Learning_Jam_Session#json)
  * [default json package](https://github.com/faringet/Go_Learning_Jam_Session#default-json-package)
  * [GJSON](https://github.com/faringet/Go_Learning_Jam_Session#gjson)
  * [SJSON](https://github.com/faringet/Go_Learning_Jam_Session#sjson)

### [Design patterns](https://github.com/faringet/Go_Learning_Jam_Session#design-patterns)

* [Strategy pattern](https://github.com/faringet/Go_Learning_Jam_Session#strategy-pattern)
* [Builder + Director pattern](https://github.com/faringet/Go_Learning_Jam_Session#builder--director-pattern)



___
Exploring the basics 
---------
### Tests
#### [**default testing package**](https://github.com/faringet/Go_Learning_Jam_Session/blob/master/Tests/Testing/sample_test.go)
To run a specific test from a set of tests, you can use the command:
```javascript 
go test -v -run TestMultiple/simple
``` 
To run a test in parallel:
```javascript 
 t.Parallel()  
 ``` 
 #### [**Go Convey**](https://github.com/faringet/Go_Learning_Jam_Session/blob/master/Tests/GoConvey/convey_test.go) 
 First acquaintance. [**View full feature tour**](http://goconvey.co/)
 
 ___
 ### Synchronization primitives
 #### [**sync.Mutex/RWMutex**](https://github.com/faringet/Go_Learning_Jam_Session/blob/master/SyncPrimitives/Mutex/main.go)
 When should it be used?
 Lock/Unlock when sure that the code of the critical section changes the guarded data, otherwise (does not change) use RWMutex
 #### [**sync.WaitGroup**](https://github.com/faringet/Go_Learning_Jam_Session/blob/master/SyncPrimitives/WaitGroup/main.go)
 Allows to wait for all goroutines to complete at the same time
 #### [**sync.Once**](https://github.com/faringet/Go_Learning_Jam_Session/blob/master/SyncPrimitives/Once/main.go)
 Allows to run the function only once
 ___
 ### Panics
An [**example**](https://github.com/faringet/Go_Learning_Jam_Session/blob/master/Panics/main.go) of how can manage and handle panics using recover()

___
 ### Reading/Writing Files
Introduction to [**io/ioutil**](https://github.com/faringet/Go_Learning_Jam_Session/blob/master/IO/main.go) and os packages

___
### Shell Commands
The [**top (table of processes) command**](https://github.com/faringet/Go_Learning_Jam_Session/blob/master/Shell/main.go) shows a real-time view of running processes in macOS/Linux and displays kernel-managed tasks. The command also provides a system information summary that shows resource utilization, including CPU and memory usage

___
### Context
Package context defines the Context type, which carries [**deadlines**](https://github.com/faringet/Go_Learning_Jam_Session/blob/master/Context/main.go), cancellation signals, and other request-scoped values across API boundaries and between processes.

___
### JSON
#### [**default json package**](https://github.com/faringet/Go_Learning_Jam_Session/blob/master/JSON/DefaultJson/main.go)
Practice in json.Marshal/Unmarshal
#### [**GJSON**](https://github.com/faringet/Go_Learning_Jam_Session/blob/master/JSON/GJSON/main.go)
Practice in GJSON package

Quick overview of the path syntax :

```json
{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Fear and Loathing in Las Vegas",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}
```
```
"name.last"          >> "Anderson"
"age"                >> 37
"children"           >> ["Sara","Alex","Jack"]
"children.#"         >> 3
"children.1"         >> "Alex"
"child*.2"           >> "Jack"
"c?ildren.0"         >> "Sara"
"fav\.movie"         >> "Fear and Loathing in Las Vegas"
"friends.#.first"    >> ["Dale","Roger","Jane"]
"friends.1.last"     >> "Craig"
```

[**View full feature tour**](https://github.com/tidwall/gjson)

#### [**SJSON**](https://github.com/faringet/Go_Learning_Jam_Session/tree/master/JSON/SJSON)
Practice in SJSON package

[**View full feature tour**](https://github.com/tidwall/sjson)

___
Design patterns
---------
### [**Strategy pattern**](https://github.com/faringet/Go_Learning_Jam_Session/blob/master/DesignPatterns/Strategy/main.go)
is a behavioralgn pattern that enables selecting an algorithm at runtime. Instead of implementing a single algorithm directly, code receives run-time instructions as to which in a family of algorithms to use.

Greate explanation from Israel Miles [**on Medum**](https://levelup.gitconnected.com/the-strategy-pattern-in-go-2072d2b9d6ae)

___
### [**Builder + Director pattern**](https://github.com/faringet/Go_Learning_Jam_Session/blob/master/DesignPatterns/Builder/main.go) 
is a creational pattern designed to provide a flexible solution to various object creation problems in object-oriented programming. The intent of the Builder design pattern is to separate the construction of a complex object from its representation.

Greate explanation from josué Parra Rosales [**on Medum**](https://medium.com/@josueparra2892/builder-pattern-in-go-56605f9e7387)

___
Free formulation of SOLID principles
---------

S - Single responsibility principle (SRP) Принцип единственной ответственности.

	Если у нас БД, пишем сервис по работе с БД. Для очередей пишем отдельный сервис по работе с очередями.
	Потом будет третье звено, которое знает и о БД и об очередях. БД об очередях не знает.
	Каждый объект программы должен знать только об одной (своей сфере). Иметь единственную ответственность.
	Ответственности прописываем на уровне пакетов (подпакетов) .pkg/database/client-db.go и .pkg/amqp/client.go

O - Open/Close principle (OCP) Принцип открытости/закрытости.

	Каждый сервис/структура должна быть открыта для расширения (добавление нового функционала), но никак для изменения.
	Для добавления новых функций - открыта. Для редактирования - закрыта.
	Один раз что-то создали и в дальнейшем это не менять.
	Как пример - создаем общий интерфейс (тем самым закрыли функционал на добавление чего-либо), а в структурах уже
	начинаем расширять данный интерфейс.

L - Liskov substitution principle (LSP) Принцип подстановки Барбары Лисков

	Функции, которые используют базовый тип должны иметь возможность использовать подтипы базового типа не зная об этом.
	Когда интерфейс можем реализовать в каком-то наследнике.


I - Interface segregation principle (ISP) Принцип разделения интерфейсов

	Клиенты, с точки зрения реализаций, не должны зависеть от методов, которые они не используют.
	Слишком "жирные" интерфейсы необходимо разделять на более мелкие. То есть если в интерфейсе будет огромная куча методов,
	то нам придется у всех структур реализовывать эти методы, которые привязаны к этому интерфейсу - что не круто и очень расходно.

D - Dependency inversion principle (DIP) Принцип инверсии зависимостей

	Модули верхних уровней не должны зависеть от модулей нижних уровней. Оба типа модулей должны зависеть от абстракций.
	Сами абстракции не должны зависеть от деталей.
	Нужно применять такой принцип, когда мы внутри структуры реализуем функционал, мы не должны реализовывать от какого-то
	объекта. Объект нужно привести к интерфейсу.
	Как пример - выполнять взаимодействие с БД только через интерфейс. Никогда не нужно выполнять через реализацию.
