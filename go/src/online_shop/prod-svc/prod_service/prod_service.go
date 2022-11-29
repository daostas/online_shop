package prod_service

import (
	"context"
	"fmt"
	"online_shop/prod-svc/config"
	"online_shop/prod-svc/pb"
	rep "online_shop/repository"
	"online_shop/repository/models"
	"strconv"
	"strings"

	"gopkg.in/reform.v1"
)

type ProdServiceServer struct {
	pb.UnimplementedProdServiceServer
	Db  *reform.DB
	Cfg *config.Config
}

func NewProdServiceServer(db *reform.DB, cfg *config.Config) *ProdServiceServer {
	return &ProdServiceServer{
		Db:  db,
		Cfg: cfg,
	}
}

func (s *ProdServiceServer) RegisterProducer(ctx context.Context, req *pb.RegProducerReq) (*pb.ProdRes, error) {
	tr, err := s.Db.Begin()
	if err != nil {
		return &pb.ProdRes{Err: "error in begining of the transaction"}, nil
	}

	producer := rep.NewProducer(req.Photos, req.Status)

	err = tr.Insert(producer)
	if err != nil {
		tr.Rollback()
		return &pb.ProdRes{Err: "error in inserting data into producers table"}, nil
	}

	var warn bool = false

	for key, value := range req.Localizations {
		num, err := strconv.Atoi(key)
		if err != nil {
			return &pb.ProdRes{Err: "invalid data"}, nil
		}

		_, err = s.Db.SelectOneFrom(models.ProducersLocalizationView, "where title = $1", value.Title)
		if err == nil {
			warn = true
		}

		loc := rep.NewLocalizaionForProducer(producer.ProducerID, int32(num), value.Title, value.Description)
		err = tr.Insert(loc)
		if err != nil {
			tr.Rollback()
			return &pb.ProdRes{Err: "error in inserting data into producers localization table"}, nil
		}
	}

	tr.Commit()
	if warn {
		return &pb.ProdRes{Err: "success, but producer with this name already exist"}, nil

	}
	return &pb.ProdRes{Err: "success"}, nil
}

func (s *ProdServiceServer) GetListOfProducers(ctx context.Context, req *pb.EmptyProdReq) (*pb.GetListOfProducersRes, error) {
	dl, err := s.Db.SelectOneFrom(models.SettingsTable, "where key = 'DefaultLanguage'")
	if err != nil {
		return &pb.GetListOfProducersRes{
			Producerslist: nil,
			Err:           "error in getting data from settings table"}, nil
	}

	num, err := strconv.Atoi(dl.(*models.Settings).Value)
	if err != nil {
		return &pb.GetListOfProducersRes{
			Producerslist: nil,
			Err:           "invalid data in setting table"}, nil
	}

	query := fmt.Sprintf(`select p.producer_id, pl.title, p.photos
						from producers p, producers_localization pl
						where p.producer_id = pl.producer_id and pl.lang_id = %d;`, num)

	rows, err := s.Db.Query(query)
	if err != nil {
		return &pb.GetListOfProducersRes{
			Producerslist: nil,
			Err:           "error in getting data from producers and producers localization table"}, nil
	}
	defer rows.Close()

	var producers []*pb.GetListOfProducersResResult

	for rows.Next() {
		var id int32
		var title string
		var photos []uint8

		if err = rows.Scan(&id, &title, &photos); err != nil {
			fmt.Println(err)
		}

		var str string
		for i := 1; i < len(photos)-1; i++ {
			str += string(photos[i])
		}
		ss := strings.Split(str, ",")

		t := &pb.GetListOfProducersResResult{
			ProducerId: id,
			Title:      title,
			Photos:     ss,
		}

		producers = append(producers, t)
	}
	if err = rows.Err(); err != nil {
		return &pb.GetListOfProducersRes{
			Producerslist: producers,
			Err:           "error in processing of the data"}, nil
	}

	return &pb.GetListOfProducersRes{
		Producerslist: producers,
		Err:           "success"}, nil
}

func (s *ProdServiceServer) ChangeProducerStatus(ctx context.Context, req *pb.ChangeStatusReq) (*pb.ProdRes, error) {
	producer, err := s.Db.SelectOneFrom(models.ProducersTable, "where producer_id = $1", req.Id)
	if err != nil {
		return &pb.ProdRes{Err: "error in getting data from producer table"}, nil
	}

	if producer.(*models.Producers).Status {
		producer.(*models.Producers).Status = false
	} else {
		producer.(*models.Producers).Status = true
	}

	err = s.Db.Save(producer.(*models.Producers))
	if err != nil {
		return &pb.ProdRes{Err: "error in saving changes in producers table"}, nil
	}
	return &pb.ProdRes{Err: "success"}, nil
}
