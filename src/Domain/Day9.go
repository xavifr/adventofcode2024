package Domain

import (
	"slices"
	"sort"
)

type D9Disk struct {
	Sectors map[int]*D9File
	Files   []*D9File
}

type D9File struct {
	Id          int
	Size        int
	StartSector int
}

type D9FilesById []*D9File

func (x D9FilesById) Len() int           { return len(x) }
func (x D9FilesById) Less(i, j int) bool { return x[i].Id < x[j].Id }
func (x D9FilesById) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type D9FilesBySector []*D9File

func (x D9FilesBySector) Len() int           { return len(x) }
func (x D9FilesBySector) Less(i, j int) bool { return x[i].StartSector < x[j].StartSector }
func (x D9FilesBySector) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func (d *D9Disk) RearrangeSectorsV1() {
	usedSectors := make([]int, 0)
	for i := range d.Sectors {
		if d.Sectors != nil {
			usedSectors = append(usedSectors, i)
		}
	}

	slices.Sort(sort.IntSlice(usedSectors))

	for i := len(usedSectors) - 1; i >= 0; i-- {
		for j := 0; j < usedSectors[i]; j++ {
			if d.Sectors[j] == nil {
				d.Sectors[j], d.Sectors[usedSectors[i]] = d.Sectors[usedSectors[i]], d.Sectors[j]
				break
			}
		}
	}

}

func (d *D9Disk) RearrangeSectorsNoFragment() {
	filesById := make(D9FilesById, len(d.Files))
	copy(filesById, d.Files)

	filesBySector := make(D9FilesBySector, len(d.Files))
	copy(filesBySector, d.Files)

	sort.Sort(sort.Reverse(filesById))
	sort.Sort(filesBySector)

	for i := 0; i < filesById.Len(); i++ {
		for j := 0; j < filesBySector.Len()-1; j++ {
			if filesBySector[j].StartSector >= filesById[i].StartSector {
				continue
			}

			if (filesBySector[j+1].StartSector - (filesBySector[j].StartSector + filesBySector[j].Size)) >= filesById[i].Size {
				filesById[i].StartSector = filesBySector[j].StartSector + filesBySector[j].Size
				sort.Sort(filesBySector)
				break
			}
		}
	}
}

func (d *D9Disk) ChecksumV1() int {
	sum := 0
	for i := 0; i < len(d.Sectors); i++ {
		if d.Sectors[i] == nil {
			break
		}
		sum += i * d.Sectors[i].Id
	}

	return sum
}

func (d *D9Disk) ChecksumV2() int {
	sum := 0

	filesBySector := make(D9FilesBySector, len(d.Files))
	copy(filesBySector, d.Files)
	sort.Sort(filesBySector)

	for _, file := range filesBySector {
		for i := 0; i < file.Size; i++ {
			sum += (file.StartSector + i) * file.Id
		}

	}

	return sum
}
