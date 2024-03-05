package models

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"
)

const selectQuery = `
SELECT id,
--        first_name->>'first_name_ru' as first_name_ru,
       first_name,
       last_name,
       middle_name,
       iin,
       rnn,
       birth_date,
       death_date,
       resident,
       phone_number,
       asp,
       esp,
       accident,
       avg_income,
       avg_pension_income,
       avg_social_income,
       birth_info,
       citizenship,
       education,
       health,
       farm_animal,
       scoring,
       financing_terr_extr_list,
       opv,
       gender,
       nationality,
       income,
       job,
       income_refund,
       criminal_record,
       kdn,
       tax_notification,
       narco_registry,
       photo,
       bmg,
       psycho_registry,
       registration_address,
       tub_registry,
       vaccination_kz,
       real_estate_object_registration,
       portal_refresh_date,
       wanted,
       real_estate_queue,
       real_estate_object_encumbrance,
       bankruptcy_application,
       debtor_mu,
       unemployed_reg,
       unemployed_reg_refresh_date
FROM crm360_person`

const insertQuery = `
INSERT INTO crm360_person (
   	   first_name,
       last_name,
       middle_name,
       iin,
       rnn,
       birth_date,
       death_date,
       resident,
       phone_number,
       asp,
       esp,
       accident,
       avg_income,
       avg_pension_income,
       avg_social_income,
       birth_info,
       citizenship,
       education,
       health,
       farm_animal,
       scoring,
       financing_terr_extr_list,
       opv,
       gender,
       nationality,
       income,
       job,
       income_refund,
       criminal_record,
       kdn,
       tax_notification,
       narco_registry,
       photo,
       bmg,
       psycho_registry,
       registration_address,
       tub_registry,
       vaccination_kz,
       real_estate_object_registration,
       portal_refresh_date,
       wanted,
       real_estate_queue,
       real_estate_object_encumbrance,
       bankruptcy_application,
       debtor_mu,
       unemployed_reg,
       unemployed_reg_refresh_date) --54
 VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
    $11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
    $21, $22, $23, $24, $25, $26, $27, $28, $29, $30,
    $31, $32, $33, $34, $35, $36, $37, $38, $39, $40,
    $41, $42, $43, $44, $45, $46, $47
);
--        first_name->>'first_name_ru' as first_name_ru,
`

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@127.0.0.1:5432/postgres?sslmode=disable"
)

func InitDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect(dbDriver, dbSource) // Измените "postgres" на нужный драйвер
	if err != nil {
		return nil, err
	}
	return db, nil
}
func GetAllPersons(c *gin.Context) {
	var persons []person
	db, err := InitDB()
	if err != nil {
		panic(err)
	}
	err = db.Select(&persons, selectQuery)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении запроса получения данных из базы"})
		return
	}

	c.JSON(http.StatusOK, persons)
}

func CreatePerson(c *gin.Context) {

	db, err := InitDB()
	if err != nil {
		panic(err)
	}
	var newPerson newPerson
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		log.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат входных данных"})
		return
	}

	// Преобразование структуры firstName в JSON
	firstNameJSON, err := json.Marshal(newPerson.FirstName)
	if err != nil {
		handleJSONError(c, err)
		return
	}

	lastNameJSON, err := json.Marshal(newPerson.LastName)
	if err != nil {
		handleJSONError(c, err)
		return
	}

	middleNameJSON, err := json.Marshal(newPerson.MiddleName)
	if err != nil {
		handleJSONError(c, err)
		return
	}

	_, err = db.Exec(insertQuery,
		firstNameJSON, lastNameJSON, middleNameJSON)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении запроса добавления данных в базу"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Запись успешно добавлена"})
}

