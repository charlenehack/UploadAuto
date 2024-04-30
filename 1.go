func getHash(str string) (strHash string) {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str + serect))
	cipherStr := md5Ctx.Sum(nil)
	strHash = hex.EncodeToString(cipherStr)
	return
}

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	//	"github.com/astaxie/beego/logs"
	ldap "gopkg.in/ldap.v3"
)

type ldapConfig struct {
	addr   string
	port   int
	baseDn string
	dnUser string
	dnPwd  string
}

var (
	ldapconfig *ldapConfig
)

func InitLdap(addr string, port string, baseDn string, dnUser string, dnPwd string) {
	p, err := strconv.Atoi(port)
	if err != nil {
		p = 389
		fmt.Println("load . ", port, " is not a lega numuber ,use default 389")
	}
	ldapconfig = &ldapConfig{addr, p, baseDn, dnUser, dnPwd}
	ldap.DefaultTimeout = 3 * time.Second
	fmt.Printf("ldapconfig:%v\n", ldapconfig)
}

// ExampleConn_Bind demonstrates how to bind a connection to an ldap user
// allowing access to restricted attributes that user has access to
func Test_Conn() error {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "10.11.12.250", 389))
	if err != nil {
		fmt.Println(err)
	}
	defer l.Close()

	return err
}

// ExampleConn_Compare demonstrates how to compare an attribute with a value
func Conn_Compare(uid, pwd string) error {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapconfig.addr, ldapconfig.port))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer l.Close()

	matched, err := l.Compare(fmt.Sprintf("cn=%s,%s", ldapconfig.dnUser, ldapconfig.baseDn), "uid", uid)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(matched)
	return nil
}

func Conn_Search(uid string) (entries []*ldap.Entry, err error) {
	entries = nil
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapconfig.addr, ldapconfig.port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	_, err = l.SimpleBind(&ldap.SimpleBindRequest{
		Username: fmt.Sprintf("cn=%v,%v", ldapconfig.dnUser, ldapconfig.baseDn),
		Password: ldapconfig.dnPwd,
	})

	searchRequest := ldap.NewSearchRequest(
		ldapconfig.baseDn, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", uid),                                                // The filter to apply
		[]string{"dn", "cn", "sn", "mail", "telephoneNumber", "displayName", "userPassword", "modifytimestamp", "title"}, // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	/**
	for _, entry := range sr.Entries {
		fmt.Printf("%s: cn:%v,sn:%v,mail:%v,phone:%v,name:%v,password:%v\n", entry.DN, entry.GetAttributeValue("cn"), entry.GetAttributeValue("sn"), entry.GetAttributeValue("mail"), entry.GetAttributeValue("telephoneNumber"), entry.GetAttributeValue("displayName"), entry.GetAttributeValue("userPassword"))
		fmt.Printf(" %v\n", entry.DN)

	}

	**/
	return sr.Entries, err
}

// ExampleConn_Bind demonstrates how to bind a connection to an ldap user
// allowing access to restricted attributes that user has access to
func Ldap_Login(uid, pwd string) (err error) {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapconfig.addr, ldapconfig.port))
	if err != nil {
		fmt.Println("dial ldap error ", err)
		return
	}
	defer l.Close()

	_, err = l.SimpleBind(&ldap.SimpleBindRequest{
		Username: fmt.Sprintf("cn=%v,%v", ldapconfig.dnUser, ldapconfig.baseDn),
		Password: ldapconfig.dnPwd,
	})

	userdn, err := findDNbyUid(l, uid)
	if err != nil {
		return
	}
	err = l.Bind(userdn, pwd)
	if err != nil {
		fmt.Println("Bind ldap error ", err)
	}
	return
}

/*
*

*
 */
