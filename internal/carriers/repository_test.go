package carriers_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/carrier"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/carriers"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGetAll(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("get_all_ok", func(t *testing.T) {
		expectedCarriers := []domain.Carrier{
			{
				ID:          1,
				CID:         "CID#1",
				CompanyName: "some name",
				Address:     "corrientes 800",
				Telephone:   "4567-4567",
				LocalityId:  6700,
			},
			{
				ID:          2,
				CID:         "CID#1",
				CompanyName: "some name",
				Address:     "corrientes 800",
				Telephone:   "4567-4567",
				LocalityId:  6701,
			},
		}
		r := carriers.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "cid", "company_name", "address", "telephone", "locality_id"})

		for _, expectedCarrier := range expectedCarriers {
			rows.AddRow(expectedCarrier.ID, expectedCarrier.CID, expectedCarrier.CompanyName, expectedCarrier.Address, expectedCarrier.Telephone, expectedCarrier.LocalityId)
		}
		mock.ExpectQuery("SELECT id, cid, company_name, address, telephone, locality_id FROM carriers").WillReturnRows(rows)

		carriersReceived, err := r.GetAll(ctx)

		assert.Equal(t, expectedCarriers, carriersReceived)
		assert.Nil(t, err)
	})

	t.Run("get_all_error", func(t *testing.T) {

		r := carriers.NewRepository(fields{db}.db)

		mock.ExpectQuery("SELECT id, cid, company_name, address, telephone, locality_id FROM carriers").
			WithArgs().
			WillReturnError(sql.ErrNoRows)

		carriersReceived, err := r.GetAll(ctx)

		assert.Equal(t, []domain.Carrier(nil), carriersReceived)
		assert.NotNil(t, err)
	})
}
func TestRepositoryExists(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("exists_true", func(t *testing.T) {

		expectedCarrier := &domain.Carrier{
			ID:          1,
			CID:         "CID#1",
			CompanyName: "some name",
			Address:     "corrientes 800",
			Telephone:   "4567-4567",
			LocalityId:  6700,
		}

		r := carriers.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"cid"}).
			AddRow(expectedCarrier.CID)

		mock.ExpectQuery("SELECT cid FROM carriers WHERE cid=?").WithArgs(expectedCarrier.CID).WillReturnRows(rows)

		carrierExists := r.Exists(ctx, expectedCarrier.CID)

		assert.True(t, carrierExists)
	})

	t.Run("exists_false", func(t *testing.T) {

		expectedCarrier := &domain.Carrier{
			ID:          1,
			CID:         "CID#1",
			CompanyName: "some name",
			Address:     "corrientes 800",
			Telephone:   "4567-4567",
			LocalityId:  6700,
		}

		r := carriers.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"cid"}).
			AddRow(expectedCarrier.CID)

		mock.ExpectQuery("SELECT cid FROM carriers WHERE cid=?").WithArgs(expectedCarrier.CID).WillReturnRows(rows)

		carrierExists := r.Exists(ctx, "CID#2")

		assert.False(t, carrierExists)
	})
}

func TestRepositorySave(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("save_ok", func(t *testing.T) {
		expectedCarrier := &domain.Carrier{
			ID:          1,
			CID:         "CID#1",
			CompanyName: "some name",
			Address:     "corrientes 800",
			Telephone:   "4567-4567",
			LocalityId:  6700,
		}

		r := carriers.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO carriers(cid, company_name, address, telephone, locality_id) VALUES (?,?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO carriers(cid, company_name, address, telephone, locality_id) VALUES (?,?,?,?,?)")).
			WithArgs(expectedCarrier.CID, expectedCarrier.CompanyName, expectedCarrier.Address, expectedCarrier.Telephone, expectedCarrier.LocalityId).
			WillReturnResult(sqlmock.NewResult(1, 1))

		id, err := r.Save(ctx, *expectedCarrier)
		assert.Equal(t, expectedCarrier.ID, id)
		assert.Nil(t, err)
	})

	t.Run("save_error_exec", func(t *testing.T) {
		expectedCarrier := &domain.Carrier{
			ID:          1,
			CID:         "CID#1",
			CompanyName: "some name",
			Address:     "corrientes 800",
			Telephone:   "4567-4567",
			LocalityId:  6700,
		}

		r := carriers.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO carriers(cid, company_name, address, telephone, locality_id) VALUES (?,?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO carriers(cid, company_name, address, telephone, locality_id) VALUES (?,?,?,?,?)")).
			WithArgs(expectedCarrier.CID, expectedCarrier.CompanyName, expectedCarrier.Address, expectedCarrier.Telephone, expectedCarrier.LocalityId).
			WillReturnError(sql.ErrNoRows)

		_, err := r.Save(ctx, *expectedCarrier)

		assert.NotNil(t, err)
	})

	t.Run("save_error_rowlsAffected0", func(t *testing.T) {

		expectedCarrier := domain.Carrier{
			ID:          1,
			CID:         "CID#1",
			CompanyName: "some name",
			Address:     "corrientes 800",
			Telephone:   "4567-4567",
			LocalityId:  6700,
		}

		r := carriers.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO carriers(cid, company_name, address, telephone, locality_id) VALUES (?,?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO carriers(cid, company_name, address, telephone, locality_id) VALUES (?,?,?,?,?)")).
			WithArgs(expectedCarrier.CID, expectedCarrier.CompanyName, expectedCarrier.Address, expectedCarrier.Telephone, expectedCarrier.LocalityId).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		_, err := r.Save(ctx, expectedCarrier)

		assert.NotNil(t, err)
	})

	t.Run("save_error_prepare", func(t *testing.T) {
		expectedCarrier := &domain.Carrier{
			ID:          1,
			CID:         "CID#1",
			CompanyName: "some name",
			Address:     "corrientes 800",
			Telephone:   "4567-4567",
			LocalityId:  6700,
		}

		r := carriers.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO carriers(cid, company_name, address, telephone, locality_id) VALUES (?,?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO carriers(cid, company_name, address, telephone, locality_id) VALUES (?,?,?,?,?)")).
			WithArgs(expectedCarrier.CID, expectedCarrier.CompanyName, expectedCarrier.Address, expectedCarrier.Telephone, expectedCarrier.LocalityId).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		_, err := r.Save(ctx, *expectedCarrier)

		assert.NotNil(t, err)

	})
}

func TestRepositoryGetLocalityById(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("get_by_locality_id_true", func(t *testing.T) {

		expectedLocality := domain.Locality{
			ID:           1,
			ProvinceName: "Teste",
			LocalityName: "Teste-2",
		}

		r := carriers.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "province_name", "locality_name"}).
			AddRow(expectedLocality.ID, expectedLocality.ProvinceName, expectedLocality.LocalityName)

		mock.ExpectQuery("SELECT localities.id, localities.province_name, localities.locality_name  FROM localities WHERE id =?").WithArgs(expectedLocality.ID).WillReturnRows(rows)

		locality, error := r.GetLocalityById(ctx, expectedLocality.ID)

		assert.Equal(t, locality, expectedLocality)
		assert.Nil(t, error)
	})

	t.Run("get_by_locality_id_false", func(t *testing.T) {

		r := carriers.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "province_name", "locality_name"})

		mock.ExpectQuery("SELECT localities.id, localities.province_name, localities.locality_name  FROM localities WHERE id =?").WillReturnRows(rows)

		locality, error := r.GetLocalityById(ctx, 10)

		assert.Equal(t, locality, domain.Locality{})
		assert.NotNil(t, error)
	})
}

