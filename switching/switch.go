package switching

import "sync"

func Switch() (func(), func() chan struct{}, func() chan struct{}) {
	once := new(sync.Once)
	red := make(chan struct{})
	blue := make(chan struct{})

	return func() {
			once = new(sync.Once)

			go once.Do(func() {
				select {
				case <-red:
					blue <- struct{}{}
				case <-blue:
					red <- struct{}{}
				}
			})
		},
		func() chan struct{} {
			return red
		},
		func() chan struct{} {
			return blue
		}
}
