package draft

import "github.com/pkg/errors"

// TODO candidates to move to go-kit package.

// Sequence returns an empty slice with the specified size.
//
//     for range Sequence(5) {
//             // do something five times
//     }
//
func Sequence(size int) []struct{} {
	return make([]struct{}, size)
}

// Safe runs the action and captures a panic as its error.
//
//     serve := make(chan error, 1)
//
//     go Safe(func() error {
//             return server.ListenAndServe()
//     }, func(err error) {
//             serve <- errors.Wrap(err, "tried to listen and serve a connection")
//             close(serve)
//     })
//
func Safe(action func() error, closer func(error)) {
	var err error
	defer func() { closer(err) }()
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("unexpected panic handled: %+v", r)
		}
	}()
	err = action()
}
