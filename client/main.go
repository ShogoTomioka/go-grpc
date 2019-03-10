package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	pb "github.com/ShogoTomioka/grpc/proto"

	"google.golang.org/grpc"
)

func main() {
	//sampleなのでwithInsecure
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing conntact massage")
		os.Exit(1)
	}
	contactCat := flag.Arg(0)

	// Clientの作成
	client := pb.NewCatClient(conn)
	// Serverに問い合わせるCatMessageを作成
	message := &pb.GetMyCatMessage{
		TargetCat: contactCat,
	}
	//空のContextと問い合わせ内容を渡す
	res, err := client.GetMyCat(context.TODO(), message)
	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
