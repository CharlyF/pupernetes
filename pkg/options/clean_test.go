// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCleanOptions(t *testing.T) {
	testCases := []struct {
		input     string
		expected  *Clean
		cliString string
	}{
		{
			"all",
			&Clean{
				common{
					true,
					false,
				},
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
			},
			"all",
		},
		{
			"none",
			&Clean{
				common{
					false,
					true,
				},
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
			},
			"",
		},
		{
			"none,all",
			&Clean{
				common{
					false,
					true,
				},
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
			},
			"",
		},
		{
			"all,none",
			&Clean{
				common{
					true,
					false,
				},
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
			},
			"all",
		},
		{
			"etcd",
			&Clean{
				common{
					false,
					false,
				},
				true,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
			},
			"etcd",
		},
		{
			"all,etcd",
			&Clean{
				common{
					true,
					false,
				},
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
			},
			"all",
		},
		{
			"etcd,binaries",
			&Clean{
				common{
					false,
					false,
				},
				true,
				true,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
			},
			"binaries,etcd",
		},
		{
			"etcd,binaries,secrets",
			&Clean{
				common{
					false,
					false,
				},
				true,
				true,
				false,
				false,
				true,
				false,
				false,
				false,
				false,
				false,
			},
			"binaries,etcd,secrets",
		},
		{
			"none,etcd",
			&Clean{
				common{
					false,
					true,
				},
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
			},
			"",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			actual := NewCleanOptions(testCase.input)
			assert.Equal(t, testCase.expected, actual)
			assert.Equal(t, testCase.cliString, actual.StringCLI())
		})
	}
}
