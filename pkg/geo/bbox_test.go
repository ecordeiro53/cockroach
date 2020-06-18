// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package geo

import (
	"fmt"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/geo/geopb"
	"github.com/stretchr/testify/require"
	"github.com/twpayne/go-geom"
)

func TestBoundingBoxFromGeom(t *testing.T) {
	testCases := []struct {
		g        geom.T
		expected *geopb.BoundingBox
	}{
		{geom.NewPointFlat(geom.XY, []float64{-15, -20}), &geopb.BoundingBox{MinX: -15, MaxX: -15, MinY: -20, MaxY: -20}},
		{geom.NewPointFlat(geom.XY, []float64{0, 0}), &geopb.BoundingBox{MinX: 0, MaxX: 0, MinY: 0, MaxY: 0}},
		{testGeomPoint, &geopb.BoundingBox{MinX: 1, MaxX: 1, MinY: 2, MaxY: 2}},
		{testGeomLineString, &geopb.BoundingBox{MinX: 1, MaxX: 2, MinY: 1, MaxY: 2}},
		{geom.NewLineStringFlat(geom.XY, []float64{-15, -20, -30, -40}), &geopb.BoundingBox{MinX: -30, MaxX: -15, MinY: -40, MaxY: -20}},
		{testGeomPolygon, &geopb.BoundingBox{MinX: 1, MaxX: 2, MinY: 1, MaxY: 2}},
		{testGeomMultiPoint, &geopb.BoundingBox{MinX: 1, MaxX: 2, MinY: 1, MaxY: 2}},
		{testGeomMultiLineString, &geopb.BoundingBox{MinX: 1, MaxX: 4, MinY: 1, MaxY: 4}},
		{testGeomMultiPolygon, &geopb.BoundingBox{MinX: 1, MaxX: 4, MinY: 1, MaxY: 4}},
		{testGeomGeometryCollection, &geopb.BoundingBox{MinX: 1, MaxX: 2, MinY: 1, MaxY: 2}},
		{emptyGeomPoint, nil},
		{emptyGeomLineString, nil},
		{emptyGeomPolygon, nil},
		{emptyGeomMultiPoint, nil},
		{emptyGeomMultiLineString, nil},
		{emptyGeomMultiPolygon, nil},
		{emptyGeomGeometryCollection, nil},
		{emptyGeomPointInGeometryCollection, &geopb.BoundingBox{MinX: 1, MaxX: 2, MinY: 1, MaxY: 2}},
		{emptyGeomObjectsInGeometryCollection, nil},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v", tc.g), func(t *testing.T) {
			bbox := boundingBoxFromGeom(tc.g)
			require.Equal(t, tc.expected, bbox)
		})
	}
}
