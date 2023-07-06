package seller

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/assert"
	"os"
	"regexp"
	"testing"
)

func Test_repository_GetAll(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx context.Context
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	sellersSerialized, _ := os.ReadFile("../../test/resources/valid_sellers.json")
	var validSellers []domain.Seller
	if err := json.Unmarshal(sellersSerialized, &validSellers); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Seller
		wantErr bool
	}{
		{
			name: "Successfully get sellers",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
			},
			want:    validSellers,
			wantErr: false,
		},
		{
			name: "Error getting sellers",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
			},
			want:    []domain.Seller{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewSellerRepository(tt.fields.db)

			rows := sqlmock.NewRows([]string{"id", "cid", "company_name", "address", "telephone", "locality_id"})
			for _, s := range validSellers {
				rows.AddRow(s.ID, s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID)
			}

			if tt.wantErr {
				rows.RowError(0, sql.ErrNoRows)
			}

			mock.ExpectQuery(GetAllSellers).WillReturnRows(rows)

			got, err := r.GetAll(tt.args.ctx)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_repository_Get(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx context.Context
		id  int
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	sellerSerialized, _ := os.ReadFile("../../test/resources/valid_seller.json")
	var validSeller domain.Seller
	if err := json.Unmarshal(sellerSerialized, &validSeller); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Seller
		wantErr bool
	}{
		{
			name:   "Successfully get locality",
			fields: fields{db},
			args: args{
				ctx: ctx,
				id:  validSeller.ID,
			},
			want:    &validSeller,
			wantErr: false,
		},
		{
			name:   "Error nonexistent locality id",
			fields: fields{db},
			args: args{
				ctx: ctx,
				id:  999,
			},
			want:    &domain.Seller{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewSellerRepository(tt.fields.db)

			rows := sqlmock.NewRows([]string{"id", "cid", "company_name", "address", "telephone", "locality_id"}).
				AddRow(tt.want.ID, tt.want.CID, tt.want.CompanyName, tt.want.Address, tt.want.Telephone, tt.want.LocalityID)

			mock.ExpectQuery(GetSellerByID).
				WithArgs(tt.want.ID).
				WillReturnRows(rows)

			got, err := r.Get(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_repository_Exists(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx context.Context
		cid int
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	validSellerSerialized, _ := os.ReadFile("../../test/resources/valid_seller.json")
	var validSeller domain.Seller
	if err := json.Unmarshal(validSellerSerialized, &validSeller); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "Exists",
			fields: fields{db},
			args: args{
				ctx: ctx,
				cid: validSeller.ID,
			},
			want: true,
		},
		{
			name:   "Does not exist",
			fields: fields{db},
			args: args{
				ctx: ctx,
				cid: 999,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewSellerRepository(tt.fields.db)

			rows := sqlmock.NewRows([]string{"cid"}).
				AddRow(tt.args.cid)

			if !tt.want {
				rows.RowError(0, sql.ErrNoRows)
			}

			mock.ExpectQuery(ExistsSellerByCID).
				WithArgs(tt.args.cid).
				WillReturnRows(rows)

			got := r.Exists(tt.args.ctx, tt.args.cid)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_repository_Save(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx context.Context
		s   domain.Seller
	}

	validSellerSerialized, _ := os.ReadFile("../../test/resources/valid_seller.json")
	var validSeller domain.Seller
	if err := json.Unmarshal(validSellerSerialized, &validSeller); err != nil {
		t.Fatal(err)
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Successfully save seller",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				s:   validSeller,
			},
			want:    validSeller.ID,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewSellerRepository(tt.fields.db)

			mock.ExpectPrepare(regexp.QuoteMeta(SaveSeller))
			mock.ExpectExec(regexp.QuoteMeta(SaveSeller)).
				WithArgs(
					tt.args.s.CID,
					tt.args.s.CompanyName,
					tt.args.s.Address,
					tt.args.s.Telephone,
					tt.args.s.LocalityID,
				).
				WillReturnResult(sqlmock.NewResult(int64(tt.want), 1))

			got, err := r.Save(tt.args.ctx, tt.args.s)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_repository_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx context.Context
		s   domain.Seller
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	sellerSerialized, _ := os.ReadFile("../../test/resources/valid_seller.json")
	var validSeller domain.Seller
	if err := json.Unmarshal(sellerSerialized, &validSeller); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Successfully update locality",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				s:   validSeller,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewSellerRepository(tt.fields.db)

			mock.ExpectPrepare(regexp.QuoteMeta(UpdateSeller))
			mock.ExpectExec(regexp.QuoteMeta(UpdateSeller)).
				WithArgs(
					tt.args.s.CID,
					tt.args.s.CompanyName,
					tt.args.s.Address,
					tt.args.s.Telephone,
					tt.args.s.LocalityID,
					tt.args.s.ID,
				).
				WillReturnResult(sqlmock.NewResult(1, 1))

			err := r.Update(tt.args.ctx, tt.args.s)

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_repository_Delete(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx context.Context
		id  int
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Successfully delete seller",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				id:  6700,
			},
			wantErr: false,
		},
		{
			name: "No seller to delete",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				id:  6700,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewSellerRepository(tt.fields.db)

			rowsAffected := int64(1)
			if tt.wantErr {
				rowsAffected = 0
			}

			mock.ExpectPrepare(regexp.QuoteMeta(DeleteSellerByID))
			mock.ExpectExec(regexp.QuoteMeta(DeleteSellerByID)).
				WithArgs(
					tt.args.id,
				).
				WillReturnResult(sqlmock.NewResult(1, rowsAffected))

			err := r.Delete(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