func handleJSONError(c *gin.Context, err error) {
	log.Println(err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при преобразовании данных в JSON"})
}

// Структура для хранения данных из таблицы
type person struct {
	ID                           *int64           `json:"id" db:"id"`
	FirstName                    *json.RawMessage `json:"first_name" db:"first_name"`
	LastName                     *json.RawMessage `json:"last_name" db:"last_name"`
	MiddleName                   *json.RawMessage `json:"middle_name" db:"middle_name"`
	IIN                          *string          `json:"iin" db:"iin"`
	RNN                          *string          `json:"rnn" db:"rnn"`
	BirthDate                    *time.Time       `json:"birth_date" db:"birth_date"`
	DeathDate                    *time.Time       `json:"death_date" db:"death_date"`
	Resident                     *string          `json:"resident" db:"resident"`
	PhoneNumber                  *string          `json:"phone_number" db:"phone_number"`
	Asp                          *json.RawMessage `json:"asp" db:"asp"`
	Esp                          *json.RawMessage `json:"esp" db:"esp"`
	Accident                     *json.RawMessage `json:"accident" db:"accident"`
	AvgIncome                    *json.RawMessage `json:"avg_income" db:"avg_income"`
	AvgPensionIncome             *json.RawMessage `json:"avg_pension_income" db:"avg_pension_income"`
	AvgSocialIncome              *json.RawMessage `json:"avg_social_income" db:"avg_social_income"`
	BirthInfo                    *json.RawMessage `json:"birth_info" db:"birth_info"`
	Citizenship                  *json.RawMessage `json:"citizenship" db:"citizenship"`
	Education                    *json.RawMessage `json:"education" db:"education"`
	Health                       *json.RawMessage `json:"health" db:"health"`
	FarmAnimal                   *json.RawMessage `json:"farm_animal" db:"farm_animal"`
	Scoring                      *json.RawMessage `json:"scoring" db:"scoring"`
	FinancingTerrExtrList        *json.RawMessage `json:"financing_terr_extr_list" db:"financing_terr_extr_list"`
	Opv                          *json.RawMessage `json:"opv" db:"opv"`
	Gender                       *json.RawMessage `json:"gender" db:"gender"`
	Nationality                  *json.RawMessage `json:"nationality" db:"nationality"`
	Income                       *json.RawMessage `json:"income" db:"income"`
	Job                          *json.RawMessage `json:"job" db:"job"`
	IncomeRefund                 *json.RawMessage `json:"income_refund" db:"income_refund"`
	CriminalRecord               *json.RawMessage `json:"criminal_record" db:"criminal_record"`
	Kdn                          *json.RawMessage `json:"kdn" db:"kdn"`
	TaxNotification              *json.RawMessage `json:"tax_notification" db:"tax_notification"`
	NarcoRegistry                *json.RawMessage `json:"narco_registry" db:"narco_registry"`
	Photo                        *json.RawMessage `json:"photo" db:"photo"`
	Bmg                          *json.RawMessage `json:"bmg" db:"bmg"`
	PsychoRegistry               *json.RawMessage `json:"psycho_registry" db:"psycho_registry"`
	RegistrationAddress          *json.RawMessage `json:"registration_address" db:"registration_address"`
	TubRegistry                  *json.RawMessage `json:"tub_registry" db:"tub_registry"`
	VaccinationKz                *json.RawMessage `json:"vaccination_kz" db:"vaccination_kz"`
	RealEstateObjectRegistration *json.RawMessage `json:"real_estate_object_registration" db:"real_estate_object_registration"`
	PortalRefreshDate            *time.Time       `json:"portal_refresh_date" db:"portal_refresh_date"`
	Wanted                       *json.RawMessage `json:"wanted" db:"wanted"`
	RealEstateQueue              *json.RawMessage `json:"real_estate_queue" db:"real_estate_queue"`
	RealEstateObjectEncumbrance  *json.RawMessage `json:"real_estate_object_encumbrance" db:"real_estate_object_encumbrance"`
	BankruptcyApplication        *json.RawMessage `json:"bankruptcy_application" db:"bankruptcy_application"`
	DebtorMu                     *json.RawMessage `json:"debtor_mu" db:"debtor_mu"`
	UnemployedReg                *string          `json:"unemployed_reg" db:"unemployed_reg"`
	UnemployedRegRefreshDate     *time.Time       `json:"unemployed_reg_refresh_date" db:"unemployed_reg_refresh_date"`
}

type newPerson struct {
	FirstName                    firstName                    `json:"first_name" db:"first_name"`
	LastName                     lastName                     `json:"last_name" db:"last_name"`
	MiddleName                   middleName                   `json:"middle_name" db:"middle_name"`
	IIN                          string                       `json:"iin" db:"iin"`
	RNN                          string                       `json:"rnn" db:"rnn"`
	BirthDate                    time.Time                    `json:"birth_date" db:"birth_date"`
	DeathDate                    time.Time                    `json:"death_date" db:"death_date"`
	Resident                     string                       `json:"resident" db:"resident"`
	PhoneNumber                  string                       `json:"phone_number" db:"phone_number"`
	Asp                          asp                          `json:"asp" db:"asp"`
	Esp                          esp                          `json:"esp" db:"esp"`
	Accident                     accident                     `json:"accident" db:"accident"`
	AvgIncome                    avgIncome                    `json:"avg_income" db:"avg_income"`
	AvgPensionIncome             avgPensionIncome             `json:"avg_pension_income" db:"avg_pension_income"`
	AvgSocialIncome              avgSocialIncome              `json:"avg_social_income" db:"avg_social_income"`
	BirthInfo                    birthInfo                    `json:"birth_info" db:"birth_info"`
	Citizenship                  citizenship                  `json:"citizenship" db:"citizenship"`
	Education                    education                    `json:"education" db:"education"`
	Health                       health                       `json:"health" db:"health"`
	FarmAnimal                   farmAnimal                   `json:"farm_animal" db:"farm_animal"`
	Scoring                      scoring                      `json:"scoring" db:"scoring"`
	FinancingTerrExtrList        financingTerrExtrList        `json:"financing_terr_extr_list" db:"financing_terr_extr_list"`
	Opv                          opv                          `json:"opv" db:"opv"`
	Gender                       gender                       `json:"gender" db:"gender"`
	Nationality                  nationality                  `json:"nationality" db:"nationality"`
	Income                       income                       `json:"income" db:"income"`
	Job                          job                          `json:"job" db:"job"`
	IncomeRefund                 incomeRefund                 `json:"income_refund" db:"income_refund"`
	CriminalRecord               criminalRecord               `json:"criminal_record" db:"criminal_record"`
	Kdn                          kdn                          `json:"kdn" db:"kdn"`
	TaxNotification              taxNotification              `json:"tax_notification" db:"tax_notification"`
	NarcoRegistry                narcoRegistry                `json:"narco_registry" db:"narco_registry"`
	Photo                        photo                        `json:"photo" db:"photo"`
	Bmg                          bmg                          `json:"bmg" db:"bmg"`
	PsychoRegistry               psychoRegistry               `json:"psycho_registry" db:"psycho_registry"`
	RegistrationAddress          registrationAddress          `json:"registration_address" db:"registration_address"`
	TubRegistry                  tubRegistry                  `json:"tub_registry" db:"tub_registry"`
	VaccinationKz                vaccinationKz                `json:"vaccination_kz" db:"vaccination_kz"`
	RealEstateObjectRegistration realEstateObjectRegistration `json:"real_estate_object_registration" db:"real_estate_object_registration"`
	PortalRefreshDate            time.Time                    `json:"portal_refresh_date" db:"portal_refresh_date"`
	Wanted                       wanted                       `json:"wanted" db:"wanted"`
	RealEstateQueue              realEstateQueue              `json:"real_estate_queue" db:"real_estate_queue"`
	RealEstateObjectEncumbrance  realEstateObjectEncumbrance  `json:"real_estate_object_encumbrance" db:"real_estate_object_encumbrance"`
	BankruptcyApplication        bankruptcyApplication        `json:"bankruptcy_application" db:"bankruptcy_application"`
	DebtorMu                     debtorMu                     `json:"debtor_mu" db:"debtor_mu"`
	UnemployedReg                string                       `json:"unemployed_reg" db:"unemployed_reg"`
	UnemployedRegRefreshDate     time.Time                    `json:"unemployed_reg_refresh_date" db:"unemployed_reg_refresh_date"`
}

type firstName struct {
	FirstNameRu string `json:"first_name_ru"`
	FirstNameKz string `json:"first_name_kz"`
	FirstNameEn string `json:"first_name_en"`
}

type lastName struct {
	LastNameRu string `json:"last_name_ru"`
	LastNameKz string `json:"last_name_kz"`
	LastNameEn string `json:"last_name_en"`
}

type middleName struct {
	MiddleNameRu string `json:"middle_name_ru"`
	MiddleNameKz string `json:"middle_name_kz"`
	MiddleNameEn string `json:"middle_name_en"`
}

type accident struct {
	AccidentName       string    `json:"accident_name"`
	AccidentType       string    `json:"accident_type"`
	AccidentDate       time.Time `json:"accident_date"`
	AccidentDamageType string    `json:"accident_damage_type"`
}

type asp struct {
	AspAmount               float64   `json:"asp_amount"`
	AspDateFrom             time.Time `json:"asp_date_from"`
	AspDateTo               time.Time `json:"asp_date_to"`
	AspNumberOfChildren     int64     `json:"asp_number_of_children"`
	AspNumberOfFamilyMember int64     `json:"asp_number_of_family_member"`
}

type esp struct {
	EspPaymentAmount float64   `json:"esp_payment_amount"`
	EspPaymentDate   time.Time `json:"esp_payment_date"`
}

type avgIncome struct {
	AvgIncomeAmount          float64   `json:"avg_income_amount"`
	AvgIncomePeriod          time.Time `json:"avg_income_period"`
	AvgIncomeType            int64     `json:"avg_income_type"`
	AvgIncomePaymentCount    int64     `json:"avg_income_payment_count"`
	AvgIncomeCalculationType int64     `json:"avg_income_calculation_type"`
	AvgIncomeRefreshDate     time.Time `json:"avg_income_refresh_date"`
}

type avgPensionIncome struct {
	AvgPensionIncomeAmount         float64   `json:"avg_pension_income_amount"`
	AvgPensionIncomePeriod         time.Time `json:"avg_pension_income_period"`
	AvgPensionIncomeType           int64     `json:"avg_pension_income_type"`
	AvgPensionIncomePaymentCount   int64     `json:"avg_pension_income_payment_count"`
	AvgPensionIncomeAssignmentDate int64     `json:"avg_pension_income_calculation_type"`
}

type avgSocialIncome struct {
	AvgSocialAmount       float64   `json:"avg_social_income_amount"`
	AvgSocialPeriod       time.Time `json:"avg_social_income_period"`
	AvgSocialType         int64     `json:"avg_social_income_type"`
	AvgSocialPaymentCount int64     `json:"avg_social_income_payment_count"`
}

type birthInfo struct {
	BirthInfoCountryCode    string `json:"birth_info_country_code"`
	BirthInfoCountryNameKz  string `json:"birth_info_country_name_kz"`
	BirthInfoCountryNameRu  string `json:"birth_info_country_name_ru"`
	BirthInfoRegionCode     string `json:"birth_info_region_code"`
	BirthInfoRegionNameKz   string `json:"birth_info_region_name_kz"`
	BirthInfoRegionNameRu   string `json:"birth_info_region_name_ru"`
	BirthInfoDistrictCode   string `json:"birth_info_district_code"`
	BirthInfoDistrictNameKz string `json:"birth_info_district_name_kz"`
	BirthInfoDistrictNameRu string `json:"birth_info_district_name_ru"`
	BirthInfoBirthCity      string `json:"birth_info_birth_city"`
}
type citizenship struct {
	CitizenshipCode   string `json:"citizenship_code"`
	CitizenshipNameKZ string `json:"citizenship_name_kz"`
	CitizenshipNameRu string `json:"citizenship_name_ru"`
}
type educationStudentCategory struct {
	EducationStudentCategoryCode   string `json:"education_student_category_code"`
	EducationStudentCategoryNameRu string `json:"education_student_category_name_ru"`
	EducationStudentCategoryNameKz string `json:"education_student_category_name_kz"`
}
type educationCourse struct {
	EducationCourseNumber string `json:"education_course_number"`
	EducationCourseNameKz string `json:"education_course_name_kz"`
	EducationCourseNameRu string `json:"education_course_name_ru"`
}
type educationDiplomaType struct {
	EducationDiplomaTypeCode   string `json:"education_diploma_type_code"`
	EducationDiplomaTypeNameKz string `json:"education_diploma_type_name_kz"`
	EducationDiplomaTypeNameRu string `json:"education_diploma_type_name_ru"`
}
type educationHeRetirement struct {
	EducationHeRetirementCode   string `json:"education_he_retirement_code"`
	EducationHeRetirementNameKz string `json:"education_he_retirement_name_kz"`
	EducationHeRetirementNameRu string `json:"education_he_retirement_name_ru"`
}
type educationHeSpeciality struct {
	EducationHeSpecialityCode   string `json:"education_he_speciality_code"`
	EducationHeSpecialityNameKz string `json:"education_he_speciality_name_kz"`
	EducationHeSpecialityNameRu string `json:"education_he_speciality_name_ru"`
}
type educationInstitute struct {
	EducationInstituteBIN    string `json:"education_institute_bin"`
	EducationInstituteNameKz string `json:"education_institute_name_kz"`
	EducationInstituteNameRu string `json:"education_institute_name_ru"`
}
type educationLanguage struct {
	EducationLanguageCode   string `json:"education_language_code"`
	EducationLanguageNameKz string `json:"education_language_name_kz"`
	EducationLanguageNameRu string `json:"education_language_name_ru"`
}
type educationStatus struct {
	EducationStatusCode   int64  `json:"education_status_code"`
	EducationStatusNameKz string `json:"education_status_name_kz"`
	EducationStatusNameRu string `json:"education_status_name_ru"`
}
type educationStudyForm struct {
	EducationStudyFormCode   string `json:"education_study_form_code"`
	EducationStudyFormNameKz string `json:"education_study_form_name_kz"`
	EducationStudyFormNameRu string `json:"education_study_form_name_ru"`
}
type educationTpeSpeciality struct {
	EducationTpeSpecialityCode   string `json:"education_tpe_speciality_code"`
	EducationTpeSpecialityNameKz string `json:"education_tpe_speciality_name_kz"`
	EducationTpeSpecialityNameRu string `json:"education_tpe_speciality_name_ru"`
}
type education struct {
	EducationIdType              int64                    `json:"education_id_type"`
	EducationDateFrom            time.Time                `json:"education_date_from_from"`
	EducationDateTo              time.Time                `json:"education_date_to"`
	EducationCountryNameRu       string                   `json:"education_country_name_ru"`
	EducationTpeRetirementReason string                   `json:"education_tpe_retirement_reason"`
	EducationStudentCatego       educationStudentCategory `json:"education_student_category"`
	EducationCourse              educationCourse          `json:"education_course"`
	EducationDiplomaType         educationDiplomaType     `json:"education_diploma_type"`
	EducationHeRetirement        educationHeRetirement    `json:"education_he_retirement"`
	EducationHeSpeciality        educationHeSpeciality    `json:"education_he_speciality"`
	EducationInstitute           educationInstitute       `json:"education_institute"`
	EducationLanguage            educationLanguage        `json:"education_language"`
	EducationStatus              educationStatus          `json:"education_status"`
	EducationStudyForm           educationStudyForm       `json:"education_study_form"`
	EducationTpeSpeciality       educationTpeSpeciality   `json:"education_tpe_speciality"`
}

type pcrTest struct {
	PcrTestFirstName          string    `json:"pcr_test_first_name"`
	PcrTestLastName           string    `json:"pcr_test_last_name"`
	PcrTestMiddleName         string    `json:"pcr_test_middle_name"`
	PcrTestIIN                string    `json:"pcr_test_iin"`
	PcrTestGender             string    `json:"pcr_test_gender"`
	PcrTestHasSymptomsCovid   string    `json:"pcr_test_has_symptoms_covid"`
	PcrTestPlaceOfStudyOrWork string    `json:"pcr_test_place_of_study_or_work"`
	PcrTestResearchResult     string    `json:"pcr_test_research_result"`
	PcrTestPhone              string    `json:"pcr_test_phone"`
	PcrTestProtocolDate       time.Time `json:"pcr_test_protocol_date"`
	PcrTestFileLocation       string    `json:"pcr_test_file_location"`
	PcrTestLivingAddress      string    `json:"pcr_test_living_address"`
	PcrTestCitizenship        string    `json:"pcr_test_citizenship"`
	PcrTestAddress            string    `json:"pcr_test_address"`
	PcrTestResident           string    `json:"pcr_test_resident"`
}
type healthDisability struct { //Нет привязки к БТ
	HealthDisability            bool  `json:"health_disability"`
	HealthDisabilityGroupNumber int64 `json:"health_disability_group_number"`
}

type health struct {
	PcrTest          pcrTest          `json:"pcr_test"`
	HealthDisability healthDisability `json:"health_disability"`
}

type farmAnimal struct {
	FarmAnimalCount           int64     `json:"farm_animal_count"`
	FarmAnimalCato            string    `json:"farm_animal_cato"`
	FarmAnimalType            int64     `json:"farm_animal_type"`
	FarmAnimalTypeDescription string    `json:"farm_animal_type_description"`
	FarmAnimalRefreshDate     time.Time `json:"farm_animal_refresh_date"`
}

type bmlPdlScoring struct {
	BmlPdlScoringScore              float64   `json:"bml_pdl_scoring_score"`
	BmlPdlDefaultScoringProbability float64   `json:"bml_pdl_default_scoring_probability"`
	BmlPdlScoringDate               time.Time `json:"bml_pdl_scoring_date"`
}

type asiaCreditBankScoring struct {
	AsiaCreditBankScoringScore          float64   `json:"asia_credit_bank_scoring_score"`
	AsiaCreditBankScoringCreditTypeCode string    `json:"asia_credit_bank_scoring_credit_type_code"`
	AsiaCreditBankScoringDate           time.Time `json:"asia_credit_bank_scoring_date"`
}

type alfaBankScoring struct {
	AlfaBankScoringScore int64     `json:"alfa_bank_scoring_score"`
	AlfaBankScoringDate  time.Time `json:"alfa_bank_scoring_date"`
}

type behaviorScoring struct {
	BehaviorScoringDate                           time.Time `json:"behavior_scoring_date"`
	BehaviorScoringScore                          int64     `json:"behavior_scoring_score"`
	BehaviorScoringRiskClassNum                   int64     `json:"behavior_scoring_risk_class_num"`
	BehaviorScoringRiskClass                      string    `json:"behavior_scoring_risk_class"`
	BehaviorScoringMaxOverdueDaysCnt13m           int64     `json:"behavior_scoring_max_overdue_days_cnt_13_m"`
	BehaviorScoringYage                           int64     `json:"behavior_scoring_yage"`
	BehaviorScoringResultCuses                    string    `json:"behavior_scoring_result_cuses"`
	BehaviorScoringClosedHardOverdContrCntl60m    int64     `json:"behavior_scoring_closed_hard_overd_contr_cntl_60_m"`
	BehaviorScoringNoVerdRatiol12m                int64     `json:"behavior_scoring_no_verd_ratiol_12_m"`
	BehaviorScoringIsOverdAmountM1000             int64     `json:"behavior_scoring_is_overd_amount_m_1000"`
	BehaviorScoringContrLEndDateMCnt              int64     `json:"behavior_scoring_contr_l_end_date_m_cnt"`
	BehaviorScoringCredCardMaxutilizationl13m     int64     `json:"behavior_scoring_cred_card_maxutilizationl_13_m"`
	BehaviorScoringActiveContrMaxTerm             int64     `json:"behavior_scoring_active_contr_max_term"`
	BehaviorScoringReqCntl60m                     int64     `json:"behavior_scoring_req_cntl_60_m"`
	BehaviorScoringStartOverdMCnt                 int64     `json:"behavior_scoring_start_overd_m_cnt"`
	BehaviorScoringNoOverdMCnt                    int64     `json:"behavior_scoring_no_overd_m_cnt"`
	BehaviorScoringIsOverddaysCntBetween15And89lm int64     `json:"behavior_scoring_is_overddays_cnt_between_15_and_89_lm"`
	BehaviorScoringActContrCntl12m                int64     `json:"behavior_scoring_act_contr_cntl_12_m"`
	BehaviorScoringIsOverdmoreEq90lm              int64     `json:"behavior_scoring_is_overdmore_eq_90_lm"`
	BehaviorScoringCredCardOustandingAmount       int64     `json:"behavior_scoring_cred_card_oustanding_amount"`
}

type bckBankScoring struct {
	BckBankScoringDate                                  time.Time `json:"bck_bank_scoring_date"`
	BckBankScoringIsActCredCntMore5Flag                 int64     `json:"bck_bank_scoring_is_act_cred_cnt_more_5_flag"`
	BckBankScoringIsCurOverdAmountMoreEq5000Flag        int64     `json:"bck_bank_scoring_is_cur_overd_amount_more_eq_5000_flag"`
	BckBankScoringIsMaxOverdInstalCntMore60dIn12mFlag   int64     `json:"bck_bank_scoring_is_max_overd_instal_cnt_more_60_d_in_12_m_flag"`
	BckBankScoringIsMaxOverdInstalCntMoreEq90dIn23mFlag int64     `json:"bck_bank_scoring_is_max_overd_instal_cnt_more_eq_90_d_in_23_m_flag"`
	BckBankScoringIsReqCntMoreEq5In39dFlag              int64     `json:"bck_bank_scoring_is_req_cnt_more_eq_5_in_39_d_flag"`
}

type applicationScoring struct {
	ApplicationScoringDate         time.Time `json:"application_scoring_date"`
	ApplicationScoringScore        int64     `json:"application_scoring_score"`
	ApplicationCardScoringCardName string    `json:"application_card_scoring_card_name"`
}

type ficoScoring struct {
	FicoScoringDate  time.Time `json:"fico_scoring_date"`
	FicoScoringScore int64     `json:"fico_scoring_score"`
}

type fraudScoring struct {
	FraudScoringScore       int64     `json:"fraud_scoring_score"`
	FraudScoringRiskClass   string    `json:"fraud_scoring_risk_class"`
	FraudScoringProbability int64     `json:"fraud_scoring_probability"`
	FraudScoringRiskZone    string    `json:"fraud_scoring_risk_zone"`
	FraudScoringDate        time.Time `json:"fraud_scoring_date"`
}

type fastCashScoring struct {
	FastCashScoringDate                  time.Time `json:"fast_cash_scoring_date"`
	FastCashScoringIsCurOvdM15d          int64     `json:"fast_cash_scoring_is_cur_ovd_m_15d"`
	FastCashScoringIsCurOvdM5d           int64     `json:"fast_cash_scoring_is_cur_ovd_m_5d"`
	FastCashScoringIsCurOvdAmountM5000t  int64     `json:"fast_cash_scoring_is_cur_ovd_amount_m_5000t"`
	FastCashScoringIsOvdAmountM5000t13m  int64     `json:"fast_cash_scoring_is_ovd_amount_m_5000t_13m"`
	FastCashScoringIsPersonPledger       string    `json:"fast_cash_scoring_is_person_pledger"`
	FastCashScoringIsPledgePriceM500000t int64     `json:"fast_cash_scoring_is_pledge_price_m_500000t"`
	FastCashScoringIsReqCntM8l30d        int64     `json:"fast_cash_scoring_is_req_cnt_m_8_l30d"`
}

type freedomFinanceBankScoring struct {
	FreedomFinanceBankScoringDate                                time.Time `json:"freedom_finance_bank_scoring_date"`
	FreedomFinanceBankScoringIsCredPaymCntMEq3In6m               int64     `json:"freedom_finance_bank_scoring_is_cred_paym_cnt_m_eq_3_in_6_m"`
	FreedomFinanceBankScoringIsReqCntM5l30d                      int64     `json:"freedom_finance_bank_scoring_is_req_cnt_m_5_l30d"`
	FreedomFinanceBankScoringIsOvdCntMEq90dAndAmountM5000tL24m   int64     `json:"freedom_finance_bank_scoring_is_ovd_cnt_m_eq_90d_and_amount_m_5000t_l24m"`
	FreedomFinanceBankScoringIsMortgageContractBorrower          int64     `json:"freedom_finance_bank_scoring_is_mortgage_contract_borrower"`
	FreedomFinanceBankScoringIsMortgageContractOther             int64     `json:"freedom_finance_bank_scoring_is_mortgage_contract_other"`
	FreedomFinanceBankScoringIsOvdDaysCntM7OrOvdAmountMEq1000t   int64     `json:"freedom_finance_bank_scoring_is_ovd_days_cnt_m_7_or_ovd_amount_m_eq_1000t"`
	FreedomFinanceBankScoringIsOvdDaysM60dAndOvdAmountM5000tL12m int64     `json:"freedom_finance_bank_scoring_is_ovd_days_m_60d_and_ovd_amount_m_5000t_l12m"`
}

type forteBankScoring struct {
	ForteBankScoringDate                 time.Time `json:"forte_bank_scoring_date"`
	ForteBankScoringIsOvdDaysCntMEq1dl5d float64   `json:"forte_bank_scoring_is_ovd_days_cnt_m_eq_1d_l5d"`
	ForteBankScoringIsOvdDaysCntM30dl12m float64   `json:"forte_bank_scoring_is_ovd_days_cnt_m_30d_l12m"`
	ForteBankScoringReqCntL7d            float64   `json:"forte_bank_scoring_req_cnt_l7d"`
	ForteBankScoringIsOvdDaysCntM10dL3m  float64   `json:"forte_bank_scoring_is_ovd_days_cnt_m_10d_l3m"`
	ForteBankScoringIsOvdDaysCntM60dL36m float64   `json:"forte_bank_scoring_is_ovd_days_cnt_m_60d_l36m"`
	ForteBankScoringIsOvdDaysCntM90dL36m float64   `json:"forte_bank_scoring_is_ovd_days_cnt_m_90d_l36m"`
	ForteBankScoringIsOvdDaysCntM30dL6m  float64   `json:"forte_bank_scoring_is_ovd_days_cnt_m_30d_l6m"`
	ForteBankScoringReqCntL30d           float64   `json:"forte_bank_scoring_req_cnt_l30d"`
	ForteBankScoringReqCntL90d           float64   `json:"forte_bank_scoring_req_cnt_l90d"`
	ForteBankScoringNegStatusListL5y     float64   `json:"forte_bank_scoring_neg_status_list_l5y"`
	ForteBankScoringHasClosedContracts   float64   `json:"forte_bank_scoring_has_closed_contracts"`
}

type jysanBankScoring struct {
	JysanBankScoringDate                                      time.Time `json:"jysan_bank_scoring_date"`
	JysanBankScoringIsCurOvdCntMEq45d                         int64     `json:"jysan_bank_scoring_is_cur_ovd_cnt_m_eq_45d"`
	JysanBankScoringIsCurOvdDaysCntMEq1d                      int64     `json:"jysan_bank_scoring_is_cur_ovd_days_cnt_m_eq_1d"`
	JysanBankScoringIsOvdDaysCntMEq360dAndOvdAmountMEq150000t int64     `json:"jysan_bank_scoring_is_ovd_days_cnt_m_eq_360d_and_ovd_amount_m_eq_150000t"`
	JysanBankScoringIsCurOvdDaysCntMEq10d                     int64     `json:"jysan_bank_scoring_is_cur_ovd_days_cnt_m_eq_10d"`
	JysanBankScoringIsCurOvdDaysCntMEq30d                     int64     `json:"jysan_bank_scoring_is_cur_ovd_days_cnt_m_eq_30d"`
	JysanBankScoringIsCurOvdDaysCntMEq90d                     int64     `json:"jysan_bank_scoring_is_cur_ovd_days_cnt_m_eq_90d"`
	JysanBankScoringIsCurOvdAmountM0tAndOvdDaysLEq30d         int64     `json:"jysan_bank_scoring_is_cur_ovd_amount_m_0t_and_ovd_days_l_eq_30d"`
}

type kaspiBankScoring struct {
	KaspiBankScoringScore                                     float64   `json:"kaspi_bank_scoring_score"`
	KaspiBankScoringDate                                      time.Time `json:"kaspi_bank_scoring_date"`
	KaspiBankScoringCurMaxOvdDaysCnt                          int64     `json:"kaspi_bank_scoring_cur_max_ovd_days_cnt"`
	KaspiBankScoringCurOutstandingAmountWithOvd               float64   `json:"kaspi_bank_scoring_cur_outstanding_amount_with_ovd"`
	KaspiBankScoringClosedContractsCntL12m                    int64     `json:"kaspi_bank_scoring_closed_contracts_cnt_l12m"`
	KaspiBankScoringHasActiveContractsL12m                    float64   `json:"kaspi_bank_scoring_has_active_contracts_l12m"`
	KaspiBankScoringCurOutstandingAmount                      float64   `json:"kaspi_bank_scoring_cur_outstanding_amount"`
	KaspiBankScoringPaymentMonthCnt                           int64     `json:"kaspi_bank_scoring_payment_month_cnt"`
	KaspiBankScoringOverdueMonthCnt                           float64   `json:"kaspi_bank_scoring_overdue_month_cnt"`
	KaspiBankScoringProviderCnt                               int64     `json:"kaspi_bank_scoring_provider_cnt"`
	KaspiBankScoringActiveTotalAmount                         float64   `json:"kaspi_bank_scoring_active_total_amount"`
	KaspiBankScoringActiveContractsMonthCnt                   int64     `json:"kaspi_bank_scoring_active_contracts_month_cnt"`
	KaspiBankScoringClosedOverdueAmount                       float64   `json:"kaspi_bank_scoring_closed_overdue_amount"`
	KaspiBankScoringActiveContractsWithOverdue                int64     `json:"kaspi_bank_scoring_active_contracts_with_overdue"`
	KaspiBankScoringActiveContractsMaxSumMonthlyPaymentAmount float64   `json:"kaspi_bank_scoring_active_contracts_max_sum_monthly_payment_amount"`
	KaspiBankScoringMonthlyPaymentAllCntrsWithTrancheloans    float64   `json:"kaspi_bank_scoring_monthly_payment_all_cntrs_with_trancheloans"`
	KaspiBankScoringMonthlyPaymentNoCreditCards               float64   `json:"kaspi_bank_scoring_monthly_payment_no_credit_cards"`
	KaspiBankScoringCntMonthOverdue                           float64   `json:"kaspi_bank_scoring_cnt_month_overdue"`
	KaspiBankScoringActiveCreditLimitAllCards                 float64   `json:"kaspi_bank_scoring_active_credit_limit_all_cards"`
	KaspiBankScoringMaxCntDaysOverdueAllContrs                int64     `json:"kaspi_bank_scoring_max_cnt_days_overdue_all_contrs"`
	KaspiBankScoringMaxCntDaysClosedContrs                    int64     `json:"kaspi_bank_scoring_max_cnt_days_closed_contrs"`
	KaspiBankScoringActiveCntContrs                           int64     `json:"kaspi_bank_scoring_active_cnt_contrs"`
	KaspiBankScoringContrsClosedCntAllHistory                 int64     `json:"kaspi_bank_scoring_contrs_closed_cnt_all_history"`
	KaspiBankScoring1stCreditAgeMonthsAll                     int64     `json:"kaspi_bank_scoring_1st_credit_age_months_all"`
	KaspiBankScoringFinOrgsCntActiveContrs                    int64     `json:"kaspi_bank_scoring_fin_orgs_cnt_active_contrs"`
	KaspiBankScoringFinOrgsCntOverdueContrs                   int64     `json:"kaspi_bank_scoring_fin_orgs_cnt_overdue_contrs"`
	KaspiBankScoringClientCreditLifeLength                    int64     `json:"kaspi_bank_scoring_client_credit_life_length"`
	KaspiBankScoringOverdueSumWithoutKaspiData                float64   `json:"kaspi_bank_scoring_overdue_sum_without_kaspi_data"`
	KaspiBankScoringAmountMaxBefore                           float64   `json:"kaspi_bank_scoring_amount_max_before"`
	KaspiBankScoringCntDays1yMaxOverdue                       int64     `json:"kaspi_bank_scoring_cnt_days_1y_max_overdue"`
	KaspiBankScoringDelinqAmount                              float64   `json:"kaspi_bank_scoring_delinq_amount"`
	KaspiBankScoringCurrentMaxCntOverdueDays                  int64     `json:"kaspi_bank_scoring_current_max_cnt_overdue_days"`
}

type quantumScoring struct {
	QuantumScoringDate               time.Time `json:"quantum_scoring_date"`
	QuantumScoringCsIsCurOvdMore30d  float64   `json:"quantum_scoring_cs_is_cur_ovd_more_30d"`
	QuantumScoringIsOvdMore30dIn6m   float64   `json:"quantum_scoring_is_ovd_more_30d_in_6m"`
	QuantumScoringIsNegStatusActCred float64   `json:"quantum_scoring_is_neg_status_act_cred"`
	QuantumScoringIsActCredCntMore7  float64   `json:"quantum_scoring_is_act_cred_cnt_more_7"`
	QuantumScoringIsReqCntMore7In30d float64   `json:"quantum_scoring_is_req_cnt_more_7_in_30d"`
	QuantumScoringIsCurOvdMore1000   float64   `json:"quantum_scoring_is_cur_ovd_more_1000"`
}

type sberbankScoring struct {
	SberbankScoringDate                                time.Time `json:"sberbank_scoring_date"`
	SberbankScoringIsCurMonCredPaysumBetw200001300000  float64   `json:"sberbank_scoring_is_cur_mon_cred_paysum_betw_200001_300000"`
	SberbankScoringIsMaxOvdMore100000                  float64   `json:"sberbank_scoring_is_max_ovd_more_100000"`
	SberbankScoringCreditHistoryLess180d               float64   `json:"sberbank_scoring_credit_history_less_180d"`
	SberbankScoringIsOvdCntBetw23                      float64   `json:"sberbank_scoring_is_ovd_cnt_betw_2_3"`
	SberbankScoringIsMaxOvdEq0                         float64   `json:"sberbank_scoring_is_max_ovd_eq_0"`
	SberbankScoringIsMaxOvdMoreEq91d                   float64   `json:"sberbank_scoring_is_max_ovd_more_eq_91d"`
	SberbankScoringIsCurMonCredPaysumLessEq100000      float64   `json:"sberbank_scoring_is_cur_mon_cred_paysum_less_eq_100000"`
	SberbankScoringIsCurMonCredPaysumBetw100001200000  float64   `json:"sberbank_scoring_is_cur_mon_cred_paysum_betw_100001_200000"`
	SberbankScoringIsMaxOvdBetw6d30d                   float64   `json:"sberbank_scoring_is_max_ovd_betw_6d_30d"`
	SberbankScoringIsCurOvdMore0                       float64   `json:"sberbank_scoring_is_cur_ovd_more_0"`
	SberbankScoringIsMaxOvdBetw50001100000             float64   `json:"sberbank_scoring_is_max_ovd_betw_50001_100000"`
	SberbankScoringIsCurMonCredPaysumBetw5000011000000 float64   `json:"sberbank_scoring_is_cur_mon_cred_paysum_betw_500001_1000000"`
	SberbankScoringIsOvdCntEq1                         float64   `json:"sberbank_scoring_is_ovd_cnt_eq_1"`
	SberbankScoringIsCurCredDebtsumMoreEq1000000eur    float64   `json:"sberbank_scoring_is_cur_cred_debtsum_more_eq_1000000eur"`
	SberbankScoringIsMaxOvdEq5d                        float64   `json:"sberbank_scoring_is_max_ovd_eq_5d"`
	SberbankScoringIsMaxOvdBetw150000                  float64   `json:"sberbank_scoring_is_max_ovd_betw_1_50000"`
	SberbankScoringIsCurMonCredPaysumBetw300001400000  float64   `json:"sberbank_scoring_is_cur_mon_cred_paysum_betw_300001_400000"`
	SberbankScoringIsCurMonCredPaysumBetw400001500000  float64   `json:"sberbank_scoring_is_cur_mon_cred_paysum_betw_400001_500000"`
	SberbankScoringIsOvdCntBetw45                      float64   `json:"sberbank_scoring_is_ovd_cnt_betw_4_5"`
	SberbankScoringIsMaxOvdBetw31d90d                  float64   `json:"sberbank_scoring_is_max_ovd_betw_31d_90d"`
}

type vtbScoring struct {
	VtbScoringDate                                     time.Time `json:"vtb_scoring_date"`
	VtbScoringIsNegDebtDelForLossesIn18m               float64   `json:"vtb_scoringis_neg_debt_del_for_losses_in_18m"`
	VtbScoringIsNegDebtJudjRefIn19m36m                 float64   `json:"vtb_scoring_is_neg_debt_judj_ref_in_19m_36m"`
	VtbScoringIsNegDebtDelOutsysIn19m36m               float64   `json:"vtb_scoring_is_neg_debt_del_outsys_in_19m_36m"`
	VtbScoringMmIsReqCntLess10In30d                    float64   `json:"vtb_scoring_mm_is_req_cnt_less_10_in_30d"`
	VtbScoringIsNegLoanDebtDelForBalanceIn19m36m       float64   `json:"vtb_scoring_is_neg_loan_debt_del_for_balance_in_19m_36m"`
	VtbScoringIsNegOvdFrom91dTo360dIn19m36m            float64   `json:"vtb_scoring_is_neg_ovd_from_91d_to_360d_in_19m_36m"`
	VtbScoringIsNegDebtForgForHoplIn19m36m             float64   `json:"vtb_scoring_is_neg_debt_forg_for_hopl_in_19m_36m"`
	VtbScoringIsNegDebtForgForDeathIn18m               float64   `json:"vtb_scoring_is_neg_debt_forg_for_death_in_18m"`
	VtbScoringIsNegDebtDelForLosses17In19m36m          float64   `json:"vtb_scoring_is_neg_debt_del_for_losses_17_in_19m_36m"`
	VtbScoringMmIsProlongOf1MoreIn19m36m               float64   `json:"vtb_scoring_mm_is_prolong_of_1_more_in_19m_36m"`
	VtbScoringIsNegDebt18ForgForDeathIn19m36m          float64   `json:"vtb_scoring_is_neg_debt_18_forg_for_death_in_19m_36m"`
	VtbScoringMmIsOvdCntLess10More9dIn18m              float64   `json:"vtb_scoring_mm_is_ovd_cnt_less_10_more_9d_in_18m"`
	VtbScoringIsNegLoanDebtDelForBalanceIn18m          float64   `json:"vtb_scoring_is_neg_loan_debt_del_for_balance_in_18m"`
	VtbScoringIsNegDebtJudjRefIn18mStatus20            float64   `json:"vtb_scoring_is_neg_debt_judj_ref_in_18m_status_20"`
	VtbScoringMmIsOvdLess30dMoreEq2000tIn12m           float64   `json:"vtb_scoring_mm_is_ovd_less_30d_more_eq_2000t_in_12m"`
	VtbScoringMmIsReqCntLess6In7d                      float64   `json:"vtb_scoring_mm_is_req_cnt_less_6_in_7d"`
	VtbScoringIsNegAssignCredRightsIn18mStatus2        float64   `json:"vtb_scoring_is_neg_assign_cred_rights_in_18m_status_2"`
	VtbScoringIsNegDebtForgForHoplIn18mStatus25        float64   `json:"vtb_scoring_is_neg_debt_forg_for_hopl_in_18m_status_25"`
	VtbScoringIsNegTransfCredRightsIn19m36mStatus26    float64   `json:"vtb_scoring_is_neg_transf_cred_rights_in_19m_36m_status_26"`
	VtbScoringMmIsOvdCntMore9More9999In18m             float64   `json:"vtb_scoring_mm_is_ovd_cnt_more_9_more_9999_in_18m"`
	VtbScoringMmIsOvdCntMore10More9999In18m            float64   `json:"vtb_scoring_mm_is_ovd_cnt_more_10_more_9999_in_18m"`
	VtbScoringIsNegOvdFrom61dTo90dIn18mStatus13        float64   `json:"vtb_scoring_is_neg_ovd_from_61d_to_90d_in_18m_status_13"`
	VtbScoringIsNegOvdFrom7dTo30dIn19m36mStatus11      float64   `json:"vtb_scoring_is_neg_ovd_from_7d_to_30d_in_19m_36m_status_11"`
	VtbScoringIsNegOvdMore360dIn19m36mStatus16         float64   `json:"vtb_scoring_is_neg_ovd_more_360d_in_19m_36m_status_16"`
	VtbScoringMmIsReqCntMore5In7d                      float64   `json:"vtb_scoring_mm_is_req_cnt_more_5_in_7d"`
	VtbScoringIsNegDebtDelForOtherReasIn19m36mStatus19 float64   `json:"vtb_scoring_is_neg_debt_del_for_other_reas_in_19m_36m_status_19"`
	VtbScoringIsNegDebtDelForOtherReasIn18mStatus19    float64   `json:"vtb_scoring_is_neg_debt_del_for_other_reas_in_18m_status_19"`
	VtbScoringIsNegDebtJudjPaidIn18mStatus23           float64   `json:"vtb_scoring_is_neg_debt_judj_paid_in_18m_status_23"`
	VtbScoringIsNegTransfCredRightsIn18mStatus26       float64   `json:"vtb_scoring_is_neg_transf_cred_rights_in_18m_status_26"`
	VtbScoringMmIsOvdCntMore9More9dMore5000In18m       float64   `json:"vtb_scoring_mm_is_ovd_cnt_more_9_more_9d_more_5000_in_18m"`
	VtbScoringMmIsOvdCntMore9More9dIn19m36m            float64   `json:"vtb_scoring_mm_is_ovd_cnt_more_9_more_9d_in_19m_36m"`
	VtbScoringIsNegOvdMore360dIn18mStatus16            float64   `json:"vtb_scoring_is_neg_ovd_more_360d_in_18m_status_16"`
	VtbScoringIsNegDebtDelOutsysIn18mStatus24          float64   `json:"vtb_scoring_is_neg_debt_del_outsys_in_18m_status_24"`
	VtbScoringIsNegAssignCredRightsIn19m36mStatus2     float64   `json:"vtb_scoring_is_neg_assign_cred_rights_in_19m_36m_status_2"`
	VtbScoringMmIsProlongIn18m                         float64   `json:"vtb_scoring_mm_is_prolong_in_18m"`
	VtbScoringIsNegOvdFrom31dTo60dIn19m36mStatus12     float64   `json:"vtb_scoring_is_neg_ovd_from_31d_to_60d_in_19m_36m_status_12"`
	VtbScoringMmIsCurOvdLess8dOrLess15000              float64   `json:"vtb_scoring_mm_is_cur_ovd_less_8d_or_less_15000"`
	VtbScoringMmIsOvdMoreEq30dMoreEq2000tIn12m         float64   `json:"vtb_scoring_mm_is_ovd_more_eq_30d_more_eq_2000t_in_12m"`
	VtbScoringMmIsCurOvdMoreEq8dOrMoreEq15000          float64   `json:"vtb_scoring_mm_is_cur_ovd_more_eq_8d_or_more_eq_15000"`
	VtbScoringIsNegOvdFrom31dTo60dIn18Status12         float64   `json:"vtb_scoring_is_neg_ovd_from_31d_to_60d_in_18_status_12"`
	VtbScoringIsNegOvdFrom7dTo30dIn18mStatus11         float64   `json:"vtb_scoring_is_neg_ovd_from_7d_to_30d_in_18m_status_11"`
	VtbScoringMmIsReqCntMore9In30d                     float64   `json:"vtb_scoring_mm_is_req_cnt_more_9_in_30d"`
	VtbScoringIsNegDebtJudjPaidIn19m36mStatus23        float64   `json:"vtb_scoring_is_neg_debt_judj_paid_in_19m_36m_status_23"`
	VtbScoringIsNegOvdFrom61dTo90dIn19m36mStatus13     float64   `json:"vtb_scoring_is_neg_ovd_from_61d_to_90d_in_19m_36m_status_13"`
	VtbScoringIsNegOvdFrom91dTo360dIn18mStatus15       float64   `json:"vtb_scoring_is_neg_ovd_from_91d_to_360d_in_18m_status_15"`
}

type creditRatingScoring struct {
	CreditRatingScoringRiskGroup int64     `json:"credit_rating_scoring_risk_group"`
	CreditRatingScoringDate      time.Time `json:"credit_rating_scoring_date"`
	CreditRatingScoringScore     int64     `json:"credit_rating_scoring_score"`
}

type creditPropensity struct {
}

type scoring struct {
	BmlPdlScoring             bmlPdlScoring             `json:"bml_pdl_scoring"`
	AsiaCreditBankScoring     asiaCreditBankScoring     `json:"asia_credit_bank_scoring"`
	AlfaBankScoring           alfaBankScoring           `json:"alfa_bank_scoring"`
	BehaviorScoring           behaviorScoring           `json:"behavior_scoring"`
	BckBankScoring            bckBankScoring            `json:"bck_bank_scoring"`
	ApplicationScoring        applicationScoring        `json:"application_scoring"`
	FicoScoring               ficoScoring               `json:"fico_scoring"`
	FraudScoring              fraudScoring              `json:"fraud_scoring"`
	FastCashScoring           fastCashScoring           `json:"fast_cash_scoring"`
	FreedomFinanceBankScoring freedomFinanceBankScoring `json:"freedom_finance_bank_scoring"`
	ForteBankScoring          forteBankScoring          `json:"forte_bank_scoring"`
	JysanBankScoring          jysanBankScoring          `json:"jysan_bank_scoring"`
	KaspiBankScoring          kaspiBankScoring          `json:"kaspi_bank_scoring"`
	QuantumScoring            quantumScoring            `json:"quantum_scoring"`
	SberbankScoring           sberbankScoring           `json:"sberbank_scoring"`
	VtbScoring                vtbScoring                `json:"vtb_scoring"`
	CreditRatingScoring       creditRatingScoring       `json:"credit_rating_scoring"`
}

type financingTerrExtrList struct {
	FinancingTerrExtrListStatus      string    `json:"financing_terr_extr_list_status"`
	FinancingTerrExtrListIncludeDate time.Time `json:"financing_terr_extr_list_include_date"`
	FinancingTerrExtrListExcludeDate time.Time `json:"financing_terr_extr_list_exclude_date"`
}

type opv struct {
	OpvAmount        float64   `json:"opv_amount"`
	OpvPeriod        time.Time `json:"opv_period"`
	OpvCompanyIinBin string    `json:"opv_company_iin_bin"`
}

type gender struct {
	GenderCode   string `json:"gender_code"`
	GenderNameRu string `json:"gender_name_ru"`
	GenderNameKz string `json:"gender_name_kz"`
}

type nationality struct {
	NationalityCode   string `json:"nationality_code"`
	NationalityNameRu string `json:"nationality_name_ru"`
	NationalityNameKz string `json:"nationality_name_kz"`
	NationalityNameEn string `json:"nationality_name_en"`
}

type income struct {
	IncomeAmount        float64   `json:"income_amount"`
	IncomePaymentDate   time.Time `json:"income_payment_date"`
	IncomePeriod        time.Time `json:"income_period"`
	IncomeCompanyIinBin string    `json:"income_company_iin_bin"`
	IncomeRefreshDate   time.Time `json:"income_refresh_date"`
}

type jobContract struct {
	JobContractDate          time.Time `json:"job_contract_date"`
	JobContractDateFrom      time.Time `json:"job_contract_date_from"`
	JobContractDateTo        time.Time `json:"job_contract_date_to"`
	JobContractCompanyIinBin string    `json:"job_contract_company_iin_bin"`
	//JobContractCompanyIinBinSha	string `json:"job_contract_company_iin_bin_sha"`
	JobContractPositionClassifierNameRu string    `json:"job_contract_position_classifier_name_ru"`
	JobContractClassifierCode           string    `json:"job_contract_classifier_code"`
	JobContractTerminationReasonId      int64     `json:"job_contract_termination_reason_id"`
	JobContractPositionNameRu           string    `json:"job_contract_position_name_ru"`
	JobContractRefreshDate              time.Time `json:"job_contract_refresh_date"`
	JobContractPositionClassifierNameKz string    `json:"job_contract_position_classifier_name_kz"`
	JobContractPositionClassifierCode   int64     `json:"job_contract_position_classifier_code"`
}

type jobHead struct {
	JobHeadCompanyIinBin string    `json:"job_head_company_iin_bin"`
	JobHeadRefreshDate   time.Time `json:"job_head_refresh_date"`
}

type job struct {
	JobContract jobContract `json:"job_contract"`
	JobHead     jobHead     `json:"job_head"`
}

type incomeRefund struct {
	IncomeRefundAmount       float64   `json:"income_refund_amount"`
	IncomeRefundDate         time.Time `json:"income_refund_date"`
	IncomeRefundRefundPeriod time.Time `json:"income_refund_refund_period"`
}

type criminalRecord struct { //Нет привязки к БТ
	CriminalRecordRefreshDate               time.Time `json:"criminal_record_refresh_date"`
	CriminalRecordArticleOfAttraction       string    `json:"criminal_record_article_of_attraction"`
	CriminalRecordCondemnationOrgRu         string    `json:"criminal_record_condemnation_org_ru"`
	CriminalRecordCondemnationDate          time.Time `json:"criminal_record_condemnation_date"`
	CriminalRecordMainPunishmentRu          string    `json:"criminal_record_main_punishment_ru"`
	CriminalRecordPunishmentExecutionTypeRu string    `json:"criminal_record_punishment_execution_type_ru"`
	CriminalRecordSentenceDateFrom          time.Time `json:"criminal_record_sentence_date_from"`
	CriminalRecordPunishmentTermDateFrom    time.Time `json:"criminal_record_punishment_term_date_from"`
	CriminalRecordReleaseDate               time.Time `json:"criminal_record_release_date"`
	CriminalRecordReleaseReasonRu           string    `json:"criminal_record_release_reason_ru"`
	CriminalRecordReleaseOrgRu              string    `json:"criminal_record_release_org_ru"`
	CriminalRecordGrantOfParole             bool      `json:"criminal_record_grant_of_parole"`
	CriminalRecordGrantOfParoleFactDateFrom time.Time `json:"criminal_record_grant_of_parole_fact_date_from"`
	CriminalRecordHoldToLiabilityOrgRu      string    `json:"criminal_record_hold_to_liability_org_ru"`
	CriminalRecordSuppressionType           string    `json:"criminal_record_suppression_type"`
	CriminalRecordDateOfArrest              time.Time `json:"criminal_record_date_of_arrest"`
	CriminalRecordCaseNumber                string    `json:"criminal_record_case_number"`
	CriminalRecordCaseDateFrom              time.Time `json:"criminal_record_case_date_from"`
	CriminalRecordArticleOfCondemnation     string    ` json:"criminal_record_article_of_condemnation"`
}

type kdn struct {
	KdnScore        float64   `json:"kdn_score"`
	KdnIdType       int64     `json:"kdn_id_type"`
	KdnRefreshDate  time.Time `json:"kdn_refresh_date"`
	KdnIncomeAmount float64   `json:"kdn_income_amount"`
	KdnDebtAmount   float64   `json:"kdn_debt_amount"`
}

type taxNotification struct {
	TaxNotificationOrgCode    string    `json:"tax_notification_org_code"`
	TaxNotificationNumber     string    `json:"tax_notification_number"`
	TaxNotificationDate       time.Time `json:"tax_notification_date"`
	TaxNotificationReturnDate time.Time `json:"tax_notification_return_date"`
	TaxNotificationNameRu     string    `json:"tax_notification_name_ru"`
}

type narcoRegistry struct {
	NarcoRegistryStatusNameRu string    `json:"narco_registry_status_name_ru"`
	NarcoRegistryStatusNameKZ string    `json:"narco_registry_status_name_kz"`
	NarcoRegistryRefreshDate  time.Time `json:"narco_registry_refresh_date"`
}

type photo struct {
	PhotoRefreshDate time.Time `json:"photo_refresh_date"`
	PhotoLocation    string    `json:"photo_location"`
}

type bmg struct {
	BmgMobileNumber string    `json:"bmg_mobile_number"`
	BmgRefreshDate  time.Time `json:"bmg_refresh_date"`
}

type psychoRegistry struct {
	PsychoRegistryStatusNameRu string    `json:"psycho_registry_status_name_ru"`
	PsychoRegistryStatusNameKz string    `json:"psycho_registry_status_name_kz"`
	PsychoRegistryRefreshDate  time.Time `json:"psycho_registry_refresh_date"`
}

type registrationAddress struct {
	RegistrationAddressCountryCode    string    `json:"registration_address_country_code"`
	RegistrationAddressCountryNameRu  string    `json:"registration_address_country_name_ru"`
	RegistrationAddressCountryNameKz  string    `json:"registration_address_country_name_kz"`
	RegistrationAddressRegionCode     string    `json:"registration_address_region_code"`
	RegistrationAddressRegionNameRu   string    `json:"registration_address_region_name_ru"`
	RegistrationAddressRegionNameKz   string    `json:"registration_address_region_name_kz"`
	RegistrationAddressDistrictCode   string    `json:"registration_address_district_code"`
	RegistrationAddressDistrictNameRu string    `json:"registration_address_district_name_ru"`
	RegistrationAddressDistrictNameKz string    `json:"registration_address_district_name_kz"`
	RegistrationAddressCity           string    `json:"registration_address_city"`
	RegistrationAddressStreet         string    `json:"registration_address_street"`
	RegistrationAddressBuilding       string    `json:"registration_address_building"`
	RegistrationAddressFlat           string    `json:"registration_address_flat"`
	RegistrationAddressCorpus         string    `json:"registration_address_corpus"`
	RegistrationAddressRefreshDate    time.Time `json:"registration_address_refresh_date"`
	RegistrationAddressName           string    `json:"registration_address_name"`
}

type tubRegistry struct {
	TubRegistryStatusNameRu string    `json:"tub_registry_status_name_ru"`
	TubRegistryStatusNameKz string    `json:"tub_registry_status_name_kz"`
	TubRegistryRefreshDate  time.Time `json:"tub_registry_refresh_date"`
}

type vaccinationKz struct {
	VaccinationKzPatientIIN            string    `json:"vaccination_kz_patient_iin"`
	VaccinationKzPatientDateOfBirth    time.Time `json:"vaccination_kz_patient_date_of_birth"`
	VaccinationKzPatientPassportDocnum string    `json:"vaccination_kz_patient_passport_docnum"`
	VaccinationKzDoctorIIN             string    `json:"vaccination_kz_doctor_iin"`
	VaccinationKzPatientFullNameRu     string    `json:"vaccination_kz_patient_full_name_ru"`
	VaccinationKzPatientFullNameKz     string    `json:"vaccination_kz_patient_full_name_kz"`
	VaccinationKzPatientFullNameEn     string    `json:"vaccination_kz_patient_full_name_en"`
	VaccinationKzPatientGender         string    `json:"vaccination_kz_patient_gender"`
	VaccinationKzDoctorFullNameRu      string    `json:"vaccination_kz_doctor_full_name_ru"`
	VaccinationKzDoctorFullNameKz      string    `json:"vaccination_kz_doctor_full_name_kz"`
	VaccinationKzDoctorFullNameEn      string    `json:"vaccination_kz_doctor_full_name_en"`
	VaccinationKzTypeofVaccinationRu   string    `json:"vaccination_kz_typeof_vaccination_ru"`
	VaccinationKzTypeofVaccinationKz   string    `json:"vaccination_kz_typeof_vaccination_kz"`
	VaccinationKzTypeofVaccinationEn   string    `json:"vaccination_kz_typeof_vaccination_en"`
	VaccinationKzDrugCountryRu         string    `json:"vaccination_kz_drug_country_ru"`
	VaccinationKzDrugCountryKz         string    `json:"vaccination_kz_drug_country_kz"`
	VaccinationKzDrugCountryEn         string    `json:"vaccination_kz_drug_country_en"`
	VaccinationKzDrugCompanyNameRu     string    `json:"vaccination_kz_drug_company_name_ru"`
	VaccinationKzDrugCompanyNameKz     string    `json:"vaccination_kz_drug_company_name_kz"`
	VaccinationKzDrugCompanyNameEn     string    `json:"vaccination_kz_drug_company_name_en"`
	VaccinationKzAddressRu             string    `json:"vaccination_kz_address_ru"`
	VaccinationKzAddressKz             string    `json:"vaccination_kz_address_kz"`
	VaccinationKzAddressEn             string    `json:"vaccination_kz_address_en"`
	VaccinationKzDrugNameRu            string    `json:"vaccination_kz_drug_name_ru"`
	VaccinationKzDrugNameKz            string    `json:"vaccination_kz_drug_name_kz"`
	VaccinationKzDrugNameEn            string    `json:"vaccination_kz_drug_name_en"`
	VaccinationKzVaccinationDateRu     string    `json:"vaccination_kz_vaccination_date_ru"`
	VaccinationKzVaccinationDateKz     string    `json:"vaccination_kz_vaccination_date_kz"`
	VaccinationKzVaccinationDateEn     string    `json:"vaccination_kz_vaccination_date_en"`
	VaccinationKzDoseRu                string    `json:"vaccination_kz_dose_ru"`
	VaccinationKzDoseKz                string    `json:"vaccination_kz_dose_kz"`
	VaccinationKzDoseEn                string    `json:"vaccination_kz_dose_en"`
	VaccinationKzSeries                string    `json:"vaccination_kz_series"`
	VaccinationKzRevacDoctorIIN        string    `json:"vaccination_kz_revac_doctor_iin"`
	VaccinationKzRevacAddressRu        string    `json:"vaccination_kz_revac_address_ru"`
	VaccinationKzRevacAddressKz        string    `json:"vaccination_kz_revac_address_kz"`
	VaccinationKzRevacAddressEn        string    `json:"vaccination_kz_revac_address_en"`
	VaccinationKzRevacDrugName         string    `json:"vaccination_kz_revac_drug_name"`
	VaccinationKzRevacDrugNameKz       string    `json:"vaccination_kz_revac_drug_name_kz"`
	VaccinationKzRevacDrugNameEn       string    `json:"vaccination_kz_revac_drug_name_en"`
	VaccinationKzRevacDrugCountryRu    string    `json:"vaccination_kz_revac_drug_country_ru"`
	VaccinationKzRevacDrugCountryKz    string    `json:"vaccination_kz_revac_drug_country_kz"`
	VaccinationKzRevacDrugCountryEn    string    `json:"vaccination_kz_revac_drug_country_en"`
	VaccinationKzRevacDrugCompanyRu    string    `json:"vaccination_kz_revac_drug_company_ru"`
	VaccinationKzRevacDrugCompanyKz    string    `json:"vaccination_kz_revac_drug_company_kz"`
	VaccinationKzRevacDrugCompanyEn    string    `json:"vaccination_kz_revac_drug_company_en"`
	VaccinationKzValidDate             time.Time `json:"vaccination_kz_valid_date"`
	VaccinationKzRevacSeries           string    `json:"vaccination_kz_revac_series"`
	VaccinationKzRevacDate             time.Time `json:"vaccination_kz_revac_date"`
	VaccinationKzRevacDose             string    `json:"vaccination_kz_revac_dose"`
	VaccinationKzStage                 string    `json:"vaccination_kz_stage"`
}

type realEstateObjectRegistration struct {
	RealEstateObjectRegistrationPersonIdentifyingDocnum string    `json:"real_estate_object_registration_person_identifying_docnum"`
	RealEstateObjectRegistrationLegalDocumentInfo       string    `json:"real_estate_object_registration_legal_document_info"`
	RealEstateObjectRegistrationObjectId                int64     `json:"real_estate_object_registration_object_id"`
	RealEstateObjectRegistrationDateFrom                time.Time `json:"real_estate_object_registration_date_from"`
	RealEstateObjectRegistrationDateTo                  time.Time `json:"real_estate_object_registration_date_to"`
	RealEstateObjectRegistrationRefreshDate             time.Time `json:"real_estate_object_registration_refresh_date"`
	RealEstateObjectRegistrationRightTypeNameRu         string    `json:"real_estate_object_registration_right_type_name_ru"`
}

type wantedChildEnforcement struct {
	WantedChildEnforcementCategory            string    `json:"wanted_child_enforcement_category"`
	WantedChildEnforcementDossierNumber       string    `json:"wanted_child_enforcement_dossier_number"`
	WantedChildEnforcementDossierCreationDate time.Time `json:"wanted_child_enforcement_dossier_creation_date"`
	WantedChildEnforcementInitiator           string    `json:"wanted_child_enforcement_initiator"`
	WantedChildEnforcementSearchDep           string    `json:"wanted_child_enforcement_search_dep"`
	WantedChildEnforcementDaysSearched        int64     `json:"wanted_child_enforcement_days_searched"`
}

type wantedDebtor struct {
	WantedDebtorCategory            string    `json:"wanted_debtor_category"`
	WantedDebtorDossierNumber       string    `json:"wanted_debtor_dossier_number"`
	WantedDebtorDossierCreationDate time.Time `json:"wanted_debtor_dossier_creation_date"`
	WantedDebtorInitiator           string    `json:"wanted_debtor_initiator"`
	WantedDebtorSearchDep           string    `json:"wanted_debtor_search_dep"`
	WantedDebtorDaysSearched        int64     `json:"wanted_debtor_days_searched"`
}

type wantedKgd struct {
	WantedKgdCriminalArticle   string `json:"wanted_kgd_criminal_article"`
	WantedKgdType              string `json:"wanted_kgd_type"`
	WantedKgdCriminalMeasure   string `json:"wanted_kgd_criminal_measure"`
	WantedKgdCriminalInfo      string `json:"wanted_kgd_criminal_info"`
	WantedKgdCriminalInitiator string `json:"wanted_kgd_criminal_initiator"`
	WantedKgdCriminalAddInfo   string `json:"wanted_kgd_criminal_add_info"`
}

type wantedDefendant struct {
	WantedDefendantCategory            string    `json:"wanted_defendant_category"`
	WantedDefendantDossierNumber       string    `json:"wanted_defendant_dossier_number"`
	WantedDefendantDossierCreationDate time.Time `json:"wanted_defendant_dossier_creation_date"`
	WantedDefendantInitiator           string    `json:"wanted_defendant_initiator"`
	WantedDefendantSearchDep           string    `json:"wanted_defendant_search_dep"`
	WantedDefendantDaysSearched        int64     `json:"wanted_defendant_days_searched"`
}

type wantedForViolation struct {
	WantedForViolationCategory            string    `json:"wanted_for_violation_category"`
	WantedForViolationDossierNumber       string    `json:"wanted_for_violation_dossier_number"`
	WantedForViolationDossierCreationDate time.Time `json:"wanted_for_violation_dossier_creation_date"`
	WantedForViolationInitiator           string    `json:"wanted_for_violation_initiator"`
	WantedForViolationSearchDep           string    `json:"wanted_for_violation_search_dep"`
	WantedForViolationDaysSearched        int64     `json:"wanted_for_violation_days_searched"`
}

type wantedGovDebtor struct {
	WantedGovDebtorCategory            string    `json:"wanted_gov_debtor_category"`
	WantedGovDebtorDossierNumber       string    `json:"wanted_gov_debtor_dossier_number"`
	WantedGovDebtorDossierCreationDate time.Time `json:"wanted_gov_debtor_dossier_creation_date"`
	WantedGovDebtorInitiator           string    `json:"wanted_gov_debtor_initiator"`
	WantedGovDebtorSearchDep           string    `json:"wanted_gov_debtor_search_dep"`
	WantedGovDebtorDaysSearched        int64     `json:"wanted_gov_debtor_days_searched"`
}

type wantedLost struct {
	WantedLostDossierNumber       string    `json:"wanted_lost_dossier_number"`
	WantedLostDossierCreationDate time.Time `json:"wanted_lost_dossier_creation_date"`
	WantedLostInitiator           string    `json:"wanted_lost_initiator"`
	WantedLostSearchDep           string    `json:"wanted_lost_search_dep"`
	WantedLostDaysSearched        int64     `json:"wanted_lost_days_searched"`
	WantedLostDisappearReason     string    `json:"wanted_lost_disappear_reason"`
}

type wanted struct {
	WantedChildEnforcement wantedChildEnforcement `json:"wanted_child_enforcement"`
	WantedDebtor           wantedDebtor           `json:"wanted_debtor"`
	WantedKgd              wantedKgd              `json:"wanted_kgd"`
	WantedDefendant        wantedDefendant        `json:"wanted_defendant"`
	WantedForViolation     wantedForViolation     `json:"wanted_for_violation"`
	WantedGovDebtor        wantedGovDebtor        `json:"wanted_gov_debtor"`
	WantedLost             wantedLost             `json:"wanted_lost"`
}

type realEstateQueue struct {
	RealEstateQueueDate            time.Time `json:"real_estate_queue_date"`
	RealEstateQueueFamilyMemberCnt int64     `json:"real_estate_queue_family_member_cnt"`
}

type realEstateObjectEncumbrance struct {
	RealEstateObjectEncumbranceObjectId     int64     `json:"real_estate_object_encumbrance_object_id"`
	RealEstateObjectEncumbranceDateFrom     time.Time `json:"real_estate_object_encumbrance_date_from"`
	RealEstateObjectEncumbranceDateTo       time.Time `json:"real_estate_object_encumbrance_date_to"`
	RealEstateObjectEncumbranceRefreshDate  time.Time `json:"real_estate_object_encumbrance_refresh_date"`
	RealEstateObjectEncumbranceHolderIinBin string    `json:"real_estate_object_encumbrance_holder_iin_bin"`
	RealEstateObjectEncumbranceHolderInfo   string    `json:"real_estate_object_encumbrance_holder_info"`
}

type bankruptcyApplicationDebt struct {
	BankruptcyApplicationDebtCreditorCompanyIin    string `json:"bankruptcy_application_debt_creditor_company_iin"`
	BankruptcyApplicationDebtCreditorCompanyNameRu string `json:"bankruptcy_application_debt_creditor_company_name_ru"`
	BankruptcyApplicationDebtAmount                string `json:"bankruptcy_application_debt_amount"`
	BankruptcyApplicationDebtDate                  string `json:"bankruptcy_application_debt_date"`
}

type bankruptcyApplication struct {
	BankruptcyApplicationTypeCode                 string                    `json:"bankruptcy_application_type_code"`
	BankruptcyApplicationTypeNameRu               string                    `json:"bankruptcy_application_type_name_ru"`
	BankruptcyApplicationIin                      string                    `json:"bankruptcy_application_iin"`
	BankruptcyApplicationTypeNameKz               string                    `json:"bankruptcy_application_type_name_kz"`
	BankruptcyApplicationTypeNameEn               string                    `json:"bankruptcy_application_type_name_en"`
	BankruptcyApplicationStatementNumber          string                    `json:"bankruptcy_application_statement_number"`
	BankruptcyApplicationFio                      string                    `json:"bankruptcy_application_fio"`
	BankruptcyApplicationDate                     time.Time                 `json:"bankruptcy_application_date"`
	BankruptcyApplicationRefusalDate              time.Time                 `json:"bankruptcy_application_refusal_date"`
	BankruptcyApplicationRefusalReasonCode        string                    `json:"bankruptcy_application_refusal_reason_code"`
	BankruptcyApplicationRefusalReasonRu          string                    `json:"bankruptcy_application_refusal_reason_ru"`
	BankruptcyApplicationRefusalReasonKz          string                    `json:"bankruptcy_application_refusal_reason_kz"`
	BankruptcyApplicationRefusalReasonRn          string                    `json:"bankruptcy_application_refusal_reason_rn"`
	BankruptcyApplicationProceedingsStartDate     time.Time                 `json:"bankruptcy_application_proceedings_start_date"`
	BankruptcyApplicationStatementAcceptDate      time.Time                 `json:"bankruptcy_application_statement_accept_date"`
	BankruptcyApplicationCourtDecisionDate        time.Time                 `json:"bankruptcy_application_court_decision_date"`
	BankruptcyApplicationTerminationDate          time.Time                 `json:"bankruptcy_application_termination_date"`
	BankruptcyApplicationTerminationReasonCode    string                    `json:"bankruptcy_application_termination_reason_code"`
	BankruptcyApplicationTerminationReasonRu      string                    `json:"bankruptcy_application_termination_reason_ru"`
	BankruptcyApplicationTerminationReasonKz      string                    `json:"bankruptcy_application_termination_reason_kz"`
	BankruptcyApplicationTerminationReasonEn      string                    `json:"bankruptcy_application_termination_reason_en"`
	BankruptcyApplicationTerminationInitiatorCode string                    `json:"bankruptcy_application_termination_initiator_code"`
	BankruptcyApplicationTerminationInitiatorRu   string                    `json:"bankruptcy_application_termination_initiator_ru"`
	BankruptcyApplicationTerminationInitiatorKz   string                    `json:"bankruptcy_application_termination_initiator_kz"`
	BankruptcyApplicationTerminationInitiatorEn   string                    `json:"bankruptcy_application_termination_initiator_en"`
	BankruptcyApplicationBankruptStartDate        time.Time                 `json:"bankruptcy_application_bankrupt_start_date"`
	BankruptcyApplicationBankruptEndDate          time.Time                 `json:"bankruptcy_application_bankrupt_end_date"`
	BankruptcyApplicationBankruptEndReasonCode    string                    `json:"bankruptcy_application_bankrupt_end_reason_code"`
	BankruptcyApplicationBankruptEndReasonRu      string                    `json:"bankruptcy_application_bankrupt_end_reason_ru"`
	BankruptcyApplicationBankruptEndReasonKz      string                    `json:"bankruptcy_application_bankrupt_end_reason_kz"`
	BankruptcyApplicationBankruptEndReasonEn      string                    `json:"bankruptcy_application_bankrupt_end_reason_en"`
	BankruptcyApplicationRefreshDate              time.Time                 `json:"bankruptcy_application_refresh_date"`
	BankruptcyApplicationStatementStatusCode      string                    `json:"bankruptcy_application_statement_status_code"`
	BankruptcyApplicationStatementStatusRu        string                    `json:"bankruptcy_application_statement_status_ru"`
	BankruptcyApplicationStatementStatusKz        string                    `json:"bankruptcy_application_statement_status_kz"`
	BankruptcyApplicationStatementStatusEn        string                    `json:"bankruptcy_application_statement_status_en"`
	BankruptcyApplicationDebt                     bankruptcyApplicationDebt `json:"bankruptcy_application_debt"`
}

type debtorMu struct {
	DebtorMuEnforcementProceedingsDate string `json:"debtor_mu_enforcement_proceedings_date"`
	DebtorMuEnforcementProceedingsNum  string `json:"debtor_mu_enforcement_proceedings_num"`
	DebtorMuDebtAmount                 string `json:"debtor_mu_debt_amount"`
	DebtorMuTypeNameRu                 string `json:"debtor_mu_type_name_ru"`
	DebtorMuTypeOfLegalUnitNameRu      string `json:"debtor_mu_type_of_legal_unit_name_ru"`
	DebtorMuCollectorName              string `json:"debtor_mu_collector_name"`
	DebtorMuCollectorIinBin            string `json:"debtor_mu_collector_iin_bin"`
}
