// ************************ testing  ************************************************

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// 	// "google.golang.org/genproto/googleapis/cloud/gkeconnect/gateway/v1"
// )

// // ###############
// // func main() {
// // 	// Create Echo instance
// // 	e := echo.New()

// // 	// //  set up middlewares
// // 	middlewares.SetUpEntryMiddlewares(e)
// // 	middlewares.SetUpCors(e)

// // 	// // Set up rate limiting middleware
// // 	// ✅ 1. Hook the Middleware into Your Echo App
// // 	middlewares.LoadAllowedIPs([]string{"127.0.0.1"}) // Add your IP here

// // 	e.Use(middlewares.IpRateLimiting(1, 3, 1*time.Minute))
// // 	// e.Use(middlewares.UserRateLimiting(1, 3, 1*time.Minute))

// // 	routes.RegisterRoutes(e)

// // 	config.SetUpServer(e)
// // }

// // // ✅ Add timeout to http.Client.
// // // ✅ Add retries if backend service is down.
// // // ✅ Add JWT auth middleware before proxying.
// // // ✅ Add logging and monitoring.
// // // ✅ Dynamic routing (using configs instead of hardcoded paths).
// // // trothling pending ?
// // // universal logs with zap pending ?
// // // bg task use in go lang asynq
// // // periodic task use go lang
// // // task monitor.

// // generate mac address
// // ###############

// // const (
// // 	age  = 31
// // 	sata = 4
// // )

// // // order struct

// // type order struct {
// // 	id        string
// // 	amount    float32
// // 	status    string
// // 	createdAt time.Time // nanosecond precision

// // }
// // enumarate types

// type OrderStatus int

// const (
// 	Recevied OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Deliverd
// )

// func getStatus(status OrderStatus) {
// 	fmt.Println("changing order status :", status)
// }

// func Counter() func() int {
// 	var count int = 0
// 	return func() int {
// 		count += 1
// 		return count
// 	}
// }

// // create interface  for manage all methods for  all feature
// type paymenter interface {
// 	pay(amount int) int
// }

// // this cls or struct use for payment
// type payment struct {
// 	//  here you can use inhiritance concept by calling  another struct
// 	gateway paymenter
// }

// // this struct use for gateway
// type razorpay struct {
// }

// type strip struct {
// }

// // create function for payment
// func (p payment) makePayment(amount int) int {
// 	//  you can  makepayment through multiple payment gateway
// 	// razorpayGateway := razorpay{}

// 	return p.gateway.pay(amount)
// 	// return razorpayGateway.pay(amount)
// 	// return

// }

// func (r razorpay) pay(amount int) int {
// 	// logic  for make payment
// 	return amount
// }

// func (r strip) pay(amount int) int {
// 	// logic  for make payment
// 	return amount
// }

// // generic uses
// func geneFunction[T any](data T) any {
// 	// this is generic function
// 	fmt.Println("this is your data", data)
// 	return data
// }

// //  go routine light weight thread using for  multithering wait , add(1)  , done

// //  lets learn  about channel
// //  concurrent programming in go lang we use channel and as well as goroutine

// // whenever we have to use interface we suffix will be er added tjere
// func pro(ch chan int) {
// 	fmt.Println("this is your data from pro function")
// 	ch <- 1
// 	ch <- 2
// 	ch <- 3
// 	ch <- 4
// }

// func proString(sttr chan string) {
// 	fmt.Println("this is your data from pro function")
// 	sttr <- "this is string data"
// 	sttr <- "this is string data 2"
// 	sttr <- "this is string data 3"
// }

// func generator(data chan int) {
// 	for i := range data {
// 		// data <- i
// 		fmt.Println("this is your data from generator function", i)
// 		time.Sleep(1 * time.Second) // simulate some work
// 	}
// 	// close(data) // close the channel when done sending data

// }

// // func main() {
// // 	// In Go, channels are a way for goroutines (lightweight threads) to communicate with each other and synchronize execution
// // 	ch := make(chan int, 10)
// // 	go generator(ch) // start a goroutine that runs the generator function
// // 	for {
// // 		ch <- rand.Intn(100) // send value to channel
// // 	}

