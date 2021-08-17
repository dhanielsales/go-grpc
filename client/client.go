package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
	"tcmedia-grpc.study.com/pb"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := pb.NewS3ServiceClient(cc)

	sendImage(c)

}

func sendImage(client pb.S3ServiceClient) {
	file, err := os.Open("../files/placeholder.mp4")
	// file, err := os.Open("../files/teste.png")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}

	defer file.Close()

	stream, err := client.UploadFile(context.Background())
	if err != nil {
		log.Fatalf("could not upload file: %v", err)
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read chunk: %v", err)
		}

		req := &pb.UploadFileRequest{
			Data: &pb.UploadFileRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Fatalf("could not send chunk to server: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}

	log.Printf("File uploaded with url: %s, size: %d", res.GetUrl(), res.GetSize())
}
