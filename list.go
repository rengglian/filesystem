package filesystem

import (
    "path/filepath"
    "sort"
)

type FileInfo struct {
	FullPath string
	Name string
}

type FileList struct {
	Items []FileInfo
}
   
func (newFile *FileList) AddItem(item FileInfo) []FileInfo {
    newFile.Items = append(newFile.Items, item)
    return newFile.Items
}

func ListFiles(root string, serachexp string, sortBy string) ([]string, error){
    items := []FileInfo{}
    fileCollection := FileList{items}
    
	files, err := filepath.Glob(root + serachexp)
    
	for _, file := range files {
			dir, fileName := filepath.Split(file)
			newFile := FileInfo{FullPath: dir + fileName, Name: fileName}
            fileCollection.AddItem(newFile)			
	}
    
	sort.Slice(fileCollection.Items, func(i, j int) bool { return fileCollection.Items[i].Name < fileCollection.Items[j].Name })	
    
    var filesSorted []string
    for _, sortedItem := range fileCollection.Items {
        filesSorted = append(filesSorted, sortedItem.FullPath)
    }

    return filesSorted, err
   
}