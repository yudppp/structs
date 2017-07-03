# structs

## how to use


```
import "github.com/yudppp/structs"

type User struct {
    Name string `example:"ichiro" default:"suzuki"`
}

example := structs.NewExample(User{}).(User)
fmt.Println(User.Name) # -> ichiro

d := structs.NewDefault(User{}).(User)
fmt.Println(User.Name) # -> suzuki
```
