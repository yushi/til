package main

import (
	"strings"

	"golang.org/x/text/encoding/japanese"

	"github.com/Sirupsen/logrus"
	"github.com/alexcesaro/quotedprintable"
	"github.com/k0kubun/pp"
)

func decode(in string) string {
	elems := strings.Split(in, "?")
	data, err := quotedprintable.DecodeString(elems[3])
	if err != nil {
		logrus.Fatalf("error: %s", err)
	}
	return string(data)
}

func asISO2022JP(in string) string {
	b := make([]byte, 1024)
	d := japanese.ISO2022JP.NewDecoder()
	n, _, err := d.Transform(b, []byte(in), true)
	if err != nil {
		logrus.Fatalf("transform: %s", err)
	}
	return string(b[:n])
}
func main() {
	logrus.Info("OK")
	okl := make([]string, 3)
	okl[0] = decode("=?iso-2022-jp?Q?=1B$B%a!<%k7oL>J8;z2=3D$1%F%9%H=1B(B_opas-test=0D__?=")
	okl[1] = decode("=?iso-2022-jp?Q?[=1B$B%A%1%C%HHV9f=1B(B_23504?=")
	okl[2] = decode("=?iso-2022-jp?Q?0961]_=1B$BH/9T$N$*CN$i$;=1B(B?=")
	pp.Println(okl)
	for _, v := range okl {
		logrus.Info(asISO2022JP(v))
	}

	logrus.Info("NG")
	ngl := make([]string, 2)
	ngl[0] = decode("=?iso-2022-jp?Q?%a!<%k7oL>J8;z2=3D$1%F%9%H=1B(B_opas-test=0D__?=")
	ngl[1] = decode("=?iso-2022-jp?Q?[=1B$B%A%1%C%HHV9f=1B(B_235040961]_=1B$BH/9T$N$*CN$i$;?=")
	pp.Println(ngl)
	for _, v := range ngl {
		logrus.Info(asISO2022JP(v))
	}
}
