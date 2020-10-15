package product_test

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BackAged/go-ddd-microservice/mock/repo"
	"github.com/BackAged/go-ddd-microservice/product"
)

var _ = Describe("Service", func() {
	var (
		mockCtrl *gomock.Controller
		mckRpo   *repo.ProductRepo
		srvc     product.Service
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mckRpo = repo.NewProductRepo(mockCtrl)
		srvc = product.NewService(mckRpo)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("NewService", func() {
		Context("With valid repo", func() {
			It("should get a service", func() {
				Expect(srvc).To(BeIdenticalTo(srvc))
			})
		})
	})

	Describe("CreateProduct", func() {
		Context("With valid product", func() {
			It("should get a product", func() {
				prd := product.Product{
					Image:           "image",
					DiscountedPrice: 2,
					Price:           2,
					Name:            "name",
					Slug:            "slug",
					Stock:           0,
					Status:          product.StatusActive,
					Version:         0,
				}
				mckRpo.EXPECT().Create(gomock.Any(), &prd).
					DoAndReturn(func(ctx context.Context, prod *product.Product) (*product.Product, error) {
						prd.ID = 2
						return &prd, nil
					}).AnyTimes()

				_, err := srvc.CreateProduct(context.Background(), &prd)
				Expect(err).ToNot(HaveOccurred())
			})
		})
		Context("With a invalid product", func() {
			It("should get a error", func() {
				prd := product.Product{
					Image:           "image",
					DiscountedPrice: -2,
					Price:           2,
					Name:            "name",
					Slug:            "slug",
					Stock:           0,
					Status:          product.StatusActive,
					Version:         0,
				}
				mckRpo.EXPECT().Create(gomock.Any(), &prd).
					DoAndReturn(func(ctx context.Context, prod *product.Product) (*product.Product, error) {
						prd.ID = 2
						return &prd, nil
					}).AnyTimes()

				_, err := srvc.CreateProduct(context.Background(), &prd)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("GetProduct", func() {
		Context("with id exist", func() {
			It("should get a product", func() {
				prd := product.Product{
					ProductID:       2,
					Image:           "image",
					DiscountedPrice: 2,
					Price:           2,
					Name:            "name",
					Slug:            "slug",
					Stock:           0,
					Status:          product.StatusActive,
					Version:         0,
				}
				mckRpo.EXPECT().GetByProductID(gomock.Any(), prd.ID).
					DoAndReturn(func(ctx context.Context, prod int64) (*product.Product, error) {
						prd.ID = 2
						return &prd, nil
					}).AnyTimes()

				prdt, err := srvc.GetProduct(context.Background(), prd.ID)
				Expect(err).ToNot(HaveOccurred())
				Expect(prdt.ID).To(Equal(int64(2)))
			})
		})
		Context("With id not exist", func() {
			It("should occur a error", func() {
				prd := product.Product{
					Image:           "image",
					DiscountedPrice: -2,
					Price:           2,
					Name:            "name",
					Slug:            "slug",
					Stock:           0,
					Status:          product.StatusActive,
					Version:         0,
				}
				mckRpo.EXPECT().GetByProductID(gomock.Any(), prd.ID).
					DoAndReturn(func(ctx context.Context, prod int64) (*product.Product, error) {
						return nil, nil
					}).AnyTimes()

				prdt, err := srvc.GetProduct(context.Background(), prd.ID)
				Expect(err).To(HaveOccurred())
				Expect(prdt).To(BeNil())
			})
		})
	})

})
