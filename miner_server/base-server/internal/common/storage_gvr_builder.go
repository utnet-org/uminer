package common

import (
	"encoding/json"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	SpiderPersistentVolumeFmt      = "spider-pv-%s"
	SpiderPersistentVolumeChaimFmt = "spider-pvc-%s"
)

type PersistentVolumeSourceExtender struct {
	v1.PersistentVolumeSource
	Capacity string `json:"capacity,omitempty"`
}

func BuildStorageSource(b []byte) (*PersistentVolumeSourceExtender, error) {
	var pcs *PersistentVolumeSourceExtender
	err := json.Unmarshal(b, &pcs)
	if err != nil {
		return nil, err
	}

	return pcs, nil
}

func GetStoragePersistentVolume(name string) string {
	return fmt.Sprintf(SpiderPersistentVolumeFmt, name)
}

func GetStoragePersistentVolumeChaim(name string) string {
	return fmt.Sprintf(SpiderPersistentVolumeChaimFmt, name)
}

func BuildStoragePersistentVolume(name, capacity string) *v1.PersistentVolume {
	nameWrapped := GetStoragePersistentVolume(name)
	return &v1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: nameWrapped,
		},
		Spec: v1.PersistentVolumeSpec{
			StorageClassName:              name,
			PersistentVolumeReclaimPolicy: v1.PersistentVolumeReclaimRetain,
			AccessModes:                   []v1.PersistentVolumeAccessMode{v1.ReadWriteMany},
			Capacity:                      map[v1.ResourceName]resource.Quantity{v1.ResourceStorage: resource.MustParse(capacity)},
		},
	}
}

func BuildStoragePersistentVolumeChaim(namespace, name string, capacity string) *v1.PersistentVolumeClaim {
	nameWrapped := GetStoragePersistentVolumeChaim(name)
	return &v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      nameWrapped,
			Namespace: namespace,
		},
		Spec: v1.PersistentVolumeClaimSpec{
			Resources: v1.ResourceRequirements{
				Limits:   map[v1.ResourceName]resource.Quantity{v1.ResourceStorage: resource.MustParse(capacity)},
				Requests: map[v1.ResourceName]resource.Quantity{v1.ResourceStorage: resource.MustParse(capacity)},
			},
			StorageClassName: &name,
			AccessModes:      []v1.PersistentVolumeAccessMode{v1.ReadWriteMany},
			VolumeName:       GetStoragePersistentVolume(name),
		},
	}
}
