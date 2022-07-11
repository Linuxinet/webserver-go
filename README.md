### Deploying to Heroku [tips]

while deploying go web app to heroku, make sure you set the port value to env variable $PORT.

**Code:**

```go
	ports := os.Getenv("PORT")
        if ports == "" {
                ports = "8080"
        }
        port := ":" + ports

	.....

	srv := &http.Server{
                Addr:     port,
                ErrorLog: errorLog,
                Handler:  app.routes(),
        }
        infoLog.Printf("Listening On Port %s", port)
        err := srv.ListenAndServe()
        errorLog.Fatal(err)
```
