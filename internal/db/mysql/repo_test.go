package mysql

import (
	"testing"

	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/repository"
)

func TestInterface(t *testing.T) {
	var _ repository.BoardRepository = new(BoardMySQLRepository)
	var _ repository.BoardUserRepository = new(BoardUserMySQLRepository)
	var _ repository.StatusRepository = new(StatusMySQLRepository)
	var _ repository.TaskRepository = new(TaskMySQLRepository)
	var _ repository.UserRepository = new(UserMySQLRepository)
}
