package user_service

import models "zetting/pkg/user/core/models"

func (s *userService) GetUsersByProject(projectId interface{}, document string) (models.Users, error) {

	users, err := s.userRepo.GetUsersByProject(projectId, document)
	if err != nil {
		return nil, err
	}
	return users, nil
}
