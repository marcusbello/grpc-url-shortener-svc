package main

import (
	"context"
	"grpc-url-shortener-svc/db"
	pb "grpc-url-shortener-svc/proto"
	"reflect"
	"testing"
)

func Test_server_Add(t *testing.T) {
	type fields struct {
		UnimplementedShortenerSvcServer pb.UnimplementedShortenerSvcServer
		db                              *db.PGData
	}
	type args struct {
		ctx context.Context
		in  *pb.AddRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.AddResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedShortenerSvcServer: tt.fields.UnimplementedShortenerSvcServer,
				db:                              tt.fields.db,
			}
			got, err := s.Add(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_GetUrl(t *testing.T) {
	type fields struct {
		UnimplementedShortenerSvcServer pb.UnimplementedShortenerSvcServer
		db                              *db.PGData
	}
	type args struct {
		ctx context.Context
		in  *pb.GetUrlRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetUrlResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedShortenerSvcServer: tt.fields.UnimplementedShortenerSvcServer,
				db:                              tt.fields.db,
			}
			got, err := s.GetUrl(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_List(t *testing.T) {
	type fields struct {
		UnimplementedShortenerSvcServer pb.UnimplementedShortenerSvcServer
		db                              *db.PGData
	}
	type args struct {
		ctx context.Context
		in  *pb.ListRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ListResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedShortenerSvcServer: tt.fields.UnimplementedShortenerSvcServer,
				db:                              tt.fields.db,
			}
			got, err := s.List(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_Remove(t *testing.T) {
	type fields struct {
		UnimplementedShortenerSvcServer pb.UnimplementedShortenerSvcServer
		db                              *db.PGData
	}
	type args struct {
		ctx context.Context
		in  *pb.RemoveRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.RemoveResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedShortenerSvcServer: tt.fields.UnimplementedShortenerSvcServer,
				db:                              tt.fields.db,
			}
			got, err := s.Remove(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_Show(t *testing.T) {
	type fields struct {
		UnimplementedShortenerSvcServer pb.UnimplementedShortenerSvcServer
		db                              *db.PGData
	}
	type args struct {
		ctx context.Context
		in  *pb.ShowRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ShowResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedShortenerSvcServer: tt.fields.UnimplementedShortenerSvcServer,
				db:                              tt.fields.db,
			}
			got, err := s.Show(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Show() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Show() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_Update(t *testing.T) {
	type fields struct {
		UnimplementedShortenerSvcServer pb.UnimplementedShortenerSvcServer
		db                              *db.PGData
	}
	type args struct {
		ctx context.Context
		in  *pb.UpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.UpdateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedShortenerSvcServer: tt.fields.UnimplementedShortenerSvcServer,
				db:                              tt.fields.db,
			}
			got, err := s.Update(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
