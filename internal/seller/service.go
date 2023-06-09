package seller

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("seller not found")
	ErrConflict = errors.New("409 Conflict: Seller with CID already exists")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Seller, error)
	Get(ctx context.Context, id int) (*domain.Seller, error)
	Save(ctx context.Context, CID int, CompanyName, Adress, Telephone string) (*domain.Seller, error)
	Update(ctx context.Context, id, CID int, CompanyName, Adress, Telephone string) (*domain.Seller, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	sellerRepository Repository
}

func NewService(r Repository) Service {
	return &service{
		sellerRepository: r,
	}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Seller, error) {
	sellers, err := s.sellerRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return sellers, nil
}

func (s *service) Get(ctx context.Context, id int) (*domain.Seller, error){
	seller, err := s.sellerRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if seller.ID == 0 {
		return nil, ErrNotFound
	}

	return &seller, nil
}


func (s *service) Save(ctx context.Context, cid int, companyName, address, telephone string) (*domain.Seller, error) {
	existingSeller := s.sellerRepository.Exists(ctx, cid)

	if existingSeller {
		return nil, ErrConflict
	}

	newSeller := domain.Seller{
		CID:         cid,
		CompanyName: companyName,
		Address:     address,
		Telephone:   telephone,
	}

	sellerId, err := s.sellerRepository.Save(ctx, newSeller)
	if err != nil {
		return nil, err
	}

	savedSeller, err := s.sellerRepository.Get(ctx, sellerId)
	if err != nil {
		return nil, err
	}

	return &savedSeller, nil
}

func (s *service) Update(ctx context.Context, id int, cid int, companyName, address, telephone string) (*domain.Seller, error) {
	seller, err := s.sellerRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if seller.ID == 0 {
		return nil, ErrNotFound
	}

	existingSeller := s.sellerRepository.Exists(ctx, cid)

	if existingSeller && cid != seller.CID {
		return nil, ErrConflict
	}

	updateSeller := domain.Seller{
		CID:         cid,
		CompanyName: companyName,
		Address:     address,
		Telephone:   telephone,
	}

	err1 := s.sellerRepository.Update(ctx, updateSeller)
	if err1 != nil {
		return nil, err1
	}

	return &updateSeller, nil
}

func (s *service) Delete(ctx context.Context, id int) error{
	seller, err := s.sellerRepository.Get(ctx, id)
	if err != nil {
		return err
	}

	if seller.ID == 0 {
		return ErrNotFound
	}

	err1 := s.sellerRepository.Delete(ctx, id)
	if err1 != nil {
		return err
	}

	return nil
}