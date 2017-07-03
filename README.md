# structs

## how to use


```
import (
    "fmt"

    "github.com/yudppp/structs"
)
type User struct {
    Name string `example:"ichiro" default:"suzuki"`
}

func main() {
    user := structs.NewExample(User{}).(User)
    fmt.Println(user.Name) # -> ichiro

    user = structs.NewDefault(User{}).(User)
    fmt.Println(user.Name) # -> suzuki
}

```
