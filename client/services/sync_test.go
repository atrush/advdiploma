package services

import (
	"advdiploma/client/model"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"testing"
)

var (
	testFile = "sync_test_case.json"
)

type syncCase struct {
	Local    []model.SecretMeta
	Remote   map[uuid.UUID]int
	SyncList []SyncTask
}

func TestSync_CalcSyncBatch(t *testing.T) {
	syncCase := readSyncCase()

	syncResult, err := SyncService{}.CalcSyncBatch(syncCase.Remote, syncCase.Local)

	require.NoError(t, err)
	for i, el := range syncCase.SyncList {
		assert.EqualValues(t, el.Ver, syncResult[i].Ver, "ver")
		assert.EqualValues(t, el.LocID, syncResult[i].LocID, "loc id")
		assert.EqualValues(t, el.SecretId, syncResult[i].SecretId, "secretID")
		assert.EqualValues(t, el.ActionID, syncResult[i].ActionID, "action id")
	}
}

func readSyncCase() *syncCase {
	file, err := ioutil.ReadFile(testFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	data := syncCase{}
	if err := json.Unmarshal([]byte(file), &data); err != nil {
		log.Fatal(err.Error())
	}

	return &data
}
func genBaseSyncCase(count int) {
	loc := make([]model.SecretMeta, count)
	rm := make(map[uuid.UUID]int, count)
	syncList := make([]SyncTask, count)

	for i := 0; i < count; i++ {
		loc[i] = getMockSecret(i + 1)
		rm[loc[i].SecretID] = loc[i].SecretVer

		syncList[i] = SyncTask{
			LocID:    loc[i].ID,
			SecretId: loc[i].SecretID,
			Ver:      loc[i].SecretVer,
			ActionID: SyncActions["UPLOAD"],
		}
	}

	js, err := json.MarshalIndent(syncCase{
		Local:    loc,
		Remote:   rm,
		SyncList: syncList,
	}, "", "   ")

	if err != nil {
		log.Fatal(err.Error())
	}

	f, err := os.Create(testFile)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if _, err := f.Write(js); err != nil {
		log.Fatal(err)
	}
}
func getMockSecret(id int) model.SecretMeta {
	return model.SecretMeta{
		ID:        int64(id),
		SecretVer: rand.Intn(20) + 1,
		SecretID:  uuid.New(),
		StatusID:  model.SecretStatuses["ACTUAL"],
	}
}
