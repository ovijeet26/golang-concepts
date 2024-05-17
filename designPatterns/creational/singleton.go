package creational

import (
	"fmt"
	"sync"
)

/*
Singleton is a creational design pattern, which ensures that only one object of its kind exists and provides a single point of access to it for any other code.
Singleton has almost the same pros and cons as global variables. Although they’re super-handy, they break the modularity of your code.
You can’t just use a class that depends on a Singleton in some other context, without carrying over the Singleton to the other context. Most of the time, this limitation comes up during the creation of unit tests.
*/

var lock sync.Mutex

type single struct {
}

var singleObject *single

func GetInstance() *single {

	if singleObject == nil {
		lock.Lock()
		defer lock.Unlock()
		// second check is to ensure, if multiple go routines bypass th first chck, duplicate instances are no created
		if singleObject == nil {
			fmt.Println("Creaing a new Singleton object")
			singleObject = new(single)
		} else {
			fmt.Println("Inner - Singleton object is already created!")
		}
	} else {
		fmt.Println("Outer - Singleton object is already created!")
	}

	return singleObject
}
