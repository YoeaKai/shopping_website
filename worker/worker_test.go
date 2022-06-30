package worker

import (
	"context"
	"shopping_website/model"
	pb "shopping_website/product"
	"testing"
)

func Test_Queue_Mouse(t *testing.T) {
	keyWord := "mouse"
	cleanupCtx := context.Background()
	pProduct := make(chan pb.ProductResponse, 1000)
	results := []pb.ProductResponse{}
	var workerConfig = model.WorkerConfig{
		MaxProduct: 200,
		WorkerNum:  2,
		SleepTime:  2,
	}

	go func() {
		for product := range pProduct {
			results = append(results, product)
		}
	}()

	Queue(cleanupCtx, keyWord, pProduct, workerConfig)

	if len(results) == 0 {
		t.Error("error in Queue, result: ", results)
	}

}
