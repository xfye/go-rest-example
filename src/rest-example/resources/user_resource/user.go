package user_resource

import (
	"github.com/emicklei/go-restful"
	//"go.etcd.io/etcd/client"
	//"fmt"
	"net/http"
	"rest-example/database"
	"rest-example/database/models"
	"strconv"
)

var g_dbop *database.DBOperator = nil

func New(dbop *database.DBOperator) *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/users").
		Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	service.Route(service.POST("").To(CreateUser))
	service.Route(service.GET("/{userId}").To(GetUser))
	service.Route(service.DELETE("/{userId}").To(DeleteUser))

	g_dbop = dbop
	return service
}

func CreateUser(request *restful.Request, response *restful.Response) {
	user := new(models.User)
	err := request.ReadEntity(&user)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}
	g_dbop.DB.Create(&user)
	response.WriteEntity(user)
}

func GetUser(request *restful.Request, response *restful.Response) {
	userId := request.PathParameter("userId")
	id, _ := strconv.Atoi(userId)

	var user models.User
	g_dbop.DB.First(&user, id)

	response.WriteEntity(user)
}

func DeleteUser(request *restful.Request, response *restful.Response) {
	userId, _ := strconv.Atoi(request.PathParameter("userId"))
	g_dbop.DB.Delete(models.User{}, "id=?", userId)
}
