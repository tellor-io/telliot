package pow

import "github.com/tellor-io/TellorMiner/db"

func deleteFromDB(DB db.DB, keys []string) error {
	for _, k := range keys {
		err := DB.Delete(k)
		if err != nil {
			return err
		}
	}
	return nil
}
func writeToDB(DB db.DB, keys []string, values [][]byte) error {
	for i, k := range keys {
		err := DB.Put(k, values[i])
		if err != nil {
			return err
		}
	}
	return nil
}
