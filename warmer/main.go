package main

import (
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

func main() {
	if err := Reconcile(&Cluster{}); err != nil {
		log.Printf("Failed to reconcile cluster: %v", err)
	}
}

func Reconcile(c *Cluster) error {
	if err := UpdateNode(c.Node); err != nil {
		log.Printf("Failed to update node: %v", err)
		return errors.Wrap(err, "failed to update node")
	}
	return nil

}

func UpdateNode(n *Node) error {
	if err := AttachVolume(n); err != nil {
		log.Printf("Failed to attach volume: %v", err)
		return errors.Wrap(err, "failed to attach volume")
	}
	return nil
}

func AttachVolume(n *Node) error {
	v, err := CreateVolume()
	if err != nil {
		log.Printf("Failed to create volume: %v", err)
		return errors.Wrap(err, "failed to create volume")
	}
	n.Volume = v
	return nil
}

func CreateVolume() (*Volume, error) {
	// simulated error
	return nil, errors.Errorf("something terrible has happened 😮")
}
