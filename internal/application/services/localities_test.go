package services

import (
	"context"
	"encoding/json"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

func Test_localityService_GetAll(t *testing.T) {
	type args struct {
		ctx                  *context.Context
		expectedGetAllResult []entities.Locality
		expectedGetAllError  error
		expectedGetAllCalls  int
	}

	ctx := context.TODO()

	var expectedLocalities []entities.Locality
	expectedLocalitiesSerialized, _ := os.ReadFile("../../../test/resources/valid_localities.json")
	if err := json.Unmarshal(expectedLocalitiesSerialized, &expectedLocalities); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		args    args
		want    []entities.Locality
		wantErr error
	}{
		{
			name: "Successfully get all localities",
			args: args{
				ctx:                  &ctx,
				expectedGetAllResult: expectedLocalities,
				expectedGetAllError:  nil,
				expectedGetAllCalls:  1,
			},
			want:    expectedLocalities,
			wantErr: nil,
		},
		{
			name: "Error getting all",
			args: args{
				ctx:                  &ctx,
				expectedGetAllResult: []entities.Locality{},
				expectedGetAllError:  assert.AnError,
				expectedGetAllCalls:  1,
			},
			want:    []entities.Locality{},
			wantErr: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sellerRepositoryMock := mocks.NewMockLocalityRepository(t)
			sellerRepositoryMock.On("GetAll", *tt.args.ctx).Return(tt.args.expectedGetAllResult, tt.args.expectedGetAllError)

			service := NewLocalityService(sellerRepositoryMock)
			got, err := service.GetAll(tt.args.ctx)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)

			sellerRepositoryMock.AssertNumberOfCalls(t, "GetAll", tt.args.expectedGetAllCalls)
		})
	}
}

func Test_localityService_Get(t *testing.T) {
	type args struct {
		ctx               *context.Context
		id                int
		expectedGetResult entities.Locality
		expectedGetError  error
		expectedGetCalls  int
	}

	ctx := context.TODO()

	var expectedLocality entities.Locality
	expectedLocalitySerialized, _ := os.ReadFile("../../../test/resources/valid_locality.json")
	if err := json.Unmarshal(expectedLocalitySerialized, &expectedLocality); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		args    args
		want    entities.Locality
		wantErr error
	}{
		{
			name: "Successfully get buyer from db",
			args: args{
				ctx:               &ctx,
				id:                expectedLocality.ID,
				expectedGetResult: expectedLocality,
				expectedGetError:  nil,
				expectedGetCalls:  1,
			},
			want:    expectedLocality,
			wantErr: nil,
		},
		{
			name: "Buyer not found in db",
			args: args{
				ctx:               &ctx,
				id:                expectedLocality.ID,
				expectedGetResult: entities.Locality{},
				expectedGetError:  ErrNotFound,
				expectedGetCalls:  1,
			},
			want:    entities.Locality{},
			wantErr: ErrNotFound,
		},
		{
			name: "Error connecting db",
			args: args{
				ctx:               &ctx,
				id:                expectedLocality.ID,
				expectedGetResult: entities.Locality{},
				expectedGetError:  assert.AnError,
				expectedGetCalls:  1,
			},
			want:    entities.Locality{},
			wantErr: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			localityRepositoryMock := mocks.NewMockLocalityRepository(t)
			localityRepositoryMock.On("Get", *tt.args.ctx, mock.AnythingOfType("int")).Return(tt.args.expectedGetResult, tt.args.expectedGetError)

			service := NewLocalityService(localityRepositoryMock)
			got, err := service.Get(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)

			localityRepositoryMock.AssertNumberOfCalls(t, "Get", tt.args.expectedGetCalls)

		})
	}
}

