package services

import (
	"advdiploma/client/model"
	"fmt"
	"log"
)

func (s *SyncService) Upload(task SyncTask) error {
	secret, err := s.db.GetSecret(task.LocID)
	if err != nil {
		return err
	}

	id, ver, err := s.provider.UploadSecret(secret.SecretData, task.SecretId, task.Ver)
	if err != nil {
		return err
	}

	secret.SecretVer = ver
	secret.SecretID = id

	if err := s.db.UpdateSecret(secret); err != nil {
		return err
	}

	return nil
}

func (s *SyncService) Download(task SyncTask) error {

	id, ver, data, err := s.provider.DownloadSecret(task.SecretId)
	if err != nil {
		return err
	}

	info := model.Info{}
	if err := info.FromEncodedData(data, s.cfg.MasterKey); err != nil {
		return fmt.Errorf("error read info drom encoded secret data: %w", err)
	}

	if task.LocID == 0 {
		_, err := s.db.AddSecret(model.Secret{
			Info:       info,
			SecretID:   id,
			SecretVer:  ver,
			SecretData: data,
			StatusID:   model.SecretStatuses["ACTUAL"],
		})
		if err != nil {
			return fmt.Errorf("error save secret data to storage: %w", err)
		}

		return nil
	}

	dbSecret, err := s.db.GetSecretByExtID(task.SecretId)
	if err != nil {
		return fmt.Errorf("error save secret data to storage: %w", err)
	}

	dbSecret.SecretData = data
	dbSecret.SecretVer = ver
	dbSecret.Info = info
	dbSecret.StatusID = model.SecretStatuses["ACTUAL"]

	if err := s.db.UpdateSecret(dbSecret); err != nil {
		return err
	}

	return nil
}

func (s *SyncService) Delete(task SyncTask) error {
	secret, err := s.db.GetSecret(task.LocID)
	if err != nil {
		return err
	}

	if err := s.provider.DeleteSecret(task.SecretId); err != nil {
		return err
	}

	secret.StatusID = model.SecretStatuses["DELETED"]
	err = s.db.UpdateSecret(secret)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *SyncService) ProcessTask(task SyncTask) {
	switch task.ActionID {

	case SyncActions["UPLOAD"]:
		if err := s.Upload(task); err != nil {
			log.Println(err.Error())
		}
	case SyncActions["DOWNLOAD"]:
		if err := s.Download(task); err != nil {
			log.Println(err.Error())
		}

	case SyncActions["SEND_DELETE"]:
		if err := s.Delete(task); err != nil {
			log.Println(err.Error())
		}
	}
}
