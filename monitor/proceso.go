package monitor

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gopkg.in/routeros.v2"
	"gopkg.in/routeros.v2/proto"
)

var (
	// Devs = []*Device{
	// 	{uuid.New(), "Dev Roberto", "172.16.201.1:8728", "api", "api"},
	// 	{uuid.New(), "Dev Jaime", "172.16.201.176:8728", "api", "api"},
	// }

	Devs = make(MapDevice)

	Pro = make(MapProcess)
)

func init() {
	id := uuid.New()
	Devs[id] = &Device{id, "Dev Roberto", "172.16.201.1:8728", "api", "api"}
	id = uuid.New()
	Devs[id] = &Device{id, "Dev Jaime", "172.16.201.176:8728", "api", "api"}

}

type chanReply struct {
	replicar bool
	reC      chan *proto.Sentence
}

type Proceso struct {
	Dev  Device `json:"dev"`
	quit chan int
	n    chan int
	info chan *proto.Sentence
	fin  bool
	chanReply
}

func Monitorizar(d *Device) {
	pro := NewProceso(*d)
	go pro.ListenMikrotik(time.Minute * 1)
	go pro.Run()
	Pro[d.ID] = pro
}

func NewProceso(d Device) *Proceso {
	return &Proceso{
		Dev:       d,
		n:         make(chan int),
		quit:      make(chan int),
		info:      make(chan *proto.Sentence),
		fin:       false,
		chanReply: chanReply{replicar: false, reC: make(chan *proto.Sentence)},
	}
}

// Procesador de eventos
func (p *Proceso) Run() {
	for {
		select {
		case numero := <-p.n:
			fmt.Println(numero)
		case info := <-p.info:
			//fmt.Println("Name:", p.Dev.Name, "ccq:", info.Map["overall-tx-ccq"])
			//fmt.Println(info)
			if p.replicar {
				p.reC <- info
			}
		case <-p.quit:
			fmt.Println("Finaliza el proceso:", p.Dev.ID)
			return
		}
	}
}

func (p *Proceso) ListenMikrotik(t time.Duration) {

	client, err := routeros.Dial(p.Dev.Address, p.Dev.User, p.Dev.Password)
	if err != nil {
		log.Println(err.Error())
		go p.Cancel()
		return
	}
	defer client.Close()
	fmt.Println("Connect to mikrotik")

	go p.CancelWithTime(t)

	listen, err := client.Listen("/interface/wireless/monitor", "=numbers=wlan1")
	if err != nil {
		log.Println(err)
	}

	// Escuchamos los datos de mikrotik
	for v := range listen.Chan() {
		if !p.fin {
			p.info <- v
		} else {
			listen.Cancel()
			return
		}
	}
	fmt.Println("Fin de listen")
}

func (p *Proceso) Cancel() {
	if !p.fin {
		p.quit <- 1
		p.fin = true
		delete(Pro, p.Dev.ID)
	}
}

func (p *Proceso) CancelWithTime(t time.Duration) {
	time.Sleep(t)
	p.Cancel()
}

func (p *Proceso) EnableReply() {
	p.replicar = true
	p.reC = make(chan *proto.Sentence)
}

func (p *Proceso) DisableReply() {
	p.replicar = false
	close(p.reC)
}

func (p *Proceso) Chan() <-chan *proto.Sentence {
	return p.reC
}
