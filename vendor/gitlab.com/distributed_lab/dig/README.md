# Dig
This library is built on gitlab.com/distributed_lab/figure, allowing you to use the default `fig` tags.

# Usage

```go
package dig

import "gitlab.com/distributed_lab/dig"

type Config struct {
	Age        int    `dig:"USER_AGE"`
	Name       string `dig:"USER_NAME,required"`
	PassportID string `dig:"USER_PASSPORT_ID,clear"` // An environment variable will be unset after being read.
}

func main() {
	config := Config{}
	err := dig.Out(&config).Where(values).Now()
	if err != nil {
		panic(err)
	}
}
```