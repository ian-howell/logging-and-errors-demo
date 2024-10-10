package main

import (
	"log"
	"os"

	"github.com/pkg/errors"
)

// NOTE(ianhowell): global state for DEMO PURPOSES ONLY.
var numAttempts = 5

type Cluster struct {
	Node *Node // we only have a single node to keep the example simple
}
type Node struct {
	Volume *Volume
}
type Volume struct{}

func main() {
	log.SetOutput(os.Stdout)
	done := false
	for !done {
		done = true
		if err := Reconcile(&Cluster{Node: &Node{}}); err != nil {
			log.Printf("Failed to reconcile cluster: %v", err)
			done = false
		}
	}
}

func Reconcile(c *Cluster) error {
	if err := UpdateNode(c.Node); err != nil {
		return errors.Wrap(err, "failed to update node")
	}
	log.Println("Reconciled cluster")
	return nil
}

func UpdateNode(n *Node) error {
	if err := AttachVolume(n); err != nil {
		return errors.Wrap(err, "failed to attach volume")
	}
	log.Println("Updated node")
	return nil
}

func AttachVolume(n *Node) error {
	v, err := CreateVolume()
	if err != nil {
		return errors.Wrap(err, "failed to create volume")
	}
	n.Volume = v
	log.Println("Attaching volume")
	return nil
}

func CreateVolume() (*Volume, error) {
	if numAttempts > 0 {
		numAttempts--
		// simulate an error
		return nil, errors.Errorf("something terrible has happened 😮")
	}
	log.Println("Created volume")
	return &Volume{}, nil
}
