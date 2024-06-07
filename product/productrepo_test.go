package product

import (
	"reflect"
	"testing"
)

func TestNewProductRepo(t *testing.T) {
	tests := []struct {
		name string
		want *ProductRepo
	}{
		{
			name: "Default",
			want: &ProductRepo{products: []Product{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductRepo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductRepo_Create(t *testing.T) {
	type fields struct {
		products []Product
	}
	type args struct {
		partial Product
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Product
	}{
		{
			name: "Create new product",
			fields: fields{
				products: []Product{},
			},
			args: args{
				partial: Product{Name: "Test Product"},
			},
			want: Product{ID: 1, Name: "Test Product"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProductRepo{
				products: tt.fields.products,
			}
			got := p.Create(tt.args.partial)
			if got.ID == 0 {
				t.Errorf("expected non-zero product ID, got %d", got.ID)
			}
			if got.Name != tt.want.Name {
				t.Errorf("Create() = %v, want %v", got.Name, tt.want.Name)
			}
		})
	}
}

func TestProductRepo_DeleteOne(t *testing.T) {
	type fields struct {
		products []Product
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Delete existing product",
			fields: fields{
				products: []Product{{ID: 1, Name: "Test Product"}},
			},
			args: args{
				id: 1,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Delete non-existing product",
			fields: fields{
				products: []Product{},
			},
			args: args{
				id: 1,
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProductRepo{
				products: tt.fields.products,
			}
			got, err := p.DeleteOne(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeleteOne() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductRepo_GetList(t *testing.T) {
	type fields struct {
		products []Product
	}
	tests := []struct {
		name   string
		fields fields
		want   []Product
	}{
		{
			name: "Get list of products",
			fields: fields{
				products: []Product{{ID: 1, Name: "Product1"}, {ID: 2, Name: "Product2"}},
			},
			want: []Product{{ID: 1, Name: "Product1"}, {ID: 2, Name: "Product2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProductRepo{
				products: tt.fields.products,
			}
			if got := p.GetList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductRepo_GetOne(t *testing.T) {
	type fields struct {
		products []Product
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Product
		wantErr bool
	}{
		{
			name: "Get existing product",
			fields: fields{
				products: []Product{{ID: 1, Name: "Test Product"}},
			},
			args: args{
				id: 1,
			},
			want:    Product{ID: 1, Name: "Test Product"},
			wantErr: false,
		},
		{
			name: "Get non-existing product",
			fields: fields{
				products: []Product{},
			},
			args: args{
				id: 1,
			},
			want:    Product{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProductRepo{
				products: tt.fields.products,
			}
			got, err := p.GetOne(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOne() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductRepo_Update(t *testing.T) {
	type fields struct {
		products []Product
	}
	type args struct {
		id      uint
		amended Product
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Product
		wantErr bool
	}{
		{
			name: "Update existing product",
			fields: fields{
				products: []Product{{ID: 1, Name: "Old Product"}},
			},
			args: args{
				id:      1,
				amended: Product{Name: "Updated Product"},
			},
			want:    Product{ID: 1, Name: "Updated Product"},
			wantErr: false,
		},
		{
			name: "Update non-existing product",
			fields: fields{
				products: []Product{},
			},
			args: args{
				id:      1,
				amended: Product{Name: "Updated Product"},
			},
			want:    Product{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProductRepo{
				products: tt.fields.products,
			}
			got, err := p.Update(tt.args.id, tt.args.amended)
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
