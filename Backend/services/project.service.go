package service

import (
	"fmt"
	m "mongoCrud/models"
	repository "mongoCrud/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProject(p *m.Proyecto, id string) error {
	//crear el proyecto en el Id
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	fmt.Println("crear el proyecto en el Id ", userId)
	// enviar a a bd
	if err := repository.CreateProject(p, userId); err != nil {
		return err
	}
	return nil

}

func GetMyProjects(id string) (m.Proyectos, error) {

	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	ps, err := repository.GetMyProjects(userId)
	if err != nil {
		return nil, err
	}
	return ps, nil

}

func GetProjectsWorkingOn(id string) (m.Proyectos, error) {
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	ps, err := repository.GetProjectsWorkingOn(userId)
	if err != nil {
		return nil, err
	}
	return ps, nil

}
func AddWorker(uid string, pid string) error {
	projectId, err := primitive.ObjectIDFromHex(pid)
	if err != nil {
		return err
	}
	userId, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		return err
	}

	if err = repository.AddWorker(userId, projectId); err != nil {
		return err
	}
	return nil

}
