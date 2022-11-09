# Go Learning Jam Session
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
