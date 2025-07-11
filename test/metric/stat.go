// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package metric

type Statistics string
type MetricValues []float64

const (
	AVERAGE                  Statistics = "Average"
	SAMPLE_COUNT             Statistics = "SampleCount"
	MINIMUM                  Statistics = "Minimum"
	MAXUMUM                  Statistics = "Maximum"
	SUM                      Statistics = "Sum"
	HighResolutionStatPeriod            = 10
	MinuteStatPeriod                    = 60
)
