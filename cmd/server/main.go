package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//PATTERN /{groupId-replace "." to "/" }/{artifactId}/{version}/maven-metadata.xml
	//PATTERN /{groupId-replace "." to "/" }/{artifactId}/{version}/{artifactId}-{version}.pom
	//PATTERN /{groupId-replace "." to "/" }/{artifactId}/{version}/{artifactId}-{version}.jar

	router.GET("/mvn-repo/info/crazylab/test/1.0-SNAPSHOT/maven-metadata.xml", func(c *gin.Context) {
		c.Header("Content-type", "text/xml")
		c.String(200, `
		<metadata modelVersion="1.1.0">
			<groupId>info.crazylab</groupId>
			<artifactId>test</artifactId>
			<version>1.0-SNAPSHOT</version>
		</metadata>`)
	})
	router.GET("/mvn-repo", func(c *gin.Context) {

		fmt.Printf("%#v", c)

		c.JSON(404, gin.H{})
	})

	//PATTERN GET /{groupId-replace "." to "/" }/{artifactId}/{version}/maven-metadata.xml
	//PATTERN GET /{groupId-replace "." to "/" }/{artifactId}/{version}/maven-metadata.sha1
	//PATTERN GET /{groupId-replace "." to "/" }/{artifactId}/{version}/maven-metadata.md5
	//PATTERN PUT /{groupId-replace "." to "/" }/{artifactId}/{version}/{artifactId}-{version}-{timestamp}-{build_number}.pom
	//PATTERN PUT /{groupId-replace "." to "/" }/{artifactId}/{version}/{artifactId}-{version}-{timestamp}-{build_number}.jar

	router.PUT("/mvn-repo", func(c *gin.Context) {

		fmt.Printf("%#v", c)

		c.JSON(404, gin.H{})
	})

	router.Run(":8055")
}
