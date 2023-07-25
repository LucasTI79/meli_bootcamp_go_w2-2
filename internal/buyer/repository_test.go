package buyer

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

func Test_buyerRepository_GetAll(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx context.Context
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	validBuyersSerialized, _ := os.ReadFile("../../test/resources/valid_buyers.json")
	var validBuyers []domain.Buyer
	if err := json.Unmarshal(validBuyersSerialized, &validBuyers); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Buyer
		wantErr bool
	}{
		{
			name: "Successfully get buyers",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
			},
			want:    validBuyers,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewBuyerRepository(tt.fields.db)

			rows := sqlmock.NewRows([]string{"cardNumberID", "card_number_id", "first_name", "last_name"})
			for _, buyer := range validBuyers {
				rows.AddRow(buyer.ID, buyer.CardNumberID, buyer.FirstName, buyer.LastName)
			}
			if tt.wantErr {
				rows.RowError(0, sql.ErrNoRows)
			}

			mock.ExpectQuery(GetAllBuyers).
				WithArgs().
				WillReturnRows(rows)

			got, err := r.GetAll(tt.args.ctx)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}

	t.Run("Error getting buyer by cardNumberID", func(t *testing.T) {

		db, _, _ := sqlmock.New()
		ctx := context.TODO()

		r := NewBuyerRepository(db)

		_, err := r.GetAll(ctx)

		assert.Equal(t, true, err != nil)

	})
}

