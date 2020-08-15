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
    if err.IsPresent() && err.ContainsCtx(godfatherr.REDIS){
          ...
    }
### Assign fatality concept
    err := someFunc()
    if err.IsPresent() && err.IsFatal(){
          ...
    }

### Perform std operations directly
    err := someFunc()
    if err.IsPresent() && err.ContainsCtx(domain.FOO){
          err.Print()
          return err.Error()
    }
    
## Currently available features (concepts)
- context
- fluency 
- operation chaining 
- watchdog environment (beta)

## Philosophy and purpose 
The Godfatherr library embodies the fluency and simplicity. The "go" approach to errors is extremely 
appealing because of the simplicity it introduces (ie. errors are defined exclusively by their error messages).
However, in my (@juricaKenda) opinion, this approach lacks context. When I visualize software constructs, 
I see engines and services on different abstraction and contextual levels. A "raw", string only, error can 
never express where it came from. The context of it stayed at the level it was produced at. 
 
## Issues to conquer 
### Error comparison
Implement the function to compare two errors. The question a client of this method will ask is: 
"Are these two Errors equal?". The method returns true if they are, false otherwise. When implementing
this method, consider the context (does all context have to be equal or is a certain "similarity" enough).
