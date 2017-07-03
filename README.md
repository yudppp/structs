# structs

## how to use


```
import github.com/yudppp/structs

type User struct {
    Name string `example:"ichiro" default:"suzuki"`
}

example := NewExample(User{}).(User)
fmt.Println(User.Name) # -> ichiro

d := NewDefault(User{}).(User)
fmt.Println(User.Name) # -> suzuki
```
