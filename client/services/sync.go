package services

import (
	"advdiploma/client/model"
	"advdiploma/client/pkg"
	"advdiploma/client/provider"
	"advdiploma/client/storage"
	"context"
	"github.com/google/uuid"
	"golang.org/x/time/rate"
	"log"
	"sync"
	"time"
)

type SyncService struct {
	db       storage.Storage
	provider provider.SecretProvider
	cfg      *pkg.Config
	limiter  *rate.Limiter
}

func NewSyncService(db storage.Storage, provider provider.SecretProvider, cfg *pkg.Config) *SyncService {
	return &SyncService{
		db:       db,
		provider: provider,
		cfg:      cfg,
		limiter:  rate.NewLimiter(rate.Limit(float64(cfg.RequestsPerMinute)/float64(60)), 1),
	}
}

func (s *SyncService) Run(ctx context.Context) error {
	ticker := time.NewTicker(s.cfg.SyncTimeout)

	go func() {
		defer func() {
			log.Println("defer")
			ticker.Stop()
		}()
		for {
			select {
			case <-ticker.C:
				if err := s.provider.PingAuth(); err != nil {
					log.Println(err.Error())
					break
				}

				batch, err := s.GetSyncBatch()
				if err != nil {
					log.Printf("error synchronization, err:%s", err.Error())
					break
				}

				if len(batch) == 0 {
					break
				}

				s.tick(ctx, batch)
			case <-ctx.Done():
				log.Println("accrual worker context done")
				return
			}
		}
	}()
	return nil
}

func (s *SyncService) tick(ctx context.Context, tasks []SyncTask) {
	wg := sync.WaitGroup{}
	wg.Add(len(tasks))

	for _, o := range tasks {
		if err := s.limiter.Wait(ctx); err != nil {
			log.Printf(err.Error())

			break
		}

		o := o
		go func() {
			defer wg.Done()

			if err := s.ProcessTask(o); err != nil {
				log.Println(err.Error())
			}
		}()
	}

	wg.Wait()
}

//  GetSyncBatch downloads remote list, gets local list, compares and returns list of tasks
func (s *SyncService) GetSyncBatch() ([]SyncTask, error) {

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

	return s.CalcSyncBatch(remList, locList)
}

//  GetSyncBatch compares local data and server meta info and returns list of tasks
func (s SyncService) CalcSyncBatch(rm map[uuid.UUID]int, loc []model.SecretMeta) ([]SyncTask, error) {
	tasks := []SyncTask{}

	// map for local server ids
	locListMap := make(map[uuid.UUID]struct{})

	for _, el := range loc {

		//  SecretID is nil - non sync local elements
		if el.SecretID == uuid.Nil {
			//  if status NEW - UPLOAD
			if el.StatusID == model.SecretStatuses["NEW"] {
				tasks = append(tasks, taskUploadNew(el.ID, el.SecretVer))
				continue
			}

			//  if status DELETED - delete locally
			if el.StatusID == model.SecretStatuses["DELETED"] {
				tasks = append(tasks, taskDeleteLocally(el.ID))
				continue
			}

			log.Printf("sync err, element id:%v with status %v has nil SecretID", el.ID, el.StatusID)
			continue
		}

		// secret id exist - add secret id to map
		locListMap[el.SecretID] = struct{}{}

		//  check remote version
		remVer, remExist := rm[el.SecretID]

		if !remExist {
			// if SecretID not nil and not exist remote - delete locally
			tasks = append(tasks, taskDeleteLocally(el.ID))
			continue
		}

		// if status DELETED and exist remote - send delete
		if el.StatusID == model.SecretStatuses["DELETED"] {
			// send delete
			tasks = append(tasks, taskDeleteRemote(el.SecretID))
			continue
		}

		//  if locally has no changes and remote have newer version - download
		if remVer > el.SecretVer && el.StatusID == model.SecretStatuses["ACTUAL"] {
			// download
			tasks = append(tasks, taskDownload(el.SecretID))
			continue
		}

		//  if locally has changes and remote have newer version
		//  add collision
		if remVer > el.SecretVer && el.StatusID == model.SecretStatuses["EDITED"] {
			// task collision
			tasks = append(tasks, taskCollision(el.ID, el.SecretID, remVer))
			continue
		}

		//  if locally has changes and version is actual
		//  upload
		if remVer == el.SecretVer && el.StatusID == model.SecretStatuses["EDITED"] {
			tasks = append(tasks, taskUpload(el.SecretID, el.SecretVer))
			continue
		}

	}

	// if exist in remote list and not exist in local - download
	for k, _ := range rm {
		_, ok := locListMap[k]
		if !ok {
			tasks = append(tasks, taskDownload(k))
		}
	}
	return tasks, nil
}

func taskUploadNew(locID int64, ver int) SyncTask {
	return SyncTask{
		LocID:    locID,
		Ver:      ver,
		ActionID: SyncActions["UPLOAD_NEW"]} //1
}
func taskUpload(secretID uuid.UUID, ver int) SyncTask {
	return SyncTask{
		SecretId: secretID,
		Ver:      ver,
		ActionID: SyncActions["UPLOAD"]} //2
}
func taskDownload(secretID uuid.UUID) SyncTask {
	return SyncTask{
		SecretId: secretID,
		ActionID: SyncActions["DOWNLOAD"]} //3
}
func taskDeleteLocally(locID int64) SyncTask {
	return SyncTask{
		LocID:    locID,
		ActionID: SyncActions["DELETE_LOCALLY"]} //4
}
func taskDeleteRemote(secretID uuid.UUID) SyncTask {
	return SyncTask{
		SecretId: secretID,
		ActionID: SyncActions["SEND_DELETE"]} //5
}
func taskCollision(locID int64, secretID uuid.UUID, rmVersion int) SyncTask {
	return SyncTask{
		LocID:    locID,
		SecretId: secretID,
		Ver:      rmVersion,
		ActionID: SyncActions["COLLISION"]} //6
}
