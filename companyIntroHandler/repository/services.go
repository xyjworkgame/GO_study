package repository

import (
	example "companyIntroHandler/companyIntro"
	"companyIntroHandler/model"
	"context"
	"fmt"
	"log"
)

var Intro *CompanyIntroRepo

/*type GetI int
type GetAll int*/
func  GetI(ctx context.Context, args *example.RequestGetI, reply *example.ResponseGetI) error {

	all, err := Intro.SelectIAll()
	if err != nil {
		log.Fatal(err)
	}
	reply.CompanyIntro = all
	fmt.Printf("call: request - %v ,\n response - %v", args, reply)
	return nil
}
func  GetAll(ctx context.Context, args *example.RequestGetAll, reply *example.ResponseGetAll) error {
	var page model.Pagination
	page.PageNum = args.Page.GetPageNum()
	page.PageSize = args.Page.GetPageSize()
	page.PageTotal = args.Page.GetPageTotal()

	all,err := Intro.Select(&page, int(args.GetStatus()))
	if err != nil{
		return err
	}
	reply.Page = &example.Page{
		PageTotal: page.PageTotal,
		PageSize: page.PageSize,
		PageNum: page.PageNum,
		Total: page.Total,
	}
	for _,v := range all{
		intro := example.CompanyIntro{}
		intro.Status = v.Status
		intro.Id = v.Id
		intro.SeqNo = v.SeqNo
		intro.Title  =v.Title
		intro.Content = v.Content
		reply.CompanyIntro = append(reply.CompanyIntro, &intro)
	}
	return nil
}