func findDNbyUid(l *ldap.Conn, uid string) (dn string, err error) {
	dn = ""
	searchRequest := ldap.NewSearchRequest(
		ldapconfig.baseDn, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", uid), // The filter to apply
		[]string{"dn"}, // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	fmt.Println(sr)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(sr.Entries) != 1 {
		fmt.Println("ldap:User does not exist or too many entries returned")
		err = errors.New("ldap :User does not exist or too many entries returned")
		return
	}

	dn = sr.Entries[0].DN
	return
}

func Ldap_Add(uid, name, cname, pwd, phone, mail string) (err error) {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapconfig.addr, ldapconfig.port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	err = l.Bind(fmt.Sprintf("cn=%s,%s", ldapconfig.dnUser, ldapconfig.baseDn), ldapconfig.dnPwd)
	if err != nil {
		fmt.Println(err)
		return
	}
	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%s,ou=Users,%s", uid, ldapconfig.baseDn), nil)
	addReq.Attribute("cn", []string{uid})
	addReq.Attribute("sn", []string{uid})
	addReq.Attribute("uid", []string{uid})
	addReq.Attribute("telephoneNumber", []string{phone})
	addReq.Attribute("mail", []string{mail})
	addReq.Attribute("displayName", []string{cname})
	addReq.Attribute("userPassword", []string{pwd})
	addReq.Attribute("objectClass", []string{"top", "person", "organizationalPerson", "inetOrgPerson"})
	fmt.Println(addReq)
	err = l.Add(addReq)
	if err != nil {
		return
	}
	//init pwd
	//passwordModifyReq := ldap.NewPasswordModifyRequest(fmt.Sprintf("cn=%s,ou=Users,%s", uid, ldapconfig.baseDn), "", pwd)
	//_, err = l.PasswordModify(passwordModifyReq)
	err = Ldap_PasswordModify(uid, pwd)
	if err != nil {
		fmt.Println("init passwd fail user = [%s], pwd = [%s]", uid, pwd)
	}
	return
}

func Ldap_Update(uid, phone, mail, cname string) (err error) {

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapconfig.addr, ldapconfig.port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	err = l.Bind(fmt.Sprintf("cn=%s,%s", ldapconfig.dnUser, ldapconfig.baseDn), ldapconfig.dnPwd)
	if err != nil {
		fmt.Println(err)
		return
	}
	modReq := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,ou=Users,%s", uid, ldapconfig.baseDn), nil)
	modReq.Replace("telephoneNumber", []string{phone})
	modReq.Replace("mail", []string{mail})
	modReq.Replace("displayName", []string{cname})
	modReq.Replace("objectClass", []string{"top", "person", "organizationalPerson", "inetOrgPerson"})
	fmt.Println(modReq)
	err = l.Modify(modReq)
	if err != nil {
		fmt.Println("modified fail for user = [%s] , err =%v", uid, err)
		return
	}
	return
}

func Ldap_Delete(uids []string) (err error) {

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapconfig.addr, ldapconfig.port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	err = l.Bind(fmt.Sprintf("cn=%s,%s", ldapconfig.dnUser, ldapconfig.baseDn), ldapconfig.dnPwd)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, u := range uids {
		delReq := ldap.NewDelRequest(fmt.Sprintf("cn=%s,ou=Users,%s", u, ldapconfig.baseDn), nil)
		err = l.Del(delReq)
	}
	if err != nil {
		fmt.Println("delete fail for user = [%s] , err =%v", uids, err)
		return
	}
	return
}

// update passwd by admin without old pwd
func Ldap_PasswordModify(uid, newPwd string) (err error) {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapconfig.addr, ldapconfig.port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	err = l.Bind(fmt.Sprintf("cn=%s,%s", ldapconfig.dnUser, ldapconfig.baseDn), ldapconfig.dnPwd)
	if err != nil {
		fmt.Println(err)
		return
	}
	userdn, err := findDNbyUid(l, uid)
	if err != nil {
		return
	}
	passwordModifyReq := ldap.NewPasswordModifyRequest(userdn, "", newPwd)
	_, err = l.PasswordModify(passwordModifyReq)

	if err != nil {
		fmt.Println("Password could not be changed:", err.Error())
	}
	return
}

