package repository

import (
	"errors"

	"github.com/donnyirianto/be-stock-app/modules/base/domain"
	"gorm.io/gorm"
)

// Act Base
func (r *BaseRepositoryImpl) GetMenu(role string) ([]*domain.BaseItem, error) {

	var resBase []*domain.BaseItem

	query := `SELECT a.id, a.nama, a.link, a.id_main, a.urut
	FROM m_menu a
	RIGHT JOIN m_menu_role b ON a.id = b.id_menu
	WHERE 
	b.id_role = ?
	AND a.nama !='';`

	err := r.mysqlConn.Raw(query, role).Scan(&resBase).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return resBase, nil
}
