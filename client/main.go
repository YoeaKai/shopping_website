// $ go run main.go iphone
package main

import (
	"context"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"sort"

	pb "shopping_website/product"

	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

var tpl *template.Template

const session = "session"

type sortMethod string

const (
	lessFirst   sortMethod = "lessFirst"
	higherFirst sortMethod = "higherFirst"
)

type searchResult struct {
	Title  string
	Result []*pb.ProductResponse
}

var dbResult = map[string]*searchResult{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/", index)
	mux.POST("/search", search)
	http.ListenAndServe(":8080", mux)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// create session
	sID := uuid.NewV4()
	sess := &http.Cookie{
		Name:  "session",
		Value: sID.String(),
	}
	http.SetCookie(w, sess)

	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func search(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	keyWord := req.FormValue("keyWord")

	// Open the config.
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	// Parse config
	var config map[string]string
	if err := json.NewDecoder(jsonFile).Decode(&config); err != nil {
		log.Fatal(err)
	}

	// connect to GRPC service
	port := config["host_ip"] // [name of process of server container]:8081  e.g.: "server:8081"
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)
	getProductClient, err := client.GetProductInfo(context.Background(), &pb.ProductRequest{KeyWord: keyWord})
	if err != nil {
		log.Fatal(err)
	}

	var result []*pb.ProductResponse

	// receive
	for {
		reply, err := getProductClient.Recv()
		if err == io.EOF {
			log.Println("Done")
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("reply : %v\n", reply)
		result = append(result, reply)
	}

	sort.Slice(result, func(i, j int) bool { return result[i].Price < result[j].Price })

	data := &searchResult{
		Title:  "Search result: " + keyWord,
		Result: result,
	}

	err = tpl.ExecuteTemplate(w, "result.gohtml", data)
	HandleError(w, err)

	sess, err := req.Cookie(session)
	if err != nil {
		log.Fatalln(err)
	}

	dbResult[sess.Value] = data
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
