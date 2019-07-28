package persist

import (
	"go_crawler/crawler/engine"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	ch := make(chan engine.Item, 1024)
	client := "ElasticClient"

	go func() {
		itemCount := 0
		for item := range ch {
			itemCount++
			log.Printf("Item Saver: Got Item #%d: %v", itemCount, item)
			err := save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: save error: %s", err)
			}
		}
	}()
	return ch, nil
}

// 返回存储的ID
func save(client, index string, item engine.Item) error {
	log.Printf("[client:%s] [index:%s] [item:%v] ", client, index, item)
	return nil
}
