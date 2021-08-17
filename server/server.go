package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"tcmedia-grpc.study.com/aws"
	"tcmedia-grpc.study.com/config"
	"tcmedia-grpc.study.com/pb"
)

type Server struct {
	s3Client aws.S3Client
	pb.UnimplementedS3ServiceServer
}

func New(s3Client aws.S3Client) *Server {
	return &Server{s3Client: s3Client}
}

func (s *Server) UploadFile(stream pb.S3Service_UploadFileServer) error {

	fileData := bytes.Buffer{}
	fileSize := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return err
		}

		chunk := req.GetChunkData()
		size := len(chunk)
		fileSize += size

		log.Printf("received a chunk with size: %d", size)

		_, err = fileData.Write(chunk)
		if err != nil {
			return err
		}
	}

	reader := bytes.NewReader(fileData.Bytes())

	fileURL, err := s.s3Client.Upload("test.mp4", reader)
	if err != nil {
		return err
	}

	res := &pb.UploadFileResponse{
		Url:  fileURL,
		Size: uint32(fileSize),
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return err
	}

	log.Printf("saved file with url: %s, size: %d", fileURL, fileSize)
	return nil

}

func main() {
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("error while loading envs: %v", err)
	}

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("error while listening to port 50051: %v", err)
	}

	s3Client, err := aws.NewS3Client(env.AccessKey, env.SecretKey, env.Region, env.Bucket)
	if err != nil {
		log.Fatalf("could not connect to aws s3: %v", err)
	}

	server := New(s3Client)

	fmt.Println("Server Start on 0.0.0.0:50051")
	fmt.Println("Wainting requests...")

	s := grpc.NewServer()
	pb.RegisterS3ServiceServer(s, server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("error while serving: %v", err)
	}
}
