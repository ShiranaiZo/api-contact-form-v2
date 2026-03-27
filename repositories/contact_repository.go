package repositories

import (
	"time"

	"github.com/ShiranaiZo/api-contact-form-v2/models"
	"gorm.io/gorm"
)

type ContactRepository interface {
	Create(contact *models.Contact) error
	FindAll() ([]models.Contact, error)
	FindByID(id uint) (*models.Contact, error)
	Update(contact *models.Contact) error
	Delete(contact *models.Contact) error
}

type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{db}
}

func (r *contactRepository) Create(contact *models.Contact) error {
	return r.db.Create(contact).Error
}

func (r *contactRepository) FindAll() ([]models.Contact, error) {
	var contacts []models.Contact
	err := r.db.Where("deleted_at = ?", "0000-00-00 00:00:00").Find(&contacts).Error
	return contacts, err
}

func (r *contactRepository) FindByID(id uint) (*models.Contact, error) {
	var contact models.Contact

	err := r.db.Where("id = ? AND deleted_at = ?", id, "0000-00-00 00:00:00").First(&contact).Error

	return &contact, err
}

func (r *contactRepository) Update(contact *models.Contact) error {
	return r.db.Save(contact).Error
}

func (r *contactRepository) Delete(contact *models.Contact) error {
	contact.DeletedAt = time.Now()

	return r.db.Save(contact).Error
}