func Test_buyerRepository_Get(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx context.Context
		id  int
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	buyerSerialized, _ := os.ReadFile("../../test/resources/valid_buyer.json")
	var validBuyer domain.Buyer
	if err := json.Unmarshal(buyerSerialized, &validBuyer); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Buyer
		wantErr bool
	}{
		{
			name:   "Successfully get buyer",
			fields: fields{db},
			args: args{
				ctx: ctx,
				id:  validBuyer.ID,
			},
			want:    validBuyer,
			wantErr: false,
		},
		{
			name:   "Error nonexistent buyer cardNumberID",
			fields: fields{db},
			args: args{
				ctx: ctx,
				id:  999,
			},
			want:    domain.Buyer{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewBuyerRepository(tt.fields.db)

			rows := sqlmock.NewRows([]string{"cardNumberID", "card_number_id", "first_name", "last_name"}).
				AddRow(tt.want.ID, tt.want.CardNumberID, tt.want.FirstName, tt.want.LastName)

			mock.ExpectQuery(GetBuyerByID).
				WithArgs(tt.want.ID).
				WillReturnRows(rows)

			got, err := r.Get(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_buyerRepository_CardNumberExists(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx          context.Context
		cardNumberID string
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	buyerSerialized, _ := os.ReadFile("../../test/resources/valid_buyer.json")
	var buyer domain.Buyer
	if err := json.Unmarshal(buyerSerialized, &buyer); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "CardNumber Exists",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:          ctx,
				cardNumberID: buyer.CardNumberID,
			},
			want: true,
		},
		{
			name: "CardNumber Does not exists",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:          ctx,
				cardNumberID: buyer.CardNumberID,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewBuyerRepository(tt.fields.db)

			rows := sqlmock.NewRows([]string{"card_number_id"}).
				AddRow(buyer.CardNumberID)

			if !tt.want {
				rows.RowError(0, sql.ErrNoRows)
			}

			mock.ExpectQuery(ExistsBuyerByID).
				WithArgs(tt.args.cardNumberID).
				WillReturnRows(rows)

			got := r.CardNumberExists(tt.args.ctx, tt.args.cardNumberID)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_buyerRepository_Save(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx   context.Context
		buyer domain.Buyer
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	buyerSerialized, _ := os.ReadFile("../../test/resources/valid_buyer.json")
	var buyer domain.Buyer
	if err := json.Unmarshal(buyerSerialized, &buyer); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Successfully save buyer",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:   ctx,
				buyer: buyer,
			},
			want:    buyer.ID,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewBuyerRepository(tt.fields.db)

			mock.ExpectPrepare(regexp.QuoteMeta(SaveBuyer))
			mock.ExpectExec(regexp.QuoteMeta(SaveBuyer)).
				WithArgs(
					tt.args.buyer.CardNumberID,
					tt.args.buyer.FirstName,
					tt.args.buyer.LastName,
				).
				WillReturnResult(sqlmock.NewResult(int64(tt.want), 1))

			got, err := r.Save(tt.args.ctx, tt.args.buyer)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}

	t.Run("Error preparing query", func(t *testing.T) {
		db, _, _ := sqlmock.New()
		ctx := context.TODO()

		r := NewBuyerRepository(db)
		db.Close()

		r.Save(ctx, domain.Buyer{})

	})

	t.Run("Error executing query", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		ctx := context.TODO()

		r := NewBuyerRepository(db)

		mock.ExpectPrepare(regexp.QuoteMeta(SaveBuyer))
		mock.ExpectExec(regexp.QuoteMeta(SaveBuyer)).
			WithArgs(
				buyer.CardNumberID,
				buyer.FirstName,
				buyer.LastName,
			).
			WillReturnError(sql.ErrNoRows)

		got, err := r.Save(ctx, buyer)

		assert.Equal(t, 0, got)
		assert.Equal(t, true, err != nil)
	})
}

func Test_buyerRepository_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx   context.Context
		buyer domain.Buyer
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	buyerSerialized, _ := os.ReadFile("../../test/resources/valid_buyer.json")
	var buyer domain.Buyer
	if err := json.Unmarshal(buyerSerialized, &buyer); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Successfully update buyer",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:   ctx,
				buyer: buyer,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewBuyerRepository(tt.fields.db)

			mock.ExpectPrepare(regexp.QuoteMeta(UpdateBuyer))
			mock.ExpectExec(regexp.QuoteMeta(UpdateBuyer)).
				WithArgs(
					tt.args.buyer.CardNumberID,
					tt.args.buyer.FirstName,
					tt.args.buyer.LastName,
					tt.args.buyer.ID,
				).
				WillReturnResult(sqlmock.NewResult(1, 1))

			err := r.Update(tt.args.ctx, tt.args.buyer)

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}

	t.Run("Error preparing query", func(t *testing.T) {
		db, _, _ := sqlmock.New()
		ctx := context.TODO()

		r := NewBuyerRepository(db)

		err := r.Update(ctx, domain.Buyer{})

		assert.Equal(t, true, err != nil)
	})

	t.Run("Error executing query", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		ctx := context.TODO()

		r := NewBuyerRepository(db)

		mock.ExpectPrepare(regexp.QuoteMeta(UpdateBuyer))
		mock.ExpectExec(regexp.QuoteMeta(UpdateBuyer)).
			WithArgs(
				buyer.CardNumberID,
				buyer.FirstName,
				buyer.LastName,
				buyer.ID,
			).
			WillReturnError(sql.ErrNoRows)

		err := r.Update(ctx, buyer)

		assert.Equal(t, true, err != nil)
	})
}

func Test_buyerRepository_Delete(t *testing.T) {
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
			name: "Successfully delete buyer",
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
			name: "No buyer to delete",
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
			r := NewBuyerRepository(tt.fields.db)

			rowsAffected := int64(1)
			if tt.wantErr {
				rowsAffected = 0
			}

			mock.ExpectPrepare(regexp.QuoteMeta(DeleteBuyerByID))
			mock.ExpectExec(regexp.QuoteMeta(DeleteBuyerByID)).
				WithArgs(
					tt.args.id,
				).
				WillReturnResult(sqlmock.NewResult(1, rowsAffected))

			err := r.Delete(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}

	t.Run("Error preparing query", func(t *testing.T) {
		db, _, _ := sqlmock.New()
		ctx := context.TODO()

		r := NewBuyerRepository(db)
		db.Close()

		err := r.Delete(ctx, 6700)

		assert.Equal(t, true, err != nil)
	})

	t.Run("Error executing query", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		ctx := context.TODO()

		r := NewBuyerRepository(db)

		mock.ExpectPrepare(regexp.QuoteMeta(DeleteBuyerByID))
		mock.ExpectExec(regexp.QuoteMeta(DeleteBuyerByID)).
			WithArgs(
				6700,
			).WillReturnError(sql.ErrNoRows)

		err := r.Delete(ctx, 6700)

		assert.Equal(t, true, err != nil)
	})
}