func Test_localityService_Create(t *testing.T) {
	type args struct {
		ctx                  *context.Context
		locality             entities.Locality
		expectedExistsResult bool
		expectedExistsCalls  int
		expectedSaveResult   int
		expectedSaveError    error
		expectedSaveCalls    int
	}

	ctx := context.TODO()

	var expectedLocality entities.Locality
	expectedLocalitySerialized, _ := os.ReadFile("../../../test/resources/valid_locality.json")
	if err := json.Unmarshal(expectedLocalitySerialized, &expectedLocality); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		args    args
		want    entities.Locality
		wantErr error
	}{
		{
			name: "Successfully create locality",
			args: args{
				ctx:                  &ctx,
				locality:             expectedLocality,
				expectedExistsResult: false,
				expectedExistsCalls:  1,
				expectedSaveResult:   expectedLocality.ID,
				expectedSaveError:    nil,
				expectedSaveCalls:    1,
			},
			want:    expectedLocality,
			wantErr: nil,
		},
		{
			name: "Error duplicated locality",
			args: args{
				ctx:                  &ctx,
				locality:             expectedLocality,
				expectedExistsResult: true,
				expectedExistsCalls:  1,
				expectedSaveResult:   0,
				expectedSaveError:    nil,
				expectedSaveCalls:    0,
			},
			want:    entities.Locality{},
			wantErr: ErrConflict,
		},
		{
			name: "Error saving locality",
			args: args{
				ctx:                  &ctx,
				locality:             expectedLocality,
				expectedExistsResult: false,
				expectedExistsCalls:  1,
				expectedSaveError:    assert.AnError,
				expectedSaveCalls:    1,
			},
			want:    entities.Locality{},
			wantErr: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			localityRepositoryMock := mocks.NewMockLocalityRepository(t)

			localityRepositoryMock.On(
				"Exists",
				*tt.args.ctx,
				mock.AnythingOfType("int"),
			).Return(
				tt.args.expectedExistsResult,
			)

			localityRepositoryMock.On(
				"Save",
				*tt.args.ctx,
				mock.AnythingOfType("entities.Locality"),
			).Return(
				tt.args.expectedSaveResult,
				tt.args.expectedSaveError,
			)

			service := NewLocalityService(localityRepositoryMock)
			got, err := service.Create(tt.args.ctx, tt.args.locality)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)

			localityRepositoryMock.AssertNumberOfCalls(t, "Exists", tt.args.expectedExistsCalls)
			localityRepositoryMock.AssertNumberOfCalls(t, "Save", tt.args.expectedSaveCalls)
		})
	}
}

