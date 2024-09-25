## How to run the unit tests

### if in root directory
` go test -v ./... `

### if in Helper directory
` go test `
` go test -v `
` go test -v -run=<TestFunction> `
` go test -v -run=<TestFunction>/<subtest> `

### don't use panic in unit tests because it will stop the next unit test