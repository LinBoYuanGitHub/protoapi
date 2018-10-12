// Code generated by "esc -o generator/data/tpl/tpl.go -modtime 0 -pkg=tpl generator/template"; DO NOT EDIT.

package tpl

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/generator/template/echo_enum.gogo": {
		local:   "generator/template/echo_enum.gogo",
		size:    383,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/3SQMWvDMBCFZ9+veIQOFjTOnpKpaaFLUmjoEjKo9tWYxmchy0MQ99+L7MSZsj2O70kf
b7XCa1cxahb2NnCFnwuc70JnXfOC7R67/QFv249DQeRs+WdrRozF5xRVicLFjaedbVkVjQSispM+IKcs
xiW8lZpRvDd8rnqopuuNjvHpFjfpkW97HnhClmCpVMkQ/Q5SIi+T6Fw1+Aq+kTo36MeASJnYlnusN2it
O87oaQIiZQ987kJrLOa8eL4WJo9MiTLPYfCC8Z9jEjqRPvJLw+YmDZLUrs28kWBG1JDSfwAAAP//H21g
KH8BAAA=
`,
	},

	"/generator/template/echo_service.gogo": {
		local:   "generator/template/echo_service.gogo",
		size:    2577,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/7xWbW/bNhD+LP2Kq5d1pBPLirciSwIj6xK764AkxWLsw8I0o6WzTVQmDYrOnEnsbx9O
cgw7TfcWo18E+l6ee+7hHeF2G05NijBGjVY6TGF4DzNrnJEzdQxnl3BxOYDe2dtBFIYzmXyQY4SiiN7V
R+/DUE1nxjpgYdCwOMbFrBGGQWOs3GQ+jBIzbWdymDuZfGhjMjGNkIdhYnReZfSmUmWQO6v0GLrQeM8Y
Y9ey9efr1m83pRBpef3iKyF2vn75jRBNIXaFaAnR7gpxIsT7298LIUr/8aa8FmJRxPHr2LfodHbQ73s6
9A/jpal/dro0nfUfTP1e39/wXSZE9MWL8iYvGRNi0elwappOcSnEIj7kTfoVp/SRnJ9s+Hb5CWNUOd4n
vPh7+gzpk9AHybg/EmJxMKJeFp39imfnW3K8GtasX6X06wCf10PJhFincrhJpa4xel4NznnzP6pTa8r5
D5+M0rN63R7WJlKrFCIqb8uPz8Jsbo8e50JEfHet4W0Jt03RtijYs8U62aletTtp6U2zi/pV60L9HEbn
89ydmulMZcgqF6fodpve0Qs5Re9B5eAmCEo7tCOZICRGO6l0DjLLKhcZrMkytHno7me4nrzKKsKgKFpg
pR4jROfoJibNwXsyRwPlMvSe0SscnRrtcOH2oFkU0Vs9m7vB/Qy958DIcjl3K1NRqBFohKhnrbFkg0bD
+zp1ZaM41Kn3vOaAOqXCPgyfZjSa6wRuV03c/iR1mqFlub2DothZmjlUbJfOPuUUYWDRza0GgmAJrPfD
gaG1gMSKU2igNBx1QeMf7FGjITlHFApdSKIflU6Z0vy4srzoglZZBRBYzGeEcWqmU6OrhguKrk5H8HJ1
Ls4xz+UYjwiiVoZxEj94YJxEP19dXrDvOvEeECwPg8AvidzJrGctFVI6+lVmKpUOGT9+cPwTpYeUJa06
618VN3P35B0DXXIl5fJqqWRu76K1YUr2QGlColt+GmOls3ncxAatOF4Wq3kFD1PkK46bwR0KNnPHw4BG
bG3eaK1+wbHKHdqN9ZrnmIIzMFQ6BWvmjhapGsJPwhlCsxqqXjIxe1BP5Gog/27HVhJcob1TCdYivLu8
GtRCYPSmN2CN6v+Tm3jf2PvMCvDNLfo89pveCprK/A/sR9v6VwAAAP//A2OiBBEKAAA=
`,
	},

	"/generator/template/echo_struct.gogo": {
		local:   "generator/template/echo_struct.gogo",
		size:    734,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/7SRQW+bQBCFz95fMeVgmShZ90xFpShgCclxWtfNpaqSLQx0W7yQYZFqrfa/V7Muim2p
x1zQ4+3sex/Dcgl3XYXQoEFSFiv4cYCeOtupXn+A7AE2DzvIs2InhehV+Vs1CM7JT0fpvRDLJRt3rRqG
jdqzZQ89XngwWBpLC07MnLsBUqZBkCuNbTWA9+zKnbYtj7I89Kyefw2dSSLn5DElej5eR1PxJS9EPZoS
FnTRFsOjanWlLC5iuJp0TtQREyDRAEkK375fBYJw4Pz/0W5A1yCnnC2+jJowIMx0DSRP4NMUoohLZsgV
qyJfZ0/b/PPXYptn7HJ3Cqrv0VQLfruG+QlFkPwVyVnuNYRj3ksCc/SxmPmzXUyU+PIKuuporyxEuFe6
jSbcd/QnZ0PeK1v+/GJJm2Zx2hW/4hebx9t1kT3l97fF+q3pJ61raNGE9Bg+wvuAQ2hHMjA/+5kuPIcE
eNaH0H9zRrfCi78BAAD//ww5MNzeAgAA
`,
	},

	"/generator/template/go/enum.gogo": {
		local:   "generator/template/go/enum.gogo",
		size:    383,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/3SQMWvDMBCFZ9+veIQOFjTOnpKpaaFLUmjoEjKo9tWYxmchy0MQ99+L7MSZsj2O70kf
b7XCa1cxahb2NnCFnwuc70JnXfOC7R67/QFv249DQeRs+WdrRozF5xRVicLFjaedbVkVjQSispM+IKcs
xiW8lZpRvDd8rnqopuuNjvHpFjfpkW97HnhClmCpVMkQ/Q5SIi+T6Fw1+Aq+kTo36MeASJnYlnusN2it
O87oaQIiZQ987kJrLOa8eL4WJo9MiTLPYfCC8Z9jEjqRPvJLw+YmDZLUrs28kWBG1JDSfwAAAP//H21g
KH8BAAA=
`,
	},

	"/generator/template/go/service.gogo": {
		local:   "generator/template/go/service.gogo",
		size:    1439,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5xUzW7bPBA8i0+xnxF8kAKXDoKeHOTSxE1TIHHQGL0GtLSWiMikQFJuU4LvXiylqLZj
o0Fvwv7Mzg5nNZnAlS4QSlRohMMCli/QGO20aOS01BdwPYf7+QJm17cLzlgj8mdRInjPH7rPEJj3Jxam
l8BDYHLdaOMgZcmolK5qlzzX60ktltaJ/HmCeaVHLGNsMiGMe7HGEEBacBWCVA7NSuQIuVZOSGVB1HVM
UcDoukZjmXtpcLt56PIs8f4DGKFKBH6HrtKFhRAYS7pxC+lqDCHNgXjwK60c/nRjkApOvee3qmnd4qXB
EDJIdetidN66Iey9XIFC4DNjtKEYjEYhjGEpfz0hxWLLkKUOVAVVoDEQKzJiOVDZJfKWxbsZHJ6bdZKg
KkiHwNhhgVatyuFp0PTpi1BFjSa1ZgPen/ThrJOtT36mHs8Sg641CghiT9gM0j9bU2kiFflE4Y90b1FG
yVUU6RJy/kmqIpUqu4iR/y5ByToCJAZtQxhXer3WKi7sqTp+TeH/4dvfobWixClBdMqkWQgdRmSc86+P
8/v04/nZGAg2Y0kSeiIbUc+MoUFS8e+iloVwmGYXr4m/UXpt6Wl1Xe8arlt38I2hN5E2/dPSSGs2fNvX
ZGVColc+jDHorPeX2KF1djYe3JqEHjGOJY67xedUrFuXsYQstuU3uvJvWErr0Oxce2uxAKdhKVUBRreO
7jqa8E15inAaTTXLKz2GzpGDIY+e/LYEj2g2MsdOhIf546ITAvnNbJGO4q/MVSGMxkdOINu9ouPYN7MB
msb8A/betf4OAAD//98TXM2fBQAA
`,
	},

	"/generator/template/go/struct.gogo": {
		local:   "generator/template/go/struct.gogo",
		size:    737,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/7SRQW+bQBCFz95fMeVgmShZ90xFpShgCclxWtfNpaqSLQx0W7yQYZFqrfa/V7Muim2p
x1zQ4+3sex/Dcgl3XYXQoEFSFiv4cYCeOtupXidN9wGyB9g87CDPip0Uolflb9UgOCc/HaX3QiyXbNy1
ahg2as+WPfR44cFgaSwtODFz7gZImQZBrjS21QDesyt32rY8yvLQs3r+NXQmiZyTx5To+XgdTcWXvBD1
aEpY0EVbDI+q1ZWyuIjhatI5UUdMgEQDJCl8+34VCMKB8/9HuwFdg5xytvgyasKAMNM1kDyBT1OIIi6Z
IVesinydPW3zz1+LbZ6xy90pqL5HUy347RrmJxRB8lckZ7nXEI55LwnM0cdi5s92MVHiyyvoqqO9shDh
Xuk2mnDf0Z+cDXmvbPnziyVtmsVpV/yKX2web9dF9pTf3xbrt6aftK6hRRPSY/gI7wMOoR3JwPzsZ7rw
HBLgWR9C/80Z3Qov/gYAAP//lnsw6+ECAAA=
`,
	},

	"/generator/template/php_client.gophp": {
		local:   "generator/template/php_client.gophp",
		size:    3889,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5RWW2/bNhR+1684MwzYDhwl3R5WxHOGNPG2AklTdM6AoSkMWj6OuUmkIlJNHIH/fSAl
0bpQtqeHICbP5Tsfz+2XX+NN7J2dwXxDBVABBNY0RHhChgmRuILlFuKES05iCsPvmAjKmZ+mb6kf8Ois
vBppG79bJSIvIMvAn9MIQSl9eXMPn+7nMLv5OPc9j5EIRUwChCzzP5EI/9Q/lJp4XioQ/ub8jfPHz9r6
VUwn5vBue31LluJxxtLI/Jl43tnJCdyhEOQJBZycnHlZdgoJYU8IfnmulBeERAjIMklliGAcKmVk6Rqo
+EDfZklSngO+SmQrAaX7x/yeJ7PXAGNJOTOqGArM9a95tFf/mkcRZy4TbKUU0CgOMUImKyoFeC/zAAAq
Uf1GMVwJTaq+0PxjoCnvF0RqDksVY97LBdNlSANYpyzQ3oEyKockScgW+gmKmDOBo1zR/N3rVX90DUMq
BMqh1f/asxh630ajiqXSGl0DPoN/S5YYQu/26sPsdvFl9nl2NZ/d9CrG9deXGypOL61JmILBOxxNanJr
niAJNtCBA4iocNMEVQFGxf3yHwwk+DdEkvk2RmggylFFMUyB4csunUp5pZrYSo3TS8N3BUeX3HcS0hWR
6LZUZ+TrN5gatYkzJp2fzghcVhrZ07JlMql6qrwj/B1NreOxj6S4qVlS7U7KQ7pd9HfE54Dtduyw1mK0
fao8911HTe/A/99C/qGo5EY4rRqWm4S/mJf5XDT+x7zph7avDXsDqz/QE4VxCfhKhexVOO2OLD9zN7tG
vALlwvoaZpkz05RqtH3IMuOvUowNulyv2qiPjhd4qiFq2k1QpglrmZ/UIt9x0TQu+aJogG67+eVxj763
OBtylTyG6aWjaHbAxscUzQF74wNF0T4ZlQzqkZ5f6b1ArwitpcAcdm0EdnJrqUOjN+BMSKgmSZb5f5Ew
dYzgHbDSr91BmmN8I2V8HVJkcuLOscXCeE7SQA77SyLwIaEwhcG7H3/2z/1z/93F+/P354OOpN6ZL1qs
3Tj+sDfD2gs00qr8Btr1Ik3owDxiAWTclpM0Qp5KI/bTeV1g1H7Dkrh+oJemCPXSBBdT8PMlq3gE1WgU
dyg33B43Kas2Cf8ji1OZp7lu1s8tojaErUJMYLozsBsnY+gv6ZtBNS4wsmaXNB0Vo1hud3oj1+ahb+2k
8+9TaZE5x3+CojndOsX27RFlH0pQ1C+V3Wkt/jJYN/5lsRjbIChb4Sv492YUCOgZ3V5HPKV2GZT11ZbN
B4/VOAi7eBj3upfnTC3BTh2LUmXDzc3t329tWO0l5sCGWNcq2XD5dOsW5OxZ4QqGiiaU7bFSawnt2f7A
/mX8hUEODnLy5DbGC+j5+7Jy3/J4lOdPHFZEkiJ5ceXX1omiV7an7K7fnV4GJAyvYqqr53lsxpDueXdy
pVQv//3w5aP5v+wBtisVzCnvvwAAAP//u9r3ZDEPAAA=
`,
	},

	"/generator/template/spring_service.gojava": {
		local:   "generator/template/spring_service.gojava",
		size:    856,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/6yRTY7iMBCF9zlFKauwGHOAbBADQiwgaMgFKkkRLKDssSuNUOS7t8JfsuhuqRFZRX5+
z++rGo/hr6kIamJyKFRBcQHrjBi0OoVZBussh/lsmasoslgesCZoW7W5/YaQRpE+WeMEjKuVt05zvXN4
orNxB3WmQhWaK4XMRlC0YbUgWaG1muv0t9aN8S97/5G3hj1NTXV5wfy/IS83b2Sb4qhLwMKLw1KgPKL3
3VTWeKIQpugJ2ggAoG3/gEOuCdSKZG8qDyE8Fb0DJlBbch+6pPxiCeLFPI8fdyYD3iS+Tl32IcSjmzok
up7ce7WtyhqxjXSJIfTFurhkMmDppCX3NzWP7sW7z5E0jnt7onmUXtUegbjq2v6ItMm2A6Z+/e9AWpAk
b4V4LvXbJ794L71HPZJC9BkAAP//3o5aN1gDAAA=
`,
	},

	"/generator/template/spring_struct.gojava": {
		local:   "generator/template/spring_struct.gojava",
		size:    585,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5SPwWoCMRCG73mKOdqD8QGWQkFbqBT1sC8w7o7baDYJk1mphLx7ibqtUmhpTjOB+f7v
n81g7luCjhwxCrWwPUFgLx6DqWCxhtW6hufFa62VCtgcsCNISW8uY86VUqYPngUa3+sdRiH+6K3eY3OI
3ml0zguK8U4vo3dzJhTP1b+ONuwDsZy+s/Z4RD2IsfrNRKmUCsPWmgYaizEWv3kZVthTzpAUAEBKU2B0
HYF+MWTbmPP5P7A5ohDsjENbTpd4xPoUzpcJdIFA6TlCyLWQszrvTzedLrSLx73BpKzeCQ+NbJCxz/nh
D6vy5N1EfaMAj78KlfWqldIdE6Zj1y+7m5IdScHWRmzhTka18phkYPczNl9jxuysPgMAAP//ioLHmkkC
AAA=
`,
	},

	"/generator/template/ts/helper.gots": {
		local:   "generator/template/ts/helper.gots",
		size:    3462,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4xWX28bxRZ/309xrnVv13bcdW5ppcrpNk3TtAlK4+AkL0QRmnjH8TTjmWVm1olxLYGg
FBCFPhQe4AEkhFTUh/YJoVYVXyZp+jHQmdldbxKnZF/sPed3fufvnNl6tepVYb3LNHQYp8A07FBBFTE0
gu0B+LGSRpKY+RZGHaothSFMaOhSHlMFnUS0DZNCg+kSA3sy4RFsU0g0jYCJAiWJmaWoQaITwvkA1X5Q
JzELjPZBKnzb399fo6rP2hSFqecxSVtGNtI9xYyhAinWBzFdaysWG68KF8//eFV4+/z3Nz8+PHj159GT
X9589fjg5XdZzl4VnObw2weHj58d/PXpwcufDx++fvPDi8NH3x9+84ezOPrpi8OvHx09e/72xedHT57O
rS6NDb98cPDqt6NfPzt4/ffRk6fFKOvecMg6ENyhIlgkel72elIsKCXVaOSxXiyVgSH0SFzQoD2MoKNk
D8s0HAbznGi9Qnp0NGpu39O+NxxSEY1GHjYWqnCLdpigESwaE8M81q0jFSiqYyk0hS4REWdix4Nq3aP7
1ikVSQ+6xsQWPvQAAG4t3J7bWF6HEKZrVrDSbN2dW4YQLk2nkptLH3600Go1WxDC5Uw437x7t7kyll9K
5Usr6wutlbnlXHMFLbKoMfAbMVGkB8Ms1NE4aNOlQLEcuagYfjaLDrKYJljOoJUGrCrZY5peE7RP1fU0
Q73HTLsLOS7QhphEV1I1Pm2CFUsLE+T5NnIAPoqaRInMR6DoPdo2Y9aIGFJ5V+cnOyvW8bg/Tk3qc0Ep
CCcMzAnnM+cKNyWs5POE4JFnf0xXyT04xjrjjbzTzesTDgS0UW7C4Ibj1UAEUIGnOEq1bm+0icCtQeKY
igiMtNsiUXxyey1BuU94IyWpZH/SnqW5OeBGa2le9mIpqDBoVAm8cQ1iTtq0XP/f5en6DquBf8OfqH5v
zqkbk9WXLtd3auD/9wztvDOunaGetsZTk7VXbjrjzTPUt5x6y68UOgE3E8YjILDRWsZd7uqK9bHt0Vhh
e5ZEVGhaojg0rHwbRxBfcWWgQLudXACnRA6v6McJ1QbkNg5RADTYCaztIuVctlJtfjmgiYOC7IDB1VYE
FgemAfnkQCfhHDYUz3p90bq4s7COKe7SQb1PeEIhJkxpy0H3SS/mtIEvmBLahlDCs9Wo17lsE96V2jSu
Tl+dLiGIqB0IYShIjzagtEfFzh5lpRoI1t51AsOIKI0Qm8UVXoeJhLPIEuYkF5AjzCkmDnZ2z32QUDXY
UPza+vVyovIpr6VFb8D6yYFnHSj/x2mLeys9B4niM8VjjJsjJspoCGFza8ZzUixmGVW71N7Np+lQ2Sd8
vEfQbdP13N6c2MqgS3RzT6wqGVNlBuVdOqgUSfDB9RCmDjZ36WBrTJmGmLFbZBiCwN7fv2+HRXYgE/uJ
iNw955/04XKfSGyTrCGJHuvrdXTINBClyOBkEIGRa7bg5YpzvJkO8Byit05534XQ1nEK/M0tvxAFUK7p
ZPDJEtn29AmfXB0EBB2pFki7W+7jGB4nzfOJiKHHNDalMxO6zTg9nY/1CCGg3dJaMzOdOQXK3Tq6Y/o0
efSf9dH5dtgznb6/1lwJ3LyzzqDcP+F2dOzNDnYQJ7pbTi+K3Qq2IfRhKr86KpVxSSvZ2chGQ1PFCGef
0GjVLbkwJb0nmSj7F3DV5qfuJLiYA27QqRDwDAdMRHS/2Sn7s74r+MX/wyz4sz40AClh6pTfYlzjkzzx
vkVP51zcqQS/HIsIwHVVgGVLKcflWyoHpuv13Nu16DmEkl356ce+1R9zGUJpjQwspnSedVsv0gWZ6TvX
7KQNW4hxLCwG9i+fHLblUKqXXD/zfKegFKCsSIUX9j8BAAD//17jOSSGDQAA
`,
	},

	"/generator/template/ts/objs.gots": {
		local:   "generator/template/ts/objs.gots",
		size:    1210,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5xSz2sTTRi+z1/xkEOTLl+Te+GDr5+NVShatHopRSabd5Mxm9l1ZtYahgHBihasFGw9
9KAnoVBoexK0/jlN0j9DZtMkG1s89CWE2Xmf55nn/VELAhZgvS00IhEThEaLJCluqIlGD+VUJSbhqSjP
wMJEGi6kBo9jmDahyQ2HNioLTaYIDRKyhUxTE0LmgKmq0dCkXoqQNAuwcJtgAS5Pvw0+v7s4/z7c/zp4
v3fx8+PYKQswyvQ/vO3vHS+t3e/v7A6PTy/P3gz3jwY7r/u/Dob7R8PD7cHB2WD3ZHj+afBlu39yePFj
hwU1xmo1kMy6mlm7AMVli1Ct+ws4x+hVmiiTA2Bt9QHvknOwDAAK+LuC4mZOGCXGwH/9+SmPM3LunwmJ
ZNNDHbP26lirjVpqeinNGFnmhq/7y4IZIQ2piId0C0eLsBZGe8mpuHPWigj0AtVV3qAYpdWl/+urzx7V
1+pL6/XlknMbm7nXidy1IhacYyyXqa6QrN7j+k7S7SayrlSifIVBwOB/+C/lindhi3mE4498e8hfImk8
p9AwBLVx5VEmQyMSiS5PC3RfQ2Wi4EvMPcwYmF9ERRslZOt69nHWyHvs3PxVI6NEoRKTQYd6fqUn4mOA
DxFh+mq1zfXDLbmmkpSU6VU61JvH3NyUudGh3maR7kNvCRO2MULb4uC9xRUyI5eU25wdaTFCrglla5FP
Gc6VF69hfCgymZJ/eALXfikm2+AfuBrwDU81KeJZbP6qX3oiOzLZQu66NIN0jE2P0/+bicXd+h0AAP//
LkxTtboEAAA=
`,
	},

	"/generator/template/ts/service_axios.gots": {
		local:   "generator/template/ts/service_axios.gots",
		size:    1740,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4xTYW8bRRP+fr9iZFWyY7nn5pXeKjgEKYIC+QCJ1PAD1ndje8N597q7lzQ6VmoIaamU
NJEaVUACDYKiiCYxSFCStKF/5s52PvUvoL29ixOEEPvpdueZZ555Zq5erTpVmO9QCS0aIFAJbWQoiEIf
mstQDgVXnIS0nMHQojzOFKFMQktwppD5MD03Ax73EVSHKFji4lNYoqoDqoMQ0KYgYhnK5C7lslwzjwJb
XGANqCpLEHgnogJ9m5zDjBTKpCJBgD5QllGFgi+gp3ItI6VZaSphSVClkBn4/HKItz1BwwKdYULBF6mP
Egg0iaQeRJK0EVpc2BZIEABhPnTJMjBEH4i/EEnVRaaAeB4XPmVtUBxkiB5tUa+QVDRhkcwHSVVEFOXM
qcL1/36cKgx7z/pPHiQvXwy2n/a/3EpOHxVDcKpgI+n6Wrr1PH24MXjeG/6yOtjen56bGXzzRfLyh8He
yptX6+nJi+Ts9WB7f3BwkBw/7D85SU8fZ8a+ebUO6c5e//DH8917w59WktffDnsrWSg9/Crd3U9OH51/
fzLYOUqOD0cF769Zbst6xdxh75lVmkM3f+5vbiV/7lh503MzVmH69LS/ezBSePTd+ddr/dW19P7v6WZv
uHpm9fT3TvobR+naH8nZY6vDfpvQr58nxxvp1vrw3npytps+OO3v/NbfPnGqdYd2Qy4UZH3UIIZp8zEn
eJdKBG0WtVts1mQBjh0AgDgWhLURrqnlEGtwrcl5AI0pqLRRzWTA94gipmMJ7vsR88xQ5ZjWefZ1mwla
1/IXZL7WTlHUrcex+25ApPyYdFHr2ebCZQ0XW/yJCGqAQnDxIWF+YPZsRNHBIERRnnScerVJpAFDutmz
/VfrjseZVFBEpqDUUSps1Ovjb/3PHb854Y6P/9+9eaMxcWPiRmnSqdchkmjNcpysBa8QaFofyYXrWluE
NWlkgNY2sYuqw32TVQq5VKUsw0RoC/BO5uJtFIvUw4+UD+5saO2D0ge35kfgyzRtvGCxTmaAzBkTzyhv
mZuZySVGrR28m5naykVCHLvW9EpIBOnKhnmZYWGkTK7WYw3Id+TtOHZnI3URgc+A4SKKd/ItCVBBJIIG
SCXMaKYuz62SG1+DUhyPrNS6lL3kGkpjk07GJVBFgln73TjOe9e6EhmKOKYtYHhhifVDa9tBHGMgUesY
7B107tJYRm2OqzrIKgIlTBXqi1OvQ8csF4KMPA+lBJ8ocgWSi8ttcQVKHiyioXMNFoiEv1k1qqzHXI8o
r1NBIf6tuJ2lQBlyJvGfyl/5DwydW6AvV5t08iXJ1uWvAAAA///fSH7BzAYAAA==
`,
	},

	"/generator/template/ts/service_fetch.gots": {
		local:   "generator/template/ts/service_fetch.gots",
		size:    1539,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4xTXW8bRRR9319xZVXyh5x1gkQVnCZSxIdqHmpLMT9gvHttT1jPLDOzCdYyUkNwC1JC
IhEh0US0CBVVFNUgQSEpgT/jtZ0n/gLambWdVDzgp13fc889956zlVLJKUGzSyW0aYBAJXSQoSAKfWj1
IR8KrjgJad7A0KI8zhShTEJbcKaQ+bDZqIHHfQTVJQp2ufgQdqnqQhuV1zXVNhdwt9lsQCRJB2VGtxhm
uqmEXUGVQgaUQbMf4pYnaKgytMGEgu9QHyUQaBFJPUto+I0KEgRAmA890geG6APxtyOpesgUEM/jwqes
A4qDDNGjbeqljNvoKRD4UUQFWiTzQVIVEUU5c0qw9P9/Tgmmw6fjrx+OXr2cnDwef348uvhydkenBLaS
HAyS4+fJF4eT58Ppz/uTk2ebjdrk0WejV99Pnuz98+dBcv5ydPn35OSZuWGy9yi9nkUuSB4MLN4ibxxs
Onxqp2fQox/HR8ejv07tyM1GzXIljy/GZz8tpr749uqbwXh/kDz4LTkaTvcvr87uT3/YGz85Hx++SAa/
jy6/uvrufHKaPaelXz4d/XGYHB9M7x+MLs+Shxfj01/HJ+dOqeLQXsiFgtgBAIhjQVgH4Zbqh1iGWy3O
A6iuQ6GDqmaA7xBF0iUkuO9FzEtvL4taZ91LthO0Lmf/IPO1dnSawx7k3Uocu28HRMp7pIda11vbMr82
1zAP2wciKAMKwcVdwvwgjcOCootBiCK/5jiVUovIFAzJ0dCuVKo4HmdSwayyDrmuUmG1Ull56w135faq
u7Lypnt7ubq6vLqcW3Mco9qbaUq3XSiEJa0dp1KBSKL9Upx2trWJ8Z0aa5pL1SOVPmwUJIod6mEVpBKU
dcrQQ9Xl/uI9JIL0ZBVsZ7EKDcF7VOKdjAI+AYY7KDYySwJUEIlgRgDr149UyLYsQzZ3Nq+45phugSoS
zEovRCkwnivKN+pbzXwZWtzvV+H9rfo9186g7X7ByiyCLrqqi6wgUML6TNM15ky9K1DyYAdTnLstOSsU
iwaqi65H0uEoxH8S3HA5RaVUIWcSZwRrjrYu2Wwucqe1Nc9QpMaZnL6bvplLuvXQ5hO0dvBjk7G5fXHs
2gwWZo7EsVtjoXVB62vOxLFbj9S88rpD2SImDzc5yvBa60YhF8eLsGmdK0NuriQ3S4dZ2Xw7JoD/BgAA
//8qr8BMAwYAAA==
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/generator": {
		isDir: true,
		local: "generator",
	},

	"/generator/template": {
		isDir: true,
		local: "generator/template",
	},

	"/generator/template/go": {
		isDir: true,
		local: "generator/template/go",
	},

	"/generator/template/ts": {
		isDir: true,
		local: "generator/template/ts",
	},
}
