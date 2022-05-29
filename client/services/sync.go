package services

import (
	"advdiploma/client/model"
	"advdiploma/client/pkg"
	"advdiploma/client/provider"
	"advdiploma/client/storage"
	"github.com/google/uuid"
	"log"
)

type SyncService struct {
	db       storage.Storage
	provider provider.SecretProvider
	cfg      *pkg.Config
}

var (
	SyncActions = map[string]int{
		"UPLOAD":      1,
		"DOWNLOAD":    2,
		"SEND_DELETE": 3,
		"DELETE":      4,
	}
)

type SyncTask struct {
	LocID    int64
	SecretId uuid.UUID
	Ver      int
	ActionID int
}

func NewSyncService(db storage.Storage, provider provider.SecretProvider, cfg *pkg.Config) *SyncService {
	return &SyncService{
		db:       db,
		provider: provider,
		cfg:      cfg,
	}
}

//
//func (s *SyncService) Run(ctx context.Context) error {
//	ticker := time.NewTicker(s.syncTimeout)
//
//	go func() {
//		defer func() {
//			log.Println("defer")
//			ticker.Stop()
//		}()
//		for {
//			select {
//			case <-ticker.C:
//				//batch, err := s.CalcSyncBatch()
//
//				if err != nil {
//					log.Printf("error synchronization, err:%s", err.Error())
//					break
//				}
//
//				if len(batch) == 0 {
//					break
//				}
//
//				//s.tick(ctx, batch, s.processAccrualForOrder)
//			case <-ctx.Done():
//				log.Println("accrual worker context done")
//				return
//			}
//		}
//	}()
//	return nil
//}

func (s *SyncService) Tick() {

}

func (s *SyncService) GetSyncBatch() ([]SyncTask, error) {
	tasks := []SyncTask{}

	// get loc secrets array
	locList, err := s.db.GetMetaList()
	if err != nil {
		return nil, err
	}

	// get remote list
	remList, err := s.provider.GetSyncList()
	if err != nil {
		return nil, err
	}

	// map for local server ids
	locListMap := make(map[uuid.UUID]struct{})

	for _, el := range locList {

		// if locally exist new secrets, without secret id
		// add upload tasks
		if el.SecretID == uuid.Nil {

			// send upload
			tasks = append(tasks, SyncTask{
				LocID:    el.ID,
				SecretId: el.SecretID,
				Ver:      el.SecretVer,
				ActionID: SyncActions["UPLOAD"]})
			continue
		}

		// add secret id to map
		locListMap[el.SecretID] = struct{}{}

		//  get remote version
		remVer, ok := remList[el.SecretID]
		//  If not exist remote, delete local
		if !ok {
			// delete local
			// current el - secret without data field
			el.StatusID = model.SecretStatuses["DELETED"]

			err := s.db.UpdateSecret(el)
			if err != nil {
				log.Println(err.Error())
			}
			continue
		}

		// if deleted locally and in remote list
		// send delete
		if el.StatusID == model.SecretStatuses["DELETED"] {
			// send delete
			tasks = append(tasks, SyncTask{
				LocID:    el.ID,
				SecretId: el.SecretID,
				Ver:      el.SecretVer,
				ActionID: SyncActions["SEND_DELETE"]})
			continue
		}

		//  if locally has no changes and remote have newer version
		//  download external
		if remVer > el.SecretVer && el.StatusID == model.SecretStatuses["ACTUAL"] {
			// download

			tasks = append(tasks, SyncTask{
				LocID:    el.ID,
				SecretId: el.SecretID,
				Ver:      remVer,
				ActionID: SyncActions["DOWNLOAD"]})
			continue
		}

		//  if locally has changes and remote have newer version
		//  add collision
		if remVer > el.SecretVer && el.StatusID == model.SecretStatuses["EDITED"] {
			// ask collision

			tasks = append(tasks, SyncTask{
				LocID:    el.ID,
				SecretId: el.SecretID,
				Ver:      remVer,
				ActionID: SyncActions["DOWNLOAD"]})
			continue
		}

		//  if locally has changes and version is actual
		//  upload
		if remVer == el.SecretVer && el.StatusID == model.SecretStatuses["EDITED"] {

			tasks = append(tasks, SyncTask{
				LocID:    el.ID,
				SecretId: el.SecretID,
				Ver:      el.SecretVer,
				ActionID: SyncActions["UPLOAD"]})
		}

	}

	// if exist in remote list and not exist in local
	// create empty and download
	for k, v := range remList {
		_, ok := locListMap[k]
		if !ok {
			tasks = append(tasks, SyncTask{
				SecretId: k,
				Ver:      v,
				ActionID: SyncActions["DOWNLOAD"]})
			// make loc and download
		}
	}

	return tasks, nil
}

func SecretToUpload() {

}
