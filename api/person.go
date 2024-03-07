package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"regexp"
	"sqlBoiler/internal/db/models"
	"sqlBoiler/util"
)

func InitDB() (*sqlx.DB, error) {
	config, err := util.LoadConfig(".")
	db, err := sqlx.Connect(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Невозможно создать соединение с базой данных ", err)
	}
	return db, nil
}

func GetAllPersons(c *gin.Context) {
	var persons []models.Person
	db, err := InitDB()
	if err != nil {
		panic(err)
	}
	err = db.Select(&persons, models.SelectQuery)
	if err != nil {
		err := errors.Wrap(err, "ERROR: error getAllPersons")
		log.Printf("%+v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении запроса получения данных из базы"})
		return
	}
	c.JSON(http.StatusOK, persons)
}

func GetPersonByIIN(c *gin.Context) {
	var person models.Person
	iin := c.Param("iin")
	if isValidIINFormat(iin) != true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат ИИН"})
		return
	}

	db, err := InitDB()
	if err != nil {
		log.Println("ERROR: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка инициализации базы данных"})
		return
	}
	err = db.Get(&person, models.SelectQueryByIIN, iin)
	if err != nil {
		err := errors.Wrap(err, "ERROR: error getPersonByIIN")
		log.Printf("%+v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Запись не найдена"})
		return
	}
	c.JSON(http.StatusOK, person)
}

func isValidIINFormat(iin string) bool {
	if len(iin) != 12 {
		return false
	}
	match, _ := regexp.MatchString("^[0-9]+$", iin)
	if !match {
		return false
	}
	return true
}

func CreatePerson(c *gin.Context) {

	db, err := InitDB()
	if err != nil {
		panic(err)
	}
	var newPerson models.NewPerson
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		err := errors.Wrap(err, "ERROR: error create person")
		log.Printf("%+v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат входных данных"})
		return
	}
	fields := prepareFields(c, newPerson)
	_, err = db.Exec(models.InsertQuery, fields...)
	if err != nil {
		err := errors.Wrap(err, "ERROR: error create person")
		log.Printf("%+v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении запроса добавления данных в базу"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Запись успешно добавлена"})
}

func prepareFields(c *gin.Context, newPerson models.NewPerson) []interface{} {
	return []interface{}{
		marshalField(c, newPerson.FirstName),
		marshalField(c, newPerson.LastName),
		marshalField(c, newPerson.MiddleName),
		newPerson.IIN,
		newPerson.RNN,
		newPerson.BirthDate,
		newPerson.DeathDate,
		newPerson.Resident,
		newPerson.PhoneNumber,
		marshalField(c, newPerson.Asp),
		marshalField(c, newPerson.Esp),
		marshalField(c, newPerson.Accident),
		marshalField(c, newPerson.AvgIncome),
		marshalField(c, newPerson.AvgPensionIncome),
		marshalField(c, newPerson.AvgSocialIncome),
		marshalField(c, newPerson.BirthInfo),
		marshalField(c, newPerson.Citizenship),
		marshalField(c, newPerson.Education),
		marshalField(c, newPerson.Health),
		marshalField(c, newPerson.FarmAnimal),
		marshalField(c, newPerson.Scoring),
		marshalField(c, newPerson.FinancingTerrExtrList),
		marshalField(c, newPerson.Opv),
		marshalField(c, newPerson.Gender),
		marshalField(c, newPerson.Nationality),
		marshalField(c, newPerson.Income),
		marshalField(c, newPerson.Job),
		marshalField(c, newPerson.IncomeRefund),
		marshalField(c, newPerson.CriminalRecord),
		marshalField(c, newPerson.Kdn),
		marshalField(c, newPerson.TaxNotification),
		marshalField(c, newPerson.NarcoRegistry),
		marshalField(c, newPerson.Photo),
		marshalField(c, newPerson.Bmg),
		marshalField(c, newPerson.PsychoRegistry),
		marshalField(c, newPerson.RegistrationAddress),
		marshalField(c, newPerson.TubRegistry),
		marshalField(c, newPerson.VaccinationKz),
		marshalField(c, newPerson.RealEstateObjectRegistration),
		newPerson.PortalRefreshDate,
		marshalField(c, newPerson.Wanted),
		marshalField(c, newPerson.RealEstateQueue),
		marshalField(c, newPerson.RealEstateObjectEncumbrance),
		marshalField(c, newPerson.BankruptcyApplication),
		marshalField(c, newPerson.DebtorMu),
		newPerson.UnemployedReg,
		newPerson.UnemployedRegRefreshDate,
	}
}
func marshalField(c *gin.Context, field interface{}) interface{} {
	jsonField, err := json.Marshal(field)
	if err != nil {
		handleJSONError(c, err)
	}
	return jsonField
}

func handleJSONError(c *gin.Context, err error) {
	log.Println(err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при преобразовании данных в JSON"})
	return
}

func handleDBError(c *gin.Context, err error) {
	log.Printf("Database error: %+v\n", err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении запроса в базу данных"})
}
