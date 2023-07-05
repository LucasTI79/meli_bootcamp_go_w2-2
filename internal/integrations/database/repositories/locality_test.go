package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"os"
	"regexp"
	"testing"
)

func Test_localityRepository_GetAll(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx context.Context
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	validLocalitiesSerialized, _ := os.ReadFile("../../../../test/resources/valid_localities.json")
	var validLocalities []entities.Locality
	if err := json.Unmarshal(validLocalitiesSerialized, &validLocalities); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.Locality
		wantErr bool
	}{
		{
			name: "Successfully get localities",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
			},
			want:    validLocalities,
			wantErr: false,
		},
		{
			name: "Error getting localities",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
			},
			want:    []entities.Locality{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewLocalityRepository(tt.fields.db)

			rows := sqlmock.NewRows([]string{"id", "country_name", "province_name", "locality_name"})
			for _, locality := range validLocalities {
				rows.AddRow(locality.ID, locality.CountryName, locality.ProvinceName, locality.LocalityName)
			}
			if tt.wantErr {
				rows.RowError(0, sql.ErrNoRows)
			}

			mock.ExpectQuery(GetAllLocalities).
				WithArgs().
				WillReturnRows(rows)

			got, err := r.GetAll(tt.args.ctx)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_localityRepository_Get(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx context.Context
		id  int
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	localitySerialized, _ := os.ReadFile("../../../../test/resources/valid_locality.json")
	var validLocality entities.Locality
	if err := json.Unmarshal(localitySerialized, &validLocality); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.Locality
		wantErr bool
	}{
		{
			name:   "Successfully get locality",
			fields: fields{db},
			args: args{
				ctx: ctx,
				id:  validLocality.ID,
			},
			want:    validLocality,
			wantErr: false,
		},
		{
			name:   "Error nonexistent locality id",
			fields: fields{db},
			args: args{
				ctx: ctx,
				id:  999,
			},
			want:    entities.Locality{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewLocalityRepository(tt.fields.db)

			rows := sqlmock.NewRows([]string{"id", "country_name", "province_name", "locality_name"}).
				AddRow(tt.want.ID, tt.want.CountryName, tt.want.ProvinceName, tt.want.LocalityName)

			mock.ExpectQuery(GetLocalityByID).
				WithArgs(tt.want.ID).
				WillReturnRows(rows)

			got, err := r.Get(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_localityRepository_Exists(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx context.Context
		id  int
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	localitySerialized, _ := os.ReadFile("../../../../test/resources/valid_locality.json")
	var locality entities.Locality
	if err := json.Unmarshal(localitySerialized, &locality); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Exists",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				id:  locality.ID,
			},
			want: true,
		},
		{
			name: "Does not exists",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				id:  locality.ID,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewLocalityRepository(tt.fields.db)

			rows := sqlmock.NewRows([]string{"id"}).
				AddRow(locality.ID)

			if !tt.want {
				rows.RowError(0, sql.ErrNoRows)
			}

			mock.ExpectQuery(ExistsLocalityByID).
				WithArgs(tt.args.id).
				WillReturnRows(rows)

			got := r.Exists(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_localityRepository_Save(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx      context.Context
		locality entities.Locality
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	localitySerialized, _ := os.ReadFile("../../../../test/resources/valid_locality.json")
	var locality entities.Locality
	if err := json.Unmarshal(localitySerialized, &locality); err != nil {
		t.Fatal(err)
	}

	id := locality.ID

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Successfully save locality",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:      ctx,
				locality: locality,
			},
			want:    id,
			wantErr: false,
		},
		// TODO Return errors
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewLocalityRepository(tt.fields.db)

			mock.ExpectPrepare(regexp.QuoteMeta(SaveLocality))
			mock.ExpectExec(regexp.QuoteMeta(SaveLocality)).
				WithArgs(
					tt.args.locality.CountryName,
					tt.args.locality.ProvinceName,
					tt.args.locality.LocalityName,
				).
				WillReturnResult(sqlmock.NewResult(int64(tt.want), 1))

			got, err := r.Save(tt.args.ctx, tt.args.locality)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_localityRepository_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx      context.Context
		locality entities.Locality
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	localitySerialized, _ := os.ReadFile("../../../../test/resources/valid_locality.json")
	var locality entities.Locality
	if err := json.Unmarshal(localitySerialized, &locality); err != nil {
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
				ctx:      ctx,
				locality: locality,
			},
			wantErr: false,
		},
		//TODO Return errors
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewLocalityRepository(tt.fields.db)

			mock.ExpectPrepare(regexp.QuoteMeta(UpdateLocality))
			mock.ExpectExec(regexp.QuoteMeta(UpdateLocality)).
				WithArgs(
					tt.args.locality.CountryName,
					tt.args.locality.ProvinceName,
					tt.args.locality.LocalityName,
					tt.args.locality.ID,
				).
				WillReturnResult(sqlmock.NewResult(1, 1))

			err := r.Update(tt.args.ctx, tt.args.locality)

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_localityRepository_Delete(t *testing.T) {
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
			name: "Successfully delete locality",
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
			name: "No locality to delete",
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
			r := NewLocalityRepository(tt.fields.db)

			rowsAffected := int64(1)
			if tt.wantErr {
				rowsAffected = 0
			}

			mock.ExpectPrepare(regexp.QuoteMeta(DeleteLocalityByID))
			mock.ExpectExec(regexp.QuoteMeta(DeleteLocalityByID)).
				WithArgs(
					tt.args.id,
				).
				WillReturnResult(sqlmock.NewResult(1, rowsAffected))

			err := r.Delete(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_localityRepository_GetNumberOfSellers(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		id  int
	}

	// Para funcionar com * na query
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	ctx := context.TODO()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Successfully get number of sellers",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				id:  6700,
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewLocalityRepository(tt.fields.db)

			rows := sqlmock.NewRows([]string{"count(*)"}).AddRow(tt.want)

			mock.ExpectQuery(CountLocalitySellersByID).
				WithArgs(tt.args.id).
				WillReturnRows(rows)

			got, err := r.GetNumberOfSellers(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)

		})
	}
}
