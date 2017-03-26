package storage

import (
	"testing"
	"time"

	"github.com/dhconnelly/rtreego"
	"github.com/stretchr/testify/assert"
)

func TestDriverStorage(t *testing.T) {
	s := New(10)
	s.Set(&Driver{
		ID: 123,
		LastLocation: Location{
			Lat: 1,
			Lon: 1,
		},
		Expiration: time.Now().Add(15).UnixNano(),
	})
	d, err := s.Get(123)
	assert.NoError(t, err)
	assert.Equal(t, d.ID, 123)
	err = s.Delete(123)
	assert.NoError(t, err)
	d, err = s.Get(123)
	assert.Equal(t, err.Error(), "does not exist")

}
func BenchmarkNearest(b *testing.B) {
	s := New(20)
	for i := 0; i < 10000; i++ {
		s.Set(&Driver{
			ID: i,
			LastLocation: Location{
				Lat: float64(i),
				Lon: float64(i),
			},
		})
	}
	for i := 0; i < b.N; i++ {
		s.Nearest(rtreego.Point{123: 123}, 10)
	}
}

func TestNearest(t *testing.T) {
	s := New(10)
	s.Set(&Driver{
		ID: 123,
		LastLocation: Location{
			Lat: 42.875799,
			Lon: 74.588279,
		},
		Expiration: time.Now().Add(15).UnixNano(),
	})
	s.Set(&Driver{
		ID: 321,
		LastLocation: Location{
			Lat: 42.875508,
			Lon: 74.588107,
		},
		Expiration: time.Now().Add(15).UnixNano(),
	})
	s.Set(&Driver{
		ID: 666,
		LastLocation: Location{
			Lat: 42.876106,
			Lon: 74.588204,
		},
		Expiration: time.Now().Add(15).UnixNano(),
	})
	s.Set(&Driver{
		ID: 2319,
		LastLocation: Location{
			Lat: 42.874942,
			Lon: 74.585908,
		},
		Expiration: time.Now().Add(15).UnixNano(),
	})
	s.Set(&Driver{
		ID: 991,
		LastLocation: Location{
			Lat: 42.875744,
			Lon: 74.584503,
		},
		Expiration: time.Now().Add(15).UnixNano(),
	})

	drivers := s.Nearest(rtreego.Point{42.876420, 74.588332}, 3)
	assert.Equal(t, len(drivers), 3)
	assert.Equal(t, drivers[0].ID, 123)
	assert.Equal(t, drivers[1].ID, 321)
	assert.Equal(t, drivers[2].ID, 666)
}
