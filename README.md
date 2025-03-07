# statez
Simple application healthiness lib


## Usage

look at examples in /example

```go
type natsService struct {
	name          string
	statezService *statez.Service
}

func main() {

	sz := statez.NewStatez("example2")

	natsSvc := natsService{
		name:          "natsService",
		statezService: statez.NewService("natsService"),
	}

	sz.RegisterService(natsSvc.statezService)

	// go natsSvc.Run or do action

	http.HandleFunc("/ready", sz.ReadinessHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
