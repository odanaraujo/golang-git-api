package repository

import (
	"github.com/odanaraujo/golang/users-api/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"os"
	"testing"
)

var (
	databaseName   = "user_database_test"
	collectionName = "user_collection_test"
)

func TestCreateUser(t *testing.T) {

	os.Setenv(MONGODB_USER_COLLECTION, collectionName)
	defer os.Clearenv()

	// cria um banco de dados mockado com as opções setadas para que ele crie esse banco mockado
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("when sending a valid domain return success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)
		user := model.NewUSerDomain("test", "test@gmail.com", "test@test", 34)
		userDomainInterface, err := repo.CreateUser(model.NewUSerDomain("test", "test@gmail.com", "test@test", 34))

		_, errID := primitive.ObjectIDFromHex(userDomainInterface.GetID())

		assert.Nil(t, err)
		assert.Nil(t, errID)

		assert.EqualValues(t, userDomainInterface.GetName(), user.GetName())
		assert.EqualValues(t, userDomainInterface.GetEmail(), user.GetEmail())
		assert.EqualValues(t, userDomainInterface.GetPassword(), user.GetPassword())
		assert.EqualValues(t, userDomainInterface.GetAge(), user.GetAge())
	})

	mt.Run("return error from database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)
		userDomainInterface, err := repo.CreateUser(model.NewUSerDomain("test", "test@gmail.com", "test@test", 34))

		assert.NotNil(t, err)
		assert.Nil(t, userDomainInterface)
	})
}
