// package main
//
// import (
//
//	"fmt"
//	"net/http"
//	"sync"
//	"time"
//
// )
//
//	func callAPI(url string, wg *sync.WaitGroup) {
//		defer wg.Done()
//
//		resp, err := http.Get(url)
//		if err != nil {
//			fmt.Printf("Error calling API: %v\n", err)
//			return
//		}
//		defer resp.Body.Close()
//
//		// Đọc và xử lý dữ liệu từ phản hồi API ở đây (tùy thuộc vào API cụ thể).
//	}
//
//	func main() {
//		const numGoroutines = 1000
//		const apiUrl = "https://api.hodo.solutions/api/go/healthcheck" // Thay thế URL của API thực tế ở đây
//
//		var wg sync.WaitGroup
//		wg.Add(numGoroutines)
//
//		for i := 0; i < numGoroutines; i++ {
//			time.Sleep(10 * time.Millisecond)
//			go callAPI(apiUrl, &wg)
//		}
//
//		// Đợi cho tất cả các goroutine kết thúc trước khi thoát chương trình.
//		wg.Wait()
//	}
// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"time"

// 	"github.com/you/client/internal/pkg/token"
// )

// func main() {
// 	prvKey, err := ioutil.ReadFile("cert/id_rsa")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	pubKey, err := ioutil.ReadFile("cert/id_rsa.pub")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	jwtToken := token.NewJWT(prvKey, pubKey)

// 	// 1. Create a new JWT token.
// 	tok, err := jwtToken.Create(time.Hour, "Can be anything")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println("TOKEN:", tok)

//		// 2. Validate an existing JWT token.
//		content, err := jwtToken.Validate(tok)
//		if err != nil {
//			log.Fatalln(err)
//		}
//		fmt.Println("CONTENT:", content)
//	}
//
// mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.0
package main

import (
	"context"
	"github/cuongcm/ws_rec/crud"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//MgShopAggreate()
	// Khởi tạo bộ định tuyến Gin
	router := gin.Default()

	// Xử lý yêu cầu GET đến /result
	router.GET("/result", func(c *gin.Context) {
		// Thực hiện các xử lý cần thiết để lấy kết quả
		result, _ := crud.PushExam()
		// Trả về kết quả dưới dạng JSON
		c.JSON(http.StatusOK, gin.H{"result": result})
	})
	router.POST("/create", func(c *gin.Context) {
		// Thực hiện các xử lý cần thiết để lấy kết quả
		crud.CreateShops()
		// Trả về kết quả dưới dạng JSON
		c.JSON(http.StatusOK, gin.H{"result": nil})
	})
	router.PUT("/update", func(c *gin.Context) {
		// Thực hiện các xử lý cần thiết để lấy kết quả
		var input crud.UpdateInput
		c.ShouldBindJSON(&input)
		crud.UpdateShopD(context.TODO(), map[string]string{
			"id":             c.Query("id"),
			"shopify_domain": c.Query("domain")}, &input)
		// Trả về kết quả dưới dạng JSON
		c.JSON(http.StatusOK, gin.H{"result": nil})
	})

	router.GET("/list", func(c *gin.Context) {
		// Thực hiện các xử lý cần thiết để lấy kết quả

		qParam := crud.BuidFactory(c)
		var count int
		result, _ := crud.GetShopM(context.TODO(), qParam, &count)
		// Trả về kết quả dưới dạng JSON
		c.JSON(http.StatusOK, gin.H{"result": result, "count": count})
	})
	router.GET("/ag-list", func(c *gin.Context) {
		// Thực hiện các xử lý cần thiết để lấy kết quả

		qParam := crud.BuidFactory(c)
		var count int
		result, _ := crud.AggregateTest(context.TODO(), qParam)
		// Trả về kết quả dưới dạng JSON
		c.JSON(http.StatusOK, gin.H{"result": result, "count": count})
	})
	router.GET("/chat-gpt", func(c *gin.Context) {
		// Thực hiện các xử lý cần thiết để lấy kết quả

		crud.ChatGPTAgg()
		// Trả về kết quả dưới dạng JSON
		c.Status(http.StatusOK)
	})
	router.GET("/my-list", func(c *gin.Context) {
		// Thực hiện các xử lý cần thiết để lấy kết quả
		qParam := crud.BuidFactory(c)
		result, count, err := crud.MyMongoDbList(context.TODO(), qParam)
		// Trả về kết quả dưới dạng JSON
		c.JSON(http.StatusOK, gin.H{"result": result, "count": count, "err ": err}) // Trả về kết quả dưới dạng JSON
	})

	router.GET("/my-func", func(c *gin.Context) {
		// Thực hiện các xử lý cần thiết để lấy kết quả
		qParam := crud.BuidFactory(c)
		result, err := crud.AggegrateWithFunc(context.TODO(), qParam)
		// Trả về kết quả dưới dạng JSON
		c.JSON(http.StatusOK, gin.H{"result": result, "err ": err}) // Trả về kết quả dưới dạng JSON
	})

	router.GET("/my-first-elk", func(c *gin.Context) {
		// Thực hiện các xử lý cần thiết để lấy kết quả
		qParam := crud.BuidFactory(c)
		result, _, err := crud.ElkQuery(context.TODO(), qParam)
		// Trả về kết quả dưới dạng JSON
		c.JSON(http.StatusOK, gin.H{"result": result, "err ": err}) // Trả về kết quả dưới dạng JSON
	})
	router.GET("/my-org-elk", func(c *gin.Context) {
		// Thực hiện các xử lý cần thiết để lấy kết quả
		qParam := crud.BuidFactory(c)
		result, _, err := crud.ElkOrgQuery(context.TODO(), qParam)
		// Trả về kết quả dưới dạng JSON
		c.JSON(http.StatusOK, gin.H{"result": result, "err ": err}) // Trả về kết quả dưới dạng JSON
	})
	// Khởi động máy chủ trên cổng 8080
	router.Run(":8080")
	//UpdateTime1()

}
