import (
	"database/sql"

	"github.com/techinscribed/repository-db/models"
)

// UserRepo implements models.UserRepository
type PasteRepo struct {
	db *sql.DB
}

// NewUserRepo ..
func NewPasteRepo(db *sql.DB) *PasteRepo {
	return &PasteRepo{
		db: db,
	}
}

// FindByID ..
func (r *PasteRepo) FindByID(ID int) (*models.User, error) {
	return &models.User{}, nil
}

// Save ..
func (r *PastPasteRepoeRpo) Save(user *models.User) error {
	return nil
}