func Ldap_BatchDelete(ids []int64) (err error) {

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapconfig.addr, ldapconfig.port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	err = l.Bind(fmt.Sprintf("cn=%s,%s", ldapconfig.dnUser, ldapconfig.baseDn), ldapconfig.dnPwd)
	if err != nil {
		fmt.Println(err)
		return
	}
	//delReq := ldap.NewDelRequest()
	return
}

func Ldap_AllocatePerm(uid, opt, perm string) (err error) {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapconfig.addr, ldapconfig.port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	err = l.Bind(fmt.Sprintf("cn=%s,%s", ldapconfig.dnUser, ldapconfig.baseDn), ldapconfig.dnPwd)
	if err != nil {
		fmt.Println(err)
		return
	}
	entries, err := Conn_Search(uid)
	if err != nil {
		return
	}
	if len(entries) == 0 {
		return errors.New("未查找到用户")
	}

	perms := entries[0].GetAttributeValues("title")

	var is_modify bool = false
	switch opt {
	case "add":
		{
			perms = append(perms, perm)
		}
	case "del":
		{
			for i, p := range perms {
				if strings.EqualFold(p, perm) {
					perms = append(perms[:i], perms[i+1:]...)
				}
			}
		}
	default:
		{
			is_modify = true
		}
	}
	if !is_modify {
		modReq := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,ou=Users,%s", uid, ldapconfig.baseDn), nil)
		modReq.Replace("title", perms)
		err = l.Modify(modReq)

		if err != nil {
			fmt.Println("allocate perm fail：", err.Error())
		}
	}
	return
}

func Ldap_SearchUsersWithPaging(uid, user_type string, pageSize, pageIndex int) (entries []*ldap.Entry, total int, err error) {
	entries = nil
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapconfig.addr, ldapconfig.port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	_, err = l.SimpleBind(&ldap.SimpleBindRequest{
		Username: fmt.Sprintf("cn=%v,%v", ldapconfig.dnUser, ldapconfig.baseDn),
		Password: ldapconfig.dnPwd,
	})

	if len(uid) == 0 {
		uid = "*"
	}

	var filter string
	if len(user_type) > 0 {
		filter = fmt.Sprintf("(title=%s)", user_type)
	}
	var searchRequest *ldap.SearchRequest
	if uid == "*" {
		searchRequest = ldap.NewSearchRequest(
			ldapconfig.baseDn, // The base dn to search
			ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
			fmt.Sprintf("(&%s(uid=%s))", filter, uid), // The filter to apply
			nil,
			nil,
		)
	} else {
		searchRequest = ldap.NewSearchRequest(
			ldapconfig.baseDn, // The base dn to search
			ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
			fmt.Sprintf("(&%s(uid=*%s*))", filter, uid), // The filter to apply
			nil,
			nil,
		)
	}

	sr, err := l.Search(searchRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	startIndex := pageSize * (pageIndex - 1) //include
	endIndex := pageSize * pageIndex         //exclude

	if len(sr.Entries) == 0 {
		return make([]*ldap.Entry, 0), 0, nil
	}

	if len(sr.Entries)-1 < startIndex {
		return nil, 0, errors.New("起始索引超过结果长度")
	}

	if endIndex > len(sr.Entries) {
		endIndex = len(sr.Entries)
	}

	return sr.Entries[startIndex:endIndex], len(sr.Entries), err
}

func main() {
	InitLdap("10.11.12.180", "389", "dc=ldap,dc=yunzhijia,dc=com", "admin", "Kingdee@2024")
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "10.11.12.180", 389))
	if err != nil {
		fmt.Println("dial ldap error ", err)
		return
	}
	defer l.Close()
	dn, e := findDNbyUid(l, "qing_hq")
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(dn)

	Ldap_AllocatePerm("qing_hq", "add", "git") // 添加权限

}
