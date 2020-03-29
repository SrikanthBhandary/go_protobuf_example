package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "github.com/srikanthbhandary/go_protobuf_example/proto"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func printRecord(w io.Writer, p *pb.Person) {
	fmt.Fprintln(w, "Person ID:", p.Id)
	fmt.Fprintln(w, "  Name:", p.Name)
	if p.Email != "" {
		fmt.Fprintln(w, "  E-mail address:", p.Email)
	}

	for _, pn := range p.Phones {
		switch pn.Type {
		case pb.Person_MOBILE:
			fmt.Fprint(w, "  Mobile phone #: ")
		case pb.Person_HOME:
			fmt.Fprint(w, "  Home phone #: ")
		case pb.Person_WORK:
			fmt.Fprint(w, "  Work phone #: ")
		}
		fmt.Fprintln(w, pn.Number)
	}
}

func listPeopleRecords(w io.Writer, book *pb.AddressBook) {
	for _, p := range book.People {
		printRecord(w, p)
	}
}

func main() {
	http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		message, _ := ioutil.ReadAll(r.Body)
		book := &pb.AddressBook{}
		if err := proto.Unmarshal(message, book); err != nil {
			log.Fatalln("Failed to parse address book:", err)
		}
		listPeopleRecords(os.Stdout, book)
	})
	log.Fatal(http.ListenAndServe(":8888", nil))
}
