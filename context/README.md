# Context in golang
- useful solution if we require some timeout, deadline
- or some cleanup after cancellation etc.
## 3 main parts of the context package
### Deadlines
- `WithTimeout(parent context.Context, timeout time.Time) (context.Context, context.CancelFunc)`
- this cancels automatically after a **duration**
- `WithDeadline(parent context.Context, timeout time.Time) (context.Context, context.CancelFunc)`
- similar to the above but with an **absolute time** at witch we cancel.

```go
const shortDuration = 1 * time.Second
const longDuration 	= 1 * time.Minute

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), longDuration)
	defer cancel()

	select{
        case <- time.After(2*time.Second):
		fmt.Println("finished after 1 second")
	case <- ctx.Done():
		fmt.Println(ctx.Err())
	}
}
// short duration -> context deadline exceeded
// long duration  -> finished after 1 second
```

### cancellation Signals
- `WithCancel(parent context.Context)      (context.Context, context.CancelFunc)`
- here we manually have to call the cancel function
```go
func main(){
	ch := make(chan struct{})
	run := func(ctx context.Context){
		// we just create a channel, where we wait for the context cancel singal, 
		// then we loop n++ and print n every half second
		n := 1
		for{
			select{
			case <-ctx.Done():
				fmt.Println("recived graceful shutdown signal, shutting down")
				close(ch)
				return
			default:
				time.Sleep(time.Millisecond*500)
				fmt.Println(n)
				n++
			}
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	// after a timer of 3 seconds we cancel our context by hand. By calling our cancel() 
	go func(){
		time.Sleep(time.Second *3)
		fmt.Println("goodbye, sending over a shutdown")
		cancel()
	}()
	go run(ctx)
	fmt.Println("only goroutines and channel left running")
	<-ch			// this blocks untill our channel is closed
	fmt.Println("main finished running")
}

// only goroutines and channel left running
// 1
// 2
// 3
// 4
// 5
// goodbye, sending over a shutdown
// 6
// recived graceful shutdown signal, shutting down
// main finished running
```

## Request scoped values
- A request scoped variable is instancinated for a each single (http)request. So a request scoped variable instance does not exist in the context of another request.


```go
type jwt string
const auth jwt = "JWT"

func main(){
	ctx := context.WithValue(context.Background(), auth, "Some Data we want to pass down")

	bearer := ctx.Value(auth)
	str, ok := bearer.(string)
	if !ok{
		log.Fatal("not a string")
	}
	log.Println("value:", str)
}
```

## Best practices
- if we need a context in a api, define the ctx context.Context as the first argument.
- MAKE SURE context.CancelFunct is called. Otherwise the application will leak memory/goroutines.