func TestRepositoryCountCarriersByLocalityId(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("get_count_ok", func(t *testing.T) {

		expectedCarriers := []domain.Carrier{
			{
				ID:          1,
				CID:         "CID#1",
				CompanyName: "some name",
				Address:     "corrientes 800",
				Telephone:   "4567-4567",
				LocalityId:  6701,
			},
			{
				ID:          2,
				CID:         "CID#1",
				CompanyName: "some name",
				Address:     "corrientes 800",
				Telephone:   "4567-4567",
				LocalityId:  6701,
			},
		}

		r := carriers.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"COUNT(id)"})

		for _, expectedCarrier := range expectedCarriers {
			rows.AddRow(expectedCarrier.ID)
		}
		mock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(id) FROM carriers WHERE locality_id = ?")).WillReturnRows(rows)

		carriersReceived, err := r.GetCountCarriersByLocalityId(ctx, 6701)

		assert.Equal(t, len(expectedCarriers), carriersReceived)
		assert.Nil(t, err)

	})

	t.Run("get_count_error", func(t *testing.T) {

		expectedCarriers := []domain.Carrier{
			{
				ID:          1,
				CID:         "CID#1",
				CompanyName: "some name",
				Address:     "corrientes 800",
				Telephone:   "4567-4567",
				LocalityId:  6701,
			},
			{
				ID:          2,
				CID:         "CID#1",
				CompanyName: "some name",
				Address:     "corrientes 800",
				Telephone:   "4567-4567",
				LocalityId:  6701,
			},
		}

		r := carriers.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"COUNT(id)"})

		for _, expectedCarrier := range expectedCarriers {
			rows.AddRow(expectedCarrier.ID)
		}
		mock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(id) FROM carriers WHERE locality_id = ?")).WillReturnRows(rows)

		carriersReceived, err := r.GetCountCarriersByLocalityId(ctx, 6701)

		assert.Equal(t, 2, carriersReceived)
		assert.Nil(t, err)
	})
	t.Run("get_count_false", func(t *testing.T) {

		r := carriers.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"COUNT(id)"})
		mock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(id) FROM carriers WHERE locality_id = ?")).WillReturnRows(rows)

		count, error := r.GetCountCarriersByLocalityId(ctx, 0)

		assert.Equal(t, count, 0)
		assert.Nil(t, error)
	})
}

func TestRepositoryCountAndDataByLocality(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("get_datas_success", func(t *testing.T) {
		actualDatas := []dtos.DataLocalityAndCarrier{
			{
				Id:           12,
				LocalityName: "Teste",
				CountCarrier: 1,
			},
		}
		r := carriers.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "locality_name", "count_carrier"})

		for _, actualData := range actualDatas {
			rows.AddRow(actualData.Id, actualData.LocalityName, actualData.CountCarrier)
		}
		mock.ExpectQuery(regexp.QuoteMeta("SELECT l.id, l.locality_name, (SELECT count(id) FROM carriers c where c.locality_id = l.id) AS count_carrier FROM localities l LIMIT 10")).WillReturnRows(rows)

		datasReceived, err := r.GetCountAndDataByLocality(ctx)

		assert.Equal(t, actualDatas, datasReceived)
		assert.Nil(t, err)
	})

	t.Run("get_datas_fail", func(t *testing.T) {

		actualDatas := []dtos.DataLocalityAndCarrier{}
		r := carriers.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "locality_name", "count_carrier"})

		mock.ExpectQuery(regexp.QuoteMeta("SELECT l.id, l.locality_name, (SELECT count(id) FROM carriers c where c.locality_id = l.id) AS count_carrier FROM localities l LIMIT 10")).WillReturnRows(rows)

		datasReceived, err := r.GetCountAndDataByLocality(ctx)

		assert.Equal(t, actualDatas, datasReceived)
		assert.Nil(t, err)
	})
}
