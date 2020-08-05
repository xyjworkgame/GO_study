package repository

import (
	"companyIntroHandler/model"
	"companyIntroHandler/utils"
	"github.com/jinzhu/gorm"
	example "companyIntroHandler/companyIntro"
)

type CompanyIntroRepo struct {
	DB *gorm.DB
}

const (
	companyIntroParams = "id,content,status,seq_no,title"
)

func NewCompanyIntro(db *gorm.DB) *CompanyIntroRepo {
	return &CompanyIntroRepo{DB: db}

}

func (c *CompanyIntroRepo) Insert(companyIntro *example.CompanyIntro) error {
	tx, err := c.routine()
	if err != nil {
		return err
	}
	if err := tx.Create(&companyIntro).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}

// 查找所有内容
func (c *CompanyIntroRepo) Select(pagination *model.Pagination, status int) ([]model.CompanyIntroJson,error) {
	db := c.DB.Model(model.CompanyIntro{})
	var companyIntros []model.CompanyIntroJson
	pagination.Check()
	offset := (pagination.PageNum - 1) * pagination.PageSize
	limit := pagination.PageSize
	var totalSize int32
	totalSize  = 0
	if status == -1 {
		if err := db.Count(&totalSize).Error; err != nil {
			return nil,  err
		}
		if err := db.Select(companyIntroParams).Offset(offset).
			Order("seq_no",false).
			Limit(limit).Scan(&companyIntros).Error; err != nil {
			return nil,  err
		}
		pagination.Total = totalSize
		utils.CalPageTotal(totalSize,pagination)
		return companyIntros,nil
	} else {
		if err := db.Where("status = ?", status).
			Count(&totalSize).Error; err != nil {
			return nil,  err
		}
		err := db.Where("status = ?", status).
			Order("seq_no",false).
			Select(companyIntroParams).
			Offset(offset).Limit(limit).Scan(&companyIntros).Error
		if err != nil {
			return nil, err
		}
		pagination.Total = totalSize
		utils.CalPageTotal(totalSize,pagination)
		return companyIntros,nil
	}
}
// select all

func (c *CompanyIntroRepo) SelectIAll() ([]*example.CompanyIntro, error) {
	db := c.DB.Model(model.CompanyIntro{})
	var companyIntro []*example.CompanyIntro

	if err := db.Order("seq_no",false).
		Where("status= ? ",1).
		Select(companyIntroParams).Scan(&companyIntro).Error; err != nil {
		return nil, err
	}

	return companyIntro, nil
}


// 根据id 删除内容
func (c *CompanyIntroRepo) DeleteId(id int) error {
	tx, err := c.routine()
	if err != nil {
		return err
	}
	if err := tx.Delete(model.CompanyIntro{}, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// 批量 or 单个 删除
func (c *CompanyIntroRepo) Deletes(ids []string) error {
	tx, err := c.routine()
	if err != nil {
		return err
	}
	if err := tx.Delete(model.CompanyIntro{}, "id IN (?)", ids).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// 更新数据
func (c *CompanyIntroRepo) UpdateContent(companyIntro *model.CompanyIntroJson) error {
	tx, err := c.routine()
	//db := c.DB.Model(model.CompanyIntro{})
	//tx := db.Begin()
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback()
	//	}
	//}()
	//if err := tx.Error; err != nil {
	//	return err
	//}
	if err != nil {
		return err
	}
	if err := tx.Where("id = ?", companyIntro.Id).Update(&companyIntro).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// 事务前期步骤,抽取为一个方法
func (c *CompanyIntroRepo) routine() (*gorm.DB, error) {
	db := c.DB.Model(model.CompanyIntro{})
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil, err
	}

	return tx, nil
}
