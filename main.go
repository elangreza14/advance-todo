package main

import (
	"log"

	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/infrastructure"
)

func main() {
	env, err := config.NewEnv()
	if err != nil {
		log.Fatal(err)
	}

	if err = infrastructure.Run(env); err != nil {
		log.Fatal(err)
	}
}

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"

// 	"bou.ke/monkey"
// )

// func main() {
// 	monkey.Patch(fmt.Println, func(a ...any) (n int, err error) {

// 		for _, v := range a {
// 			if strings.Contains(fmt.Sprint(v), "hell") {
// 				return fmt.Fprintln(os.Stdout, "censored")
// 			}
// 		}

// 		return fmt.Fprintln(os.Stdout, a...)
// 	})
// 	fmt.Println("what the hell?")
// }

// package main

// import (
// 	"fmt"
// 	"net"
// 	"net/http"
// 	"reflect"

// 	"bou.ke/monkey"
// )

// func main() {
// 	var d *net.Dialer
// 	monkey.PatchInstanceMethod(reflect.TypeOf(d), "Dial", func(_ *net.Dialer, _, _ string) (net.Conn, error) {
// 		return nil, fmt.Errorf("no dialing allowed")
// 	})
// 	_, err := http.Get("http://google.com")
// 	fmt.Println(err) // Get http://google.com: no dialing allowed
// }

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"reflect"
// 	"strings"

// 	"bou.ke/monkey"
// )

// func main() {
// 	var guard *monkey.PatchGuard
// 	guard = monkey.PatchInstanceMethod(reflect.TypeOf(http.DefaultClient), "Get", func(c *http.Client, url string) (*http.Response, error) {
// 		guard.Unpatch()
// 		defer guard.Restore()

// 		if !strings.HasPrefix(url, "https://") {
// 			return nil, fmt.Errorf("only https requests allowed")
// 		}

// 		return c.Get(url)
// 	})

// 	_, err := http.Get("http://google.com")
// 	fmt.Println(err) // only https requests allowed
// 	resp, err := http.Get("https://google.com")
// 	fmt.Println(resp.Status, err) // 200 OK <nil>
// }
