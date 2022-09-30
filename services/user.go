package services

import (
	"context"
	"fmt"
	"github.com/diegofsousa/lab_grpc/pb"
	"io"
	"log"
	"time"
)

//type UserServiceServer interface {
//	AddUser(context.Context, *pb.User) (*pb.User, error)
//	AddUserVerbose(ctx context.Context, in *pb.User, opts ...grpc.CallOption) (UserService_AddUserVerboseClient, error)
//	mustEmbedUnimplementedUserServiceServer()
//}

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	// Insert database
	fmt.Println(req.Name)

	return &pb.User{Id: "123", Name: req.GetName(), Email: req.GetEmail()}, nil
}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	stream.SendMsg(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.SendMsg(&pb.UserResultStream{
		Status: "Inserting",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.SendMsg(&pb.UserResultStream{
		Status: "Completed",
		User: &pb.User{
			Id:    "123",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	return nil
}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	var users []*pb.User

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}

		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
		})

		fmt.Println("Adding", req.GetName())
	}
}

func (*UserService) AddUsersStreamBoth(stream pb.UserService_AddUsersStreamBothServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving stream from the client: %v", err)
		}

		err = stream.Send(&pb.UserResultStream{
			Status: "Added",
			User:   req,
		})

		if err != nil {
			log.Fatalf("Error sending stream for the client: %v", err)
		}
	}
}
