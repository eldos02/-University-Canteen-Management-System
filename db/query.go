package db

import (
	"UniversityCanteenManagementSystem/pkg/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// Создаем нового пользователя
func CreateUser(user models.User, requests *gin.Context) {
	//При форс мажорах узнаем в чем ошибка
	if err := database.Create(&user).Error; err != nil {
		requests.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
		return
	}
}

// Проверяем наличие юзера в базе данных
func CheckUser(creds models.Credentials, requests *gin.Context) {
	var user models.User
	//Для откладки
	log.Println("Checking user with email:", creds.Email)

	// Ищем юзера по его емайлу
	if result := database.Where("email = ?", creds.Email).First(&user); result.Error != nil {
		log.Println("User not found:", result.Error)
		requests.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	//Хэшируем пароль и проверяем его правильность
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		log.Println("Password comparison failed:", err)
		requests.JSON(http.StatusUnauthorized, gin.H{"error": "Неправильный пароль ты выбрал дружок!"})
		return
	}

	//После успешного захода сохроняем сессию
	session := sessions.Default(requests)
	session.Set("user_id", user.ID)
	if err := session.Save(); err != nil {
		log.Println("Сессия не сохранилась", err)
		requests.JSON(http.StatusInternalServerError, gin.H{"error": "Провал при сохранений сессий"})
		return
	}
	//Для проверки добавил логи
	log.Println("Юзер успешно авторизовался!")

	//Перенаправляем юзера
	//Если админ - в админску страницу. Если обычный юзер то в Dashboard
	if user.IsAdmin {
		requests.Redirect(http.StatusFound, "/adminPage")
	} else {
		requests.Redirect(http.StatusFound, "/")
	}
}

// Просто находим юзера по его айди
func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	// Отправка запроса в базу данных и загрузка первой найденной записи в `user`
	result := database.Where("id = ?", id).First(&user)
	if result.Error != nil {
		// Возвращаем nil и ошибку, если пользователь не найден или произошла другая ошибка
		return nil, result.Error
	}
	// Возвращаем найденного пользователя и nil, если пользователь найден успешно
	return &user, nil
}
