package game

import (
	"context"
	"runtime"
	"time"

	"game/models"

	"github.com/satori/go.uuid"
)

type Player struct {
	uid string
}

type Game struct {
	id    string
	teams [][]Player
}

type Dummy struct {
	uid string
	c   chan *Game
}

var (
	queue = models.NewQueue()
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				dummies := queue.PopMulti(2)
				if dummies != nil {
					g := &Game{
						id: uuid.NewV1().String(),
					}
					g.teams = [][]Player{
						{
							Player{uid: dummies[0].(Dummy).uid},
						},
						{
							Player{uid: dummies[1].(Dummy).uid},
						},
					}
					for _, d := range dummies {
						d.(Dummy).c <- g
						close(d.(Dummy).c)
					}
				} else {
					select {
					case <-time.After(2 * time.Second):
					}
				}
			}
		}()
	}
}

func QueueGame(ctx context.Context, uid string) (*Game, error) {
	c := make(chan *Game, 1)
	e := queue.Push(Dummy{uid, c})
	select {
	case <-ctx.Done():
		if queue.Remove(e) != nil {
			// timeout
			return nil, ctx.Err()
		}
		g := <-c
		return g, nil
	case g := <-c:
		return g, nil
	}
}
