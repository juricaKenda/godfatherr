# Godfatherr
![Godfather icon](https://github.com/juricaKenda/godfatherr/blob/master/visuals/godfatherr-icon-smallest.png)
*I'm gonna make you an offer you can't refuse..*  

## Code samples
### Replace nils w/ optional errors
    err := someFunc()
    if err.IsPresent(){
          ...
    }
### Assign additional context
    err := someFunc()
    if err.IsPresent() && err.ContainsCtx("redis"){
          ...
    }
### Assign fatality concept
    err := someFunc()
    if err.IsPresent() && err.IsFatal(){
          ...
    }

### Perform std operations directly
    err := someFunc()
    if err.IsPresent() && err.ContainsCtx("redis"){
          err.Print()
          return err.Error()
    }
