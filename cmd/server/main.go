package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"
)

var storagePath string = path.Base(path.Join(".", "storage"))

func main() {
	router := gin.Default()
	router.Static("/mvn-repo", storagePath)

	//GET
	//PATTERN /{groupId-replace "." to "/" }/{artifactId}/{version}/maven-metadata.xml
	//PATTERN /{groupId-replace "." to "/" }/{artifactId}/{version}/{artifactId}-{version}.pom
	//PATTERN /{groupId-replace "." to "/" }/{artifactId}/{version}/{artifactId}-{version}.jar

	//PUT
	//PATTERN GET /{groupId-replace "." to "/" }/{artifactId}/{version}/maven-metadata.xml
	//PATTERN GET /{groupId-replace "." to "/" }/{artifactId}/{version}/maven-metadata.xml.sha1
	//PATTERN GET /{groupId-replace "." to "/" }/{artifactId}/{version}/maven-metadata.xml.md5
	//PATTERN PUT /{groupId-replace "." to "/" }/{artifactId}/{version}/{artifactId}-{version}-{timestamp}-{build_number}.pom
	//PATTERN PUT /{groupId-replace "." to "/" }/{artifactId}/{version}/{artifactId}-{version}-{timestamp}-{build_number}.jar

	router.PUT("/mvn-repo/*saveInfo", func(c *gin.Context) {
		saveInfo := c.Param("saveInfo")
		saveInfoSlice := strings.Split(saveInfo, "/")
		artifact := saveInfoSlice[len(saveInfoSlice) - 1]
		artifactDir := strings.Join(saveInfoSlice[:len(saveInfoSlice) - 1], "/")

		if _, err := os.Stat(path.Join(storagePath, artifactDir)); os.IsNotExist(err) {
			log.Println(err)
			c.String(404, "Not Found")
			return
		}

		f, err := os.OpenFile(path.Join(storagePath, saveInfo), os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			log.Println(err)
			c.String(503, "Server Error")
			return
		}
		defer f.Close()

		log.Println(artifact, c.Request.ContentLength)

		body := c.Request.Body
		defer body.Close()

		data, err := ioutil.ReadAll(body)
		if err != nil {
			log.Println(err)
			c.String(503, "Server Error")
			return
		}

		_, err = f.Write(data)
		if err != nil {
			log.Println(err)
			c.String(503, "Server Error")
			return
		}

		c.String(200, "Success Save")
	})

	srv := &http.Server{
		Addr:    ":8055",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServeTLS("./certs/cert.pem", "./certs/key.pem"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
