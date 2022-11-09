# Go Learning Jam Session
Exploring the basics 
---------
### [**Tests**](https://github.com/faringet/Go_Learning_Jam_Session/blob/master/lesson_Testing/sample_test.go)
To run a specific test from a set of tests, you can use the command:
```javascript 
go test -v -run TestMultiple/simple
``` 
To run a test in parallel:
```javascript 
 t.Parallel()  
 ``` 
 ___
 ### [**Synchronization primitives**](https://github.com/faringet/Go_Learning_Jam_Session/blob/master/lesson_Testing/sample_test.go)
 #### sync.Mutex/RWMutex
 When should it be used?
 Lock/Unlock when sure that the code of the critical section changes the guarded data, otherwise (does not change) use RWMutex
 #### sync.WaitGroup
 Allows to wait for all goroutines to complete at the same time
 #### sync.Once
 Allows to run the function only once
 ___
 ### [**Panics**](https://github.com/faringet/Go_Learning_Jam_Session/blob/master/lesson_Testing/sample_test.go)
An example of how can manage and handle panics using recover()
