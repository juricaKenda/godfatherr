package godfatherr

const (
	WATCHED string = "godfatherr-watchdog-watched"
)

//WatchDog enables the early exit principle scope by watching the compliant Errors
func WatchDog(f func()) (e *Error) {
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case *Error:
				err := r.(*Error)
				if err.isWatched() {
					e = err.unwatch()
				} else {
					panic(err)
				}
			default:
				panic(r)
			}
		}
	}()
	f()
	return Empty()
}