func Test_localityService_Update(t *testing.T) {
	type args struct {
		ctx                   *context.Context
		id                    int
		updateLocalityRequest dtos.UpdateLocalityRequestDTO
		expectedGetResult     entities.Locality
		expectedGetError      error
		expectedGetCalls      int
		expectedExistsResult  bool
		expectedExistsCalls   int
		expectedUpdateError   error
		expectedUpdateCalls   int
	}

	ctx := context.TODO()

	var originalLocality entities.Locality
	originalLocalitySerialized, _ := os.ReadFile("../../../test/resources/valid_locality.json")
	err := json.Unmarshal(originalLocalitySerialized, &originalLocality)
	if err != nil {
		t.Fatal(err)
	}

	newCountryName := "Brasil"
	newProvinceName := "Rio de Janeiro"
	newLocalityName := "Rio de Janeiro"

	updateLocalityRequest := dtos.UpdateLocalityRequestDTO{
		CountryName:  &newCountryName,
		ProvinceName: &newProvinceName,
		LocalityName: &newLocalityName,
	}

	updatedLocality := entities.Locality{
		ID:           originalLocality.ID,
		CountryName:  newCountryName,
		ProvinceName: newProvinceName,
		LocalityName: newLocalityName,
	}

	tests := []struct {
		name    string
		args    args
		want    entities.Locality
		wantErr error
	}{
		{
			name: "Successfully updating all fields",
			args: args{
				ctx:                   &ctx,
				id:                    originalLocality.ID,
				updateLocalityRequest: updateLocalityRequest,
				expectedGetResult:     originalLocality,
				expectedGetError:      nil,
				expectedGetCalls:      1,
				expectedExistsResult:  false,
				expectedExistsCalls:   1,
				expectedUpdateError:   nil,
				expectedUpdateCalls:   1,
			},
			want:    updatedLocality,
			wantErr: nil,
		},
		{
			name: "Error locality doesn't exists",
			args: args{
				ctx:                   &ctx,
				id:                    originalLocality.ID,
				updateLocalityRequest: updateLocalityRequest,
				expectedGetResult:     entities.Locality{},
				expectedGetError:      ErrNotFound,
				expectedGetCalls:      1,
			},
			want:    entities.Locality{},
			wantErr: ErrNotFound,
		},
		{
			name: "Error duplicated locality_id",
			args: args{
				ctx:                   &ctx,
				id:                    originalLocality.ID,
				updateLocalityRequest: updateLocalityRequest,
				expectedGetResult:     originalLocality,
				expectedGetError:      nil,
				expectedGetCalls:      1,
				expectedExistsResult:  true,
				expectedExistsCalls:   1,
			},
			want:    entities.Locality{},
			wantErr: ErrConflict,
		},
		{
			name: "Error updating locality",
			args: args{
				ctx:                   &ctx,
				id:                    originalLocality.ID,
				updateLocalityRequest: updateLocalityRequest,
				expectedGetResult:     originalLocality,
				expectedGetError:      nil,
				expectedGetCalls:      1,
				expectedExistsResult:  false,
				expectedExistsCalls:   1,
				expectedUpdateError:   assert.AnError,
				expectedUpdateCalls:   1,
			},
			want:    entities.Locality{},
			wantErr: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			localityRepositoryMock := mocks.NewMockLocalityRepository(t)
			service := NewLocalityService(localityRepositoryMock)

			localityRepositoryMock.On("Get", *tt.args.ctx, mock.AnythingOfType("int")).Return(tt.args.expectedGetResult, tt.args.expectedGetError)
			localityRepositoryMock.On("Exists", *tt.args.ctx, mock.AnythingOfType("int")).Return(tt.args.expectedExistsResult)
			localityRepositoryMock.On("Update", *tt.args.ctx, mock.AnythingOfType("entities.Locality")).Return(tt.args.expectedUpdateError)

			newBuyer, err := service.Update(tt.args.ctx, tt.args.id, tt.args.updateLocalityRequest)

			assert.Equal(t, tt.want, newBuyer)
			assert.Equal(t, tt.wantErr, err)

			localityRepositoryMock.AssertNumberOfCalls(t, "Get", tt.args.expectedGetCalls)
			localityRepositoryMock.AssertNumberOfCalls(t, "Exists", tt.args.expectedExistsCalls)
			localityRepositoryMock.AssertNumberOfCalls(t, "Update", tt.args.expectedUpdateCalls)

		})
	}
}
func Test_localityService_Delete(t *testing.T) {
	type args struct {
		ctx                 *context.Context
		id                  int
		expectedGetResult   entities.Locality
		expectedGetError    error
		expectedGetCalls    int
		expectedDeleteError error
		expectedDeleteCalls int
	}

	ctx := context.TODO()

	var expectedLocality entities.Locality
	expectedLocalitySerialized, _ := os.ReadFile("../../../test/resources/valid_locality.json")
	if err := json.Unmarshal(expectedLocalitySerialized, &expectedLocality); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Successfully deleting locality",
			args: args{
				ctx:                 &ctx,
				id:                  1,
				expectedGetResult:   expectedLocality,
				expectedGetError:    nil,
				expectedGetCalls:    1,
				expectedDeleteError: nil,
				expectedDeleteCalls: 1,
			},
			wantErr: nil,
		},
		{
			name: "Error getting locality",
			args: args{
				ctx:                 &ctx,
				id:                  1,
				expectedGetResult:   entities.Locality{},
				expectedGetError:    assert.AnError,
				expectedGetCalls:    1,
				expectedDeleteError: nil,
				expectedDeleteCalls: 0,
			},
			wantErr: assert.AnError,
		},
		{
			name: "Error deleting locality",
			args: args{
				ctx:                 &ctx,
				id:                  1,
				expectedGetError:    nil,
				expectedGetCalls:    1,
				expectedDeleteError: assert.AnError,
				expectedDeleteCalls: 0,
			},
			wantErr: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			localityRepositoryMock := mocks.NewMockLocalityRepository(t)
			localityRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(tt.args.expectedGetResult, tt.args.expectedGetError)
			localityRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(tt.args.expectedDeleteError)

			service := NewLocalityService(localityRepositoryMock)
			err := service.Delete(&ctx, tt.args.id)

			assert.Equal(t, tt.wantErr, err)
			localityRepositoryMock.On("Get", tt.args.expectedGetCalls)
			localityRepositoryMock.On("Delete", tt.args.expectedDeleteCalls)
		})
	}
}

func Test_localityService_CountSellers(t *testing.T) {
	type args struct {
		ctx                        *context.Context
		id                         int
		expectedCountSellersResult int
		expectedCountSellersError  error
		expectedCountSellersCalls  int
	}

	ctx := context.TODO()

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name: "Successfully count sellers",
			args: args{
				ctx:                        &ctx,
				id:                         1,
				expectedCountSellersResult: 1,
				expectedCountSellersError:  nil,
				expectedCountSellersCalls:  1,
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "Error count sellers",
			args: args{
				ctx:                        &ctx,
				id:                         1,
				expectedCountSellersResult: 0,
				expectedCountSellersError:  assert.AnError,
				expectedCountSellersCalls:  1,
			},
			want:    0,
			wantErr: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			localityRepositoryMock := mocks.NewMockLocalityRepository(t)
			localityRepositoryMock.On("CountSellers", ctx, mock.AnythingOfType("int")).Return(tt.args.expectedCountSellersResult, tt.args.expectedCountSellersError)

			service := NewLocalityService(localityRepositoryMock)
			got, err := service.CountSellers(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)

			localityRepositoryMock.On("CountSellers", tt.args.expectedCountSellersCalls)

		})
	}
}