// // sttr := make(chan string, 3)
// // go proString(sttr)
// // go pro(ch) // start a goroutine that runs the pro function
// // RECIVE
// // msg1 := <-ch // receive value from channel
// // msg2 := <-ch // receive value from channel
// // msg3 := <-ch // receive value from channel
// // msg4 := <-ch // receive value from channel
// // fmt.Println("this is your value from channel", msg1)
// // fmt.Println("this is your value from channel", msg2)
// // fmt.Println("this is your value from channel", msg3)
// // fmt.Println("this is your value from channel", msg4)
// //
// // get string data from channel
// //  RECIVE STRING DATA
// // msg5 := <-sttr // receive value from channel
// // msg6 := <-sttr // receive value from channel
// // msg7 := <-sttr // receive value from channel
// // fmt.Println("this is your value from string channel", msg5)
// // fmt.Println("this is your value from string channel", msg6)
// // fmt.Println("this is your value from string channel", msg7)
// // create payment with a concrete gateway implementation
// // ************
// // part of interface
// // createPayment := payment{gateway: razorpay{}}
// // data := createPayment.makePayment(1000)
// // fmt.Println("this your data for make payment", data)
// // ***************
// // part of generic function wheerer  you camn handlee all type of  datya
// // type of  sting  data  data
// // geneFunction("this is string data")
// // listData := []int{1, 2, 3, 4, 5}
// // geneFunction(listData)
// // listStingData := []string{"datish", "fatish", "satiadah"}
// // geneFunction(listStingData)

// // }

// // this type of function we are using for testing for class oops  concept part of  struct

// // func (o *order) changeStatus(status string) {
// // 	o.status = status
// // 	fmt.Println("order data", o.status)

// // }

// // if we get input and  output in the same way then we can call then we we have to  use  structuc pointer in the both side
// // func newOrder(id string, amount float32, status string) *order {
// // 	FewOrder := order{
// // 		id:     id,
// // 		amount: amount,
// // 		status: status,
// // 	}
// // 	// fmt.Println()
// // 	return &FewOrder

// // }

// // func main() {
// // const age = 30
// // // age = 31
// // fmt.Println("take", age)
// // fmt.Println("asdfasdfa", sata)
// // // i := 1
// // // for i := 0; i <= 3; i++ {
// // // 	fmt.Println(i)
// // // 	// continue
// // // 	// if i == 2 {

// // // 	// 	continue
// // // 	// }
// // // 	// i = i + 1

// // // }
// // // for i:=  range 3 {
// // // fmt.Println("data", i)
// // // }
// // // does not ternary oprations
// // // if age := 28; age < 20 {
// // // 	fmt.Println("here you arre in ", age)
// // // } else if age <= 23 {
// // // 	fmt.Println("here you arre in ", age)
// // // } else {
// // // 	fmt.Println("here you are", age)
// // // }
// // i := 3
// // switch i {
// // case 1:
// // 	fmt.Println("this is one type")
// // case 2:
// // 	fmt.Println("this is two type")
// // case 3:
// // 	fmt.Println("this is therer adfga[dhf[dhffa[eppd]]]")
// // }
// // switch time.Now().Weekday() {
// // case time.Saturday, time.Sunday:
// // 	fmt.Println("this is my weekend!")
// // default:
// // 	fmt.Println("its workday right!")
// // }

// // test := func(i interface{}) {
// // 	switch i.(type) {

// // 	case int:
// // 		fmt.Println("its integer")
// // 	case string:
// // 		fmt.Println("its string bai !")
// // 	case bool:
// // 		fmt.Println(("its  bool right!"))

// // 	default:
// // 		fmt.Println("thaiefasdpfhabpdfb ")
// // 	}

// // }
// // test("dtighgsefgaed")
// // // arrays
// // // define  array how data size  4
// // var num [4]int
// // num[0] = 1
// // fmt.Println(len(num), num)
// // ar := [4]int{1, 2, 3, 433}
// // fmt.Println(ar)

// // // 2d array
// // nums := [2][3]int{{1, 23}, {1, 3, 4}}
// // fmt.Println(nums)
// // data := make([]int, 2)
// // data = append(data, 1)
// // data = append(data, 2)
// // data = append(data, 23)
// // fmt.Println(data)
// // //  if you unaware about elements then you can use make
// // m := make(map[string]string)
// // m["aatish"] = "aatish"
// // m["satish"] = "satish"
// // fmt.Println("this is dictonory data that might be you light it", m, len(m))
// // // if you aware about the elements then you can use
// // // only map function
// // fixzw := map[string]int{"aatish": 12, "satish": 32}
// // fmt.Println(fixzw)
// // da := []int{5, 6, 7}
// // aa := make(map[string]int)
// // aa["aatish"] = 2
// // fmt.Println(aa["aatish"])

