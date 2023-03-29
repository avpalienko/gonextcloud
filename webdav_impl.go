package gonextcloud

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"sort"

	"github.com/avpalienko/gowebdav"
)

// ShareType is the nextcloud shares types enum :
type AclType string

const (
	UserAcl  AclType = "user"
	GroupAcl AclType = "group"
)

type Acl struct {
	Type        AclType
	Id          string
	DisplayName string
	Mask        int
	Permissions SharePermission
}

type webDav struct {
	*gowebdav.Client
}

func newWebDav(url string, user string, password string) *webDav {
	wb := gowebdav.NewClient(url, user, password)
	return &webDav{Client: wb}
}

// Implementation adapted from filepath.Walk

// Walk walks the file tree rooted at root, calling walkFn for each file or
// directory in the tree, including root. All errors that arise visiting files
// and directories are filtered by walkFn. The files are walked in lexical
// order, which makes the output deterministic but means that for very
// large directories Walk can be inefficient.
// Walk does not follow symbolic links.
func (wd *webDav) Walk(root string, walkFn filepath.WalkFunc) error {
	info, err := wd.Stat(root)
	if err != nil {
		err = walkFn(root, nil, err)
	} else {
		err = wd.walk(root, info, walkFn)
	}
	if err == filepath.SkipDir {
		return nil
	}
	return err
}

// walk recursively descends path, calling walkFn.
func (wd *webDav) walk(path string, info os.FileInfo, walkFn filepath.WalkFunc) error {
	if !info.IsDir() {
		return walkFn(path, info, nil)
	}
	fis, err := wd.readDir(path)
	err1 := walkFn(path, info, err)
	// If err != nil, walk can't walk into this directory.
	// err1 != nil means walkFn want walk to skip this directory or stop walking.
	// Therefore, if one of err and err1 isn't nil, walk will return.
	if err != nil || err1 != nil {
		// The caller's behavior is controlled by the return value, which is decided
		// by walkFn. walkFn may ignore err and return nil.
		// If walkFn returns SkipDir, it will be handled by the caller.
		// So walk should return whatever walkFn returns.
		return err1
	}

	for _, fi := range fis {
		filename := filepath.Join(path, fi.Name())
		err = wd.walk(filename, fi, walkFn)
		if err != nil {
			if !fi.IsDir() || err != filepath.SkipDir {
				return err
			}
		}
	}
	return nil
}

// readDir reads the directory and returns
// a sorted list of directory entries.
func (wd *webDav) readDir(dirname string) ([]os.FileInfo, error) {
	fs, err := wd.ReadDir(dirname)
	if err != nil {
		return nil, err
	}
	sort.Slice(fs, func(i, j int) bool {
		return sort.StringsAreSorted([]string{fs[i].Name(), fs[j].Name()})
	})
	return fs, nil
}

func (wd *webDav) GetAclList(path string) ([]Acl, error) {
	var acList []Acl
	parse := func(resp interface{}) error {
		r := resp.(*response)
		for _, acl := range r.NcProp.AclLists {
			var a Acl
			a.Type = AclType(acl.MapType)
			a.Id = acl.Id
			a.Mask = acl.Mask
			a.Permissions = SharePermission(acl.Permissions)
			a.DisplayName = acl.DispName
			acList = append(acList, a)
		}
		return nil
	}
	err := wd.FindProps(path, "xmlns:nc='http://nextcloud.org/ns'", []string{"nc:acl-list", "nc:group-folder-id"}, &response{}, parse)
	if err != nil {
		return nil, err
	}
	return acList, err
}

func (wd *webDav) SetAclList(path string, acList []Acl) error {
	acls := aclList{Acl: make([]acl, len(acList))}
	for i, a := range acList {
		acls.Acl[i] = acl{
			MapType:     string(a.Type),
			Id:          a.Id,
			Mask:        a.Mask,
			Permissions: int(a.Permissions),
		}
	}
	err := wd.PatchProps(path, "xmlns:nc='http://nextcloud.org/ns'", acls)

	return err

}

type response struct {
	Href   string `xml:"DAV: href"`
	NcProp ncProp `xml:"DAV: propstat>prop"`
}

type ncProp struct {
	AclLists []acl `xml:"http://nextcloud.org/ns acl-list>acl"`
	GfId int `xml:"http://nextcloud.org/ns group-folder-id"`
}

type aclList struct {
	XMLName xml.Name `xml:"http://nextcloud.org/ns acl-list"`
	Acl []acl `xml:"acl"`
}

type acl struct {
	MapType     string `xml:"acl-mapping-type"`
	Id          string `xml:"acl-mapping-id"`
	DispName    string `xml:"acl-mapping-display-name",omitempty`
	Mask        int    `xml:"acl-mask"`
	Permissions int    `xml:"acl-permissions"`
}
