package main

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

type Cluster struct {
	Node *Node // we only have a single node to keep the example simple
}
type Node struct {
	Volume *Volume
}
type Volume struct{}

func toStack(err error) string {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	st := errors.Cause(err).(stackTracer).StackTrace()
	return fmt.Sprintf("%v%+v", err, st[:len(st)-2]) // ignore frames outside of main
}

func main() {
	if err := Reconcile(&Cluster{Node: &Node{}}); err != nil {
		log.Printf("Failed to reconcile cluster: %v", toStack(err))
	}
}

func Reconcile(c *Cluster) error {
	if err := UpdateNode(c.Node); err != nil {
		return errors.Wrap(err, "failed to update node")
	}
	return nil

}

func UpdateNode(n *Node) error {
	if err := AttachVolume(n); err != nil {
		return errors.Wrap(err, "failed to attach volume")
	}
	return nil
}

func AttachVolume(n *Node) error {
	v, err := CreateVolume()
	if err != nil {
		return errors.Wrap(err, "failed to create volume")
	}
	n.Volume = v
	return nil
}

func CreateVolume() (*Volume, error) {
	// simulated error
	return nil, errors.Errorf("something terrible has happened 😮")
}
