package grpcserver

import (
	"app/database"
	"app/user"
)

type GrpcServer struct {
	DBHandler *database.GORMHandler
}

func NewGrpcServer() (*GrpcServer, error) {
	connection, err := database.CreateConnection()
	if err != nil {
		return nil, err
	}

	return &GrpcServer{
		DBHandler: connection,
	}, nil
}

func (server *GrpcServer) GetPeople(r *user.Request, stream user.UserService_GetPeopleServer) error {
	people, err := server.DBHandler.GetPeople()
	if err != nil {
		return err
	}

	for _, person := range people {
		grpcPerson := convertToGrpcPerson(person)
		err := stream.Send(grpcPerson)
		if err != nil {
			return err
		}
	}
	return nil
}

func convertToGrpcPerson(u database.User) *user.User {
	return &user.User{
		Id:     int32(u.Id),
		Name:   u.Name,
		Family: u.Family,
	}
}
