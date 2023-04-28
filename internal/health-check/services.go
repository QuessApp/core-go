package healthcheck

import (
	"log"

	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/internal/auth"
	"github.com/quessapp/core-go/internal/questions"
	"github.com/quessapp/core-go/internal/users"
	"github.com/quessapp/core-go/tests/mocks"
)

// Run is a function to run health check services.
// This function will create a fake user, fake question, and delete them.
func Run(handlerCtx *configs.HandlersCtx, authRepository *auth.AuthRepository, questionsRepository *questions.QuestionsRepository, usersRepository *users.UsersRepository) error {
	fakeUser1 := mocks.NewUserMock()
	fakeUser2 := mocks.NewUserMock()

	fakeUser1.Name = "USER_CREATED_BY_HEALTH_CHECK"
	fakeUser2.Name = "USER_CREATED_BY_HEALTH_CHECK"

	// we believe that health check endpoints is not only to return a 200 status code,
	// but also to create real data in the database and check if the data is created successfully.
	firstUser, err := authRepository.SignUp(&auth.SignUpUserDTO{
		Email:    fakeUser1.Email,
		Password: fakeUser1.Password,
		Nick:     fakeUser1.Nick,
		Name:     fakeUser1.Name,
		Locale:   fakeUser1.Locale,
	})

	if err != nil {
		log.Fatalf("Error when creating user %s for health check: %s \n", firstUser.ID, err)
		return err
	}

	secondUser, err := authRepository.SignUp(&auth.SignUpUserDTO{
		Email:    fakeUser2.Email,
		Password: fakeUser2.Password,
		Nick:     fakeUser2.Nick,
		Name:     fakeUser2.Name,
		Locale:   fakeUser2.Locale,
	})

	if err != nil {
		log.Fatalf("Error when creating user %s for health check: %s \n", secondUser.ID, err)
		return err
	}

	fakeQuestion := mocks.NewQuestionMock()

	if err = questionsRepository.Create(&questions.CreateQuestionDTO{
		Content:     fakeQuestion.Content,
		SendTo:      secondUser.ID,
		SentBy:      firstUser.ID,
		IsAnonymous: false,
	}); err != nil {
		log.Fatalf("Error when creating question for health check: %s \n", err)
		return err
	}

	var page int64 = 1
	var sort string = "desc"
	var filter string = "sent"

	q, err := questionsRepository.GetAll(&page, &sort, &filter, firstUser.ID)

	if err != nil {
		log.Fatalf("Error when listing questions created by user %s for health check: %s \n", firstUser.ID, err)
		return err
	}

	questions := *q.Questions

	if err := questionsRepository.Delete(questions[0].ID); err != nil {
		log.Fatalf("Error when deleting question %s for health check: %s \n", questions[0].ID, err)
		return err
	}

	if err := usersRepository.Delete(firstUser.ID); err != nil {
		log.Fatalf("Error when deleting user %s for health check: %s \n", firstUser.ID, err)
		return err
	}

	if err := usersRepository.Delete(secondUser.ID); err != nil {
		log.Fatalf("Error when deleting user %s for health check: %s \n", secondUser.ID, err)
		return err
	}

	return nil
}
