package service

import (
	"context"
	"fmt"
	demo "rpc/grpc/pb"
)

type ProductServiceServer struct{}

func (p *ProductServiceServer) Gets(ctx context.Context, req *demo.GetsProductReq, res *demo.GetsProductRes) error {
	return nil
}
func (p *ProductServiceServer) GetsDo(ctx context.Context, req *demo.GetsProductDoReq, res *demo.GetsProductDoRes) error {
	return nil
}

func (p *ProductServiceServer) GetSearchDo(ctx context.Context, req *demo.GetProductSearchDoReq, res *demo.GetProductSearchDoRes) error {
	return nil
}

func (p *ProductServiceServer) PutStatusDo(ctx context.Context, req *demo.PutProductStatusDoReq, res *demo.PutProductStatusDoRes) error {
	return nil
}

func (p *ProductServiceServer) PostProductDo(ctx context.Context, req *demo.PostProductDoReq, res *demo.PostProductDoRes) error {
	return nil
}

func (p *ProductServiceServer) GeById(ctx context.Context, req *demo.GetProductByIdReq, res *demo.GetProductByIdRes) error {
	fmt.Println("成功")
	fmt.Println(req)
	res.Status = 1

	return nil
}
