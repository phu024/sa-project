package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phu024/sa-64/entity"
)

// POST /patien
func CreatePatient(c *gin.Context) {

	var patient entity.Patient
	var gender entity.Gender
	var allergy entity.Allergy
	var underlying_disease entity.Underlying_disease
	var recorder entity.Recorder

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3: ค้นหา recorderID ด้วย id
	if tx := entity.DB().Where("id = ?", patient.RecorderID).First(&recorder); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	// 9: ค้นหา genderID ด้วย id
	if tx := entity.DB().Where("id = ?", patient.GenderID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	// 10: ค้นหา Allergy ด้วย id
	if tx := entity.DB().Where("id = ?", patient.AllergyID).First(&allergy); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Allergy not found"})
		return
	}

	// 11: ค้นหา underlying_disease ด้วย id
	if tx := entity.DB().Where("id = ?", patient.Underlying_diseaseID).First(&allergy); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "underlying_disease not found"})
		return
	}
	// 12: สร้าง patient
	wv := entity.Patient{
		Recorder:           recorder,           // โยงความสัมพันธ์กับ Entity Recorder
		Gender:             gender,             // โยงความสัมพันธ์กับ Entity gender
		Allergy:            allergy,            // โยงความสัมพันธ์กับ Entity allergy
		Underlying_disease: underlying_disease, // โยงความสัมพันธ์กับ Entity underlying_disease
		Id_card:            patient.Id_card,    // ตั้งค่าฟิลด์ Id_card
		FirstName:          patient.FirstName,  // ตั้งค่าฟิลด์ Firstname
		LastName:           patient.LastName,   // ตั้งค่าฟิลด์ Lastname
		Birthdate:          patient.Birthdate,  // ตั้งค่าฟิลด์ Birthdate
		Age:                patient.Age,        // ตั้งค่าฟิลด์ Age
	}

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /patient/:id
func GetPatient(c *gin.Context) {
	var patient entity.Patient
	id := c.Param("id")
	if err := entity.DB().Preload("gender").Preload("allergy").Preload("underlying_disease").Preload("recorder").Raw("SELECT * FROM patients WHERE id = ?", id).Find(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patient})
}

// GET /watch_underlying_diseases
func ListPatients(c *gin.Context) {
	var patients []entity.Patient
	if err := entity.DB().Preload("gender").Preload("allergy").Preload("underlying_disease").Preload("recorder").Raw("SELECT * FROM patients").Find(&patients).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patients})
}

// DELETE /watch_underlying_diseases/:id
func DeletePatient(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM patients WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /watch_underlying_diseases
func UpdatePatient(c *gin.Context) {
	var patient entity.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", patient.ID).First(&patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	if err := entity.DB().Save(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patient})
}
