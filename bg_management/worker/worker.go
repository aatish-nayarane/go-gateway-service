// main file in scheduler package

package main

import (
	"fmt"
	"net/smtp"
	// "os"
	// "os/signal"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

// Make a redis pool
var redisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle:   5,
	Wait:      true,
	Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", ":6379")
	},
}

type Context struct {
	email  string
	userId int64
}

// Make an enqueuer with a particular namespace
// var enqueuer = work.NewEnqueuer("application_namespace", redisPool)

func main() {
	// Enqueue a job named "send_email" with the specified parameters.

	// intialize que and adn  job for  that so they can take input from  that  q and submit into job you can use context
	// pool, err := enqueuer.Enqueue("send_welcome_email", work.Q{"email_address": "test@example.com", "user_id": 4})
	pool := work.NewWorkerPool(Context{}, 10, "application_namespace", redisPool)
	pool.Middleware((*Context).Log)
	pool.Middleware((*Context).FindCustomer)
	pool.Job("send_welcome_email", (*Context).SendWelcomeEmail)
	pool.Start()
	// Wait for a signal to quit:
	// signalChan := make(chan os.Signal, 1)
	// signal.Notify(signalChan, os.Interrupt, os.Kill)
	// <-signalChan

	// Stop the pool
	pool.Stop()
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting  job", job.Name)
	fmt.Println("this is your job id ", job.ID)
	return next()
}

func (c *Context) FindCustomer(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("pointer for job, next, context", job, next, c)
	// if job.ArgString("emai_address")
	if _, ok := job.Args["user_id"]; ok {
		c.userId = job.ArgInt64("user_id")
		c.email = job.ArgString("email_address")
		if err := job.ArgError(); err != nil {
			return err
		}
	}
	return next()
}

func (c *Context) SendWelcomeEmail(job *work.Job, next *work.NextMiddlewareFunc) error {
	addr := job.ArgString("email_address")
	if err := job.ArgError(); err != nil {
		return err
	}
	fmt.Println("sending email", c)
	from, password := "xyz@gmail.com", "sthkxlpixjferhfd"
	to := []string{
		addr,
	}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	message := []byte("This is a Welcome email message.")
	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// sending email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("email sent successfully !")
	return nil
}