// // for i, _ := range da {
// // 	fmt.Println("this is your data", i)
// // }
// // // if you want get index aswell

// // // creata map not dictionary
// // dic := map[string]string{"aatish": "aa", "satish": "sa", "datish": "da"}
// // for i, v := range dic {
// // 	fmt.Println(i, ": ", v)
// // }
// // for ind, num := range da {
// // 	fmt.Println(ind, num)
// // }
// // for i, c := range "rannge" {
// // 	fmt.Println("dte", i, c)
// // }
// // aatishfun := func(a, b int) int {
// // 	fmt.Println(a + b)
// // 	return a + b
// // }

// // aatishfun(2, 3)
// // // s := []int {1, 3, 4, 45}
// // // veridic function

// // // clouser functipn bassically user for store data in memory without changing
// // increament := Counter()
// // fmt.Println(increament())
// // fmt.Println(increament())
// // fmt.Println(increament())
// // // lets learn about pointersz
// // // custum data strcture using  go  ;ang its struct
// // ord := order{
// // 	id:        "1",
// // 	amount:    5045.12,
// // 	status:    "recivres",
// // 	createdAt: time.Now(),
// // }
// // fmt.Println("that's your order", ord)
// // // time.Sleep(10)
// // ord.createdAt = time.Now()
// // fmt.Println("that's your order", ord)
// // ord.changeStatus("test")
// // fmt.Println("that's your order", ord)
// // fmt.Println(newOrder("safsd", 12.2, "ssdfsadf"))
// // // fmt.Println(ord)
// // ord.changeStatus("test1")
// // fmt.Println("that's your order", ord)
// // }

// // mutex is used to protect shared data from being accessed by multiple goroutines at the same time
// // it is a way to ensure that only one goroutine can access the shared data at a time
// // mutex stands for mutual exclusion, which means that only one goroutine can access the shared
// // when we used mutex when we have to use shared data between multiple goroutines then we have to use mutex
// // when the multiple  processes are trying to access the same resource at the same time, it can lead to race conditions

// // package main

// type post struct {
// 	views int
// 	mu    sync.Mutex
// }

// func (p *post) incrementViews(wg *sync.WaitGroup) {
// 	defer wg.Done() // Decrement the wait group counter when the goroutine completes
// 	p.mu.Lock() // Lock the mutex to ensure exclusive access to the views
// 	p.views++
// 	fmt.Println("Views:", p.views)
// 	p.mu.Unlock() // Unlock the mutex to allow other goroutines to access the views

// }
// func main() {
// 	var wg sync.WaitGroup
// 	p := &post{}
// 	// p.incrementViews()
// 	// p.incrementViews()
// 	// p.incrementViews()

// 	for i := 0; i < 1000000; i++ {
// 		wg.Add(1)
// 		go p.incrementViews(&wg) // Increment the wait group counter for each goroutine
// 	}
// 	wg.Wait() // Wait for all goroutines to finish
// 	fmt.Println("Post views:", p.views)
// }
// ************************ testing ************************************************

//  Add This all things for  testing perpose

package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

// GetIPFromMAC finds the IP address for a given MAC address using `ip neigh`
func GetIPFromMAC(mac string) (string, error) {
	// Run `ip neigh`
	cmd := exec.Command("ip", "neigh")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to run ip neigh: %v", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	// Normalize MAC to lowercase for matching
	mac = strings.ToLower(mac)
	fmt.Println("this ::=====", mac)
	// Regex: <IP> dev <iface> lladdr <MAC> <state>
	re := regexp.MustCompile(`^(\S+)\s+dev\s+\S+\s+lladdr\s+([0-9a-f:]+)\s+.*$`)
	fmt.Println("iss it ::=====", re)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) == 3 {
			ip := matches[1]
			foundMAC := strings.ToLower(matches[2])
			if foundMAC == mac {
				return ip, nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("no IP found for MAC %s", mac)
}

func main() {
	// Example usage
	ip, err := GetIPFromMAC("52:54:00:8b:90:bf")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("IP Address:", ip)
	}
}