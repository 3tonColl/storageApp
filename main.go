package main

import (
	/*"fmt"
	"log"
	"math"
	"net/http"
	"strconv" */

	"main/handler"
	"main/storage"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	memoryStorage := storage.NewMemoryStorage()
	handler := handler.NewHandler(memoryStorage)

	router := gin.Default()

	// Путь к директории с изображениями
	tmplPath := filepath.Join("img")

	// Настройка маршрута для статических файлов

	router.StaticFS("/static/", http.Dir(tmplPath))
	router.POST("/employee", handler.CreateEmployee)
	router.GET("/", handler.LebonPage)
	router.GET("/employee/:id", handler.GetEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)

	router.Run()

	/*
			type LoggingMiddleware struct {
				handler http.Handler
			}

			func NewLoggingMiddleware(handler http.Handler) *LoggingMiddleware {
				return &LoggingMiddleware{
					handler: handler,
				}
			}

			func (l *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
				log.Println(req.URL, "requested")

				l.handler.ServeHTTP(w, req)
			}

			func Haa(w http.ResponseWriter, req *http.Request) {
				_, err := w.Write([]byte("hello!!!"))
				if err != nil {
					log.Println(err)
				}

		}

			func RadiusFunc(w http.ResponseWriter, req *http.Request) {
				rad, _ := strconv.ParseFloat(req.URL.Query().Get("radius"), 64)
				answ := rad * math.Pi * rad
				_, err := w.Write([]byte(fmt.Sprint(answ)))
				if err != nil {
					log.Println(err)
				}
			}
	*/

	/*
		r := http.NewServeMux()
		r.HandleFunc("/hello", Haa)
		r.HandleFunc("/zalupa", func(w http.ResponseWriter, req *http.Request) {
			_, err := w.Write([]byte("hui"))
			if err != nil {
				log.Println(err)
			}	ццк
		})
		r.HandleFunc("/face", func(w http.ResponseWriter, req *http.Request) {
			_, err := w.Write([]byte(":)"))
			if err != nil {
				log.Println(err)
			}

		})
		r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			_, err := w.Write([]byte("База"))
			if err != nil {
				log.Println(err)
			}

		})
		r.HandleFunc("/echo", func(w http.ResponseWriter, req *http.Request) {
			_, err := fmt.Fprint(w, req.URL.Query().Get("message"))
			if err != nil {
				log.Println(err)
			}

		})
		r.HandleFunc("/radius", RadiusFunc)

		log.Println(http.ListenAndServe(":8080", NewLoggingMiddleware(r)))
	*/

}
