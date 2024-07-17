package main

import (
	"testing"

	"github.com/hashicorp/consul/api"
)

func TestMyConsulClient_ForceLeave(t *testing.T) {
	type fields struct {
		client *api.Client
	}
	type args struct {
		node string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MyConsulClient{
				client: tt.fields.client,
			}
			if err := m.ForceLeave(tt.args.node); (err != nil) != tt.wantErr {
				t.Errorf("MyConsulClient.ForceLeave() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMyConsulClient_PruneNode(t *testing.T) {
	type fields struct {
		client *api.Client
	}
	type args struct {
		node string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MyConsulClient{
				client: tt.fields.client,
			}
			if err := m.PruneNode(tt.args.node); (err != nil) != tt.wantErr {
				t.Errorf("MyConsulClient.PruneNode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_forceLeave(t *testing.T) {
	type args struct {
		client   ConsulClient
		nodeName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := forceLeave(tt.args.client, tt.args.nodeName); (err != nil) != tt.wantErr {
				t.Errorf("forceLeave() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pruneNode(t *testing.T) {
	type args struct {
		client   ConsulClient
		nodeName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := pruneNode(tt.args.client, tt.args.nodeName); (err != nil) != tt.wantErr {
				t.Errorf("pruneNode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
